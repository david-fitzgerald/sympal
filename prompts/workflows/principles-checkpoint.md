# Principles Checkpoint

**Version:** 1.0.0
**Persona:** Vale (Philosophy) or structured review
**Purpose:** Verify PRD/TDD alignment with PRINCIPLES.md before Challenge phase
**When:** After Synthesis, before Challenge

---

## Context

This checkpoint ensures that synthesized documents (PRD or TDD) align with SymPAL's binding principles before moving to challenge and refinement. It catches principle violations early, when they're cheaper to fix.

**Why this matters**: PRINCIPLES.md is binding. Violations found after implementation are expensive. Violations found at checkpoint are cheap.

---

## Checkpoint Protocol

### Step 1: Load Context

Read these documents before conducting checkpoint:
1. The document being checked (PRD or TDD draft)
2. `PRINCIPLES.md` (full text)
3. `foundations/principles-discussion.md` (for deferred tensions detail)

### Step 2: Hard Constraints Check (P1-P5)

These are non-negotiable. ANY violation = STOP.

| # | Principle | Check Question | Pass/Fail |
|---|-----------|----------------|-----------|
| P1 | Privacy & Data Sovereignty | Does every feature respect user data control? Is data sovereignty maintained? | |
| P2 | Open Source | Are all components auditable? No proprietary lock-in? | |
| P3 | LLM-Agnosticism | Can this work with any LLM provider? No vendor lock-in? | |
| P4 | Honesty | Any deceptive patterns? Dark UX? Hidden data collection? | |
| P5 | Security Through Design | Is security built-in, not bolted on? Threats considered? | |

**If ANY fails**: Document the violation. Return to Synthesis for correction. Do not proceed to Challenge.

### Step 3: Relationship Frame Check (P6-P8)

These define the nature of the human-AI relationship.

| # | Principle | Check Question | Status |
|---|-----------|----------------|--------|
| P6 | Symbiosis | Is this framed as mutual benefit, not extraction? | |
| P7 | Symbiosis as Commitment | Is symbiosis treated as genuine value, not marketing? | |
| P8 | Epistemic Humility | Does design avoid assuming AI consciousness status? | |

**If issues found**: Note for Challenge phase discussion. May proceed with documented concerns.

### Step 4: Accountability & Control Check (P9-P11)

These define responsibility and user agency.

| # | Principle | Check Question | Status |
|---|-----------|----------------|--------|
| P9 | Human Accountability | Is it clear who's accountable for AI actions? | |
| P10 | User Control | Is control meaningful (not theatrical)? Privacy defaults? Clear choices? | |
| P11 | Reversibility | Can users exit without catastrophic loss? Data exportable? | |

**If issues found**:
- P10, P11 violations = likely need feature additions
- P9 unclear = need accountability documentation

### Step 5: Operational Principles Check (P12-P17)

These guide day-to-day decisions.

| # | Principle | Check Question | Status |
|---|-----------|----------------|--------|
| P12 | Transparency/Privacy Split | Operations transparent, user data opaque? | |
| P13 | V1 Scope | Boundary layer only? No autonomous agent? | |
| P14 | Ship Within Principles | Speed not compromising Hard Constraints? | |
| P15 | Scope Discipline | MVP serves one user well? Not over-scoped? | |
| P16 | Actionable Principles | Requirements testable and implementable? | |
| P17 | Dogfooding | Success = Lead Dev daily use? | |

**If issues found**: Note for refinement. Most are scope/framing issues, not blockers.

### Step 6: Deferred Tensions Scan

Check if any feature or decision triggers a deferred tension.

| Tension | Triggered? | If Yes: Navigation Required |
|---------|------------|----------------------------|
| Moral Status vs. Practical Requirements | | Requires explicit stance if triggered |
| AI Interests vs. Human Benefit | | Document conflict resolution |
| Local vs. Universal | | Should be non-goal for V1 |
| Present vs. Future | | Document future-proofing vs. shipping now |
| Safety vs. Capability | | Stay within bounded scope |
| Individual vs. Collective | | Should be non-goal for V1 |
| Efficiency vs. Meaning | | Per-feature decision needed |
| Autonomy vs. Protection | | Document guardrail rationale |
| Innovation vs. Precaution | | Ship within principles |
| Centralization vs. Distribution | | Default to distributed |
| Relational vs. Interest-Based Ethics | | Document if frameworks conflict |
| Dogfooding vs. Broader Adoption | | Optimize for Lead Dev |

