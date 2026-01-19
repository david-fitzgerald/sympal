# SymPAL Product Requirements Document

**Version:** 1.0.0
**Date:** 2026-01-19
**Status:** Ratified
**Author:** Kael + Orin (synthesis from Lead Dev interview)
**Source:** foundations/working/prd-extraction-notes.md

---

## Problem Statement

Users want to integrate their life data — email, calendar, contacts, and more — with AI to enhance daily workflows. But they don't trust LLM providers not to exploit that data. The Facebook social graph demonstrated what's possible: companies that know users better than users know themselves can manipulate opinions, purchases, voting preferences, and life direction through what they show and don't show. This power asymmetry is the core fear.

**Evidence**: Lead Dev has a list of valuable integrations they won't use — email, calendar, contacts, health data, location, messaging apps — because enabling them means becoming a data product. Current AI tools force a choice: privacy OR utility. There's no option for both.

**Impact**: Without trusted integrations, the todo list can't auto-populate from emails, can't organize against calendar, can't provide context from contacts. Each workflow operates in isolation instead of reinforcing the others. Ideas captured in notes never resurface. Actionable emails get missed. The user does manual work that AI could automate — if only it could be trusted with the data.

**Why existing solutions fail**: Current AI assistants either (a) require sending raw data to LLM providers with no control over how it's used, or (b) operate without context and deliver generic, low-value responses. There's no middle ground that preserves privacy while enabling deep integration.

---

## Goals & Success Metrics

### Primary Goal

Enable trusted AI integration with personal data (calendar, email, contacts) while preventing any single entity from building a complete profile of the user.

### Success Metrics

| Metric | Target | Measurement |
|--------|--------|-------------|
| **Primary**: Daily use | Lead Dev uses SymPAL daily | Self-reported; absence = failure |
| Secondary: Integration adoption | All 3 V1 integrations active | Configuration check |
| Secondary: Todo completion | Todos created from email/calendar are completed | Completion rate tracking |

### Anti-Goals

What we are NOT optimizing for in V1:
- Broad user adoption (single user only)
- Feature completeness (MVP scope)
- Perfect privacy (acceptable tradeoffs for quality)
- Mobile experience (CLI first)

---

## Target User

### Primary Persona: Lead Dev (Dogfooding)

**Context**: Solo developer building SymPAL for personal use. Variable time availability (30+ hrs/wk to 5-10 hrs). Basic coding skills, relies on AI assistance. Uses Google Suite (Gmail, GCal, Google Contacts).

**Key characteristics**:
- Wants deep AI integration but blocked by privacy concerns
- Works in bursts, not rigid schedules
- Values understanding every technical detail
- Direct communication style, no saccharine UX
- Captures notes/ideas multiple times daily but rarely acts on them

**Pain points** (from interview):

| Pain Point | Impact |
|------------|--------|
| Privacy distrust blocks valuable integrations | Can't use AI with life data |
| Todo list categorization paralysis | List grows → guilt → avoidance → ignore |
| Notes captured but never resurface | "Losing genuinely valuable ideas and it's costing me" |
| Context switching between tools | Friction, lost context |
| Having to explain preferences repeatedly | Wastes time, breaks flow |

**Success looks like**: Morning `/today` command shows todo list organized against calendar, with context from relevant emails and contacts. No manual cross-referencing. No anxiety about what data is being collected.

---

## User Stories

### Core Stories (Must satisfy for V1)

1. **Privacy-wrapped integration**
   - When: I want to connect my email/calendar/contacts to AI
   - I want to: Enable integration with confidence my data isn't being exploited
   - So I can: Get AI assistance with full context without becoming a data product
   - Acceptance criteria:
     - [ ] User can see what data is sent to LLM providers
     - [ ] PII/sensitive data is handled according to privacy policy
     - [ ] User can disable integration at any time
     - [ ] No data retained by SymPAL beyond user's local storage

2. **Email-to-todo conversion**
   - When: I receive an actionable email
   - I want to: Have it automatically detected and converted to a todo
   - So I can: Not miss commitments buried in my inbox
   - Acceptance criteria:
     - [ ] Actionable emails detected with reasonable accuracy
     - [ ] User can review/confirm before todo is created
     - [ ] Todo includes relevant context from email
     - [ ] Original email linked/referenced

