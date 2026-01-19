# SymPAL Technical Design Document

**Version:** 1.0.0
**Date:** 2026-01-19
**Status:** Draft (Fresh synthesis, awaiting Checkpoint)
**Author:** Kael + Ryn (synthesis from Lead Dev interview + delta)
**PRD Reference:** foundations/prd.md (v0.3.0)
**Privacy Reference:** foundations/privacy-innovations.md (v3.0.0)

---

## Problem Statement

**Product problem** (from PRD): Users want AI integration with their life data but don't trust LLM providers not to exploit it. Current tools force a choice: privacy OR utility.

**Technical challenge**: Build a privacy layer that enables AI-assisted workflows (calendar, todos) while preventing any single entity from building a complete profile of the user.

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
| **DSL Compilation** | Privacy tier where cloud LLM generates SymQL code instead of receiving data. Query sent without data; code runs locally. Zero data exposure. |
| **Ephemeral Slots** | Privacy technique using single-use random placeholders (`[E1]`, `[E2]`) with a session-scoped legend. Prevents cross-query entity correlation. |
| **Knowledge Graph** | SQLite-based entity/relationship store tracking people, organizations, projects, and their connections. |
| **Legend** | Session-scoped mapping from placeholder (`[E1]`) to entity info. Minimum viable detail for task. Destroyed after rehydration. |
| **Local LLM** | Ollama running Llama 3.2 3B locally. Used for content tasks (summarize, draft) where data never leaves device. |
| **Modular Monolith** | Single binary with internal module boundaries. Not microservices, but plugin-ready interfaces for V2. |
| **Privacy Tier** | The routing layer that classifies queries and directs them to DSL (structured), Ephemeral Slots (reasoning), or Local (content). |
| **Query Classifier** | Keyword cascade algorithm that routes queries to appropriate privacy tier. Deterministic, zero latency. |
| **Rehydration** | Replacing placeholders in LLM responses with actual entities using the legend. Session-scoped, legend destroyed after. |
| **SymQL** | Domain-specific query language for structured data operations. Compiled from natural language by cloud LLM, executed locally. |

---

## Goals & Non-Goals

### Technical Goals

| Goal | Success Criteria |
|------|------------------|
| Three-tier privacy routing | Queries routed to correct tier >80% of time |
| Zero exposure for structured queries | DSL compilation executes locally, no data sent |
| Pattern-only exposure for reasoning | Ephemeral Slots strip identity, preserve reasoning context |
| Local processing for content | Ollama handles summarization, drafting |
| Acceptable latency | ≤1.5x Claude Code baseline |
| Dogfood-ready | Lead Dev uses daily after M5 |

### Non-Goals (V1)

| Non-Goal | Rationale |
|----------|-----------|
| Full LLM agnosticism | V1 = Claude + local LLM; abstraction layer is V2 |
| Perfect anonymity | Practical obscurity is goal; nation-state attacks out of scope |
| PaaP (Prompt-as-Program) | Deferred to V1.5; reduces risk surface for V1 |
| Plugin architecture | Modular monolith; plugins are V2+ |
| Mobile/web interface | CLI only for V1 |
| Multi-user support | Single user dogfooding |
| Email/Contacts integration | Deferred until privacy proven with Calendar |

### V1 Coverage

With PaaP deferred, V1 achieves:
- **~40-50% zero exposure** (DSL Compilation for structured queries)
- **~40-50% pattern-only exposure** (Ephemeral Slots for reasoning)
- **~10% local-only** (content generation)

---

## Architecture Overview

### Design Principles

1. **Local-first**: All user data stored locally; nothing persisted remotely
2. **Privacy by default**: No "fast mode" bypass; all queries route through privacy tier
3. **Default to private**: If classification uncertain, route to Local LLM (most private)
4. **Modular monolith**: Single binary, internal module boundaries, plugin-ready interfaces for V2
5. **Fail safe**: If privacy tier fails, error — don't fall back to raw data exposure

### System Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLI (Go)                                │
│  sympal today | sympal todo add | sympal calendar               │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Core Router                                │
│  - Parse command                                                │
│  - Load context from knowledge graph                            │
│  - Route to appropriate handler                                 │
└─────────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┼───────────────┐
              ▼               ▼               ▼
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────────────┐
│  Integrations   │ │   Capabilities  │ │      Privacy Tier       │
│  - Calendar(RW) │ │   - Todo CRUD   │ │  ┌─────────────────┐    │
│                 │ │   - Day Planner │ │  │ Query Classifier│    │
│  (Gmail, Cont-  │ │                 │ │  │ (keyword cascade)│   │
│   acts = V1.5+) │ │                 │ │  └────────┬────────┘    │
└─────────────────┘ └─────────────────┘ │           │             │
        │                   │           │     ┌─────┴─────┐       │
        │                   │           │     ▼     ▼     ▼       │
        │                   │           │  ┌───┐ ┌───┐ ┌─────┐   │
        │                   │           │  │DSL│ │E.S│ │Local│   │
        │                   │           │  └───┘ └───┘ └─────┘   │
        │                   │           └─────────────────────────┘
        ▼                   ▼                       │
