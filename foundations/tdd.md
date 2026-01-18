# SymPAL Technical Design Document

**Version:** 0.2.0
**Date:** 2026-01-18
**Status:** Draft (Synthesis complete, awaiting Checkpoint)
**Author:** Kael + Ryn (synthesis from Lead Dev interview)
**PRD Reference:** foundations/prd.md (v0.2.0)
**Privacy Reference:** foundations/privacy-innovations.md (v0.2.1)

---

## Problem Statement

**Product problem** (from PRD): Users want AI integration with their life data but don't trust LLM providers not to exploit it. Current tools force a choice: privacy OR utility.

**Technical challenge**: Build a privacy layer that enables AI-assisted workflows (email, calendar, contacts) while preventing any single entity from building a complete profile of the user.

**Key constraints**:
- P1: User data never leaves user control without explicit consent
- P3: No vendor lock-in — works with any LLM provider (V2+ full; V1 Claude + local)
- P5: Security by design, not bolted on
- Lead Dev has basic coding skills — architecture must be buildable with AI assistance
- Single user dogfooding — no scale concerns for V1

---

## Glossary

| Term | Definition |
|------|------------|
| **Knowledge Graph** | SQLite-based entity/relationship store tracking people, organizations, projects, and their connections. Used for context and semantic projection. |
| **LLM-as-Compiler** | Privacy tier where LLM generates executable code instead of receiving raw data. Query sent without data; code runs locally. Zero data exposure. |
| **Local LLM** | Ollama running Llama 3.2 3B locally. Used for content tasks (summarize, draft) where data never leaves device. |
| **Modular Monolith** | Single binary with internal module boundaries. Not microservices, but plugin-ready interfaces for V2. |
| **Privacy Tier** | The routing layer that classifies queries and directs them to Compiler (structured), Projection (reasoning), or Local (content). |
| **Projection** | See Semantic Projection. |
| **Rehydration** | Replacing typed placeholders in LLM responses with actual entities. Inverse of projection. Session-scoped mapping. |
| **Semantic Projection** | Privacy technique replacing real entities with typed placeholders (`[PERSON:colleague,senior]`). Preserves reasoning context, strips identity. |
| **Typed Placeholder** | Token like `[PERSON:colleague]` that preserves semantic type while hiding identity. Richer than simple `[REDACTED]`. |

---

## Goals & Non-Goals

### Technical Goals

| Goal | Success Criteria |
|------|------------------|
| Three-tier privacy routing | Queries routed to correct tier >80% of time |
| Zero exposure for structured queries | LLM-as-compiler executes locally, no data sent |
| Pattern-only exposure for reasoning | Semantic projection strips identity, preserves reasoning |
| Local processing for content | Ollama handles summarization, drafting |
| Acceptable latency | ≤1.5x Claude Code baseline |
| Dogfood-ready | Lead Dev uses daily within 2 milestones |

### Non-Goals (V1)

| Non-Goal | Rationale |
|----------|-----------|
| Full LLM agnosticism | V1 = Claude + local LLM; abstraction layer is V2 |
| Perfect anonymity | Practical obscurity is goal; nation-state attacks out of scope |
| Plugin architecture | Modular monolith; plugins are V2+ |
| Mobile/web interface | CLI only for V1 |
| Multi-user support | Single user dogfooding |
| Cloud storage/sync | Local only |

### Future Work (V2+)

These are explicitly deferred but inform V1 architecture decisions:

| Feature | Why Deferred | V1 Preparation |
|---------|--------------|----------------|
| Full LLM agnosticism | Complexity; V1 proves privacy architecture | Provider abstraction layer interface defined |
| Plugin system | YAGNI for single user | Modular monolith with clean boundaries |
| REPL mode | Subcommands sufficient for dogfooding | Session state design can extend |
| Mobile/web interface | CLI-first validates core value | API layer separable from CLI |
| P2P query mixing | Requires multiple users | Documented in privacy-innovations.md |
| Notepad/resurfacing | Personal knowledge management scope creep | Knowledge graph can support later |
| Multi-provider routing | One cloud LLM sufficient for V1 | Abstraction layer anticipates this |
| Advanced NER | Local small model sufficient for V1 | Entity extraction interface extensible |

