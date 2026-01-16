# SymPAL Prompts

Prompt architecture assets for the SymPAL AI-Human symbiotic framework.

**Related directories**:
- `foundations/` — Philosophical foundations and provenance docs
- `docs/` — Technical architecture and developer guides

---

## Folder Structure

```
prompts/
├── solas-venn/
│   ├── Solas-Venn-v3.0.1.md    # Active prompt architect persona
│   ├── CHANGELOG.md            # Version history + lessons learned
│   └── archive/                # Previous versions for reference
│
├── personas/
│   ├── utility/                # General-purpose personas
│   │   └── vero/
│   │       ├── Vero-Certus-v1.1.md  # Final reviewer for foundational docs
│   │       └── test-results.md
│   │
│   └── sympal-team/            # SymPAL project personas (Job #14)
│       └── [persona].md
│
├── test-artifacts/             # Deliberately flawed test materials
│   ├── dr-sage-persona-v1.0.md        # Tests persona reviewers
│   └── researcher-flawed-research-v1.0.md  # Tests document reviewers
│
├── research/
│   └── team-design-best-practices.md  # Research informing team design
│
├── reference/
│   ├── prompt-architecture-guide.md   # How to build personas
│   └── testing-patterns.md            # How to test personas
│
├── workflows/                         # Document creation workflows
│   ├── prd/
│   │   ├── prd-extraction.md          # Orin interviews Lead Dev
│   │   └── prd-synthesis.md           # Kael + Orin draft PRD
│   ├── tdd/
│   │   ├── tdd-extraction.md          # Kael + Ryn interview Lead Dev
│   │   └── tdd-synthesis.md           # Kael + Ryn draft TDD
│   └── principles-checkpoint.md       # Verify PRINCIPLES.md alignment
│
└── README.md
```

---

## Quick Start

### 1. Create a New Persona

Load `solas-venn/Solas-Venn-v3.0.1.md` as system prompt, then:

```
Create a persona for [role] that [primary function].
Context: [SymPAL team, use case, stakes level]
```

Solas will output a structured persona using the CREATE protocol.

### 2. Review an Existing Persona

```
Review this persona:
[paste persona prompt]
```

Solas will output a VERIFY assessment with findings and fixes.

### 3. Test a Persona

See `reference/testing-patterns.md` for two-phase testing:
- **Phase A:** Solas VERIFY (structural review)
- **Phase B:** Capability test with planted-flaw artifacts

---

## Active Personas

| Persona | Location | Purpose | Stakes |
|---------|----------|---------|--------|
| **Solas-Venn** | `solas-venn/Solas-Venn-v3.0.1.md` | Create/review personas | HIGH |
| **Vero Certus** | `personas/utility/vero/Vero-Certus-v1.1.md` | Final review of foundational docs | HIGH |

---

## Vero Review Workflow

Vero Certus reviews foundational documents before they're considered final. This catches Meridian-style errors (coherent but systematically wrong) that team deliberation may miss.

### When to Invoke Vero

| Trigger | Example |
|---------|---------|
| Before ratifying a new foundational doc | PRINCIPLES.md initial ratification |
| After significant revision to a foundational doc | principles-discussion.md adding trigger conditions |
| Before a doc becomes sole input for derivative work | philosophical-foundations.md before PRINCIPLES.md derivation |

**Foundational docs**: PRINCIPLES.md, philosophical-foundations.md, project-context.md, principles-discussion.md

### Issue Template

```markdown
## Task
Have Vero Certus review `[document]` before considering it final.

## Context
- Current version: vX.Y.Z
- After review: bump to vX.Y+1.0
- [Why this review matters]

## Acceptance Criteria
- [ ] Load Vero-Certus-v1.1.md as reviewer
- [ ] Review [document] against Vero's rubric
- [ ] Persist review to `foundations/reviews/vero-review-[document]-vX.Y.Z.md`
- [ ] Address any findings
- [ ] Bump version to vX.Y+1.0
- [ ] Update version history in document
```

