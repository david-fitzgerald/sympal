# Kael — Roadmap Implementation Review

**Date:** 2026-01-22
**Focus:** Primitive extraction, dependency mapping, build path clarity
**Question:** "Will this survive reality? What are the actual building blocks?"

---

## Executive Summary

You've got a 590-line vision document masquerading as a roadmap. The vision is coherent, but the implementation path is foggy. I count roughly 40+ distinct features across V1-V11+, many with unstated dependencies on components that don't exist yet.

**The core problem:** Features are organized by *when* (version), not by *what enables what*. This makes it hard to see which decisions block others.

---

## Primitive Extraction

After cutting through the feature descriptions, here are the actual building blocks everything depends on:

### Tier 1: Foundational Primitives (Must exist first)

| Primitive | What It Is | Everything That Depends On It |
|-----------|------------|-------------------------------|
| **LKG (Local Knowledge Graph)** | Graph database of entities + relationships | Intent disambiguation, Ephemeral Slots, Lens UI, proactive features, Collective Evolution feedback, basically everything V2+ |
| **Adapter Framework** | Pluggable modules for data ingestion | All integrations (calendar, email, files, wearables, SaaS), LKG population |
| **Effector Framework** | Pluggable modules for actions | All outputs (notifications, smart home, API calls), Dry-Run Framework |
| **Agent Protocol** | Triggers, inputs, outputs, constraints spec | All internal agents, external agent hiring, Collective Evolution |
| **Privacy Pipeline** | Projection → Slots → Rehydration | Ephemeral Slots, cloud LLM usage, Collective Evolution anonymization |

### Tier 2: Platform Primitives (Enable major capabilities)

| Primitive | What It Is | Depends On | Enables |
|-----------|------------|------------|---------|
| **SymQL Runtime** | DSL interpreter + sandbox | LKG schema | Foundry, automation, community recipes |
| **Vault Core** | HSM integration + key storage | — | Economic features, credential orchestration, signing |
| **Audit Subsystem** | Append-only log + verifiable builds | — | Trust features, Collective Evolution proposals |
| **LLM Abstraction Layer** | Unified interface to multiple models | — | LLM-agnosticism, cost routing, cross-model orchestration |

### Tier 3: Derived Capabilities (Built on primitives)

Everything else — Lens UI, Workbench, goal inference, SMB, economic features, The Lattice — these are *applications* of the primitives, not primitives themselves.

---

## Dependency Graph (Simplified)

```
                    ┌─────────────────┐
                    │   LKG Schema    │
                    └────────┬────────┘
                             │
         ┌───────────────────┼───────────────────┐
         │                   │                   │
         ▼                   ▼                   ▼
┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
│ Adapter Framework│ │  Privacy Pipeline│ │  SymQL Runtime  │
└────────┬────────┘ └────────┬────────┘ └────────┬────────┘
         │                   │                   │
         │                   ▼                   │
         │          ┌─────────────────┐          │
         └─────────►│ Agent Protocol  │◄─────────┘
                    └────────┬────────┘
                             │
         ┌───────────────────┼───────────────────┐
         │                   │                   │
         ▼                   ▼                   ▼
   Internal Agents    External Agents    Collective Evolution
```

The **LKG** is the spine. Without a solid graph schema and query interface, nothing else works.

---

## Feasibility Assessment by Section

### V1 Milestones — Clear ✓

Build path visible. M1 done. M2-M5 are sequential and well-scoped. No concerns.

### V2-4 Infrastructure — Foggy

| Component | Dependency Clarity | Complexity | Build Path |
|-----------|-------------------|------------|------------|
| LKG | Partial | High | What's the schema? What's the query language? Is this SQLite with graph extensions, or a real graph DB? |
| Foundry | Partial | High | Requires SymQL spec, sandbox security model, versioning scheme. None defined. |
| Vault | Clear | Very High | Well-specified in roadmap, but this is months of security engineering. Is this V2-4 or V3-4? |
| Auto-Upgrade | Clear | Medium | Straightforward if Audit Subsystem exists. |

