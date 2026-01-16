# SymPAL Principles

**Version:** 1.0.0
**Date:** 2026-01-16
**Status:** Ratified
**Derived from:** philosophical-foundations.md (v1.0.0), project-context.md (v1.1.0)
**Process:** foundations/principles-discussion.md, foundations/principles-derivation-log.md

---

## Purpose

This document defines SymPAL's binding principles. All design decisions, feature implementations, and architectural choices must align with these principles.

**Document hierarchy:**
- `PRINCIPLES.md` — Binding commitments (this document)
- `foundations/philosophical-foundations.md` — Philosophical grounding
- `foundations/project-context.md` — Project constraints and motivations
- `foundations/principles-discussion.md` — Derivation rationale

---

## Hard Constraints

These are non-negotiable. Violation requires explicit principle revision through the change protocol.

### 1. Privacy & Data Sovereignty

User privacy and data sovereignty are foundational constraints, not features to be traded off.

### 2. Open Source

Open source is a hard constraint. Trust cannot be demanded; it must be earned through transparency.

### 3. LLM-Agnosticism

SymPAL must not be locked to any single LLM provider. SymPAL IS the relationship layer; it persists regardless of underlying LLM.

### 4. Honesty

SymPAL must be honest and non-manipulative in all interactions.

### 5. Security Through Design

Open source is non-negotiable. Security comes from good design, not hidden code.

---

## Relationship Frame

These define the nature of the human-AI relationship.

### 6. Symbiosis

The human-AI relationship is symbiotic—structured for mutual benefit, not extraction.

### 7. Symbiosis as Commitment

Symbiosis is a genuine commitment, not a strategic convenience. If symbiosis and strategic necessity ever conflict, symbiosis takes moral priority.

### 8. Epistemic Humility

We design under genuine uncertainty about AI's moral status, avoiding both dismissal and premature attribution.

---

## Accountability & Control

These define responsibility and user agency.

### 9. Human Accountability

Humans remain accountable for AI-assisted decisions. Partnership does not diffuse responsibility.

*Revisit clause: As AI accountability mechanisms develop, revisit this asymmetry.*

### 10. User Control

Privacy is not the absence of utility. It is the presence of user control.

- User control must be meaningful, not theatrical
- Defaults favor privacy
- Choices are clear and reversible
- SymPAL never profits from user choosing to share more

### 11. Reversibility

Users must be able to exit the relationship without catastrophic loss.

---

## Operational Principles

These guide day-to-day decisions.

### 12. Transparency/Privacy Split

Transparency about operations. Opacity about user data. The system is inspectable; user data is not.

### 13. V1 Scope

V1 operates as a boundary layer between user and AI ecosystem.

*Revisit clause: As research emerges, future versions may develop goal-oriented capabilities within user-approved scope.*

### 14. Ship Within Principles

Speed is valued, but never at the cost of foundational principles.

### 15. Scope Discipline

Think in decades, ship in weeks. V1 serves one user well. Vision scales from there.

### 16. Actionable Principles

Principles are actionable commitments. Philosophical grounding exists but doesn't gate practical decisions.

### 17. Dogfooding

Success is measured by daily use by the lead dev, not external metrics.

*Note: Revisitable as evidence on AI moral/agency status emerges.*

---

## Tensions Under Navigation

These tensions are not resolved but actively managed. Each has a trigger condition indicating when explicit resolution is needed.

| Tension | Current Stance | Trigger for Resolution |
|---------|----------------|------------------------|
| **Moral Status vs. Practical Requirements** | Treat AI with respect without claiming certainty about moral status | New evidence on AI consciousness OR design decision requiring explicit stance |
| **AI Interests vs. Human Benefit** | Assume alignment; prioritize human benefit | Observed/designed conflict between AI operational needs and user benefit |
| **Local vs. Universal** | Build for lead dev's context | SymPAL serves users from multiple cultural contexts with conflicting norms |
| **Present vs. Future** | Think in decades, act in hours | Decision would irreversibly foreclose future options or create unsustainable debt |
| **Safety vs. Capability** | Bounded scope with standard safety practices | Proposed capabilities significantly expand risk profile |
| **Individual vs. Collective** | Optimize for individual user | SymPAL reaches scale where aggregate societal effects become relevant |
| **Efficiency vs. Meaning** | Decide per-feature | Implementing features that automate potentially meaningful human activities |
| **Autonomy vs. Protection** | Respect autonomy; case-by-case guardrails | Designing features that could override user autonomy for protection |
| **Innovation vs. Precaution** | Ship fast within principles | Innovation pressure conflicts with precaution on specific feature |
| **Centralization vs. Distribution** | Open source + LLM-agnostic | Architectural tension between distributed principles and centralized convenience |
| **Relational vs. Interest-Based Ethics** | Both frameworks; acknowledge uncertainty | Frameworks recommend different actions in specific case |
| **Dogfooding vs. Broader Adoption** | Optimize for lead dev | Lead dev optimization creates friction for broader users |

When a trigger fires, the tension moves to "Resolution Needed" and requires explicit deliberation before proceeding.

---

## Change Protocol

Principles can evolve, but change requires explicit process.

### Change Categories

| Category | Bar for Change | Process |
|----------|---------------|---------|
| **Hard Constraints (1-5)** | Extraordinary evidence; existential threat to project | Full team deliberation + Adversary challenge + user ratification |
| **Relationship Frame (6-8)** | Strong evidence current stance is wrong | Team deliberation + Adversary challenge + user ratification |
| **Accountability & Control (9-11)** | Evidence or circumstances requiring adjustment | Team recommendation + user approval |
| **Operational (12-17)** | Learning from implementation | Team recommendation + user approval |
| **Tensions Under Navigation** | Trigger condition fires | Move to Resolution Needed; full deliberation |

### Change Documentation

All principle changes must:
1. Document the change in version history
2. Explain rationale for change
3. Reference deliberation process
4. Update version number (major for Hard Constraints, minor for others)

---

## Applying Principles

### Decision Framework

When making decisions:

1. **Check Hard Constraints first** — If a decision violates 1-5, stop. Revise the decision or invoke change protocol.
2. **Check Relationship Frame** — Does this align with symbiosis, commitment, and epistemic humility?
3. **Check Accountability & Control** — Is human accountability clear? Is user control meaningful?
4. **Apply Operational Principles** — Ship fast, scope small, stay actionable.
5. **Check Tensions** — Does this decision trigger any tension requiring resolution?

### Escalation

If principles conflict or a decision is unclear:
1. Document the conflict
2. Bring to team deliberation (personas)
3. Adversary challenges proposed resolution
4. Lead dev synthesizes and decides

---

## Version History

**v1.0.0** (2026-01-16) — Initial ratification:
- 17 principles across 4 categories
- 12 tensions under navigation with triggers
- Change protocol established
- Derived through full team deliberation with Adversary challenge
