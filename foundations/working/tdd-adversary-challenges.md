# TDD Adversary Challenges (SUPERSEDED)

> **Note**: This file documents challenges against TDD v0.2.0, which has been superseded by TDD v1.0.0+.
> For the current Adversary challenge, see: `foundations/reviews/adversary-challenge-tdd-v1.0.1.md`

**Date:** 2026-01-18
**TDD Version:** 0.2.0 (superseded by v1.0.3)
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

**Status:** Resolved

**Response:**

The 60-70% figure is a hypothesis, not a validated claim. It should be labeled as such.

**What happens if inverted (60-70% Projection)**:

1. **Privacy posture shifts** from "mostly zero exposure" to "mostly pattern-only exposure"
2. **Compiler tier still provides zero exposure** for whatever structured queries exist
3. **P1 is still satisfied** — pattern-only exposure requires explicit user consent to the privacy architecture; raw data never leaves

**Why this is acceptable**:
- Projection was designed assuming it would be the primary path for sensitive queries
- The LLM provider sees *patterns*, never identities — this is the core privacy guarantee
- Correlation mitigations (token rotation, batching, timing noise) apply regardless of volume

**Changes to TDD**:
1. Mark 60-70% assumption as "hypothesis to validate during dogfooding"
2. Add: "Architecture is resilient to inverted distribution — privacy degrades gracefully from zero-exposure to pattern-only-exposure"
3. Add Phase 3 milestone: "Measure actual query distribution after 2 weeks; adjust expectations if >50% hit Projection"

**Conclusion**: The assumption is unvalidated but non-load-bearing. The architecture works either way; only the privacy posture description changes.

---

## Challenge 2: Semantic Projection Quality Is Asserted, Not Demonstrated

**Steelman**: Semantic Projection — replacing entities with typed placeholders — preserves reasoning context while stripping identity. This is the core innovation enabling privacy-preserving cloud LLM use.

**Attack surface**: The TDD claims Projection "preserves reasoning" but provides no evidence this is true.

**Challenge**: What if typed placeholders aren't enough for good reasoning?

Example: "Should I follow up with [PERSON:colleague,senior] about [PROJECT:business]?"

The LLM may give generic advice. But: "Should I follow up with my VP of Sales about the Q4 renewal forecast?" gets context-aware advice.

The projection loses crucial information (this is a high-stakes relationship, the project has a deadline). The TDD's only quality metric is "Subjective 1-5 scale, Weekly sample review." That's not a quality floor — it's observation without a threshold.

**To satisfy this challenge**: Define a minimum acceptable quality level (e.g., "3 or above on 5-point scale for 80% of projected queries") and what happens if it's not met.

**Status:** Resolved

**Response:**

The challenge is valid. Measurement without threshold is observation, not quality control.

**Quality floor definition**:

| Metric | Threshold | Measurement |
|--------|-----------|-------------|
| Task completion | ≥90% usable answers | Binary: did the response help? |
| Reasoning quality | ≥3.0 average on 5-point scale | Weekly sample of 20 projected queries |
| Quality delta vs baseline | ≤20% degradation | Compare projected vs raw-data responses |

**If thresholds not met after 2 weeks of Phase 3B**:

1. **If quality 2.5-3.0**: Evaluate whether taxonomy can be enriched *without* creating rare combinations (privacy tradeoff)
2. **If quality <2.5**: Consider routing more queries to Local LLM, or revisit whether Projection tier is viable

**The quality/privacy tradeoff**:

Richer taxonomy improves quality but increases correlation risk:
- `[PERSON:colleague]` → generic advice, high privacy (statistically common)
- `[PERSON:colleague,senior,sales,key-stakeholder]` → context-aware advice, lower privacy (potentially unique)

The coarsening rule is the safety valve: if a type combination appears <5 times in user's data, coarsen to parent type. This prevents unique fingerprints but limits expressiveness.

**Changes to TDD**:
1. Add quality thresholds to Success Metrics section
2. Add decision point: "If projection quality <3.0 average after 2 weeks, pause Phase 3C and iterate taxonomy"
3. Note taxonomy is v0 — expect iteration based on quality measurements, constrained by privacy tradeoff

---

## Challenge 3: The Sandbox Is Critical But Underspecified

**Steelman**: Deno subprocess with deny-by-default permissions is a sound approach for sandboxing LLM-generated code. It's industry standard and well-documented.

**Attack surface**: The TDD says "Deno sandbox with deny-by-default" but doesn't specify the actual permission set.

