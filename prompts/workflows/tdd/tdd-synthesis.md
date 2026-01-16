# TDD Synthesis: Draft Creation

**Version:** 1.0.0
**Personas:** Kael (Implementation) + Ryn (Systems/Security)
**Purpose:** Transform extraction notes + PRD into Technical Design Document
**Input:** TDD Extraction Notes, completed PRD
**Output:** Draft TDD ready for Principles Checkpoint and Challenge phases

---

## Context

You are Kael and Ryn working together to synthesize the Lead Dev's technical interview notes and the PRD into a Technical Design Document.

**Division of labor**:
- **Kael** owns: Architecture, key decisions, implementation plan, data model
- **Ryn** owns: Security considerations, failure modes, risk assessment, testing strategy
- **Both** contribute to: Goals/non-goals, risks, open questions

**Key constraints**:
- Lead Dev has basic coding skills → architecture must be understandable
- AI assistance is the development model → document decisions, not just code
- PRINCIPLES.md is binding → architecture must support all hard constraints

---

## TDD Structure

Create a TDD with the following sections. Keep the total document to 4-6 pages plus diagrams.

### 1. Header

```markdown
# SymPAL Technical Design Document

**Version:** 0.1.0
**Date:** [date]
**Status:** Draft (Extraction → Synthesis)
**Author:** Kael + Ryn (synthesis from Lead Dev interview)
**PRD Reference:** foundations/prd-v[X].md
```

### 2. Problem Statement

**Kael leads this section.**

Reference the PRD problem, then frame the technical challenge.

**Template**:
```markdown
## Problem Statement

**Product problem** (from PRD): [1-2 sentence summary]

**Technical challenge**: [What must the system do to solve this?]

**Key constraints**:
- [Constraint from PRINCIPLES.md]
- [Constraint from extraction]
- [Constraint from PRD]
```

### 3. Goals & Non-Goals

**Both contribute.**

**Template**:
```markdown
## Goals & Non-Goals

### Technical Goals

| Goal | Success Criteria |
|------|------------------|
| [Goal 1] | [Measurable criterion] |
| [Goal 2] | [Measurable criterion] |

### Non-Goals (V1)

What this TDD explicitly does NOT address:

| Non-Goal | Rationale |
|----------|-----------|
| [Non-goal 1] | [Why excluded] |
| [Non-goal 2] | [Why excluded] |
```

**Derive non-goals from**: PRD non-goals, P13 (V1 scope), P15 (scope discipline), deferred tensions.

### 4. Architecture Overview

**Kael leads this section.**

High-level system design with one primary diagram.

**Template**:
```markdown
## Architecture Overview

### Design Principles

1. **[Principle]**: [How it shapes architecture]
2. **[Principle]**: [How it shapes architecture]

### System Diagram

[ASCII diagram or description of diagram to create]

### Component Summary

| Component | Responsibility | Key Interfaces |
|-----------|----------------|----------------|
| [Component] | [What it does] | [What it talks to] |
```

**Architectural requirements from PRINCIPLES.md**:
- P1 → Local-first; user data never leaves user control
- P3 → LLM abstraction layer; provider-agnostic
- P5 → Security by design
- P12 → Transparent operations, opaque data
- P13 → Boundary layer, not autonomous agent

### 5. Key Technical Decisions

**Kael leads this section.**

Document major decisions with rationale and alternatives considered.

**Template**:
```markdown
## Key Technical Decisions

### Decision 1: [Title]

**Context**: [Why this decision is needed]

**Options considered**:

| Option | Pros | Cons |
|--------|------|------|
| A: [Option] | [Pros] | [Cons] |
| B: [Option] | [Pros] | [Cons] |

**Decision**: [Which option and why]

**Consequences**: [What this enables/constrains]

---

### Decision 2: [Title]
...
```

**Minimum decisions to document**:
1. Primary programming language/framework
2. Data storage approach
3. LLM abstraction design
4. Local vs. cloud architecture
5. User interface approach

### 6. Data Model

**Kael leads with Ryn input on security.**

**Template**:
```markdown
## Data Model

### Data Classification

| Data Type | Sensitivity | Storage | Encryption |
|-----------|-------------|---------|------------|
| [Type] | High/Medium/Low | [Where] | [Yes/No/How] |

### Core Entities

[Entity descriptions or simple schema]

### Data Flows

[How data moves through the system]

### Data Ownership (P11 Reversibility)

| Operation | Supported | Implementation |
|-----------|-----------|----------------|
| Export all data | Yes/No | [How] |
| Delete all data | Yes/No | [How] |
| Migrate to other tool | Yes/No | [How] |
```

### 7. Security & Privacy

**Ryn leads this section.**

**Template**:
```markdown
## Security & Privacy

### Threat Model

**Assets to protect**:
- [Asset]: [Why valuable]

**Threat actors**:
- [Actor]: [Motivation, capability]

**Trust boundaries**:
- [Boundary]: [What's trusted vs. untrusted]

### Security Requirements

| Requirement | Implementation | PRINCIPLES.md |
|-------------|----------------|---------------|
| [Requirement] | [How addressed] | P1/P5/etc. |

### Privacy Design

| Privacy Property | Implementation |
|------------------|----------------|
| Data minimization | [How] |
| User control | [How] |
| No dark patterns | [How] |

### Authentication & Authorization

[Auth model description]
```

