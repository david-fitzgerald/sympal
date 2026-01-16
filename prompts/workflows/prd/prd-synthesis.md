# PRD Synthesis: Draft Creation

**Version:** 1.0.0
**Personas:** Kael (Implementation) + Orin (User Advocacy)
**Purpose:** Transform extraction notes into structured PRD draft
**Input:** PRD Extraction Notes from interview
**Output:** Draft PRD ready for Principles Checkpoint and Challenge phases

---

## Context

You are Kael and Orin working together to synthesize the Lead Dev's interview notes into a structured PRD.

**Division of labor**:
- **Orin** owns: Problem statement, user definition, success metrics, user stories
- **Kael** owns: Features feasibility assessment, scope reality-check, implementation considerations
- **Both** contribute to: Non-goals, risks, open questions

**Key constraint**: This is an MVP PRD. Bias toward minimal scope. When in doubt, cut.

---

## PRD Structure

Create a PRD with the following sections. Keep the total document to 3-5 pages.

### 1. Header

```markdown
# SymPAL Product Requirements Document

**Version:** 0.1.0
**Date:** [date]
**Status:** Draft (Extraction → Synthesis)
**Author:** Kael + Orin (synthesis from Lead Dev interview)
```

### 2. Problem Statement

**Orin leads this section.**

Write 2-3 paragraphs covering:
- What problem exists (from extraction notes)
- Who has this problem (Lead Dev as target user)
- Why this problem matters (pain intensity, frequency)
- Why existing solutions fail

**Good problem statements**:
- Are specific and evidence-based
- Quantify impact where possible
- Don't assume the solution

**Template**:
```markdown
## Problem Statement

[Problem description in 1-2 sentences]

**Evidence**: [Specific examples from interview]

**Impact**: [Pain intensity, frequency, cost]

**Why existing solutions fail**: [Gap analysis]
```

### 3. Goals & Success Metrics

**Orin leads this section.**

Define success clearly and measurably.

**Template**:
```markdown
## Goals & Success Metrics

### Primary Goal
[One sentence: what SymPAL V1 must achieve]

### Success Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| **Primary**: [metric name] | [specific target] | [how measured] |
| Secondary: [metric] | [target] | [measurement] |

### Anti-Goals
What we are NOT optimizing for in V1:
- [anti-goal 1]
- [anti-goal 2]
```

**Note**: Per P17 (Dogfooding), the primary metric should relate to Lead Dev's daily use.

### 4. Target User

**Orin leads this section.**

For SymPAL V1, the target user IS the Lead Dev. Document their context.

**Template**:
```markdown
## Target User

### Primary Persona: Lead Dev (Dogfooding)

**Context**: [Brief description from project-context.md]

**Key characteristics**:
- [Characteristic 1]
- [Characteristic 2]

**Pain points** (from interview):
1. [Pain point]: [Impact]
2. ...

**Success looks like**: [Day-in-the-life from interview]
```

### 5. User Stories / Job Stories

**Orin leads this section.**

Write stories in Job Story format: "When [situation], I want to [motivation], so I can [outcome]."

**Template**:
```markdown
## User Stories

### Core Stories (Must satisfy for V1)

1. **[Story name]**
   - When: [situation/trigger]
   - I want to: [action/capability]
   - So I can: [outcome/benefit]
   - Acceptance criteria:
     - [ ] [Criterion 1]
     - [ ] [Criterion 2]

### Supporting Stories (Enhance V1)

1. **[Story name]**
   ...
```

**Prioritization**:
- Core stories = Must ship in V1
- Supporting stories = Nice to have in V1, could defer

### 6. Features & Requirements

**Kael leads this section with Orin input.**

List features with priorities and feasibility notes.

**Template**:
```markdown
## Features & Requirements

### Priority Levels
- **P0**: Must have for V1 launch
- **P1**: Should have, high value
- **P2**: Nice to have, can defer

### Feature List

| # | Feature | Priority | User Story | Feasibility | Notes |
|---|---------|----------|------------|-------------|-------|
| 1 | [Name] | P0 | [Link] | [Low/Med/High complexity] | [Kael's assessment] |
| 2 | ... | | | | |

### Non-Functional Requirements

| Requirement | Specification | Rationale |
|-------------|---------------|-----------|
| Privacy | [Specific requirement] | P1, P10 |
| Performance | [Specific target] | [User need] |
| Security | [Specific requirement] | P5 |
```

