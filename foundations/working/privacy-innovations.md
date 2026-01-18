# Privacy Innovations — Novel Approaches for SymPAL

**Version**: 0.2.0
**Date**: 2026-01-18
**Status**: Concepts defined with critical analysis, awaiting implementation
**Context**: These are novel approaches developed during PRD challenge phase, distinct from existing research. See `privacy-research.md` for survey of existing techniques.

---

## Why This Document Exists

The privacy research revealed that existing approaches have significant tradeoffs:
- **Local LLMs**: 10-20% quality gap vs frontier
- **PII Redaction**: Loses context ("swap don't blank" helps but limited)
- **Differential Privacy**: 20-40% accuracy loss
- **HE/MPC**: Not practical for interactive use

We developed novel approaches that aren't in the literature. This document captures them for V1 implementation and future development.

**Important**: These are applied innovations, not fundamental research. Section-specific limitations, assumptions, and failure modes are documented throughout. See [Open Questions](#open-questions) for unresolved issues.

---

## Threat Model

### What We're Defending Against

| Adversary | Concern | Design Response |
|-----------|---------|-----------------|
| LLM Provider (passive) | Profile building, data monetization | Semantic Projection — they see patterns, not identities |
| LLM Provider (active) | Deanonymization via correlation | Token rotation, query batching, timing noise |
| Nation-state (subpoena) | Forced disclosure of accumulated data | Provider has patterns only — can't reconstruct original data |
| Local device compromise | Full access to real data | Out of scope for V1 |
| Network observer | Query interception | Standard TLS (not a novel concern) |

### What We're NOT Defending Against

- **Perfect anonymity**: A motivated nation-state with long observation may correlate patterns. Goal is practical obscurity.
- **Malicious local code**: If user's device is compromised, all bets are off.
- **Provider collusion**: If LLM provider actively cooperates with adversary targeting you specifically.
- **Side-channel attacks**: Timing, power analysis, etc. beyond query-level protections.

### Threat Model Assumptions

- LLM providers are honest-but-curious (follow protocols, but may analyze data)
- TLS is not broken
- User's local environment is trusted
- Attackers don't have access to our source code's internal mappings

---

## Core Insight

**LLMs are pattern processors.** They don't need YOUR data — they need data with the SAME PATTERNS as yours.

If we can preserve reasoning-relevant structure while removing identity, we get utility without surveillance.

---

## Novel Approach 1: Semantic Projection ("Shadow Data")

### The Idea

Maintain two parallel representations:

**Real Data (local only)**:
```
Email from John Smith (CEO, Acme Corp)
about Q4 merger with Beta Inc
sent Tuesday, marked urgent
```

**Semantic Shadow (sent to LLM)**:
```
Email from [CONTACT:executive,important,frequent]
about [TOPIC:business,sensitive,time-bound]
sent [TIME:recent], marked [PRIORITY:high]
```

The LLM reasons over the shadow. The reasoning transfers to real data. The LLM never knew identities.

### Why It's Different

| Approach | What LLM Sees | What's Lost |
|----------|---------------|-------------|
| PII Redaction | "Email from *** about ***" | Context destroyed |
| Pseudonymization | "Email from Person_A7" | Reversible; LLM still accumulates |
| **Semantic Projection** | Typed relationship pattern | Identity, but NOT reasoning value |

### How It Works

1. **Local Knowledge Graph**: Typed graph of your life data
   - Contacts: `executive`, `friend`, `family`, `vendor`
   - Topics: `business`, `personal`, `health`, `financial`
   - Relationships: `reports-to`, `collaborates`, `family-of`

2. **Projection Function**: Map real entities to typed placeholders
   - "John Smith" → `[PERSON:colleague,senior,trusted,frequent]`
   - "Q4 merger" → `[PROJECT:business,confidential,deadline-driven]`

3. **LLM Reasoning**: Cloud LLM operates on projected patterns

4. **Rehydration**: Map LLM response back to real entities

### Prior Art & What's Actually Novel

**Established techniques**:
- Pseudonymization with typed placeholders (arXiv:2502.15233 and others)
- Named Entity Recognition for PII detection (mature field)
- Context-preserving anonymization (active research area)

**What's novel in our approach**:
- **Richer type vocabulary** derived from personal knowledge graph (not generic NER categories)
- **Rehydration workflow** that maps LLM responses back to real entities
- **Privacy-first framing** — existing work optimizes for accuracy; we optimize for non-reversibility

**Honest framing**: This is applied innovation combining proven techniques in a novel workflow, not fundamental research.

### What The LLM Provider Accumulates

Over time they learn you have:
- Important executive contacts
- Time-sensitive business matters
- Mix of personal and professional priorities

But they can NEVER know:
- WHO those contacts are
- WHAT companies you deal with
- WHAT specifically you're working on

You become statistically indistinguishable from millions of others.

### Failure Modes & Limitations

| Failure Mode | Severity | Mitigation | Residual Risk |
|--------------|----------|------------|---------------|
| Type vocabulary leaks identity | High | Limit specificity; coarsen rare combinations | Highly specific types (e.g., `[PERSON:colleague,neurosurgeon,marathon-runner]`) may be unique |
| Cross-query correlation | High | Token rotation + batching | Long-term patterns still observable |
| NER accuracy gap | Medium | User review of first mentions | Missed entities leak to LLM |
| Rehydration ambiguity | Medium | Session-scoped mappings with context | "The colleague" may be ambiguous if multiple mentioned |
| Novel entities (no prior type) | Medium | Default to generic type + user prompt | New contacts require classification |
| Content tasks bypass projection | Low | Route to local LLM | Quality tradeoff for content tasks |

### Assumptions

| Assumption | Status | Validation Path |
|------------|--------|-----------------|
| Local models can type entities accurately | Unvalidated | Benchmark on personal data samples |
| Type vocabulary balances richness ↔ genericity | Unvalidated | User study on projection quality |
| Users accept first-mention classification prompts | Unvalidated | Dogfooding feedback |
| Rehydration rarely ambiguous in practice | Unvalidated | Log ambiguity rate post-launch |

### V1 Implementation Priority

**Phase B of M3 (Privacy milestone)**

---

## Novel Approach 2: LLM as Compiler

### The Idea

For structured queries, the LLM never touches your data. Instead:

**Traditional**:
```
User: "Find emails about deadlines conflicting with vacation"
→ Send emails + calendar to LLM
→ LLM returns answer
```

**Compiler approach**:
```
User: "Find emails about deadlines conflicting with vacation"
→ Send query PATTERN to LLM (describe data shape, not data)
→ LLM returns: executable filter logic + conflict detection algorithm
→ Run locally against real data
→ Get answer
```

The LLM is a **reasoning consultant** who helps build tools but never sees your files.

### Why It's Different

- Zero data exposure (not reduced — zero)
- LLM output is reusable (same logic applies to future queries)
- Builds a library of local capabilities over time

### Prior Art & What's Actually Novel

**Established techniques**:
- NL2SQL is mature — AWS, Oracle, Microsoft have production systems
- Text-to-code generation (Codex, StarCoder, etc.)
- Schema-based query generation

**What's novel in our approach**:
- **Privacy framing**: Existing systems optimize for accuracy; we optimize for zero data exposure
- **Personal data context**: Schema descriptions include relationship semantics, not just column types
- **Hybrid routing**: Compiler for structured, projection for reasoning, local for content

**Honest framing**: This is proven technology in a novel application context. The innovation is the routing architecture, not the compilation itself.

### What It Works For

**Good fit (~60-70% of queries)**:
- Scheduling and conflict detection
- Filtering and search
- Priority scoring based on rules
- Pattern matching
- Data transformation

**Poor fit**:
- Summarization (needs content)
- Creative tasks (needs context)
- Ambiguous queries (needs data to disambiguate)

### Example

**User query**: "Show me all meetings next week with people I haven't emailed in 30 days"

**LLM receives**:
```
Data shapes available:
- calendar_events: {title, start, end, attendees[]}
- emails: {from, to, date, subject}
- contacts: {name, email}

Query: Find calendar events in date range where attendees
have no email activity in past N days
```

**LLM returns**:
```python
def find_neglected_meeting_contacts(calendar, emails, date_range, inactive_days):
    # Filter calendar events in range
    upcoming = [e for e in calendar if e.start in date_range]

    # Get attendee emails
    attendees = flatten([e.attendees for e in upcoming])

    # Find last email date per contact
    last_contact = {}
    for email in emails:
        for addr in [email.from] + email.to:
            if addr not in last_contact or email.date > last_contact[addr]:
                last_contact[addr] = email.date

    # Filter to inactive
    cutoff = today() - days(inactive_days)
    neglected = [a for a in attendees if last_contact.get(a, MIN_DATE) < cutoff]

    return [e for e in upcoming if any(a in neglected for a in e.attendees)]
```

**User runs locally** with real data. LLM never saw calendar or emails.

### Failure Modes & Limitations

| Failure Mode | Severity | Mitigation | Residual Risk |
|--------------|----------|------------|---------------|
| Vulnerable code generation | High (12-65% of LLM code has vulnerabilities per research) | AST sandbox with strict allowlist | Some vulnerability classes may slip through |
| Package hallucination | High (44% per research) | No external imports allowed | Limits expressiveness |
| Logic errors in generated code | Medium | Test with sample data; user verification | Silent wrong answers |
| Query too complex for compiler | Medium | Fallback to projection | May require multiple round trips |
| Schema description leaks info | Low | Describe shapes, not examples | Very generic schemas only |

### Sandbox Architecture

**V1 Specification: AST Parsing with Allowlist**

1. **Parse** generated code as Python AST
2. **Validate** against strict allowlist:
   - **Allowed**: list comprehensions, dict operations, basic math, datetime operations, string methods
   - **Disallowed**:
     - Imports (except: `datetime`, `math`, `collections`)
     - File I/O (`open`, `read`, `write`, path operations)
     - Network (`socket`, `urllib`, `requests`, etc.)
     - Exec/eval (`exec`, `eval`, `compile`)
     - System access (`os`, `sys`, `subprocess`)
   - **Capped**: Max 10,000 iterations per loop
3. **Execute** in restricted namespace (only our data accessors available)
4. **Timeout**: 5 seconds max execution

**On validation failure**: Return error, don't execute, log for analysis.

**Future V2**: Consider WebAssembly sandbox for stronger isolation.

### Assumptions

| Assumption | Status | Validation Path |
|------------|--------|-----------------|
| 60-70% of queries are structured | No citation | Measure query distribution post-launch |
| AST allowlist catches dangerous patterns | Unvalidated | Security review + fuzzing |
| 5-second timeout sufficient | Reasonable | Adjust based on real query complexity |
| Schema descriptions don't leak identity | Probable | Audit schema generation |

### V1 Implementation Priority

**Phase A of M3 (Privacy milestone)** — implement first because:
- Lower complexity than semantic projection
- Zero data exposure proves concept immediately
- Builds local execution infrastructure needed for projection
- Query patterns inform knowledge graph schema

---

## Novel Approach 3: P2P Query Mixing (V2+)

### The Idea

Multiple users' queries get pooled and shuffled before hitting the LLM. The LLM sees queries but can't attribute them to individuals.

```
User A: query about calendar
User B: query about email        →  Mixer  →  Batch to LLM  →  Unmixer  →  Responses
User C: query about contacts
```

### Why It's Interesting

- Anonymity through crowd
- No single point sees both identity and query content
- Mixer knows who asked but not responses
- LLM sees queries and responses but not who

### Limitations (V2+ Scope)

| Limitation | Severity | Mitigation |
|------------|----------|------------|
| Need critical mass of simultaneous users | High | Defer to V2 when user base exists |
| Queries should be similar enough to batch | Medium | Batch by query type |
| Timing attacks if low activity | High | Add dummy queries during low periods |
| Trust in mixer | High | Distributed mixing protocol needed |

### Alternative: Distributed Entity Typing

Use P2P network specifically for semantic projection typing:

```
Query: "Meeting with John Smith about Acme merger"

Local extraction: [John Smith], [Acme], [merger]

P2P distribution:
  → Peer 1 sees: "John Smith" → votes: PERSON
  → Peer 2 sees: "Acme" → votes: ORGANIZATION
  → Peer 3 sees: "merger" → votes: BUSINESS_EVENT

Result: Type classifications without any peer seeing full query
```

### V2+ Implementation

Requires multiple users. Not applicable for V1 dogfooding.

---

## V1 Privacy Architecture

```
┌─────────────────────────────────────────────┐
│                  YOUR QUERY                  │
└─────────────────────────────────────────────┘
                      ↓
         ┌──────────────────────┐
         │   Query Classifier   │
         │   (local, simple)    │
         └──────────────────────┘
                      ↓
    ┌─────────────────┼─────────────────┐
    ↓                 ↓                 ↓
STRUCTURED        REASONING         CONTENT
(filter, search,  (prioritize,      (summarize,
 schedule)        plan, advise)      draft)
    ↓                 ↓                 ↓
┌─────────┐    ┌─────────────┐    ┌─────────┐
│ LLM as  │    │  Semantic   │    │ Local   │
│Compiler │    │ Projection  │    │  LLM    │
└─────────┘    └─────────────┘    └─────────┘
    ↓                 ↓                 ↓
Cloud LLM        Cloud LLM         Ollama
returns          reasons on        processes
LOGIC            SHADOWS           CONTENT
    ↓                 ↓                 ↓
Run locally      Rehydrate         Direct
                 locally           response
    ↓                 ↓                 ↓
└─────────────────────┼─────────────────┘
                      ↓
              ┌──────────────┐
              │   RESPONSE   │
              └──────────────┘
```

---

## Implementation Sequence

### V1 (M3: Privacy Milestone)

**Phase A: LLM as Compiler**
- Query classifier (structured vs reasoning vs content)
- Schema description generator (describe data shapes to LLM)
- Code execution sandbox (run LLM-generated logic safely)
- Test with todo + calendar data

**Phase B: Semantic Projection**
- Entity extraction from queries
- Local knowledge graph (contacts, topics, relationships)
- Projection function (real → typed placeholders)
- Rehydration function (placeholder responses → real entities)
- Test with calendar events, todo items

**Phase C: Local LLM Fallback**
- Ollama integration
- Content task routing
- Quality comparison logging (to validate tradeoffs)

### PRD Integration

| PRD Milestone | Privacy Phase | What's Tested |
|---------------|---------------|---------------|
| M1-M2 | — | Low sensitivity data only |
| M3 | A: Compiler | Structured queries with zero exposure |
| M3 | B: Projection | Reasoning queries with shadow data |
| M3 | C: Local LLM | Content tasks with local processing |
| M5-M6 | Full | Email + contacts (high sensitivity) |

**Gate**: M5 (Email) only proceeds after M3 privacy architecture validated.

### V2+

- P2P query mixing (when users exist)
- Distributed entity typing
- Federated learning across SymPAL instances
- Advanced correlation attack mitigations

---

## Correlation Attack Mitigations

### The Risk

Even with perfect projection, query patterns over time can be correlated:
- Same type combinations appearing together
- Query timing patterns
- Response latency signatures

### V1 Mitigations

| Mitigation | Implementation | Privacy Gain |
|------------|----------------|--------------|
| Token rotation | Daily epoch marker in placeholders | Prevents cross-session correlation |
| Query batching | Batch 2-5 queries when possible | Obscures individual query patterns |
| Timing noise | 100-500ms random delay | Defeats timing analysis |

### V2 Mitigations (If Needed)

| Mitigation | Implementation | Privacy Gain |
|------------|----------------|--------------|
| Query padding | Add plausible dummy queries | Obscures true query volume |
| Multi-provider distribution | Split queries across providers | No single provider sees all patterns |
| Decoy sessions | Periodic automated query bursts | Camouflages real usage patterns |

### Accepted Limitation

A motivated nation-state with long observation windows and access to provider logs may still correlate patterns. **Goal is practical obscurity, not perfect anonymity.**

This is acceptable because:
- Most users aren't nation-state targets
- Provider passive accumulation is the primary threat
- Even partial correlation requires significant effort

---

## Quality Measurement Methodology

### Why Measurement Matters

Privacy approaches are worthless if they destroy utility. We must measure quality impact empirically.

### Metrics

| Dimension | Metric | Target | How Measured |
|-----------|--------|--------|--------------|
| Task completion | Binary success | >95% | Did the query return a usable answer? |
| Reasoning quality | 1-5 subjective scale | >3.5 average | User rating on sample |
| Latency | End-to-end time | <5s simple, <15s complex | Instrumentation |
| Routing accuracy | Correct tier? | >80% | Manual audit of sample |
| Compiler validity | Code executes? | >90% | Execution success rate |
| Projection fidelity | Rehydration correct? | >95% | Spot-check sample |

### Protocol

1. **Log** all queries, tier decisions, execution results (locally)
2. **Weekly sample** 20 queries for manual quality rating
3. **Compare** to baseline (raw Claude API with full data)
4. **Adjust** thresholds based on findings

### Baseline Comparison

For each query type, run same query through:
- Raw LLM (full data exposure) — quality ceiling
- Privacy tier (compiler/projection/local) — actual

Quality delta should be <10% for projection, <20% for local. Compiler should be ~0% (logic is logic).

---

## Entity Taxonomy Specification

### Purpose

Standardize the type vocabulary for semantic projection. Balance expressiveness with privacy (too specific = identifying).

### Categories

| Category | Types | Example Placeholder |
|----------|-------|---------------------|
| PERSON | colleague, friend, family, acquaintance, vendor, client, executive | `[PERSON:colleague,senior]` |
| ORGANIZATION | employer, client, vendor, partner, government | `[ORG:client]` |
| PROJECT | business, personal, learning, health, financial | `[PROJECT:business,deadline-driven]` |
| LOCATION | home, work, travel, event | `[LOCATION:work]` |
| TIME | recent, upcoming, deadline, recurring, historical | `[TIME:upcoming,deadline]` |
| PRIORITY | urgent, important, routine, low | `[PRIORITY:urgent]` |
| TOPIC | business, personal, health, financial, legal, technical | `[TOPIC:business,sensitive]` |

### Rules

1. **Max 3 type modifiers** per entity — more is likely identifying
2. **Coarsen rare combinations** — if a type combo appears <5 times in your data, use parent type
3. **No proper nouns** ever in types — "startup" not "YCombinator-company"
4. **Daily token rotation** — placeholder tokens reset each epoch
5. **Session-scoped mappings** — rehydration maps valid only within session

### Examples

| Real Entity | Projection | Why |
|-------------|------------|-----|
| "John Smith, CEO of Acme" | `[PERSON:colleague,senior,frequent]` | Role + relationship, not identity |
| "Dr. Sarah Chen, my therapist" | `[PERSON:professional,health]` | Coarsened — "therapist" alone might be identifying |
| "Project Nightingale Q4 launch" | `[PROJECT:business,deadline-driven]` | Generic business project |
| "mom's birthday party" | `[PERSON:family] + [EVENT:personal,upcoming]` | Split into typed components |

---

## Open Questions

### Resolved in This Version

| Question | Resolution |
|----------|------------|
| Compiler output safety | AST sandbox with allowlist (see [Sandbox Architecture](#sandbox-architecture)) |
| Cross-query correlation | Token rotation + batching + timing noise (see [Correlation Attack Mitigations](#correlation-attack-mitigations)) |
| Quality measurement | Protocol defined (see [Quality Measurement Methodology](#quality-measurement-methodology)) |
| Projection granularity | Max 3 modifiers, coarsen rare combos (see [Entity Taxonomy](#entity-taxonomy-specification)) |

### Still Open

1. **Query classification accuracy**: How reliably can we route queries to the right tier? Requires implementation to measure.

2. **Knowledge graph bootstrapping**: How does the graph get populated initially? Options:
   - User effort (explicit classification)
   - Inference from existing data
   - Hybrid with user confirmation

   Defer decision until Phase B.

3. **Local LLM quality floor**: At what quality level does local LLM become unusable? Need empirical testing with Ollama/Llama models.

4. **Provider trust verification**: How do we verify providers aren't doing undisclosed analysis? May require cryptographic approaches (V2+).

---

## Relationship to PRD Challenges

These innovations are the response to:
- **Challenge 1**: Privacy layer is a black box → Now defined as three-tier architecture with explicit threat model
- **Challenge 2**: Quality degradation may be fantasy → Honest tradeoffs by task type with measurement protocol

---

## References

### Privacy & Anonymization

- arXiv:2502.15233 — Pseudonymization techniques with typed placeholders
- "A Survey on Privacy-Preserving Techniques for NLP" — context-preserving anonymization approaches
- k-anonymity, l-diversity literature — foundational privacy concepts

### NL2SQL & Code Generation

- AWS Athena, Oracle APEX, Microsoft Copilot — production NL2SQL systems
- "A Survey on Text-to-SQL Parsing" — comprehensive NL2SQL review
- StarCoder, CodeLlama — open code generation models

### LLM Security

- "Do Users Write More Insecure Code with AI Assistants?" — 12-65% vulnerability rate
- "An Empirical Study on LLM Hallucination in Code Generation" — 44% package hallucination
- OWASP LLM Top 10 — security considerations for LLM applications

### Sandbox & Isolation

- Python RestrictedPython — AST-based sandboxing
- WebAssembly isolation — stronger sandbox option
- seccomp, AppArmor — OS-level isolation (may be relevant for V2)

### Correlation Attacks

- Differential privacy literature — formal privacy guarantees
- Traffic analysis research — timing and pattern correlation
- "Website Fingerprinting" literature — analogous correlation attacks

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 0.1.0 | 2026-01-18 | Initial concepts: Semantic Projection, LLM as Compiler, P2P Mixing |
| 0.2.0 | 2026-01-18 | Added critical analysis: threat model, prior art acknowledgment, failure modes, assumptions, sandbox spec, correlation mitigations, quality measurement, entity taxonomy |

---

*Created: 2026-01-18*
*Last updated: 2026-01-18*
