# Vero Certus Review: TDD v1.0.2

**Date**: 2026-01-19
**Document version reviewed**: 1.0.2
**Reviewer**: Vero Certus (Final Reviewer)

---

## Activation

- **Document purpose**: "Build a privacy layer that enables AI-assisted workflows (calendar, todos) while preventing any single entity from building a complete profile of the user."
- **Downstream use**: Implementation (Phase 3). Developers will use this document to build SymPAL V1. This is the authoritative technical specification.
- **Stakes**: HIGH — errors propagate into implemented code, security vulnerabilities, and privacy failures
- **Review date**: 2026-01-19
- **Prior reviews**: Vale checkpoint (v1.0.0), Adversary challenge (v1.0.1) — both passed, fixes integrated

---

## Phase 1: Purpose Alignment

| Claimed Purpose | Fulfilled? | Evidence |
|-----------------|------------|----------|
| Privacy layer enabling AI workflows | YES | Three-tier architecture (DSL, Ephemeral Slots, Local) with explicit privacy guarantees per tier |
| Prevent profile building | YES | Ephemeral Slots with session-scoped legends, correlation mitigations documented |
| Technical spec for V1 | YES | Implementation Plan (M1-M5), schema, dependencies, test strategy all present |
| Align with PRINCIPLES.md | YES | P1, P3, P4, P5, P10, P11, P12, P17 explicitly referenced |

**Findings**: Purpose alignment is strong. Document does what it claims.

---

## Phase 2: Internal Consistency

| Position A | Position B | Relationship |
|------------|------------|--------------|
| "Default to private" (line 83) | "UNCERTAIN → Local LLM" (line 175) | CONSISTENT — Local is most private tier |
| "No vendor lock-in" (line 20) | "V1 = Claude + local" (line 61) | TENSION (acknowledged) — explicitly noted as deferred to V2 |
| "Zero exposure for structured" (line 51) | "DSL Compilation" tier (line 201) | CONSISTENT — DSL sends query + schema, not data |
| Calendar "RW" (line 108) | "Create only" (line 498) | TENSION (acknowledged) — diagram shows RW, text limits to create-only. Diagram shows aspiration, text is V1 scope |
| ">80% routing accuracy" (line 50) | ">90% DSL execution" (line 602) | CONSISTENT — different metrics for different stages |
| "≤1.5x latency" (line 54) | Cost table shows 1-5s (lines 305-310) | CONSISTENT — Claude baseline ~2s, so 1-5s within 1.5x envelope |

**Findings**:

One visual inconsistency: System diagram (line 108) shows "Calendar(RW)" but V1 is create-only. This is confusing for implementers.

**Severity**: MINOR — text is authoritative; diagram should match.

---

## Phase 3: Implementation Verification

Changes claimed from Vale checkpoint + Adversary challenge:

| Claimed Change | Implemented? | Location |
|----------------|--------------|----------|
| `sympal log` added to M1 | YES | Lines 483-484 |
| Cloud LLM decision clarified | YES | Lines 625, 663 |
| UNCERTAIN fallback rate metric | YES | Lines 594, 606 |
| Rehydration failure categories | YES | Lines 279-288 |
| Correlation claim caveat | YES | Line 416 |
| M5 gate made objective | YES | Line 533 |
| SymQL limitations documented | YES | Lines 220-225 |
| Calendar write controls | YES | Lines 497-500 |

**Findings**: All claimed changes verified. No partial implementations detected.

---

## Phase 4: Downstream Readiness

**Use case**: V1 implementation by lead dev with AI assistance

**Requirements check**:

| Need | Addressed? | How/Where |
|------|------------|-----------|
| Language choice | YES | Go (line 621) |
| Schema definition | YES | Lines 349-390 |
| External dependencies | YES | Lines 648-655 |
| Implementation order | YES | M1-M5 with dependency graph (lines 536-551) |
| Success criteria | YES | Lines 582-613 |
| Test strategy | YES | Lines 556-578 |
| Security requirements | YES | Lines 432-438 |
| What to build first | YES | M1: Foundation (lines 477-486) |
| How to validate privacy tier | YES | Gate criteria for M3, M4, M5 |
| What to do when uncertain | YES | "Default to Local" (line 175) |
| Error handling | YES | User-Facing Failure Model (lines 316-331) |
| OAuth flow | PARTIAL | Mentioned (line 490) but no detail on refresh handling |

**Findings**:

1. **OAuth token refresh** is mentioned as an open question in PRD (line 337 PRD) but not addressed in TDD. Implementer will need to make this decision.

   **Severity**: MINOR — well-trodden path, implementer can handle.

2. **NER library choice** not specified. Which NER will be used? SpaCy? Go-native?

   **Severity**: MINOR — implementation detail, but worth noting for planning.

---

## Phase 5: Meridian Risk Assessment

Scanning for sections with extended confident claims and low uncertainty markers.

| Section | Confidence Level | Uncertainty Markers | Risk |
|---------|------------------|---------------------|------|
| Problem Statement | HIGH | 0 markers (~100 words) | OK — problem statement, not technical claim |
| Query Classifier | HIGH | 1 marker ("UNCERTAIN" category exists) | OK — hedges by including default path |
| DSL Compilation | HIGH | 1 marker ("Intentional Limitations") | OK — acknowledges constraints |
| Ephemeral Slots | HIGH | "Partial" for wrong placeholder detection | OK — uncertainty surfaced |
| Accepted Risks | HIGH | Multiple "Why we accept it" | OK — explicit risk acknowledgment |
| Success Metrics | HIGH | "Circuit Breakers" for failure modes | OK — failure paths explicit |
| Cost Model | MEDIUM | None, but uses "~" estimates | OK — estimates marked as estimates |
| Threat Model | HIGH | "Out of scope" for nation-state | OK — explicit scoping |

**Uncertainty markers found**:
- "may" (line 289)
- "~40-50%" (lines 72-74, estimates marked)
- "Partial" (line 284)
- "UNCERTAIN" as classification category
- "If >30% ... expand" (line 225)
- "should" (lines 464, 465)

**Total for ~4000 words substantive content**: ~10-15 markers. This exceeds the 2 per 500 words threshold for low-risk.

**Findings**: No Meridian risk detected. Document appropriately hedges technical claims and explicitly acknowledges what it doesn't know.

---

## Summary

**Strengths** (preserve these):
1. Explicit privacy guarantees per tier with measurable criteria
2. Circuit breakers with specific thresholds and actions
3. User-facing failure model with clear principles
4. Rehydration failure categorization with severity levels
5. Acknowledged risks with mitigations
6. Clear implementation sequence with dependency graph

**Findings by Severity**:
- **BLOCKING**: 0
- **MAJOR**: 0
- **MINOR**: 3
  1. System diagram shows Calendar(RW) but V1 is create-only
  2. OAuth token refresh handling not specified
  3. NER library choice not specified

**Recommendation**: **SHIP**

**Confidence**: HIGH — Document is well-structured, internally consistent, and appropriately humble about uncertainties. Vale and Adversary challenges have strengthened it. The minor issues are implementation details that don't block starting work.

**What would change this assessment**:
- If implementation revealed the Query Classifier routes incorrectly >30% of the time (but this is measurable post-implementation)
- If rehydration accuracy can't reach >90% (testable hypothesis)
- If latency exceeds 1.5x baseline consistently (testable hypothesis)

---

## Closing

Does anything in this assessment seem off given your context? I may be missing domain knowledge that changes the recommendation.

---

*Reviewed by Vero Certus — "Your deliverable: certainty that the foundation is sound."*
