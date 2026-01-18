# PRD Adversary Challenges

**Date**: 2026-01-18
**Status**: Challenges 3-6 addressed; Challenges 1-2 addressed via research + novel approaches
**PRD Version**: 0.1.0 → 0.2.0 (pending refinement)

---

## Addressed Challenges (Via Research + Novel Approaches)

### Challenge 1: Privacy Layer is a Black Box (Critical)

**Issue**: PRD commits to P0 feature with "approach TBD" and "may require novel solutions."

**To satisfy**: Either (a) spike privacy approach in TDD before committing to PRD, or (b) define fallback V1 that works without novel privacy layer (e.g., transparency-only).

**Response**: ACCEPT challenge. Research completed + novel approaches developed.

**Research findings** (see `foundations/privacy-research.md`):
- No Free Lunch Theorem: Perfect privacy + perfect utility is mathematically impossible
- Existing approaches have significant tradeoffs (local LLM: 10-20% gap, DP: 20-40% loss)
- Industry converging on hybrid local+cloud architectures
- "Swap don't blank" for PII redaction preserves context

**Novel approaches developed** (see `foundations/privacy-innovations.md`):
1. **Semantic Projection**: LLM sees typed patterns, not identities. Preserves reasoning value.
2. **LLM as Compiler**: For structured queries, LLM returns logic, never sees data. Zero exposure.
3. **P2P Query Mixing**: V2+ approach for multi-user anonymity.

**V1 Privacy Architecture**:
```
Query → Classifier → Route by task type:
  - STRUCTURED → LLM as Compiler (zero exposure)
  - REASONING → Semantic Projection (pattern exposure only)
  - CONTENT → Local LLM (quality tradeoff accepted)
```

**PRD Change**: Replace "approach TBD" with three-tier architecture specification.

---

### Challenge 2: Quality Degradation May Be Fantasy (High)

**Issue**: 2-3% imperceptible degradation specified with no evidence it's achievable. Hidden unbounded tradeoff.

**To satisfy**: Define exposure-quality tradeoff explicitly. What data acceptable to send raw? What's minimum protection that counts? Have a degradation budget.

**Response**: ACCEPT challenge. 2-3% target was unrealistic. Honest tradeoffs now defined.

**Research-based tradeoffs**:

| Task Type | Approach | Data Exposure | Quality Impact |
|-----------|----------|---------------|----------------|
| Structured queries | LLM as Compiler | **Zero** | 0% (logic returned, not answers) |
| Reasoning tasks | Semantic Projection | Pattern only | ~5-10% (some context loss) |
| Content tasks | Local LLM | **Zero** | 10-20% vs frontier |

**Key insight**: The tradeoff varies by task type, not globally. Some queries can have zero exposure with zero quality loss (compiler approach). Others require accepting quality gap (local LLM for content).

**PRD Change**:
- Remove "2-3% imperceptible degradation" claim
- Add task-type-specific tradeoff table
- Be honest: content tasks on local LLM will be noticeably different from frontier

---

## Scope Revision (Post-Research)

Privacy architecture is now the core of V1. Other features simplified to support privacy testing.

**Revised milestones**:

| Milestone | Scope | Notes |
|-----------|-------|-------|
| M0: Foundation | CLI + Claude + local storage | Infrastructure |
| M1: Todo | Simple CRUD only | No categorization — just enough to test privacy |
| M2: Calendar | Read-only Google Calendar | Just fetch events — no write operations |
| M3: Privacy | Three-tier architecture | **The core innovation** |
| --- | **V1 SHIP LINE** | --- |
| M4: Todo+ | Categorization, /today improvements | Post-ship enhancement |
| M5: Email | User-triggered email-to-todo | Requires proven privacy layer |
| M6: Contacts | Context enrichment | Requires proven privacy layer |

**Rationale**: Todo + Calendar provide enough real data to meaningfully test all three privacy approaches. Don't build elaborate integrations until privacy is proven.

---

## Addressed Challenges

### Challenge 3: Value Prop Depends on Unproven Capability (High)

**Issue**: If privacy layer fails, you've built a worse Claude Code. Only differentiator is the thing you don't know how to build.

**To satisfy**: Identify at least one V1 capability clearly better than Claude Code even without novel privacy layer.

**Response**: MODIFY challenge framing

The Adversary assumed external success metrics. But per P17 (Dogfooding), success = lead dev uses daily, not "beats competitors."

- Learning goals about AI and coding are a valid outcome even if product is mediocre
- Personalized workflows for lead dev don't have to beat Claude Code for everyone
- "Worse Claude Code" is measured against external users; dogfooding measures against personal utility

