# Orin — Roadmap User Value Audit

**Date:** 2026-01-22
**Focus:** User value, what matters, what's bloat, privacy alignment
**Question:** "Is this trajectory serving users? What do they actually need?"

---

## Executive Summary

This roadmap describes a system I'd want to use. But I'm a power user who cares about privacy and has technical skills. Most users aren't like me.

The roadmap conflates what's *philosophically interesting* with what *users need*. Some features are high-value user problems; others are architecture astronautics that won't help anyone until V11+.

**My assessment:** Strong V1-V4 user value. V5-7 questionable. V8+ is speculation.

---

## User Value Assessment by Version

### V1: Foundation — High Value ✓

| Feature | User Problem | How Much Users Care | Score |
|---------|--------------|---------------------|-------|
| Todo CRUD | "I need to track tasks" | High — everyone has tasks | 10/12 |
| Calendar integration | "What's on my plate today?" | High — core daily workflow | 10/12 |
| `sympal today` | "Give me my day at a glance" | High — this is the killer feature | 11/12 |
| Ephemeral Slots | "Don't give my data to AI companies" | Medium — privacy-conscious users only | 8/12 |
| Local LLM | "Keep everything on my machine" | Medium — power users only | 7/12 |

**V1 verdict:** Solid. The combination of local-first + privacy + daily driver is a real value prop. Ship it.

---

### V2-4: Personal Automation — Mixed Value

| Feature | User Problem | How Much Users Care | Score | Notes |
|---------|--------------|---------------------|-------|-------|
| **LKG** | "Connect my life data" | Medium | 8/12 | Users don't know they want this until they have it |
| **Foundry** | "Create my own automations" | Low-Medium | 6/12 | Power users only; most users won't write scripts |
| **Intent disambiguation** | "Just understand what I mean" | High | 10/12 | This is what users expect from AI |
| **Vault** | "Secure my keys/passwords" | Medium | 7/12 | Competes with established solutions (1Password, etc.) |
| **Glass CLI (TUI)** | "Interactive terminal" | Low | 5/12 | CLI users are a niche |
| **Lens (Web UI)** | "See my data visually" | High | 9/12 | This should be the primary interface, not CLI |
| **Workbench** | "Build automations visually" | Medium | 7/12 | Good, but assumes users want to automate |
| **Companion (Mobile)** | "Access on the go" | High | 10/12 | Mobile is where users are; this is essential |
| **Auto-upgrade** | "Keep things updated" | High | 9/12 | Users expect this; not doing it is friction |

**V2-4 concerns:**

1. **CLI-first bias**: The roadmap treats CLI as "heritage." But most users aren't terminal people. The Lens (web UI) and Companion (mobile) should be primary, CLI secondary.

2. **Vault complexity**: The Vault spec includes Shamir recovery, duress protocols, dead man's switches. As a user, I just want: (a) secure key storage, (b) password manager integration. The crypto custody features assume users hold crypto. Most don't.

3. **Foundry is for power users**: "Personal API factory" sounds exciting to developers. Regular users want pre-built solutions that just work.

**V2-4 verdict:** Good ideas, but prioritize Lens and Companion over Glass CLI. Simplify Vault for V2; crypto features can wait.

---

### V5-7: Proactive Adaptation — Questionable Value

| Feature | User Problem | How Much Users Care | Score | Notes |
|---------|--------------|---------------------|-------|-------|
| **Goal inference** | "Help me stay on track" | Medium | 7/12 | Users are skeptical of AI "coaching" |
| **Workflow synthesis** | "Notice my patterns" | Low | 5/12 | Creepy unless done right |
| **Digital tidy-up** | "Manage my digital clutter" | Medium | 7/12 | Useful if non-intrusive |
| **Wearable integration** | "Use my health data" | Low-Medium | 6/12 | Privacy-sensitive; many won't opt in |
| **Smart home integration** | "Control my home" | Medium | 7/12 | Competes with existing solutions |

**V5-7 concerns:**

1. **Proactive = annoying**: Users have been burned by "proactive" systems (Clippy, notification spam). The line between helpful and annoying is razor-thin.

2. **Goal tracking is personal**: "I noticed you haven't worked on your side project" feels invasive. Some users will love it; many will disable it.

3. **Health data is sensitive**: Wearable integration needs to be radically opt-in with clear value. "We'll tell you your sleep quality" isn't compelling enough for most users to share health data.

**V5-7 verdict:** Approach with caution. Ship as opt-in features, not defaults. Let users discover them rather than pushing them.

---

### V8-10: Protocol Alignment — Speculative

| Feature | User Problem | How Much Users Care | Score | Notes |
|---------|--------------|---------------------|-------|-------|
| **Cross-model orchestration** | "Use the best AI for each task" | Low | 4/12 | Users don't think in models |
| **Credential orchestration** | "Manage all my API keys" | Low | 4/12 | Developer problem, not user problem |
| **SMB (P2P messaging)** | "Coordinate with other SymPAL users" | Very Low | 3/12 | Assumes other users exist |
| **Agent delegation** | "Hire AI agents to do tasks" | Low | 4/12 | Requires economic infrastructure users don't have |
| **Trusted services registry** | "Find services to use" | Low | 4/12 | Marketplace problems |

**V8-10 concerns:**

1. **These are B2B features disguised as B2C**: "Agent delegation," "trusted services registry," "credential orchestration" — these solve problems for businesses and developers, not regular users.

2. **SMB assumes network effects**: P2P coordination only works if other people use SymPAL. As a user, why would I care about scheduling meetings via SMB when I can just use Google Calendar?

3. **Economic features require crypto adoption**: The roadmap assumes users have crypto wallets, understand micropayments, and want to participate in token economies. Most don't.