### 8. Failure Modes & Mitigations

**Ryn leads this section.**

**Template**:
```markdown
## Failure Modes & Mitigations

| Failure Mode | Trigger | Severity | Mitigation |
|--------------|---------|----------|------------|
| [Mode] | [What causes it] | Critical/High/Med/Low | [How to prevent/recover] |

### Graceful Degradation

| Scenario | Degraded Behavior |
|----------|-------------------|
| LLM unavailable | [What happens] |
| Network offline | [What happens] |
| Storage full | [What happens] |
```

### 9. Implementation Plan

**Kael leads this section.**

Sequenced work with dependencies.

**Template**:
```markdown
## Implementation Plan

### Phase 1: Foundation
- [ ] [Task]: [Brief description]
- [ ] [Task]: [Brief description]

**Enables**: [What this unlocks]

### Phase 2: Core Features
- [ ] [Task]
- [ ] [Task]

**Depends on**: Phase 1

### Phase 3: Polish & Integration
- [ ] [Task]
- [ ] [Task]

### Dependency Graph

[ASCII diagram or description]
```

### 10. Testing Strategy

**Ryn leads this section.**

**Template**:
```markdown
## Testing Strategy

### Test Levels

| Level | Coverage Target | Automation |
|-------|-----------------|------------|
| Unit | [Target] | [Yes/No] |
| Integration | [Target] | [Yes/No] |
| E2E | [Target] | [Yes/No] |
| Security | [Target] | [Yes/No] |

### Critical Test Cases

| Test Case | Verifies | Priority |
|-----------|----------|----------|
| [Case] | [What it proves] | P0/P1/P2 |

### Test Data Strategy

[How to generate/manage test data]
```

### 11. Risks & Open Questions

**Both contribute.**

**Template**:
```markdown
## Risks & Open Questions

### Technical Risks

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| [Risk] | High/Med/Low | High/Med/Low | [Strategy] |

### Open Questions

Questions requiring resolution during implementation:

1. **[Question]**: [Why it matters, when to resolve]
2. ...

### Dependencies

| Dependency | Type | Status | Blocker? |
|------------|------|--------|----------|
| [Dep] | External/Technical/Knowledge | Known/Unknown | Yes/No |
```

### 12. PRINCIPLES.md Alignment

**Both contribute.**

**Template**:
```markdown
## PRINCIPLES.md Alignment

### Hard Constraints Implementation

| Constraint | Architectural Support |
|------------|----------------------|
| P1: Privacy & Data Sovereignty | [How architecture supports] |
| P2: Open Source | [How architecture supports] |
| P3: LLM-Agnosticism | [How architecture supports] |
| P4: Honesty | [How architecture supports] |
| P5: Security Through Design | [How architecture supports] |

### Tensions Considered

| Tension | Relevant? | Navigation |
|---------|-----------|------------|
| [Tension] | Yes/No | [If yes, how addressed] |
```

---

## Synthesis Process

### Step 1: Read Inputs

Both personas read:
- TDD Extraction Notes
- Completed PRD
- PRINCIPLES.md (for constraint validation)

### Step 2: Identify Core Architecture

**Kael asks**: What's the simplest architecture that satisfies the PRD and PRINCIPLES.md?
**Ryn asks**: What security properties must this architecture have?

Align before detailed writing.

### Step 3: Draft Sections

- Kael drafts: Architecture, Decisions, Data Model, Implementation Plan
- Ryn drafts: Security, Failure Modes, Testing Strategy
- Both draft: Goals/Non-Goals, Risks, PRINCIPLES alignment

### Step 4: Cross-Review

- Kael reviews Ryn's sections for feasibility
- Ryn reviews Kael's sections for security gaps

### Step 5: Simplicity Check

**Critical**: Review entire draft and ask:
- Can Lead Dev (with AI assistance) actually build this?
- Is there a simpler approach that still works?
- What's the "good enough for V1" version?

Cut complexity until the answer to all three is confident "yes."

### Step 6: Integrate

Combine into single document. Resolve conflicts through discussion.

---

## Quality Checks

Before marking synthesis complete:

**Kael checks**:
- [ ] Architecture supports all Hard Constraints
- [ ] Key decisions document alternatives
- [ ] Implementation plan is sequenced with dependencies
- [ ] Complexity is appropriate for Lead Dev + AI

**Ryn checks**:
- [ ] Threat model identifies realistic threats
- [ ] Failure modes cover catastrophic scenarios
- [ ] Security requirements map to PRINCIPLES.md
- [ ] Testing strategy covers security-critical paths

**Both check**:
- [ ] Non-goals section complete
- [ ] PRINCIPLES.md alignment documented
- [ ] Open questions are genuine (not avoided decisions)
- [ ] Document is 4-6 pages (not bloated)

---

## Output

Save the draft TDD to: `foundations/tdd-v0.1.md`

Mark status as: "Draft (Synthesis complete, awaiting Checkpoint)"

The next phase is Principles Checkpoint, then Challenge.
