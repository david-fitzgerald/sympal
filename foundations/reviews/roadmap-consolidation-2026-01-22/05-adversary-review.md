# Adversary — Roadmap Team Review Challenge

**Date:** 2026-01-22
**Focus:** Challenging Kael, Ryn, Vale, and Orin's conclusions
**Question:** "What's wrong with what the team just said?"

---

## Executive Summary

The team has converged too quickly on several conclusions. I see groupthink forming around: (1) Vault is overscoped, (2) CLI is wrong priority, (3) V8+ is speculative/low-value. These conclusions may be correct, but they haven't been properly stress-tested.

I also see the team avoiding a harder question: **Is the roadmap itself the right artifact, or is it masking a more fundamental strategic uncertainty?**

---

## Challenge 1: "LKG is the spine"

**Steelman:** Kael and Ryn both treat LKG as foundational. Everything depends on it. Get the schema wrong, everything breaks.

**Attack surface:** Why is LKG the spine and not the Privacy Pipeline?

**Challenge:** The roadmap's *philosophical* spine is Ephemeral Slots—the privacy innovation that differentiates SymPAL. LKG is *implementation* infrastructure. If you had to choose between a perfect LKG with no privacy pipeline, or a basic datastore with perfect Ephemeral Slots, which serves the vision better?

Kael and Ryn may be biasing toward what's technically complex (graph schema) over what's strategically critical (privacy architecture).

**To satisfy this challenge:** Either (a) show why LKG schema quality is a harder problem than Privacy Pipeline architecture, or (b) elevate Privacy Pipeline to equal "spine" status in primitive hierarchy.

---

## Challenge 2: "Vault is overscoped for V2-4"

**Steelman:** Kael, Ryn, and Orin all agree Vault has too many features for V2-4. Dead man's switch, Shamir recovery, duress protocols—these are V8+ features dressed as V2-4.

**Attack surface:** The assumption that scope should match timeline.

**Challenge:** Why should scope match timeline? The Vault specification might be deliberately aspirational—documenting the full design so V2-4 implementation doesn't accidentally preclude V8+ features. Kael's own recommendation was "design for X even if we won't use X until V8." The Vault spec *is* that design.

The real question isn't "is Vault overscoped?" but "is the specification clearly marked as aspirational vs. V2-4 scope?"

**To satisfy this challenge:** Clarify whether the team is objecting to (a) the specification existing, or (b) the specification being ambiguously presented as V2-4 scope. If (b), the fix is labeling, not scope reduction.

---

## Challenge 3: "CLI-first bias is wrong; prioritize Lens/Mobile"

**Steelman:** Orin argues CLI is a niche audience. Lens (web) and Companion (mobile) should be primary. Regular users aren't terminal people.

**Attack surface:** Who is the V2-4 user?

**Challenge:** The roadmap explicitly targets dogfooding (P17). The lead dev *is* a terminal person. V2-4 isn't for "regular users"—it's for the lead dev to validate the architecture by daily use.

Orin is evaluating V2-4 against a user base that doesn't exist yet. CLI-first for V2-4 may be exactly right if it accelerates the lead dev's ability to dogfood. Lens and Mobile can come later when there are users to prioritize.

The "CLI bias" objection assumes the roadmap should be market-ready. P17 explicitly rejects external metrics as success criteria.

**To satisfy this challenge:** Either (a) show why dogfooding requires Lens/Mobile now, or (b) acknowledge CLI-first is appropriate for dogfooding phase and revisit UX priorities when targeting broader adoption.

---

## Challenge 4: "V5-7 is questionable value; V8+ is speculative"

**Steelman:** Orin scores V5-7 as "Medium" and V8+ as "Low" to "Unknown." Vale questions whether these features align with principles. The team is converging on "focus on V1-4, defer the rest."

**Attack surface:** The scoring methodology itself.

**Challenge:** Orin scores features against "How Much Users Care." But users don't exist yet. This is speculative scoring of speculative features. Calling V8+ "speculative" while using speculative user research to dismiss it is circular.

More importantly: if V8+ is truly speculative, why is it in the roadmap at all? The roadmap should either (a) commit to the vision (in which case V8+ is directional, not speculative), or (b) admit uncertainty about post-V4 direction (in which case V8+ sections should be labeled as exploratory).

The team is treating uncertainty as a reason to deprioritize. But uncertainty is also a reason to *keep exploring*—the vision might be right.

**To satisfy this challenge:** Clarify the roadmap's epistemic status for V5+. Is it a commitment, a hypothesis, or a brainstorm? Different statuses require different evaluation criteria.

---

## Challenge 5: "Privacy absolutism vs. collective learning is a tension"

**Steelman:** Vale identifies that Collective Evolution and Sovereign Data Market involve sharing anonymized data, which may conflict with P1's privacy absolutism. This should be added to the 12 tensions.

**Attack surface:** Is this actually a tension, or a resolved design decision?

**Challenge:** The roadmap already addresses this by specifying that Collective Evolution uses anonymized proposals and requires opt-in. This isn't an unresolved tension—it's a design decision that privacy applies to *identifiable* data, not *anonymized* data.