**Challenge**: Deno's permission model is granular: `--allow-read`, `--allow-write`, `--allow-net`, `--allow-env`, etc.

What permissions does the sandbox actually grant? The Testing Strategy lists "Sandbox blocks imports", "Sandbox blocks network", "Sandbox blocks file I/O" — but these are test cases, not a specification.

If an implementation mistakenly grants `--allow-read=.` (current directory), LLM-generated code could read `config.yaml` containing API keys.

**To satisfy this challenge**: Specify the exact Deno permission flags. E.g., "Deno runs with `--deny-all` (no flags granted) plus explicit `--allow-read` only for the schema description file."

**Status:** Resolved

**Response:**

**Note**: Discrepancy identified — `privacy-innovations.md` specifies Python AST sandbox, TDD specifies Deno. TDD is the later decision (Go + Deno), so Deno is correct. `privacy-innovations.md` needs updating to align.

**Exact Deno permission specification**:

```bash
deno run --no-prompt --no-npm --no-remote [script.ts]
```

| Flag | Effect |
|------|--------|
| (no `--allow-*` flags) | All permissions denied by default |
| `--no-prompt` | Don't ask user for permissions interactively |
| `--no-npm` | Disallow npm package imports |
| `--no-remote` | Disallow fetching remote modules |

**What's denied (by default)**:
- `--allow-read` — no filesystem access
- `--allow-write` — no filesystem writes
- `--allow-net` — no network access
- `--allow-env` — no environment variables
- `--allow-run` — no subprocess spawning
- `--allow-ffi` — no foreign function interface

**Data passing**: Via stdin/stdout, not filesystem. Go process pipes data to Deno subprocess, receives result.

**Timeout**: Go enforces 5-second execution timeout; kills subprocess if exceeded.

**On permission request**: Deno exits with error (due to `--no-prompt`). Log and fail the query.

**Changes to TDD**:
1. Add "Sandbox Specification" subsection under Security & Privacy with exact flags
2. Update `privacy-innovations.md` to align with Deno (not Python AST)

---

## Challenge 4: "Default to Most Private" Isn't Architecturally Enforced

**Steelman**: The Failure Modes section says misclassified queries default to the most private tier (local). This is a reasonable fail-safe.

**Attack surface**: Vale flagged this: the behavior is stated in Failure Modes but not in the architecture or component descriptions.

**Challenge**: If "default to private" is critical to P1 compliance, it should be an architectural invariant, not a mitigation note. Where in the code path is this enforced? What component owns this behavior? Without explicit architectural placement, an implementer could reasonably build a classifier that errors on ambiguity rather than defaulting.

**To satisfy this challenge**: Add "default to local on uncertain classification" to the Privacy Tier component description in Architecture Overview.

**Status:** Resolved

**Response:**

Valid. This is an architectural invariant, not just a mitigation.

**Change to TDD**:

In Component Summary table, update Privacy Tier row:

| Component | Responsibility | Key Interfaces |
|-----------|----------------|----------------|
| Privacy Tier | Query classification, routing, projection. **On uncertain classification, defaults to Local (most private).** | Capabilities, LLM Providers |

Also add to Design Principles (section 4.1):

> 5. **Default to private**: If query classification is uncertain, route to Local LLM. Never fail-open to less private tiers.

**Rationale**: This makes the invariant visible to implementers at the architectural level, not buried in failure modes.

---

## Challenge 5: Correlation Attack Mitigations Are Absent

**Steelman**: The TDD acknowledges "LLM Provider (active) — Deanonymization" as a threat in the Threat Model.

**Attack surface**: The TDD has no mitigations for correlation attacks.

**Challenge**: `privacy-innovations.md` v0.2.1 specifies correlation mitigations: token rotation (daily), query batching, timing noise. These are referenced in the TDD header but not incorporated into the architecture.

If an LLM provider correlates query patterns over time ("this user asks about [ORG:employer] → [PERSON:colleague,senior] every Monday at 9am"), they can profile the user despite projection.

**To satisfy this challenge**: Either (a) incorporate correlation mitigations from privacy-innovations.md into the Security & Privacy section, or (b) explicitly state this is deferred and update the Threat Model to mark "LLM Provider (active)" as "Not mitigated in V1."

**Status:** Resolved

**Response:**

Incorporate mitigations with honest scope of what they do/don't prevent.

**Add to TDD Security & Privacy section — "Correlation Mitigations" subsection:**

