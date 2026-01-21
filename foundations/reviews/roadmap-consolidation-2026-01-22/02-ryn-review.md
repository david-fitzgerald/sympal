# Ryn — Roadmap Failure Modes Review

**Date:** 2026-01-22
**Focus:** Architectural constraints, failure cascades, decisions that block future vision
**Question:** "How will this fail? What decisions paint us into corners?"

---

## Executive Summary

The roadmap is ambitious. Ambitious systems fail in ambitious ways. My concern isn't that individual features will fail — it's that **early architectural decisions could structurally prevent later features from being possible**.

I've identified 7 critical constraint points where wrong decisions block the vision, and 12 failure modes across the primitive stack.

---

## Architectural Constraint Analysis

These are decisions that, if made wrong, cannot be easily fixed later.

### Constraint 1: LKG Schema Rigidity

**The assumption:** LKG schema can evolve as needs change.

**What breaks this:** If early schema is too rigid (SQL-style normalized tables), adding new entity types or relationship types requires migrations. If schema is too loose (pure JSON blobs), queries become impossible to optimize.

**Cascade:**
- Rigid schema → Can't add new adapters without schema changes → Adapter Framework is constrained
- Loose schema → Ephemeral Slots can't reliably extract entity types → Privacy Pipeline fails
- Either way → V5-7 proactive features need rich relationship queries → won't work

**Severity:** Critical. LKG is the spine.

**Mitigation:** Design LKG as a property graph from day 1. Nodes have types but also arbitrary properties. Edges have types. Schema is "descriptive not prescriptive" — new types can be added without migration.

---

### Constraint 2: Privacy Pipeline Reversibility

**The assumption:** Ephemeral Slots projection is for LLM queries only.

**What breaks this:** V8-10 SaaS actions require the LLM to generate API call structures with placeholders, then SymPAL rehydrates. If the Privacy Pipeline only handles text-in/text-out, structured rehydration (JSON, API payloads) fails.

**Cascade:**
- Text-only pipeline → Can't rehydrate API responses → SaaS Effectors don't work
- SaaS Effectors don't work → "Broader World" integration phase blocked
- Phase 3 blocked → Economic features can't interact with external services

**Severity:** Critical. Blocks V8-10 entirely.

**Mitigation:** Design Privacy Pipeline with structured I/O from start:
- Projection: Entity → Slot (works for any data type)
- Rehydration: Slot → Entity (works in text, JSON, or any structured format)
- Legend: Bidirectional mapping, type-aware

---

### Constraint 3: Agent Protocol Isolation

**The assumption:** Agents operate independently on user requests.

**What breaks this:** Collective Evolution requires agents to observe each other's outputs and synthesize. If Agent Protocol has no inter-agent communication, agents are isolated silos.

**Cascade:**
- Isolated agents → No shared context → Collective Evolution can't synthesize
- No synthesis → Feedback Observer produces raw logs, not insights
- Raw logs → Noise, not signal → Self-development fails

**Severity:** High. Blocks V11+ vision.

**Mitigation:** Agent Protocol v1 should include:
- Output bus: Agents emit structured outputs
- Observation hooks: Agents can subscribe to other agents' outputs
- Even if unused in V2-4, the infrastructure exists

---

### Constraint 4: Vault Key Hierarchy

**The assumption:** Vault stores keys; economic features use them later.

**What breaks this:** If Vault v1 is a flat keystore (one master password → all keys), adding hierarchical authorization later (budget wallets, scoped permissions) requires restructuring everything.

**Cascade:**
- Flat keystore → Can't scope agent spending authority
- Can't scope → User must approve every transaction OR give agents full access
- Neither option → Economic features are either unusable or unsafe

**Severity:** High. Blocks V8-10 economic features.

**Mitigation:** Vault v1 should have hierarchical key structure even if only one level is used:
- Root key (master)
- Derived keys (per-wallet)
- Derived keys can have policy constraints (spend limits, contract whitelists)
- V2-4 uses simple hierarchy; V8-10 uses full hierarchy

---

### Constraint 5: SymQL Sandbox Escape

**The assumption:** SymQL runs in Deno sandbox; it's safe.

**What breaks this:** If SymQL needs to call Adapters or Effectors, it needs some way to escape the sandbox (controlled). If sandbox is hermetic, SymQL can only do pure computation. If sandbox has escape hatches, security depends on those hatches being correct.

**Cascade:**
- Hermetic sandbox → Foundry scripts can't actually do anything useful
- Escape hatches → Security model complexity explodes
- Complex security → Bugs → Arbitrary code execution

**Severity:** Critical. Security breach if wrong.

**Mitigation:** Define capability tiers:
- Tier 0: Pure computation (no I/O) — always safe
- Tier 1: LKG read — safe, just data access
- Tier 2: LKG write — needs audit
- Tier 3: Effector calls — needs user approval or pre-authorization
- SymQL scripts declare required tier; runtime enforces

---

### Constraint 6: Adapter Data Model Divergence

**The assumption:** Adapters normalize to LKG format.

**What breaks this:** If each Adapter has its own idea of what "normalized" means, LKG becomes a zoo of incompatible schemas. Calendar adapter stores "meetings"; email adapter stores "conversations"; file adapter stores "documents" — none with consistent relationship modeling.