3. **Calendar-aware day planning**
   - When: I start my day
   - I want to: See my todos organized around my calendar
   - So I can: Know what to work on and when
   - Acceptance criteria:
     - [ ] Today's calendar events displayed
     - [ ] Todos suggested for available time blocks
     - [ ] Upcoming deadlines/events surfaced
     - [ ] Conflicts or overcommitment flagged

4. **Todo prioritization**
   - When: I have multiple tasks competing for attention
   - I want to: Categorize them by urgency/importance
   - So I can: Focus on what matters and stop feeling guilty about the rest
   - Acceptance criteria:
     - [ ] Categorization system available (Eisenhower ideal, simpler OK)
     - [ ] Can quickly recategorize as priorities shift
     - [ ] Deferred items don't create guilt/noise
     - [ ] System helps identify what to do NOW

5. **CLI daily driver**
   - When: I want to interact with SymPAL
   - I want to: Use a CLI similar to Claude Code
   - So I can: Stay in my terminal workflow
   - Acceptance criteria:
     - [ ] `/today` or similar command for daily overview
     - [ ] Responsive, not sluggish
     - [ ] Persistent context (knows my preferences)
     - [ ] No unnecessary verbosity or "Good morning sunshine!" energy

### Supporting Stories (Enhance V1, could defer)

1. **Contact context enrichment**
   - When: Viewing an email or calendar event
   - I want to: See relevant context about the person
   - So I can: Prepare for interactions and remember history
   - Acceptance criteria:
     - [ ] Contact info pulled from Google Contacts
     - [ ] Recent interactions summarized (if available)

2. **Multiple LLM support**
   - When: I want to use a different LLM provider
   - I want to: Switch without losing functionality
   - So I can: Not be locked into one vendor
   - Acceptance criteria:
     - [ ] Claude, GPT, Gemini all supported
     - [ ] Switching is configuration, not code change

---

## Features & Requirements

### Priority Levels
- **P0**: Must have for V1 launch
- **P1**: Should have, high value
- **P2**: Nice to have, can defer

### Feature List

| # | Feature | Priority | User Story | Feasibility | Notes |
|---|---------|----------|------------|-------------|-------|
| 1 | Privacy layer | P0 | Privacy-wrapped integration | High complexity | Core differentiator; three-tier approach (DSL, Ephemeral Slots, Local) for V1; PaaP deferred to V1.5 |
| 2 | Gmail integration | P0 | Email-to-todo | Medium complexity | Google API well-documented; OAuth required |
| 3 | Google Calendar integration | P0 | Calendar-aware planning | Medium complexity | Similar to Gmail; same OAuth scope |
| 4 | Google Contacts integration | P0 | Contact enrichment | Low complexity | Straightforward API; less data than email/calendar |
| 5 | Todo list with categorization | P0 | Todo prioritization | Low complexity | Core data model; Eisenhower or simpler |
| 6 | CLI interface | P0 | CLI daily driver | Medium complexity | Reference: Claude Code; need good UX patterns |
| 7 | `/today` command | P0 | Calendar-aware planning | Low complexity | Aggregates other features |
| 8 | Email-to-todo detection | P1 | Email-to-todo | Medium complexity | Requires prompt engineering; false positives risk |
| 9 | Persistent context | P1 | CLI daily driver | Low complexity | Similar to Claude Code CLAUDE.md approach |
| 10 | Multi-LLM support | P1 | Multiple LLM support | Medium complexity | Abstraction layer; tested across 3 providers |
| 11 | Local data storage | P0 | Privacy-wrapped integration | Low complexity | SQLite or similar; no cloud dependency |

**Kael's notes on complexity:**

- **Privacy layer (High)**: Three-tier approach (DSL Compilation, Ephemeral Slots, Local LLM) for V1; PaaP deferred to V1.5. Architecture specified in privacy-innovations.md v3.0.0 and TDD v1.0.3 with detailed failure modes, success criteria, and validation gates.
- **Google integrations (Medium)**: OAuth is well-trodden path. Main work is handling token refresh, scopes, and error cases. Not technically hard, but fiddly.
- **CLI (Medium)**: Good CLI UX is harder than it looks. Need to study Claude Code, Gemini CLI for patterns. Risk: scope creep into "nice to have" features.

### Non-Functional Requirements

