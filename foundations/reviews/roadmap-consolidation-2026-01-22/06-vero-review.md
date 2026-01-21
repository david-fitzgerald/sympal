# Vero Certus — Roadmap Systematic Error Review

**Date:** 2026-01-22
**Focus:** Systematic errors, coherence masking gaps, Meridian-style risks
**Downstream use:** Roadmap consolidation and prioritization decisions

---

## Activation

- **Document purpose:** "V1 milestones, long-term vision, V2+ ideas" (from CONTEXT.md)
- **Downstream use:** Inform implementation priorities, prevent architectural decisions that block future vision
- **Stakes:** HIGH — Roadmap errors compound across years of development
- **Review approach:** Adapted five-phase protocol for strategic document

---

## Phase 1: Purpose Alignment

| Claimed Purpose | Fulfilled? | Evidence |
|-----------------|------------|----------|
| V1 milestones | YES | M1-M5 clearly defined with scope |
| Long-term vision | PARTIAL | V2-V11+ described but relationship to vision unclear |
| V2+ ideas | YES | Extensive feature descriptions present |

**Findings:**

The roadmap conflates three distinct functions:
1. **Implementation plan** (M1-M5) — what to build next
2. **Architecture vision** (primitives, frameworks) — how pieces connect
3. **Possibility space** (V11+) — where this could go

These are interleaved rather than separated. The reader cannot easily distinguish "we will build this" from "this could exist someday."

**Severity:** MAJOR — Ambiguous commitment level creates planning uncertainty.

---

## Phase 2: Internal Consistency

| Position A | Position B | Relationship |
|------------|------------|--------------|
| "LLM-Agnosticism" (V1, Principle 3) | "Cross-Model Orchestration" (V8-10) | CONSISTENT — orchestration assumes multiple models |
| "Local-first" (V1, privacy) | "SaaS Effectors" (V8-10) | TENSION (acknowledged in Ephemeral Slots design) |
| "Agents propose only" (Collective Evolution) | "Self-developing through human-AI symbiosis" (V11+) | CONTRADICTION (unacknowledged) — noted by Vale |
| "Privacy absolutism" (P1) | "Sovereign Data Market" (V11+) | TENSION (partially acknowledged) — needs explicit navigation |
| "V2-4 Infrastructure" scope | Vault full specification | CONTRADICTION (unacknowledged) — Vault spec exceeds V2-4 timeline |

**Findings:**

1. The "agents propose only" vs. "self-developing" contradiction is unresolved. The roadmap doesn't specify when or how the constraint relaxes.