**V1 value prop stands even without novel privacy**:
1. Learning goals accomplished (AI, coding, systems)
2. Personalized workflows (suped-up Claude Code for lead dev specifically)
3. Integrated life data (Gmail/Calendar/Contacts in one place)
4. Cross-domain persistent context
5. Proactive surfacing via `/today`

**PRD Change**: Add "Value Proposition" section stating:
- Primary: Learning + personalized daily driver
- Secondary: Privacy-wrapped integrations (if achieved)

---

### Challenge 4: Email-to-Todo Detection Hand-Waved (Medium)

**Issue**: "Reasonable accuracy" undefined. Email classification is notoriously hard.

**To satisfy**: Define acceptable false positive/negative rates, OR descope to user-triggered email-to-todo instead of auto-detection.

**Response**: ACCEPT challenge. Descope to user-triggered for V1.

**V1 Behavior**:
- User explicitly triggers conversion: "make this a todo"
- System extracts context from email, suggests todo text
- User confirms/edits before creation

**Rationale**:
- Auto-detection is genuinely hard (spam-like classification problem)
- No data yet to define achievable FP/FN rates
- User-triggered still delivers core value (email context → todo) without classification risk
- Can learn patterns from what user actually converts

**V1.1 Enhancement** (after V1 proves value):
- "Suggest, don't auto-create" — flag emails that look actionable
- User still confirms — never auto-create without approval
- Define FP/FN targets once we have usage data

**PRD Change**:
- Move "Email-to-todo detection" from P1 to P2
- Update User Story 2: "User can convert email to todo" (not "automatically detected")
- Update acceptance criteria: remove "detected with reasonable accuracy"

---

### Challenge 5: Scope Too Large for Timeline (Medium)

**Issue**: 6+ substantial systems for solo dev with basic skills, variable time, April-August travel.

**To satisfy**: Either cut scope further (single LLM, simpler privacy, no auto-detection), OR define explicit milestones with "shippable at this point" gates.

**Response**: ACCEPT challenge. Define shippable milestones + cut multi-LLM from V1.

**V1 Milestones** (each shippable gate):

| Milestone | Deliverable | Shippable? |
|-----------|-------------|------------|
| M0: Foundation | CLI skeleton + Claude + local storage | No |
| M1: Todo Core | Todo CRUD + categorization + `/today` | **Yes** |
| M2: Calendar | Google Calendar integration | **Yes** |
| M3: Privacy | Whatever research determines | **Yes** |
| M4: Email | Gmail + user-triggered email-to-todo | **Yes** |
| M5: Contacts | Google Contacts + context enrichment | **Yes** |

**Key decisions**:
- Single LLM (Claude) for all of V1
- Privacy comes BEFORE email/contacts — don't send sensitive data to LLM until privacy solved
- Calendar acceptable risk before privacy (lower sensitivity)
- Travel contingency: M1-M2 before April = usable tool

**V2+ (deferred)**:
- Multi-LLM support (GPT, Gemini)
- Self-hosting / local models (OS models lag frontier ~7 months; more powerful by V2)
- Auto email-to-todo detection

**PRD Change**:
- Add Milestones section with shippable gates
- Move Multi-LLM from P1 to V2/Non-Goals
- Reorder features to match milestone sequence

---

### Challenge 6: Persistent Context Underspecified (Low)

**Issue**: Could be trivial (config file) or major system (learning user model). Unclear what V1 means.

**To satisfy**: Specify what "persistent context" means for V1 — config file like CLAUDE.md, or something more?

**Response**: ACCEPT challenge. V1 = Config + local history (Option B).

**V1 Persistent Context**:

| Component | What It Is | Complexity |
|-----------|-----------|------------|
| User config | `SYMPAL.md` — preferences, communication style, priorities | Low |
| Todo history | Completed todos, patterns, recurring items | Low |
| Interaction history | Past queries, what user asked about whom/what | Low-Medium |
| Contact notes | User's notes on contacts (manual, not inferred) | Low |

**V1 Behavior**:
- User edits `SYMPAL.md` to set explicit preferences
- System remembers todos and their history (local DB)
- System can reference "you completed X yesterday" — recall, not inference
- User can add notes to contacts manually
- No magic learning — explicit > implicit

**V2+ (deferred)**:
- Auto-learning user model (observes and infers preferences)
- Cross-session memory synthesis ("you mentioned X last week")
- Predictive context (anticipating what user will need)

**PRD Change**:
- Specify persistent context as "SYMPAL.md config + local history"
- Add to Non-Goals: "Auto-learning user model (V2)"
- Update CLI acceptance criteria to reflect this scope

---

## Response Template

For each challenge, provide:
1. **Accept / Reject / Modify** the challenge framing
2. **Response**: How you address it
3. **PRD Change**: What changes in the PRD (if any)