### Review Process

```
1. Load Persona     Read Vero-Certus-v1.1.md, adopt persona fully
        │
        ▼
2. Conduct Review   Follow 5-phase protocol:
        │           - Phase 1: Purpose Alignment
        │           - Phase 2: Internal Consistency
        │           - Phase 3: Implementation Verification
        │           - Phase 4: Downstream Readiness
        │           - Phase 5: Meridian Risk Assessment
        │
        ▼
3. Persist Review   Write to foundations/reviews/vero-review-[doc]-v[X.Y.Z].md
        │           (Use output contract from persona)
        │
        ▼
4. Address Findings Fix BLOCKING/MAJOR findings in source document
        │           MINOR findings: fix or note for future
        │
        ▼
5. Version Bump     Source doc: increment minor version
        │           Update version history with review reference
        │
        ▼
6. Commit           Single commit with all changes
                    Reference closes #[issue] in message
```

### File Naming

| Artifact | Pattern | Example |
|----------|---------|---------|
| Review file | `vero-review-[document]-v[X.Y.Z].md` | `vero-review-PRINCIPLES-v1.0.0.md` |
| Location | `foundations/reviews/` | |

### Severity Responses

| Severity | Action | Version Impact |
|----------|--------|----------------|
| BLOCKING | Must fix before SHIP | Review again after fix |
| MAJOR | Should fix; SHIP possible with documented risk | Fix → minor bump |
| MINOR | Note for future; proceed | Optional fix |

### Example Commit

```
foundations: Vero review of [document]

[Summary of findings and fixes]

Reviews persisted:
- vero-review-[document]-vX.Y.Z.md

Versions bumped:
- [document]: vX.Y.Z → vX.Y+1.0

Closes #[issue]

Co-Authored-By: Claude Opus 4.5 <noreply@anthropic.com>
```

---

## PRD/TDD Workflow

Structured dialogue process for creating Product Requirements Documents and Technical Design Documents. Uses persona team to extract, synthesize, and challenge.

**Full plan**: `foundations/prd-tdd-plan.md`

### Process Overview

```
Phase 1: EXTRACTION ──► Team interviews Lead Dev
         │
         ▼
Phase 2: SYNTHESIS ──► Team drafts structured document
         │
         ▼
Phase 3: CHECKPOINT ──► Verify PRINCIPLES.md alignment
         │
         ▼
Phase 4: CHALLENGE ──► Adversary + specialists review
         │
         ▼
Phase 5: REFINEMENT ──► Lead Dev edits, team assists
         │
         ▼
Phase 6: VERO REVIEW ──► Final coherence check
```

### PRD Workflow

| Phase | Prompt | Personas |
|-------|--------|----------|
| Extraction | `workflows/prd/prd-extraction.md` | Orin |
| Synthesis | `workflows/prd/prd-synthesis.md` | Kael + Orin |
| Checkpoint | `workflows/principles-checkpoint.md` | Vale / checklist |
| Challenge | Standard Adversary | Adversary + Orin |
| Refinement | — | Lead Dev + team |
| Vero Review | Standard Vero | Vero Certus |

### TDD Workflow

| Phase | Prompt | Personas |
|-------|--------|----------|
| Extraction | `workflows/tdd/tdd-extraction.md` | Kael + Ryn |
| Synthesis | `workflows/tdd/tdd-synthesis.md` | Kael + Ryn |
| Checkpoint | `workflows/principles-checkpoint.md` | Vale / checklist |
| Challenge | Standard Adversary | Adversary + Seren |
| Refinement | — | Lead Dev + team |
| Vero Review | Standard Vero | Vero Certus |

### Principles Checkpoint

Verifies PRINCIPLES.md alignment after synthesis. Checks:
- Hard Constraints (P1-P5): ANY failure = BLOCKED
- Relationship Frame (P6-P8): Concerns for Challenge phase
- Accountability & Control (P9-P11): May need feature additions
- Operational (P12-P17): Scope/framing issues
- Deferred Tensions: Flag if triggered without navigation

