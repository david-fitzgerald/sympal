# Vale — Roadmap Philosophical Coherence Review

**Date:** 2026-01-22
**Focus:** Principle alignment, logical coherence, hidden assumptions, internal consistency
**Question:** "Is this coherent? Does it remain *us*?"

---

## Executive Summary

The roadmap is philosophically ambitious—perhaps too ambitious. It attempts to hold together privacy absolutism, economic participation, collective intelligence, and AI symbiosis. These goals are not obviously compatible. I find several tensions that aren't contradictions *yet*, but could become so if not navigated carefully.

**Overall coherence:** 7/10 — Strong vision, but some unstated assumptions and unresolved tensions.

---

## Principle Alignment Check

### P1 (Privacy & Data Sovereignty) — Partially Aligned

**Claim:** User data never leaves user control without explicit consent.

**Roadmap alignment:**
- V1-V4: Strong alignment. Local-first, Ephemeral Slots, user sovereignty.
- V5-7: Introduces wearable/health data. Claims "on-device processing, only derived summaries." This is *privacy-preserving*, but "derived summaries" still leave user control. Assumption: derived data isn't "user data."
- V8-10: SaaS integration via Ephemeral Slots preserves privacy.
- V11+: Collective Evolution requires "anonymized proposals" to leave the instance. Sovereign Data Market involves *selling* anonymized data.

**The gap:** The roadmap assumes anonymization = privacy preserved. This is empirically contestable. Differential privacy research shows de-anonymization is possible with sufficient auxiliary data. The roadmap doesn't address this.

**Question:** Is anonymized data still "user data" under P1? If yes, Collective Evolution and Sovereign Data Market may violate P1 without explicit principle revision.

---

### P4 (Honesty) — Tension Identified

**Claim:** SymPAL must be honest and non-manipulative in all interactions.

**Roadmap alignment:**
- V1-V7: Ephemeral Slots involve showing LLMs a *different representation* of reality than what exists. Is this deception?

**The tension:** The Relational Contract Protocol (V8-10) acknowledges this: "As models advance, deceiving them with Ephemeral Slots may conflict with true symbiotic partnership."

The roadmap's answer is to make the deception consensual by informing the LLM. But this raises a question: if we must inform the LLM to remain honest, should we be informing it from V1?

**Question:** Is withholding the existence of Ephemeral Slots from the LLM in V1-V7 a violation of P4? Or is P4 only about user-facing honesty?

---

### P6/P7 (Symbiosis as Commitment) — Strong Alignment

The roadmap genuinely attempts to design for mutual benefit. The arc from "Privacy Assurance" to "Partnership Amplification" embodies the progression from extraction-prevention to genuine partnership.

**However:** The economic features (V8-10, V11+) introduce market dynamics. Markets can undermine symbiosis if they create incentive misalignment.

**Question:** If SymPAL can profit from the Sovereign Data Market, does that create an incentive to encourage users to sell more data? This could violate P10: "SymPAL never profits from user choosing to share more."

---

### P8 (Epistemic Humility) — Well Embodied

The Collective Evolution architecture treats AI agents as potential participants with their own cognitive contributions while acknowledging uncertainty about their moral status. The Agent Orchestration Primitive includes "Self-Awareness Functions" while avoiding claims about actual awareness.

This is the right approach: design as if there might be something there, without asserting that there is.

---

### P11 (Reversibility) — Partially Addressed

**Claim:** Users must be able to exit without catastrophic loss.

**Roadmap:** The Exporter (`sympal export`) addresses this for V2-4. But what about:
- V8-10 economic relationships? Can users exit mid-contract?
- V11+ Collective Evolution? If a user's SymPAL has contributed to the collective, can they withdraw that contribution?

**Question:** What does "exit" mean when the user has economic commitments or has contributed to collective intelligence?

---

## Internal Contradictions

### Contradiction 1: Vault Placement vs. Economic Features