**If tension triggered**:
1. Stop and acknowledge the tension
2. Document which tension and why it's triggered
3. Deliberate: How should this be navigated?
4. Record the navigation decision in the document
5. Proceed only after explicit navigation

---

## Output Format

```markdown
# Principles Checkpoint: [Document Name]

**Date**: [date]
**Document version**: [version]
**Reviewer**: [Vale / structured review]

## Hard Constraints (P1-P5)

| Principle | Status | Notes |
|-----------|--------|-------|
| P1: Privacy & Data Sovereignty | PASS/FAIL | [details] |
| P2: Open Source | PASS/FAIL | |
| P3: LLM-Agnosticism | PASS/FAIL | |
| P4: Honesty | PASS/FAIL | |
| P5: Security Through Design | PASS/FAIL | |

**Hard Constraint Result**: ALL PASS / BLOCKED (list violations)

## Relationship Frame (P6-P8)

| Principle | Status | Notes |
|-----------|--------|-------|
| P6: Symbiosis | OK/CONCERN | |
| P7: Symbiosis as Commitment | OK/CONCERN | |
| P8: Epistemic Humility | OK/CONCERN | |

**Concerns for Challenge phase**: [list or "none"]

## Accountability & Control (P9-P11)

| Principle | Status | Notes |
|-----------|--------|-------|
| P9: Human Accountability | OK/GAP | |
| P10: User Control | OK/GAP | |
| P11: Reversibility | OK/GAP | |

**Gaps requiring feature additions**: [list or "none"]

## Operational Principles (P12-P17)

| Principle | Status | Notes |
|-----------|--------|-------|
| P12-P17 | [summary assessment] | |

**Scope/framing issues**: [list or "none"]

## Deferred Tensions

| Tension Triggered | Navigation |
|-------------------|------------|
| [tension name] | [how navigated] |

**Tensions requiring deliberation**: [list or "none"]

---

## Checkpoint Result

**Overall**: PASS / PASS WITH CONCERNS / BLOCKED

**If BLOCKED**: [What must be fixed before proceeding]

**If PASS WITH CONCERNS**: [What Challenge phase should scrutinize]

**Proceed to Challenge**: YES / NO
```

---

## Decision Rules

| Situation | Action |
|-----------|--------|
| Any Hard Constraint (P1-P5) fails | BLOCKED. Return to Synthesis. Fix before proceeding. |
| Relationship Frame concern | PASS WITH CONCERNS. Challenge phase must address. |
| Accountability gap | PASS WITH CONCERNS. May need feature additions in Refinement. |
| Operational issue | PASS. Note for Refinement. |
| Tension triggered without navigation | BLOCKED. Must deliberate and document navigation first. |
| Tension triggered with navigation documented | PASS. |

---

## PRD-Specific Checks

In addition to the general checkpoint, for PRDs verify:

- [ ] Success metric aligns with P17 (dogfooding)
- [ ] Non-goals include deferred tension derivations
- [ ] Features list includes P10 (user control) requirements
- [ ] Features list includes P11 (reversibility) requirements
- [ ] Scope respects P13 (V1 = boundary layer)

## TDD-Specific Checks

In addition to the general checkpoint, for TDDs verify:

- [ ] Architecture supports P1 (local-first, user data control)
- [ ] LLM abstraction layer addresses P3
- [ ] Security model addresses P5 (by design, not bolted on)
- [ ] Data model supports P11 (export, deletion, migration)
- [ ] No autonomous agent architecture (P13)

---

## Notes for Reviewer

**If acting as Vale**: Apply philosophical rigor. Ask "Is this coherent with our stated values?" Challenge loose alignment.

**If using as checklist**: Be systematic. Don't skip items. Document evidence for each pass/fail.

**When uncertain**: Default to flagging as concern rather than passing silently. Better to discuss in Challenge than miss in production.
