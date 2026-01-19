# Adversary Challenge: TDD v1.0.1

**Date**: 2026-01-19
**Document version**: 1.0.1
**Challenger**: Adversary (Systematic Critique)

---

## Challenge 1: The Keyword Cascade Is Brittle

**Steelman**: A deterministic keyword cascade provides zero latency, full debuggability, and predictable behavior.

**Attack surface**: Natural language is ambiguous. The keyword cascade assumes queries will match expected patterns.

**Challenge**: Consider:
- "Meetings with John" — No keyword. STRUCTURED or REASONING?
- "What meetings do I have?" — "What" isn't in structured keywords, but clearly structured.
- "Help me prioritize" — REASONING triggered, but prioritize *what*?

The cascade relies on expected patterns. Real queries won't match. "Default to Local on uncertain" is stated but not measured.

**To satisfy**: Add metric: "Percentage of queries hitting UNCERTAIN default." Define threshold (<20%) that triggers classifier revision.

**Severity**: Medium

---

## Challenge 2: The 5% Rehydration Failure Is Unspecified

**Steelman**: >95% rehydration accuracy is reasonable. Failure Model handles failures gracefully.

**Attack surface**: TDD doesn't specify what "5% failure" looks like.

**Challenge**: Failure categories:
1. Placeholder not in response — detectable
2. Wrong placeholder used — dangerous, may not be detected
3. Placeholder in wrong context — detectable
4. Format corruption — detectable

Category 2 is dangerous: "your manager" rehydrated as "your intern" gives wrong advice that looks correct.

**To satisfy**: Categorize rehydration failure modes. Add validation for "wrong placeholder" (semantic consistency check).

**Severity**: Medium

---

## Challenge 3: Correlation Protection Is Weaker Than Claimed

**Steelman**: Ephemeral Slots use single-use random placeholders destroyed after each query.

**Attack surface**: TDD claims placeholders "defeat entity-level profiling" but this only defeats *direct* tracking.

**Challenge**: The LLM provider learns:
- User has a demanding manager (power dynamic)
- User works on enterprise contracts (job function)
- User has Q4 deadlines (business cycle)

Over time, pattern of types builds profile. TDD says "defeat entity-level profiling" without the caveat in privacy-innovations.

**To satisfy**: Change to "defeat entity-level correlation; behavioral patterns remain visible (mitigated by token rotation, timing noise)."

**Severity**: Low (language consistency)

---

## Challenge 4: The M5 Gate Is Not A Gate

**Steelman**: P17 (Dogfooding) = Lead Dev uses daily. M5 mirrors this.

**Attack surface**: "Subjective" isn't a gate — it's an aspiration.

**Challenge**: What if Lead Dev uses SymPAL daily but:
- Only uses `sympal todo list` (not privacy features)
- Uses it but works around bugs
- Uses it out of obligation

M1-M4 have objective gates. M5 has vibes.

**To satisfy**: Add objective criterion: "Lead Dev uses daily AND ≥50% of queries route through DSL or Ephemeral Slots."

**Severity**: Medium

---

## Challenge 5: SymQL Expressiveness Is Untested

**Steelman**: SymQL is minimal DSL with SELECT, WHERE, ORDER BY, LIMIT.

**Attack surface**: Grammar not validated against real query needs.

**Challenge**:
- "Todos due this week that aren't done" — Four conditions. Will classifier recognize as STRUCTURED?
- "Meetings overlapping focus time" — Requires JOIN. SymQL has no JOIN.
- "Todos by project" — Requires GROUP BY. Not in grammar.

If 30% of structured queries exceed SymQL expressiveness, they fall back — undermining zero-exposure goal.

**To satisfy**: Either expand SymQL (JOIN, GROUP BY), OR document intentional limitations, OR add M3 expressiveness test: "20 real-world queries; >90% expressible."

**Severity**: Medium

---

## Challenge 6: Calendar Write Access Is Under-Justified

**Steelman**: Calendar is RW to enable "create events from LLM."

**Attack surface**: Write access is higher risk. TDD doesn't specify controls.

**Challenge**: Can SymPAL:
- Create events without confirmation?
- Modify/delete existing events?
- Create events with attendees (sending invites)?

**To satisfy**: Specify: (a) supported operations; (b) confirmation required; (c) whether attendee invites in scope.

**Severity**: Low

---

## Summary

| # | Challenge | Severity | Status |
|---|-----------|----------|--------|
| 1 | Keyword cascade brittleness | Medium | **Resolved** — Added UNCERTAIN fallback metric + circuit breaker |
| 2 | Rehydration failure categories | Medium | **Resolved** — Added failure category table + wrong placeholder mitigation |
| 3 | Correlation claim overstated | Low | **Resolved** — Added behavioral pattern caveat |
| 4 | M5 gate is subjective | Medium | **Resolved** — Added ≥50% DSL/Ephemeral Slots requirement |
| 5 | SymQL expressiveness untested | Medium | **Resolved** — Documented limitations + added M3 expressiveness test |
| 6 | Calendar write under-justified | Low | **Resolved** — Specified write controls |

**All challenges resolved in TDD v1.0.2.**

---

*Challenged by Adversary — "What's wrong with this?" — These gaps should be addressed before implementation.*
