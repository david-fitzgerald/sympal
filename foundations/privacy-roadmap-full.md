# Privacy Roadmap — Full V2+ Research

**Version**: 1.0.0
**Date**: 2026-01-21
**Status**: V2+ reference (extracted from privacy-innovations.md)
**Purpose**: Detailed research and design for post-V1 privacy features

> This document contains the full V2+ privacy research originally developed in privacy-innovations.md.
> V1 implementation details remain in privacy-innovations.md.
> ROADMAP.md provides the summary index.

---

## Table of Contents

1. [Ephemeral Slots Evolution](#ephemeral-slots-evolution)
2. [P2P Query Mixing](#p2p-query-mixing)
3. [UX Controls V2](#ux-controls-v2)
4. [Security Controls V2](#security-controls-v2)
5. [Advanced Privacy Techniques](#advanced-privacy-techniques)
6. [Correlation Mitigations V2](#correlation-mitigations-v2)

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


---

## UX Controls V2

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

---

## Security Controls V2

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

---

## Advanced Privacy Techniques

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


---

## Correlation Mitigations V2

### V2 Mitigations (If Needed)

| Mitigation | Implementation | Privacy Gain |
|------------|----------------|--------------|
| Query padding | Add plausible dummy queries | Obscures true query volume |
| Multi-provider distribution | Split queries across providers | No single provider sees all patterns |
| Decoy sessions | Periodic automated query bursts | Camouflages real usage patterns |


---

## Version History

| Version | Date | Changes |
|---------|------|---------||
| 1.0.0 | 2026-01-21 | Initial extraction from privacy-innovations.md v3.0.0 |

