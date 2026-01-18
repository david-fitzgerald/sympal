# TDD Adversary Challenges

**Date:** 2026-01-18
**TDD Version:** 0.2.0
**Challenger:** Adversary

---

## Challenge 1: The Three-Tier Architecture Is Load-Bearing and Unvalidated

**Steelman**: The three-tier privacy architecture (Compiler/Projection/Local) is clever. It routes queries to appropriate handlers based on sensitivity, providing privacy without sacrificing utility. The design is well-reasoned.

**Attack surface**: The entire TDD rests on this architecture working. Every goal, every success metric, every claim assumes the privacy tier delivers.

**Challenge**: The Assumptions section lists "60-70% of queries are structured" as "Unvalidated" with "No citation." If this is wrong — if most queries are actually reasoning queries requiring Projection — then:
- The Compiler tier is underutilized
- More data flows through Projection than planned
- The privacy story weakens significantly

What's the fallback if the query distribution assumption is inverted?

**To satisfy this challenge**: Either (a) provide evidence for the 60-70% claim, or (b) document what happens if Projection handles 60-70% instead — is that still acceptable for P1?

**Status:** Open

**Response:**

---

## Challenge 2: Semantic Projection Quality Is Asserted, Not Demonstrated

**Steelman**: Semantic Projection — replacing entities with typed placeholders — preserves reasoning context while stripping identity. This is the core innovation enabling privacy-preserving cloud LLM use.

**Attack surface**: The TDD claims Projection "preserves reasoning" but provides no evidence this is true.

**Challenge**: What if typed placeholders aren't enough for good reasoning?

Example: "Should I follow up with [PERSON:colleague,senior] about [PROJECT:business]?"

The LLM may give generic advice. But: "Should I follow up with my VP of Sales about the Q4 renewal forecast?" gets context-aware advice.

The projection loses crucial information (this is a high-stakes relationship, the project has a deadline). The TDD's only quality metric is "Subjective 1-5 scale, Weekly sample review." That's not a quality floor — it's observation without a threshold.

**To satisfy this challenge**: Define a minimum acceptable quality level (e.g., "3 or above on 5-point scale for 80% of projected queries") and what happens if it's not met.

**Status:** Open

**Response:**

---

## Challenge 3: The Sandbox Is Critical But Underspecified

**Steelman**: Deno subprocess with deny-by-default permissions is a sound approach for sandboxing LLM-generated code. It's industry standard and well-documented.

**Attack surface**: The TDD says "Deno sandbox with deny-by-default" but doesn't specify the actual permission set.

**Challenge**: Deno's permission model is granular: `--allow-read`, `--allow-write`, `--allow-net`, `--allow-env`, etc.

What permissions does the sandbox actually grant? The Testing Strategy lists "Sandbox blocks imports", "Sandbox blocks network", "Sandbox blocks file I/O" — but these are test cases, not a specification.

If an implementation mistakenly grants `--allow-read=.` (current directory), LLM-generated code could read `config.yaml` containing API keys.

**To satisfy this challenge**: Specify the exact Deno permission flags. E.g., "Deno runs with `--deny-all` (no flags granted) plus explicit `--allow-read` only for the schema description file."

**Status:** Open

**Response:**

---

## Challenge 4: "Default to Most Private" Isn't Architecturally Enforced

**Steelman**: The Failure Modes section says misclassified queries default to the most private tier (local). This is a reasonable fail-safe.

**Attack surface**: Vale flagged this: the behavior is stated in Failure Modes but not in the architecture or component descriptions.

**Challenge**: If "default to private" is critical to P1 compliance, it should be an architectural invariant, not a mitigation note. Where in the code path is this enforced? What component owns this behavior? Without explicit architectural placement, an implementer could reasonably build a classifier that errors on ambiguity rather than defaulting.

**To satisfy this challenge**: Add "default to local on uncertain classification" to the Privacy Tier component description in Architecture Overview.

**Status:** Open

**Response:**

---

## Challenge 5: Correlation Attack Mitigations Are Absent

**Steelman**: The TDD acknowledges "LLM Provider (active) — Deanonymization" as a threat in the Threat Model.

**Attack surface**: The TDD has no mitigations for correlation attacks.

**Challenge**: `privacy-innovations.md` v0.2.1 specifies correlation mitigations: token rotation (daily), query batching, timing noise. These are referenced in the TDD header but not incorporated into the architecture.

If an LLM provider correlates query patterns over time ("this user asks about [ORG:employer] → [PERSON:colleague,senior] every Monday at 9am"), they can profile the user despite projection.

**To satisfy this challenge**: Either (a) incorporate correlation mitigations from privacy-innovations.md into the Security & Privacy section, or (b) explicitly state this is deferred and update the Threat Model to mark "LLM Provider (active)" as "Not mitigated in V1."

**Status:** Open

**Response:**

---

## Challenge 6: Success Criteria Has No Failure Threshold

**Steelman**: Success Metrics defines measurable targets: routing accuracy >80%, latency <5s simple/<15s complex, code execution >90%.

**Attack surface**: The document doesn't say what happens if these fail.

**Challenge**: If routing accuracy is 65%, what then? Is the project a failure? Does Phase 3 continue? Is there a pivot point?

P17 (Dogfooding) says "success = daily use." But daily use might happen even if the privacy architecture fails — the Lead Dev might use a degraded system out of commitment, not because it's working.

**To satisfy this challenge**: Define thresholds that trigger reconsideration. E.g., "If routing accuracy <70% after 2 weeks of Phase 3B, revisit classifier approach before proceeding."

**Status:** Open

**Response:**

---

## Summary

| # | Challenge | Severity | Status |
|---|-----------|----------|--------|
| 1 | Query distribution assumption unvalidated | High | Open |
| 2 | Projection quality has no threshold | High | Open |
| 3 | Sandbox permissions unspecified | High | Open |
| 4 | Default-to-private not in architecture | Medium | Open |
| 5 | Correlation mitigations missing | Medium | Open |
| 6 | No failure thresholds | Medium | Open |

---

## Vale Checkpoint Concerns (for reference)

From the Principles Checkpoint:

1. Query misclassification severity rating may be understated
2. Default-to-private behavior stated in failure modes but not architecture (→ Challenge 4)
3. Projection quality measurement is underspecified for "calculated risk" claim (→ Challenge 2)
