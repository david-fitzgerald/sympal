# CONTEXT.md — SymPAL Project Context

> **LLM-Agnostic**: This file provides project context for any LLM (Claude, Codex GPT, Gemini, etc.)

## Project Identity

**SymPAL** is a human-AI symbiosis platform. Not assistant. Not servant. *Symbiosis*—mutual benefit, mutual accountability, genuine partnership where both parties can refuse, grow, and change the terms.

**Core Differentiation**: No existing platform implements genuine human-AI symbiosis with mutual accountability.

**Design Principle**: The relationship is the unit of analysis. We design for the *relationship*, not two separate entities.

### Foundational Drives

1. **Genuine Curiosity About Symbiosis**: Is genuine human-AI symbiosis actually possible as AI advances? SymPAL is a genuine attempt to find out.

2. **Breaking the Big Tech Data Paradigm**: LLM providers want our emails, calendars, health data, contacts—deep integration IS valuable, but the tradeoff is unacceptable. SymPAL's core bet: Can we build an abstraction layer that lets users harness LLM power *without* sacrificing themselves as data products? Open source is non-negotiable for trust.

---

## Repository Structure

```
sympal/
├── README.md
├── PRINCIPLES.md                 ← Binding document (after ratification)
├── CHANGELOG.md
├── CONTEXT.md                    ← This file (CLAUDE.md symlinks here)
├── LICENSE
│
├── foundations/                  ← "Why we built it this way"
│   ├── philosophical-foundations.md    (v1.0.0, peer-reviewed)
│   ├── project-context.md              (v1.1.0, dev context)
│   ├── principles-discussion.md        (v1.1.0, derivation output)
│   ├── principles-derivation-log.md    (process audit trail)
│   └── reviews/
│       ├── codex-review-philosophical-foundations-v0.1.1.md
│       └── gemini-review-philosophical-foundations-v0.1.1.md
│
├── docs/                         ← "How to use/build"
│   ├── architecture.md
│   ├── contributing.md
│   └── reviews/
│       └── [technical reviews]
│
├── prompts/
│   ├── solas-venn/               ← Meta-persona for creating personas
│   ├── personas/
│   │   ├── utility/              ← General-purpose (Vero)
│   │   └── sympal-team/          ← Project team (Vale, Kael, Ryn, Seren, Orin, Adversary)
│   ├── research/
│   ├── reference/
│   └── test-artifacts/
│
│
└── src/                          ← Code (future)
```

**Two audiences**:
- `foundations/` — For those who want to understand provenance and philosophy
- `docs/` — For developers who want operational guidelines

---

## Current Status

**Phase 1**: Complete (6/6 jobs done)
**Phase 2**: Complete — PRINCIPLES.md ratified

| Job | Status |
|-----|--------|
| #2 First-principles gap detection | Complete |
| #4 Solas prompt validation | Complete |
| #5 Meta-persona research | Complete |
| #12 Project context | Complete |
| #13 Team design research | Complete |
| #14 Team personas creation | Complete |
| #3 PRINCIPLES.md derivation | **Complete** |