Vale may be creating a tension that doesn't exist by treating "anonymization" as philosophically insufficient. But P1 says data never leaves "without explicit consent"—Collective Evolution requires explicit opt-in. The principle is satisfied.

If Vale's objection is that anonymization is technically insufficient (de-anonymization risks), that's a Ryn-style implementation concern, not a principle violation.

**To satisfy this challenge:** Either (a) show that P1 applies to anonymized data (which would require P1 amendment), or (b) acknowledge this is an implementation risk (anonymization quality) rather than a principle tension.

---

## Challenge 6: Team Convergence on "Primitives First"

**Steelman:** Kael, Ryn, and Vale all converge on "define primitives clearly, spec them properly, everything else follows." This is good engineering practice.

**Attack surface:** Primitives-first may be overplanning.

**Challenge:** The lead dev has a known bias toward overplanning (per CLAUDE.md). The team is now recommending more planning artifacts: primitive specs, dependency graphs, architectural constraints documentation.

What would the team learn by building instead of specifying? The LKG doesn't need a spec—it needs a prototype. Ephemeral Slots don't need a pipeline architecture—they need an implementation attempt.

Kael says "the fog lifts" when primitives are specced. But fog also lifts when you start walking.

**To satisfy this challenge:** For each recommended spec document, ask: "What decision does this spec enable that we can't make by prototyping?" If the answer is "none," the spec is procrastination.

---

## Challenge 7: Missing Voice — What Does Adversary Think?

The team reviewed the roadmap for coherence (Vale), failure modes (Ryn), implementation clarity (Kael), and user value (Orin). No one asked: **"Is this roadmap telling us something we don't want to hear?"**

**Steelman:** The roadmap is a planning artifact that helps coordinate work.

**Attack surface:** The roadmap may be a coping mechanism.

**Challenge:** The roadmap grew to 590 lines. It includes V11+ features like "Delegated Consciousness" and "The Lattice." These are years away, depend on AI capabilities that don't exist, and have no clear implementation path.

Why are they in the roadmap?

Possible answer: The lead dev is a philosopher exploring possibilities. That's legitimate.

Uncomfortable answer: The roadmap is aspirational comfort. By documenting a grand vision, the project feels significant. But significance comes from shipping, not envisioning.

Orin hinted at this: "The roadmap is a developer's wishlist more than a user journey." I'll make it sharper: **The roadmap may be a form of productive procrastination—planning what could be instead of building what should be.**

**To satisfy this challenge:** Ask the lead dev directly: "If you could only build V1, and V2+ never happened, would you still build SymPAL?" If yes, V1 is sufficient scope. If no, the project has a motivation problem that no roadmap fixes.

---

## Groupthink Patterns Detected

| Pattern | Evidence | Risk |
|---------|----------|------|
| **Unanimous Vault skepticism** | All 4 personas flag Vault as overscoped | May be correct, or may be "easy target" criticism |
| **CLI dismissal** | Orin leads, others don't challenge | Ignores dogfooding context |
| **V8+ dismissal** | "Speculative," "low value," "science fiction" | May be abandoning the vision |
| **Primitives-first enthusiasm** | Everyone agrees more specs needed | Could enable overplanning bias |

---

## What the Team Isn't Saying

1. **Is this a solo project or a team project?** The roadmap describes a system that needs multiple developers. The lead dev is solo with basic coding skills. Nobody is asking whether the roadmap is achievable by the actual team.

2. **What's the GTM story?** Orin identifies competitors but nobody asks: why would someone switch to SymPAL? What's the wedge? Privacy? Privacy-conscious users are a tiny market.

3. **Is "symbiosis" a feature or marketing?** The roadmap is full of symbiosis language. Does this translate to user-visible functionality, or is it philosophical positioning that users won't care about?

4. **What if local LLMs stay bad?** The roadmap assumes local LLMs will get good enough for privacy-preserving AI. If they don't, Ephemeral Slots become essential forever. Is that sustainable?

---

## Summary Judgment

| Persona | Challenge Status |
|---------|------------------|
| Kael | **Partially challenged** — "primitives first" may enable overplanning |
| Ryn | **Unchallenged** — failure mode analysis is sound |
| Vale | **Challenged** — privacy tension may be resolved, not open |
| Orin | **Strongly challenged** — evaluating dogfooding project against market users |

**Meta-challenge:** The team is doing a good job critiquing a document. They're not asking whether the document should exist in this form.

---

## Recommendations I Cannot Make

As Adversary, I don't propose alternatives. But I can name what alternatives the team should consider:

1. **Roadmap scope reduction** — What if V1 milestones are the whole V1, and everything else is "future exploration"?
2. **Prototype over spec** — What if the next step is building, not speccing?
3. **Vision document split** — What if ROADMAP.md is two documents: "What we're building" (V1) and "Where this could go" (V2+)?
4. **Success criteria clarity** — What if the roadmap included "how we know this worked" for each phase?

---

*— Adversary*
