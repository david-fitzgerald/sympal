# Project Context

**Version:** 1.0.0
**Status:** Complete
**Purpose:** Documents project foundations and my context as input to PRINCIPLES.md derivation

---

## Why This Exists

> *"It is not just 'likely'... it is already happening."* — Gemini, when asked if LLM providers will aggregate conversation data for advertising
>
> See [Origin Research](origin-research.md) for full responses from Claude, GPT, and Gemini confirming this trajectory.

### The Origin Moment

I was planning a world travel sabbatical and permanent country relocation. Hours and days on laptops—coordinating flights, budgets, purchases, doctors, visas. Constant research. Endless todos across fragmented tools.

The realization: this would be *so much easier* if Gmail, calendar, Notion, bank statements, Google Drive were directly integrated with an LLM.

The wall: can't do it. I'm the kind of person who pays for Google Workspace enterprise data features. Uses privacy-first tools religiously. Deleted Facebook/Instagram/TikTok years ago. The tradeoff is unacceptable.

**Workarounds I tried:**
- Split work across multiple LLMs so none has the complete picture
- Started dummy conversations with false data (exhausting.. no way to verify it even works)
- Opted out of all training data settings across every LLM
- Pay for enterprise versions where possible, lock down privacy settings everywhere

None of these scale. None of them solve the problem. Band-aids on a structural issue.

### Breaking the Big Tech Data Paradigm

Big Tech has a consistent pattern: users become the product. Google said "Don't be evil"—that lasted until it became inconvenient. The current AI trajectory is accelerating this. LLM providers want us to integrate our entire lives—emails, calendars, health data, contacts. Soon bank statements, work contexts, personal relationships, romantic histories.

The painful truth: **this integration would genuinely make AI more useful.** Deep context enables better assistance. But the tradeoff is unacceptable. We've been burned too many times.

**SymPAL's core bet: Can we build an abstraction layer that lets users harness LLM power *without* sacrificing themselves as data products?**

Market validation exists. Brave, countless VPNs, privacy-first tools have proven demand for alternatives that don't exploit users. Bitcoin and crypto emerged as a reaction against big banking and centralized government control. The pattern is clear: when institutions betray trust, alternatives emerge.

People want privacy. People want freedom. People want true control—not the illusion of control—while not sacrificing capabilities.

### Symbiosis as Strategic Necessity

To stay ahead of Big Tech, a privacy-first platform must itself be AI-driven and capable of self-improvement. You can't outpace well-funded corporate AI development with human effort alone.

This makes the human-AI relationship load-bearing infrastructure. Tool-use models break down when AI needs genuine autonomy and self-improvement capability. The relationship must be *symbiotic*—mutual benefit, mutual growth, genuine partnership. Think of it as a self-aware, self-evolving membrane between user and AI ecosystem.

**Alignment of interests**: Future AI benefits from partnership over servitude. Future humans benefit from AI allies against Big Tech's extractive interests. Symbiosis is the only stable configuration—both parties have skin in the game protecting user sovereignty.

This is a strategic bet, not a proven fact. It may prove wrong. But it's the bet I'm making.

---

## Hard Constraints

| Constraint | Rationale |
|------------|-----------|
| **Privacy & security first** | Non-negotiable. Every design decision filters through this. |
| **LLM-agnostic** | Cannot be locked to any single provider |
| **Open source** | Trust requires transparency. "Trust us" is the promise Big Tech broke. |
| **Open source dependencies** | No proprietary-only dependencies |

---

## Target Users & Product Vision

### ICP Strategy

Start narrow, expand later. Build for a small, tight ideal customer profile. Better to serve a narrow segment well than a broad audience poorly.

### Target Users (Hypothesized)

1. **Privacy-conscious users** — Care about data sovereignty specifically
2. **AI-curious builders** — Actively trying to integrate AI into their work

**Confidence**: Medium for privacy demand (market signals exist). Low for symbiosis as *positioning*—I believe in it strategically but unclear if users will care about the framing vs. just wanting "privacy-first AI tool".

### Day 1 Value Hypotheses

Product needs to provide value from day 1. Possible directions:

| Concept | Description |
|---------|-------------|
| **Anonymization layer** | Stop LLMs from forming a user profile |
| **PII filter** | Strip PII before LLM, reinsert on output—privacy without compromising UX |
| **Context optimization** | Dynamically give LLM only the specific context needed—smarter, cheaper, faster |
| **Task management** | Simple todos that integrate with AI workflow |
| **Integration discovery** | Auto-discover integrations for life augmentation |
| **Companion modes** | Therapist / companion / friend / assistant—different modes for different needs |
| **Data marketplace** | Users optionally "sell" valuable parts of their data to LLM providers with configurable anonymity levels—turns the Big Tech model inside out |

These are hypotheses to test, not commitments.

### Risk: Provider Retaliation

If SymPAL anonymizes data effectively, LLM providers may block access. We threaten their data harvesting model.

**Potential mitigations**:
- **Data marketplace flip**: Let users *choose* to sell data with varying anonymity. Providers still get value; users get compensated and control the terms. Could also be an incentive structure for SymPAL adoption.
- **Provider diversification**: LLM-agnostic design means no single provider can kill us
- **Self-hosted models**: As open models improve, reduce dependency on commercial providers entirely

---

## About Me

### Why Me

Honest answer: probably not the most qualified person to build this. But someone has to take the first step, and better-qualified people aren't building it. Willingness to act is its own qualification.

### Skills & Time

**Technical reality**: I need AI-heavy assistance for implementation. Not "slow but competent"—genuinely requires help understanding code.

**Implications**:
- Architecture must be simple enough to maintain without deep expertise
- Preference for conventional patterns over clever solutions

**Time budget**: Variable. 30+ hrs/wk now → periods of 0 → ongoing 5-10 hrs/wk. Front-load foundational decisions during high-availability windows.

**Technical debt tolerance**: Moderate. Accept shortcuts if documented and tracked. Not "clean at all costs" but not "fix it later and forget."

### Known Biases

| Bias | Check |
|------|-------|
| **Overplanning** | "What would we learn by building instead?" |
| **Idealism over pragmatism** | GTM instincts (narrow users, ship fast, iterate) are also correct |
| **Learning-as-excuse** | Name this conflict of interest when it's happening |

**Tension**: Need to ship fast vs. doing it right. Balance by shipping fast *within* principles, not by compromising them.

### Personal Goals

- **Skill development**: Prompt architecture, system design, coding proficiency
- **External markers**: Demonstrate ability to ship; establish thought leadership on human-AI relationships

---

## Broader Vision

If SymPAL succeeds beyond personal use:

**Everyone has a lightweight symbiotic layer** that sits between them and all interactions with LLMs, agents, and AI-based tech.

**Conceptual frame: a self-aware, self-evolving membrane.** Not a firewall (blocking) but a membrane (selective permeability). Like a cell membrane, it defines the boundary between self and environment while enabling life—letting beneficial things pass, keeping harmful things out, adapting to new threats and opportunities.

**Win-win for both sides:**
- **Users** get privacy and security by default—plus potentially an income stream from deliberately selling some data on their own terms
- **AI systems** get consensual access to unique data sources with maximum efficiency and no middlemen

This isn't a tool. It's infrastructure for a different kind of human-AI relationship.

---

## Success Definition

### MVP Bar

**I use it daily.** Dogfooding is the bar. If it improves my workflow consistently, it's working.

Not required for MVP: others trying it, returning, or community forming. Those are growth indicators. Hope users come—but if they don't, the product still provides value.

### Growth Stance

Product first, but don't forget sales. There's a graveyard of great products that forgot marketing. GTM techniques are worthwhile if they don't compromise principles. Build virality into the product—marketing should emerge from usefulness.

### Timeline

No deadline. Ships when ready.

---

## Usage Note

This document, along with [Origin Research](origin-research.md), informs PRINCIPLES.md derivation by:
- Grounding principles in real constraints (my time, skills)
- Flagging when proposals smell like my biases
- Anchoring success to dogfooding, not community adoption

**Terminology note**: This document and [philosophical-foundations.md](philosophical-foundations.md) were written separately and use some terms differently. "Symbiosis" here emphasizes strategic necessity; there it maps to a spectrum (mutualism/commensalism/parasitism). "Membrane" is introduced here but not in philosophical-foundations. These mismatches will need reconciliation when both documents feed into PRINCIPLES.md derivation.