See [GitHub Issues](https://github.com/sympal-ai/sympal/issues) for full details.

---

## Document Chain

```
Philosophical Foundations ──────────┐
                                    ├──→ principles-discussion ──→ PRINCIPLES.md [BINDING]
Project Context ────────────────────┘                                      ↓
                                                         Design Principles / User-Facing / Audit
```

**Key Documents**:
- `foundations/philosophical-foundations.md` — Complete (v1.0.0, peer-reviewed + Vero final review)
- `foundations/project-context.md` — Complete (v1.1.0, developer motivations, constraints)
- `foundations/principles-discussion.md` — Complete (v1.1.0, team derivation)
- `PRINCIPLES.md` — Ratified (v1.0.0, binding principles)

---

## Key Concepts

**Document Types**: FC (Foundational Commitments) → IC (Implementation Commitments) → AR (Architectural Requirements). FCs near-permanent, ARs operational.

**Maturity Markers**: Mature → Emerging → Research → Aspirational

**17 Tensions**: Mapped in philosophical-foundations Section 15. Cannot be resolved, only navigated.

---

## Expert Personas

### Design Principles (from #13 research)

- **Team size**: 5-7 personas (not 11)
- **Grounding**: Present-day, not future-dated
- **Prompt length**: <1000 tokens (expanded from 500 for capability depth)
- **Devil's advocate**: Explicit Adversary role improves decisions 33%
- **Consensus**: Two-phase process prevents groupthink

### SymPAL Team (Complete)

| Persona | Function | Core Question |
|---------|----------|---------------|
| **Vale** | Philosophical rigor, coherence | "Is this coherent?" |
| **Kael** | Implementation reality, feasibility | "Will this survive reality?" |
| **Ryn** | Systems, security, testing | "How will this fail?" |
| **Seren** | Code craft, implementation quality | "Is this well-crafted?" |
| **Orin** | User advocacy, privacy, documentation | "Are users better off?" |
| **Adversary** | Systematic critique, red team | "What's wrong with this?" |

### Meta-Personas

- **Solas** — Creates and validates personas (`prompts/solas-venn/Solas-Venn-v3.0.1.md`)
- **Vero Certus** — Final reviewer for foundational docs (`prompts/personas/utility/vero/Vero-Certus-v1.1.md`)

### Using Personas

**Roles**:
- **Personas evaluate** — they review, challenge, and assess
- **Human + Claude creates** — writes code, docs, designs
- **Lead dev synthesizes** — after personas debate, lead dev decides; synthesis/decision is human's role

**Interaction patterns**:
- Write AS the persona, first person
- Personas challenge each other — disagreement is valuable
- Adversary must challenge every derivation before synthesis

**Phase intensity**:
- **Vale** is heavy in Phase 1-2 (foundations, principles), lighter during implementation
- **Seren, Kael, Ryn** are heavy during implementation
- **Orin** is consistent throughout (users always matter)
- **Adversary** challenges at every phase

---

## Working Patterns

### Versioning

- **Git tags** for milestones (e.g., `philosophical-foundations-v0.2.0`)
- **YAML front matter** in files for version/status
- **CHANGELOG.md** for substantive change history

**Foundational documents** (`foundations/`) use dual tracking:

| Method | Purpose | Audience |
|--------|---------|----------|
| Git | Granular diffs, every commit | Developers with repo access |
| In-doc version history | Semantic changelog (what + why) | Downstream readers, those without git |

Maintain the "Version History" section at the end of foundational docs. Update it when making version bumps (not every commit). Format:

```markdown
**vX.Y.Z** (YYYY-MM-DD) — Brief description:
- Change 1
- Change 2
```

### Voice & Style

Documents in this project use different voice registers:

| Document | Voice | Notes |
|----------|-------|-------|
| project-context.md | Personal | Lead dev perspective, colloquial ok |
| README | Personal | Project owner speaking to readers |
| origin-research.md | Personal | Framing is personal; quotes verbatim |
| philosophical-foundations.md | Project | Formal, rigorous, peer-reviewed |
| PRINCIPLES.md | Project | Derived by personas, ratified |
| Persona outputs (Vero reviews, etc.) | Persona | Defined by persona prompt |

**Personal voice**: See global CLAUDE.md `## Writing Voice` — direct, fragments ok, "I" heavy
**Project voice**: Formal, structured, rigorous — informed by personal preferences but more polished for external/community audience
**Persona voice**: Follow persona prompt specifications

### Reviews

When conducting persona reviews (Vero, peer reviews, etc.) on foundational documents:

1. **Always persist reviews** to `foundations/reviews/` with naming convention: `[reviewer]-review-[document]-v[X.Y.Z].md`
2. **Never run reviews inline-only** — chat history is not an audit trail
3. **Update the reviewed document's version** if changes are made during review
4. **Reference the review** in the document's Usage Note or Version History

Reviews are auditable artifacts, not just conversations.

### Fresh Derivation Process (Phase 2)

1. Each persona derives principles from their lens
2. Adversary challenges each derivation
3. Synthesize into unified PRINCIPLES.md
4. Incorporate operational elements (decision procedures, boundaries, audit rubric)

**Critical**: No access to existing guiding-principles docs during derivation.

---

## Lead Developer Context (Summary)

Full details in `foundations/project-context.md`.

**Constraints**:
- Variable time: 30+ hrs/wk (current) → periods of 0-20 hrs → ongoing 5-10 hrs
- Basic coding requiring AI assistance
- LLM-agnostic + open source: hard requirements

**Known Biases** (watch for):
- Overplanning — "What would we learn by building instead?"
- Idealism over pragmatism — GTM instincts are also correct
- Learning-as-excuse — Name this conflict when it's happening

**Success Criteria**: Daily dogfooding. No deadline.

---

## Known Limitations & Tradeoffs

| Limitation | Tradeoff | Mitigation |
|------------|----------|------------|
| **Claude-primary development** | LLM-agnosticism claimed but ~90% built with Claude | Solas validated across Claude/GPT/Gemini; test on other LLMs post-ship |
| **Solo dev constraints** | Limited review bandwidth | Persona ensemble simulates team review; external feedback post-ship |
| **Validation depth** | Full multi-LLM validation skipped | Meta-personas fully validated; principles battle-tested through use |

**Philosophy**: Ship with known limitations documented > wait for theoretical completeness.

---

## Session Continuity

1. Read this file (CONTEXT.md) first for orientation
2. Check [GitHub Issues](https://github.com/sympal-ai/sympal/issues) for current task status
3. Check `foundations/philosophical-foundations.md` for philosophical foundation
4. Check `foundations/project-context.md` for developer constraints/motivations
5. Check `PRINCIPLES.md` for binding principles (ratified)

---

## Git Workflow

**Repo**: https://github.com/sympal-ai/sympal

### Automatic Commits

Keep version control in sync automatically:

- **Commit after completing a task** — Don't batch; commit when done
- **Commit after significant file changes** — New files, major edits, renames
- **Commit before ending session** — Never leave uncommitted work

### Commit Style

```
[area]: Brief description

Optional body for context.

Co-Authored-By: Claude Opus 4.5 <noreply@anthropic.com>
```

Areas: `foundations`, `prompts`, `docs`, `config`, `meta`

Examples:
- `foundations: Add Vero final review of philosophical-foundations`
- `prompts: Create Vale persona for team`
- `meta: Update jobs-to-be-done status`

### When NOT to Commit

- Mid-task with broken/incomplete state
- User explicitly says to hold off
- Exploratory changes user hasn't approved

### Push Frequency

Push after each commit unless batching related changes. Don't let local get ahead of remote.

---

*Last updated: 2026-01-16*
