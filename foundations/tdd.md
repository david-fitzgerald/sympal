# SymPAL Technical Design Document

**Version:** 1.2.0
**Date:** 2026-01-27
**Status:** Final (team + Vero reviewed; all findings addressed)
**Author:** Kael + Ryn (synthesis from Lead Dev interview + delta)
**PRD Reference:** foundations/prd.md (v1.0.0)
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
| Contacts integration | Deferred until privacy proven with Calendar + Email |

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
│  - Calendar*    │ │   - Todo CRUD   │ │  ┌─────────────────┐    │
│  - Gmail**      │ │   - Day Planner │ │  │ Query Classifier│    │
│  (Contacts V2+) │ │                 │ │  │ (keyword cascade)│   │
│                 │ │                 │ │  └────────┬────────┘    │
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
│  - todos, cached calendar, cached email (M5)                    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                     LLM Providers                               │
│  - Claude API (DSL compilation, Ephemeral Slots reasoning)      │
│  - Ollama/Llama 3.2 3B (content tasks, local)                   │
└─────────────────────────────────────────────────────────────────┘

Legend: DSL = DSL Compilation, E.S = Ephemeral Slots, Local = Local LLM
        * = Read + Create only (see Calendar Write Controls)
        ** = Read only; requires NER >90% accuracy gate (see M5)
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

### Classification Categories

| Category | Route To | Exposure | Trigger Keywords |
|----------|----------|----------|------------------|
| STRUCTURED | DSL Compilation | Zero | "how many", "count", "list", "filter", "search", "find all", "show me", "when is" |
| HYBRID | DSL → Ephemeral Slots | Zero → Protected | Structured keyword + content/reasoning verb |
| REASONING | Ephemeral Slots | Protected | "should I", "prioritize", "advise", "recommend", "help me decide", "what's important" |
| CONTENT | Local LLM | Zero | "draft", "write", "compose", "summarize", "rewrite", "edit" |
| UNCERTAIN | Local LLM (default) | Zero | No confident match |

### Classification Algorithm

**Order of operations** (privacy-first — check zero-exposure paths first):

```
CLASSIFY(query):
    1. CHECK STRUCTURED PATTERNS (→ DSL Compilation, zero exposure)
       Keywords: "how many", "count", "list", "filter", "search", "find all", "show me", "when is"
       Confidence: HIGH if keyword match + no content verbs

    2. CHECK HYBRID PATTERNS (→ DSL then Ephemeral Slots)
       Detection: Structured keyword + reasoning/content verb
       Example: "Prioritize my meetings today" → filter first, then reason
       Process: Execute DSL, pass filtered results to Ephemeral Slots

    3. CHECK REASONING PATTERNS (→ Ephemeral Slots, protected exposure)
       Keywords: "should I", "prioritize", "advise", "recommend", "help me decide"
       Confidence: HIGH if reasoning keyword present

    4. CHECK CONTENT PATTERNS (→ Local LLM, zero exposure)
       Keywords: "draft", "write", "compose", "summarize", "rewrite"
       Confidence: HIGH if content verb present

    5. DEFAULT (→ Local LLM with quality warning)
       Log: "Uncertain classification, defaulting to most private"
       Always safe fallback — never expose data on uncertainty
```

### Confidence Thresholds

| Confidence | Behavior |
|------------|----------|
| HIGH (strong keyword match) | Route directly |
| MEDIUM (partial match) | Route with logging for review |
| LOW (no match) | Default to Local LLM (safest) |

### Hybrid Query Decomposition

Some queries combine structured operations with reasoning/content. Decompose them:

**Example**: "Summarize my urgent emails about Project Titan"

Without decomposition: Routes to Local LLM (content task), all data exposed locally.

With decomposition:
1. DSL extracts: `FILTER emails WHERE priority = 'urgent' AND subject CONTAINS 'Project Titan'`
2. Returns 3 emails (filtered from 500)
3. Only those 3 sent to Local LLM for summarization

