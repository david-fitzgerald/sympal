# Privacy Innovations — Novel Approaches for SymPAL

**Date**: 2026-01-18
**Status**: Concepts defined, awaiting implementation
**Context**: These are novel approaches developed during PRD challenge phase, distinct from existing research. See `privacy-research.md` for survey of existing techniques.

---

## Why This Document Exists

The privacy research revealed that existing approaches have significant tradeoffs:
- **Local LLMs**: 10-20% quality gap vs frontier
- **PII Redaction**: Loses context ("swap don't blank" helps but limited)
- **Differential Privacy**: 20-40% accuracy loss
- **HE/MPC**: Not practical for interactive use

We developed novel approaches that aren't in the literature. This document captures them for V1 implementation and future development.

---

## Core Insight

**LLMs are pattern processors.** They don't need YOUR data — they need data with the SAME PATTERNS as yours.

If we can preserve reasoning-relevant structure while removing identity, we get utility without surveillance.

---

## Novel Approach 1: Semantic Projection ("Shadow Data")

### The Idea

Maintain two parallel representations:

**Real Data (local only)**:
```
Email from John Smith (CEO, Acme Corp)
about Q4 merger with Beta Inc
sent Tuesday, marked urgent
```

**Semantic Shadow (sent to LLM)**:
```
Email from [CONTACT:executive,important,frequent]
about [TOPIC:business,sensitive,time-bound]
sent [TIME:recent], marked [PRIORITY:high]
```

The LLM reasons over the shadow. The reasoning transfers to real data. The LLM never knew identities.

### Why It's Different

| Approach | What LLM Sees | What's Lost |
|----------|---------------|-------------|
| PII Redaction | "Email from *** about ***" | Context destroyed |
| Pseudonymization | "Email from Person_A7" | Reversible; LLM still accumulates |
| **Semantic Projection** | Typed relationship pattern | Identity, but NOT reasoning value |

### How It Works

1. **Local Knowledge Graph**: Typed graph of your life data
   - Contacts: `executive`, `friend`, `family`, `vendor`
   - Topics: `business`, `personal`, `health`, `financial`
   - Relationships: `reports-to`, `collaborates`, `family-of`

2. **Projection Function**: Map real entities to typed placeholders
   - "John Smith" → `[PERSON:colleague,senior,trusted,frequent]`
   - "Q4 merger" → `[PROJECT:business,confidential,deadline-driven]`

3. **LLM Reasoning**: Cloud LLM operates on projected patterns

4. **Rehydration**: Map LLM response back to real entities

### What The LLM Provider Accumulates

Over time they learn you have:
- Important executive contacts
- Time-sensitive business matters
- Mix of personal and professional priorities

But they can NEVER know:
- WHO those contacts are
- WHAT companies you deal with
- WHAT specifically you're working on

You become statistically indistinguishable from millions of others.

### Challenges

1. **Projection quality**: Local model must correctly type entities
2. **Novel entities**: First mention of someone needs inference or user input
3. **Cross-reference attacks**: Correlation across queries could leak identity
   - Mitigation: Rotate placeholder tokens periodically
4. **Content tasks**: "Summarize this email" needs actual text
   - Mitigation: Use local LLM for content tasks

### V1 Implementation Priority

**Phase B of M3 (Privacy milestone)**

---

## Novel Approach 2: LLM as Compiler

### The Idea

For structured queries, the LLM never touches your data. Instead:

**Traditional**:
```
User: "Find emails about deadlines conflicting with vacation"
→ Send emails + calendar to LLM
→ LLM returns answer
```

**Compiler approach**:
```
User: "Find emails about deadlines conflicting with vacation"
→ Send query PATTERN to LLM (describe data shape, not data)
→ LLM returns: executable filter logic + conflict detection algorithm
→ Run locally against real data
→ Get answer
```

The LLM is a **reasoning consultant** who helps build tools but never sees your files.

### Why It's Different

- Zero data exposure (not reduced — zero)
- LLM output is reusable (same logic applies to future queries)
- Builds a library of local capabilities over time

### What It Works For

**Good fit (~60-70% of queries)**:
- Scheduling and conflict detection
- Filtering and search
- Priority scoring based on rules
- Pattern matching
- Data transformation

**Poor fit**:
- Summarization (needs content)
- Creative tasks (needs context)
- Ambiguous queries (needs data to disambiguate)

### Example

**User query**: "Show me all meetings next week with people I haven't emailed in 30 days"

**LLM receives**:
```
Data shapes available:
- calendar_events: {title, start, end, attendees[]}
- emails: {from, to, date, subject}
- contacts: {name, email}

Query: Find calendar events in date range where attendees
have no email activity in past N days
```