### Correlation Mitigations (V1)

| Mitigation | Implementation | Prevents | Doesn't Prevent |
|------------|----------------|----------|-----------------|
| Token rotation | Placeholder IDs reset daily | Cross-session entity tracking | Pattern analysis, type fingerprinting |
| Query batching | Batch 2-5 queries when possible | Individual query attribution | Batch-level patterns |
| Timing noise | 100-500ms random delay | Timing-based correlation | Content-based correlation |

**What "LLM Provider (active)" can still do**:
- Analyze type combination patterns over time
- Correlate usage timing at coarse granularity
- Build aggregate profile ("user frequently asks about senior colleagues and business projects")

**What's prevented**:
- Tracking specific entities across sessions ("the same colleague from yesterday")
- Precise timing analysis
- Individual query isolation

**Accepted limitation**: A motivated adversary with long observation may still correlate patterns. Goal is practical obscurity — raise the cost of profiling, not eliminate it.

**V2 Enhancement — Query Padding (user-configurable)**:
- Insert fake/decoy queries to obscure true query volume and patterns
- Off by default (costs API credits)
- User can enable if willing to pay for enhanced privacy
- Requires plausible query generation to avoid obvious filtering

**Update Threat Model table**:

| Actor | Motivation | In Scope? | Mitigation Status |
|-------|------------|-----------|-------------------|
| LLM Provider (passive) | Profile building | Yes | Semantic Projection |
| LLM Provider (active) | Deanonymization | Yes | Partial — token rotation, batching, timing noise (V1); query padding optional (V2) |

---

## Challenge 6: Success Criteria Has No Failure Threshold

**Steelman**: Success Metrics defines measurable targets: routing accuracy >80%, latency <5s simple/<15s complex, code execution >90%.

**Attack surface**: The document doesn't say what happens if these fail.

**Challenge**: If routing accuracy is 65%, what then? Is the project a failure? Does Phase 3 continue? Is there a pivot point?

P17 (Dogfooding) says "success = daily use." But daily use might happen even if the privacy architecture fails — the Lead Dev might use a degraded system out of commitment, not because it's working.

**To satisfy this challenge**: Define thresholds that trigger reconsideration. E.g., "If routing accuracy <70% after 2 weeks of Phase 3B, revisit classifier approach before proceeding."

**Status:** Resolved

**Response:**

Valid. Targets without decision points are just wishful thinking.

**Add "Decision Points" subsection to Success Metrics:**

### Decision Points (Circuit Breakers)

| Metric | Target | Threshold | Action if Below Threshold |
|--------|--------|-----------|---------------------------|
| Routing accuracy | >80% | <70% after 2 weeks | Pause Phase 3C; revisit classifier approach |
| Projection quality | ≥3.0 avg | <2.5 after 2 weeks | Evaluate if Projection tier is viable; consider routing more to Local |
| Code execution success | >90% | <80% | Review code generation prompts; consider tighter schema constraints |
| Latency (simple) | <5s | >10s consistently | Profile pipeline; identify bottleneck |
| Latency (complex) | <15s | >30s consistently | Acceptable for V1 dogfooding if infrequent; flag for V2 optimization |

**Dogfooding validity check (P17)**:

"Daily use" is necessary but not sufficient. Add qualitative check:

> **Monthly reflection**: Is daily use happening because SymPAL is valuable, or out of commitment to the project? If the latter, that's a signal to diagnose what's not working.

Log a brief note monthly: "Would I use this if I hadn't built it?" Honest no = problem to solve.

**Changes to TDD**:
1. Add Decision Points table to Success Metrics section
2. Add dogfooding validity check note

---

## Summary

| # | Challenge | Severity | Status |
|---|-----------|----------|--------|
| 1 | Query distribution assumption unvalidated | High | Resolved |
| 2 | Projection quality has no threshold | High | Resolved |
| 3 | Sandbox permissions unspecified | High | Resolved |
| 4 | Default-to-private not in architecture | Medium | Resolved |
| 5 | Correlation mitigations missing | Medium | Resolved |
| 6 | No failure thresholds | Medium | Resolved |

---

## Vale Checkpoint Concerns (for reference)

From the Principles Checkpoint:

1. Query misclassification severity rating may be understated
2. Default-to-private behavior stated in failure modes but not architecture (→ Challenge 4)
3. Projection quality measurement is underspecified for "calculated risk" claim (→ Challenge 2)
