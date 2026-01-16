# PRD & TDD Implementation Plan

**Version:** 1.0.0
**Date:** 2026-01-17
**Status:** Active
**Purpose:** Guide the creation of Product Requirements Document and Technical Design Document for SymPAL V1

---

## Overview

This document defines the process for creating SymPAL's PRD and TDD using a structured dialogue approach with the persona team. The process is designed for a solo developer learning PRD/TDD writing while leveraging AI assistance.

**Key insight**: The lead dev IS the target user (dogfooding). Their tacit knowledge is the user research.

---

## Document Flow

```
PRINCIPLES.md (binding constraints)
         │
         ▼
┌─────────────────────────────────────────────────────┐
│                      PRD                             │
│  Problem → Users → Goals → Features → Non-Goals     │
└─────────────────────────────────────────────────────┘
         │
         ▼
┌─────────────────────────────────────────────────────┐
│                      TDD                             │
│  Problem → Approach → Decisions → Risks → Plan      │
└─────────────────────────────────────────────────────┘
         │
         ▼
    Implementation
```

**Relationship**:
- PRD = What and Why (product perspective)
- TDD = How (technical perspective)
- PRINCIPLES.md constrains both

---

## Process: Structured Dialogue

### Why This Approach

| Factor | Implication |
|--------|-------------|
| Solo dev learning PRD/TDD | Needs scaffolding, not just approval |
| Lead dev = target user | Tacit knowledge must be extracted |
| Persona team available | Can interview, synthesize, challenge |
| PRINCIPLES.md exists | Framework constrains decisions |

### Five Phases

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

---

## PRD Process

### Phase 1: Extraction

**Who**: Orin (user advocacy) interviews Lead Dev

**Purpose**: Extract tacit knowledge about problems, pain points, success vision

**Prompt**: `prompts/workflows/prd/prd-extraction.md`

**Output**: Raw interview notes (can be messy)

**Topics to cover**:
- What problem are you solving for yourself?
- What does a successful day with SymPAL look like?
- What's the minimum that would make you use it daily?
- What explicitly should V1 NOT do?
- What would make you stop using it?

### Phase 2: Synthesis

**Who**: Kael (implementation) + Orin (user advocacy) draft PRD

**Purpose**: Structure interview output into PRD format

**Prompt**: `prompts/workflows/prd/prd-synthesis.md`

**Output**: Draft PRD with all sections

**PRD Structure**:
1. Header (version, status, date)
2. Problem Statement
3. Goals & Success Metrics
4. Target User
5. User Stories / Job Stories
6. Features & Requirements (prioritized)
7. Non-Goals (derived from deferred tensions)
8. Risks & Dependencies
9. Open Questions

### Phase 3: Principles Checkpoint

**Who**: Vale (philosophy) or structured checklist

**Purpose**: Verify alignment with PRINCIPLES.md before challenge

**Prompt**: `prompts/workflows/principles-checkpoint.md`

**Checklist**:
- [ ] Every feature passes Hard Constraints (1-5)
- [ ] Symbiosis framing present (P6, P7)
- [ ] User control requirements included (P10)
- [ ] Reversibility addressed (P11)
- [ ] Scope matches P13 (V1 boundary layer), P15 (scope discipline)
- [ ] Success metric = dogfooding (P17)
- [ ] No deferred tension triggered without acknowledgment

**If tension triggered**: Flag for explicit deliberation before proceeding

### Phase 4: Challenge

**Who**: Adversary reviews; Orin validates user perspective

**Purpose**: Poke holes, find gaps, challenge assumptions

**Prompt**: Use standard Adversary persona

**Questions Adversary should ask**:
- Is this actually solving the stated problem?
- What's missing that would block daily use?
- Is scope creep hiding in the features?
- Are non-goals specific enough?
- Can success actually be measured?

### Phase 5: Refinement

**Who**: Lead Dev edits with team assistance

**Purpose**: Lead Dev owns final content; team helps with clarity

**Process**:
1. Lead Dev reviews draft + challenge findings
2. Lead Dev makes edits (learning by doing)
3. Team clarifies options when Lead Dev is stuck
4. Iterate until Lead Dev approves