**V8-10 verdict:** These aren't user features. They're platform features that might enable user value later. Don't sell them as user benefits.

---

### V11+: Partnership Amplification — Pure Speculation

| Feature | User Problem | Score | Notes |
|---------|--------------|-------|-------|
| **Co-evolutionary goals** | "AI that grows with me" | ?/12 | Can't assess; depends on AI capabilities that don't exist |
| **Delegated consciousness** | "Long-running AI assistant" | ?/12 | Same |
| **The Lattice** | "Collective intelligence" | ?/12 | Same |
| **Sovereign data market** | "Monetize my data" | 5/12 | Users generally don't want this hassle |

**V11+ verdict:** This is science fiction. Fine to have a vision, but don't plan around it.

---

## What's Missing (User Perspective)

### 1. Onboarding

The roadmap describes features but not how users discover them. A new user lands on SymPAL and sees... what? A CLI? The value prop needs to be obvious in 30 seconds.

**Recommendation:** Add "First 5 Minutes Experience" to V2-4. What does a user do immediately after install?

### 2. Migration Path

If I'm switching from Todoist + Google Calendar + Notion, how do I import my data? The LKG is great, but it's empty on day 1.

**Recommendation:** Add importers for common tools to V2-4.

### 3. Sharing & Collaboration

The roadmap is very solo-user focused. Real life involves sharing: "Share this task with my spouse." "Collaborate on this project."

SMB (V8-10) addresses SymPAL-to-SymPAL communication, but what about SymPAL users working with non-SymPAL users?

**Recommendation:** Add basic sharing (export task list, share calendar view) to V2-4.

### 4. Why Not Just Use [Existing Tool]?

The roadmap never addresses competitors. As a user, I'm comparing SymPAL to:
- Todoist (tasks)
- Notion (notes + tasks + docs)
- Apple Reminders/Calendar (built-in)
- Obsidian (local-first notes)

What's the switching cost? What's the unique value?

**Recommendation:** Add competitive positioning (not in the roadmap itself, but somewhere user-visible).

---

## Privacy Red Line Check

Checking roadmap features against my privacy red lines:

| Red Line | V1-V4 | V5-7 | V8-10 | V11+ |
|----------|-------|------|-------|------|
| Data without user benefit | ✓ Clear | ⚠️ Health data could feel invasive | ✓ Clear | ⚠️ Collective learning unclear |
| Third-party data flows | ✓ Local-first | ✓ On-device processing | ⚠️ SaaS integration | ⚠️ Sovereign data market |
| Retention | Not specified | Not specified | Not specified | Not specified |
| Dark patterns | ✓ None seen | ⚠️ "Proactive" could be pushy | ⚠️ Economic incentives | ⚠️ Market incentives |
| Contradicts SymPAL commitment | ✓ Aligned | ✓ Aligned | ⚠️ See below | ⚠️ See below |

**V8-10/V11+ privacy concern:**

The Sovereign Data Market lets users sell anonymized data. This is opt-in and consent-based. But:
- Does SymPAL take a cut? If so, SymPAL profits from users sharing data — potential violation of P10.
- Does SymPAL show users how much their data is worth? Could create pressure to share.
- "Anonymized" is a technique, not a guarantee. Users may not understand re-identification risks.

**Recommendation:** If Sovereign Data Market proceeds, SymPAL must:
1. Take zero cut (or donate cut to user-chosen charity)
2. Show risks clearly ("your data may be de-anonymized")
3. Default to not participating; require active opt-in for each data sale

---

## Bloat Identification

Features that could be cut without users noticing:

| Feature | Why It's Bloat | Alternative |
|---------|----------------|-------------|
| **Duress Protocol** (Vault) | Most users won't face coercion | Cut from V2-4; add if crypto custody becomes common |
| **Dead Man's Switch** (Vault) | Inheritance planning is rare | Cut from V2-4 |
| **Pre-Transaction Simulation** (Vault) | Assumes active crypto trading | Cut from V2-4 |
| **Glass CLI TUI** | Niche audience | Keep basic CLI; skip TUI for V2 |
| **Visual Script Builder** | Power feature | Can be community-contributed later |
| **Neural/Cybernetic Effectors** | Science fiction | Remove from roadmap entirely |
| **SymPAL Hardware Bridge** | Science fiction | Remove from roadmap entirely |

---

## Prioritization Recommendation

If I were a user, here's what I'd want first:

**Must Have (V2):**
1. `sympal today` with calendar + todos
2. Basic web UI (Lens v1)
3. Mobile companion (read-only)
4. Import from common tools
5. Simple export (reversibility)

**Should Have (V3-4):**
1. LKG (but invisible to users; it powers features)
2. Intent disambiguation ("remind me to call Sarah")
3. Email integration (headers only)
4. Basic Foundry (pre-built recipes, not custom scripting)
5. Full mobile companion (capture + display)

**Could Have (V5+):**
1. Proactive suggestions (opt-in)
2. Advanced Vault features
3. Goal tracking
4. Everything else

---

## Summary Judgment

| Version | User Value | Notes |
|---------|------------|-------|
| V1 | **High** | Solid foundation, ship it |
| V2-4 | **Medium-High** | Good but CLI-biased; simplify Vault; prioritize Lens + Mobile |
| V5-7 | **Medium** | Proceed cautiously; proactive features are risky |
| V8-10 | **Low** | Platform features, not user features |
| V11+ | **Unknown** | Science fiction; fine as vision, don't plan around it |

**Bottom line:** The roadmap is a developer's wishlist more than a user journey. Reframe V2-4 around what users do in their first week, not what the architecture can support.

---

*— Orin*