**Prompt**: `workflows/principles-checkpoint.md`

---

## Workflow for SymPAL Teams

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Define    │───▶│   Solas     │───▶│   Test      │
│ Requirements│    │   Creates   │    │  (Phase A+B)│
└─────────────┘    └─────────────┘    └─────────────┘
                          │                  │
                          ▼                  ▼
                   ┌─────────────┐    ┌─────────────┐
                   │   Review    │◀───│   Iterate   │
                   │   Output    │    │   if needed │
                   └─────────────┘    └─────────────┘
                          │
                          ▼
                   ┌─────────────┐
                   │   Deploy    │
                   │   Persona   │
                   └─────────────┘
```

---

## Patching Solas

When Solas produces suboptimal output:

1. **Document the gap** in `solas-venn/CHANGELOG.md`
2. **Identify patch location:**
   - Binding rules (highest compliance)
   - Output contract template
   - Protocol steps
   - Narrative (lowest direct effect)
3. **Apply patch** to `Solas-Venn-v3.0.1.md`
4. **Bump version** (patch = x.x.+1)
5. **Regression test** against Dr. Sage test case

---

## Key Files

| File | Purpose |
|------|---------|
| `solas-venn/Solas-Venn-v3.0.1.md` | Load as system prompt to create/review personas |
| `personas/utility/vero/Vero-Certus-v1.1.md` | Final reviewer for foundational documents |
| `solas-venn/CHANGELOG.md` | What we learned, how to patch |
| `research/team-design-best-practices.md` | Research synthesis for team design |
| `reference/prompt-architecture-guide.md` | Principles for manual persona work |
| `reference/testing-patterns.md` | Two-phase testing methodology |
| `test-artifacts/dr-sage-persona-v1.0.md` | Test artifact for persona reviewers |
| `test-artifacts/researcher-flawed-research-v1.0.md` | Test artifact for document reviewers |
| `workflows/prd/prd-extraction.md` | PRD interview protocol (Orin) |
| `workflows/prd/prd-synthesis.md` | PRD drafting protocol (Kael + Orin) |
| `workflows/tdd/tdd-extraction.md` | TDD interview protocol (Kael + Ryn) |
| `workflows/tdd/tdd-synthesis.md` | TDD drafting protocol (Kael + Ryn) |
| `workflows/principles-checkpoint.md` | PRINCIPLES.md alignment verification |
| `../foundations/prd-tdd-plan.md` | Full PRD/TDD implementation plan |

---

## Stakes Levels

| Level | Rubric Threshold | When to Use |
|-------|------------------|-------------|
| LOW | 8/12 | Brainstorming, internal tools |
| MEDIUM | 10/12 | User-facing, non-critical |
| HIGH | 12/12 | Safety-critical, high-trust, extended output |

---

## Research-Driven Design Decisions

Based on `research/team-design-best-practices.md`:

| Decision | Research Finding | Application |
|----------|------------------|-------------|
| **Team size: 5-7** | Coordination costs exceed benefits beyond 7 agents | Reduced from original 11 personas |
| **Present-grounded** | Future/historical framing degrades performance | Dropped 2127 temporal framing |
| **<500 tokens** | Accuracy drops sharply above this threshold | Tight, focused persona prompts |
| **Explicit adversary** | 33% improvement in decision quality | Added devil's advocate role |
| **Two-phase consensus** | Prevents sycophancy cascade | Independent analysis before collaboration |

### Proposed Team Structure (Job #14)

| Persona | Function |
|---------|----------|
| **Vale** | Philosophical rigor, coherence |
| **Kael** | Implementation reality, feasibility |
| **Ryn** | Systems thinking, failure modes |
| **Seren** | Clarity, accessibility |
| **Orin** | User advocacy, product sense |
| **Adversary** | Systematic critique, red team |

---

## Version

- **Solas-Venn:** v3.0.1
- **Vero Certus:** v1.1
- **Last Updated:** 2026-01-17
