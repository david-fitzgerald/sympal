# CONTEXT.md — SymPAL Project Context

> **LLM-Agnostic**: This file provides project context for any LLM (Claude, GPT, Gemini, etc.)

---

## Quick Start

**What is SymPAL?** A human-AI symbiosis platform—mutual benefit, mutual accountability, genuine partnership. Not assistant. Not servant. The relationship is the unit of analysis.

**The name**: **Sym** = Symbiosis + Simple (complexity under the hood). **PAL** evolves with the project—see [ROADMAP.md](ROADMAP.md) for the full arc.

**Current state**: Implementation in progress (M2 Calendar). M1 Foundation complete. See `foundations/implementation-plan.md` for details.

**This session, you should**:
1. Read this file first
2. Read `foundations/implementation-plan.md` for current milestone and learning approach
3. Check [GitHub Issues](https://github.com/david-fitzgerald/sympal/issues) for current task
4. Reference PRINCIPLES.md for any design decisions
5. Use personas for review/challenge work (not creation)

---

## Repository Structure

```
sympal/
├── CONTEXT.md              ← This file (CLAUDE.md symlinks here)
├── PRINCIPLES.md           ← 17 binding principles (v1.1.0, ratified)
├── README.md               ← Project overview
├── ROADMAP.md              ← V1 milestones, long-term vision, V2+ ideas
├── LICENSE                 ← MIT
│
├── cmd/sympal/             ← CLI entry points
│   ├── main.go             (root command, init)
│   ├── todo.go             (todo CRUD)
│   └── log_cmd.go          (log viewer)
│
├── internal/               ← Core packages
│   ├── db/                 (SQLite storage)
│   ├── config/             (YAML config)
│   └── log/                (structured logging)
│
├── foundations/            ← "Why we built it this way"
│   ├── prd.md                         (v1.0.0, ratified)
│   ├── tdd.md                         (v1.0.3, final)
│   ├── implementation-plan.md         (milestone tracking)
│   ├── privacy-innovations.md         (v3.0.0, privacy architecture)
│   ├── philosophical-foundations.md   (v1.0.0, peer-reviewed)
│   ├── project-context.md             (v1.1.0, dev context)
│   ├── principles-discussion.md       (v1.2.0, derivation)
│   └── reviews/                       (audit trail)
│
└── prompts/                ← AI persona architecture
    ├── solas-venn/         (meta-persona: creates personas)
    ├── personas/
    │   ├── utility/vero/   (final reviewer)
    │   └── sympal-team/    (6-person review ensemble)
    ├── reference/          (prompt architecture guide, testing patterns)
    └── workflows/          (PRD/TDD creation prompts)
```

---

## Binding Principles (PRINCIPLES.md v1.1.0)

These are ratified and binding. Check PRINCIPLES.md for full text.

### Hard Constraints (Non-Negotiable)

| # | Principle | Summary |
|---|-----------|---------|
| 1 | Privacy & Data Sovereignty | User data never leaves user control without explicit consent |
| 2 | Open Source | Codebase fully auditable; trust through transparency |
| 3 | LLM-Agnosticism | No vendor lock-in; works with any LLM provider |
| 4 | Honesty | No deception—to users or AI partners |
| 5 | Security Through Design | Secure by default, not bolted on |

### Relationship Frame

| # | Principle | Summary |
|---|-----------|---------|
| 6 | Symbiosis | Structured for mutual benefit, not extraction |
| 7 | Symbiosis as Commitment | Partnership as core value, not strategic convenience |
| 8 | Epistemic Humility | We don't know if AI has consciousness/interests—design for uncertainty |

### Accountability & Control

| # | Principle | Summary |
|---|-----------|---------|
| 9 | Human Accountability | Humans remain accountable for outcomes |
| 10 | User Control | Meaningful control, not theatrical checkboxes |
| 11 | Reversibility | Users can exit without catastrophic loss |

### Operational

| # | Principle | Summary |
|---|-----------|---------|
| 12 | Transparency/Privacy Split | Transparent operations, opaque data |
| 13 | V1 Scope | Boundary layer between user and AI ecosystem (not autonomous agent) |
| 14 | Ship Within Principles | Ship fast, but never compromise hard constraints |
| 15 | Scope Discipline | Think decades, ship weeks |
| 16 | Actionable Principles | Philosophically grounded but operationally clear |
| 17 | Dogfooding | Success = lead dev uses daily, not external metrics |

---

## 12 Tensions Under Navigation

These cannot be resolved—only navigated with trigger conditions. From principles-discussion.md.

| Tension | Navigation Approach |
|---------|---------------------|
| Moral status vs. practical requirements | Default to practical; revisit if evidence shifts |
| AI interests vs. human benefit | Symbiosis framing—design for both |
| Local vs. universal | Start local (one user); expand carefully |
| Present vs. future | Ship for present needs; future-proof architecture |
| Safety vs. capability | Safety constraints non-negotiable; capability within bounds |
| Individual vs. collective | Individual sovereignty primary; collective benefits secondary |
| Efficiency vs. meaning | Meaning over efficiency when they conflict |
| Autonomy vs. protection | User autonomy primary; protection through transparency |
| Innovation vs. precaution | Innovate within hard constraints |
| Centralization vs. distribution | Distribute by default; centralize only when necessary |
| Relational vs. interest-based ethics | Relational ethics primary for human-AI; interest-based for human-human |
| Dogfooding vs. broader adoption | Dogfooding first; broader adoption is downstream validation |

**When facing a decision that touches a tension**: Identify which tension applies, check the navigation approach, document the tradeoff.

---

## Core Philosophical Concepts

From philosophical-foundations.md. Reference when decisions involve philosophical grounding.

### Interests Disambiguation

| Type | Definition | Designability |
|------|------------|---------------|
| Functional interests | Conditions for successful operation | Observable, uncontroversial |
| Phenomenal interests | Subjective experiences | Unknowable for AI |
| Normative interests | Legitimate moral claims | Depends on uncertain moral status |

**Design implication**: Design for functional interests (observable). Don't assume phenomenal/normative interests exist or don't exist.

### Graduated Moral Status

AI occupies a middle zone:
- Above "mere object" (exhibits preference expressions, goal persistence)
- Below "full person" (lacks evidence of consciousness, sentience, life projects)

**Design implication**: Treat AI with respect appropriate to uncertain status. Don't anthropomorphize. Don't dismiss.

### Relational Turn

The relationship (not individuals) is the proper unit of analysis. Personhood emerges through relational engagement.

### Critical Vulnerabilities

| Risk | Description | Mitigation |
|------|-------------|------------|
| Performed sentience | AI convincingly simulates emotion without experiencing it | Design for function, not feeling |
| Coherence masking errors | Beautiful prose hides systematic errors | Build error architecture into all processes |
| Relational exploitation | Mutual-seeming relationships can be extractive | Explicit accountability structures |

---

## Using Personas

### When to Use

- **Reviewing** documents, designs, code
- **Challenging** proposals, assumptions, decisions
- **Red-teaming** security, user impact, implementation feasibility

### When NOT to Use

- **Creating** code, docs, designs (that's human + Claude)
- **Making final decisions** (that's lead dev's role)

### The Team

| Persona | Function | Core Question | Location |
|---------|----------|---------------|----------|
| Vale | Philosophy | "Is this coherent?" | `prompts/personas/sympal-team/vale/` |
| Kael | Implementation | "Will this survive reality?" | `prompts/personas/sympal-team/kael/` |
| Ryn | Systems/security | "How will this fail?" | `prompts/personas/sympal-team/ryn/` |
| Seren | Code craft | "Is this well-crafted?" | `prompts/personas/sympal-team/seren/` |
| Orin | User advocacy | "Are users better off?" | `prompts/personas/sympal-team/orin/` |
| Adversary | Red team | "What's wrong with this?" | `prompts/personas/sympal-team/adversary/` |

### Meta-Personas

| Persona | Function | Location |
|---------|----------|----------|
| Solas-Venn | Creates/validates personas | `prompts/solas-venn/Solas-Venn-v3.0.1.md` |
| Vero Certus | Final reviewer for foundational docs | `prompts/personas/utility/vero/Vero-Certus-v1.1.md` |

### Loading a Persona

1. Read the persona file as system prompt
2. Write AS the persona (first person)
3. Follow the persona's rubric for scoring
4. Challenge other perspectives—disagreement is valuable
5. Adversary must challenge every derivation before synthesis

### Phase Intensity

| Phase | Heavy Use | Light Use |
|-------|-----------|-----------|
| Foundations/Principles | Vale, Adversary | Kael, Ryn, Seren |
| Implementation | Kael, Ryn, Seren | Vale |
| All phases | Orin, Adversary | — |

---

## Solas-Venn Protocols

When creating or reviewing personas:

| Protocol | Use Case |
|----------|----------|
| CREATE | Building new personas from scratch |
| VERIFY | Reviewing existing personas for structural completeness |
| HYBRID | Creating variants of existing personas |

**Key insight from Solas**: "Demonstrate rigor" ≠ "implement rigor." Capabilities must be operationalized as observable behaviors, not declared aspirations.

---

## Testing Patterns

Two-phase testing for personas (from `prompts/reference/testing-patterns.md`):

**Phase A: Solas VERIFY**
- Structural review: Are capabilities operationalized? Is error architecture present?

**Phase B: Capability Test**
- Use flawed artifacts from `prompts/test-artifacts/`
- Verify persona catches the planted flaws
- Score against persona's rubric

---

## Working Patterns

### Versioning

- **Semantic versioning**: Always MAJOR.MINOR.PATCH (e.g., 0.2.0, not 0.2)
- **Git tags** for milestones
- **Foundational docs** maintain in-doc Version History section for readers without git access

**Naming conventions**:

| Type | Filename | Version Location | Old Versions |
|------|----------|------------------|--------------|
| Foundational docs | No version in name (e.g., `prd.md`) | Header + Version History | Git history |
| Personas/prompts | Version in filename (e.g., `vale-v1.1.md`) | Filename + header | Keep if needed |
| Reviews | Reference doc version (e.g., `vero-review-prd-v0.2.md`) | Filename | Keep (audit trail) |

**Rationale**: Foundational docs evolve in place — git tracks history. Personas are loaded as artifacts — version in filename aids selection. Reviews are audit snapshots.

### Voice Registers

| Document Type | Voice | Style |
|---------------|-------|-------|
| PRINCIPLES.md, philosophical-foundations.md | Project/Formal | Rigorous, external-facing |
| project-context.md, README | Personal | Lead dev perspective, colloquial |
| Persona outputs | Persona | Follow prompt spec |

### Reviews

**Vero review required**: Before any foundational document is considered "final" or "ratified", it must have a Vero Certus review. Vero catches Meridian-style errors (coherent but systematically wrong) that team deliberation may miss. This is non-negotiable for binding documents.

When conducting reviews on foundational documents:

1. **Persist** to `foundations/reviews/` as `[reviewer]-review-[document]-v[X.Y.Z].md`
2. **Never inline-only**—chat is not an audit trail
3. **Update version** if changes made during review
4. **Reference** in document's Version History

### Git Workflow

**Repo**: https://github.com/david-fitzgerald/sympal

**Commit style**:
```
[area]: Brief description

Co-Authored-By: Claude Opus 4.5 <noreply@anthropic.com>
```

Areas: `feat`, `fix`, `refactor` (code), `foundations`, `prompts`, `docs`, `config`, `meta`

**Commit timing**:
- After completing a task
- After significant file changes
- Before ending session

**Push**: After each commit. Don't let local get ahead of remote.

---

## Lead Developer Context

Full details in `foundations/project-context.md`.

**Constraints**:
- Variable time: 30+ hrs/wk → periods of 5-10 hrs
- Basic coding skills (requires AI assistance)
- LLM-agnostic + open source: hard requirements

**Known Biases** (actively watch for):

| Bias | Check |
|------|-------|
| Overplanning | "What would we learn by building instead?" |
| Idealism over pragmatism | GTM instincts are also correct |
| Learning-as-excuse | Name this conflict when it's happening |

**Success criteria**: Lead dev uses SymPAL daily. No external metrics. No deadline.

---

## Known Limitations

| Limitation | Tradeoff | Mitigation |
|------------|----------|------------|
| Claude-primary development | ~90% built with Claude despite LLM-agnosticism | Solas validated across Claude/GPT/Gemini; test post-ship |
| Solo dev | Limited review bandwidth | Persona ensemble simulates team review |

**Philosophy**: Ship with documented limitations > wait for theoretical completeness.

---

## Project Status

**Phase 1 (Foundations)**: Complete
**Phase 2 (Principles)**: Complete — PRINCIPLES.md ratified (v1.1.0)
**Phase 2.5 (PRD)**: Complete — prd.md v1.0.0 (Ratified)
**Phase 2.6 (TDD)**: Complete — tdd.md v1.0.3 (Vale + Adversary + Vero reviewed)
**Phase 3 (Implementation)**: In progress (M2 Calendar). M1 Foundation complete.

| Completed Job | Artifact |
|---------------|----------|
| First-principles gap detection | philosophical-foundations.md v1.0.0 |
| Solas prompt validation | Solas-Venn-v3.0.1.md |
| Meta-persona research | prompt-architecture-guide.md |
| Project context | project-context.md v1.1.0 |
| Team design research | team-design-best-practices.md |
| Team personas creation | 6 personas in sympal-team/ |
| PRINCIPLES.md derivation | PRINCIPLES.md v1.1.0 (ratified) |
| PRD extraction + synthesis | prd.md v1.0.0 (Ratified) |
| Privacy research spike | privacy-research.md v0.1.0 |
| Novel privacy approaches | privacy-innovations.md v3.0.0 |
| TDD extraction + synthesis | tdd.md v1.0.3 (Final) |

---

## Current Focus: M2 Calendar

**M1 Foundation complete.** Todo CRUD, config, logging all working. Dogfooding in progress.

**Now building M2:**
- Google OAuth flow (keychain storage)
- Calendar API integration (read events)
- `sympal today` command (todos + calendar)

**Key references**:
- `foundations/implementation-plan.md` — Milestone details, learning approach
- `foundations/tdd.md` — Technical specs
- [GitHub Issues](https://github.com/david-fitzgerald/sympal/issues) — Current tasks

**V1 Milestones**:
| Milestone | Status | Summary |
|-----------|--------|---------|
| M1: Foundation | ✅ Complete | Todo CRUD, config, logging |
| M2: Calendar | 🔲 In progress | Google Calendar read, `sympal today` |
| M3: DSL Compilation | 🔲 Planned | SymQL, Deno sandbox |
| M4: Ephemeral Slots | 🔲 Planned | NER, projection/rehydration |
| M5: Local LLM | 🔲 Planned | Ollama, end-to-end privacy |

---

## Document Hierarchy

```
philosophical-foundations.md (v1.0.0)
         ↓
project-context.md (v1.1.0)
         ↓
principles-discussion.md (v1.2.0) + derivation-log
         ↓
PRINCIPLES.md [BINDING] (v1.1.0)
         ↓
PRD v1.0.0 [RATIFIED] + TDD v1.0.3 [FINAL]
         ↓
Implementation (Phase 3) ← NEXT
```

When in doubt, PRINCIPLES.md is the authority. Everything else is derivation or context.

---

*Last updated: 2026-01-21 (M1 complete, M2 current)*