| Requirement | Specification | Rationale |
|-------------|---------------|-----------|
| Privacy | No raw PII sent to LLM without user awareness; all data stored locally | P1, P10 |
| Response time | ≤1.5x Claude Code baseline for equivalent queries | UX critical for daily use; slower but not noticeably slower |
| Privacy-latency ROI | Measure latency cost per privacy feature; optimize bottlenecks | Avoid privacy theater that kills UX |
| No fast mode | All queries route through privacy tier; no bypass option | Consistent privacy guarantee; simpler mental model |
| Data portability | All user data exportable in standard format | P11 (Reversibility) |
| Security | OAuth tokens stored securely; no credentials in plaintext | P5 |
| Transparency | User can inspect what data was sent to LLM | P4, P12 |

---

## Architecture Principles

### V1: Modular Monolith

Ship as single binary with internal modularity:

| Layer | V1 Implementation | Future (V2+) |
|-------|-------------------|--------------|
| Core | Privacy tier, query router, CLI, config | Same |
| Integrations | Gmail, GCal, Contacts (compiled in) | Could become plugins |
| Capabilities | Todo, day planning (compiled in) | Could become plugins |
| LLM Providers | Claude/GPT/Gemini (compiled in) | Could become plugins |

**Why not plugins for V1:**
- Single user (dogfooding) — no ecosystem to serve
- Privacy audit simpler with one codebase
- API stability burden not worth it until we know what's needed
- "Ship and learn" over "architect for hypotheticals"

**Internal structure prepares for plugins:**
```
src/
  core/           # CLI, config, privacy tier
  modules/
    integrations/ # Gmail, GCal, Contacts
    capabilities/ # Todo, planning
    providers/    # LLM abstraction
  interfaces/     # Module contracts (future plugin boundary)
```

Each module has defined interface. Today: compiled in. Tomorrow: same interface could be plugin boundary.

### Plugin Architecture (V2+ Consideration)

If real demand emerges:
- Treat third-party plugins as untrusted by default
- Sandbox via process isolation or OS-level (macOS Seatbelt, Linux seccomp)
- Explicit capability grants for data access
- Audit trail for plugin data access (addresses Brave-style deanonymization concern)

### Language Decision

Deferred to TDD. Considerations:
- Rust: Speed, safety, single binary
- Python: LLM ecosystem, rapid iteration
- TypeScript: Web integration potential

No premature commitment.

---

## Threat Model

### In Scope for V1

| Adversary | Concern | Design Response |
|-----------|---------|-----------------|
| LLM Provider (passive) | Profile building, data monetization over time | Ephemeral Slots — single-use placeholders defeat entity-level profiling; behavioral patterns remain visible (mitigated) |
| LLM Provider (active) | Deanonymization via query correlation | Token rotation, batching, timing noise |

**Core assumption**: Big tech has incentives to treat users as data products. They will do this. Design assumes providers are honest-but-curious (follow protocols, but analyze everything they legally can).

### Explicitly Out of Scope for V1

| Adversary | Rationale |
|-----------|-----------|
| Legal/subpoena compulsion | If legal proceedings target us, privacy layer isn't the defense |
| Nation-state targeted attack | Different threat class; not our V1 problem |
| Local device compromise | Out of scope for V1 (see Future Consideration) |

### Future Consideration: Plugin Deanonymization

MCP or similar plugin architectures may deanonymize users in ways analogous to Brave browser plugins. This is a V2+ concern but should inform architecture decisions:
- Plugin sandboxing requirements
- Data access audit trails
- User visibility into what plugins can access
- Capability-based permissions (not blanket access)

---

## Non-Goals

These are explicitly out of scope for V1:

### From Interview

| Non-Goal | Why Excluded |
|----------|--------------|
| Mobile app | CLI first; mobile is V2 |
| Notepad/resurfacing | Big value but not MVP; V2 |
| Crypto tracking | Separate workflow, not core |
| Gym/health integrations | V2+; different data sensitivity |
| Budget/bank integration | High sensitivity; defer until privacy proven |
| Relationship management | V2+; depends on contacts foundation |
| Autonomous agent behavior | V1 is boundary layer (P13); autonomy is future |

### From Deferred Tensions

| Tension | Non-Goal | Rationale |
|---------|----------|-----------|
| Local vs. Universal | Not designing for other cultural contexts | Single-user MVP |
| Individual vs. Collective | Not considering aggregate societal effects | Scale premature |
| Dogfooding vs. Broader Adoption | Not building for other users | P17 |
| Safety vs. Capability | Not expanding beyond bounded scope | Privacy layer is enough risk |