**Benefit**: Minimizes data processed even within zero-exposure tiers.

#### State Machine

```
CLASSIFY_HYBRID
    │
    ▼
COMPILE_FILTER (generate SymQL for filter portion)
    │
    ▼
EXECUTE_FILTER (run SymQL locally)
    │
    ▼
CHECK_FILTER_RESULT
    │
    ├── If meaningful reduction (>50% filtered) → ROUTE_CONTENT
    │
    └── If minimal reduction (<50% filtered) → ROUTE_FULL_TO_LOCAL
```

**Meaningful reduction threshold**: 50%. If DSL doesn't meaningfully reduce the dataset, skip decomposition and route full query to Local LLM.

**V1 Scope**: Explicit hybrid patterns only (content verb + structured keyword). Expand detection with confidence in V1.5.

---

## Privacy Tier Implementation

### Tier 1: DSL Compilation (Zero Exposure)

**How it works**:
1. Query + schema description sent to Claude (no data)
2. Claude generates SymQL code
3. Code validated by parser
4. Executed by local Go interpreter (primary) or Deno sandbox (fallback)
5. Results returned to user

#### Execution Architecture

**Primary: Go Interpreter (~95% of queries)**

LLM generates constrained SymQL. Executed by local Go interpreter — no sandbox needed.

Why DSL over full code generation:
- 99.9% parse success vs 60-90% for general code (Anka DSL study)
- Eliminates code injection entirely (no arbitrary execution)
- No package hallucination (no imports exist)
- Easier to audit and validate
- Interpreter runs in-process — no subprocess, no IPC overhead

**Fallback: Deno Sandbox (<5% of queries)**

For queries exceeding SymQL expressiveness:

```bash
deno run --no-prompt --no-npm --no-remote [script.ts]
```

| Flag | Effect |
|------|--------|
| (no `--allow-*` flags) | All permissions denied by default |
| `--no-prompt` | Don't ask user for permissions interactively |
| `--no-npm` | Disallow npm package imports |
| `--no-remote` | Disallow fetching remote modules |

Data passed via stdin (JSON), results via stdout. 5-second timeout enforced by Go context.

#### SymQL Grammar

```
FILTER <table>
  WHERE <conditions>
JOIN <table2> ON <condition>
SCORE BY <field> [ASC|DESC]
AGGREGATE <function>(<field>)
WINDOW <date_range>
LIMIT <n>
RETURN <table|fields>
```

**Supported Operations**:
- `FILTER` — Select from table with WHERE conditions
- `JOIN` — Combine tables on field relationships
- `SCORE` — Rank results by fields (ASC/DESC)
- `AGGREGATE` — COUNT, SUM, AVG, MIN, MAX
- `WINDOW` — Date ranges (@today, @this_week, @next_week, @last_30_days)
- `LIMIT` — Cap result count
- `RETURN` — Specify output table/fields

**Comparison Operators**: =, !=, >, <, >=, <=, CONTAINS, IN

**Example**:
```
FILTER calendar_events
  WHERE start >= @this_week_start
    AND start <= @this_week_end
AGGREGATE COUNT(*)
RETURN result
```

**V1 Intentional Limitations**:
- No nested subqueries
- Simple JOINs only (single level)
- No complex GROUP BY

If >30% of "structured-looking" queries can't be expressed, expand grammar in V1.5.

#### DSL Hardening (V1)

**Golden-Set Test Harness**

Before running compiled SymQL on real data, validate against synthetic test cases:
1. Auto-generate small synthetic datasets matching schema
2. Run compiled SymQL against synthetic data
3. Verify output matches expected results
4. Only then execute against real data

**Benefit**: Prevents data exposure during testing; catches logic errors before real execution.

**Semantic Type Guards**

Validate types and bounds before execution:
- Date fields receive valid dates (not strings)
- Numeric comparisons use numbers
- LIMIT values are reasonable (1-1000)
- Field names exist in schema

