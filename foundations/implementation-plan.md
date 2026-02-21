# Implementation Plan

**Version:** 1.1.0
**Date:** 2026-01-27
**Status:** Hibernating (resuming ~early April 2026)
**Purpose:** Session-start reference for implementation phase
**TDD Reference:** foundations/tdd.md v1.2.0

---

## Session Start Checklist

1. Read this file
2. Check current milestone progress (below)
3. Review any open GitHub issues
4. Pick up where we left off

---

## Current Status

**Phase:** M3 DSL Compilation
**Progress:** Ready to start

**M1 Foundation:** ✅ Complete (2026-01-21)
**M2 Calendar:** ✅ Complete + Polished (2026-01-27)
**M3 DSL Compilation:** 🔲 Not started
**M4 Ephemeral Slots:** 🔲 Not started
**M5 Local LLM + Email:** 🔲 Not started

### Resume Point (2026-02-21)

**Hibernating.** Other projects taking priority. Resuming ~early April 2026.

**When resuming:** M3 DSL Compilation ready to start. TDD v1.2.0 finalized with all review findings addressed. Implementation plan rebuilt from TDD.

---

## Milestone Wrap-Up Procedure

After completing each milestone:

### 1. Manual Testing (All Milestones)
- Run through all new commands/features
- Test edge cases (empty inputs, invalid IDs, missing files)
- Verify nothing broke from previous milestones

### 2. Code Review (Security-Critical Milestones)
Use personas for M3+ (sandboxing, LLM integration, privacy tier):

| Persona | Focus |
|---------|-------|
| **Ryn** | Security review — "how will this fail?" |
| **Seren** | Code craft — "is this well-crafted?" |
| **Kael** | Implementation — "will this survive reality?" |

**Not needed for:** M1, M2 (low security surface)

### 3. Dogfooding (Starts M1)
- Add sympal to PATH: `export PATH="$PATH:/Users/df/pg/sympal"`
- Use daily for real tasks
- Log friction points to ROADMAP.md

### 4. Documentation Update
- Update this file (Current Status section)
- Update CONTEXT.md if project status changed
- Git tag: `git tag v0.x.0-mN`

### 5. Vero Review
**Not for code** — Vero reviews foundational docs only. Code review uses team personas above.

---

## Learning-First Approach

**Goal:** Build coding confidence through real implementation, not just receive code.

**Skill baseline:** Intermediate prompt architecture, basic coding. Moving toward confident independent writing.

### How We Work

| Mode | Trigger | Style |
|------|---------|-------|
| **Teach** (default) | New concepts, architecture decisions, anything you'll debug later | Explain *why* before code, flag transferable patterns, you write/modify |
| **Ship** | Boilerplate, well-trodden patterns, "just do it" signal | I write, you review, move fast |
| **Deep dive** | Security-critical, novel patterns, genuine curiosity | Go deep with sources, build full understanding |

### The Tension to Watch

> **Learning-as-excuse**: Planning serves learning even when shipping serves project.

**Legitimate deep dives:**
- Security-critical (sandboxing, NER privacy implications)
- Architecture decisions with tradeoffs
- Patterns you'll use repeatedly

**Probably learning-as-excuse:**
- Fully understanding X before writing simple code that uses X
- Researching edge cases for MVP features
- Perfecting understanding when "good enough" would ship

**When in doubt:** Ask "What would we learn by building instead?"

---

## Milestones

### M1: Foundation ✅

**Status:** Complete (2026-01-21)

**Deliverables:**
- [x] Go project scaffolding (Cobra CLI)
- [x] SQLite setup with schema
- [x] Todo CRUD (`sympal todo add/list/done/delete`)
- [x] Config file handling (`~/.sympal/config.yaml`)
- [x] Logging infrastructure (`~/.sympal/sympal.log`)
- [x] `sympal log` command (view recent queries)

**Gate:** ✅ Todo CRUD works end-to-end

**Learning outcomes:**
- Go fundamentals (syntax, error handling, packages)
- CLI structure patterns (Cobra)
- Project scaffolding (transferable to any Go project)
- SQLite basics
- Config/logging patterns (slog for structured logging)