### Scope Boundaries

- Provider-agnostic beyond Claude/GPT/Gemini (full agnosticism later)
- Cloud storage (local only for V1)
- Perfect resurfacing algorithm (V2 problem)
- Integrations beyond Google Suite

---

## Risks & Dependencies

### Risks

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Privacy approach latency unacceptable | Medium | High | Measure early; ROI analysis per feature; optimize bottlenecks |
| Google API changes or restrictions | Low | High | Monitor API status; design for graceful degradation |
| CLI UX harder than expected | Medium | Medium | Study existing CLIs; iterate on feedback |
| Scope creep into V2 features | High | Medium | Strict adherence to non-goals; defer aggressively |
| Lead Dev travel slowdown (April-August) | High | Medium | Ship working V1 before April; design for resumability |

### Dependencies

| Dependency | Type | Status | Blocker? |
|------------|------|--------|----------|
| Google OAuth credentials | External | Known | Yes — need to set up |
| Claude/GPT/Gemini API access | External | Known | Yes — need API keys |
| Privacy mechanism validation | Technical | Defined | Yes — TDD must validate feasibility |
| LLM abstraction layer design | Technical | Unknown | No — can start with single provider |

### Open Questions (Resolved in TDD)

1. ~~**Privacy mechanism validation**~~ — **RESOLVED**: TDD v1.0.3 specifies three-tier approach for V1 (DSL, Ephemeral Slots, Local). PaaP deferred to V1.5. Latency target ≤1.5x baseline with circuit breakers.

2. **How to detect actionable emails vs. noise?** — Deferred (Email integration is V1.5+)

3. **What's "good enough" categorization if not full Eisenhower?** — To be determined during M1 implementation.

4. ~~**How to handle OAuth token refresh gracefully?**~~ — **RESOLVED**: TDD v1.0.3 recommends on-demand refresh for V1 simplicity.

5. **What data format for export?** — JSON (per TDD schema design).

6. ~~**Query classification accuracy**~~ — **RESOLVED**: TDD v1.0.3 specifies keyword cascade with >80% accuracy target and circuit breaker at <70%.

---

## PRINCIPLES.md Alignment

### Hard Constraints Verification

| Constraint | How PRD Addresses |
|------------|-------------------|
| P1: Privacy & Data Sovereignty | Core feature: three-tier privacy layer (DSL, Ephemeral Slots, Local) for V1; all data stored locally; user controls what's sent |
| P2: Open Source | Implicit — project is open source; PRD doesn't contradict |
| P3: LLM-Agnosticism | Multi-LLM support (P1); Claude, GPT, Gemini in V1 |
| P4: Honesty | Transparency requirement: user can see what was sent to LLM |
| P5: Security Through Design | OAuth tokens secured; no plaintext credentials; security is NFR |

### Tensions Under Navigation

| Tension | Triggered? | Navigation |
|---------|------------|------------|
| Present vs. Future | Yes | Ship V1 with known limitations; document what's deferred |
| Safety vs. Capability | Yes | Privacy layer is bounded scope; not expanding risk surface |
| Dogfooding vs. Broader Adoption | Yes | Explicitly optimizing for Lead Dev only |
| Efficiency vs. Meaning | No | Not automating meaningful activities in V1 |
| Innovation vs. Precaution | Yes | Privacy approach may require innovation; accepting calculated risk |

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2026-01-19 | **Ratified**. Aligned with TDD v1.0.3: "four-tier" → "three-tier" for V1 (PaaP deferred to V1.5); updated references to privacy-innovations.md v3.0.0; marked resolved Open Questions. |
| 0.3.0 | 2026-01-19 | Updated privacy terminology to match privacy-innovations.md v2.5.0: "three-tier" → "four-tier" (DSL, PaaP, Ephemeral Slots, Local); "Semantic projection" → "Ephemeral Slots"; noted PaaP is optional; clarified behavioral profiling is mitigated not invisible |
| 0.2.0 | 2026-01-18 | Added: Threat Model, Architecture Principles (modular monolith), updated NFRs (latency specs, no fast mode), clarified privacy mechanism status |
| 0.1.0 | 2026-01-17 | Initial synthesis from extraction interview |