┌─────────────────────────────────────────────────────────────────┐
│                     Data Layer (SQLite)                         │
│  ~/.sympal/data.db                                              │
│  - entities, relationships (knowledge graph)                    │
│  - todos, cached calendar                                       │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                     LLM Providers                               │
│  - Claude API (DSL compilation, Ephemeral Slots reasoning)      │
│  - Ollama/Llama 3.2 3B (content tasks, local)                   │
└─────────────────────────────────────────────────────────────────┘

Legend: DSL = DSL Compilation, E.S = Ephemeral Slots, Local = Local LLM
```

### Component Summary

| Component | Responsibility | Key Interfaces |
|-----------|----------------|----------------|
| CLI | Parse commands, display output | stdin/stdout, Core Router |
| Core Router | Command dispatch, context loading | CLI, Capabilities, Privacy Tier |
| Integrations | OAuth, API calls, data sync | Google APIs, Data Layer |
| Capabilities | Business logic (todo, planning) | Core Router, Data Layer, Privacy Tier |
| Privacy Tier | Query classification, routing, Ephemeral Slots. **On uncertain classification, defaults to Local.** | Capabilities, LLM Providers |
| Data Layer | SQLite access, knowledge graph | All components |
| LLM Providers | API abstraction | Privacy Tier |

---

## Query Classifier

The Query Classifier is a deterministic keyword cascade that routes queries to the appropriate privacy tier. Zero latency, fully debuggable.

### Classification Algorithm

```
CLASSIFY(query):
    1. CHECK STRUCTURED PATTERNS (→ DSL Compilation)
       Keywords: "how many", "count", "list", "filter", "search", "find all"
       Confidence: HIGH if keyword match + no content verbs

    2. CHECK HYBRID PATTERNS (→ DSL then Ephemeral Slots)
       Detection: Structured keyword + reasoning verb ("prioritize the meetings")
       Process: Execute DSL first, pass results to Ephemeral Slots

    3. CHECK REASONING PATTERNS (→ Ephemeral Slots)
       Keywords: "should I", "prioritize", "advise", "recommend", "what's important"
       Confidence: HIGH if reasoning keyword present

    4. CHECK CONTENT PATTERNS (→ Local LLM)
       Keywords: "draft", "write", "compose", "summarize", "rewrite"
       Confidence: HIGH if content verb present

    5. DEFAULT (→ Local LLM with quality warning)
       Log: "Uncertain classification, defaulting to most private"