---

### M2: Calendar Integration ✅

**Status:** Complete + Polished (2026-01-27)

**Deliverables:**
- [x] Google OAuth flow (system keychain storage)
- [x] Calendar API integration (read events)
- [x] `sympal today` command (todos + calendar)
- [x] Basic day view (no LLM yet)
- [x] Token refresh on 401
- [x] HTTP status code checking

**Gate:** ✅ Can view today's calendar and todos together

**Learning outcomes:**
- OAuth flow (reusable for any Google/OAuth API)
- API integration patterns
- Token management (refresh, storage)
- HTTP client error handling

---

### M3: DSL Compilation

**Status:** Not started

**Deliverables:**
- [ ] Knowledge graph schema (entities, relationships tables)
- [ ] Query Classifier (keyword cascade routing)
- [ ] Schema description generator
- [ ] Claude API integration
- [ ] SymQL lexer and parser
- [ ] SymQL executor (Go interpreter)
- [ ] Deno sandbox fallback (with security validation)
- [ ] DSL hardening (golden-set testing, type guards)
- [ ] Security controls (taint tracking, policy gate, egress firewall)
- [ ] `sympal query "..."` command

**Gates:**
- >90% structured queries return correct results
- 20-query test suite with >80% expressible in SymQL
- Deno sandbox passes all isolation tests (see Security Validation below)
- All security controls active and tested

**Knowledge Graph Scope** (define before Chunk 1):
- **Entities**: People (from calendar attendees), Projects (extracted from event titles), Organizations
- **Relationships**: `attends` (person→event), `works_on` (person→project), `part_of` (project→org)
- **Connection to M4**: Entities populate NER known-entity list; relationships inform legend construction

**Implementation Chunks:**

| # | Component | Mode | Depends On | Notes |
|---|-----------|------|------------|-------|
| 1 | Knowledge Graph Schema | Teach | — | SQLite tables for entities/relationships; first new domain |
| 2 | Security Foundation | Deep dive | — | Taint tracking types, policy gate stubs, egress interface |
| 3 | Query Classifier | Teach | 1 | Foundation for routing; keyword patterns + confidence |
| 4 | Claude API Client | Teach | 2 | Standard HTTP through egress; first LLM API integration |
| 5 | Schema Generator | Teach | 1 | Describe tables without exposing data |
| 6 | SymQL Lexer | Teach | — | Tokenize SymQL; reusable parsing skill |
| 7 | SymQL Parser | Teach | 6 | Tokens → AST; valuable skill |
| 8 | SymQL Executor | Teach | 7, 2 | Execute AST against SQLite with taint tracking |
| 9 | DSL Hardening | Teach | 8 | Golden-set testing, type guards, canonicalization |
| 10 | Deno Sandbox | Deep dive | 2 | Build sandbox + run isolation tests (must pass before enabling) |
| 11 | Security Hardening | Deep dive | 2, 10 | Full policy implementation, egress firewall rules |
| 12 | Integration + CLI | Ship | all | Wire together, add `sympal query` command |

**Deno Sandbox Isolation Tests** (must pass before Chunk 10 complete):

| Test | Attack | Expected |
|------|--------|----------|
| Network isolation | HTTP request, DNS lookup, localhost | All blocked |
| Filesystem isolation | File read, write, directory traversal | All blocked |
| Process isolation | Subprocess spawn, FFI | All blocked |
| Resource limits | CPU >5s, Memory >100MB | Killed cleanly |

**Test procedure**: Create malicious SymQL payloads for each test. Run in sandbox. Verify blocked + clean termination. Test BEFORE wiring into query command.

**Learning focus:**
- Parser construction (lexer → parser → AST → executor)
- Sandboxing and security (Deno deny-by-default)
- LLM API integration and prompt engineering
- Compile-don't-interpret pattern

**Security-critical (go deep):**
- Deno sandbox isolation — must pass tests before enabled
- Taint-tracked serialization
- Egress firewall (single exit point)
- Policy-as-code gate