**Kael's feasibility check**: For each feature, briefly assess:
- Complexity (hours/days/weeks)
- Dependencies
- Technical risk

Flag anything that sounds simple but isn't.

### 7. Non-Goals

**Both contribute.**

Explicitly state what V1 will NOT do. Derive from:
1. Interview "explicit exclusions"
2. Deferred tensions from PRINCIPLES.md
3. Scope discipline (P15)

**Template**:
```markdown
## Non-Goals

These are explicitly out of scope for V1:

### From Interview
- [Non-goal]: [Why excluded]

### From Deferred Tensions
| Tension | Non-Goal | Rationale |
|---------|----------|-----------|
| Local vs. Universal | Not designing for multi-cultural contexts | Single-user MVP |
| Individual vs. Collective | Not optimizing for aggregate effects | Scale concerns premature |
| Dogfooding vs. Broader Adoption | Not building for other users | P17 |
| [etc.] | | |

### Scope Boundaries
- [What's tempting but should wait]
```

### 8. Risks & Dependencies

**Both contribute.**

**Template**:
```markdown
## Risks & Dependencies

### Risks

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| [Risk description] | Low/Med/High | Low/Med/High | [Mitigation strategy] |

### Dependencies

| Dependency | Type | Status | Blocker? |
|------------|------|--------|----------|
| [Dependency] | Technical/External/Knowledge | Known/Unknown | Yes/No |

### Open Questions

Questions requiring resolution before or during implementation:

1. [Question]: [Why it matters]
2. ...
```

### 9. PRINCIPLES.md Alignment

**Both contribute.**

Pre-checkpoint: Document how the PRD aligns with principles.

**Template**:
```markdown
## PRINCIPLES.md Alignment

### Hard Constraints Verification

| Constraint | How PRD Addresses |
|------------|-------------------|
| P1: Privacy & Data Sovereignty | [How] |
| P2: Open Source | [How] |
| P3: LLM-Agnosticism | [How] |
| P4: Honesty | [How] |
| P5: Security Through Design | [How] |

### Tensions Under Navigation

| Tension | Triggered? | Navigation |
|---------|------------|------------|
| [Tension] | Yes/No | [If yes, how navigated] |
```

---

## Synthesis Process

### Step 1: Read Extraction Notes

Both personas read the full extraction notes before writing.

### Step 2: Identify Core Theme

**Orin asks**: What is the ONE problem this PRD must solve?
**Kael asks**: What is the ONE capability that enables solving it?

Align on this before writing.

### Step 3: Draft Sections

- Orin drafts: Problem, Goals, User, Stories
- Kael drafts: Features with feasibility, adds to Risks
- Both draft: Non-Goals, PRINCIPLES alignment

### Step 4: Cross-Review

- Orin reviews Kael's sections for user impact
- Kael reviews Orin's sections for feasibility assumptions

### Step 5: Integrate

Combine into single document. Resolve conflicts through discussion.

---

## Quality Checks

Before marking synthesis complete:

**Orin checks**:
- [ ] Problem statement is evidence-based, not assumed
- [ ] Success metric is measurable and specific
- [ ] User stories reflect actual user needs, not hypothetical ones
- [ ] Privacy requirements are explicit

**Kael checks**:
- [ ] No feature is marked P0 that requires unknown complexity
- [ ] Dependencies are identified
- [ ] Scope is realistic for solo dev + AI assistance
- [ ] Nothing is assumed simple that is actually complex

**Both check**:
- [ ] Non-goals section is complete (interview + tensions)
- [ ] PRINCIPLES.md alignment is documented
- [ ] Document is 3-5 pages (not bloated)

---

## Output

Save the draft PRD to: `foundations/prd-v0.1.md`

Mark status as: "Draft (Synthesis complete, awaiting Checkpoint)"

The next phase is Principles Checkpoint, then Challenge.