2. The Vault specification (Shamir, dead man's switch, duress protocols) appears in V2-4 section but describes V8+ capability. This is not scope ambiguity—it's a contradiction between placement and content.

3. The "Collective Evolution" section assumes a trust model that isn't specified. How do thousands of instances coordinate without a central authority? The Beads/Git mention is insufficient.

**Severity:**
- #1: MAJOR — Architectural decision needed
- #2: MINOR — Labeling fix sufficient
- #3: BLOCKING for V11+ — Cannot assess feasibility without trust model

---

## Phase 3: Implementation Verification

*Note: This phase typically verifies claimed changes from peer review. Adapting for roadmap context.*

| Claimed Primitive | Spec Exists? | Dependencies Clear? |
|-------------------|--------------|---------------------|
| LKG (Local Knowledge Graph) | NO | Partial — Kael identified |
| Adapter Framework | NO | NO |
| Effector Framework | NO | NO |
| Agent Protocol | NO | NO |
| Privacy Pipeline | NO | Partial — described in privacy-innovations.md |
| SymQL Runtime | NO | NO |
| Vault Core | PARTIAL | YES — detailed in roadmap, but not standalone spec |
| Audit Subsystem | NO | NO |
| LLM Abstraction Layer | NO | NO |

**Findings:**

8 of 9 primitives lack specification documents. The roadmap describes them but doesn't define interfaces, data models, or extension points.

This isn't necessarily wrong—primitives can be specced during implementation. But the roadmap presents them as defined ("LKG v1: Schema spec, SQLite+FTS backend") when they're actually undefined.

**Severity:** MAJOR — False sense of definition could lead to incorrect implementation assumptions.

---

## Phase 4: Downstream Readiness

**Use case:** Determine what to build next after M5

**Requirements check:**

| Need | Addressed? | How/Where |
|------|------------|-----------|
| Clear V2 scope | PARTIAL | Features listed, but scope vs. aspiration unclear |
| Primitive dependencies | NO | Kael diagrammed, but not in roadmap |
| Resource requirements | NO | Solo dev constraint not addressed |
| Success criteria by version | NO | P17 (dogfooding) applies globally, no version-specific criteria |
| Exit criteria for V1 completion | NO | M5 deliverables listed, no "V1 done" definition |

**Findings:**

The roadmap lacks:
1. **V1 completion criteria** — When is V1 "done"? After M5 ships? After dogfooding period?
2. **V2 entry criteria** — What must be true before V2 work begins?
3. **Resource alignment** — The roadmap doesn't acknowledge the solo-dev constraint. Features are designed for a team.

**Severity:** MAJOR — Without completion criteria, "V1" becomes perpetual.

---

## Phase 5: Meridian Risk Assessment

| Section | Confidence Level | Uncertainty Markers | Risk |
|---------|------------------|---------------------|------|
| V1 Milestones | HIGH (appropriate) | Few needed — concrete scope | OK |
| V2-4 Infrastructure | HIGH | "requires," "assumes" — but no "might fail" | ⚠️ ELEVATED |
| V5-7 Proactive Features | MEDIUM | "optional," "could" present | OK |
| V8-10 Economic Features | HIGH | Vault spec very detailed; no uncertainty | [MERIDIAN RISK] |
| V11+ Vision | HIGH | "will," "enables" — reads as certain | [MERIDIAN RISK] |
| Collective Evolution | HIGH | "thousands of instances" — no hedging | [MERIDIAN RISK] |

**Findings:**

Three sections show Meridian Risk patterns:

1. **V8-10 Economic Features**: The Vault specification is extremely detailed (HSM, Argon2, Shamir, guardian protocols). This detail creates false confidence. The spec reads as "this is how it will work" rather than "this is one possible design." No uncertainty about:
   - Whether HSM integration is feasible on consumer hardware
   - Whether users will accept social recovery complexity
   - Whether the threat model matches actual threats

2. **V11+ Vision**: "Delegated Consciousness," "The Lattice," "Co-Evolutionary Goals" are presented as future features, not speculative possibilities. Language like "SymPAL becomes a trusted partner in managing..." asserts outcomes that depend on AI capabilities that may never exist.

3. **Collective Evolution**: "Thousands of SymPAL instances" coordinating via Git is presented without uncertainty about:
   - Whether Git-based coordination scales
   - Whether anonymization is sufficient
   - Whether humans will actually review proposals
   - Whether the persona architecture produces meaningful synthesis

**Root cause:** The roadmap was created through expansion (brainstorming), not filtering (critique). Each idea was added in its most optimistic form.

**Severity:** BLOCKING for ratification — The roadmap cannot be treated as a commitment document without uncertainty markers.

---

## Summary

**Strengths** (preserve these):
1. V1 milestones are well-scoped and achievable
2. Primitive identification (via Kael's analysis) provides real architectural clarity
3. Privacy architecture (Ephemeral Slots, Privacy Pipeline) is genuinely novel
4. The I/O architecture (Adapter/Effector duality) is clean and extensible

**Findings by Severity:**

| Severity | Count | Items |
|----------|-------|-------|
| BLOCKING | 2 | Collective Evolution trust model undefined; Meridian risks in V8-V11+ |
| MAJOR | 4 | Purpose conflation; "agents propose only" contradiction; primitives undefined; completion criteria missing |
| MINOR | 2 | Vault scope labeling; terminology inconsistency ("agent") |

---

## Recommendation: REVISE

The roadmap is a valuable brainstorm and vision document. It cannot be treated as an implementation plan or commitment document in its current form.

**Required revisions before roadmap can guide V2+ planning:**

1. **Separate commitment from possibility**
   - V1 (M1-M5): COMMITTED — this is what we're building
   - V2-4: PLANNED — intended direction, subject to learning from V1
   - V5+: EXPLORATORY — possibilities, not commitments

2. **Add uncertainty markers to V8+**
   - "If HSM integration proves feasible..."
   - "Assuming local LLMs reach sufficient capability..."
   - "This depends on network effects that may not materialize..."

3. **Define V1 completion criteria**
   - What must be true for V1 to be "done"?
   - What's the handoff condition to V2?

4. **Resolve the "agents propose only" vs. "self-developing" contradiction**
   - Either commit to permanent human approval, or
   - Specify conditions under which constraint relaxes

5. **Acknowledge resource constraints**
   - Solo dev with basic coding skills
   - LLM-assisted development
   - What this means for timeline and scope

**Confidence:** HIGH — These are structural issues, not interpretation disagreements.

**What would change this assessment:**
- If the roadmap is explicitly labeled as "vision document, not commitment" — then Meridian risks become acceptable (visions can be aspirational)
- If primitive specs are created — then "undefined" findings resolve
- If completion criteria are added — then "perpetual V1" risk resolves

---

## Team Review Integration

| Persona | Key Finding | Vero Assessment |
|---------|-------------|-----------------|
| Kael | Primitives need specs | **CONFIRMED** — Phase 3 shows 8/9 undefined |
| Ryn | Architectural constraints could block future | **CONFIRMED** — But constraints are well-identified |
| Vale | Privacy vs. collective tension | **PARTIALLY CONFIRMED** — Tension exists; may be resolvable |
| Orin | V8+ is speculative/low-value | **REFRAMED** — Issue is confidence level, not value |
| Adversary | Team converging too quickly | **CONFIRMED** — Phase 5 shows systematic optimism |

---

## Closing

This review applies systematic error detection to a strategic document. The findings are structural—the roadmap's content is valuable, but its form creates risks.

Does anything in this assessment seem off given your context? The lead developer has knowledge about project priorities and constraints that may change these recommendations.

---

*— Vero Certus*