**Deterministic Canonicalization**

Normalize queries to reduce variance in Claude's output:
- Standard field ordering
- Consistent whitespace
- Predictable placeholder naming

Reduces prompt sensitivity and improves cache hit rates.

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

**Escalation Framework**:

| Level | Detail | When Used |
|-------|--------|-----------|
| Ultra-minimal | Type only | Calendar scheduling (low reasoning) |
| Minimal | Type + role | Simple queries |
| Standard | Type + role + context | Reasoning tasks |
| Detailed | Full relationship | Complex drafting |

**Escalation rules**:
- Start at task-appropriate default
- Escalate only if quality degrades
- Track escalation frequency — if >30% of queries escalate, re-evaluate defaults
- Escalation is per-task, not per-entity (same entity can have different legend detail for different tasks)

**Leakage warning**: Over-escalation defeats privacy. Detailed legends can identify entities without names. Even slotted descriptions like "direct manager at my current employer, strained relationship" may be as identifying as a real name in a small organization. **This is a navigated tension, not a solved problem** — minimize legend detail aggressively.

**NER Confidence Handling** (measured by F1 score on held-out test set):
- HIGH (F1 >0.8): Auto-accept extraction
- MEDIUM (F1 0.5-0.8): Accept, log for review
- LOW (F1 <0.5): Skip review for known entities; prompt for truly new entities

**Rehydration Requirements**:
- LLM response must use exact placeholder format: `[E1]`, `[E2]`
- Structured output (JSON) preferred for reliable parsing
- Fallback cascade: retry prompt → local LLM → partial result with warning

**Rehydration Failure Categories**:

| Category | Example | Detectable? | Severity |
|----------|---------|-------------|----------|
| Missing placeholder | LLM says "your colleague" instead of `[E1]` | Yes (regex) | Low — shows anonymized |
| Wrong placeholder | `[E1]` used where `[E2]` meant | Partial (heuristics only) | High — wrong advice looks correct |
| Self-reference | `[E1] should meet with [E1]` | Yes (duplicate check) | Medium |
| Format corruption | `[E 1]` or `E1` | Yes (regex) | Low |

**Mitigation for "wrong placeholder"**: If response references placeholder in semantically inconsistent context (e.g., legend says `[E1]` is a person but response says "the project [E1]"), flag for user review rather than auto-rehydrating.

**Rehydration Pipeline** (5 steps):
1. **Validate response** — Check JSON structure, placeholder format
2. **Detect unbound placeholders** — Flag any `[Ex]` not in legend
3. **Handle possessives** — `[E1]'s` → expand correctly
4. **Case variations** — Match `[E1]`, `[e1]`, etc.
5. **Escape user input** — Prevent injection via entity names

**Fallback cascade** (if rehydration fails):
1. Retry with stricter prompt (explicit format instructions)
2. Route to Local LLM (full data, zero exposure)
3. Return partial result with warning (show anonymized response)

**Known Failure Cases** (route to Local LLM instead):

| Pattern | Why It Fails | Route To |
|---------|--------------|----------|
| Follow-up queries | Requires previous context | Local LLM |
| Reminders/recurring | Needs entity consistency over time | Local LLM |
| Multi-turn conversations | Consistent references across queries | Local LLM |
| Longitudinal summaries | "How has X changed?" needs history | Local LLM |

These patterns require consistent entity references that Ephemeral Slots (single-use placeholders) cannot provide.

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
| **Hybrid (DSL → Local)** | 2-4s | 1 cloud | ~$0.002 | High | Zero exposure |
| **Ephemeral Slots** | 1-3s | 1 cloud | ~$0.003 | High | Pattern-only |
| **Hybrid (DSL → Ephemeral)** | 2-4s | 2 cloud | ~$0.005 | High | Zero → Pattern |
| **Local LLM only** | 2-5s | 0 | $0 | Medium-Low | Zero exposure |

