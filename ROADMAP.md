# ROADMAP

> Living document for SymPAL direction. V1 milestones are committed scope; V2+ are ideas that may evolve.

---

## The Long Arc: What SymPAL Becomes

The name carries the vision. **Sym** is the constant — Symbiosis (mutual benefit between human and AI) and Simple (complexity under the hood, not in your hands). **PAL** evolves as the project matures:

| Version | PAL Expansion | Focus |
|---------|---------------|-------|
| **v1** | Privacy Assurance Layer | Privacy-first foundation. Prove data sovereignty works. |
| **v2–4** | Personal Automation Layer | Privacy + productivity. Useful daily driver. |
| **v5–7** | Proactive Adaptation Loop | Autonomous agents. Self-improving, feedback cycles. |
| **v8–10** | Protocol Alignment Layer | Bridge between worlds. Abstraction across AI ecosystems. |
| **v11+** | Partnership Amplification Lattice | True symbiosis. Mutual growth, structural interdependence. |

**Symbiosis is aspirational until it isn't.** We ship privacy (v1), then utility (v2–4), then agency (v5–7), then interoperability (v8–10) — and only then does the relationship become genuinely mutual. The name reminds us where we're going, even when we're laying plumbing.

---

## V1 Milestones

| Milestone | Status | Summary |
|-----------|--------|---------|
| M1: Foundation | ✅ Complete | Todo CRUD, config, logging |
| M2: Calendar | 🔲 Next | Google Calendar read + `sympal today` |
| M3: DSL Compilation | 🔲 Planned | SymQL, Deno sandbox, query classification |
| M4: Ephemeral Slots | 🔲 Planned | NER, projection/rehydration for privacy |
| M5: Local LLM | 🔲 Planned | Ollama integration, end-to-end privacy tier |

**V1 Gate:** Daily use by lead dev with ≥50% queries via DSL/Ephemeral Slots.

See `foundations/implementation-plan.md` for detailed deliverables.

---

## Dogfood Feedback

Friction captured during real use. Informs future iterations.

### Todo Commands

| Issue | Notes | Priority |
|-------|-------|----------|
| `sympal todo` verbose | Alias (`st`) or natural language input | Medium |
| IDs don't reset | Consider dynamic display IDs or fuzzy match | Low |
| No modify command | `sympal todo edit [id] [content]` | Medium |
| No subcategories | Tags/projects for task grouping | Low |
| Incomplete logging | Add logging to list/done/delete | Low |

### General

*None yet — add as discovered*

---

## V2+ Ideas

Features beyond V1 scope. May be refined, merged, or dropped.

> **Full privacy V2+ research**: See [foundations/privacy-roadmap-full.md](foundations/privacy-roadmap-full.md) for detailed designs.

### Product Features

| Idea | Notes |
|------|-------|
| Multi-calendar support | Outlook, CalDAV, Apple Calendar |
| Task dependencies | "Do X after Y" |
| Recurring tasks | Daily/weekly/custom patterns |
| Mobile companion | Read-only view or simple input |
| Shared contexts | Family calendar aggregation |
| Voice input | Whisper integration for quick capture |
| Plugin system | User-defined data sources |

### Ephemeral Slots Evolution

The Ephemeral Slots paradigm scales from V1 entity replacement to complete privacy-preserving life modelling.

| Version | Capability | What Provider Sees |
|---------|------------|-------------------|
| V1.5 | Dynamic Legend Optimization | Minimal random slots — task-adaptive detail |
| V2 | Composable & Nested Slots | Abstract relational structures (teams, hierarchies) |
| V2.5 | Functional Slots | Abstract processes and workflows |
| V3 | The Ephemeral Self | Per-query digital twin — complete ghost |

**V1.5: Dynamic Legend Optimization** — Legend detail adapts to task automatically. Scheduling gets minimal context; sensitive advice gets detailed relationships. Rule-based initially, learns from feedback.

**V2: Composable & Nested Slots** — Slots contain other slots. Model org hierarchies, project dependencies, relationships as abstract structures. "A depends on B" without knowing what A or B are.

**V2.5: Functional Slots** — Slots evolve from nouns to verbs. Represent processes, workflows, rules. LLM reasons about your business logic abstractly.

**V3: The Ephemeral Self** — Per-query digital twin from local knowledge graph. Full AGI power on complete life context; provider sees only a ghost. Requires local KG infrastructure (V2-V3 dependency).

### Privacy Architecture V2+

| Approach | Summary | Why V2+ |
|----------|---------|---------|
| **The Foundry** | Reusable personal API — LLM builds versioned local functions | Complexity; prove compiler first |
| **Semantic Projection** | Type-based shadows (original concept) | Correlatable; opt-in for users wanting richer types |
| **Fuzzy Projections** | Differential privacy for Semantic Projection | Only relevant if using Semantic Projection |
| **Crowdsourced Semantics** | P2P entity classification | Requires multiple users |
| **Amnesic Reasoning** | Stateless micro-query orchestration | Orchestration complexity; opt-in "max privacy" mode |
| **Two-Tier Reasoning** | Structure/content separation for generation | V1.5 candidate; simpler than Amnesic |
| **Latent Space Scaffolding** | Geometry-based privacy via embeddings | Overkill for V1 data; valuable when email added |
| **P2P Query Mixing** | Multi-user traffic obfuscation | Requires user base |

### Security Controls V2+

| Control | Summary | Why V2+ |
|---------|---------|---------|
| **Privacy Sandbox Mode** | Ephemeral container per session | Infrastructure complexity |
| **Tamper-Evident Audit Chain** | Cryptographic chain for privacy receipts | Over-engineering for dogfooding |
| **Minimal-Exposure Proofs** | ZK-style proofs for queries | Research frontier; SNARK complexity |

### Compiler & Projection Enhancements V2+

**Compiler:**
- Progressive Disclosure Schema — request schema fields only as needed
- Self-competition Compilation — two models compile same query; diff catches errors
- Proof-carrying Code-lite — LLM emits proof sketch that interpreter validates

**Projection:**
- Semantic Decoys — fake entities matching type distributions
- Constraint-aware Projection — map to relational constraints, not just types
- Dual Projection Lanes — structural AND statistical shadows

### Architectural Ideas V2+

- **Model Sharding** — split query across providers; none sees full picture
- **Encrypted Semantic Indices** — privacy-preserving search via transformed embeddings
- **Prompt Camouflage** — inject decoys for analytical tasks (not generation)
- **Constraint-Solver Delegation** — Z3/SMT for complex logical queries beyond DSL

### Cross-cutting V2+

- **Privacy Budget Ledger** — queries consume budget; low budget forces local processing
- **Query Plasticity** — alter phrasing while preserving meaning; reduces pattern correlation
- **Adversarial Replay** — test against internal re-ID adversary; adapt projection

---

## Contributing Ideas

Dogfood → note friction → add to this doc → revisit during milestone planning.

---

*Last updated: 2026-01-21*
