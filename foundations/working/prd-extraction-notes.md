# PRD Extraction Notes

**Date**: 2026-01-17
**Interviewer**: Orin
**Interviewee**: Lead Dev

---

## Problem Space

### Core Problem
Want to integrate life data with AI (email, calendar, contacts, health, location, etc.) but don't trust LLM providers not to exploit that data eventually.

**The real fear** (Facebook social graph example): Companies that know you better than you know yourself can manipulate opinions, voting preferences, purchases, life direction — through what they show and don't show. It's not about the data itself; it's about the *power* that comes from deep profiling.

### Pain Intensity
- Privacy concern: Existential — blocks all valuable integrations
- Notepad problem: "Losing genuinely valuable ideas and it's costing me"
- Todo problem: Categorization paralysis → guilt → avoidance → ignore whole list

### Current State
- Todo: Simple system, struggles to categorize priority correctly
- Notes: Notion page, crude (new lines at top, no dates, no structure), captures multiple times daily, notes "die" — never resurface, never become actionable
- Tried reminders to self — ignores them

### Root Cause
No way to get deep AI integration AND retain control. Current tools force a choice: privacy OR utility.

---

## Success Vision

### Day-in-the-Life

**Morning** (`/today` command):
- Todo list organized against calendar
- Resurfaced notes with added context
- CLI initially, mobile later

**Resurfacing cadence** (concept, not exact spec):
- Top 2 from yesterday
- Top 1 from last 3 days (excl. yesterday)
- Top 1 from last 10 days (excl. last 3)
- Option to convert any to actionable task

**Selection mechanism**: Hard problem. Cross-domain ideas often most valuable. Something about extracting structure from one domain, applying to another.

**During day**: Capture is quick/passive. Value comes later through augmentation + resurfacing.

**End of day**: Not formal. Brief if anything. Works in bursts, not rigid blocks. "I don't know" is the honest answer.

### Success Metric
**Primary**: "I use it daily" (aligns with P17 Dogfooding)

### Minimum Viable
Privacy-preserving integrations + simple todo list

---

## The Core Loop

```
Capture (notes/links/screenshots)
      ↓
AI augmentation (research, summarize, categorize)
      ↓
Proactive resurfacing
      ↓
Some become actionable → Tasks
      ↓
Eisenhower prioritization
      ↓
Execution
```

Privacy layer wraps around all of it.

---

## Features

### Must-Haves (V1)

| Feature | Detail |
|---------|--------|
| Privacy layer | Wrapping calendar, email, contacts — prevent any single entity from building complete picture |
| Todo list | With categorization (Eisenhower ideal, simpler OK to start) |
| Calendar integration | Google Calendar — plan day, remind of upcoming |
| Email integration | Gmail — detect actionable emails → create todos |
| Contacts integration | Google Contacts — context for emails/calendar |
| CLI interface | Similar feel to Claude Code / Gemini CLI / Codex |

### Nice-to-Haves (V2)

| Feature | Notes |
|---------|-------|
| Notepad capture + resurfacing | "Biggest pain" but can wait |
| Mobile capture | Hacked solution fine, doesn't need polish |
| Crypto tracking | Separate workflow |
| Gym/health integrations | With wearables |
| Budget/bank integration | High-sensitivity, defer until privacy proven |
| Researcher workflow | May integrate with notepad loop |
| Relationship management | SymPAL's own contacts list |

---

## Non-Goals (V1)

| Non-Goal | Rationale |
|----------|-----------|
| Multi-user support | Dogfooding — V1 serves Lead Dev only |
| Mobile app | CLI first; mobile capture V2 |
| Notepad/resurfacing | V2 — big value but not MVP |
| Integrations beyond Google Suite | Scope discipline — prove privacy layer first |
| Gym/health integrations | V2 or later |
| Crypto tracking | Separate workflow |
| Budget/bank integration | High-sensitivity, defer |
| Autonomous agent behavior | P13 — V1 is boundary layer (important for later, not now) |
| Provider-agnostic beyond Claude/GPT/Gemini | Full agnosticism later |
| Cloud storage | Local first |
| Perfecting resurfacing algorithm | V2 problem |

---

## Privacy Approach (Exploratory)

**Goal**: Prevent any single entity from building a complete picture.

**Ideas floated**:
1. **Privacy firewall**: Strip PII before prompt → send sanitized → reinsert on return
2. **Prompt fragmentation**: Break prompt into pieces → different LLMs → recombine

**Quality threshold**: 2-3% max degradation, imperceptible. If noticeable, too much. Willing to accept some exposure for quality.

**Key insight**: Maybe quality loss can be recovered through better orchestration (thinking loops, harnesses) — Claude Code vs raw Claude proves this.

**Risk**: May require novel approaches. Currently known methods might not be sufficient.

---

## Anti-Patterns (Must Avoid)

- Context switching — hate it
- Lack of persistent context — never make user explain same thing twice
- Claude Code model (global + project context) is good reference
- Overly cheery personality — match Lead Dev's direct, dry style

---

## Constraints

| Constraint | V1 Scope |
|------------|----------|
| LLM providers | Claude, GPT, Gemini |
| Data storage | Local only |
| Interface | CLI |
| Integrations | Google Suite (Gmail, GCal, Google Contacts) |
| Technical understanding | Lead Dev must understand every detail |

---

## Risks & Open Questions

### Risks

| Risk | Type | Severity |
|------|------|----------|
| Privacy approach may not be achievable with known methods | Technical | High — existential if unsolved |
| Learning appetite vs. time available | Resource | Medium |
| Travel slow-down (April-August) | Timeline | Medium — slow, not abandon |

### Open Questions

1. What is the actual privacy mechanism? Firewall? Fragmentation? Something novel?
2. How to select which notes to resurface? Cross-domain value detection?
3. How to detect actionable emails vs. noise?
4. What's "good enough" categorization if not full Eisenhower?

---

## Top 3 Priorities (V1)

1. **Privacy** — the enabler
2. **Todo** — the daily driver
3. **CLI UX** — the experience

---

## Dealbreaker

"If I could get a better experience from just using Claude Code" — SymPAL must add clear value beyond what could be hacked together with Claude Code alone. Privacy-wrapped integrations are supposed to be that value.

---

## Raw Notes

- Notepad problem explicitly called out as "biggest problem I have with my current workflow" but deprioritized for MVP
- Relationship management / CRM-like features mentioned (birthdays, touchpoints, notes on people) — V2+
- Kindle integration for highlights mentioned — V2+
- The "self-aware, self-evolving" framing from project-context.md aligns with future autonomous agent capability — deferred, not dismissed
- Works in bursts, not rigid blocks — don't force formal wrap-up rituals