**LLM returns**:
```python
def find_neglected_meeting_contacts(calendar, emails, date_range, inactive_days):
    # Filter calendar events in range
    upcoming = [e for e in calendar if e.start in date_range]

    # Get attendee emails
    attendees = flatten([e.attendees for e in upcoming])

    # Find last email date per contact
    last_contact = {}
    for email in emails:
        for addr in [email.from] + email.to:
            if addr not in last_contact or email.date > last_contact[addr]:
                last_contact[addr] = email.date

    # Filter to inactive
    cutoff = today() - days(inactive_days)
    neglected = [a for a in attendees if last_contact.get(a, MIN_DATE) < cutoff]

    return [e for e in upcoming if any(a in neglected for a in e.attendees)]```
    

**User runs locally** with real data. LLM never saw calendar or emails.

### V1 Implementation Priority

**Phase A of M3 (Privacy milestone)** — implement first because:
- Lower complexity than semantic projection
- Zero data exposure proves concept immediately
- Builds local execution infrastructure needed for projection
- Query patterns inform knowledge graph schema

---


## Novel Approach 3: P2P Query Mixing (V2+)

### The Idea

Multiple users' queries get pooled and shuffled before hitting the LLM. The LLM sees queries but can't attribute them to individuals.

```
User A: query about calendar
User B: query about email        →  Mixer  →  Batch to LLM  →  Unmixer  →  Responses
User C: query about contacts
```

### Why It's Interesting

- Anonymity through crowd
- No single point sees both identity and query content
- Mixer knows who asked but not responses
- LLM sees queries and responses but not who

### Challenges

- Need critical mass of simultaneous users
- Queries should be similar enough to batch
- Timing attacks if low activity
- Trust in mixer (or need distributed mixing)

### Alternative: Distributed Entity Typing

Use P2P network specifically for semantic projection typing:

```
Query: "Meeting with John Smith about Acme merger"

Local extraction: [John Smith], [Acme], [merger]

P2P distribution:
  → Peer 1 sees: "John Smith" → votes: PERSON
  → Peer 2 sees: "Acme" → votes: ORGANIZATION
  → Peer 3 sees: "merger" → votes: BUSINESS_EVENT

Result: Type classifications without any peer seeing full query
```

### V2+ Implementation

Requires multiple users. Not applicable for V1 dogfooding.

---

## V1 Privacy Architecture

```
┌─────────────────────────────────────────────┐
│                  YOUR QUERY                  │
└─────────────────────────────────────────────┘
                      ↓
         ┌──────────────────────┐
         │   Query Classifier   │
         │   (local, simple)    │
         └──────────────────────┘
                      ↓
    ┌─────────────────┼─────────────────┐
    ↓                 ↓                 ↓
STRUCTURED        REASONING         CONTENT
(filter, search,  (prioritize,      (summarize,
 schedule)        plan, advise)      draft)
    ↓                 ↓                 ↓
┌─────────┐    ┌─────────────┐    ┌─────────┐
│ LLM as  │    │  Semantic   │    │ Local   │
│Compiler │    │ Projection  │    │  LLM    │
└─────────┘    └─────────────┘    └─────────┘
    ↓                 ↓                 ↓
Cloud LLM        Cloud LLM         Ollama
returns          reasons on        processes
LOGIC            SHADOWS           CONTENT
    ↓                 ↓                 ↓
Run locally      Rehydrate         Direct
                 locally           response
    ↓                 ↓                 ↓
└─────────────────────┼─────────────────┘
                      ↓
              ┌──────────────┐
              │   RESPONSE   │
              └──────────────┘
```

---

## Implementation Sequence

### V1 (M3: Privacy Milestone)

**Phase A: LLM as Compiler**
- Query classifier (structured vs reasoning vs content)
- Schema description generator (describe data shapes to LLM)
- Code execution sandbox (run LLM-generated logic safely)
- Test with todo + calendar data

**Phase B: Semantic Projection**
- Entity extraction from queries
- Local knowledge graph (contacts, topics, relationships)
- Projection function (real → typed placeholders)
- Rehydration function (placeholder responses → real entities)
- Test with calendar events, todo items

**Phase C: Local LLM Fallback**
- Ollama integration
- Content task routing
- Quality comparison logging (to validate tradeoffs)

### V2+

- P2P query mixing (when users exist)
- Distributed entity typing
- Federated learning across SymPAL instances
- Projection token rotation (anti-correlation)

---

## Open Questions

1. **Query classification accuracy**: How reliably can we route queries to the right tier?

2. **Projection granularity**: How specific should type tags be? `[PERSON]` vs `[PERSON:colleague,senior,trusted]`?

3. **Knowledge graph bootstrapping**: How does the graph get populated initially? User effort vs inference?

4. **Compiler output safety**: How do we sandbox LLM-generated code safely?

5. **Cross-query correlation**: How often do we rotate placeholder tokens to prevent correlation attacks?

6. **Quality measurement**: How do we measure actual quality impact of each approach?

---

## Relationship to PRD Challenges

These innovations are the response to:
- **Challenge 1**: Privacy layer is a black box → Now defined as three-tier architecture
- **Challenge 2**: Quality degradation may be fantasy → Honest tradeoffs by task type

---

*Created: 2026-01-18*
*Last updated: 2026-01-18*