V1 target: 60-70% of queries hit DSL or Ephemeral Slots (cloud, high quality, privacy-preserved).

**Cost optimization**: DSL + Local Hybrid paths achieve high quality at ~$0.002 with zero exposure by filtering first, then processing locally.

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
| Policy gate blocks | "This query would send [field] to cloud. Allow once / Always allow / Use local instead?" | Prompt for consent, log decision |

**Design principles**:
1. Never silently fail
2. Never silently degrade privacy
3. Never expose raw data on failure
4. Always offer a path forward
5. Explain in plain language what happened

---

## User Control Mechanisms (V1)

### Progressive Consent Ladder

Micro-consent per data type. Permissions build incrementally.

**Flow example**:
1. First calendar query: "Allow SymPAL to read calendar titles?" [Yes/No]
2. User allows calendar titles
3. Later: "This query needs attendee names. Allow?" [Yes/No]
4. Permissions accumulate only as needed

**Defaults**: No access. User grants explicitly.

**Scope**: Each grant is scoped to:
- Data type (calendar, todos, etc.)
- Field set (titles, attendees, etc.)
- Time window (this session, 24 hours, always)

### Privacy Receipts + Undo

After each remote call, show receipt of what was sent.

**Receipt anatomy**:
```
Privacy Receipt #47
─────────────────────
Sent to: Claude API
Purpose: Calendar reasoning
Fields sent: [event types, time ranges]
Fields NOT sent: [titles, attendees, locations]
Timestamp: 2026-01-27 21:15:03
─────────────────────
```

**Undo mechanism**: "Undo last 24h" revokes all permissions granted in that window.

**Access**: `sympal receipts` shows recent receipts. `sympal receipts --detail` shows full payloads.

### Consent Recipes

Named preset policies for common use cases.

| Recipe | What It Allows |
|--------|---------------|
| Calendar-only | Calendar metadata to cloud; no todos, no content |
| Strict local | All queries route to Local LLM; no cloud |
| Balanced | DSL + Ephemeral Slots for reasoning; Local for content |

**Setup**: User selects recipe on first run. Can change anytime.

**Escalation**: If query needs more than recipe allows, one-line prompt: "This needs [X]. Allow once?"

### Time-Boxed Access (TTL)

Every consent grant has a time-to-live.

| TTL Option | Duration |
|------------|----------|
| This query only | Single use |
| This session | Until CLI exits |
| 1 hour | Auto-revoke after |
| 24 hours | Auto-revoke after |
| Always | Persists (can revoke manually) |

**Default**: 24 hours. Prevents permission creep.

**UI**: Countdown badge shows time remaining on active grants.

---

## Data Model

### Data Classification

| Data Type | Sensitivity | Storage | Sent to Cloud? |
|-----------|-------------|---------|----------------|
| Todos | Low | SQLite | Via Ephemeral Slots only |
| Calendar events | Medium | SQLite (cached) | Via Ephemeral Slots only |
| Email (M5) | High | SQLite (cached) | Via Ephemeral Slots only; requires NER >90% gate |
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