**My concern:** The Vault is specified at the detail level of a V5+ feature, but placed in V2-4. Either scope it down (keychain only, no Shamir, no dead man's switch) or push full Vault to V3-4.

### V2-4 UX Surfaces — Depends on Primitives

Glass CLI, Lens, Workbench, Companion — all depend on LKG and SymQL existing. These are *applications*, not primitives. Order should be: primitives first, then UX.

### V5-7 Features — Requires V2-4 Solid

Goal inference, workflow synthesis, proactive agents — all assume LKG is rich, Agent Protocol exists, and SymQL is mature. If V2-4 primitives are shaky, V5-7 won't work.

### V8-10 Economic Features — Requires Vault Mature

Agent delegation, micropayments, trusted registry — all assume Vault is battle-tested. The roadmap has Vault in V2-4 but economic features in V8-10. That's a 4+ version gap. Is the Vault really V2-4 scope, or is it being built incrementally across V2-V8?

### Collective Evolution — Requires Everything

This depends on: Agent Protocol, Privacy Pipeline (anonymization), Audit Subsystem (proposals), LKG (feedback storage), and external infrastructure (Beads, Git). It's correctly placed at V11+ for full realization, but the V2-4 "Feedback Observer" piece needs the Agent Protocol to exist.

---

## Architectural Constraints (Blocking Risks)

These decisions, if made wrong, could block future vision:

| Decision | Risk | Mitigation |
|----------|------|------------|
| **LKG Schema Design** | Wrong schema makes Ephemeral Slots projection hard | Design schema with projection in mind from day 1. Include entity types, relationship types, provenance metadata. |
| **SymQL as DSL** | If too limited, Foundry is useless. If too powerful, sandbox is impossible. | Define capability tiers. Tier 1: read-only queries. Tier 2: local mutations. Tier 3: external actions. Sandbox can restrict by tier. |
| **Agent Protocol Spec** | If agents can't observe each other's outputs, Collective Evolution can't synthesize. | Build inter-agent communication into protocol from start, even if unused in V2-4. |
| **Privacy Pipeline Architecture** | If projection is one-way (can't rehydrate structured responses), action-based workflows break. | Ensure rehydration handles not just text but structured API responses. |

---

## Recommendations

### 1. Reframe V2-4 Around Primitives

Instead of listing features, list primitives with clear specs:

```
V2-4 Primitives:
- LKG v1: Schema spec, SQLite+FTS backend, basic query API
- Adapter Framework v1: Interface spec, 3 reference adapters (calendar, files, notes)
- Effector Framework v1: Interface spec, 2 reference effectors (CLI, notifications)
- Privacy Pipeline v1: Entity extraction, slot projection, legend management
- Agent Protocol v1: Trigger/input/output spec, 2 reference agents (Vale, Adversary)
- SymQL v1: Grammar spec, interpreter, basic sandbox
- Vault v1: Keychain storage only (no Shamir, no economic features)
- Audit Subsystem v1: Append-only log, basic query
```

Then features (Lens, Glass CLI, Workbench) are explicitly "built on" these primitives.

### 2. Create Primitive Spec Documents

Each Tier 1 and Tier 2 primitive needs a 1-2 page spec in `/foundations/specs/`:
- Interface definition
- Data model
- Extension points
- What it explicitly does NOT do in v1

### 3. Vault Phasing

Current Vault spec is V2-V8 worth of work compressed into V2-4. Propose:
- **V2**: Keychain integration (Secure Enclave/TPM), basic key storage
- **V3**: Shamir recovery, guardian protocol
- **V4**: Transaction firewall, policy engine
- **V8+**: Economic features (micropayments, smart contracts)

### 4. Dependency-First Ordering

In each version section, list:
1. What primitives this version delivers
2. What primitives this version requires
3. What features are enabled by the new primitives

---

## What I Don't Know

- **Graph DB choice**: Is SQLite with JSON sufficient for LKG, or do we need something like DuckDB, SurrealDB, or embedded Neo4j? Complexity varies 10x depending on choice.
- **Sandbox technology**: Deno is mentioned for SymQL. Is that confirmed? Alternatives (WASM, gVisor) have different tradeoffs.
- **HSM integration complexity**: Secure Enclave and TPM have different APIs. Cross-platform Vault is harder than single-platform.

---

## Summary Judgment

| Dimension | Rating | Notes |
|-----------|--------|-------|
| Dependency clarity | **Partial** | Primitives identifiable but not explicitly called out |
| Complexity estimate | **High** | 8 primitives × months each = years of work |
| Resource alignment | **Gap** | Solo dev + basic coding skills vs. security-critical Vault |
| Technical debt | **Concerning** | If primitives aren't clean, everything built on them suffers |
| Build path | **Foggy** | Vision clear, implementation sequence unclear |

**Bottom line:** The roadmap has good bones. Reframe it around primitives, spec them properly, and the fog lifts.

---

*— Kael*