**Package structure** (M3 new packages):
```
internal/
├── graph/              # Knowledge graph (entities, relationships)
├── llm/
│   └── claude/         # Claude API client
├── privacy/
│   ├── classifier/     # Query routing
│   ├── schema/         # Schema description generator
│   ├── symql/          # Lexer, parser, executor
│   ├── sandbox/        # Deno fallback
│   └── security/       # Taint tracking, policy gate, egress
cmd/sympal/
└── query.go            # sympal query command
```

---

### M4: Ephemeral Slots

**Status:** Not started

**Deliverables:**
- [ ] NER library selection (F1 >0.80 on benchmark + <10% false negatives on domain test set)
- [ ] Entity extraction with confidence scoring
- [ ] Ephemeral placeholder generation (single-use random IDs)
- [ ] Legend construction (task-based defaults + escalation)
- [ ] Projection function (real → placeholders)
- [ ] Rehydration pipeline (5-step process)
- [ ] Known failure case routing (follow-ups → Local LLM)
- [ ] User consent mechanisms (progressive consent, receipts)
- [ ] Consent gate passed (flow tested, no confusion about what's shared)

**Gates:**
- NER library achieves F1 >0.80 on CoNLL-2003 or OntoNotes
- NER false negative rate <10% on domain-specific test set (calendar/todo)
- Rehydration accuracy >95%
- Legend escalation rate <30%
- Consent flow tested with clear user understanding

**NER Confidence Thresholds** (per entity type):

| Entity Type | Min Confidence | Rationale |
|-------------|----------------|-----------|
| PERSON | ≥0.90 | Highest privacy risk |
| ORG | ≥0.90 | Highest privacy risk |
| PROJECT | ≥0.85 | Medium risk, often ambiguous |
| DATE/TIME | ≥0.80 | Lower risk, structured |

**Compound entity policy**: If ANY entity in query fails its threshold → route entire query to Local LLM (all-or-nothing per query).

**Legend detail limits**:
- Task type: Allowed (e.g., "meeting", "email")
- Role: Allowed (e.g., "manager", "colleague")
- Relationship context: Minimal (e.g., "professional")
- Content details: NEVER include in legend

**Domain-Specific Test Set** (build before Chunk 1):
1. Create 50-100 synthetic calendar/todo examples with known entities
2. Include edge cases: nicknames, project codes, ambiguous references
3. Evaluate candidate libraries on THIS set, not just CoNLL
4. CoNLL is rough filter; domain set is pass/fail gate

**Implementation Chunks:**

| # | Component | Mode | Depends On | Notes |
|---|-----------|------|------------|-------|
| 0 | Domain Test Set | Ship | — | Build before library evaluation |
| 1 | NER Library Evaluation | Deep dive | 0 | Benchmark on domain set; privacy-critical |
| 2 | NER Integration | Teach | 1 | Extract entities with confidence scores |
| 3 | Placeholder Generation | Teach | — | Cryptographic randomness, single-use |
| 4 | Legend Construction | Teach | — | Task-based defaults, escalation framework |
| 5 | Projection Function | Teach | 2, 3, 4 | Entity → placeholder mapping |
| 6 | Rehydration Pipeline | Teach | 5 | 5-step process: validate, detect, possessives, case, escape |
| 7 | Failure Case Routing | Teach | 2 | Detect follow-ups, reminders → route to Local |
| 8 | Consent Mechanisms | Teach | — | Progressive consent ladder, privacy receipts |
| 9 | Integration | Ship | all | Wire into privacy tier, test with reasoning queries |

**Learning focus:**
- NER/text processing (transferable skill)
- Privacy-preserving state management
- Projection/rehydration pattern (novel technique)

**Privacy-critical (go deep):**
- NER accuracy is THE privacy dependency — if it fails, real data leaks
- Legend minimization — over-detail defeats privacy
- False negative handling — route to Local LLM on low confidence

**Package structure** (cumulative — M3 packages shown for context):
```
internal/
├── graph/              # (M3) Knowledge graph
├── llm/
│   └── claude/         # (M3) Claude API client
├── privacy/
│   ├── classifier/     # (M3) Query routing
│   ├── schema/         # (M3) Schema description generator
│   ├── symql/          # (M3) Lexer, parser, executor
│   ├── sandbox/        # (M3) Deno fallback
│   ├── security/       # (M3) Taint tracking, policy gate, egress
│   ├── ner/            # (M4) NER wrapper + confidence scoring
│   ├── slots/          # (M4) Placeholder generation, legend, rehydration
│   └── consent/        # (M4) Progressive consent, receipts
```

---

### M5: Local LLM + Email Integration

**Status:** Not started

**Deliverables:**
- [ ] Ollama integration (Llama 3.2 3B)
- [ ] Content task routing
- [ ] Quality comparison logging (Local vs Claude)
- [ ] End-to-end privacy tier (all three tiers integrated)
- [ ] Time-boxed access (TTL on permissions)
- [ ] Privacy receipts (`sympal receipts` command)
- [ ] Gmail integration (requires NER >90% F1 gate; read-only)
- [ ] Email-to-todo detection (actionable email identification)

**Gates:**
- Local LLM achieves >70% task success on summarization/classification
- NER achieves >90% F1 on email test set (BLOCKING for email integration)
- ≥50% of LLM queries route through DSL or Ephemeral Slots
- Daily use sustained 2+ weeks with positive weekly reflections
- Gmail integration working (only if NER gate passed)

**Dogfooding Structure** (gate, not implementation chunk):
- Week 1-2: Daily `sympal today` use (logged to `~/.sympal/usage.log`)
- Weekly reflection questions:
  1. Did I use sympal daily?
  2. Did it provide value beyond direct tools?
  3. What friction points emerged?
- Log reflections to ROADMAP.md dogfood-feedback section

**Implementation Chunks:**

| # | Component | Mode | Depends On | Notes |
|---|-----------|------|------------|-------|
| 1 | Ollama Client | Ship | — | HTTP client for local Ollama |
| 2 | Local LLM Benchmarking | Teach | 1 | Run benchmarks, decide routing strategy |
| 3 | Content Routing | Teach | 2 | Route content tasks to Local LLM |
| 4 | Privacy Tier Integration | Teach | 3 | Wire DSL + Ephemeral Slots + Local together |
| 5 | Time-Boxed Access | Teach | — | TTL on permissions, countdown tracking |
| 6 | Privacy Receipts | Teach | — | Receipt generation, `sympal receipts` command |
| 7 | NER Email Validation | Deep dive | — | Must pass >90% F1 BEFORE email enabled |
| 8 | Gmail OAuth Expansion | Ship | 7 passes | Add Gmail scope (only if NER gate passed) |
| 9 | Email Caching | Ship | 8 | Fetch and cache emails locally |
| 10 | Email-to-Todo Detection | Teach | 9 | Identify actionable emails |

**Learning focus:**
- Local LLM orchestration
- Multi-model routing (quality vs privacy tradeoffs)
- Email API integration (extends OAuth skills from M2)

**Privacy-critical (go deep):**
- NER on email — highest sensitivity data, requires >90% accuracy
- If NER <90% on email, require per-query entity review

**Package structure** (cumulative — M3+M4 packages shown for context):
```
internal/
├── graph/              # (M3) Knowledge graph
├── llm/
│   ├── claude/         # (M3) Claude API client
│   └── ollama/         # (M5) Ollama client
├── email/              # (M5) Gmail API client
├── privacy/
│   ├── classifier/     # (M3) Query routing
│   ├── schema/         # (M3) Schema description generator
│   ├── symql/          # (M3) Lexer, parser, executor
│   ├── sandbox/        # (M3) Deno fallback
│   ├── security/       # (M3) Taint tracking, policy gate, egress
│   ├── ner/            # (M4) NER wrapper + confidence scoring
│   ├── slots/          # (M4) Placeholder generation, legend, rehydration
│   ├── consent/        # (M4) Progressive consent
│   └── receipts/       # (M5) Privacy receipts generation
cmd/sympal/
├── query.go            # (M3) sympal query command
├── receipts.go         # (M5) sympal receipts command
└── email.go            # (M5) email-related commands (if any)
```

---

## Dependency Graph

```
M1 (Foundation) ✅
    │
    ▼
M2 (Calendar) ✅
    │
    ▼
M3 (DSL Compilation)
    │
    ▼
M4 (Ephemeral Slots)
    │
    ▼
M5 (Local LLM + Email)
```

---

## Technical Stack

| Tool | Purpose | Milestone |
|------|---------|-----------|
| Go | Main language | M1+ |
| SQLite | Local data storage | M1+ |
| Deno | Sandbox for generated code | M3 |
| Claude API | Cloud LLM (DSL compilation, reasoning) | M3+ |
| Ollama + Llama 3.2 3B | Local LLM | M5 |
| Google Calendar API | Calendar integration | M2+ |
| Gmail API | Email integration | M5 |

---

## Quality Gates Summary

From TDD v1.2.0:

| Milestone | Gate | Target | Failure Action |
|-----------|------|--------|----------------|
| **M3** | SymQL expressiveness | >80% of structured queries | Expand grammar |
| **M3** | DSL execution success | >90% | Debug compiler |
| **M3** | Deno isolation | Pass all security tests | Block fallback until fixed |
| **M4** | NER library benchmark | F1 >0.80 | Try different library |
| **M4** | NER false negatives | <10% | Escalate to per-query review |
| **M4** | Legend escalation rate | <30% | Re-evaluate task defaults |
| **M4** | Rehydration accuracy | >95% | Review NER, simplify |
| **M5** | NER accuracy (email) | >90% F1 | Require per-query review for email |
| **M5** | Local LLM quality | >70% task success | Route fuzzy tasks to Ephemeral Slots |
| **M5** | Privacy tier coverage | ≥50% via DSL/Ephemeral Slots | Expand classifier |
| **M5** | Dogfooding | 2+ weeks daily use with value | Diagnose friction |
| **M5** | Email integration | Working with NER gate | Defer email if NER <90% |

---

## Team Usage During Implementation

| Persona | When to Use |
|---------|-------------|
| **Kael** (implementation) | Architecture decisions, "will this survive reality?" |
| **Ryn** (systems/security) | Security review, failure modes, "how will this fail?" |
| **Seren** (code craft) | Code review before major commits, "is this well-crafted?" |
| **Orin** (user advocacy) | UX decisions, "are users better off?" |
| **Adversary** | Challenge assumptions, red-team security |

**Not for:** Every small change, creating code (that's us).

---

## Working Patterns

- Commit after completing tasks
- Push after each commit
- End sessions with everything committed and pushed
- Update "Current Status" section when milestone progress changes

---

## Dogfood Feedback

See [ROADMAP.md](/ROADMAP.md#dogfood-feedback) — single source of truth for friction points and future ideas.

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.1.0 | 2026-01-27 | **Team + Vero review fixes.** M3: Reordered chunks (security foundation before sandbox); added Deno isolation tests spec; added knowledge graph scope; added dependency column. M4: Added NER confidence thresholds per entity type; added domain-specific test set requirement; added consent gate; cumulative file structure. M5: Reordered chunks (NER validation before email); removed dogfooding from chunks (it's a gate); added dogfooding structure; cumulative file structure. |
| 1.0.0 | 2026-01-27 | **Rebuilt from TDD v1.2.0.** M3: Added knowledge graph, DSL hardening, security gates. M4: Added NER benchmarks, rehydration pipeline, consent mechanisms. M5: Added Gmail integration, email gates, 2-week dogfooding. Added Quality Gates Summary table. |
| 0.5.0 | 2026-01-27 | M2 polished (token refresh, todos in today). M3 plan detailed with 10 implementation chunks. |
| 0.4.0 | 2026-01-26 | M2 complete: OAuth flow, calendar API, `sympal auth`, `sympal today` |
| 0.3.0 | 2026-01-24 | M2 in progress (~30%): keyring + config complete, auth skeleton |
| 0.2.0 | 2026-01-21 | M1 complete, added Milestone Wrap-Up Procedure |
| 0.1.0 | 2026-01-20 | Initial plan created |