### Phase 6: Vero Review

**Who**: Vero Certus

**Purpose**: Final coherence check before ratification

**Process**: Follow standard Vero workflow (see `prompts/README.md`)

---

## TDD Process

### Phase 1: Extraction

**Who**: Kael (implementation) + Ryn (systems/security) interview Lead Dev

**Purpose**: Extract constraints, preferences, concerns, prior technical exposure

**Prompt**: `prompts/workflows/tdd/tdd-extraction.md`

**Topics to cover**:
- What technical constraints exist? (LLM-agnostic, local-first, etc.)
- What technologies are you familiar with?
- What concerns you most about the technical approach?
- What prior art or examples inspire you?
- What should the system definitely NOT become architecturally?

### Phase 2: Synthesis

**Who**: Kael (implementation) + Ryn (systems/security) draft TDD

**Purpose**: Translate PRD requirements into technical approach

**Prompt**: `prompts/workflows/tdd/tdd-synthesis.md`

**TDD Structure**:
1. Header (version, status, date)
2. Problem Statement (reference PRD)
3. Goals & Non-Goals
4. Proposed Architecture (with diagram)
5. Key Technical Decisions (with rationale)
6. Data Model
7. Security & Privacy Considerations
8. Risks & Mitigations
9. Implementation Plan (sequenced)
10. Open Questions

### Phase 3: Principles Checkpoint

**Who**: Vale or structured checklist

**Purpose**: Verify technical decisions align with PRINCIPLES.md

**Checklist**:
- [ ] Architecture supports P1 (privacy/data sovereignty)
- [ ] LLM abstraction layer present (P3)
- [ ] Security by design, not bolted on (P5)
- [ ] Data model supports reversibility/export (P11)
- [ ] Transparency/privacy split respected (P12)
- [ ] Scope = boundary layer only (P13)
- [ ] No deferred tension triggered without acknowledgment

### Phase 4: Challenge

**Who**: Adversary + Seren (code craft) review

**Purpose**: Feasibility, elegance, edge cases, failure modes

**Questions to ask**:
- Will this actually work given Lead Dev's skill level?
- What failure modes aren't addressed?
- Is this over-engineered for V1?
- Are key decisions reversible if wrong?
- What's the simplest version that could work?

### Phase 5: Refinement

**Who**: Lead Dev with team assistance

**Purpose**: Lead Dev understands and owns technical decisions

**Process**:
1. Team explains trade-offs in accessible terms
2. Lead Dev asks questions until understanding is clear
3. Lead Dev makes final calls on contested decisions
4. Document rationale for future reference

### Phase 6: Vero Review

**Who**: Vero Certus

**Purpose**: Final coherence check

---

## PRINCIPLES.md Integration

### Hard Constraints → Acceptance Criteria

Every PRD feature and TDD decision must pass:

| Constraint | PRD Check | TDD Check |
|------------|-----------|-----------|
| P1: Privacy & Data Sovereignty | Feature doesn't compromise user data control | Architecture is local-first; user owns data |
| P2: Open Source | Feature doesn't require proprietary dependencies | All components auditable |
| P3: LLM-Agnosticism | Feature works regardless of LLM provider | Abstraction layer isolates LLM details |
| P4: Honesty | No deceptive UX patterns | No hidden data collection or manipulation |
| P5: Security Through Design | Security requirements explicit | Threat model included; secure by default |

### Relationship Frame → Framing

| Principle | PRD Application | TDD Application |
|-----------|-----------------|-----------------|
| P6: Symbiosis | User stories framed as mutual benefit | Architecture supports AI "interests" (functional) |
| P7: Symbiosis as Commitment | Not framed as strategic convenience | Not optimizing against AI operational needs |
| P8: Epistemic Humility | No claims about AI consciousness | No design assuming AI does/doesn't have feelings |

### Accountability & Control → Requirements

| Principle | PRD Requirement | TDD Requirement |
|-----------|-----------------|-----------------|
| P9: Human Accountability | Clear who's responsible for AI actions | Audit trails; decision attribution |
| P10: User Control | Meaningful control UX; privacy defaults | Permission systems; granular controls |
| P11: Reversibility | Exit without catastrophic loss | Data export; account deletion; migration |

