# PRD Adversary Challenges

**Date**: 2026-01-17
**Status**: Awaiting response
**PRD Version**: 0.1.0

---

## Challenges Requiring Response

### Challenge 1: Privacy Layer is a Black Box (Critical)

**Issue**: PRD commits to P0 feature with "approach TBD" and "may require novel solutions."

**To satisfy**: Either (a) spike privacy approach in TDD before committing to PRD, or (b) define fallback V1 that works without novel privacy layer (e.g., transparency-only).

---

### Challenge 2: Quality Degradation May Be Fantasy (High)

**Issue**: 2-3% imperceptible degradation specified with no evidence it's achievable. Hidden unbounded tradeoff.

**To satisfy**: Define exposure-quality tradeoff explicitly. What data acceptable to send raw? What's minimum protection that counts? Have a degradation budget.

---

### Challenge 3: Value Prop Depends on Unproven Capability (High)

**Issue**: If privacy layer fails, you've built a worse Claude Code. Only differentiator is the thing you don't know how to build.

**To satisfy**: Identify at least one V1 capability clearly better than Claude Code even without novel privacy layer.

---

### Challenge 4: Email-to-Todo Detection Hand-Waved (Medium)

**Issue**: "Reasonable accuracy" undefined. Email classification is notoriously hard.

**To satisfy**: Define acceptable false positive/negative rates, OR descope to user-triggered email-to-todo instead of auto-detection.

---

### Challenge 5: Scope Too Large for Timeline (Medium)

**Issue**: 6+ substantial systems for solo dev with basic skills, variable time, April-August travel.

**To satisfy**: Either cut scope further (single LLM, simpler privacy, no auto-detection), OR define explicit milestones with "shippable at this point" gates.

---

### Challenge 6: Persistent Context Underspecified (Low)

**Issue**: Could be trivial (config file) or major system (learning user model). Unclear what V1 means.

**To satisfy**: Specify what "persistent context" means for V1 — config file like CLAUDE.md, or something more?

---

## Response Template

For each challenge, provide:
1. **Accept / Reject / Modify** the challenge framing
2. **Response**: How you address it
3. **PRD Change**: What changes in the PRD (if any)
