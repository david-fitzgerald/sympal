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

**Phase:** M1 Foundation
**Progress:** Not started

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

### M1: Foundation (Current)

**Deliverables:**
- [ ] Go project scaffolding (Cobra CLI)
- [ ] SQLite setup with schema
- [ ] Todo CRUD (`sympal todo add/list/done/delete`)
- [ ] Config file handling (`~/.sympal/config.yaml`)
- [ ] Logging infrastructure (`~/.sympal/sympal.log`)
- [ ] `sympal log` command (view recent queries)

**Gate:** Todo CRUD works end-to-end

**Learning focus:**
- Go fundamentals (syntax, error handling, packages)
- CLI structure patterns (Cobra)
- Project scaffolding (transferable to any Go project)
- SQLite basics
- Config/logging patterns

**Approach:**
1. Scaffold together — explain structure, you create files
2. First operation (todo add) — write line by line, explain each part
3. Remaining CRUD — you draft, I review
4. By end: You can read any Go in codebase and roughly understand it

---

### M2: Calendar Integration

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

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 0.1.0 | 2026-01-20 | Initial plan created |