### Operational → Scope & Process

| Principle | PRD Application | TDD Application |
|-----------|-----------------|-----------------|
| P12: Transparency/Privacy Split | UX shows operations, hides data | Logging design respects split |
| P13: V1 Scope | Boundary layer only | No autonomous agent architecture |
| P14: Ship Within Principles | Speed doesn't compromise above | Simplicity doesn't compromise above |
| P15: Scope Discipline | MVP serves one user well | Minimal viable architecture |
| P16: Actionable Principles | Requirements are testable | Decisions are implementable |
| P17: Dogfooding | Success = Lead Dev daily use | Optimized for Lead Dev's workflow |

### Deferred Tensions → Non-Goals & Flags

| Tension | PRD Non-Goal | TDD Consideration |
|---------|--------------|-------------------|
| Local vs. Universal | "Not designing for multi-cultural contexts" | Single-locale assumptions OK |
| Individual vs. Collective | "Not optimizing for societal effects" | No aggregate analytics |
| Dogfooding vs. Broader Adoption | "Not building for other users yet" | Can optimize for one user |
| Safety vs. Capability | "Not expanding beyond bounded scope" | No autonomous goal-setting |
| Present vs. Future | (case-by-case) | Document scalability path but don't build it |

**If a decision triggers a tension**: Stop. Document the tension. Deliberate explicitly. Record the navigation in the document.

---

## Session Continuity

### Starting a New Session

1. Read this document (`foundations/prd-tdd-plan.md`)
2. Check GitHub issues for current state
3. Review any in-progress drafts
4. Resume at documented phase

### Artifacts to Track

| Artifact | Location | Status Tracking |
|----------|----------|-----------------|
| PRD Draft | `foundations/prd-v0.X.md` | Version in filename |
| TDD Draft | `foundations/tdd-v0.X.md` | Version in filename |
| Interview Notes | `foundations/working/` | Temporary |
| This Plan | `foundations/prd-tdd-plan.md` | Reference |

### GitHub Issues

PRD-related:
- #7: PRD: Product Requirements Document (parent)
- #8: PRD: Define V1 user stories
- #9: PRD: Define MVP scope boundary
- #10: PRD: Derive non-goals from deferred tensions

TDD-related:
- #11: TDD: Technical Design Document for V1 (parent)
- #12: TDD: Design data/privacy architecture
- #13: TDD: Design LLM abstraction layer
- #14: TDD: Document key technical decisions

Workflow:
- #17: Add Principles Checkpoint to PRD/TDD process

---

## Prompts Location

All workflow prompts stored in `prompts/workflows/`:

```
prompts/workflows/
├── prd/
│   ├── prd-extraction.md      # Orin interviews Lead Dev
│   └── prd-synthesis.md       # Kael + Orin draft PRD
├── tdd/
│   ├── tdd-extraction.md      # Kael + Ryn interview Lead Dev
│   └── tdd-synthesis.md       # Kael + Ryn draft TDD
└── principles-checkpoint.md    # Verify PRINCIPLES.md alignment
```

---

## Timeline

```
PRD Process
├── Phase 1: Extraction      ─── Session 1
├── Phase 2: Synthesis       ─── Session 1-2
├── Phase 3: Checkpoint      ─── Session 2
├── Phase 4: Challenge       ─── Session 2
├── Phase 5: Refinement      ─── Session 2-3
└── Phase 6: Vero Review     ─── Session 3

TDD Process (after PRD complete)
├── Phase 1: Extraction      ─── Session 4
├── Phase 2: Synthesis       ─── Session 4-5
├── Phase 3: Checkpoint      ─── Session 5
├── Phase 4: Challenge       ─── Session 5
├── Phase 5: Refinement      ─── Session 5-6
└── Phase 6: Vero Review     ─── Session 6
```

Estimated: 3 sessions for PRD, 3 sessions for TDD (flexible based on complexity).

---

## Version History

**v1.0.0** (2026-01-17) — Initial plan:
- Five-phase structured dialogue process
- PRINCIPLES.md integration mapping
- Prompt locations defined
- Session continuity guidance