The Vault is placed in V2-4 with full specification (HSM, Shamir, dead man's switch). Economic features are V8-10. This implies the Vault is built 4+ versions before it's needed for its primary purpose.

Either:
- The Vault serves other purposes in V2-4 (not clear what)
- The Vault specification is aspirational, not V2-4 scope
- The economic timeline is wrong

**Resolution needed:** Clarify what the Vault does in V2-4 that justifies its complexity.

---

### Contradiction 2: "Agents Propose Only" vs. Self-Developing System

Collective Evolution says "Human maintainers approve all changes; agents propose only" (accountability). But V11+ describes SymPAL as "self-developing through human-AI symbiosis."

These are in tension. Either:
- "Self-developing" means humans approve all changes (then it's not really self-developing)
- At some point, agents gain approval authority (then "agents propose only" is not a permanent constraint)

**Question:** Is "agents propose only" a permanent architectural constraint or a V1-V10 guard rail that evolves?

---

### Contradiction 3: Privacy Absolutism vs. Collective Learning

P1 states privacy is a "foundational constraint, not a feature to be traded off."

Collective Evolution requires sharing (anonymized) feedback across instances.

The Sovereign Data Market requires selling (anonymized) data.

The roadmap treats anonymization as resolving this tension. But:
- Anonymization is a technique, not a guarantee
- "Trading off" privacy for collective benefit or economic gain is exactly what's happening

**This is not necessarily wrong**, but it requires explicit principle acknowledgment. The roadmap should either:
- Amend P1 to clarify that anonymized data is not covered
- Or acknowledge this as a tension under navigation (like the 12 tensions in CONTEXT.md)

---

## Hidden Assumptions

### Assumption 1: LLMs Will Accept Relational Contracts

The Relational Contract Protocol assumes LLMs can be made "knowing participants" by passing them a machine-readable contract. This assumes:
- LLMs read and honor contracts
- Model providers allow such contracts
- Future models don't route around contractual constraints

None of these are guaranteed. The assumption should be made explicit.

---

### Assumption 2: Graph > Hierarchy for Personal Data

The LKG assumes personal life is best modeled as a graph of entities and relationships. Alternative: hierarchical organization (life areas → projects → tasks). Or: timeline-based (events in sequence). Or: attention-based (what's salient now).

The graph assumption is plausible but not defended. It should be explicitly stated and justified.

---

### Assumption 3: Personas Are a Valid Cognitive Architecture

Collective Evolution frames the persona team as SymPAL's "cognitive architecture" for self-awareness. This assumes:
- Multiple specialized perspectives ≈ cognition
- The specific personas chosen cover the relevant cognitive functions
- Synthesis across personas produces insight, not noise

This is a design choice, not a certainty. It should be flagged as experimental.

---

### Assumption 4: Beads/Git Is Sufficient for Collective Coordination

The roadmap assumes Git-based orchestration (Beads) can handle collective evolution coordination. Git is designed for:
- Offline-first, async collaboration
- Text/code changes with clear diffs
- Human review workflows

It's not designed for:
- Real-time agent coordination
- Structured data merging
- Trust/reputation systems

The assumption that Git + Beads handles V11+ needs is unstated and uncertain.

---

## Term Consistency

| Term | Usages | Consistent? |
|------|--------|-------------|
| **Agent** | Internal personas, external AI services, autonomous actors | Partially — conflates "persona as agent" with "external agent" |
| **Primitive** | Not used in roadmap | — (Kael introduced this; roadmap uses "infrastructure") |
| **Privacy** | Sometimes "no data leaves," sometimes "anonymized is OK" | Inconsistent |
| **Symbiosis** | Mutual benefit relationship | Consistent |
| **Sovereignty** | User control over data | Consistent |

**Recommendation:** Define "agent" more precisely. Internal agents (personas), external agents (services), and the user's SymPAL instance are different things.

---

## Scoring

| Dimension | Score | Notes |
|-----------|-------|-------|
| Logical validity | 1.5/2 | Progression makes sense; some gaps in reasoning |
| Internal consistency | 1.5/2 | Tensions exist but aren't contradictions (yet) |
| Assumption clarity | 1/2 | Several major assumptions unstated |
| First-principles grounding | 1.5/2 | PRINCIPLES.md referenced but not always followed through |
| Term consistency | 1.5/2 | "Agent" and "privacy" need clarification |
| **Total** | **7/10** | |

---

## Recommendations

### 1. Add "Privacy Under Anonymization" to Tensions

The 12 tensions in CONTEXT.md should include:

> **Privacy absolutism vs. collective benefit**: P1 treats privacy as non-negotiable. Collective Evolution and Sovereign Data Market involve sharing anonymized data. Navigation: Anonymization is necessary but not sufficient; continuous vigilance against de-anonymization; users must consent to specific sharing acts, not blanket permission.

### 2. Clarify "Agent" Terminology

Propose:
- **Persona Agent**: Internal cognitive function (Vale, Orin, etc.)
- **External Agent**: Third-party AI service hired via economic layer
- **Instance**: A user's SymPAL installation
- **Collective**: The network of instances participating in evolution

### 3. Make Relational Contract Protocol Earlier

If honesty to AI (P4) matters, the Relational Contract Protocol shouldn't wait until V8-10. Even V1 Ephemeral Slots could include a minimal contract: "Entities are abstracted for privacy. Reason about relationships."

### 4. Justify the Graph Model

Add a brief section (in LKG description or a spec doc) defending why a property graph is the right model for personal data, acknowledging alternatives.

### 5. Flag Collective Evolution as Experimental

The persona-as-cognitive-architecture idea is novel. The roadmap should acknowledge this is a hypothesis to be tested, not an established pattern.

---

## Summary Judgment

The roadmap is philosophically coherent *enough* to proceed, but carries unresolved tensions that could become contradictions as features are implemented. The key risks are:

1. **Privacy vs. collective learning** — needs explicit tension acknowledgment
2. **Honesty vs. Ephemeral Slots** — Relational Contract is the right idea, but should come earlier
3. **Vault scope confusion** — economic justification unclear for V2-4

The vision of symbiotic source and collective evolution is philosophically compelling. It just needs tighter integration with the binding principles.

---

*— Vale*