---

## Architecture Overview

### Design Principles

1. **Local-first**: All user data stored locally; nothing persisted remotely
2. **Privacy by default**: No "fast mode" bypass; all queries route through privacy tier
3. **Modular monolith**: Single binary, internal module boundaries, plugin-ready interfaces for V2
4. **Fail safe**: If privacy tier fails, error — don't fall back to raw data exposure

### System Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLI (Go)                                 │
│  sympal today | sympal todo add | sympal calendar               │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Core Router                                 │
│  - Parse command                                                │
│  - Load context from knowledge graph                            │
│  - Route to appropriate handler                                 │
└─────────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┼───────────────┐
              ▼               ▼               ▼
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│  Integrations   │ │   Capabilities  │ │  Privacy Tier   │
│  - Gmail (R)    │ │   - Todo CRUD   │ │  - Classifier   │
│  - Calendar(RW) │ │   - Day Planner │ │  - Compiler     │
│  - Contacts (R) │ │                 │ │  - Projection   │
└─────────────────┘ └─────────────────┘ │  - Local LLM    │
        │                   │           └─────────────────┘
        │                   │                   │
        ▼                   ▼                   ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Data Layer (SQLite)                         │
│  ~/.sympal/data.db                                              │
│  - entities, relationships (knowledge graph)                    │
│  - todos, cached calendar/email/contacts                        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                     LLM Providers                                │
│  - Claude API (structured, reasoning via projection)            │
│  - Ollama/Llama 3.2 3B (content tasks, local)                   │
└─────────────────────────────────────────────────────────────────┘
```

### Component Summary

| Component | Responsibility | Key Interfaces |
|-----------|----------------|----------------|
| CLI | Parse commands, display output | stdin/stdout, Core Router |
| Core Router | Command dispatch, context loading | CLI, Capabilities, Privacy Tier |
| Integrations | OAuth, API calls, data sync | Google APIs, Data Layer |
| Capabilities | Business logic (todo, planning) | Core Router, Data Layer, Privacy Tier |
| Privacy Tier | Query classification, routing, projection | Capabilities, LLM Providers |
| Data Layer | SQLite access, knowledge graph | All components |
| LLM Providers | API abstraction | Privacy Tier |

---

## Key Technical Decisions

### Decision 1: Go + Deno Sandbox

**Context**: Need a language for main application + secure execution of LLM-generated code.

**Options considered**:

| Option | Pros | Cons |
|--------|------|------|
| Rust | Fast, safe, single binary | Steep learning curve |
| Python | LLM ecosystem, familiar | Distribution painful |
| Go | Simple syntax, single binary, LLM-friendly, great CLI ecosystem | Smaller ML ecosystem |
| Go + Deno | Go's strengths + Deno's built-in sandbox | Two runtimes |

**Decision**: Go for main application, Deno subprocess for sandboxed code execution.

**Consequences**:
- LLM-generated code is TypeScript/JavaScript (runs in Deno)
- Deno's deny-by-default permissions provide sandbox
- Single Go binary for distribution; Deno required as runtime dependency
- Good LLM code generation quality for both languages

---

### Decision 2: SQLite with Graph-like Schema

**Context**: Need to store todos, cached data, AND a knowledge graph for semantic projection.

**Options considered**:

| Option | Pros | Cons |
|--------|------|------|
| JSON files | Simplest | No query capability |
| SQLite (flat) | Simple, queryable | No graph relationships |
| SQLite (graph schema) | Queryable + relationships | Slightly more complex |
| Embedded graph DB | Native graph queries | Additional dependency |

**Decision**: SQLite with entities + relationships tables.

**Consequences**:
- Schema: `entities(id, name, type, metadata)`, `relationships(source_id, target_id, relation_type, metadata)`
- JOINs for relationship queries (sufficient for V1 scale)
- Upgrade path to dedicated graph DB if needed
- All data in `~/.sympal/data.db`

---

### Decision 3: Three-Tier Privacy Architecture

**Context**: Different query types have different privacy requirements.

**Options considered**:

| Option | Pros | Cons |
|--------|------|------|
| Always local LLM | Maximum privacy | 10-20% quality gap |
| Always cloud with PII redaction | Better quality | Context destroyed |
| Three-tier routing | Right tool for job | Routing complexity |

**Decision**: Three-tier routing (from privacy-innovations.md):

| Tier | Query Type | Data Exposure | Example |
|------|------------|---------------|---------|
| Compiler | Structured (filter, search) | Zero — LLM returns code | "Find meetings next week" |
| Projection | Reasoning (prioritize, plan) | Patterns only — semantic shadows | "What should I focus on today?" |
| Local | Content (summarize, draft) | None — runs on Ollama | "Summarize this email thread" |

**Consequences**:
- Query classifier needed (local, simple rules + small model)
- Compiler requires Deno sandbox
- Projection requires knowledge graph + entity taxonomy
- Local LLM requires Ollama + Llama 3.2 3B

---

### Decision 4: System Keychain for OAuth Tokens

**Context**: Google OAuth tokens must be stored securely.

**Options considered**:

| Option | Pros | Cons |
|--------|------|------|
| Plaintext file | Simplest | Insecure |
| Encrypted file | Secure | Key management burden |
| System keychain | OS-managed security | Platform-specific code |

**Decision**: System keychain (macOS Keychain, Linux secret-service).

**Consequences**:
- Platform-specific code for keychain access (Go libraries exist)
- No custom encryption to manage
- Follows industry standard (gh, gcloud, aws-cli)

---

### Decision 5: Subcommand CLI (Not REPL)

**Context**: Need a daily-driver interface.

**Options considered**:

| Option | Pros | Cons |
|--------|------|------|
| Subcommands | Simple, no state | No conversational flow |
| REPL | Conversational, context | State management, more complex |
| Both | Flexible | More code |

**Decision**: Subcommands only for V1 (`sympal today`, `sympal todo add`).

**Consequences**:
- Simpler implementation
- No session state to manage
- REPL can be V2 if wanted
- Matches user expectation (gh, git, docker)

---

## Data Model

### Data Classification

| Data Type | Sensitivity | Storage | Sent to Cloud? |
|-----------|-------------|---------|----------------|
| Todos | Low | SQLite | Via projection only |
| Calendar events | Medium | SQLite (cached) | Via projection only |
| Emails | High | SQLite (cached) | Via projection only |
| Contacts | High | SQLite (cached) | Never (context only) |
| Knowledge graph | Medium | SQLite | Never |
| OAuth tokens | Critical | System keychain | Never |
| Query logs | Medium | Log file | Never |

### Core Entities

```sql
-- Knowledge graph
CREATE TABLE entities (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,        -- PERSON, ORG, PROJECT, etc.
    modifiers TEXT,            -- JSON array: ["colleague", "senior"]
    metadata TEXT,             -- JSON blob
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE relationships (
    source_id TEXT REFERENCES entities(id),
    target_id TEXT REFERENCES entities(id),
    relation_type TEXT NOT NULL,  -- works_with, reports_to, etc.
    metadata TEXT,
    PRIMARY KEY (source_id, target_id, relation_type)
);

-- Domain data
CREATE TABLE todos (
    id TEXT PRIMARY KEY,
    content TEXT NOT NULL,
    due_date DATE,
    priority TEXT CHECK(priority IN ('high', 'medium', 'low')),
    status TEXT CHECK(status IN ('not_started', 'done')),
    source_type TEXT,          -- 'manual', 'email', 'calendar'
    source_id TEXT,            -- reference to source
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Cached integration data
CREATE TABLE calendar_events (...);
CREATE TABLE emails (...);
CREATE TABLE contacts (...);
```

### Data Flows

```
Google APIs → OAuth → Local Cache (SQLite)
                           ↓
                    Knowledge Graph (entity extraction)
                           ↓
              User Query → Privacy Tier → LLM (projected data only)
                                            ↓
                                    Response → Rehydration → User
```

### Data Ownership (P11 Reversibility)

| Operation | Supported | Implementation |
|-----------|-----------|----------------|
| Export all data | Yes | `sympal export` → JSON dump of all tables |
| Delete all data | Yes | `rm -rf ~/.sympal/` or `sympal reset` |
| Migrate to other tool | Yes | Export provides standard JSON |

---

## Security & Privacy

### Threat Model

**Assets to protect**:
- Personal data (email, calendar, contacts)
- Query patterns (what user asks about)
- Identity (who the user is)

**Threat actors**:

| Actor | Motivation | In Scope? |
|-------|------------|-----------|
| LLM Provider (passive) | Profile building, monetization | Yes |
| LLM Provider (active) | Deanonymization | Yes |
| Nation-state | Targeted surveillance | No |
| Local attacker | Device access | No |

**Trust boundaries**:
- Local device is trusted
- Network is untrusted (but TLS assumed valid)
- LLM providers are honest-but-curious

### Security Requirements

| Requirement | Implementation | Principle |
|-------------|----------------|-----------|
| No raw PII to cloud | Privacy tier enforced, no bypass | P1 |
| Secure token storage | System keychain | P5 |
| Safe code execution | Deno sandbox with deny-by-default | P5 |
| Audit trail | Query log shows what was sent | P12 |
| User can inspect | `sympal log` shows recent queries | P4, P10 |

### Privacy Design

| Property | Implementation |
|----------|----------------|
| Data minimization | Only cache what's needed; sync minimal scope |
| User control | Disable integrations anytime; see what's sent |
| No dark patterns | Explicit prompts for any data sharing |
| Transparency | Full query log locally; open source |

### Authentication & Authorization

- **Google OAuth**: Read (Gmail, Contacts), Read+Write (Calendar)
- **LLM APIs**: API keys stored in config (user provides)
- **CLI**: No auth (OS login is boundary)

---

## Failure Modes & Mitigations

| Failure Mode | Trigger | Severity | Mitigation |
|--------------|---------|----------|------------|
| Query misclassification | Ambiguous query | Medium | Default to most private tier (local) |
| Sandbox escape | Malicious LLM output | Critical | Deno deny-by-default; if validation fails, don't execute |
| Entity leak in projection | NER miss | High | User review of first mentions; log for audit |
| Rehydration wrong entity | Ambiguous placeholder | Medium | Session-scoped mappings; ask for clarification |
| OAuth token expired | Token not refreshed | Low | Background refresh; graceful error with re-auth prompt |
| Local LLM unavailable | Ollama not running | Low | Error message with setup instructions |
| Cloud LLM unavailable | API down | Medium | Retry with backoff; fallback to local for more queries |
| Knowledge graph empty | New user | Low | Bootstrap prompts; infer from first queries |

### Graceful Degradation

| Scenario | Degraded Behavior |
|----------|-------------------|
| LLM unavailable | Structured queries still work (compiler runs local); reasoning/content blocked |
| Network offline | Local data accessible; integrations show cached data with staleness warning |
| Ollama not running | Content tasks fail with "start Ollama" message; other tiers work |
| SQLite locked | Retry with backoff; error if persists |

---

## Implementation Plan

### Phase 1: Foundation (M1)

- [ ] Go project scaffolding (cobra CLI, config loading)
- [ ] SQLite setup with schema
- [ ] Basic todo CRUD (`sympal todo add/list/done/delete`)
- [ ] Config file handling (`~/.sympal/config.yaml`)
- [ ] Logging infrastructure (single logger to `~/.sympal/sympal.log`)

**Enables**: Working CLI, data persistence, foundation for everything else

### Phase 2: Calendar Integration (M2)

- [ ] Google OAuth flow (system keychain storage)
- [ ] Calendar API integration (read events)
- [ ] Calendar write (create events)
- [ ] `sympal today` command (todos + calendar)
- [ ] Basic day planning (manual, no LLM yet)

**Depends on**: Phase 1
**Enables**: Integrated view, OAuth infrastructure reusable for Gmail

### Phase 3: Privacy Tier (M3) — Core Innovation

**Phase 3A: LLM as Compiler**
- [ ] Query classifier (structured vs reasoning vs content)
- [ ] Schema description generator
- [ ] Claude API integration
- [ ] Deno sandbox setup
- [ ] Code generation → validation → execution pipeline
- [ ] Test with todo + calendar queries

**Phase 3B: Semantic Projection**
- [ ] Entity extraction (local NER or small model)
- [ ] Knowledge graph population
- [ ] Projection function (real → typed placeholders)
- [ ] Rehydration function (response → real entities)
- [ ] Entity taxonomy implementation
- [ ] Test with reasoning queries

**Phase 3C: Local LLM**
- [ ] Ollama integration
- [ ] Content task routing
- [ ] Quality logging (for comparison)

**Depends on**: Phase 2
**Enables**: Privacy-preserving LLM integration — the core value prop

### Phase 4: Email Integration (M4)

- [ ] Gmail API integration (read only)
- [ ] Email-to-todo detection (using privacy tier)
- [ ] Email context enrichment

**Depends on**: Phase 3 (privacy must be proven first)

### Phase 5: Contacts Integration (M5)

- [ ] Contacts API integration (read only)
- [ ] Contact context in day planning
- [ ] Knowledge graph enrichment from contacts

**Depends on**: Phase 3

### Dependency Graph

```
Phase 1 (Foundation)
    │
    ▼
Phase 2 (Calendar)
    │
    ▼
Phase 3A (Compiler) ──→ Phase 3B (Projection) ──→ Phase 3C (Local LLM)
    │                           │                         │
    └───────────────────────────┼─────────────────────────┘
                                ▼
                    Phase 4 (Email) ──→ Phase 5 (Contacts)
```

---

## Testing Strategy

### Test Levels

| Level | Coverage Target | Automation |
|-------|-----------------|------------|
| Unit | Security boundaries, privacy tier, knowledge graph | Yes |
| Integration | Google API smoke tests | Yes (limited) |
| E2E | Dogfooding | Manual |
| Security | Sandbox validation | Yes |

### Critical Test Cases

| Test Case | Verifies | Priority |
|-----------|----------|----------|
| Sandbox blocks imports | Deno denies unauthorized access | P0 |
| Sandbox blocks network | Generated code can't exfiltrate | P0 |
| Sandbox blocks file I/O | Generated code can't access filesystem | P0 |
| Projection strips PII | No real names/emails in output | P0 |
| Rehydration correct | Placeholders map back correctly | P0 |
| Query classifier routes correctly | Sample queries hit right tier | P1 |
| OAuth token refresh | Expired token triggers refresh | P1 |
| Todo CRUD works | Basic operations succeed | P1 |

### Test Data Strategy

- **Unit tests**: Synthetic data, fixtures
- **Integration tests**: Test Google account (or mocks)
- **Dogfooding**: Real data (that's the point)
- **Regression**: Add test when bug found and fixed

---

## Deployment

### Build & Distribution

| Aspect | V1 Approach | Rationale |
|--------|-------------|-----------|
| Build system | `go build` | Simple, single binary output |
| Distribution | Manual install (download binary) | Dogfooding only; package managers are V2 |
| Target platforms | macOS (primary), Linux (secondary) | Lead Dev's machines |
| Dependencies | Deno (user installs), Ollama (user installs) | Document setup; don't bundle |

### Installation Steps (V1)

```bash
# 1. Download binary (or build from source)
go build -o sympal ./cmd/sympal

# 2. Move to PATH
mv sympal /usr/local/bin/

# 3. Install dependencies
brew install deno     # or appropriate method
brew install ollama   # or appropriate method

# 4. Run setup
sympal init           # Creates ~/.sympal/, prompts for API keys
```

### Configuration

| File | Purpose |
|------|---------|
| `~/.sympal/config.yaml` | API keys, integration settings, preferences |
| `~/.sympal/data.db` | SQLite database (todos, cache, knowledge graph) |
| `~/.sympal/sympal.log` | Log file |
| System keychain | OAuth tokens |

---

## Rollback & Recovery

### Data Backup

User responsibility. Simple approach:

```bash
# Backup
cp -r ~/.sympal ~/.sympal.backup.$(date +%Y%m%d)

# Restore
rm -rf ~/.sympal && cp -r ~/.sympal.backup.YYYYMMDD ~/.sympal
```

### Version Rollback

Since distribution is manual binary download:

1. Keep previous binary: `mv sympal sympal.prev`
2. If new version fails: `mv sympal.prev sympal`

No migrations in V1 — schema changes are breaking (acceptable for dogfooding).

### Recovery Scenarios

| Scenario | Recovery |
|----------|----------|
| Corrupted database | Delete `data.db`, re-sync integrations, re-add todos |
| Lost OAuth tokens | `sympal auth refresh` — re-authenticate with Google |
| Bad config | Delete `config.yaml`, run `sympal init` |
| Complete reset | `rm -rf ~/.sympal && sympal init` |

---

## Risks & Open Questions

### Technical Risks

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Query classification unreliable | Medium | Medium | Start with rules; add ML if needed; default to most private |
| Deno adds too much latency | Low | Medium | Benchmark early; subprocess reuse |
| Projection quality insufficient | Medium | High | Measure against baseline; iterate taxonomy |
| Ollama too slow on M1 8GB | Medium | Low | Upgrade path to M4 Pro; use smallest viable model |
| Knowledge graph bootstrap painful | Medium | Medium | Start minimal; grow organically |

### Open Questions

1. **Query classifier implementation**: Rules-based vs small model vs hybrid? Start with rules, measure, iterate.

2. **Knowledge graph bootstrapping**: Auto-infer from data vs user prompts vs hybrid? Defer to Phase 3B.

3. **Deno installation**: Require user to install, or bundle? Research Go→Deno embedding options.

4. **Ollama model selection**: Llama 3.2 3B is starting point; may need to test others for quality floor.

### Dependencies

| Dependency | Type | Status | Blocker? |
|------------|------|--------|----------|
| Google OAuth credentials | External | Not set up | Yes — need before Phase 2 |
| Claude API key | External | Have | No |
| Ollama installation | External | Need to document | Yes — need before Phase 3C |
| Deno installation | External | Need to research | Yes — need before Phase 3A |

---

## Assumptions

These assumptions underpin the design. If any prove false, the affected areas need revisiting.

| Assumption | Status | Impact if Wrong | Validation |
|------------|--------|-----------------|------------|
| Local LLMs can type entities accurately | Unvalidated | Projection quality degrades | Benchmark on personal data samples |
| Type vocabulary can be rich enough for reasoning yet generic enough for privacy | Unvalidated | Privacy/utility tradeoff shifts | User study during dogfooding |
| 60-70% of queries are structured (filterable) | Unvalidated | Compiler tier underutilized | Measure query distribution post-launch |
| Deno subprocess sandboxing is secure | Assumed (Deno's design) | Critical security boundary | Follow Deno security advisories |
| Llama 3.2 3B quality sufficient for content tasks | Unvalidated | May need larger model or accept quality gap | Compare outputs to Claude baseline |
| M1 8GB RAM can run Ollama + SymPAL concurrently | Untested | Latency/reliability issues | Test on target hardware |
| Go + Deno subprocess IPC performs adequately | Untested | Latency exceeds 1.5x target | Benchmark early in Phase 3A |
| System keychain libraries work cross-platform | Assumed (libraries exist) | Custom encryption needed | Test on Linux secret-service |
| Lead Dev can build with AI assistance | Assumed | Progress stalls | Ongoing — session-by-session feedback |

---

## PRINCIPLES.md Alignment

### Hard Constraints Implementation

| Constraint | Architectural Support |
|------------|----------------------|
| P1: Privacy & Data Sovereignty | Three-tier privacy architecture; all data local; user controls what's sent |
| P2: Open Source | All code open; no proprietary dependencies |
| P3: LLM-Agnosticism | Provider abstraction layer; V1 = Claude + Ollama; V2 = full abstraction |
| P4: Honesty | Query log shows exactly what was sent; no hidden collection |
| P5: Security Through Design | Deno sandbox, system keychain, threat model from day 1 |

### Tensions Considered

| Tension | Relevant? | Navigation |
|---------|-----------|------------|
| Present vs. Future | Yes | Ship V1 with known limitations; plugin-ready interfaces |
| Safety vs. Capability | Yes | Privacy tier is bounded; no bypass option |
| Dogfooding vs. Adoption | Yes | Optimizing for Lead Dev only; broader use is V2+ |
| Innovation vs. Precaution | Yes | Novel privacy approach; accepting calculated risk with measurement |
| Efficiency vs. Meaning | No | Not automating meaningful human activities |

---

## Success Metrics

### Primary Metric (P17 Dogfooding)

**Lead Dev uses SymPAL daily** — measured by:
- `sympal today` run each morning
- Todos managed through SymPAL, not external apps
- Calendar integration actively used

### Technical Metrics

| Metric | Target | Measurement Method |
|--------|--------|-------------------|
| Query routing accuracy | >80% correct tier | Log analysis: did query hit appropriate tier? |
| Latency (simple queries) | <5s end-to-end | Instrumented timing |
| Latency (complex reasoning) | <15s end-to-end | Instrumented timing |
| Latency vs baseline | ≤1.5x raw Claude API | Comparison benchmark |
| Code execution success | >90% valid output | Sandbox execution logs |
| Projection quality | Subjective 1-5 scale | Weekly sample review |

### Quality Protocol

1. **Logging**: All queries, tier decisions, execution results logged
2. **Weekly review**: Sample 10-20 queries, rate quality, check routing
3. **Baseline comparison**: Monthly comparison to raw Claude API for same queries
4. **Iteration**: If metrics miss targets, diagnose and adjust

---

## References

### Project Documents

| Document | Location | Relevance |
|----------|----------|-----------|
| Product Requirements | `foundations/prd.md` | What we're building and why |
| Binding Principles | `PRINCIPLES.md` | Hard constraints and navigation guidance |
| Privacy Architecture | `foundations/privacy-innovations.md` | Three-tier design, threat model, entity taxonomy |
| Privacy Research | `foundations/privacy-research.md` | Survey of existing approaches |
| Extraction Notes | `foundations/working/tdd-extraction-notes.md` | Audit trail of technical decisions |

### External Dependencies

| Dependency | Documentation | Version |
|------------|---------------|---------|
| Go | https://go.dev/doc/ | 1.21+ |
| Deno | https://deno.land/manual | 1.40+ |
| Ollama | https://ollama.ai/docs | Latest |
| Llama 3.2 3B | Ollama model registry | llama3.2:3b |
| SQLite | https://sqlite.org/docs.html | 3.x |
| Google Calendar API | https://developers.google.com/calendar | v3 |
| Gmail API | https://developers.google.com/gmail | v1 |
| Google Contacts API | https://developers.google.com/people | v1 |

### Go Libraries (Planned)

| Library | Purpose |
|---------|---------|
| spf13/cobra | CLI framework |
| mattn/go-sqlite3 | SQLite driver |
| golang.org/x/oauth2 | OAuth 2.0 |
| zalando/go-keyring | System keychain access |

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 0.2.0 | 2026-01-18 | Added: Glossary, Assumptions, Future Work, Deployment, Rollback, Success Metrics, References |
| 0.1.0 | 2026-01-18 | Initial synthesis from extraction interview |