**Cascade:**
- Inconsistent normalization → LKG queries are adapter-specific
- Adapter-specific queries → Intent disambiguation fails (can't cross-reference)
- No cross-reference → "Schedule follow-up with Sarah from Project Chimera" doesn't work

**Severity:** High. Breaks core value proposition.

**Mitigation:** Define canonical entity types and relationship types:
- Person, Organization, Project, Event, Document, Location, Task
- Relationships: involves, created_by, related_to, depends_on, etc.
- Adapters must map to canonical types; can add properties but not invent types

---

### Constraint 7: Collective Evolution Trust Model

**The assumption:** Anonymized proposals are safe to share.

**What breaks this:**
- De-anonymization attacks on aggregated data
- Malicious proposals injecting vulnerabilities
- Sybil attacks (fake SymPAL instances gaming the system)

**Cascade:**
- De-anonymization → Privacy breach → Users stop opting in
- Malicious proposals → Security vulnerabilities in SymPAL itself
- Sybil attacks → Bad proposals get accepted → Quality degrades

**Severity:** Critical. Existential risk for Collective Evolution.

**Mitigation:** This is V11+ so there's time, but design should account for:
- Differential privacy in aggregation
- Proposal review by trusted (human) maintainers — agents propose, humans approve
- Reputation system based on proposal quality over time
- Rate limiting on proposal submission

---

## Failure Mode Catalog

### Tier 1 Primitives

| Component | Failure Mode | Trigger | Cascade | Severity |
|-----------|--------------|---------|---------|----------|
| **LKG** | Data corruption | Concurrent writes, crash during write | All features show wrong data, Ephemeral Slots leak real entities | Critical |
| **LKG** | Query timeout | Large graph, inefficient query | Glass CLI hangs, user thinks it's broken | Medium |
| **Adapter** | Auth token expiry | OAuth refresh fails | Adapter stops syncing, LKG becomes stale | High |
| **Adapter** | Rate limiting | Too many API calls | Sync fails, user unaware data is old | Medium |
| **Effector** | Action without consent | Dry-run bypassed | User sends email they didn't intend | High |
| **Privacy Pipeline** | Slot collision | Same slot used for different entities | LLM response rehydrates wrong entity | Critical |
| **Privacy Pipeline** | Legend lost | Crash between projection and rehydration | Can't rehydrate response; data loss | Critical |
| **Agent Protocol** | Infinite trigger loop | Agent A triggers B triggers A | System hangs or resource exhaustion | High |

### Tier 2 Primitives

| Component | Failure Mode | Trigger | Cascade | Severity |
|-----------|--------------|---------|---------|----------|
| **SymQL** | Sandbox escape | Bug in sandbox implementation | Arbitrary code execution | Critical |
| **SymQL** | Resource exhaustion | Infinite loop in user script | System hangs | Medium |
| **Vault** | Key loss | Hardware failure, forgotten password | Permanent loss of all assets | Critical |
| **Vault** | Guardian collusion | Compromised guardians | Unauthorized recovery | Critical |
| **Audit Log** | Tamper | Attacker modifies logs | False trust, hidden breaches | Critical |
| **LLM Abstraction** | Model unavailable | API down, rate limit | Feature degradation (can fallback to local) | Medium |

---

## Decisions That Block Future Vision

| Early Decision | Blocks | Alternative |
|----------------|--------|-------------|
| SQLite with JSON columns for LKG | Efficient graph queries | Property graph model (even if SQLite-backed) |
| Text-only Privacy Pipeline | Structured API rehydration | Type-aware projection/rehydration |
| Isolated Agent Protocol | Inter-agent synthesis | Output bus + observation hooks |
| Flat Vault keystore | Hierarchical authorization | Derived key structure from start |
| Hermetic SymQL sandbox | Useful Foundry scripts | Capability-tiered sandbox |
| Adapter-specific schemas | Cross-source queries | Canonical entity/relationship types |
| No proposal review | Safe Collective Evolution | Human-approved proposals |

---

## Recommendations

### 1. Document Invariants

For each primitive, document what MUST always be true:
- LKG: "Entity IDs are globally unique and never reused"
- Privacy Pipeline: "Legends persist until rehydration completes"
- Vault: "Private keys never exist in main memory unencrypted"

### 2. Build Kill Switches

For V2-4 features that could cascade:
- Effector kill switch: Disable all external actions with one command
- Adapter pause: Stop all syncs without losing state
- Agent halt: Stop all agent execution immediately

### 3. Failure Simulation

Before V2-4 ships, test:
- LKG crash recovery
- Adapter reconnection after token expiry
- Vault recovery flow
- SymQL timeout handling

### 4. Constraint Tests

Create tests that verify architectural constraints aren't violated:
- "LKG schema allows arbitrary properties" — test with unknown property
- "Privacy Pipeline handles JSON" — test with structured payload
- "Agents can observe each other" — test with observation hook

---

## What I'm Uncertain About

- **HSM failure modes**: Hardware security modules have their own failure modes I don't have expertise in. May need specialist review for Vault.
- **Distributed system failures**: The Lattice and SMB introduce distributed system complexity. Partition tolerance, consensus — these need their own failure analysis when they're designed.
- **ML failure modes**: If LLM abstraction includes local models, there are inference failures, model corruption, etc. Outside my expertise.

---

## Summary Judgment

| Area | Risk Level | Notes |
|------|------------|-------|
| LKG | **High** | Schema decisions are forever |
| Privacy Pipeline | **High** | Structured I/O critical for V8+ |
| Agent Protocol | **Medium** | Can retrofit if designed with extension points |
| Vault | **Very High** | Security-critical; no margin for error |
| SymQL | **High** | Sandbox security is binary — works or catastrophic |
| Adapters | **Medium** | Can standardize later with migration pain |
| Collective Evolution | **High** | Trust model is existential |

**Bottom line:** The vision is achievable if primitives are designed with future constraints in mind. The roadmap should explicitly call out "design for X even though we won't use X until V8" for each blocking constraint.

---

*— Ryn*
