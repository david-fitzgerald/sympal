# Implementation Plan

**Version:** 0.1.0
**Date:** 2026-01-20
**Status:** Active
**Purpose:** Session-start reference for implementation phase

---

## Session Start Checklist

1. Read this file
2. Check current milestone progress (below)
3. Review any open GitHub issues
4. Pick up where we left off

---

## Current Status

**Phase:** M2 Calendar Integration
**Progress:** In progress (~30%)

**M1 Foundation:** ✅ Complete (2026-01-21)

### M2 Progress Detail

| Component | Status | Notes |
|-----------|--------|-------|
| `internal/keyring/` | ✅ Complete | Token save/load via system keychain |
| `internal/config/` | ✅ Complete | GoogleConfig struct added |
| `internal/auth/google.go` | 🔲 Skeleton | URL builder only, needs: state gen, callback server, token exchange |
| Calendar API client | 🔲 Not started | — |
| `sympal auth` command | 🔲 Not started | — |
| `sympal today` command | 🔲 Not started | — |

**Next steps:**
1. Implement secure state generation (crypto/rand)
2. Build callback HTTP server (localhost:8080)
3. Token exchange with Google
4. Wire up keyring storage
5. Add CLI commands

---

## Milestone Wrap-Up Procedure

After completing each milestone:

### 1. Manual Testing (All Milestones)
- Run through all new commands/features
- Test edge cases (empty inputs, invalid IDs, missing files)
- Verify nothing broke from previous milestones

### 2. Code Review (Security-Critical Milestones)
Use personas for M3+ (sandboxing, OAuth, LLM integration):

| Persona | Focus |
|---------|-------|
| **Ryn** | Security review — "how will this fail?" |
| **Seren** | Code craft — "is this well-crafted?" |
| **Kael** | Implementation — "will this survive reality?" |

**Not needed for:** M1, M2 (low security surface)

### 3. Dogfooding (Starts M1)
- Add sympal to PATH: `export PATH="$PATH:/Users/df/pg/sympal"`
- Use daily for real tasks
- Log friction points for future improvement

### 4. Documentation Update
- Update this file (Current Status section)
- Update CONTEXT.md if project status changed
- Git tag: `git tag v0.x.0-mN`

### 5. Vero Review
**Not for code** — Vero reviews foundational docs only. Code review uses team personas above.

---

## Learning-First Approach

**Goal:** Build coding confidence through real implementation, not just receive code.

**Skill baseline:** Intermediate prompt architecture, basic coding. Moving toward confident independent writing.

### How We Work

| Mode | Trigger | Style |
|------|---------|-------|
| **Teach** (default) | New concepts, architecture decisions, anything you'll debug later | Explain *why* before code, flag transferable patterns, you write/modify |
| **Ship** | Boilerplate, well-trodden patterns, "just do it" signal | I write, you review, move fast |
| **Deep dive** | Security-critical, novel patterns, genuine curiosity | Go deep with sources, build full understanding |

### The Tension to Watch

> **Learning-as-excuse**: Planning serves learning even when shipping serves project.

**Legitimate deep dives:**
- Security-critical (sandboxing, OAuth token handling)
- Architecture decisions with tradeoffs
- Patterns you'll use repeatedly

**Probably learning-as-excuse:**
- Fully understanding X before writing simple code that uses X
- Researching edge cases for MVP features
- Perfecting understanding when "good enough" would ship

**When in doubt:** Ask "What would we learn by building instead?"

---

## Milestones

### M1: Foundation ✅

**Status:** Complete (2026-01-21)

**Deliverables:**
- [x] Go project scaffolding (Cobra CLI)
- [x] SQLite setup with schema
- [x] Todo CRUD (`sympal todo add/list/done/delete`)
- [x] Config file handling (`~/.sympal/config.yaml`)
- [x] Logging infrastructure (`~/.sympal/sympal.log`)
- [x] `sympal log` command (view recent queries)

**Gate:** ✅ Todo CRUD works end-to-end

**Learning outcomes:**
- Go fundamentals (syntax, error handling, packages)
- CLI structure patterns (Cobra)
- Project scaffolding (transferable to any Go project)
- SQLite basics
- Config/logging patterns (slog for structured logging)

---

### M2: Calendar Integration (Current)

**Deliverables:**
- [ ] Google OAuth flow (system keychain storage)
- [ ] Calendar API integration (read events)
- [ ] `sympal today` command (todos + calendar)
- [ ] Basic day view (no LLM yet)

**Gate:** Can view today's calendar and todos together

**Learning focus:**
- OAuth flow (reusable for any Google/OAuth API)
- API integration patterns
- Token management (refresh, storage)
- Graceful degradation

---

### M3: DSL Compilation

**Deliverables:**
- [ ] Query Classifier (keyword cascade)
- [ ] Schema description generator
- [ ] Claude API integration
- [ ] SymQL code generation
- [ ] Deno sandbox setup
- [ ] Code validation and execution pipeline

**Gate:** >90% structured queries return correct results

**Learning focus:**
- Code generation patterns
- Sandboxing (security-critical — go deep)
- LLM API integration
- Compile-don't-interpret pattern

---

### M4: Ephemeral Slots

**Deliverables:**
- [ ] Entity extraction (NER)
- [ ] Ephemeral placeholder generation
- [ ] Legend construction (task-based defaults)
- [ ] Projection function (real → placeholders)
- [ ] Rehydration function (response → real)

**Gate:** >95% rehydration accuracy

**Learning focus:**
- NER/text processing
- State management
- Projection/rehydration pattern (novel privacy technique)

---

### M5: Local LLM + Integration

**Deliverables:**
- [ ] Ollama integration
- [ ] Content task routing
- [ ] End-to-end privacy tier
- [ ] Quality logging

**Gate:** Daily use + ≥50% queries via DSL/Ephemeral Slots

**Learning focus:**
- Local LLM orchestration
- Multi-model routing
- Quality/privacy tradeoffs

---

## Technical Stack

| Tool | Purpose |
|------|---------|
| Go | Main language |
| SQLite | Local data storage |
| Deno | Sandbox for generated code |
| Ollama + Llama 3.2 3B | Local LLM |
| Claude API | Cloud LLM |
| Google Calendar API | Calendar integration |

---

## Team Usage During Implementation

| Persona | When to Use |
|---------|-------------|
| **Kael** (implementation) | Architecture decisions, "will this survive reality?" |
| **Ryn** (systems/security) | Security review, failure modes, "how will this fail?" |
| **Seren** (code craft) | Code review before major commits, "is this well-crafted?" |
| **Orin** (user advocacy) | UX decisions, "are users better off?" |
| **Adversary** | Challenge assumptions, red-team security |

**Not for:** Every small change, creating code (that's us).

---

## Working Patterns

- Commit after completing tasks
- Push after each commit
- End sessions with everything committed and pushed
- Update "Current Status" section when milestone progress changes

---

## Dogfood Feedback

See [ROADMAP.md](/ROADMAP.md#dogfood-feedback) — single source of truth for friction points and future ideas.

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 0.1.0 | 2026-01-20 | Initial plan created |
| 0.2.0 | 2026-01-21 | M1 complete, added Milestone Wrap-Up Procedure |
| 0.3.0 | 2026-01-24 | M2 in progress (~30%): keyring + config complete, auth skeleton |
