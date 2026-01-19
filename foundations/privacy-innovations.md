# Privacy Innovations — Novel Approaches for SymPAL

**Version**: 3.0.0
**Date**: 2026-01-19
**Status**: Architecture finalized, awaiting implementation
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
| LLM Provider (passive) | Profile building, data monetization | Ephemeral Slots — single-use placeholders defeat entity-level profiling |
| LLM Provider (active) | Deanonymization via correlation | Single-use randomness defeats entity correlation; behavioral profiling remains possible |

**Core assumption**: Big tech has incentives to treat users as data products. They will do this. Design assumes providers are honest-but-curious (follow protocols, but analyze everything they legally can).

### Explicitly Out of Scope

| Adversary | Rationale |
|-----------|-----------|
| Legal/subpoena compulsion | If legal proceedings target us, privacy layer isn't the defense |
| Nation-state targeted attack | Different threat class; not our V1 problem |
| Local device compromise | Out of scope for V1; future concern for plugin architecture |
| Provider collusion | If provider actively cooperates with adversary targeting you specifically |

**Accepted limitation**: A motivated nation-state with long observation may correlate patterns. Goal is practical obscurity, not perfect anonymity.

### Threat Model Assumptions

- LLM providers are honest-but-curious (follow protocols, but analyze everything they legally can)
- TLS is not broken
- User's local environment is trusted
- Attackers don't have access to our source code's internal mappings

---

## Core Insight

**The false dichotomy**: Either give the LLM clean context and it builds a profile, OR fracture the context and quality suffers.

**The solution**: Give the LLM perfect, rich context in a way that **defeats entity-level profiling** — the provider cannot link specific people, projects, or relationships across queries.

The inspiration is the **one-time pad** from cryptography: single-use randomness. For every API call, we replace all private entities with brand-new, single-use, randomly generated placeholders. Then we give the LLM a "legend" that defines what those placeholders mean — for that query only.

The LLM gets full context. The provider gets disconnected, single-use pairs that cannot be correlated to build an entity graph.

