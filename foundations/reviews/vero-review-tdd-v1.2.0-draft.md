# Vero Certus Review: TDD v1.2.0-draft

**Reviewer**: Vero Certus
**Date**: 2026-01-27
**Document**: foundations/tdd.md v1.2.0-draft
**Stakes**: HIGH

---

## Activation

- **Document purpose**: Technical Design Document providing implementation specifications for V1 milestones (M1-M5) with privacy architecture, security controls, and quality gates
- **Downstream use**: Implementation guidance for M3-M5 development; serves as source of truth for privacy tier behavior and success criteria
- **Stakes**: HIGH — errors propagate into code that handles sensitive user data

---

## Phase 1: Purpose Alignment

| Claimed Purpose | Fulfilled? | Evidence |
|-----------------|------------|----------|
| Three-tier privacy routing (DSL, Ephemeral Slots, Local) | YES | Sections on each tier with full specifications |
| Security controls | YES | Defense-in-depth section with 5 layers |
| Quality gates | YES | Consolidated table with targets and failure actions |
| Implementation milestones | PARTIAL | M1-M2 show checkboxes but are marked complete; M3-M5 need detail alignment |

**Findings**:
1. **[MINOR]** M1 and M2 sections still show unchecked boxes (`- [ ]`) but status indicates complete. Cosmetic inconsistency.

---

## Phase 2: Internal Consistency

| Position A | Position B | Relationship |
|------------|------------|--------------|
| PRD: Email integration is P0 (lines 157-160) | TDD: Email deferred until privacy proven (line 67, Non-Goals) | **CONTRADICTION (unacknowledged)** |
| Query Classifier: 5 categories (STRUCTURED, HYBRID, REASONING, CONTENT, UNCERTAIN) | Privacy-innovations: 6 categories (adds FUZZY_DATA for PaaP) | **TENSION (acknowledged)** — PaaP deferred to V1.5 |
| TDD: "three-tier approach" (line 14) | Architecture diagram shows 4 paths (DSL, E.S, Local, + Hybrid) | CONSISTENT — Hybrid is composition, not separate tier |

**Findings**:
1. **[MAJOR]** PRD-TDD misalignment on email. PRD v1.0.0 lists "Gmail integration" as P0 required for V1 launch, but TDD explicitly defers email to V1.5+. One document must change. Either:
   - Update PRD to defer email (breaking change to ratified doc), OR
   - Update TDD to include email path (significant scope expansion)

   *This was flagged by Seren (team review). Requires explicit resolution.*

---

## Phase 3: Implementation Verification

No claimed changes from prior review to verify (this is v1.2.0-draft, first review of full sync).

**Spot-check**: Traced 3 claims from privacy-innovations.md to TDD:

| Claimed Sync | Implemented? | Location |
|--------------|--------------|----------|
| DSL golden-set testing | YES | Lines 325-333 |
| Legend escalation framework | YES | Lines 373-388 |
| Rehydration pipeline (5 steps) | YES | Lines 411-422 |

**Findings**: None. Sync appears complete.

---

## Phase 4: Downstream Readiness

**Use case**: Implementation guidance for M3-M5 development

**Requirements check**:

| Need | Addressed? | How/Where |
|------|------------|-----------|
| What code to write for Query Classifier | YES | Lines 152-201, explicit algorithm |
| SymQL grammar specification | YES | Lines 285-320 |
| How to detect NER misses | PARTIAL | States "user review" but no detection algorithm |
| What happens if Deno sandbox is breached | YES | Deny-by-default flags, no permissions granted |
| Success criteria for M3 | YES | Lines 836-839 |
| Knowledge graph schema | YES | Lines 571-612 |
| NER library choice | DEFERRED | Line 851 — "to be determined during implementation" |

**Findings**:
1. **[MAJOR]** NER library undecided. M4 (Ephemeral Slots) depends critically on NER quality. TDD defers choice to "during implementation" but provides no decision criteria or benchmark requirements. This is underspecified for a privacy-critical component.

   *Flagged by Kael. Needs decision criteria added (e.g., "NER library must achieve >X% F1 on named entity benchmark before selection").*

2. **[MINOR]** NER miss detection is understated. Line 393-394 says "HIGH (>0.8): Auto-accept" but doesn't specify what metric (precision? recall? F1?). Ambiguous for implementer.

---

## Phase 5: Meridian Risk Assessment

| Section | Confidence Level | Uncertainty Markers | Risk |
|---------|------------------|---------------------|------|
| Query Classifier (152-201) | High | 3 ("may", "if", confidence thresholds) | OK |
| DSL Hardening (323-349) | High | 2 ("should", "if") | OK |
| Ephemeral Slots (352-432) | Medium-High | 6 (extensive failure modes documented) | OK |
| Security Controls (676-758) | High | 1 ("prevent") | **BORDERLINE** |
| NER as Privacy Foundation | Very High | 0 markers in 769-792 | **[MERIDIAN RISK]** |

**Findings**:
1. **[MAJOR] [MERIDIAN RISK]** Section "Explicitly Accepted Risks" (762-792) makes confident claims about NER being "the Achilles' Heel" and lists mitigations, but uses zero uncertainty markers. The claim "User review catches misses for new entities" assumes user diligence that may not hold. The mitigation "Conservative extraction (prefer false positives)" is stated without evidence it's achievable.

   *This is exactly the kind of confident prose that can mask systematic error. NER accuracy is THE critical dependency for Ephemeral Slots — if it fails, real data goes to cloud. The document acknowledges this is high risk but presents mitigations with unwarranted confidence.*

   **Recommendation**: Add uncertainty markers. Change "User review catches misses" to "User review MAY catch misses if users are diligent — this is an untested hypothesis." Add fallback if user review proves unreliable.

2. **[MINOR]** Security Controls section (676-758) lists 5 defense layers with confident "prevents" language. While defense-in-depth is good practice, the document doesn't acknowledge that implementation errors could undermine multiple layers simultaneously.

---

## Summary

**Strengths** (preserve these):
1. Comprehensive sync with privacy-innovations.md — all major innovations captured
2. Clear quality gates with explicit targets and failure actions
3. Defense-in-depth security architecture with multiple layers
4. Honest "Explicitly Accepted Risks" section (rare and valuable)

**Findings by Severity**:
- **BLOCKING**: 0
- **MAJOR**: 3
  1. PRD-TDD email misalignment (P0 vs deferred)
  2. NER library decision criteria missing
  3. [MERIDIAN RISK] NER mitigation overconfidence
- **MINOR**: 3
  1. M1/M2 checkbox inconsistency
  2. NER confidence metric ambiguous
  3. Security controls confidence language

**Recommendation**: **REVISE**

Address the 3 MAJOR issues before finalizing:
1. Align PRD and TDD on email scope (decision required)
2. Add NER library decision criteria
3. Add uncertainty markers to NER risk section

**Confidence**: **MEDIUM** — The document is structurally sound and comprehensive. Major issues are resolvable. The Meridian Risk finding requires careful attention — NER is load-bearing for privacy.

**What would change this assessment**:
- If email scope is resolved (either PRD updated or TDD expanded) → removes one MAJOR
- If NER library benchmarks are specified → removes one MAJOR
- If NER risk section gains appropriate uncertainty → removes Meridian Risk flag
- All three resolved → SHIP

---

## Closing

Does anything in this assessment seem off given your context? I may be missing domain knowledge that changes the recommendation.
