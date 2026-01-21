# Roadmap Consolidation — Team Synthesis

**Date:** 2026-01-22
**Reviews:** Kael, Ryn, Vale, Orin, Adversary, Vero

---

## Consensus Findings

### 1. Core Primitives Identified

The team agrees on 8-9 foundational primitives. Everything else is built on these:

| Tier | Primitive | Purpose | Review Coverage |
|------|-----------|---------|-----------------|
| **1** | LKG (Local Knowledge Graph) | Data spine | Kael, Ryn |
| **1** | Privacy Pipeline | Ephemeral Slots projection/rehydration | Kael, Ryn, Adversary |
| **1** | Adapter Framework | Data ingestion | Kael, Ryn |
| **1** | Effector Framework | Actions/outputs | Kael, Ryn |
| **1** | Agent Protocol | Internal agent coordination | Kael, Ryn |
| **2** | SymQL Runtime | DSL + sandbox | Kael, Ryn |
| **2** | Vault Core | Key storage | Kael, Ryn, Orin |
| **2** | Audit Subsystem | Append-only log | Kael |
| **2** | LLM Abstraction Layer | Model-agnostic interface | Kael |

**Adversary challenge:** Is LKG really the spine, or is Privacy Pipeline the philosophical spine? Unresolved.

### 2. Vault Scope Confusion

All reviewers flagged Vault as overscoped for V2-4:

| Reviewer | Position |
|----------|----------|
| Kael | Vault spec is V2-V8 work compressed into V2-4 |
| Ryn | Vault key hierarchy critical; flat keystore blocks V8+ |
| Orin | Duress protocol, dead man's switch are bloat |
| Vero | Vault spec in V2-4 section contradicts placement |
| Adversary | Confusion may be labeling, not scope — spec can be aspirational |

**Resolution:** Label Vault spec as full design; mark V2-4 deliverable as keychain integration only.

### 3. Architectural Constraints That Block Future

Ryn identified 7 constraints; Kael confirmed via dependency analysis:

| Constraint | Risk | Mitigation |
|------------|------|------------|
| LKG schema rigidity | Blocks V5+ proactive features | Property graph from day 1 |
| Privacy Pipeline text-only | Blocks V8+ SaaS rehydration | Structured I/O from start |
| Agent Protocol isolation | Blocks Collective Evolution | Output bus + observation hooks |
| Vault flat keystore | Blocks hierarchical authorization | Derived key structure |
| SymQL hermetic sandbox | Foundry useless | Capability-tiered sandbox |
| Adapter schema divergence | Cross-source queries fail | Canonical entity types |
| Collective Evolution trust | Existential risk | Human-approved proposals |

**Consensus:** Design for V8+ constraints now, even if not used until later.

### 4. V1-4 Strong, V5+ Uncertain

| Version | Value Assessment | Confidence |
|---------|------------------|------------|
| V1 (M1-M5) | High — solid, achievable | High |
| V2-4 | Medium-High — good ideas, needs prioritization | Medium |
| V5-7 | Medium — proactive features are risky | Low |
| V8-10 | Low — platform features, not user features | Low |
| V11+ | Unknown — depends on AI capabilities that don't exist | Very Low |

**Adversary challenge:** Dismissing V5+ as "speculative" while using speculative user research is circular. The real issue is epistemic status, not value.

---

## Contested Findings

### 1. CLI-First Bias

| Position | Held By |
|----------|---------|
| CLI is wrong priority; Lens/Mobile should be primary | Orin |
| CLI-first is appropriate for dogfooding phase | Adversary |

**Vero assessment:** Both are correct for different timeframes. V2-4 = dogfooding = CLI appropriate. Post-V4 = users = Lens/Mobile.

### 2. Privacy vs. Collective Learning Tension

| Position | Held By |
|----------|---------|
| This is an unresolved tension that needs explicit navigation | Vale |
| This is a resolved design decision (anonymization + consent) | Adversary |

**Vero assessment:** The tension is real but may be navigable. Needs explicit statement in CONTEXT.md 12 tensions.

### 3. "Agents Propose Only" Permanence

| Position | Held By |
|----------|---------|
| This may conflict with "self-developing" vision | Vale |
| Needs clarification: permanent constraint or V1-V10 guardrail? | Vale |

**Vero assessment:** BLOCKING — Cannot assess Collective Evolution feasibility without resolving.

---

## Meta-Findings

### 1. Roadmap Purpose Confusion (Vero)

The roadmap conflates three functions:
1. Implementation plan (M1-M5)
2. Architecture vision (primitives)
3. Possibility space (V11+)

**Recommendation:** Split or clearly label sections by commitment level.

### 2. Overplanning Risk (Adversary)

The team converged on "more specs needed." But:
- Lead dev has known overplanning bias
- Primitives can be specced via prototyping
- 590-line roadmap may already be over-planned

**Adversary's uncomfortable question:** Is the roadmap productive procrastination?

### 3. Resource Constraint Not Acknowledged (Vero)

The roadmap describes a multi-developer system. The team is one person with basic coding skills + AI assistance. Nobody addressed this.

### 4. Completion Criteria Missing (Vero)

No definition of "V1 done" or "V2 entry criteria." Risk: perpetual V1.

---

## Recommended Actions

### Immediate (Before Next Session)

1. **Label roadmap sections by commitment level**
   - COMMITTED: V1 (M1-M5)
   - PLANNED: V2-4
   - EXPLORATORY: V5+

2. **Add V1 completion criteria**
   - When is M5 "done"?
   - What triggers V2 work?

3. **Resolve "agents propose only" vs. "self-developing"**
   - Make a decision and document it

### Short-Term (V2-4 Planning)

4. **Create primitive interface specs** (but start minimal)
   - LKG: Schema, query API
   - Privacy Pipeline: Structured I/O interface
   - Agent Protocol: Trigger/output spec with observation hooks

5. **Re-label Vault**
   - Full spec = design vision
   - V2-4 scope = keychain integration only

6. **Add privacy/collective tension to CONTEXT.md 12 tensions**

### Deferred (V5+ Planning)

7. **Define Collective Evolution trust model**
   - How do proposals get reviewed?
   - How is anonymization verified?
   - What prevents Sybil attacks?

8. **Add uncertainty markers to V8+**
   - Currently reads as commitment, should read as possibility

---

## Synthesis Recommendation

**Vero verdict: REVISE before treating roadmap as implementation guide.**

The roadmap is valuable as a vision document. It should not guide V2+ planning until:
1. Commitment levels are explicit
2. Completion criteria exist
3. Key contradictions are resolved

**Adversary's final challenge stands:**

> "If you could only build V1, and V2+ never happened, would you still build SymPAL?"

If yes → V1 is sufficient scope. Focus there.
If no → The project needs V2+ to be meaningful. That's important to know.

---

## Files Generated

| File | Persona | Key Finding |
|------|---------|-------------|
| `01-kael-review.md` | Kael | 8 primitives; foggy build path; Vault overscoped |
| `02-ryn-review.md` | Ryn | 7 critical constraints; failure cascade analysis |
| `03-vale-review.md` | Vale | 7/10 coherence; privacy/collective tension |
| `04-orin-review.md` | Orin | V1-4 mixed value; CLI bias; V8+ speculative |
| `05-adversary-review.md` | Adversary | Team convergence challenged; meta-questions raised |
| `06-vero-review.md` | Vero | REVISE recommendation; structural issues |

---

*Synthesis complete. Lead dev decision required on actions.*
