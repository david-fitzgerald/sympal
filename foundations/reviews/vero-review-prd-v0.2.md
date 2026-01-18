# Review: PRD v0.2.0

## Activation

- **Document purpose**: "Product Requirements Document... Problem → Users → Goals → Features → Non-Goals"
- **Downstream use**: TDD creation (technical design for privacy architecture and implementation)
- **Stakes**: HIGH — TDD will implement what PRD specifies; errors propagate into architecture
- **Review date**: 2026-01-18
- **Prior review**: Vale Principles Checkpoint (PASS, 9.5/10)
- **Note**: Review only, not ratification. PRD remains draft to allow iteration during TDD.

---

## Phase 1: Purpose Alignment

| Claimed Purpose | Fulfilled? | Evidence |
|-----------------|------------|----------|
| Problem Statement | YES | Lines 11-19: Clear problem (privacy vs utility tradeoff), evidence (lead dev's blocked integrations), impact (manual work that AI could automate) |
| Goals & Success Metrics | YES | Lines 23-44: Primary goal stated, metrics table, anti-goals explicit |
| Target User | YES | Lines 47-71: Persona defined with context, characteristics, pain points, success vision |
| User Stories | YES | Lines 74-145: 5 core stories with acceptance criteria, 2 supporting stories |
| Features & Requirements | YES | Lines 148-187: Priority levels defined, feature table with feasibility notes, NFRs with specifications |
| Non-Goals | YES | Lines 271-301: From interview, from deferred tensions, scope boundaries |
| Risks & Dependencies | YES | Lines 305-339: Risk table with mitigations, dependencies with blockers, open questions |
| Architecture Principles | YES | Lines 191-238: Modular monolith rationale, plugin considerations, language deferral |
| Threat Model | YES | Lines 242-267: In-scope adversaries, out-of-scope with rationale, future considerations |
| PRINCIPLES.md Alignment | YES | Lines 342-362: Hard constraints verification, tensions navigation |

**Findings**: None. All sections present and substantive.

---

## Phase 2: Internal Consistency

| Position A | Position B | Relationship |
|------------|------------|--------------|
| "No fast mode — all queries route through privacy tier" (NFR, line 184) | Three-tier routing: structured/reasoning/content (Architecture diagram) | **CONSISTENT** — "no fast mode" means no bypass, not no routing |
| "≤1.5x Claude Code baseline" (NFR, line 182) | Local LLM for content tasks (10-20% quality gap per research) | **TENSION (acknowledged)** — Quality tradeoff for privacy is stated in research; latency spec is separate from quality |
| "Modular monolith" for V1 (line 193-208) | "Could become plugins" for V2+ (line 200-202) | **CONSISTENT** — Internal modularity now, plugin capability later |
| "Multi-LLM support" as P1 feature (line 168) | "Language Decision Deferred to TDD" (line 231-238) | **CONSISTENT** — LLM abstraction is independent of implementation language |
| "P0: Privacy layer" (line 159) | "Privacy mechanism validation" as open question (line 328) | **POTENTIAL TENSION** — P0 means must-have, but validation is still open. See Finding 1. |

**Finding 1 (MAJOR)**: The Privacy layer is listed as P0 (must have for V1 launch) but Open Question #1 states "TDD must validate feasibility." This creates a logical tension: if TDD validation fails, what happens to V1? The PRD should either:
- (a) Define a fallback if three-tier approach proves infeasible, or
- (b) Acknowledge that V1 scope depends on TDD findings

This is not a contradiction—it's an unstated dependency. The PRD assumes validation will succeed. That assumption should be explicit.

---

## Phase 3: Implementation Verification

The PRD incorporates changes from:
- Adversary challenges (prd-adversary-challenges.md)
- Lead Dev clarifications (threat model, latency, architecture)

| Claimed Change | Implemented? | Location |
|----------------|--------------|----------|
| Threat model added | YES | Lines 242-267 |
| Architecture principles (modular monolith) | YES | Lines 191-238 |
| Latency spec (≤1.5x baseline) | YES | Line 182 |
| No fast mode | YES | Line 184 |
| Privacy-latency ROI measurement | YES | Line 183 |
| Plugin deanonymization concern noted | YES | Lines 261-267 |
| Three-tier approach reference | YES | Lines 159, 173, 328, 348 |

**Findings**: None. All claimed changes verified.

---

## Phase 4: Downstream Readiness

**Use case**: TDD creation — technical design for privacy architecture and V1 implementation

**Requirements check**:

| TDD Need | Addressed? | How/Where |
|----------|------------|-----------|
| What problem to solve | YES | Problem Statement |
| What features to build | YES | Features table with priorities |
| What NOT to build | YES | Non-Goals section |
| Success criteria | YES | Success Metrics table |
| User stories with acceptance criteria | YES | User Stories section |
| Privacy requirements | YES | Threat Model + NFRs |
| Latency requirements | YES | NFR: ≤1.5x baseline |
| Architecture constraints | YES | Modular monolith, no plugins V1 |
| Data model hints | PARTIAL | Local storage (SQLite mentioned), but schema not specified |
| API/integration requirements | PARTIAL | Google OAuth mentioned, but API scope not detailed |
| Security requirements | YES | P5, Threat Model, OAuth security |
| LLM abstraction requirements | YES | P3, multi-LLM support as P1 feature |
| Quality measurement approach | YES | Referenced in privacy-innovations.md |

**Finding 2 (MINOR)**: Data model and API scopes are mentioned but not detailed. This is acceptable for a PRD—TDD will specify. But TDD authors should know these are intentionally deferred, not forgotten.

**Finding 3 (MINOR)**: The PRD references `privacy-innovations.md` for the three-tier approach details but doesn't summarize the key technical decisions (sandbox spec, entity taxonomy, etc.). TDD will need to read both documents. Acceptable but worth noting.

---

## Phase 5: Meridian Risk Assessment

I will assess each major section for confident claims without uncertainty markers.

| Section | Confidence Level | Uncertainty Markers | Risk |
|---------|------------------|---------------------|------|
| Problem Statement | HIGH | "They don't trust" — framed as user belief, not absolute claim | OK |
| Goals & Success Metrics | HIGH | "Primary," "Secondary" — hierarchical, not absolute | OK |
| Target User | HIGH | Specific to Lead Dev — appropriately narrow | OK |
| User Stories | MEDIUM | Acceptance criteria use checkboxes (testable) | OK |
| Features & Requirements | MEDIUM | Feasibility column acknowledges complexity; notes section hedges | OK |
| NFRs | HIGH | Specific numbers (≤1.5x, <3s) | **ATTENTION** — see below |
| Architecture Principles | MEDIUM | "Why not plugins for V1" provides reasoning; "deferred to TDD" acknowledges limits | OK |
| Threat Model | MEDIUM | "Explicitly Out of Scope" acknowledges limits; "Future Consideration" hedges | OK |
| Non-Goals | HIGH | Explicit exclusions with rationale | OK |
| Risks & Dependencies | MEDIUM | Likelihood/Impact assessed; mitigations stated | OK |
| Open Questions | N/A | Questions are uncertainty markers themselves | OK |

**NFR Section Assessment**:
- "≤1.5x Claude Code baseline" — specific but based on user preference, not measurement
- "No fast mode" — policy decision, not empirical claim
- These are requirements, not predictions. Appropriate confidence for a PRD.

**Meridian Risk Check**: I examined whether any section makes confident empirical claims that could be systematically wrong.

- The PRD references research findings (privacy-research.md) but doesn't restate them as fact
- Quality-privacy tradeoffs are acknowledged, not hidden
- The "three-tier approach" is stated as the plan, with validation explicitly required

**Findings**: No Meridian Risk detected. The PRD appropriately distinguishes between decisions (confident) and predictions (hedged/deferred).

---

## Summary

**Strengths** (preserve these):
1. Clear alignment between Problem → Goals → Features → Non-Goals chain
2. Honest about what's out of scope and why
3. Threat model is explicit about adversaries AND non-goals
4. Architecture principles justify decisions rather than just stating them
5. Open Questions section acknowledges genuine uncertainty
6. PRINCIPLES.md alignment is verified, not assumed

**Findings by Severity**:
- **BLOCKING**: 0
- **MAJOR**: 1 — Privacy layer P0 status vs. "TDD must validate" tension (Finding 1)
- **MINOR**: 2 — Data model/API deferred without explicit note (Finding 2); TDD needs both PRD + privacy-innovations.md (Finding 3)

---

## Recommendation

**SHIP (with noted concern)**

The PRD is ready for TDD work to begin. Finding 1 (P0 vs. validation tension) should be addressed, but does not block TDD—in fact, TDD is where validation happens.

**Suggested resolution for Finding 1**: Add to Open Questions or Risks:

> "If three-tier privacy approach proves infeasible during TDD, V1 scope may need revision. Fallback options: (a) transparency-only mode, (b) local-LLM-only for sensitive data, (c) scope reduction to low-sensitivity integrations only."

This makes the dependency explicit without changing the plan.

---

## Confidence

**HIGH** — The PRD is well-structured, internally consistent, and appropriately scoped. Prior Vale checkpoint passed. My concerns are minor and addressable.

**What would change this assessment**:
- If the three-tier approach has known blockers not documented here
- If there are dependencies on external systems beyond Google APIs
- If Finding 1 is considered a true contradiction rather than unstated dependency

---

## Closing

**Note per user request**: This review does NOT constitute ratification. PRD v0.2 remains draft status to allow iteration during TDD.

Does anything in this assessment seem off given your context? I may be missing domain knowledge that changes the recommendation.

---

*Reviewed by Vero Certus, 2026-01-18*