```

### Classification Categories

| Category | Route To | Example Queries |
|----------|----------|-----------------|
| STRUCTURED | DSL Compilation | "How many meetings this week?", "List high-priority todos" |
| HYBRID | DSL → Ephemeral Slots | "Prioritize my meetings today" (filter then reason) |
| REASONING | Ephemeral Slots | "Should I reschedule?", "What should I focus on?" |
| CONTENT | Local LLM | "Draft a meeting agenda", "Summarize this" |
| UNCERTAIN | Local LLM (default) | Ambiguous queries |

### Confidence Thresholds

| Confidence | Behavior |
|------------|----------|
| HIGH (keyword match) | Route directly |
| MEDIUM (partial match) | Route with logging |
| LOW (no match) | Default to Local LLM |

---

## Privacy Tier Implementation

### Tier 1: DSL Compilation (Zero Exposure)

**How it works**:
1. Query + schema description sent to Claude (no data)
2. Claude generates SymQL code
3. Code validated and executed in Deno sandbox
4. Results returned to user

**SymQL Grammar** (subset):

```
query       := SELECT fields FROM source [WHERE conditions] [ORDER BY field] [LIMIT n]
fields      := "*" | field ("," field)*
source      := "todos" | "calendar" | "entities"
conditions  := condition (("AND" | "OR") condition)*
condition   := field operator value
operator    := "=" | "!=" | ">" | "<" | ">=" | "<=" | "CONTAINS" | "IN"
```

**Example**:
- Input: "How many meetings do I have next week?"
- Generated: `SELECT COUNT(*) FROM calendar WHERE start_date >= '2026-01-20' AND start_date < '2026-01-27' AND type = 'meeting'`

**Sandbox Specification** (Deno):

```bash
deno run --no-prompt --no-npm --no-remote [script.ts]
```

| Flag | Effect |
|------|--------|
| (no `--allow-*` flags) | All permissions denied by default |
| `--no-prompt` | Don't ask user for permissions interactively |
| `--no-npm` | Disallow npm package imports |
| `--no-remote` | Disallow fetching remote modules |

Data passed via stdin/stdout. 5-second timeout enforced by Go.

### Tier 2: Ephemeral Slots (Pattern-Only Exposure)

**How it works**:
1. Extract entities from query (NER)
2. Replace entities with single-use placeholders (`[E1]`, `[E2]`)
3. Build minimum-viable legend based on task type
4. Send anonymized query + legend to Claude
5. Receive response with placeholders
6. Rehydrate: replace placeholders with real entities
7. Destroy legend

**Legend Construction** (Task-Based Defaults):

| Task Type | Legend Detail | Example |
|-----------|---------------|---------|
| COUNTING | Bare minimum | `[E1] is a person` |
| SCHEDULING | Contact level | `[E1] is a contact` |
| SUMMARIZING | Role level | `[E1] is a colleague` |
| REASONING | Relationship + context | `[E1] is my manager, high-stakes relationship` |
| DRAFTING | Relationship + tone | `[E1] is my client, formal tone preferred` |

Escalation only if quality degrades (rare with good defaults).

**NER Confidence Handling**:
- HIGH (>0.8): Auto-accept
- MEDIUM (0.5-0.8): Accept, log for review
- LOW (<0.5): Skip review for known entities; prompt for truly new entities

**Rehydration Requirements**:
- LLM response must use exact placeholder format: `[E1]`, `[E2]`
- Structured output (JSON) preferred for reliable parsing
- Fallback cascade: retry prompt → local LLM → partial result with warning

### Tier 3: Local LLM (Zero Exposure)

**How it works**:
1. Full query + data sent to Ollama (never leaves device)
2. Ollama processes with Llama 3.2 3B
3. Response returned to user

**Use cases**: Content generation, summarization, drafting — tasks where full context needed.

**Quality tradeoff**: ~10-20% quality gap vs Claude, but zero privacy risk.

---

## Cost of Privacy Model

| Path | Latency | API Calls | Est. Cost | Quality | Privacy |
|------|---------|-----------|-----------|---------|---------|
| **Heuristic/Template** | ~50ms | 0 | $0 | Low (rigid) | Zero exposure |
| **DSL Compilation** | 1-2s | 1 cloud | ~$0.002 | High (logic) | Zero exposure |
| **Ephemeral Slots** | 1-3s | 1 cloud | ~$0.003 | High | Pattern-only |
| **Local LLM only** | 2-5s | 0 | $0 | Medium-Low | Zero exposure |

V1 target: 60-70% of queries hit DSL or Ephemeral Slots (cloud, high quality, privacy-preserved).

---

## User-Facing Failure Model

**Philosophy**: Fail visibly, fail safely, offer alternatives.

| Failure Type | User Sees | System Does |
|--------------|-----------|-------------|
| DSL compilation fails | "Couldn't understand that query. Try rephrasing, or I can use local AI (slower)." | Log error, offer Local fallback |
| NER misses entity | "I found [entity] — is this correct?" | Prompt for confirmation (new entities only) |
| Rehydration fails | "Got a response but couldn't match it back to your data. Here's the raw result: [...]" | Show anonymized response, log for debugging |
| Sandbox timeout | "Query took too long. Try a simpler question." | Kill process, log |
| Local LLM unavailable | "Local AI not running. Start Ollama, or I can try cloud (less private)." | Offer cloud fallback with explicit consent |

**Design principles**:
1. Never silently fail
2. Never expose raw data on failure
3. Always offer a path forward

---

## Data Model

### Data Classification

| Data Type | Sensitivity | Storage | Sent to Cloud? |
|-----------|-------------|---------|----------------|
| Todos | Low | SQLite | Via Ephemeral Slots only |
| Calendar events | Medium | SQLite (cached) | Via Ephemeral Slots only |
| Knowledge graph | Medium | SQLite | Never |
| OAuth tokens | Critical | System keychain | Never |
| Query logs | Medium | Log file | Never |

### Core Schema

```sql
-- Knowledge graph
CREATE TABLE entities (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,        -- PERSON, ORG, PROJECT, etc.
    modifiers TEXT,            -- JSON array: ["colleague", "senior"]
    metadata TEXT,             -- JSON blob
    first_seen TIMESTAMP,
    last_seen TIMESTAMP
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
    source_type TEXT,          -- 'manual', 'calendar'
    source_id TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE calendar_events (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    attendees TEXT,            -- JSON array of entity IDs
    metadata TEXT,
    synced_at TIMESTAMP
);
```

### Data Ownership (P11 Reversibility)

| Operation | Supported | Implementation |
|-----------|-----------|----------------|
| Export all data | Yes | `sympal export` → JSON dump |
| Delete all data | Yes | `rm -rf ~/.sympal/` or `sympal reset` |
| Migrate to other tool | Yes | Export provides standard JSON |

---

## Security & Privacy

### Threat Model

**Assets to protect**:
- Personal data (calendar, todos, entities)
- Query patterns (what user asks about)
- Identity (who the user is)

**Threat actors**:

| Actor | Motivation | In Scope? | Mitigation |
|-------|------------|-----------|------------|
| LLM Provider (passive) | Profile building | Yes | Ephemeral Slots — single-use placeholders defeat entity tracking |
| LLM Provider (active) | Deanonymization via correlation | Yes | Token rotation (daily), query batching, timing noise |
| Nation-state | Targeted surveillance | No | Out of scope for V1 |
| Local attacker | Device access | No | Out of scope for V1 |

### Correlation Mitigations

| Mitigation | Implementation | Prevents | Doesn't Prevent |
|------------|----------------|----------|-----------------|
| Ephemeral placeholders | New IDs each query | Cross-query entity tracking | — |
| Token rotation | Placeholder IDs reset daily | Cross-session tracking | Pattern analysis |
| Query batching | Batch 2-5 queries when possible | Individual query attribution | Batch-level patterns |
| Timing noise | 100-500ms random delay | Timing correlation | Content correlation |

### Security Requirements

| Requirement | Implementation | Principle |
|-------------|----------------|-----------|
| No raw PII to cloud | Privacy tier enforced, no bypass | P1 |
| Secure token storage | System keychain | P5 |
| Safe code execution | Deno sandbox, deny-by-default | P5 |
| Audit trail | Query log shows what was sent | P12 |
| User can inspect | `sympal log` shows recent queries | P4, P10 |

---

## Explicitly Accepted Risks

### Risk 1: NER is the Achilles' Heel (CRITICAL)

**The risk**: If NER misses an entity, real data goes to cloud LLM.

**Why we accept it**:
- No perfect NER exists
- User review catches misses for new entities
- Known entities auto-recognized (skip review)
- Alternative (no cloud) sacrifices too much quality

**Mitigations**:
- Conservative extraction (prefer false positives)
- User confirmation for new/uncertain entities
- Audit log for post-hoc review

### Risk 2: Legend Leakage if Escalation Always Required

**The risk**: If task-based defaults always need escalation, legends become too detailed.

**Why we accept it**:
- Task-based defaults should minimize escalation
- Dogfooding will reveal if defaults are wrong
- Can tighten defaults based on data

**Mitigations**:
- Start with task-appropriate defaults
- Track escalation frequency
- Iterate taxonomy based on usage

---

## Implementation Plan

### Phase 1: Foundation (M1)

- [ ] Go project scaffolding (Cobra CLI, config loading)
- [ ] SQLite setup with schema
- [ ] Basic todo CRUD (`sympal todo add/list/done/delete`)
- [ ] Config file handling (`~/.sympal/config.yaml`)
- [ ] Logging infrastructure

**Gate**: Todo CRUD works end-to-end

### Phase 2: Calendar Integration (M2)

- [ ] Google OAuth flow (system keychain storage)
- [ ] Calendar API integration (read events)
- [ ] `sympal today` command (todos + calendar)
- [ ] Basic day view (no LLM yet)

**Gate**: Can view today's calendar and todos together

### Phase 3: DSL Compilation (M3)

- [ ] Query Classifier (keyword cascade)
- [ ] Schema description generator
- [ ] Claude API integration
- [ ] SymQL code generation
- [ ] Deno sandbox setup
- [ ] Code validation → execution pipeline
- [ ] Test with todo + calendar queries

**Gate**: >90% of structured queries return correct results

### Phase 4: Ephemeral Slots (M4)

- [ ] Entity extraction (NER)
- [ ] Ephemeral placeholder generation
- [ ] Legend construction (task-based defaults)
- [ ] Projection function (real → placeholders)
- [ ] Rehydration function (response → real)
- [ ] Test with reasoning queries

**Gate**: >95% rehydration accuracy

### Phase 5: Local LLM + Integration (M5)

- [ ] Ollama integration
- [ ] Content task routing
- [ ] End-to-end privacy tier
- [ ] Quality logging
- [ ] Daily use begins

**Gate**: Lead Dev uses daily (subjective)

### Dependency Graph

```
M1 (Foundation)
    │
    ▼
M2 (Calendar)
    │
    ▼
M3 (DSL Compilation)
    │
    ▼
M4 (Ephemeral Slots)
    │
    ▼
M5 (Local LLM + Integration)
```

---

## Testing Strategy

### Test Levels

| Level | Coverage Target | Automation |
|-------|-----------------|------------|
| Unit | Privacy tier, sandbox, knowledge graph | Yes |
| Integration | Google API smoke tests | Yes (limited) |
| E2E | Dogfooding | Manual |
| Security | Sandbox validation | Yes |

### Critical Test Cases

| Test Case | Verifies | Priority |
|-----------|----------|----------|
| Sandbox blocks network | Generated code can't exfiltrate | P0 |
| Sandbox blocks file I/O | Generated code can't access filesystem | P0 |
| Placeholder generation unique | No reuse within session | P0 |
| Rehydration correct | Placeholders map back correctly | P0 |
| Legend destroyed after use | No persistence across queries | P0 |
| Query classifier routes correctly | Sample queries hit right tier | P1 |
| NER extracts known entities | Previously seen entities recognized | P1 |
| Default to local on uncertain | Ambiguous queries don't hit cloud | P1 |

---

## Success Metrics

### Primary Metric (P17 Dogfooding)

**Lead Dev uses SymPAL daily** — measured by `sympal today` run each morning.

### Technical Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| Query routing accuracy | >80% | Log analysis |
| DSL execution success | >90% | Sandbox logs |
| Rehydration accuracy | >95% | Sample review |
| Latency (simple) | <5s | Instrumentation |
| Latency (complex) | <15s | Instrumentation |
| Latency vs baseline | ≤1.5x Claude | Benchmark |

### Decision Points (Circuit Breakers)

| Metric | Target | Threshold | Action if Below |
|--------|--------|-----------|-----------------|
| DSL execution | >90% | <80% | Debug compiler, tighten SymQL grammar |
| Rehydration accuracy | >95% | <90% | Review NER, simplify entity types |
| Routing accuracy | >80% | <70% | Revisit classifier keywords |
| Latency (simple) | <5s | >10s | Profile pipeline |

### Dogfooding Validity

Monthly reflection: "Would I use this if I hadn't built it?"

If answer is "no" — diagnose what's not working. Daily use from commitment ≠ daily use from value.

---

## Technical Decisions Summary

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Main language | Go | Simple, LLM-friendly, single binary |
| Sandbox | Deno subprocess | Built-in deny-by-default |
| Database | SQLite (graph schema) | Queryable + relationships |
| Local LLM | Ollama + Llama 3.2 3B | Fits 8GB RAM |
| OAuth storage | System keychain | OS-managed security |
| CLI style | Subcommands | Simple, no state |
| Query classifier | Keyword cascade | Zero latency, deterministic |
| Legend strategy | Task-based defaults | Balanced privacy/quality |
| NER review | Skip for known entities | Reduce friction |

---

## References

### Project Documents

| Document | Location | Relevance |
|----------|----------|-----------|
| Product Requirements | `foundations/prd.md` (v0.3.0) | What we're building |
| Binding Principles | `PRINCIPLES.md` (v1.1.0) | Hard constraints |
| Privacy Architecture | `foundations/privacy-innovations.md` (v3.0.0) | Full privacy design |
| Extraction Notes | `foundations/working/tdd-extraction-notes.md` | Decision audit trail |

### External Dependencies

| Dependency | Version |
|------------|---------|
| Go | 1.21+ |
| Deno | 1.40+ |
| Ollama | Latest |
| Llama 3.2 3B | llama3.2:3b |
| SQLite | 3.x |
| Google Calendar API | v3 |

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-19 | Fresh synthesis from extraction notes + privacy-innovations v3.0.0 + delta interview. Key changes: Three-tier → DSL/Ephemeral Slots/Local (no PaaP in V1); Semantic Projection → Ephemeral Slots; Added Query Classifier spec; Added cost model; Added failure UX; Added milestone gates; Task-based legend defaults; Skip NER review for known entities. |
| 0.2.0 | 2026-01-18 | Previous version (superseded) |
| 0.1.0 | 2026-01-18 | Initial synthesis |