**Honest limitation**: Account-level behavioral profiling (task types, timing patterns, legend semantic content) remains possible. Ephemeral Slots protects *who and what*, not *how you work*. See [Limitations & Residual Risks](#limitations--residual-risks).

---

## Novel Approach 1: Ephemeral Slots (Primary V1 Privacy Layer)

### The Idea

For every single API call, replace all private entities with brand-new, single-use, randomly generated placeholders. Include a "legend" at the start of the prompt that defines what those placeholders mean — for that query only.

**User request**: "Draft an email to my manager Jane about the Project Phoenix delay."

**What gets sent to LLM**:
```
You are a helpful assistant. For this request, use the following legend:

Legend:
- {{PERSON_kestrel}} is my manager.
- {{PROJECT_sparrow}} is a high-priority internal project.

Task:
Draft a professional email to {{PERSON_kestrel}} apologizing for a delay on {{PROJECT_sparrow}}.
```

**LLM response**:
```
Dear {{PERSON_kestrel}},

I'm writing to update you on the timeline for {{PROJECT_sparrow}}...
```

**Local fulfillment**: Structured output with slot tokens, then local substitution (see [Rehydration Requirements](#rehydration-requirements-v1)).

### Why It's Different

The key insight: **The next time you ask about Jane, she might be `{{PERSON_falcon}}`. The next time you ask about Project Phoenix, it might be `{{PROJECT_finch}}`.** The slots are ephemeral — they live for one query and then cease to exist.

| Approach | What Provider Sees Over Time | Can Build Entity Graph? |
|----------|------------------------------|------------------------|
| Direct API | "Jane", "Project Phoenix" repeatedly | Yes — full profile |
| Pseudonymization | "Person_A7", "Project_B3" repeatedly | Yes — consistent IDs link queries |
| Semantic Projection | `[PERSON:manager]`, `[PROJECT:high-priority]` repeatedly | Partially — type patterns correlate |
| **Ephemeral Slots** | `kestrel→sparrow`, then `falcon→finch`, then `wren→robin`... | **No** — disconnected single-use pairs |

A profiler sees that `kestrel` is linked to `sparrow`. In the next query, they see `falcon` is linked to `finch`. They cannot know that `kestrel` and `falcon` are the same person. The relationship graph they build is a sea of disconnected, single-use pairs — **useless for entity-level profiling**.

**What the provider CAN still see** (account-level):
- Task types and query patterns ("this user does a lot of scheduling")
- Legend semantic content ("manager", "high-priority project", "strained relationship")
- Timing and frequency patterns
- API key / account correlation across all queries

Ephemeral Slots defeats entity graphs, not behavioral fingerprinting.

### How It Works

```
1. LOCAL ENTITY IDENTIFICATION
   ├── Input: "Draft email to manager Jane about Project Phoenix delay"
   ├── NER identifies: Jane (person), Project Phoenix (project)
   └── Entities tagged with minimal context needed

2. EPHEMERAL SLOT GENERATION (The Core)
   ├── For THIS API call only, generate random human-readable placeholders:
   │   ├── Jane → {{PERSON_kestrel}}
   │   └── Project Phoenix → {{PROJECT_sparrow}}
   ├── Slots are single-use — never reused across queries
   └── Next query about Jane uses completely different slot

3. LEGEND CONSTRUCTION
   ├── Build MINIMUM VIABLE context for each entity:
   │   ├── {{PERSON_kestrel}}: "is my manager" (not "direct manager at current employer")
   │   └── {{PROJECT_sparrow}}: "is a work project" (not "high-priority internal project")
   └── Escalate detail ONLY if quality degrades — never default to rich context

4. SELF-CONTAINED PROMPT
   ├── Legend + Task sent as single prompt
   ├── LLM has everything needed for full-context reasoning
   └── No multi-stage orchestration required

5. LOCAL FULFILLMENT
   ├── LLM returns structured response with slot tokens (JSON format)
   ├── Validate all returned slots exist in legend
   ├── Handle possessives, case variations, edge cases
   └── Substitute real values back into response
```

### Why Ephemeral Slots Over Semantic Projection (Design History)

During early design, we considered Semantic Projection — typed placeholders like `[PERSON:manager,senior]` with a consistent vocabulary. The comparison below explains why Ephemeral Slots was chosen instead.

| Dimension | Semantic Projection (considered) | Ephemeral Slots (chosen) |
|-----------|----------------------------------|--------------------------|
| Cross-query correlation | Possible via type patterns | **Defeated** — random single-use IDs |
| Context quality | Types only ("manager") | Full relationship ("is my manager") |
| Task coverage | Struggles with generation | Broad (but see limitations) |
| Infrastructure | Taxonomy, knowledge graph, rehydration | NER + random gen + structured output |
| API calls | 1 | 1 |
| Complexity | Medium | Lower |

**Why Semantic Projection was rejected**: The type vocabulary `[PERSON:manager,senior,frequent]` is consistent across queries. A provider can correlate "queries involving a senior manager" over time and build a partial profile.

**Why Ephemeral Slots wins**: `{{PERSON_kestrel}}` in one query has zero relationship to `{{PERSON_falcon}}` in the next. The provider cannot correlate entity identities across queries.

**Note**: Semantic Projection is available as an opt-in V2+ enhancement for users who want richer type vocabulary and accept the correlation tradeoff. See [V2+: Semantic Projection](#semantic-projection-type-based-shadows-v2-enhancement-to-ephemeral-slots).

### Legend Minimization

The legend reveals relationship context — **this is the primary remaining leakage channel**. Legend content can be minimized per task:

| Detail Level | Example | When to Use |
|--------------|---------|-------------|
| Detailed | "{{PERSON_kestrel}} is my direct manager at my current employer" | Complex reasoning needed |
| Standard | "{{PERSON_kestrel}} is my manager" | Most tasks |
| Minimal | "{{PERSON_kestrel}} is a professional superior" | Sensitive queries |
| Ultra-minimal | "{{PERSON_kestrel}} is someone I communicate professionally with" | Maximum privacy |

The legend is **task-specific, not person-specific**. You reveal only what this task requires.

#### Detail Escalation Decision Framework

**Default**: Always start with minimal or ultra-minimal legend.

**Escalate only when:**
1. **Response quality fails** — LLM produces generic/unhelpful response with minimal legend
2. **Task explicitly requires relationship context** — e.g., "advice about my relationship with my manager" needs relationship detail
3. **User explicitly requests richer context** — user overrides minimal default

**Escalation threshold**: If response doesn't meet task requirements after one attempt with minimal legend, retry with one level higher. Never jump straight to detailed.

| Task Type | Starting Level | Escalate If |
|-----------|---------------|-------------|
| Scheduling, reminders | Ultra-minimal | Response uses wrong pronouns or makes nonsensical suggestions |
| Professional email draft | Minimal | Tone is inappropriate for relationship (too formal/informal) |
| Career advice, conflict resolution | Standard | Response is generic, doesn't account for relationship dynamics |
| Sensitive personal decisions | Standard → Detailed | Response misses key context that changes the advice |

**Leakage risk**: Even slotted, "direct manager at my current employer, strained relationship" is sensitive information. Legend content is visible to provider and can be profiled at account level. Default to minimal; escalate detail only when quality requires it.

**M3 gate**: Legend minimization logic must be implemented and validated before M5 (email). If minimization consistently requires escalation to detailed level (>30% of queries), re-evaluate approach.

### Failure Modes & Limitations

| Failure Mode | Severity | Mitigation | Residual Risk |
|--------------|----------|------------|---------------|
| NER misses entity | **High** | User review before send (mandatory V1); deterministic redaction backup | Missed entities leak — nicknames, code names, implicit refs |
| Legend too detailed | Medium | Default to minimal; escalate only when quality drops | Over-specification leaks sensitive context |
| Rehydration fails (paraphrase/omission) | Medium | Require structured JSON output; validate all slots returned | Quality degradation on complex generation |
| Placeholder collision/injection | Medium | Escaping + strict format validation; reject unbound slots | User input mimicking slots; hallucinated slots |
| Follow-up queries need stable identity | Medium | Route to Local LLM; warn user | Multi-turn conversations lose slot context |
| Intra-query correlation | Accepted | Within one query, LLM sees full relationship graph | Unavoidable for quality; isolated to single query |
| Slot collision | Very Low | Use sufficiently random generation | Astronomically unlikely |

### Limitations & Residual Risks

**What Ephemeral Slots protects against:**
- Cross-query correlation of named entities (names, specific identifiers)
- Building explicit relationship graphs from entity names

**What Ephemeral Slots does NOT fully protect against:**

| Risk | Category | Example |
|------|----------|---------|
| **Entity inference via unique descriptions** | Entity-level | "direct manager at current employer, strained relationship" is uniquely identifying even without the name |
| Behavioral profiling | Account-level | Task types, timing patterns, query frequency |
| Legend semantic profiling | Account-level | Recurring relationship types, role descriptions |
| Longitudinal queries | Functional | Follow-ups, reminders, multi-turn conversations |

**Critical insight**: Unique legend descriptions can enable **entity-level** profiling by description, not just account-level. "The person I report to who's been difficult lately" may be as identifying as "Jane Smith" in your context.

**Residual attack surface:**
The legend is the primary leakage channel. The defense is **minimum viable legend** — use the least specific description that maintains quality. "is my manager" not "is my direct manager at my current employer who I've had conflicts with."

**Default rule**: Start with minimal legend. Escalate detail only when quality degrades. Never default to rich context.

**Known failure cases (route to Local LLM instead):**
- Follow-up queries referencing previous entities ("What about that project we discussed?")
- Reminders with entity context ("Remind me about this meeting")
- Longitudinal summaries ("What have I discussed with Jane this month?")
- Multi-turn conversations requiring consistent entity references

### Assumptions

| Assumption | Status | Validation Path |
|------------|--------|-----------------|
| NER can identify entities reliably | **Medium confidence overall** — see breakdown below | Personal data includes nicknames, code names, implicit refs — will miss some |
| Random slot generation is sufficient | High confidence | Cryptographic randomness available |
| Structured output rehydration works | Medium confidence | Requires JSON format + validation; paraphrases/omissions possible |
| Legend context sufficient for quality | To validate | Dogfooding feedback |
| User will review entities before send | **Hypothesis — major UX cost** | See below |

**NER Confidence by Data Sensitivity:**

| Data Type | NER Confidence | Required Accuracy | Rationale |
|-----------|----------------|-------------------|-----------|
| Calendar (titles, attendees) | Medium-High | 80% | Structured data, predictable formats |
| Todos (simple text) | Medium | 75% | Less structure, but user-written |
| **Email (bodies, subjects)** | **Medium** | **90%** | Free-form text, nicknames, implicit refs — highest leakage risk |
| Contacts | High | 85% | Structured fields |

**M5 Gate** (email): NER must achieve 90% accuracy on email content before email capability is enabled. If NER falls below this threshold on email data, require per-query entity review for email routes (higher friction, but necessary for high-sensitivity data).

**V1 Tradeoff: Entity Review Friction**

Mandatory entity review before every query is a significant UX cost. Options:

| Approach | Privacy | UX | Recommendation |
|----------|---------|-----|----------------|
| Review every query | High | Poor | Not recommended for V1 |
| Review first mention only | Medium | Good | **V1 default** — review when entity first seen, then trust NER |
| Review sensitive routes only | Medium | Good | Alternative — review for email/contacts, skip for calendar/todos |
| No review (trust NER) | Low | Best | Not recommended — leakage risk too high |

**V1 default**: Review on first mention. User confirms "Jane = Jane Smith (manager)" once. Future queries use that mapping without re-confirmation. User can always edit via Privacy Receipts.

**Gate for M5 (Email)**: If NER accuracy on email data < 90%, require per-query review for email routes. Measure in dogfooding.

### Rehydration Requirements (V1)

**Structured output mandatory for generation tasks:**
- Require JSON response format with explicit slot tokens
- Example: `{"body": "Dear {{PERSON_kestrel}}, ...", "slots_used": ["PERSON_kestrel"]}`
- Reject free-form prose for slot-containing responses

**Post-processing pipeline:**
1. Validate all returned slots exist in legend
2. Detect unbound/hallucinated slots → error, don't silently pass through
3. Handle possessives: `{{PERSON_kestrel}}'s` → `Jane's`
4. Handle case variations: normalize to canonical form
5. Escape user input that mimics slot patterns

**Fallback cascade when rehydration fails:**

```
1. RETRY WITH STRICTER PROMPT
   └── Add explicit instruction: "You MUST use only these slots: [list]"
   └── If still fails → step 2

2. ROUTE TO LOCAL LLM
   └── Re-run same query through Ollama with real entities
   └── User sees: "Falling back to local processing for privacy"
   └── If local unavailable → step 3

3. RETURN PARTIAL WITH WARNING
   └── Show response with unresolved slots visible: "Dear {{PERSON_kestrel}}, ..."
   └── User prompt: "Some names couldn't be filled in. Edit manually?"
   └── Never silently drop or guess
```

**User-visible failure**: The user must always know when rehydration fails. Silent failures are unacceptable — they erode trust and may expose slots in final output.

**Escaping specification:**
- User input containing `{{` must be escaped before prompt construction
- Model output validated against known slot IDs only
- Unrecognized slot patterns rejected or flagged

### Implementation

**Requirements**:
1. Entity detection (NER) with user review/correction flow
2. Random slot generator
3. Legend constructor — template-based, minimal entity metadata
4. Structured output schema (JSON with slot fields)
5. Response validation and rehydration pipeline

**What's NOT needed**:
- Type taxonomy
- Knowledge graph
- Cross-session consistency
- Multi-stage orchestration
- Multiple API calls

**Complexity note**: Rehydration is not trivial. Naive find-and-replace breaks on possessives, punctuation, paraphrases. Budget for structured output handling and edge cases.

### V1 Implementation Priority

**Phase B of M3 (Privacy milestone)** — Primary privacy layer for all generative/reasoning tasks.

---

### V1 Complement: Witness-Only Reasoning (Analytical Queries)

For analytical queries where you need advice based on patterns, not specific entities.

LLM never sees entities — only computed/derived facts from local analysis.

**How it works**:

Local engine computes a "witness set" of statements from your data:
```
Witness Set:
- "3 meetings next week with overlaps >30 minutes"
- "2 of those include attendee with role=client"
- "Priority ranking: client meetings > internal meetings"
- "Conflict density: high (multiple back-to-back days)"
```

LLM receives only the witness set, produces advice/summary. Never sees who, what, when specifically.

**When to use instead of Ephemeral Slots**:
- Query is purely analytical ("Am I overbooked?")
- No specific entities needed in response
- Maximum privacy required

**Relationship to Ephemeral Slots**:
- Ephemeral Slots: Full context via legend, entities in response
- Witness-Only: Aggregate facts only, no entities at all

Both are V1. Witness-Only is stronger privacy but narrower applicability.

---

### V1 Complement: Heuristic-First + Templates (No-LLM Path)

For a subset of common tasks, use local scoring rules + hard-coded response templates. No LLM call at all.

**Note**: This is a complement to DSL Compilation, not a replacement. The 40-50% DSL coverage estimate in the coverage table *includes* heuristic-first paths. Heuristics handle the simplest queries; DSL handles more complex structured queries.

**How it works**:

```
Query: "What should I prioritize today?"

Local computation:
├── priority_score = urgency * 2 + importance + role_weight
├── Sort items by score
└── Top item: "Review Q4 budget" (score: 8.5)

Template:
"Your top priority is {{item}}. It scores highest due to {{factors}}."

Output:
"Your top priority is reviewing the Q4 budget. It scores highest
 due to urgency and stakeholder importance."
```

No LLM call. Zero exposure. Instant response.

**Best for**: Simple queries that don't need LLM creativity:
- Prioritization ("What's most urgent?")
- Conflict detection ("Any scheduling conflicts?")
- Counts and summaries ("How many meetings this week?")
- Reminders ("What's due soon?")

**Implementation path**:
- V1: Build heuristics for top 10-15 common query patterns
- LLM only used when heuristics fail or user explicitly opts in
- Track which queries hit heuristics vs LLM to expand coverage

---

### The Ephemeral Slots Roadmap: From Technique to Paradigm

Ephemeral Slots isn't just a V1 technique — it's a foundational principle that scales from simple entity replacement to a complete paradigm for human-AI interaction. This roadmap shows how the concept evolves.

```
V1: Entity Replacement
    └── "Jane" → {{PERSON_kestrel}}
    └── Single-use random placeholders defeat profiling

V1.5: Dynamic Legend Optimization
    └── Legend detail adapts to task requirements
    └── "Just-enough-context" further hardens privacy

V2: Composable & Nested Slots
    └── Teams, hierarchies, dependencies as abstract structures
    └── Relational modeling without concrete knowledge

V2.5: Functional Slots
    └── Processes, workflows, rules as slots
    └── LLM reasons about your systems abstractly

V3: The Ephemeral Self
    └── Per-query digital twin from composable slots
    └── Full AGI power on complete life context
    └── "Ghost in the machine"
```

---

#### V1.5: Dynamic Legend Optimization

Make legends task-adaptive automatically, not manually.

**The insight**: Not all tasks require the same legend detail. A scheduling query needs "is a contact"; career advice needs "is my direct manager with a strained relationship."

**How it works**:

| Task Type | Legend Detail | Example |
|-----------|---------------|---------|
| Scheduling | Minimal | "{{PERSON_kestrel}} is a contact" |
| Professional email | Standard | "{{PERSON_kestrel}} is my manager" |
| Sensitive advice | Detailed | "{{PERSON_kestrel}} is my direct manager; we have a strained relationship" |

**Implementation approach**:
1. Task classifier determines required detail level
2. Rule-based initially: `scheduling → minimal`, `advice → detailed`
3. Learn from user feedback: "that response needed more context"
4. User override always available

**Why V1.5**: Refinement of existing V1 mechanism, not new architecture. Low risk, high value.

---

#### V2: Composable & Nested Slots

Slots can contain other slots. Represent entire systems abstractly.

**The insight**: Relationships themselves can be ephemeral. The LLM sees "A depends on B" without knowing what A or B are.

**Example legend**:
```
Legend:
- {{TEAM_sparrow}} is a team. Members: {{PERSON_kestrel}}, {{PERSON_finch}}.
- {{TEAM_condor}} is a team. Leader: {{PERSON_falcon}}.
- {{PROJECT_ibis}} is assigned to {{TEAM_sparrow}}.
- {{PROJECT_ibis}} has a dependency on {{TEAM_condor}}.

Task: "Draft a message from {{PERSON_kestrel}} to {{PERSON_falcon}}
       to inquire about the dependency status for {{PROJECT_ibis}}."
```

The LLM reasons about hierarchy, team membership, and project dependencies — all on abstract structures.

**What you can model**:
- Organizational hierarchies
- Project dependency graphs
- Family/social relationships
- Any relational structure

**Privacy consideration**: Complex legends reveal structural patterns. A legend with 10 nested slots reveals organizational complexity. Still far better than raw data, but worth noting.

**Mitigations**:
- Structural noise injection (fake relationships)
- Partial legends (only include what's needed)
- Complexity caps per query

**Why V2**: Significant conceptual extension, but still "just text" to the LLM. Tractable complexity.

---

#### V2.5: Functional Slots

Slots evolve from nouns (entities) to verbs (processes, workflows, rules).

**The insight**: You can teach the LLM your business logic without revealing what that logic actually governs.

**Example legend**:
```
Legend:
- {{PROCESS_aurora}} is our mandatory code review process.
  Requires two approvals from senior members.
- {{PERSON_kestrel}} is a junior engineer.
- {{PERSON_falcon}} is a senior engineer.
- {{TICKET_osprey}} is a high-priority bug fix.

Task: "Generate a step-by-step plan for {{PERSON_kestrel}} to get
       {{TICKET_osprey}} deployed, ensuring {{PROCESS_aurora}} is followed."
```

The LLM generates a checklist, suggests who to contact, structures a release plan — without knowing what your code review process actually is, what the bug was, or who the engineers are.

**What functional slots capture**:
- Inputs/outputs
- Constraints/rules
- Roles involved
- Steps/sequence

**Implementation consideration**: May need structured legend syntax or process modeling layer. More complex than entity slots.

**Why V2.5**: Conceptual leap from nouns to verbs. Requires process representation framework. High value for workflow automation.

---

#### V3: The Ephemeral Self

The culmination: per-query digital twin constructed entirely from composable and functional ephemeral slots.

**The insight**: For complex, holistic queries, SymPAL generates a complete, disposable life-snapshot. The LLM performs deeply personal strategic analysis on a "complete fiction."

**Example**:

Query: "Given my current workload, personal commitments, and career goals, should I accept the offer to lead the new 'Odyssey' initiative?"

```
The Ephemeral Self Prompt:

You are a world-class executive coach. Analyze the following
life-snapshot and provide your recommendation.

Legend:
- {{PERSON_self}} is the user, whose career goal is {{GOAL_leadership}}.
- {{PROJECT_sparrow}} and {{PROJECT_finch}} are current high-effort work projects.
- {{COMMITMENT_alpha}} is a recurring, non-negotiable personal commitment.
- {{INITIATIVE_omega}} is a new leadership opportunity.
  It would advance {{GOAL_leadership}} but significantly increase workload.
- {{RELATIONSHIP_kestrel}} is a key professional relationship
  that would be enhanced by accepting {{INITIATIVE_omega}}.
- {{CONSTRAINT_zeta}} is a hard constraint: "must not work on weekends."

Task:
"Should {{PERSON_self}} accept {{INITIATIVE_omega}}?
 Analyze tradeoffs considering:
 - Impact on {{PROJECT_sparrow}} and {{PROJECT_finch}}
 - Potential conflicts with {{COMMITMENT_alpha}}
 - Effect on {{RELATIONSHIP_kestrel}}
 - Whether it can be accomplished without violating {{CONSTRAINT_zeta}}"
```

The LLM performs holistic life analysis — on a ghost. The provider learns nothing of your life, work, or goals.

**Construction: Three-Step Pipeline**

The Ephemeral Self is a query-specific, temporary, abstract view of a persistent **Local Knowledge Graph (KG)** — the ground truth of your life, built over time from emails, calendar, files, and notes.

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│ 1. SEED         │ ──► │ 2. TRAVERSE     │ ──► │ 3. EPHEMERALIZE │
│ IDENTIFICATION  │     │ LOCAL GRAPH     │     │                 │
└─────────────────┘     └─────────────────┘     └─────────────────┘
     ↓                        ↓                        ↓
Local NER finds          Hop from seeds           Generate legend:
seed entities            to find relevant         random slots for
from query               sub-graph                all nodes & edges
```

1. **Seed Identification**: Local NER extracts seed entities from query. "Should I lead the 'Odyssey' initiative?" → seeds: `[me]`, `[lead]`, `['Odyssey' initiative]`

2. **Local Graph Traversal**: Starting from seeds, traverse KG to find relevant sub-graph:
   - From `[me]` → connected `[PROJECT]` nodes (workload), `[EVENT]` nodes (commitments)
   - From `['Odyssey' initiative]` → properties (type: leadership_opportunity), relationships (proposed_by: `[PERSON:Jane]`)
   - Result: complete, context-rich "slice" of your life for this query

3. **Ephemeralization**: Iterate through every node and edge in sub-graph, assign single-use random slot names. Jane → `{{PERSON_kestrel}}`, proposed_by → `{{ACTION_zeta}}`, etc.

The Ephemeral Self = the resulting legend (textual description of temporary, anonymized graph).

**Consistency Model: Burn After Reading**

| Requirement | Mechanism |
|-------------|-----------|
| Same entity → same slot *within* query | Prompt-scoped mapping object (dictionary) held in memory during construction |
| Different slots *across* queries | Mapping object destroyed the moment API call completes. Next query starts blank. |

No persistent link between `{{PERSON_kestrel}}` from query 1 and whatever slot Jane gets in query 2.

**Structural Fingerprinting Defenses**

A unique graph structure could itself become an identifying fingerprint. Three defenses:

| Defense | Mechanism | Example |
|---------|-----------|---------|
| **Coarsening** | Summarize structure instead of enumerating | "15 tasks, 4 high-priority" instead of 15 task nodes |
| **Noise Injection** | Add plausible fake nodes/edges | Fake project in workload, invented dependency |
| **Pruning** | Minimum necessary structure | Budget query doesn't need team reporting structure |

**Coarsening** (abstraction): Zoom out. Instead of detailed task dependency graph, legend states: `{{PROJECT_sparrow}} has 15 associated tasks, 4 high-priority`. Preserves reasoning context ("complex project") while destroying unique fingerprint.

**Noise Injection** (camouflage): Add fictitious entities. Fake project in workload. False dependency between real but unrelated projects. Team of 5 when there are 4. Local client filters output knowing which were decoys.

**Pruning** (minimum structure): Default to shallow traversal. Query about Project Phoenix's budget doesn't need team's other projects or reporting hierarchy.

*Quality caveat*: Noise injection risks LLM reasoning about fakes. Mitigation: restrict noise to adjacent/context nodes rather than nodes central to reasoning. Or: inject noise only for high-fingerprinting-risk queries (many entities, deep traversal).

**Dependency**: Local Knowledge Graph construction (entity linking, relationship extraction from emails/calendar/files) is a V2-V3 project unto itself. KG schema design is prerequisite.

**Why V3**: Requires robust V2-V2.5 infrastructure (Composable Slots, Functional Slots). High complexity, but transformative value.

---

#### The Vision: Ghost in the Machine

This is where the Ephemeral Slots paradigm leads:

> **Full AGI power on the most intimate aspects of your life — with complete trust.**

The system is designed, from its V1 seed, to make you a ghost in the machine. The provider sees phantoms that evaporate after each query. You get personalized, holistic, strategic AI assistance without surveillance.

This is not just privacy. It's a new form of human-AI interaction where the boundary between "what AI knows" and "what AI can help with" is completely decoupled.

**Privacy escalation across versions**:

| Version | What Provider Sees | What Provider Can Build |
|---------|-------------------|------------------------|
| V1 | Random slots + legend | Nothing — single-use |
| V1.5 | Minimal random slots | Nothing — even less structure |
| V2 | Abstract relational structure | Structural patterns only |
| V2.5 | Abstract process structure | Workflow complexity only |
| V3 | Complete abstract life model | A ghost — nothing actionable |

The more powerful the capability, the more complete the privacy. That's the Ephemeral Slots paradigm.

---

## Novel Approach 2: Query Compilation (Zero-Data Exposure)

### The Umbrella Insight

**The LLM should produce *programs*, not *answers*.**

For any task that can be formalized — whether as structured queries OR as processing instructions — the cloud LLM generates the program, and a local executor runs it against real data. Zero exposure.

Two compilation paths:
- **Path A: DSL Compilation** — For structured queries (filter, search, aggregate). Cloud LLM generates SymQL, local interpreter executes.
- **Path B: Prompt-as-Program** — For fuzzy data-processing (summarize, extract, classify). Cloud LLM generates an optimal prompt, local LLM executes.

Same principle, different target runtimes.

---

### Path A: DSL Compilation (Structured Queries)

#### The Idea

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

#### Why It's Different

- Zero data exposure (not reduced — zero)
- LLM output is reusable (same logic applies to future queries)
- Builds a library of local capabilities over time

#### Prior Art & What's Actually Novel

**Established techniques**:
- NL2SQL is mature — AWS, Oracle, Microsoft have production systems
- Text-to-code generation (Codex, StarCoder, etc.)
- Schema-based query generation

**What's novel in our approach**:
- **Privacy framing**: Existing systems optimize for accuracy; we optimize for zero data exposure
- **Personal data context**: Schema descriptions include relationship semantics, not just column types
- **Hybrid routing**: Compiler for structured, Ephemeral Slots for reasoning, local for content

**Honest framing**: This is proven technology in a novel application context. The innovation is the routing architecture, not the compilation itself.

#### What DSL Compilation Works For

**Good fit (estimated 40-50% of queries)**:
- Scheduling and conflict detection
- Filtering and search
- Priority scoring based on rules
- Pattern matching
- Data transformation

**Note**: The 60-70% "zero-exposure" V1 target includes DSL (40-50%) + PaaP (10-20%) + Local-only (10-20%). DSL alone doesn't cover the full target.

**Poor fit for DSL** (see Path B: Prompt-as-Program):
- Summarization
- Classification/sentiment
- Extraction

#### Example

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

**LLM returns** (TypeScript for Deno sandbox):
```typescript
// Deno sandbox: no network, no filesystem, data passed via stdin
interface CalendarEvent { start: Date; attendees: string[]; }
interface Email { from: string; to: string[]; date: Date; }

function findNeglectedMeetingContacts(
  calendar: CalendarEvent[],
  emails: Email[],
  dateRange: { start: Date; end: Date },
  inactiveDays: number
): CalendarEvent[] {
  // Filter calendar events in range
  const upcoming = calendar.filter(e =>
    e.start >= dateRange.start && e.start <= dateRange.end
  );

  // Get all attendee emails
  const attendees = upcoming.flatMap(e => e.attendees);

  // Find last email date per contact
  const lastContact = new Map<string, Date>();
  for (const email of emails) {
    for (const addr of [email.from, ...email.to]) {
      const prev = lastContact.get(addr);
      if (!prev || email.date > prev) {
        lastContact.set(addr, email.date);
      }
    }
  }

  // Filter to inactive
  const cutoff = new Date(Date.now() - inactiveDays * 24 * 60 * 60 * 1000);
  const neglected = new Set(
    attendees.filter(a => (lastContact.get(a) ?? new Date(0)) < cutoff)
  );

  return upcoming.filter(e => e.attendees.some(a => neglected.has(a)));
}
```

**User runs locally** with real data via Deno subprocess. LLM never saw calendar or emails.

#### Hybrid Compilation: Decomposing Semi-Structured Queries (V1)

The basic compiler model is binary: a query is either fully structured (compiler handles it) or it isn't (route to Ephemeral Slots/Local). Hybrid Compilation handles the middle ground.

**The insight**: Many queries combine a structured operation (filter, select) with a content operation (summarize, draft). Decompose them.

**Example**:

User query: "Summarize my urgent emails about Project Titan"

**Without hybrid compilation**: Entire query routes to Ephemeral Slots or Local (content task). All email content exposed to that tier.

**With hybrid compilation**:

```
Step 1: Compiler
├── LLM generates: getEmails(project='Project Titan', priority='urgent')
├── Code runs locally
└── Returns: 3 specific emails (filtered from 500)

Step 2: Local LLM
├── Only the 3 filtered emails go to Ollama
├── Summarization runs locally
└── Returns: Summary
```

**Privacy gain**: Instead of exposing 500 emails to summarization, expose 3. The compiler does maximum filtering at zero exposure before handing off.

**How it works**:

1. **Query analysis**: Classifier identifies decomposable queries (structured filter + content operation)
2. **Compile filter step**: LLM generates filter code, runs in sandbox
3. **Hand off filtered results**: Only filtered data goes to Ephemeral Slots or Local tier
4. **Content operation**: Summarize/draft/analyze on minimal dataset

**What qualifies for hybrid compilation**:

| Query Pattern | Filter (Compiler) | Content (Ephemeral Slots/Local) |
|---------------|-------------------|---------------------------|
| "Summarize emails about X" | Filter by topic/sender/date | Summarize filtered set |
| "Draft reply to urgent messages" | Filter by priority/recency | Draft based on filtered |
| "What should I prioritize from Y?" | Filter by source/project | Reasoning on filtered set |
| "Find and explain Z" | Filter/search | Explain filtered results |

**What doesn't qualify** (route normally):
- Pure structured: "How many meetings next week?" → Compiler only
- Pure content: "Summarize this document" → Local only
- Pure reasoning: "What should I focus on today?" → Ephemeral Slots only

**Implementation notes**:

- Classifier needs to detect "filter + content" pattern
- Pipeline orchestration: compile → execute → route filtered results
- Error handling: if compile step fails, fall back to routing whole query
- Latency: two-step adds ~2-3s; acceptable given privacy gain

**V1 scope**: Start with explicit patterns (emails about X, meetings with Y). Expand detection as confidence grows.

#### Hybrid Compilation State Machine

Formal definition of the hybrid pipeline as a state machine.

```
                    ┌─────────────────────────────────────────┐
                    │              START                       │
                    │  Input: user_query, data_context         │
                    └─────────────────────────────────────────┘
                                       │
                                       ▼
                    ┌─────────────────────────────────────────┐
                    │           CLASSIFY_HYBRID               │
                    │  Is query decomposable?                  │
                    │  Check: filter_pattern + content_pattern │
                    └─────────────────────────────────────────┘
                           │                    │
                    [YES: decomposable]    [NO: route normally]
                           │                    │
                           ▼                    └──────► EXIT to normal routing
                    ┌─────────────────────────────────────────┐
                    │          COMPILE_FILTER                  │
                    │  Send filter description to cloud LLM    │
                    │  Receive: DSL/code for filter operation  │
                    └─────────────────────────────────────────┘
                           │                    │
                    [compile success]     [compile fail]
                           │                    │
                           ▼                    └──────► FALLBACK_FULL_ROUTE
                    ┌─────────────────────────────────────────┐
                    │          EXECUTE_FILTER                  │
                    │  Run compiled code in sandbox            │
                    │  Result: filtered_data subset            │
                    └─────────────────────────────────────────┘
                           │                    │
                    [execution success]   [execution fail]
                           │                    │
                           ▼                    └──────► FALLBACK_FULL_ROUTE
                    ┌─────────────────────────────────────────┐
                    │          CHECK_FILTER_RESULT             │
                    │  Is filtered_data non-empty?             │
                    │  Is reduction meaningful? (< 50% of full)│
                    └─────────────────────────────────────────┘
                           │                    │
                    [meaningful reduction] [no benefit]
                           │                    │
                           ▼                    └──────► FALLBACK_FULL_ROUTE
                    ┌─────────────────────────────────────────┐
                    │        ROUTE_CONTENT_OPERATION           │
                    │  Classify content operation type:        │
                    │  SUMMARIZE → Local LLM                   │
                    │  DRAFT → Local LLM or Ephemeral Slots    │
                    │  REASON → Ephemeral Slots                │
                    └─────────────────────────────────────────┘
                           │
                           ▼
                    ┌─────────────────────────────────────────┐
                    │        EXECUTE_CONTENT                   │
                    │  Run content operation on filtered_data  │
                    │  Only filtered subset exposed            │
                    └─────────────────────────────────────────┘
                           │
                           ▼
                    ┌─────────────────────────────────────────┐
                    │              END                         │
                    │  Output: result to user                  │
                    └─────────────────────────────────────────┘


FALLBACK_FULL_ROUTE:
  ┌─────────────────────────────────────────┐
  │  Hybrid compilation failed or no benefit │
  │  Route full query to standard tier:      │
  │  - Local LLM (content tasks)             │
  │  - Ephemeral Slots (reasoning tasks)     │
  │  Log: "hybrid_fallback" for analysis     │
  └─────────────────────────────────────────┘
```

**State Transitions:**

| From State | Condition | To State |
|------------|-----------|----------|
| START | Query received | CLASSIFY_HYBRID |
| CLASSIFY_HYBRID | Has filter + content pattern | COMPILE_FILTER |
| CLASSIFY_HYBRID | No decomposition possible | EXIT (normal routing) |
| COMPILE_FILTER | DSL generated successfully | EXECUTE_FILTER |
| COMPILE_FILTER | Compilation failed | FALLBACK_FULL_ROUTE |
| EXECUTE_FILTER | Code executed, results returned | CHECK_FILTER_RESULT |
| EXECUTE_FILTER | Execution error | FALLBACK_FULL_ROUTE |
| CHECK_FILTER_RESULT | Filtered < 50% of original | ROUTE_CONTENT_OPERATION |
| CHECK_FILTER_RESULT | No meaningful reduction | FALLBACK_FULL_ROUTE |
| ROUTE_CONTENT_OPERATION | Content type classified | EXECUTE_CONTENT |
| EXECUTE_CONTENT | Complete | END |
| FALLBACK_FULL_ROUTE | Always | Standard routing path |

**Key invariants:**
1. Filter step NEVER sees content — only structural description
2. Content step only sees filtered subset
3. Any failure falls back to standard (safe) routing
4. "Meaningful reduction" threshold (50%) prevents wasteful decomposition

#### Failure Modes & Limitations (DSL)

| Failure Mode | Severity | Mitigation | Residual Risk |
|--------------|----------|------------|---------------|
| DSL too limited for query | Medium | Fallback to Deno sandbox, then Ephemeral Slots | May require multiple round trips |
| Logic errors in generated DSL | Medium | Golden-Set test harness; user verification | Silent wrong answers if tests incomplete |
| Query too complex for compiler | Medium | Route to Ephemeral Slots tier | Loses zero-exposure benefit |
| Schema description leaks info | Low | Describe shapes, not examples | Very generic schemas only |
| Deno fallback vulnerabilities | Low (rare usage) | Deny-by-default sandbox | Only triggered for edge cases |

#### Execution Architecture

**V1 Specification: Local DSL with Deterministic Interpreter (Primary)**

LLM generates constrained Domain-Specific Language. Executed by local Go interpreter — no sandbox needed.

**Why DSL over full code generation**:
- Research shows 99.9% parse success vs 60-90% for general code (Anka DSL study)
- Eliminates code injection entirely (no arbitrary execution)
- No package hallucination (no imports exist)
- Easier to audit, validate, and formally verify
- Constrained decoding can guarantee valid output

**DSL Grammar** (SymQL):

```
FILTER calendar_events
  WHERE start >= @next_week_start
    AND start <= @next_week_end
JOIN contacts ON attendees CONTAINS contacts.email
FILTER contacts
  WHERE last_email_date < @thirty_days_ago
SCORE BY priority DESC, recency DESC
LIMIT 10
RETURN calendar_events
```

**Supported operations**:
- `FILTER` — select from table with WHERE conditions
- `JOIN` — combine tables on field relationships
- `SCORE` — rank results by fields (ASC/DESC)
- `AGGREGATE` — COUNT, SUM, AVG, MIN, MAX
- `WINDOW` — date ranges (@next_week, @last_30_days)
- `LIMIT` — cap result count
- `RETURN` — specify output table/fields

**Interpreter implementation**:
- ~3K lines of Go (based on GitLab Lingo framework patterns)
- Parser validates syntax before execution
- Type checker validates field references against schema
- Deterministic execution — same input always produces same output

**Data passing**: Interpreter runs in-process. No subprocess, no IPC overhead.

**Timeout**: 5-second execution limit enforced by Go context.

---

**Fallback: Deno Sandbox (Edge Cases)**

For queries that exceed DSL expressiveness but are clearly structured:

```bash
deno run --no-prompt --no-npm --no-remote [script.ts]
```

| Flag | Effect |
|------|--------|
| (no `--allow-*` flags) | All permissions denied by default |
| `--no-prompt` | Don't ask user for permissions interactively |
| `--no-npm` | Disallow npm package imports |
| `--no-remote` | Disallow fetching remote modules |

**Deny-by-default blocks:**
- `--allow-read`: Cannot read filesystem (data passed via stdin)
- `--allow-write`: Cannot write anywhere
- `--allow-net`: Cannot make network requests
- `--allow-env`: Cannot read environment variables
- `--allow-run`: Cannot execute other processes

**When to use Deno fallback**:
- Query requires conditional logic DSL can't express
- Complex algorithm needed (optimal scheduling, etc.)
- DSL compilation fails but query is clearly structured

**Expected usage**: <5% of structured queries. If higher, expand DSL grammar.

**Future V2**: Consider WebAssembly sandbox for even stronger isolation.

#### V1 Hardening Measures (DSL)

**1. Golden-Set Test Harness**

Before running compiled logic on real data, validate against synthetic test cases:
- Auto-generate small synthetic datasets matching schema
- Define expected results for each test case
- Compiled logic must pass all tests before execution on real data
- Catches silent logic errors without exposing real data

**2. Deterministic Canonicalization**

Normalize queries before sending to LLM:
- Sort clauses alphabetically
- De-duplicate filters
- Standardize date/time formats
- Benefits: reduces prompt variance, improves caching, lowers correlation fingerprinting

**3. Semantic Type Guards**

Runtime validation before execution:
- Validate types match schema (dates are dates, counts are integers)
- Enforce bounds (dates within reasonable range, counts non-negative)
- Check string lengths and patterns
- Reject execution if validation fails
- Prevents subtle injection via type confusion

**4. Constrained Decoding (Optional Enhancement)**

Use LLM structured output mode to enforce DSL grammar at generation time:
- Define DSL grammar as JSON schema or CFG
- LLM cannot produce invalid syntax
- 100% parse success guaranteed
- Supported by Claude and GPT APIs

#### Assumptions (DSL)

| Assumption | Status | Validation Path |
|------------|--------|-----------------|
| 60-70% of queries are structured | Hypothesis | Measure query distribution post-launch |
| DSL grammar covers V1 query needs | High confidence | Based on V1 scope (filter, join, score) |
| LLM can reliably generate valid DSL | High confidence | Research shows 99%+ for constrained languages |
| DSL interpreter is secure | High confidence | Deterministic, no external calls, in-process |
| 5-second timeout sufficient | Reasonable | Adjust based on real query complexity |
| Schema descriptions don't leak identity | Probable | Audit schema generation |
| Deno fallback rarely needed | Hypothesis | Target <5% of structured queries |

#### V1 Implementation Priority (DSL)

**Phase A of M3 (Privacy milestone)** — implement first because:
- Lower complexity than Ephemeral Slots
- Zero data exposure proves concept immediately
- Builds local execution infrastructure needed for PaaP
- Query patterns inform knowledge graph schema

---

### Path B: Prompt-as-Program (Fuzzy Data-Processing)

#### The Insight

DSL Compilation achieves zero exposure for structured queries. But what about fuzzy tasks — summarization, extraction, classification, sentiment analysis?

**Traditional approach**: Route to Local LLM (quality loss) or Ephemeral Slots (some exposure).

**PaaP insight**: A prompt **IS** a program for an LLM. For fuzzy tasks, the cloud LLM can compile to a prompt that the local LLM executes.

```
User: "Summarize my emails from last week"

Traditional (exposure):
→ Send email content to cloud LLM
→ Cloud returns summary

PaaP (zero exposure):
→ Send task description to cloud LLM (no data)
→ Cloud returns: optimal prompt program for summarization
→ Local LLM executes prompt on real emails
→ Get summary
```

The cloud LLM is a **prompt engineer**, not a summarizer.

#### How It Works

**Step 1: Task Analysis** (cloud, zero data)

User provides task type and data schema (not data):
```
Task: Summarize emails
Schema: {from, to, date, subject, body, labels[]}
Constraints: Focus on action items, max 3 sentences per email
```

**Step 2: Prompt Compilation** (cloud returns program)

Cloud LLM generates an optimized prompt:
```
You are summarizing an email. Follow these rules:
1. Identify the primary ask or action item
2. Note the sender's urgency level from tone
3. Capture key dates or deadlines mentioned
4. Output format: "[URGENCY] [ACTION]: [KEY DETAILS]"

Email to summarize:
{{EMAIL_CONTENT}}
```

**Step 3: Local Execution** (local LLM, real data)

Local LLM (Ollama) runs the prompt on each email. Real content never leaves device.

#### Why It's Different

| Approach | Cloud Sees | Local Sees | Quality |
|----------|------------|------------|---------|
| Traditional cloud | All content | Nothing | High |
| Local-only | Nothing | All content | Medium |
| Ephemeral Slots | Slotted content | Real content | High |
| **Prompt-as-Program** | Task schema only | All content | Medium-High |

PaaP is the **zero-exposure** option for fuzzy tasks. The tradeoff is local LLM quality (improving rapidly).

#### What PaaP Works For

**Good fit**:
- Summarization (optimal extraction prompts)
- Classification / sentiment analysis
- Entity extraction (with prompt-specified patterns)
- Content filtering / tagging
- Template-based generation

**Poor fit**:
- Complex reasoning (needs cloud LLM capability)
- Creative tasks requiring world knowledge
- Multi-step analysis with intermediate decisions

**Key insight**: Many "fuzzy" tasks are actually deterministic once you have the right prompt. PaaP separates prompt design (requires intelligence) from prompt execution (requires compute).

#### Example: Email Triage

**User query**: "Triage my inbox by urgency"

**Cloud receives** (zero data):
```
Task: Email triage
Schema: {from, to, date, subject, body, labels[], reply_chain_length}
Classification: urgent / needs-response / informational / ignore
Criteria: Urgency based on sender relationship, explicit deadlines, tone markers
```

**Cloud returns** (prompt program):
```markdown
# Email Triage Prompt

Classify this email into one of: URGENT, NEEDS-RESPONSE, INFO, IGNORE

Classification rules:
- URGENT: Contains words like "ASAP", "deadline today", "critical"; from manager/executive; mentions meetings within 24h
- NEEDS-RESPONSE: Direct question to recipient; explicit request for input; "please" + action verb
- INFO: CC'd only; newsletter pattern; no questions directed at recipient
- IGNORE: Automated notifications; marketing; reply-all chains > 5

Output exactly one word: URGENT, NEEDS-RESPONSE, INFO, or IGNORE

Email:
{{EMAIL}}
```

**Local LLM runs** on each email with real content. Returns classifications.

#### Prompt Optimization Techniques

The cloud LLM isn't just generating prompts — it's optimizing them:

1. **Task Decomposition**: Break complex task into prompt pipeline
2. **Chain-of-Thought Injection**: Add reasoning scaffolds for local LLM
3. **Few-Shot Example Generation**: Create synthetic examples that guide local LLM
4. **Output Format Enforcement**: Constrain local LLM output for reliable parsing
5. **Error Recovery Prompts**: Fallback prompts for when primary fails

**Example pipeline for complex task**:
```
User: "Extract key decisions from meeting notes"

Cloud generates prompt pipeline:
1. Segment prompt (split notes into sections)
2. Decision detection prompt (identify decision statements)
3. Attribution prompt (who made/owns decision)
4. Formatting prompt (standardize output)

Local runs pipeline sequentially, each prompt zero-exposure.
```

#### PaaP Schema v1.0

The formal JSON structure for Prompt-as-Program outputs from the cloud LLM.

**Single-step program:**
```json
{
  "version": "1.0",
  "type": "single",
  "task": "email_triage",
  "prompt": {
    "template": "Classify this email: {{INPUT}}\n\nOutput: URGENT|RESPONSE|INFO|IGNORE",
    "input_var": "INPUT",
    "output_format": "enum",
    "output_values": ["URGENT", "RESPONSE", "INFO", "IGNORE"]
  },
  "validation": {
    "output_must_match": "^(URGENT|RESPONSE|INFO|IGNORE)$",
    "retry_on_fail": true,
    "max_retries": 2
  }
}
```

**Multi-step pipeline:**
```json
{
  "version": "1.0",
  "type": "pipeline",
  "task": "extract_decisions",
  "steps": [
    {
      "id": "segment",
      "prompt": {
        "template": "Split this text into logical sections:\n{{INPUT}}\n\nOutput JSON array of sections.",
        "input_var": "INPUT",
        "output_format": "json_array"
      },
      "output_var": "SECTIONS"
    },
    {
      "id": "detect",
      "for_each": "SECTIONS",
      "prompt": {
        "template": "Does this section contain a decision? Answer YES or NO.\n\n{{SECTION}}",
        "input_var": "SECTION",
        "output_format": "enum",
        "output_values": ["YES", "NO"]
      },
      "filter": "YES",
      "output_var": "DECISION_SECTIONS"
    },
    {
      "id": "extract",
      "for_each": "DECISION_SECTIONS",
      "prompt": {
        "template": "Extract the decision from this text:\n{{SECTION}}\n\nFormat: {decision, owner, date}",
        "input_var": "SECTION",
        "output_format": "json_object"
      },
      "output_var": "DECISIONS"
    }
  ],
  "final_output": "DECISIONS"
}
```

**Schema fields:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `version` | string | Yes | Schema version (currently "1.0") |
| `type` | enum | Yes | "single" or "pipeline" |
| `task` | string | Yes | Human-readable task identifier |
| `prompt` / `steps` | object/array | Yes | Single prompt or array of steps |
| `prompt.template` | string | Yes | Prompt text with `{{VAR}}` placeholders |
| `prompt.input_var` | string | Yes | Variable name for input data |
| `prompt.output_format` | enum | Yes | "text", "enum", "json_object", "json_array" |
| `prompt.output_values` | array | If enum | Valid output values for enum format |
| `validation.output_must_match` | regex | No | Regex to validate output |
| `validation.retry_on_fail` | bool | No | Retry with stricter prompt on validation fail |
| `steps[].id` | string | Yes | Unique step identifier |
| `steps[].for_each` | string | No | Variable to iterate over |
| `steps[].filter` | string | No | Keep only items matching this value |
| `steps[].output_var` | string | Yes | Variable name for step output |
| `final_output` | string | If pipeline | Variable containing final result |

**Variable passing:**
- Variables are scoped to the pipeline execution
- `{{VAR}}` in templates replaced with variable values
- `for_each` iterates and binds each item to the specified `input_var`
- `filter` removes items that don't match the specified output value

**Validation:**
- `output_format: enum` requires output to be one of `output_values`
- `output_must_match` regex validated before accepting output
- Failed validation triggers retry with appended "Output MUST be exactly: [values]"

#### Failure Modes & Limitations (PaaP)

| Failure Mode | Severity | Mitigation | Residual Risk |
|--------------|----------|------------|---------------|
| Local LLM quality insufficient | Medium | Fall back to Ephemeral Slots | User notices quality gap |
| Prompt too complex for local | Medium | Simplify prompt; add CoT scaffolding | Multi-step prompts may fail |
| Task not decomposable | Low | Route to Ephemeral Slots | Some tasks need cloud intelligence |
| Prompt injection via data | Low | Sanitize input; structured output mode | Edge cases may leak |

#### Assumptions (PaaP)

| Assumption | Status | Validation Path |
|------------|--------|-----------------|
| Local LLMs can follow optimized prompts | Medium confidence | Llama 3, Mistral, Phi-3 capable but quality varies |
| Prompt optimization adds meaningful quality | **Hypothesis** | A/B test vs naive prompts |
| Most fuzzy tasks decompose to prompts | **Hypothesis** | Research on prompt engineering |
| Prompt generalizes across local model quirks | **Low confidence** | Cloud LLM may generate prompt that local LLM handles poorly |
| User accepts local LLM quality tradeoff | **Hypothesis** | User preference varies; may prefer Ephemeral Slots |

**PaaP is speculative for V1**. DSL Compilation is proven technology. PaaP is novel and requires validation before broad coverage claims.

**V1 gate**: PaaP is **optional**. If local model quality benchmarks show < 70% task success rate on summarization/classification, V1 ships without PaaP. Fuzzy tasks route to Ephemeral Slots instead (protected, not zero-exposure).

#### V1 Implementation Priority (PaaP)

**Phase B of M3** — implement after DSL, but **gated on local LLM quality**:
- DSL proves zero-exposure concept first
- Local LLM infrastructure needed (shared with content fallback)
- Benchmark local models before committing to PaaP coverage
- If benchmarks fail, skip PaaP for V1 — fuzzy tasks use Ephemeral Slots

---

### Combined Coverage

With both compilation paths, zero-exposure extends to many queries:

| Query Type | Path | Exposure |
|------------|------|----------|
| Filter/search/aggregate | DSL Compilation | Zero |
| Summarize/extract/classify | Prompt-as-Program | Zero (if local LLM adequate) |
| Complex reasoning | Ephemeral Slots | Protected (slots visible, entities hidden) |
| Content generation | Ephemeral Slots or Local | Protected or Zero |

**Coverage estimates (hypothesis — requires measurement):**

| Path | Estimated Coverage | Confidence | V1 Status |
|------|-------------------|------------|-----------|
| DSL Compilation | 40-50% | Medium (structured queries only) | Core |
| PaaP Compilation | 10-20% | **Low** (local LLM quality dependent) | **Optional** — ships only if benchmarks pass |
| Ephemeral Slots | 20-30% | Medium (reasoning tasks) | Core |
| Local LLM only | 10-20% | High (content tasks) | Core |

**V1 target**: 60-70% of queries via zero-exposure paths (DSL + PaaP + Local). If PaaP is cut, target drops to 50-70% (DSL + Local only); fuzzy tasks route to Ephemeral Slots instead.

**Gate**: Measure actual query distribution in dogfooding. Adjust estimates before M5. PaaP gate: 70% task success on summarization/classification benchmarks.

### Cost of Privacy Model

Privacy has costs: latency, quality, and API spend. This table makes the tradeoffs explicit.

**Estimated costs per query (V1):**

| Path | Latency | API Calls | API Cost* | Quality | Privacy |
|------|---------|-----------|-----------|---------|---------|
| **Heuristic/Template** | ~50ms | 0 | $0 | Low (rigid) | Zero exposure |
| **DSL Compilation** | 1-2s | 1 cloud | ~$0.002 | High (logic) | Zero exposure |
| **PaaP Compilation** | 2-4s | 1 cloud + local | ~$0.002 | Medium | Zero exposure |
| **Ephemeral Slots** | 1-3s | 1 cloud | ~$0.003 | High | Protected |
| **Local LLM only** | 2-5s | 0 | $0 | Medium-Low | Zero exposure |
| **Hybrid (DSL→Slots)** | 2-4s | 2 cloud | ~$0.005 | High | Zero→Protected |
| **Interactive Scaffolding** | 4-10s | 2-4 cloud | ~$0.008 | High | Zero→Protected |

*Cost estimates assume Claude Sonnet pricing (~$3/1M input, $15/1M output tokens) for typical query sizes.

**Key insights:**
- **Heuristics are free** — invest in expanding template coverage
- **DSL is cheap and fast** — maximize structured query coverage
- **Local LLM trades quality for zero cost** — acceptable for drafts
- **Ephemeral Slots is the quality ceiling** — use when reasoning genuinely needed
- **Interactive Scaffolding is expensive but surgical** — pays off for complex queries that would otherwise need massive legends

**Cost optimization strategy:**
1. Route as much as possible to DSL/Heuristics (fast, cheap, zero exposure)
2. Use PaaP for fuzzy data tasks if local quality is acceptable
3. Reserve Ephemeral Slots for genuine reasoning needs
4. Use Interactive Scaffolding only for complex multi-part queries

---

### Interactive Scaffolding (V1.5): Intelligent Orchestration for Complex Queries

#### The Insight

A query the system is "uncertain" about is not a failure case — it's an opportunity for guided, collaborative reasoning.

Instead of routing uncertain/complex queries to a single massive Ephemeral Slots prompt, the system engages the remote LLM in an **abstract planning dialogue** to build a custom execution plan that composes existing V1 tools.

**This makes "uncertain" the smartest path, not a fallback.**

#### The Problem It Solves

Complex queries mix multiple task types:

> "Looking at my recent emails with the design team, are we on track with the mockups for Project Sparrow, and can you draft a quick update for Jane?"

This combines:
- Data processing (emails)
- Reasoning (are we on track?)
- Generation (draft update)

A simple classifier struggles. The current approach: route to Ephemeral Slots with a large legend containing all context.

**The problem**: Large legends leak more structural information and require more context than necessary.

#### How It Works

**Phase 1: Abstract Planning Dialogue (Zero Content)**

The local client engages the remote LLM in a short, abstract conversation. **No private data or content is exchanged** — only data shapes, yes/no answers, and role information.

```
Client → LLM: "Complex query.
  Goal: Assess project status from communications and draft progress report for superior.
  Available Data Shapes: [email_collection], [project_plan], [contact_A], [contact_B]"

LLM → Client: "To create a plan, I need information.
  Do the emails contain explicit status keywords like 'approved', 'delayed', 'completed'?"

Client → LLM: (Runs local grep) "Yes."

LLM → Client: "Is [contact_A] or [contact_B] the superior?"

Client → LLM: (Checks local KG) "[contact_B] has role 'manager'."

LLM → Client: "I have enough to create an execution plan."
```

**Phase 2: Execution Program**

The output is not an answer but a **program** that composes V1 tools:

```json
{
  "plan_name": "status_assessment_and_report_draft",
  "steps": [
    {
      "step": 1,
      "tool": "PaaP",
      "description": "Extract status keywords from emails",
      "program": {
        "engine": "local_llm",
        "system_prompt": "Extract one of: 'approved', 'delayed', 'completed', 'at-risk'",
        "input": "{{emails_with_design_team}}"
      },
      "output_variable": "$STATUS"
    },
    {
      "step": 2,
      "tool": "Ephemeral_Slots",
      "description": "Draft the status update email",
      "prompt": {
        "legend": [
          "{{RECIPIENT}} is the user's manager",
          "{{PROJECT}} is the project in question",
          "{{CURRENT_STATUS}} is: $STATUS"
        ],
        "task": "Draft concise email to {{RECIPIENT}} about {{PROJECT}}. State status is {{CURRENT_STATUS}} and suggest next steps."
      }
    }
  ]
}
```

The local client executes this program:
1. PaaP extracts status locally (zero exposure)
2. Ephemeral Slots drafts email with minimal, targeted legend (protected)

#### Why It's Elegant

| Benefit | Explanation |
|---------|-------------|
| **Makes "uncertain" smartest** | Cloud LLM's intelligence resolves uncertainty and creates optimal plan |
| **Hardens reasoning** | Multi-step with targeted legends beats single massive prompt |
| **Meta-innovation** | Composes existing tools, doesn't add new primitives |
| **Reduces prompt engineering** | User can be vague; dialogue refines into precise program |
| **Smaller legends** | Messy context processed locally via PaaP; final legend is minimal |

#### Privacy Model

| Phase | What's Exchanged | Exposure |
|-------|------------------|----------|
| Phase 1 | Data shapes, yes/no answers, role info | **Zero content** — abstract structural info only |
| Phase 2 | Execution plan (JSON) | **Zero** — it's a program, not data |
| Execution: PaaP step | Nothing sent remotely | **Zero** — runs locally |
| Execution: Slots step | Minimal legend + slotted task | **Protected** — targeted, small legend |

**Net effect**: Less exposure than routing directly to Ephemeral Slots with full context.

#### Concerns & Mitigations

| Concern | Severity | Mitigation |
|---------|----------|------------|
| **Latency** | Medium | 2-4 dialogue exchanges = 4-8 seconds. Acceptable for complex queries. |
| **Dialogue reliability** | Medium | Constrain LLM to structured question format; validate questions are answerable |
| **Question-based leakage** | Low | "Do emails mention salary disputes?" reveals info even if no. Limit question types to structural (data shapes, counts, roles) |
| **Plan validation** | Medium | Validate generated plan against allowed tool set; reject unknown steps |
| **Dialogue hallucination** | Medium | If LLM asks unanswerable questions, abort and fall back to standard Ephemeral Slots |

#### V1.5 Scope

**Why V1.5 (not V1)**:
- Requires proven DSL + PaaP + Ephemeral Slots infrastructure
- Dialogue protocol needs design and testing
- Adds latency — validate user acceptance first
- Core V1 can ship without it; uncertain queries use standard Ephemeral Slots

**Implementation priority**:
1. V1: Simple classifier routes to DSL / PaaP / Ephemeral Slots / Local
2. V1.5: Add Interactive Scaffolding for "uncertain" route
3. Measure: Does scaffolding produce smaller legends? Better quality? Acceptable latency?

**Success criteria**:
| Metric | Target | Measurement |
|--------|--------|-------------|
| Legend size reduction | 50% smaller than naive Ephemeral Slots | Token count comparison |
| Quality equivalence | No degradation vs full-legend approach | User rating (1-5 scale) |
| Latency acceptance | < 10 seconds for planning dialogue | Time from query to plan |
| User satisfaction | 80%+ prefer scaffolded path | A/B test preference |

If any criterion fails, re-evaluate whether V1.5 is ready to ship.

---

## Novel Approach 3: P2P Query Mixing (V2+)

> **V2+ scope**: This section documents a future enhancement requiring multiple users. Not applicable for V1 dogfooding and can be skipped if focusing on V1 implementation.

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
┌─────────────────────────────────────────────────────────┐
│                       YOUR QUERY                         │
└─────────────────────────────────────────────────────────┘
                            ↓
               ┌──────────────────────┐
               │   Query Classifier   │
               │   (local, simple)    │
               └──────────────────────┘
                            ↓
    ┌───────────┬───────────┼───────────┬───────────┐
    ↓           ↓           ↓           ↓           ↓
STRUCTURED   FUZZY-DATA   HYBRID    REASONING   UNCERTAIN
(filter,     (summarize,  (filter+  (prioritize, (default
 search)     classify)    content)  plan)        private)
    ↓           ↓           ↓           ↓           ↓
    │           │      ┌────┴────┐      │           │
    │           │      │DSL Comp │      │           │
    │           │      │(filter) │      │           │
    │           │      └────┬────┘      │           │
    │           │           ↓           │           │
    │           │      Run locally      │           │
    │           │           ↓           │           │
    │           │      Filtered         │           │
    │           │      results          │           │
    │           ↓           ↓           ↓           ↓
┌─────────┐ ┌─────────┐ ┌─────────────┐ ┌─────────────────┐
│   DSL   │ │  PaaP   │ │  Ephemeral  │ │    Local LLM    │
│ Compile │ │ Compile │ │    Slots    │ │    (Ollama)     │
└─────────┘ └─────────┘ └─────────────┘ └─────────────────┘
    ↓           ↓             ↓                 ↓
Cloud LLM   Cloud LLM    Cloud LLM          Processes
returns     returns      reasons on         CONTENT
SymQL       PROMPT       LEGEND+SLOTS       locally
    ↓           ↓             ↓                 ↓
Run in      Run in       Find-replace       Direct
Go interp   Local LLM    locally            response
    ↓           ↓             ↓                 ↓
└───────────────┴─────────────┴─────────────────┘
                            ↓
                    ┌──────────────┐
                    │   RESPONSE   │
                    └──────────────┘
```

**Zero-exposure paths** (no data leaves device):
- **DSL Compilation**: Cloud returns SymQL code, executed by local Go interpreter
- **Prompt-as-Program (PaaP)**: Cloud returns optimized prompt, executed by local LLM

**Protected path** (slots only, no real values):
- **Ephemeral Slots**: Cloud sees single-use placeholders + legend. Local find-replace restores real values.

**Hybrid path**: DSL Compilation filters first (zero exposure), then filtered results go to Ephemeral Slots or Local for content operation. Maximizes zero-exposure surface.

**V1.5 upgrade — Interactive Scaffolding**: The "uncertain" path becomes an intelligent orchestrator. Instead of routing to Ephemeral Slots with large legend, the system engages the cloud LLM in an abstract planning dialogue (zero content) to build a custom execution plan that composes DSL + PaaP + Ephemeral Slots. Makes "uncertain" the smartest path, not a fallback.

**Privacy by Default — Routing Principle**:
- **Uncertain queries default to Local LLM**, not Ephemeral Slots
- Classifier must be **confident** to route to higher-exposure paths
- If classification fails or is ambiguous, fall back to local processing with quality warning
- User can override to request cloud processing if local quality is insufficient

This ensures the system fails-safe toward privacy, not toward capability.

### Query Classifier Specification (V1)

The classifier is the entry point for all queries. It determines which privacy path handles the request. **This is the most critical component** — misclassification can expose data unnecessarily or degrade quality.

#### V1 Implementation: Keyword Cascade

V1 uses a simple, deterministic keyword-based classifier. No ML, no cloud calls — purely local pattern matching.

```
CLASSIFY(query):

    1. CHECK STRUCTURED PATTERNS (→ DSL Compilation)
       Keywords: "how many", "count", "list", "filter", "search",
                 "find all", "show me", "when is", "what time"
       Patterns: Questions about quantity, filtering, scheduling
       Confidence: HIGH if keyword match + no content verbs

    2. CHECK FUZZY-DATA PATTERNS (→ PaaP Compilation)
       Keywords: "summarize", "extract", "classify", "categorize",
                 "sort by importance", "triage", "label"
       Patterns: Data processing without generation
       Confidence: HIGH if keyword match + data reference
       Gate: Only if PaaP enabled (local LLM quality passed benchmark)

    3. CHECK HYBRID PATTERNS (→ DSL then Ephemeral Slots)
       Patterns: Structured query + content operation
       Examples: "Summarize my meetings with [filter]"
                 "What's most urgent in [date range]?"
       Detection: Structured keyword + content verb
       Confidence: MEDIUM (requires both patterns)

    4. CHECK REASONING PATTERNS (→ Ephemeral Slots)
       Keywords: "should I", "prioritize", "advise", "recommend",
                 "what do you think", "help me decide", "plan"
       Patterns: Requests requiring judgment, planning, advice
       Confidence: HIGH if keyword match

    5. CHECK CONTENT PATTERNS (→ Local LLM)
       Keywords: "draft", "write", "compose", "rewrite", "edit",
                 "rephrase", "translate"
       Patterns: Pure generation with user-provided content
       Confidence: HIGH if keyword match + quoted/referenced content

    6. DEFAULT (→ Local LLM with quality warning)
       If no pattern matches with HIGH confidence
       User sees: "Processing locally for privacy. Quality may vary."
       User can override: "Use cloud processing" → Ephemeral Slots
```

#### Classification Categories

| Category | Trigger Patterns | Privacy Path | Exposure Level |
|----------|------------------|--------------|----------------|
| STRUCTURED | Quantity, filter, schedule queries | DSL Compilation | Zero |
| FUZZY_DATA | Summarize, extract, classify (no generation) | PaaP Compilation | Zero |
| HYBRID | Filter + content operation | DSL → Ephemeral Slots | Zero → Protected |
| REASONING | Judgment, planning, advice | Ephemeral Slots | Protected |
| CONTENT | Draft, write, compose | Local LLM | Zero |
| UNCERTAIN | No confident match | Local LLM (default) | Zero |

#### Confidence Thresholds

| Confidence | Definition | Action |
|------------|------------|--------|
| HIGH | Strong keyword match, unambiguous intent | Route to matched path |
| MEDIUM | Partial match or mixed signals | Route with caution; log for review |
| LOW | Weak match, ambiguous | Default to Local LLM |
| NONE | No patterns match | Default to Local LLM with warning |

**Threshold behavior**: Routes to higher-exposure paths (Ephemeral Slots, PaaP) only on HIGH confidence. MEDIUM or below defaults toward privacy.

#### Order of Operations (Why This Order?)

1. **Structured first**: Zero exposure is best; check if query can be fully handled by DSL
2. **Fuzzy-data second**: Also zero exposure; check if PaaP can handle it
3. **Hybrid third**: Maximize zero-exposure portion before falling back
4. **Reasoning fourth**: Protected exposure only when reasoning genuinely needed
5. **Content fifth**: Local LLM for generation (zero exposure, may have quality tradeoff)
6. **Default last**: Always have a safe fallback

#### Examples

| Query | Classification | Reasoning |
|-------|----------------|-----------|
| "How many meetings do I have next week?" | STRUCTURED | "How many" + time filter → DSL |
| "Summarize my unread emails" | FUZZY_DATA | "Summarize" + data reference → PaaP |
| "What's the most urgent email from this week?" | HYBRID | Filter (this week) + judgment (urgent) |
| "Should I accept this meeting invitation?" | REASONING | "Should I" → needs judgment |
| "Draft a reply declining politely" | CONTENT | "Draft" + generation → Local LLM |
| "What's the deal with Project Phoenix?" | UNCERTAIN | Ambiguous intent → Local LLM default |

#### Failure Modes

| Failure | Consequence | Mitigation |
|---------|-------------|------------|
| False positive for STRUCTURED | DSL compilation fails | Fallback cascade to Ephemeral Slots |
| False positive for FUZZY_DATA | PaaP produces poor quality | User sees quality warning; can retry with Ephemeral Slots |
| False negative (under-classification) | Query goes to Local LLM unnecessarily | User can override; quality acceptable for most cases |
| False positive for higher exposure | Data sent when it could have been local | Logging + weekly audit to tune patterns |

**Design principle**: False negatives (routing to Local LLM when unnecessary) are acceptable. False positives (routing to higher exposure when unnecessary) should be minimized.

#### V1.5: Confidence Scoring

V1.5 may add lightweight confidence scoring:
- Multiple keyword matches increase confidence
- Conflicting patterns decrease confidence
- Query length and complexity as signals
- User feedback loop to tune thresholds

For V1, the simple keyword cascade is sufficient. Complexity can be added based on dogfooding data.

---

## UX-Centric Privacy Controls

Technical privacy means nothing if users can't understand or control it. These UX patterns make privacy tangible.

### User-Facing Failure Model (V1)

Trust is built on transparently handling failure. When something goes wrong, the user should understand what happened and what options they have.

**Philosophy**: Fail visibly, fail safely, offer alternatives.

#### Failure Types and User Experience

| Failure | What User Sees | Options Offered |
|---------|----------------|-----------------|
| **DSL compilation fails** | "Couldn't convert this to a structured query." | "Try with AI reasoning instead?" (→ Ephemeral Slots) |
| **PaaP quality too low** | "Local processing produced uncertain results." | "Review and accept?" / "Retry with cloud?" (→ Ephemeral Slots) |
| **NER misses entity** | "Found these entities: [list]. Add any I missed?" | Edit entity list before sending |
| **Rehydration fails** | "Couldn't fill in all the names. Raw response:" | Show partial result with unfilled slots visible |
| **Local LLM unavailable** | "Local AI not running. Some features limited." | "Start Ollama?" / "Use cloud with privacy protection?" |
| **Cloud API error** | "Couldn't reach AI service." | "Retry?" / "Try local-only?" |
| **Classification uncertain** | "Not sure how to handle this privately." | "Process locally (may be lower quality)?" / "Use cloud with privacy protection?" |

#### Design Principles

1. **Never silently degrade**: If privacy or quality is worse than expected, say so
2. **Always offer alternatives**: Every failure has a fallback path
3. **Explain in plain language**: No technical jargon in user-facing messages
4. **Default to privacy**: When in doubt, offer the more private option first
5. **Make repair easy**: If user catches an error (e.g., missed entity), one-click fix

#### Example Failure Flows

**DSL Compilation Failure:**
```
User: "What's the vibe of my meetings this week?"

System: "I can't convert this to a structured query.
        This question needs AI reasoning.

        Option 1: Use cloud AI with privacy protection
                  (names hidden, relationships shown)
        Option 2: Process locally (slower, may be less accurate)

        [Cloud with protection] [Local only] [Cancel]"
```

**NER Miss (Critical):**
```
User: "Draft an email to JohnBoy about the Phoenix thing"

System: "I found these entities to protect:
        • JohnBoy → will become {{PERSON_kestrel}}
        • Phoenix → will become {{PROJECT_sparrow}}

        Did I miss anything? (nicknames, code names?)

        [Looks good] [Add more] [Show what will be sent]"
```

**Rehydration Failure:**
```
System: "The AI response couldn't be fully processed.
        Some placeholders weren't filled in:

        'Dear {{PERSON_kestrel}}, regarding {{PROJECT_sparrow}}...'

        [Edit manually] [Retry] [Show full response]"
```

#### Error Logging

All failures logged locally with:
- Timestamp
- Failure type
- Query (sanitized)
- Path attempted
- Fallback taken
- User choice

This supports both debugging and transparency (user can audit their failure history).

---

### V1: Progressive Consent Ladder (A)

Every new data type unlocks only after explicit micro-consent.

**How it works**:
- First time a query needs a data type, system asks in one sentence:
  - "Allow calendar titles only for conflict detection?"
- Each consent step is scoped to: **purpose + field set + time window**
- Defaults to "no" and offers a local alternative
- User builds up permissions incrementally, understanding each exposure

**Example flow**:
```
Query: "What meetings conflict with my vacation?"

System: "This needs calendar event times. Allow?"
  → [Yes, for scheduling queries] [No, use local-only]

User: Yes

System: "This also needs calendar titles to show results. Allow?"
  → [Yes, for this session] [Yes, always for scheduling] [No]
```

**Why it works**:
- Meaningful consent — user understands what they're allowing
- Scope-limited — not "allow all calendar access forever"
- Graceful degradation — local alternative always available
- Aligns with Principle #10 (User Control)

**Implementation notes**:
- Consent stored locally with scope metadata
- Revocable at any time via Privacy Receipts
- First-run is more clicks; steady-state is smooth

---

### V1: Privacy Receipts + Undo (C)

Every exposure logged in human language with 1-click revoke.

**How it works**:
- After each remote call, show a receipt:
  ```
  Sent: 3 event titles + attendee roles (no names)
  Purpose: conflict check
  Provider: Claude API
  Time: 2 minutes ago
  ```
- Receipt log accessible anytime (local storage)
- "Undo last 24h" revokes permissions granted in that window and clears caches

**Receipt anatomy**:

| Field | Example | Why |
|-------|---------|-----|
| What sent | "3 event titles, 5 attendee roles" | Concrete, not vague |
| What NOT sent | "No names, no event content" | Reassurance |
| Purpose | "Conflict detection" | Links to consent |
| Provider | "Claude API" | Know who saw it |
| Timestamp | "2 min ago" | Audit trail |

**Undo mechanics**:
- Revokes consent grants from the undo window
- Clears any cached projections/responses
- Cannot un-send what was sent — but stops future exposure
- User informed: "Permissions revoked. Data already sent cannot be recalled."

**Why it works**:
- Transparency without requiring user to read logs
- Reversibility — mistakes are recoverable
- Accountability — system proves what it did
- Aligns with Principles #11 (Reversibility) and #12 (Transparency/Privacy Split)

**Implementation notes**:
- Receipts stored locally (privacy-respecting audit trail)
- Compact UI: expandable cards, not walls of text
- Export option for users who want full audit

---

### V1: Consent Recipes (A)

Named presets for common privacy configurations. Move decisions to setup time, not query time.

**How it works**:
- A "recipe" is a named policy bundle:
  ```
  Recipe: "Calendar-only (titles)"
  ├── Allows: event titles, start/end times, attendee count
  ├── Blocks: descriptions, attendee names, locations
  └── Route: DSL preferred, Ephemeral Slots fallback

  Recipe: "Email metadata (no bodies)"
  ├── Allows: subject lines, sender/recipient domains, dates
  ├── Blocks: email bodies, attachments, full addresses
  └── Route: Ephemeral Slots only

  Recipe: "Strict local only"
  ├── Allows: nothing to cloud
  ├── Blocks: all remote calls
  └── Route: Local LLM + Heuristics only
  ```

**Flow**:
1. User selects recipe on first run or per task type
2. System applies recipe silently for matching queries
3. If query needs more than recipe allows → one-line escalation choice
4. User can switch recipes anytime

**Why it works**:
- Decisions at setup time, not every query
- Human-readable: "Allows titles + dates; blocks bodies"
- Quick switching for different contexts (work vs personal)
- Reduces consent fatigue without sacrificing control

**Relationship to other controls**:

| Layer | Granularity | When Set |
|-------|-------------|----------|
| Progressive Consent | Per-field | Per-query |
| Recipes | Per-task-type | Setup time |
| Trust Dial | Global posture | Rarely changed |

**Implementation notes**:
- Recipes are just saved policy configs with friendly names
- UI: dropdown selector + preview card showing what's allowed/blocked
- Ship with 3-5 default recipes; user can create custom
- Recipe violations trigger minimal escalation UI

---

### V1: Field-Reveal Chip UI (B)

User explicitly sees and controls which fields leave the device.

**How it works**:
- Chips represent field types:
  ```
  [event titles] [attendee roles] [time ranges] [email subjects]
  ```
- Chips are initially gray (off); user taps to enable
- Recommended fields highlighted based on query
- Live preview shows exactly "what will be sent"

**Flow**:
1. User enters query
2. System identifies required fields, shows chips with recommendations
3. User toggles chips on/off
4. Preview updates: "Will send: 3 event titles, 5 attendee roles"
5. User confirms or adjusts

**Example UI**:
```
Query: "Who should I follow up with after my meetings?"

Required fields:
  [●] event titles (recommended)
  [●] attendee roles (recommended)
  [ ] attendee names
  [ ] event descriptions
  [●] time ranges (recommended)

Preview: "Will send 3 event titles, 5 attendee roles, time ranges.
         Will NOT send: names, descriptions."

[Send] [Reduce exposure]
```

**Why it works**:
- User never guesses what's leaving device
- Explicit and granular, not abstract
- "What you see is what you send"
- Builds trust through transparency

**Relationship to other controls**:
- Taint Tracking = code-level field tagging (invisible)
- Field Chips = user-facing visualization (visible)
- Together = technical enforcement + user understanding

**Implementation notes**:
- Chips map directly to taint-tracked fields
- "Recommended" = minimum for query success
- Can be collapsed to "trust recommendations" for power users
- Integrates with Privacy Receipts (chips shown match receipt)

---

### V1: Time-Boxed Access (D)

All consent grants have TTL. Prevents permission creep by default.

**How it works**:
- Every scope grant is time-bound:
  ```
  "Allow email metadata for: [15 min] [1 hour] [1 day]"
  ```
- Countdown badge shows remaining time
- On expiry, access auto-revokes
- Next query requiring that scope prompts again

**Default durations**:

| Context | Default TTL | Rationale |
|---------|-------------|-----------|
| One-off query | 15 minutes | Task-scoped |
| Work session | 1 hour | Reasonable work block |
| Daily workflow | 1 day | Persistent but bounded |
| Recipe-based | Recipe default | Task-type appropriate |

**Flow**:
1. User approves scope with duration
2. Badge shows: "Email metadata: 47 min remaining"
3. On expiry: silent revocation, no interruption
4. Next relevant query: "Email metadata access expired. Renew?"

**Why it works**:
- Prevents silent, long-term access accumulation
- Keeps user awareness fresh
- Forces periodic re-evaluation of what's needed
- Android/iOS pattern for location — proven UX

**Implementation notes**:
- Simple TTL on scope grants in local state
- Timer runs client-side
- Expiry logged in Privacy Receipts
- "Always" option available but discouraged (requires extra click)

---

### V1.5: Privacy Budget Meter (C)

Ambient exposure indicator. Shows cost before action.

**How it works**:
- Meter shows exposure level: Low → Medium → High
- Each query shows estimated cost:
  ```
  "This query: Medium exposure
   Includes: titles + attendee roles
   Excludes: names, bodies

   [Proceed] [Reduce exposure]"
  ```
- "Reduce exposure" shows lower-cost alternatives

**Exposure scoring** (simple, not formal DP):

| Factor | Score |
|--------|-------|
| Public fields only | +0 |
| Derived fields | +1 each |
| Sensitive fields | +3 each |
| Remote call | +2 |
| Local only | +0 |

**Why V1.5 (not V1)**:
- "Budget" metaphor needs care — implies scarcity
- V1 has Receipts (after) + Chips (before) — sufficient visibility
- Budget adds continuous feedback but also complexity
- Refine framing: "exposure level" not "budget" to avoid confusion

**Clarification**: Budget is informational, not limiting. High exposure doesn't block — just informs. User always has final choice.

**When to promote**:
- If users want more guidance on exposure tradeoffs
- If "reduce exposure" variants prove useful
- After V1 UX patterns are validated

---

### V1.5: Trust Dial + Auto-Routing (B)

Single slider that sets privacy/capability default. For users who've learned the system and want less friction.

**How it works**:
- Single slider with three positions:
  ```
  [Private] ←——————→ [Balanced] ←——————→ [Powerful]
  ```
- Each position maps to routing behavior:

| Position | Default Routing | What It Means |
|----------|-----------------|---------------|
| Private | Local-only, Heuristics | No data leaves device; accept quality limits |
| Balanced | DSL + Ephemeral Slots | Context via legend, not identities; good quality |
| Powerful | Full context when needed | Maximum capability; trust the provider |

- System explains consequences: "Balanced allows titles, not content."
- User can still override per-query

**Why V1.5 (not V1)**:
- Progressive Consent (A) is better for learning the system
- Trust Dial is the "I get it, just set my default" graduation
- Introducing both at once creates confusion
- After users understand via A, offer B as simplification

**Interaction with Consent Ladder**:
- Trust Dial sets the *default* for new data types
- Consent Ladder still gates *first access* to each type
- Dial = "what's my posture?" / Ladder = "do I allow this specific thing?"

---

### V2: Privacy Sandbox Mode (D)

"Try it with fake data" onboarding flow.

**How it works**:
- On first run, user can choose: "Start with demo data"
- System populates with realistic synthetic data (fake calendar, fake contacts)
- User explores full functionality with zero real exposure
- When ready: "Switch to real data" triggers consent ladder

**Why it matters**:
- Privacy tradeoffs become concrete, not abstract
- User sees exactly what gets sent before risking real data
- Reduces "I didn't realize it would do that" regrets
- Builds trust before asking for trust

**Why V2 (not V1)**:
- V1 is dogfooding — lead dev already understands the system
- Sandbox requires generating convincing synthetic data
- More valuable when onboarding users who don't know SymPAL

**Implementation notes**:
- Synthetic data generator (calendar, todos, contacts)
- Clear visual indicator: "DEMO MODE" banner
- One-click transition to real data
- Option to keep sandbox for testing new features

---

## Security-Centric Privacy Controls

Defense in depth. Multiple layers that each independently prevent data leakage.

### V1: Policy-as-Code Gate (A)

Every outbound payload passes through a local policy engine before leaving.

**How it works**:
- Rules engine evaluates every outbound request:
  - Which fields can leave?
  - Which endpoints are allowed?
  - Which user consent scopes are required?
- Rules are versioned, logged, and auditable
- Denial = request blocked + logged + user notified

**Example rules**:
```yaml
rules:
  - name: "no-email-bodies"
    match: { data_type: "email.body" }
    action: deny
    message: "Email bodies never sent to cloud"

  - name: "calendar-titles-only"
    match: { data_type: "calendar.*" }
    allow: ["title", "start", "end", "attendee_count"]
    deny: ["attendees", "description", "location"]

  - name: "require-consent"
    match: { data_type: "contacts.*" }
    require_scope: "contacts_read"
    action: deny_unless_consented
```

**Why it works**:
- Hard boundary — policy engine is the gatekeeper
- Auditable — every decision logged with rule that triggered it
- Versioned — rule changes are tracked, can rollback
- Declarative — rules are readable, not buried in code

**Implementation notes**:
- Single enforcement point (all outbound goes through gate)
- Rules stored as YAML/JSON, version controlled
- Hot-reload rules without restart
- Integrates with Privacy Receipts (log shows which rule allowed/denied)

---

### V1: Strict Egress Firewall (C)

All outbound requests go through a single process that enforces shape, size, and destination.

**How it works**:
- Single egress point for all external requests
- Enforces:
  - **Allowed endpoints**: whitelist of permitted API URLs
  - **Request shape**: only expected field structures pass
  - **Token budget**: max tokens per request, per query, per session
  - **Rate limits**: max requests per minute/hour

**Enforcement layers**:

| Layer | What It Checks | Block Condition |
|-------|----------------|-----------------|
| Endpoint whitelist | URL destination | Unknown endpoint |
| Schema validation | Request structure | Unexpected fields |
| Size limits | Token count | Exceeds budget |
| Rate limiting | Request frequency | Too many requests |

**Example configuration**:
```yaml
egress:
  allowed_endpoints:
    - "https://api.anthropic.com/v1/*"
    - "https://api.openai.com/v1/*"

  limits:
    max_tokens_per_request: 4000
    max_tokens_per_query: 8000
    max_requests_per_minute: 10

  schema:
    required_fields: ["model", "messages"]
    forbidden_fields: ["raw_data", "full_context"]
```

**Why it works**:
- Defense in depth — even if policy logic has bugs, firewall catches anomalies
- Single chokepoint — all egress auditable in one place
- Prevents exfiltration — unexpected payloads blocked
- Catches bugs — malformed requests fail fast

**Implementation notes**:
- Implemented as middleware/interceptor on HTTP client
- Logs all blocked requests for debugging
- Alerts on repeated blocks (possible bug or attack)
- Works independently of Policy Gate (belt AND suspenders)

---

### V1: Deterministic Redaction Layer (D) — Narrow Scope

Pattern-based PII removal as fail-safe before Ephemeral Slots.

**How it works**:
- Simple regex/pattern matching on all outbound text
- Catches obvious PII that NER might miss
- Runs BEFORE Ephemeral Slots layer
- Deterministic, testable, no ML uncertainty

**V1 scope (narrow)**:
```
Patterns:
├── Email addresses: *@*.* → [EMAIL]
├── Phone numbers: +1-xxx-xxx-xxxx variants → [PHONE]
├── SSN: xxx-xx-xxxx → [SSN]
├── Credit cards: 16-digit patterns → [CC]
└── URLs with usernames: user:pass@host → [CREDENTIAL]
```

**What it's NOT**:
- Not a replacement for Ephemeral Slots (doesn't preserve meaning)
- Not comprehensive NER (doesn't catch "John Smith")
- Not ML-based (deterministic patterns only)

**Why narrow scope**:
- Simple patterns = high confidence, low false positives
- Names/entities handled by Ephemeral Slots layer
- This is the safety net, not the primary privacy mechanism

**Example flow**:
```
Original: "Email john@acme.com about the meeting"
                    ↓
Redaction: "Email [EMAIL] about the meeting"
                    ↓
Ephemeral Slots: "Email [EMAIL] about {{EVENT_sparrow}}"
                    ↓
Egress: Passes policy gate + firewall
```

**Why it works**:
- Defense in depth — catches what NER misses
- Deterministic — same input always produces same output
- Testable — can verify all patterns with test suite
- Low effort — ~50 lines of regex, not a major system

**Implementation notes**:
- Runs before Ephemeral Slots in pipeline
- Logging shows what was redacted (locally)
- False positive rate should be near-zero with narrow patterns
- Expand patterns in V1.5 if gaps discovered

---

### V1: Taint-Tracked Serialization (A)

Mark fields as sensitive at the data model level. Serialization refuses to include them without explicit policy grant.

**How it works**:
- Every field in data models has a taint level:
  ```go
  type CalendarEvent struct {
      ID          string    `taint:"public"`
      Title       string    `taint:"sensitive"`
      Description string    `taint:"sensitive"`
      Start       time.Time `taint:"derived"`  // can derive "morning" from exact time
      Attendees   []string  `taint:"sensitive"`
      AttendeeCount int     `taint:"public"`   // derived, safe
  }
  ```
- Custom serializer checks taint tags before including fields
- Sensitive fields only serialize if policy explicitly allows
- No policy grant = field omitted (not errored, just absent)

**Taint levels**:

| Level | Meaning | Default Behavior |
|-------|---------|------------------|
| `public` | Safe to send | Always included |
| `derived` | Computed from sensitive | Include if derivation policy allows |
| `sensitive` | Contains PII or private content | Omit unless explicit policy |
| `forbidden` | Never send under any circumstance | Always omitted, logged if attempted |

**Example flow**:
```
Serializing CalendarEvent for LLM request:
├── ID: public → included
├── Title: sensitive → check policy → "calendar_titles" scope granted → included
├── Description: sensitive → check policy → no grant → OMITTED
├── Start: derived → check policy → "time_windows" allows → included as "morning"
├── Attendees: sensitive → no grant → OMITTED
└── AttendeeCount: public → included

Result: {id, title, start: "morning", attendee_count}
```

**Why it works**:
- Privacy by construction — can't accidentally include sensitive fields
- Deterministic — same policy always produces same serialization
- Testable — unit tests verify taint behavior
- Cheap — struct tags + custom serializer, ~200 lines

**Relationship to other controls**:
- Taint Tracking = serialize-time prevention (can't even try)
- Policy Gate = runtime inspection (catches what slips through)
- Together = defense in depth

**Implementation notes**:
- Go struct tags for taint levels
- Custom JSON marshaler respects tags
- Policy context passed to serializer
- Logging shows which fields omitted and why

---

### V1: Schema-Hash Gate (B)

Only payloads matching pre-registered schema hashes can leave. Blocks unreviewed data shapes.

**How it works**:
- Every allowed request schema is hashed and registered in code:
  ```go
  var allowedSchemas = map[string]string{
      "calendar_query":  "sha256:a1b2c3...",  // {model, messages, max_tokens}
      "projection_req":  "sha256:d4e5f6...",  // {model, messages, context}
  }
  ```
- Before sending, compute hash of actual payload schema (field names + types)
- If hash not in registry → block + log + alert
- New fields require code change to update hash → forces review

**Schema hashing**:
```
Payload: {model: "claude-3", messages: [...], max_tokens: 1000}

Schema extraction:
├── model: string
├── messages: array
└── max_tokens: number

Canonical form: "max_tokens:number,messages:array,model:string"
Hash: sha256("max_tokens:number,messages:array,model:string") → "a1b2c3..."

Check: "a1b2c3..." in allowedSchemas? → Yes → proceed
```

**What it catches**:
- Accidental new field added to payload
- Typo creates unexpected field name
- Refactoring changes payload structure without review
- Injection attack adds fields

**Why it works**:
- Type system for egress — only reviewed shapes can exist
- Forces code review for data shape changes
- Deterministic — hash is reproducible
- Zero runtime overhead once computed

**Relationship to other controls**:
- Schema-Hash = compile-time shape approval (what CAN exist)
- Egress Firewall = runtime shape validation (what DOES exist)
- Together = defense in depth

**Implementation notes**:
- Hash computed at build time or first run
- Schema extraction ignores values, only structure
- Nested objects recursively hashed
- Array contents typed but not enumerated
- CI can verify hashes match registered values

---

### V2: Tamper-Evident Audit Chain (D)

Hash-chain every outbound payload for forensic accountability.

**How it works**:
- Every outbound payload is added to a local hash chain:
  ```
  H(0) = hash("genesis")
  H(1) = hash(H(0) + payload_1)
  H(2) = hash(H(1) + payload_2)
  ...
  H(n) = hash(H(n-1) + payload_n)
  ```
- Chain stored locally with timestamps
- Any unauthorized payload insertion breaks the chain
- Can prove: "these are ALL the payloads that left, in order"

**What it provides**:
- **Detection**: If something silently exfiltrated data, chain is broken
- **Forensics**: After incident, prove exactly what was sent
- **Trust**: Third party can verify chain integrity
- **Compliance**: Auditable record of all egress

**Chain entry structure**:
```json
{
  "index": 47,
  "timestamp": "2026-01-19T14:32:01Z",
  "payload_hash": "sha256:...",
  "payload_summary": "calendar query, 3 events, ephemeral slots tier",
  "previous_hash": "sha256:...",
  "chain_hash": "sha256:..."
}
```

**Why V2 (not V1)**:
- Detection, not prevention — A + B already prevent exfil
- More valuable for multi-user scenarios (prove integrity to others)
- Adds storage overhead (chain grows with usage)
- Solo dogfooding doesn't need forensic proof
- Valuable when: regulatory compliance, shared instances, incident response

**When to promote to V1.5**:
- If handling data for others (not just yourself)
- If regulatory requirements demand audit trail
- If you want to prove to yourself what was sent (paranoia mode)

---

### V2: Minimal-Exposure Proofs (B)

Justification-driven serialization. Every field must have a "need proof" to leave.

**How it works**:
- Before sending any field, system must produce a short proof:
  ```
  Field: event.title
  Proof: "User query 'summarize my meetings' requires titles for summarization"
  Allowed: Yes
  ```
- If no proof can be generated, field is stripped
- Proofs are logged for audit

**Proof structure**:
```
{
  "field": "event.title",
  "query": "summarize my meetings",
  "justification": "Summarization requires content to summarize",
  "consent_scope": "calendar_titles",
  "proof_type": "query_requirement"
}
```

**Proof types**:
| Type | Meaning | Example |
|------|---------|---------|
| query_requirement | Query explicitly needs this field | "Show titles" needs titles |
| implicit_requirement | Query implicitly needs this | "Summarize" needs content |
| user_override | User explicitly allowed | "Send full context" |
| no_proof | Cannot justify | Field stripped |

**Why V2 (not V1)**:
- Adds friction to every outbound field
- Requires building proof generation logic
- Valuable for formal compliance ("prove you needed X")
- Overkill for dogfooding where you trust yourself
- V1 Policy Gate + Redaction + Firewall is sufficient defense

**When to promote to V1.5**:
- If handling very sensitive data (health, financial)
- If regulatory compliance requires audit trail of justification
- If Ephemeral Slots proves insufficient and need tighter control

---

## Implementation Sequence

### V1 (M3: Privacy Milestone)

**Phase A: Query Compilation — DSL Path**
- Query classifier (structured vs fuzzy vs hybrid vs reasoning)
- Schema description generator (describe data shapes to LLM)
- SymQL interpreter (Go, deterministic execution)
- Deno sandbox fallback (for complex structured queries)
- Hybrid decomposition (detect filter+content queries, split execution)
- Test with todo + calendar data

**Phase B: Query Compilation — PaaP Path**
- Prompt compilation endpoint (task schema → optimized prompt)
- Local LLM infrastructure (Ollama integration)
- Prompt caching (reuse compiled prompts for same task types)
- Test with email summarization, classification

**Phase C: Ephemeral Slots**
- Entity extraction from queries (NER)
- Random slot generator (single-use placeholders)
- Legend constructor (minimal context per entity)
- Find-replace fulfillment (slot responses → real entities)
- Test with calendar events, todo items

**Phase D: Local LLM Content Path**
- Pure local routing for content-only queries
- Quality comparison logging (to validate tradeoffs)

### PRD Integration

| PRD Milestone | Privacy Phase | What's Tested |
|---------------|---------------|---------------|
| M1-M2 | — | Low sensitivity data only |
| M3 | A: DSL Compilation | Structured queries with zero exposure |
| M3 | B: PaaP Compilation | Fuzzy data tasks with zero exposure |
| M3 | C: Ephemeral Slots | Reasoning queries with single-use slots |
| M3 | C: Local LLM | Content tasks with local processing |
| M5-M6 | Full | Email + contacts (high sensitivity) |

**Gate**: M5 (Email) only proceeds after M3 privacy architecture validated.

### V2+

> **Note for V1 implementation**: The sections below document future enhancements and are provided for architectural context only. They are **not in scope for V1** and can be skipped if focusing on current implementation. Consider these a roadmap for post-V1 evolution.

#### The Foundry: Reusable Personal API

Instead of generating one-off scripts, the LLM incrementally builds a stable, versioned, local API for the user's data.

**How it works**:
- When user asks "Show me meetings with people I haven't emailed in 30 days," the LLM generates a clean, reusable function like `getNeglectedMeetings(dateRange, inactivityPeriod)`
- Function is validated, sandboxed, and saved to a local library
- Next similar query calls the existing trusted function — no LLM round trip

**Benefits**:
- Efficiency: Avoid redundant LLM calls for similar queries
- Compounding: User's SymPAL instance becomes more capable over time
- Reduced attack surface: Vetted functions don't need sandbox scrutiny

**Open questions**:
- Query matching: How to recognize semantically similar queries? (May require LLM)
- Schema drift: How to handle versioning when data schema changes?
- Trust model: What makes a function "vetted"? User approval? N successful runs?

**Why V2+**: Adds significant complexity (function registry, semantic matching, versioning). V1 should prove compiler works before optimizing for reuse.

---

#### Semantic Projection: Type-Based Shadows (V2+ Enhancement to Ephemeral Slots)

The original privacy layer concept, now an optional enhancement for users who want richer type vocabulary.

**How it differs from Ephemeral Slots**:

| Dimension | Ephemeral Slots (V1) | Semantic Projection (V2+) |
|-----------|---------------------|---------------------------|
| Placeholders | Random single-use (`{{PERSON_kestrel}}`) | Typed consistent (`[PERSON:manager,senior]`) |
| Entity-level correlation | Defeated | Possible via type patterns |
| Context richness | Legend-based (task-specific) | Type vocabulary (entity-specific) |
| Infrastructure | NER + random gen | Knowledge graph + taxonomy |

**When Semantic Projection adds value**:
- User wants richer type vocabulary than legends provide
- Queries require consistent entity typing across session
- User accepts some correlation risk for better reasoning
- Knowledge graph infrastructure is built (V2+)

**How it works** (if enabled):
1. Maintain typed knowledge graph of entities
2. Map entities to type placeholders: "John Smith" → `[PERSON:colleague,senior,trusted]`
3. Type vocabulary consistent within session
4. Provider can correlate type patterns over time

**Why V2+ (not V1)**:
- **Ephemeral Slots defeats entity correlation** — single-use random placeholders
- Semantic Projection's type patterns are correlatable, which was its fatal flaw
- For users who want Projection-style typing, offer as opt-in enhancement
- Only makes sense once knowledge graph infrastructure exists

**Implementation path**:
- V1: Ephemeral Slots (broad coverage, simpler, defeats entity correlation)
- V2+: Semantic Projection as opt-in mode for users who prefer richer types and accept tradeoff

---

#### Fuzzy Projections: Probabilistic Mapping (Enhancement to Semantic Projection)

If using Semantic Projection (V2+), add differential privacy to prevent perfect profile reconstruction.

**Note**: This enhances Semantic Projection, not Ephemeral Slots. Ephemeral Slots already defeats profiling via single-use randomness.

**How it works**:
- Projection function becomes probabilistic
- "John Smith" doesn't always map to `[PERSON:colleague,senior]`:
  - 80% → `[PERSON:colleague,senior]`
  - 15% → `[PERSON:colleague]`
  - 5% → `[PERSON:professional]`
- LLM provider sees a fuzzy, shifting shadow graph

**Benefits**:
- Stronger anonymity for Semantic Projection users
- Resilience to unique entities: Rare types get coarsened more aggressively

**Open questions**:
- Rehydration: If same entity maps to different types within session, how to map responses back?
- Quality impact: Inconsistent context may degrade LLM reasoning
- Tuning: What probability distribution balances privacy vs quality?

**Why V2+**: Only relevant if using Semantic Projection. Ephemeral Slots makes this unnecessary for most users.

---

#### Crowdsourced Semantics: P2P Knowledge Graph Building

Addresses cold-start problem by letting a peer network help classify entities without seeing context.

**How it works**:
- New unclassified entity (e.g., "Neurodyne Systems") is sent to random peers
- Peers see only the isolated entity name, vote on type:
  - Peer 1: "ORGANIZATION"
  - Peer 2: "ORGANIZATION, tech"
  - Peer 3: "BUSINESS"
- System aggregates votes, classifies as ORGANIZATION
- No peer sees the context in which entity was encountered

**Benefits**:
- Reduces user burden: Automates tedious classification
- Collective intelligence: Network improves over time
- Community symbiosis: Extends partnership principle to user community

**Open questions**:
- Entity name leakage: "Neurodyne Systems" might itself be sensitive (stealth startup, acquisition target)
- Malicious peers: Sybil attacks on classification? Deliberate misclassification?
- Public vs private: Works for public entities; fails for internal project codenames

**Why V2+**: Requires multiple users — not applicable for V1 dogfooding. Should be scoped to obviously-public entities only; private entities still need user classification.

---

#### Amnesic Reasoning: Stateless Micro-Query Orchestration

**Core insight**: Even with Ephemeral Slots, the LLM sees relationships between entities within a single query (via the legend). Amnesic Reasoning prevents this by **never putting related entities in the same context**.

**Note**: Ephemeral Slots defeats cross-query correlation. Amnesic Reasoning defeats intra-query correlation. Most queries don't need this — Ephemeral Slots is sufficient. But for highly sensitive multi-entity queries, Amnesic adds another layer.

**How it works**:

The local SymPAL client acts as an orchestrator. A high-level query is decomposed into atomic, disconnected micro-queries, each sent in a separate stateless API session.

Example: "Should I be worried about the Project Titan deadline?"

```
Micro-Query #1 (Temporal Analysis) — Fresh session
├── Prompt: "Given target date and event dates, what is density?"
├── Input: Target: 2026-01-30. Events: ["2026-01-26", "2026-01-27", ...]
├── Response: "High density (5 events in 5 days)"
└── Session destroyed. LLM has no idea these were meetings.

Micro-Query #2 (Dependency Analysis) — Fresh session
├── Prompt: "A project has 5 tasks. T1 blocked by T2. T3 blocked by T4..."
├── Response: "Medium risk. Two independent dependency chains."
└── Session destroyed. LLM has no idea what the tasks are.

Micro-Query #3 (Risk Classification) — Fresh session
├── Prompt: "Classify risk in snippets: 'Waiting on legal', 'Need approval'..."
├── Response: "Snippet 1: High uncertainty. Snippet 2: Dependency risk."
└── Session destroyed. LLM doesn't know snippets relate to same project.

Local Synthesis (Ollama)
├── Aggregates: High temporal density + Medium path risk + Identified blockers
└── Output: "Yes, you should be worried. Dense schedule, multiple risks."
```

**Benefits**:
- Cross-entity linking defeated: Entities never in same context
- No new tech required: Standard stateless API calls
- Provider-agnostic: Works with any LLM with stateless API
- Auditable: Each micro-query is inspectable

**When it works best**:
- Temporal/quantitative analysis ("Am I overbooked?")
- Dependency/graph analysis ("What's blocking progress?")
- Risk classification ("What are the risks in these items?")
- Pattern matching without relationship context

**When it fails**:
- Holistic reasoning ("Should I prioritize John over Sarah?")
- Creative synthesis requiring full context
- Queries where relationships ARE the point

**Implementation approach**:
- Decomposition via local LLM (Ollama) — keeps full context local
- Pre-defined micro-query templates for common analysis types
- Synthesis via local LLM — aggregates results with full context
- Accept 3-5x latency for privacy gain

**Open questions**:
- Decomposition quality: How well can local LLM fragment queries?
- Template coverage: Which query types have good micro-query templates?
- Cost/latency tradeoff: 3-5 API calls per query — acceptable for high-sensitivity?

**Verdict**: High novelty, genuinely prevents intra-query entity linking. Valuable for analytical queries on highly sensitive multi-entity data where even legend exposure is too much. Not needed for most queries — Ephemeral Slots is sufficient.

**Why V2**: Adds orchestration complexity. V1 Ephemeral Slots handles cross-query correlation by design. Amnesic is for users who also want intra-query isolation and accept latency/cost tradeoff. Implement as opt-in "maximum privacy" mode.

---

#### Two-Tier Reasoning: Structure/Content Separation (V1.5)

**Core insight**: An LLM doesn't need to know the content of your project to give you a good structure for an email about it. And it doesn't need the overall structure to polish a single sentence. Exploit this separation.

**How it works**:

Split every generative task into two distinct, un-linkable phases:

```
Phase 1: Abstract Structural Pass — Fresh session
├── Prompt: "Generate a template for a professional email to a superior about a project delay. Use placeholders."
├── Response: Template with {{GREETING}}, {{PROJECT_NAME}}, {{REASON}}, {{NEXT_STEPS}}
└── Session destroyed. LLM knows only "project delay email structure" — generic, low-value.

Phase 2: Concrete Fulfillment Pass — Fresh session(s)
├── Prompt 2a: "Generate professional greeting for 'Jane' (manager)" → "Dear Jane,"
├── Prompt 2b: "State 'API integration issues' professionally" → "unexpected complexities during API integration"
└── Sessions destroyed. LLM sees entities but not relationships.

Local Assembly
├── Client stitches template + filled slots
└── Only local client ever sees both structure AND content
```

**Privacy outcome**: The relationship between manager and project is never exposed in a single context. Structure pass sees generic request type; content pass sees isolated entities; neither sees both.

**Relationship to Ephemeral Slots**: Two-Tier is complementary. Ephemeral Slots defeats cross-query correlation. Two-Tier defeats intra-query relationship exposure for generation tasks. Can be combined: use Ephemeral Slots for entity masking within each tier.

**Why it's better than single-call approaches for sensitive generation**:
- Single Ephemeral Slots call: Legend reveals "manager + project delay" relationship (even with random slots)
- Two-Tier: Relationship pattern never linked to specific entities — structure is generic, content is isolated

**When it works best**:
- Templated generation (emails, reports, plans)
- Tasks with clear structure/content separation
- When relationship patterns are sensitive

**When it fails**:
- Creative tasks where structure IS the content
- Very short outputs (overhead not worth it)
- Holistic reasoning requiring full context

**Refinement**: Phase 2 can be hybrid — local/rule-based for trivial slots (greetings, names), cloud LLM only for complex polishing. Reduces API calls and correlation risk.

**Open questions**:
- Cross-session correlation: Same API key + timing could link phases. Token rotation helps but doesn't eliminate.
- Template library coverage: How many structural templates needed for common tasks?

**Verdict**: Strong candidate for V1.5 for users who want stronger privacy on generation tasks than Ephemeral Slots alone provides. Fills gap between Ephemeral Slots (legend reveals relationships) and Amnesic (fragmented quality).

**Why V1.5**: Simpler than Amnesic, works for generation (unlike Prompt Camouflage), genuine privacy gain for sensitive relationships. Orchestration is tractable — two-phase workflow with template library. Most users won't need this — Ephemeral Slots is sufficient.

---

#### Latent Space Scaffolding: Geometry-Based Privacy (V2)

**Core insight**: Stop sending words to the LLM entirely. Use embedding geometry as the abstraction layer — LLM reasons on structure, never sees content.

**How it works**:

```
1. Local Embedding
├── Embed all documents locally (sentence-transformers)
├── Embed query
└── Text never leaves device

2. Local Structural Analysis
├── Find documents most similar to query
├── Cluster by thematic similarity
├── Identify outliers (potential risks/anomalies)
├── Assign meaningless IDs: vec_A1, cluster_X

3. Thematic Labeling (Local)
├── Metadata-first: folder paths, email headers, calendar fields
├── TF-IDF keywords: top terms per cluster
└── Local classifier: zero-shot labeling (risk, decision, question)

4. Scaffold Prompt (to Remote LLM)
├── "cluster_A (15 vectors) labeled 'internal_comms' is relevant"
├── "vec_A9, vec_B4 are outliers with 'risk' label"
├── LLM sees ONLY structure — no content, no identities
└── LLM returns structural plan: {"action": "describe_risk", "source": "vec_A9"}

5. Two-Tier Fulfillment (Composition)
├── Retrieve local content for vec_A9
├── Isolated polishing call: "Rewrite as professional risk summary: [content]"
├── LLM has no idea this relates to broader context
└── Context-free micro-task

6. Local Assembly
├── Stitch polished pieces per plan
└── Final output assembled locally
```

**Why it works**:
- **Planning** comes from geometric analysis + scaffold reasoning
- **Content generation** comes from Two-Tier-style isolated polishing
- **Assembly** is local — only client sees both structure AND content

**Thematic labeling solutions** (all local, V1-feasible):

| Method | How It Works | Best For |
|--------|--------------|----------|
| Metadata-First | Email headers, folder paths, calendar fields | Structured data |
| TF-IDF Keywords | Top statistically significant terms per cluster | Document clusters |
| Local Classifier | Zero-shot classification (~400MB models) | Abstract labels (risk, decision) |

**What it provides over Two-Tier alone**:
- Two-Tier: Generic template request → fill
- Latent Space: Structure **informed by your actual data relationships** → fill

The scaffold tells LLM about geometric relationships in YOUR data, not just a generic template.

**When it works best**:
- Document-heavy workflows (email triage, file summarization)
- Multi-document synthesis (reports from multiple sources)
- Finding patterns across large document sets

**Tradeoffs**:
- More infrastructure (embedding, geometric analysis, scaffold construction)
- More API calls (scaffold + N fulfillment)
- Higher latency and cost

**Why V2 (not V1)**:
- V1 data is simple (todos, calendar) — overkill
- When email capability added (M5+), this becomes valuable
- Composition with Two-Tier is elegant but adds complexity

**Implementation path**:
- V1: Simpler approaches for structured data
- V2: Add Latent Space when document workflows arrive
- Composition with Two-Tier for fulfillment is key

---

#### Other V2+ Ideas

**Compiler Enhancements:**
- **Progressive Disclosure Schema**: Start with minimal schema, request more only if needed ("Do you have attendees field?"). Reduces exposure through schema descriptions.
- **Self-competition Compilation**: Two models (or temperature variants) compile same query; diff outputs. If disagreement, require clarification or fall back. Lowers error rate for borderline queries.
- **Proof-carrying Code-lite**: LLM emits proof sketch (invariants, edge cases, complexity bounds) that interpreter validates. Cheap trust checks without full formal verification.

**Projection Enhancements:**
- **Semantic Decoys**: Insert fake entities matching same type distributions into context, discard during rehydration. Defeats profiling while keeping model behavior stable.
- **Constraint-aware Projection**: Map to relational constraints ("A is senior to B", "X and Y have timeline conflict") rather than just types. Preserves reasoning logic without identity.
- **Dual Projection Lanes**: Create structural shadow (roles, relations) AND statistical shadow (frequencies, priorities). Require LLM response consistent across both.

**Cross-cutting:**
- **Privacy Budget Ledger**: Each query consumes privacy budget. When budget low, force more local processing. Principled accounting.
- **Query Plasticity**: Slightly alter query phrasing/structure while preserving meaning. Reduces provider pattern correlation.
- **Adversarial Replay**: Store shadow queries, periodically test against internal re-ID adversary model. Adapt projection granularity based on success rate.

**Architectural:**

- **Model Sharding (Multi-Provider Distribution)**: Split a single query across multiple LLM providers, each missing key context. Example: Provider A gets goals + constraints, Provider B gets schema + timeline, Provider C gets decision criteria. Client merges outputs locally. No single provider sees full picture. Works even without Projection — it's an architectural privacy layer. Concerns: query decomposition, output merging, cost (multiple providers), capability differences.

- **Encrypted Semantic Indices**: Build local search indices (embeddings) with privacy-preserving transforms: quantization, random rotation, bucketing, salting. Remote LLM receives only index hits ("doc 12, 45, 77") plus brief non-identifying metadata. Search/selection is local; remote reasons over de-identified results. More rigid than Compiler but covers search/retrieval tasks. Research direction — established techniques in privacy-preserving IR but complex to implement well.

- **Prompt Camouflage (Noise Injection)**: Inject plausible but fictitious "decoy" entities and relationships into prompts alongside real (projected) data. LLM reasons over mix of real and fake; profiler learns false graph structure. Inspired by code obfuscation and military camouflage — "offensive privacy" that feeds profilers garbage data. **Scope limitation**: Works for classification, ranking, and analysis tasks where decoys don't affect real output. Does NOT work for generation tasks (emails, messages) — decoys leak into output and corrupt it. Extraction of real results from camouflaged response is non-trivial for generation. Use Projection for generation tasks; consider Camouflage for analytical tasks where relationship patterns are sensitive.

- **Constraint-Solver Delegation**: Alternative to DSL for complex logical queries. LLM writes formal constraints (SAT/SMT/logic); local solver (e.g., Z3) executes on private data. Declarative style: `find meetings M where M.date ∈ next_week ∧ ∃p ∈ M.attendees : p.last_email < 30_days`. Better than DSL for: complex logical conditions with quantifiers, optimization problems ("find best time slot"), queries with interdependencies. DSL remains better for simple filter/join/score (lower complexity). Consider if DSL proves limiting for complex queries. Z3 integration is tractable.

**Infrastructure:**
- P2P query mixing (when users exist)
- Federated learning across SymPAL instances
- User-configurable query padding (decoy queries)

---

## Correlation Attack Mitigations

### The Risk

Even with Ephemeral Slots defeating entity correlation, behavioral patterns over time can still be correlated:
- Legend semantic patterns (types of relationships, role descriptions)
- Query timing and frequency patterns
- Response latency signatures
- Task type distribution

**Key insight**: Ephemeral Slots defeats **entity** correlation. Behavioral correlation requires separate mitigations.

### V1 Mitigations

| Mitigation | Implementation | Privacy Gain |
|------------|----------------|--------------|
| **Prompt template normalization** | Fixed prompt structures, not free-form | Obscures query complexity patterns |
| **Length caps** | Normalize legend/query to standard buckets | Defeats length-based profiling |
| Token rotation | Daily epoch marker in placeholders | Prevents cross-session slot correlation |
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
- Entity identity (the highest-value target) is protected by Ephemeral Slots

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
| Slot fulfillment fidelity | Find-replace correct? | >99% | Spot-check sample |

### Protocol

1. **Log** all queries, tier decisions, execution results (locally)
2. **Weekly sample** 20 queries for manual quality rating
3. **Compare** to baseline (raw Claude API with full data)
4. **Adjust** thresholds based on findings

### Failure Thresholds

| Metric | Threshold | Action If Breached |
|--------|-----------|-------------------|
| Routing accuracy | < 70% | Simplify classifier; add explicit routing hints to UI |
| Compiler validity | < 85% | Expand test harness; consider stricter DSL grammar |
| Slot fulfillment | < 95% | Require structured output; add validation pipeline |
| Task completion | < 90% | Review failure modes; may need to expand legend detail defaults |

**Review cadence**: Weekly during dogfooding, monthly after stabilization.

### Baseline Comparison

For each query type, run same query through:
- Raw LLM (full data exposure) — quality ceiling
- Privacy tier (DSL/PaaP/Ephemeral Slots/Local) — actual

Quality delta should be <10% for Ephemeral Slots, <20% for Local LLM. DSL/PaaP should be ~0% (logic is logic).

---

## Entity Detection & Legend Construction

### Purpose

Guide entity detection (NER) and legend construction for Ephemeral Slots. The goal is to identify entities reliably and construct minimal legends that provide sufficient context for LLM reasoning.

### Entity Categories (for NER)

| Category | Examples | Slot Format |
|----------|----------|-------------|
| PERSON | names, nicknames, titles | `{{PERSON_kestrel}}` |
| ORGANIZATION | company names, team names | `{{ORG_sparrow}}` |
| PROJECT | project names, codenames | `{{PROJECT_finch}}` |
| LOCATION | addresses, place names | `{{LOCATION_wren}}` |
| EVENT | meeting names, event titles | `{{EVENT_robin}}` |
| DATE | specific dates (not relative) | `{{DATE_hawk}}` or temporal jittering |
| TIME | specific times | `{{TIME_owl}}` or temporal jittering |

### Legend Templates

For each entity category, use minimal context templates:

| Category | Standard Template | Minimal Template |
|----------|-------------------|------------------|
| PERSON | "{{PERSON_X}} is my [relationship]" | "{{PERSON_X}} is a professional contact" |
| ORGANIZATION | "{{ORG_X}} is a [type] I work with" | "{{ORG_X}} is an organization" |
| PROJECT | "{{PROJECT_X}} is a [priority] [type] project" | "{{PROJECT_X}} is a project I'm working on" |
| EVENT | "{{EVENT_X}} is a [type] meeting" | "{{EVENT_X}} is an upcoming event" |

**Legend minimization principle**: Include only what the task requires. For "draft professional email", you need relationship context. For "count meetings", you may need nothing beyond "is a meeting."

### Legend Construction Rules

Formal mapping of (Entity Type × Task Type) → Legend Template. These rules guide the automatic legend generator.

#### Task Types

| Task Type | Classifier Keywords | Context Needed |
|-----------|---------------------|----------------|
| SCHEDULING | "schedule", "conflict", "when", "available" | Time relationships only |
| COUNTING | "how many", "count", "number of" | Existence only |
| DRAFTING | "draft", "write", "email", "message" | Relationship + tone |
| REASONING | "should I", "prioritize", "advise" | Full relationship context |
| SUMMARIZING | "summarize", "overview", "recap" | Role/type only |

#### Legend Template Matrix

| Entity Type | SCHEDULING | COUNTING | DRAFTING | REASONING | SUMMARIZING |
|-------------|------------|----------|----------|-----------|-------------|
| **PERSON** | "is a contact" | "is someone" | "is my [relationship], [tone hint]" | "is my [relationship], [context]" | "is a [role]" |
| **ORGANIZATION** | "is an org" | "is an org" | "is a [type] we work with" | "is a [type], [relationship]" | "is a [type]" |
| **PROJECT** | "is a project" | "is a project" | "is a [priority] project" | "is a [priority] [type] project, [status]" | "is a [type] project" |
| **EVENT** | "is an event at [time bucket]" | "is an event" | "is a [type] meeting" | "is a [type] meeting with [attendee roles]" | "is a [type] event" |
| **LOCATION** | "is a location" | "is a location" | "is a [type] location" | "is [our office / client site / etc.]" | "is a location" |

**Reading the matrix:**
- Columns = task types (detected by classifier)
- Rows = entity types (detected by NER)
- Cells = legend template to use
- `[bracketed]` = fill from entity metadata or escalate

#### Escalation Rules

When to move from minimal to detailed:

```
1. START with column template (task-appropriate minimum)

2. ESCALATE if:
   - LLM response is generic/unhelpful (quality < threshold)
   - Task explicitly requires context (e.g., "relationship advice")
   - User requests more detail

3. ESCALATION PATH:
   COUNTING → SCHEDULING → SUMMARIZING → DRAFTING → REASONING
   (Each step adds context; REASONING is maximum detail)

4. NEVER:
   - Include proper names in legend (that's what slots hide)
   - Include unique identifying details ("the one with the red hair")
   - Jump straight to REASONING level without trying simpler first
```

#### Example: Same Entity, Different Tasks

Entity: Jane Smith (manager, strained relationship)

| Task | Legend Generated |
|------|------------------|
| "How many meetings with Jane?" | "{{PERSON_kestrel}} is someone" |
| "When is my next meeting with Jane?" | "{{PERSON_kestrel}} is a contact" |
| "Summarize my interactions with Jane" | "{{PERSON_kestrel}} is my manager" |
| "Draft an email to Jane about the deadline" | "{{PERSON_kestrel}} is my manager, formal tone preferred" |
| "Should I escalate the issue to Jane?" | "{{PERSON_kestrel}} is my direct manager, relationship is strained, she prefers directness" |

**Key insight**: The same entity gets vastly different legends depending on what the task needs. This is the core of minimization — context is task-scoped, not entity-scoped.

### V1 Hardening Measures

**1. Temporal Jittering**

Round or bucket timestamps unless query requires precision:
- "3:47 PM" → "afternoon" or "late afternoon"
- "Tuesday, January 21" → "next week" or "early next week"
- "Q4 2026" → keep as-is (already coarse)

When to preserve precision:
- Query explicitly asks "what time exactly"
- Scheduling conflicts require exact times
- Default: jitter unless precision required

Benefits: reduces fingerprinting via unique time patterns.

**2. Slot Uniqueness Verification**

Before sending prompt:
- Verify all slot names are unique within the prompt
- Verify no slot names collide with real entity names
- Use sufficiently random generation (bird names, colors, etc.)

**3. Legend Audit Logging**

Log legends locally for user review:
- What context was included for each entity
- User can refine templates based on what felt over-specified
- Builds understanding of what's exposed

### Examples

| User Query | Entities Detected | Legend Constructed |
|------------|-------------------|-------------------|
| "Draft email to Jane about Phoenix delay" | Jane (person), Phoenix (project) | "{{PERSON_kestrel}} is my manager. {{PROJECT_sparrow}} is a high-priority project." |
| "When's my next meeting with the CEO?" | CEO (person) | "{{PERSON_falcon}} is a senior executive I meet with." |
| "Summarize the Acme proposal" | Acme (org) | "{{ORG_finch}} is a potential client." |
| "What should I prioritize today?" | (none — analytical query) | (no legend needed — use Witness-Only or Heuristics) |

---

## Open Questions

### Resolved in This Version

| Question | Resolution |
|----------|------------|
| Compiler output safety | DSL interpreter + Deno sandbox fallback (see [Execution Architecture](#execution-architecture)) |
| Quality measurement | Protocol defined (see [Quality Measurement Methodology](#quality-measurement-methodology)) |
| Legend construction | Template-based with minimization principle (see [Entity Detection & Legend Construction](#entity-detection--legend-construction)) |

### Partially Resolved

| Question | Status | Remaining Risk |
|----------|--------|----------------|
| Entity-level correlation | **Name correlation defeated** — single-use random placeholders | **Description correlation remains** — unique legend content can identify entities without names (see Limitations) |

### Still Open

#### High Priority (Block M3 Completion)

| Question | Why It Matters | Validation Path |
|----------|----------------|-----------------|
| **Residual profiling from legend patterns** | Account-level correlation via semantic content remains possible | Analyze legend entropy across 1000 queries; define acceptable threshold |
| **Rehydration failure rate** | Paraphrases, omissions, possessives break naive substitution | A/B test structured vs free-form output; measure failure rate |
| **Minimum local-LLM capability for PaaP** | Determines PaaP coverage and user quality experience | Benchmark Llama 3 8B, Mistral 7B, Phi-3 on summarization/classification tasks |
| **Threshold for "privacy reliable"** | When to trust Ephemeral Slots for sensitive data (email, contacts) | Define acceptance criteria before M5; may require external audit |
| **Legend minimization calibration** | Over-specify = entity-level leakage via description. Under-specify = quality drop. | V1 ships with conservative defaults; measure quality impact in dogfooding |

#### Medium Priority (Inform V1.5+ Decisions)

1. **Query classification accuracy**: How reliably can we route queries to the right tier? Requires implementation to measure.

2. **NER accuracy on personal data**: How well do off-the-shelf NER models work on personal communications (nicknames, code names, implicit refs)? User correction flow is mandatory, but how much friction?

3. **Provider trust verification**: How do we verify providers aren't doing undisclosed analysis? May require cryptographic approaches (V2+).

### Explicitly Accepted Project Risks

These are the highest-severity risks that **cannot be fully mitigated** in V1. They are accepted with clear eyes.

#### Risk 1: NER is the Achilles' Heel of Ephemeral Slots (CRITICAL)

**The risk**: A single missed entity — a nickname, project codename, or typo — bypasses the entire Ephemeral Slots mechanism and leaks the raw identifier in an otherwise-secure prompt.

**Why it's severe**: This is the most likely vector for an accidental, high-severity data leak in V1. User diligence at "review on first mention" cannot catch entities introduced later in conversation.

**Acceptance rationale**:
- Perfect NER is impossible on personal data (nicknames, code names, implicit references)
- User review provides a defense layer, even if imperfect
- The alternative (require review for every query) destroys UX
- Leaking occasional entity names is less severe than leaking entire context

**Mitigations in place**:
- "Review on first mention" for new entities
- User can add custom entities to watch list
- Logging shows what was sent for post-hoc audit
- For email (M5+), require per-query entity review given higher stakes

**Residual exposure**: A motivated user reviewing their logs may discover occasional entity leaks. This is acceptable for dogfooding. For broader deployment, NER accuracy must improve or stricter review gates are needed.

---

#### Risk 2: PaaP Depends on Local Model Quality (EXTERNAL DEPENDENCY)

**The risk**: The zero-exposure promise for fuzzy tasks via PaaP is contingent on local LLM quality. If Llama/Mistral/Phi-3 cannot reliably follow cloud-generated prompts, the entire path becomes a dead end.

**Why it's severe**: This is an external dependency we cannot control. Local model quality is improving rapidly, but may stagnate or prove insufficient for specific tasks.

**Acceptance rationale**:
- PaaP is explicitly **optional** — V1 can ship without it
- If PaaP fails benchmarks, fuzzy tasks route to Ephemeral Slots (protected, not zero-exposure)
- This bet is on the broader AI industry trajectory, which favors us
- Even partial PaaP success expands zero-exposure coverage

**Mitigations in place**:
- 70% task success benchmark gate before enabling PaaP
- Fallback cascade: PaaP failure → Ephemeral Slots
- User-visible quality warnings when PaaP produces uncertain results
- Logging compares PaaP vs Ephemeral Slots quality for same tasks

**Residual exposure**: If local models disappoint, V1 coverage estimate drops from 60-70% zero-exposure to 50-60%. Privacy guarantee shifts more toward "protected" than "zero." This is acceptable.

---

#### Risk 3: Legend Leakage if Escalation Always Required (QUALITY DEPENDENCY)

**The risk**: The minimization framework assumes minimal legends are often sufficient. If dogfooding reveals that acceptable quality always requires detailed legends (e.g., "is my manager with a strained relationship"), then the privacy benefit of Ephemeral Slots is reduced to merely hiding the name.

**Why it's severe**: This undermines a core claim — that Ephemeral Slots protects not just names but relationship context. If detailed legends are always needed, we're back to "hiding names only."

**Acceptance rationale**:
- We don't know until we try — this is empirical
- Even "names only" protection defeats the most valuable profiling (identity correlation)
- Dogfooding will reveal the truth before broader deployment
- We can adjust legend templates based on findings

**Mitigations in place**:
- Conservative defaults (start minimal, escalate only on quality failure)
- Legend audit logging shows what context was actually needed
- Weekly review of escalation frequency during dogfooding
- Template library can be tuned based on real data

**Residual exposure**: Worst case, Ephemeral Slots provides strong name protection but weaker relationship protection. This is still better than raw data exposure and competitive with any existing approach.

---

### Mitigated Risks (Active Defenses in V1)

| Risk | Category | V1 Mitigations |
|------|----------|----------------|
| **Account-level behavioral profiling** | Mitigated | Template normalization, length caps, timing noise, batching — see below |
| Legend semantic content leakage | Partially mitigated | Minimum viable legend default, escalation framework, user control |

### Accepted Limitations (No V1 Defense)

| Limitation | Why Accepted | Future Mitigation |
|------------|--------------|-------------------|
| Multi-turn conversation identity | Ephemeral Slots are single-query by design | Route follow-ups to Local LLM |
| Motivated nation-state adversary | Different threat class; requires traffic padding, multi-provider | V2+ if needed |

### Behavioral Profiling Mitigations (V1)

Behavioral profiling is the **dominant remaining risk** after Ephemeral Slots. Entity correlation is defeated; behavior correlation is actively mitigated but not fully eliminated.

| Mitigation | Mechanism | V1 Status |
|------------|-----------|-----------|
| Timing noise | Random delays (100-500ms) between queries | Implement |
| Request batching | Batch multiple queries where possible | Implement |
| **Prompt template normalization** | Use fixed prompt structures, not free-form | Implement |
| **Length caps** | Normalize legend/query lengths to standard buckets | Implement |
| Query type mixing | Inject dummy queries of different types | V2 (needs measurement) |
| Traffic padding | Constant-rate request patterns | V2 (bandwidth cost) |

**V1 focus**: Template normalization and length caps are higher impact than timing noise. A provider learns more from "this user sends long complex queries about work relationships" than from timing patterns.

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
| 0.2.1 | 2026-01-18 | Aligned threat model with PRD v0.2 (nation-state, legal, local device explicitly out of scope) |
| 0.3.0 | 2026-01-19 | Aligned sandbox architecture with TDD: Python AST → Deno subprocess with deny-by-default |
| 0.4.0 | 2026-01-19 | Added Hybrid Compilation (V1); added V2+ future work: Foundry, Fuzzy Projections, Crowdsourced Semantics; updated architecture diagram |
| 0.5.0 | 2026-01-19 | Added V1 hardening: Golden-Set testing, Canonicalization, Type Guards, DSL alternative (Compiler); Entropy Minimums, Temporal Jittering, Active Rehydration (Projection). Expanded V2+ ideas. |
| 0.6.0 | 2026-01-19 | DSL now primary execution approach (SymQL); Deno sandbox demoted to fallback. Added DSL grammar spec, constrained decoding option. |
| 0.7.0 | 2026-01-19 | Added Amnesic Reasoning (V2) — stateless micro-query orchestration to prevent cross-entity linking. |
| 0.8.0 | 2026-01-19 | Added V2+ ideas: Model Sharding, Encrypted Semantic Indices. Added Lens Schema guidance to Projection. |
| 0.9.0 | 2026-01-19 | Added Prompt Camouflage (V2+) — noise injection for classification/analysis tasks only; not suitable for generation. |
| 0.10.0 | 2026-01-19 | Added Constraint-Solver Delegation (V2+) — formal logic alternative to DSL for complex queries. |
| 0.11.0 | 2026-01-19 | Added Two-Tier Reasoning (V1.5) — structure/content separation for generation tasks. |
| 0.12.0 | 2026-01-19 | Added Ephemeral Slot-Fill Sessions as V1 alternative to auto-Projection — manual, simpler, no type vocabulary. |
| 0.13.0 | 2026-01-19 | Added Witness-Only Reasoning (computed facts) and Heuristic-First + Templates (no-LLM path) as V1 alternatives. |
| 0.14.0 | 2026-01-19 | Added UX-Centric Privacy Controls: Progressive Consent Ladder (V1), Privacy Receipts + Undo (V1), Trust Dial (V1.5), Privacy Sandbox Mode (V2). |
| 0.15.0 | 2026-01-19 | Added Security-Centric Privacy Controls: Policy-as-Code Gate (V1), Strict Egress Firewall (V1), Deterministic Redaction (V1 narrow), Minimal-Exposure Proofs (V2). |
| 0.16.0 | 2026-01-19 | Added more Security Controls: Taint-Tracked Serialization (V1), Schema-Hash Gate (V1), Tamper-Evident Audit Chain (V2). |
| 0.17.0 | 2026-01-19 | Added more UX Controls: Consent Recipes (V1), Field-Reveal Chip UI (V1), Time-Boxed Access (V1), Privacy Budget Meter (V1.5). |
| **1.0.0** | **2026-01-19** | **MAJOR RESTRUCTURE**: Ephemeral Slots replaces Semantic Projection as primary V1 privacy layer. Core insight: single-use random placeholders + legend defeats entity-level profiling by preventing cross-query correlation. Simpler, more private, broad task coverage (see limitations). Semantic Projection moved to V2+ as opt-in enhancement. Updated architecture diagram, implementation sequence, all references. |
| 1.1.0 | 2026-01-19 | Added "The Ephemeral Slots Roadmap" — vision for V1.5 through V3: Dynamic Legend Optimization, Composable/Nested Slots, Functional Slots, and The Ephemeral Self. Shows how Ephemeral Slots scales from technique to paradigm. "Ghost in the machine" framing. |
| **2.0.0** | **2026-01-19** | **MAJOR RESTRUCTURE**: "LLM as Compiler" becomes "Query Compilation (Zero-Data Exposure)" with two paths — DSL Compilation (structured queries, SymQL) and Prompt-as-Program (fuzzy data-processing). PaaP insight: prompts ARE programs — cloud LLM generates optimal prompt, local LLM executes on real data. Zero exposure extended to summarization, classification, extraction tasks. Updated architecture diagram to show both compilation paths. |
| 2.1.0 | 2026-01-19 | Expanded V3 Ephemeral Self with implementation details: three-step construction pipeline (Seed → Traverse → Ephemeralize), consistency model ("burn after reading"), structural fingerprinting defenses (Coarsening, Noise Injection, Pruning). Noted Local Knowledge Graph as V2-V3 dependency. |
| 2.2.0 | 2026-01-19 | **Honest claims reframing**: Fixed overstated privacy guarantees. "Defeats entity-level correlation" not "impossible to profile". Added explicit account-level behavioral profiling as accepted limitation. Added Limitations & Residual Risks section. Added Rehydration Requirements (structured output mandatory, escaping spec). Downgraded NER confidence to "medium". Updated coverage estimates (60-70% target, marked as hypothesis). Expanded Open Questions with high-priority validation gates. Removed "universal" and "trivial" claims. |
| 2.3.0 | 2026-01-19 | **Further hardening**: (1) Legend now explicitly "minimum viable" — escalate only if quality degrades. (2) Entity-level profiling via unique descriptions called out as remaining risk. (3) Rehydration fallback cascade defined (retry → local → partial with warning). (4) NER review friction addressed with "review first mention only" default. (5) Legend minimization elevated to M3 gating criterion. (6) PaaP explicitly optional — skip if local LLM quality < threshold. (7) Behavioral profiling mitigations expanded (template normalization, length caps). (8) Entity-level correlation reclassified as "partially resolved". (9) Updated correlation mitigations to reference Ephemeral Slots, not projection. |
| 2.4.0 | 2026-01-19 | Added **Interactive Scaffolding (V1.5)**: Meta-innovation for complex/uncertain queries. Instead of routing to single massive Ephemeral Slots prompt, system engages cloud LLM in abstract planning dialogue (zero content) to build custom execution plan composing DSL + PaaP + Ephemeral Slots. Makes "uncertain" the smartest path, not a fallback. Includes two-phase process (planning dialogue → execution program), privacy model, concerns/mitigations, success criteria. Updated architecture diagram. |
| 2.5.0 | 2026-01-19 | **Consistency review**: (1) Reframed "Why It Replaces Semantic Projection" as design history. (2) Removed 80%+ claim from v2.0.0 version entry. (3) Added detail escalation decision framework with thresholds. (4) Reconciled NER confidence with M5 gate via per-tier requirements. (5) Clarified PaaP optional status in coverage table. (6) Reframed behavioral profiling as "mitigated risk" not "accepted limitation". (7) Added "Privacy by Default" routing principle. (8) Made Deno sandbox permissions explicit. (9) Added Interactive Scaffolding success criteria table. (10) Added failure thresholds to quality measurement. (11) Fixed outdated "projection" reference in baseline comparison. (12) Fixed inconsistent coverage claims (60-70% → 40-50% for DSL). (13) Fixed "impossible" → "defeated" in Amnesic Reasoning. (14) Fixed "universal" → "broad" in v1.0.0 version entry. |
| **3.0.0** | **2026-01-19** | **Critical review response**: (1) Added **Query Classifier Specification** — keyword cascade algorithm with confidence thresholds, classification categories, order of operations, failure modes. (2) Added **Cost of Privacy Model** — latency/API cost/quality/privacy tradeoffs per execution path. (3) Added **User-Facing Failure Model** — "fail visibly, fail safely, offer alternatives" with UX for each failure type. (4) Added **PaaP Schema v1.0** — formal JSON structure for single-step and pipeline programs. (5) Added **Legend Construction Rules** — (Entity Type × Task Type) → Template matrix with escalation path. (6) Added **Hybrid Compilation State Machine** — formal state diagram with transitions and fallback routing. (7) Added **Explicitly Accepted Project Risks** section — NER as Achilles' heel (CRITICAL), PaaP external dependency, Legend leakage quality dependency. (8) Fixed Deno sandbox example (Python → TypeScript). (9) Added V2+ navigation notes. |

---

*Created: 2026-01-18*
*Last updated: 2026-01-19*