-- Email (M5)
CREATE TABLE emails (
    id TEXT PRIMARY KEY,
    subject TEXT NOT NULL,
    sender TEXT,               -- Entity ID if resolved
    recipients TEXT,           -- JSON array of entity IDs
    body_preview TEXT,         -- First 500 chars for display
    received_at TIMESTAMP,
    is_actionable BOOLEAN,     -- Detected by classifier
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
- Personal data (calendar, todos, email, entities)
- Query patterns (what user asks about)
- Identity (who the user is)

**Threat actors**:

| Actor | Motivation | In Scope? | Mitigation |
|-------|------------|-----------|------------|
| LLM Provider (passive) | Profile building | Yes | Ephemeral Slots defeat entity-level correlation; behavioral patterns remain visible (mitigated by token rotation, timing noise) |
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
| Prompt template normalization | Fixed structures, not free-form | Prompt fingerprinting | — |
| Length caps | Consistent padding/truncation | Length-based profiling | Content analysis |

**Prompt Template Normalization**:
- Use fixed prompt structures for each query type
- Avoid free-form user text in prompts where possible
- Consistent ordering of schema elements

**Length Caps**:
- Pad short queries to minimum length
- Truncate long queries to maximum
- Defeats length-based fingerprinting

**Accepted limitation**: Nation-state actor with long-term observation may still correlate patterns. Practical obscurity, not perfect anonymity.

### Security Requirements

| Requirement | Implementation | Principle |
|-------------|----------------|-----------|
| No raw PII to cloud | Privacy tier enforced, no bypass | P1 |
| Secure token storage | System keychain | P5 |
| Safe code execution | Deno sandbox, deny-by-default | P5 |
| Audit trail | Query log shows what was sent | P12 |
| User can inspect | `sympal log` shows recent queries | P4, P10 |

### Defense-in-Depth Security Controls (V1)

Multiple layers prevent data exposure even if one layer fails.

#### Taint-Tracked Serialization

Mark fields with sensitivity levels. Custom serializer checks tags before including fields.

```go
type CalendarEvent struct {
    ID        string    `taint:"public"`
    Title     string    `taint:"sensitive"`
    StartTime time.Time `taint:"derived"`
    Attendees []string  `taint:"sensitive"`
}
```

| Taint Level | Behavior |
|-------------|----------|
| `public` | Always serializable |
| `derived` | Serializable (computed from sensitive) |
| `sensitive` | Requires policy grant |
| `forbidden` | Never serializable |

**Benefit**: Prevents accidental sensitive field inclusion in payloads.

#### Policy-as-Code Gate

Rules engine evaluates every outbound request before sending.

```yaml
# Example policy rules
rules:
  - name: no-raw-calendar-titles
    match: { destination: "claude-api" }
    deny: { fields: ["title", "attendees"] }

  - name: require-ephemeral-slots
    match: { tier: "reasoning" }
    require: { transform: "ephemeral-slots" }
```

**Enforcement**: Denial = block + log + notify user. No silent bypass.

#### Strict Egress Firewall

Single exit point for all external requests.

| Control | Implementation |
|---------|----------------|
| Endpoint whitelist | Only Claude API, Google APIs allowed |
| Schema validation | Request shape must match registered schema |
| Size limits | Token budget per request |
| Rate limits | Prevent runaway queries |

**Benefit**: Belt-and-suspenders with policy gate.

#### Deterministic Redaction Layer

Pattern-based PII removal as fail-safe before Ephemeral Slots.

| Pattern | Example | Action |
|---------|---------|--------|
| Email addresses | `user@example.com` | Redact to `[EMAIL]` |
| Phone numbers | `555-123-4567` | Redact to `[PHONE]` |
| SSN patterns | `123-45-6789` | Redact to `[SSN]` |
| Credit cards | `4111-1111-1111-1111` | Redact to `[CC]` |

**Scope**: Narrow — catches obvious patterns NER might miss. Not replacement for Ephemeral Slots.

#### Schema-Hash Gate

Every allowed request schema hashed and registered in code.

```go
var allowedSchemas = map[string]bool{
    "sha256:abc123...": true,  // DSL compilation request
    "sha256:def456...": true,  // Ephemeral Slots request
}
```

**Enforcement**: Unregistered schema = block. Forces code review for any payload change.

---

## Explicitly Accepted Risks

### Risk 1: NER is the Achilles' Heel (CRITICAL — Privacy Failure Mode)

**The risk**: If NER misses an entity, real data goes to cloud LLM. **This is not a quality issue — it is a privacy failure.**

**Why we accept it** (with significant uncertainty):
- No perfect NER exists — this is a fundamental limitation
- User review *may* catch misses for new entities — **this is an untested hypothesis**
- Known entities auto-recognized (skip review) — reduces friction but may miss context-dependent references
- Alternative (no cloud) sacrifices too much quality for dogfooding value

**Mitigations** (effectiveness uncertain until validated):
- Conservative extraction (prefer false positives) — **feasibility depends on library choice**
- User confirmation for new/uncertain entities — **adds friction; user compliance unknown**
- Audit log for post-hoc review — catches issues after the fact, not before
- **Privacy fallback**: If NER confidence is LOW for a query, route entire query to Local LLM rather than risk exposure

**What would change this assessment**: If dogfooding reveals >5% undetected entity leakage, escalate to per-query review or restrict to Local LLM for high-sensitivity data.

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

### Phase 1: Foundation (M1) ✅ Complete

- [x] Go project scaffolding (Cobra CLI, config loading)
- [x] SQLite setup with schema
- [x] Basic todo CRUD (`sympal todo add/list/done/delete`)
- [x] Config file handling (`~/.sympal/config.yaml`)
- [x] Logging infrastructure (`~/.sympal/sympal.log`)
- [x] `sympal log` command (view recent queries, supports P10 user control)

**Gate**: ✅ Todo CRUD works end-to-end

### Phase 2: Calendar Integration (M2) ✅ Complete

- [x] Google OAuth flow (system keychain storage)
- [x] Calendar API integration (read events)
- [x] `sympal today` command (todos + calendar)
- [x] Basic day view (no LLM yet)
- [x] Token refresh on 401 (on-demand refresh)
- [x] HTTP status code checking before JSON parsing

**Gate**: ✅ Can view today's calendar and todos together

**OAuth Implementation Note**: Token refresh implemented as on-demand (refresh on 401). Token expiry tracking deferred — system relies on refresh token validity rather than proactive expiry management.

**Calendar Write Controls** (V1):
- **Supported**: Create events only (no modify, no delete in V1)
- **Confirmation**: All event creation requires explicit user confirmation before API call
- **Attendees**: Out of scope for V1 (no sending invites on user's behalf)

### Phase 3: DSL Compilation (M3)

- [ ] Knowledge graph schema and basic CRUD (entities, relationships tables)
- [ ] Query Classifier (keyword cascade + hybrid detection)
- [ ] Schema description generator
- [ ] Claude API integration
- [ ] SymQL lexer and parser
- [ ] SymQL executor (Go interpreter)
- [ ] Deno sandbox fallback (security validation required before use)
- [ ] DSL hardening (golden-set testing, type guards)
- [ ] Security controls (taint tracking, policy gate, egress firewall)
- [ ] Code validation → execution pipeline
- [ ] Test with todo + calendar queries

**Gate**: >90% of structured queries return correct results; test suite includes 20 real-world queries with >80% expressible in SymQL

**Security gate**: All security controls active. Deno sandbox must pass isolation tests (no network, no filesystem) before fallback is enabled.

### Phase 4: Ephemeral Slots (M4)

- [ ] Entity extraction (NER)
- [ ] Ephemeral placeholder generation (single-use random IDs)
- [ ] Legend construction (task-based defaults + escalation framework)
- [ ] Projection function (real → placeholders)
- [ ] Rehydration pipeline (5-step process)
- [ ] Known failure case routing (follow-ups, reminders → Local)
- [ ] User consent mechanisms (progressive consent, receipts)
- [ ] Test with reasoning queries

**NER Implementation Note**: Library choice to be determined during implementation. Options: Go-native (prose, go-ner) vs. subprocess to Python (spaCy). Prefer Go-native for single-binary goal unless accuracy gap is significant.

**NER Library Selection Criteria** (must pass before M4 complete):
- Minimum F1 score >0.80 on standard NER benchmark (CoNLL-2003 or OntoNotes)
- Must handle common entity types: PERSON, ORG, PROJECT (custom training may be needed for PROJECT)
- False negative rate <10% on calendar/todo test set (false negatives = privacy failures)
- If no library meets criteria, escalate to per-query user review for all Ephemeral Slots routes

**NER Confidence by Data Type**:
| Data Type | Required Accuracy | Notes |
|-----------|-------------------|-------|
| Calendar titles | >85% | Medium sensitivity |
| Todos | >85% | Medium sensitivity |
| Email (M5+) | >90% | High sensitivity — require per-query review if below |

**Gate**: >95% rehydration accuracy; <30% legend escalation rate

### Phase 5: Local LLM + Email Integration (M5)

- [ ] Ollama integration (Llama 3.2 3B)
- [ ] Content task routing
- [ ] Quality comparison logging (Local vs Claude baseline)
- [ ] End-to-end privacy tier (all three tiers integrated)
- [ ] Time-boxed access (TTL on permissions)
- [ ] Privacy receipts
- [ ] Gmail integration (read-only; OAuth scope expansion)
- [ ] Email-to-todo detection (actionable email identification)
- [ ] Daily use begins

**Local LLM Quality Benchmarking**:
Before M5 complete, benchmark Llama 3.2 3B on:
- Summarization tasks (target: >70% quality vs Claude)
- Classification tasks (target: >70% accuracy)

If benchmarks fail, route fuzzy data-processing tasks to Ephemeral Slots (protected exposure) rather than Local LLM.

**Email Integration Requirements**:
- NER accuracy on email content must be >90% F1 before email integration is enabled
- If NER accuracy <90%, require per-query entity review for email routes
- Email-to-todo detection is P1 (nice-to-have); manual email queries are P0

**Gate**: Lead Dev uses daily AND ≥50% of LLM queries route through DSL or Ephemeral Slots (not Local-only fallback)

**Quality gate**: Local LLM achieves >70% task success on summarization/classification benchmarks

**Email gate**: NER achieves >90% F1 on email test set before email integration enabled

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
M5 (Local LLM + Email Integration)
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

**Lead Dev uses SymPAL daily** — measured by:
1. `sympal today` run each morning (instrumented; logged to `~/.sympal/usage.log`)
2. Weekly reflection: "Did SymPAL provide value this week beyond what I'd get from direct Claude/calendar?"

**Success definition**: Daily use sustained for 2+ weeks with positive weekly reflections. If use drops or reflections turn negative, diagnose friction points before declaring M5 complete.

### Technical Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| Query routing accuracy | >80% | Log analysis |
| DSL execution success | >90% | Sandbox logs |
| Rehydration accuracy | >95% | Sample review |
| UNCERTAIN fallback rate | <20% | Log analysis |
| Latency (simple) | <5s | Instrumentation |
| Latency (complex) | <15s | Instrumentation |
| Latency vs baseline | ≤1.5x Claude | Benchmark |

### Decision Points (Circuit Breakers)

| Metric | Target | Threshold | Action if Below |
|--------|--------|-----------|-----------------|
| DSL execution | >90% | <80% | Debug compiler, tighten SymQL grammar |
| Rehydration accuracy | >95% | <90% | Review NER, simplify entity types |
| Routing accuracy | >80% | <70% | Revisit classifier keywords |
| UNCERTAIN fallback | <20% | >30% | Expand classifier keywords or add local LLM classifier |
| Latency (simple) | <5s | >10s | Profile pipeline |

### Consolidated Quality Gates

All gates that must pass before milestone completion:

| Milestone | Gate | Target | Failure Action |
|-----------|------|--------|----------------|
| **M3** | SymQL expressiveness | >80% of structured queries expressible | Expand grammar |
| **M3** | DSL execution success | >90% | Debug compiler |
| **M4** | Legend escalation rate | <30% | Re-evaluate task defaults |
| **M4** | Rehydration accuracy | >95% | Review NER, simplify |
| **M5** | NER accuracy (email) | >90% F1 | Require per-query review for email |
| **M5** | Local LLM quality | >70% task success (summarization/classification) | Route fuzzy tasks to Ephemeral Slots |
| **M5** | Privacy tier coverage | ≥50% queries via DSL/Ephemeral Slots | Expand classifier, improve DSL |
| **M5** | Dogfooding | Daily use with value (2+ weeks sustained) | Diagnose friction points |
| **M5** | Email integration | Gmail read working with NER gate passed | Defer email if NER <90% |

**Measurement timing**:
- M3 gates: Measured with 20-query test suite
- M4 gates: Measured over 100 reasoning queries
- M5 gates: Measured over 2-week dogfooding period

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
| Cloud LLM | Claude (hardcoded V1) | V1 simplicity; abstraction layer is V2 scope (P3 deferred) |
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
| Gmail API | v1 |

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.2.0 | 2026-01-27 | **Team + Vero review complete.** Fixes: (1) Added Gmail integration to M5, aligning with PRD P0; (2) Added knowledge graph to M3 deliverables; (3) Added NER library selection criteria with F1 benchmarks; (4) Reframed NER risk as privacy failure with uncertainty markers; (5) Strengthened legend leakage warning; (6) Clarified rehydration "wrong placeholder" as partial detection; (7) Added Deno sandbox security validation requirement; (8) Clarified success metric with measurement approach; (9) Marked M1/M2 as complete; (10) Specified NER confidence metric as F1; (11) Updated OAuth note with token expiry approach. |
| 1.2.0-draft | 2026-01-27 | **Full sync with privacy-innovations.md v3.0.0**: Added DSL hardening (golden-set, type guards); Hybrid decomposition state machine; Security controls (taint tracking, policy gate, egress firewall, deterministic redaction, schema-hash gate); UX controls (progressive consent, privacy receipts, consent recipes, time-boxed access); Legend escalation framework; Ephemeral Slots known failure cases; Rehydration pipeline; M5 local LLM benchmarking; Consolidated quality gates; Expanded correlation mitigations. **Pending team + Vero review.** |
| 1.1.0 | 2026-01-27 | Synced M3 spec with privacy-innovations.md v3.0.0: (1) Execution architecture — Go interpreter primary, Deno fallback; (2) SymQL grammar expanded (FILTER/JOIN/SCORE/AGGREGATE/WINDOW/RETURN); (3) Query Classifier spec expanded with 5 categories and hybrid decomposition. |
| 1.0.3 | 2026-01-19 | Vero review fixes (all MINOR): (1) Fixed diagram — Calendar* with note for R+Create only; (2) Added OAuth implementation note (prefer on-demand refresh); (3) Added NER implementation note (prefer Go-native unless accuracy gap). Status → Final. |
| 1.0.2 | 2026-01-19 | Adversary challenge fixes: (1) Added UNCERTAIN fallback rate metric (<20%) with circuit breaker; (2) Added rehydration failure categories with "wrong placeholder" mitigation; (3) Fixed correlation claim to include behavioral pattern caveat; (4) Made M5 gate objective (≥50% queries via DSL/Ephemeral Slots); (5) Documented SymQL intentional limitations, added expressiveness test to M3; (6) Specified calendar write controls (create only, confirmation required, no attendees). |
| 1.0.1 | 2026-01-19 | Vale checkpoint fixes: (1) Added `sympal log` command to M1 for P10 user control; (2) Clarified Cloud LLM decision — V1 hardcodes Claude, abstraction layer is V2 scope (P3 deferred). |
| 1.0.0 | 2026-01-19 | Fresh synthesis from extraction notes + privacy-innovations v3.0.0 + delta interview. Key changes: Three-tier → DSL/Ephemeral Slots/Local (no PaaP in V1); Semantic Projection → Ephemeral Slots; Added Query Classifier spec; Added cost model; Added failure UX; Added milestone gates; Task-based legend defaults; Skip NER review for known entities. |
| 0.2.0 | 2026-01-18 | Previous version (superseded) |
| 0.1.0 | 2026-01-18 | Initial synthesis |
