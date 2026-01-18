# TDD Extraction Notes

**Date:** 2026-01-18
**Interviewer:** Kael + Ryn
**Interviewee:** Lead Dev

---

## Principles Alignment

### P1: Privacy
- Local-only for raw data
- Explicit user control for transformed data leaving device
- No over-engineering (e.g., user-held encryption keys deferred)

### P3: LLM-Agnostic
- V1: Claude + local LLM (Ollama) only
- Full LLM-agnostic is V2+ goal
- Keep it simple for V1

### P5: Security by Design
- **OAuth tokens**: System keychain (macOS Keychain, Linux secret-service)
- **Sandbox**: Deno subprocess with deny-by-default permissions (required)
- **Encryption at rest**: Skip — rely on OS-level (FileVault/BitLocker)
- **CLI auth**: Skip — OS login is the security boundary

---

## Technical Stack

| Component | Decision | Rationale |
|-----------|----------|-----------|
| Main language | Go | Simple syntax, LLM-friendly, single binary, great CLI ecosystem |
| Sandbox runtime | Deno subprocess | Built-in deny-by-default permissions |
| Sandboxed code | TypeScript/JavaScript | Runs in Deno |
| Local LLM runtime | Ollama | Standard, well-supported |
| Local LLM model | Llama 3.2 3B | Fits 8GB RAM, fast, good enough for entity typing |
| Database | SQLite | Graph-like schema (entities + relationships) |
| Data location | `~/.sympal/` | Simple, obvious, platform-agnostic |

---

## Data Model

### Todo (Minimal)
- Due date (optional)
- Priority: High / Medium / Low
- Status: Not started / Done
- CRUD operations only

### Knowledge Graph
- Entities table (id, name, type, modifiers, metadata)
- Relationships table (source, target, relation_type, metadata)
- Enables semantic projection for privacy tier

---

## CLI UX

| Element | Decision | Rationale |
|---------|----------|-----------|
| Command style | Subcommands (`sympal <cmd>`) | Simplest to implement |
| Output | Contextual (echo what happened) | No decisions about what to omit |
| Memory | Session-only | No file persistence for V1 |
| Errors | Exit codes + stderr + contextual messages + `--verbose` | Best practice |

---

## Google API Scopes

| Integration | Access | Use Case |
|-------------|--------|----------|
| Calendar | Read + Write | View events, create from LLM |
| Gmail | Read only | Detect actionable emails → todos |
| Contacts | Read only | Context enrichment |

---

## Testing

| Area | Approach |
|------|----------|
| Sandbox | Unit tested (required — security boundary) |
| Privacy tier | Unit tested (required — hard to verify manually) |
| Knowledge graph | Unit tested (required — data integrity) |
| Google APIs | Smoke tests only |
| Other | Unit tests for learning; dogfooding for UX |

---

## Logging

| Element | Decision |
|---------|----------|
| Approach | Single logger, everything to file |
| File | `~/.sympal/sympal.log` |
| Screen | stdout (output), stderr (errors), `--verbose` (debug) |
| Format | `TIMESTAMP LEVEL MESSAGE` |

Rationale: Privacy architecture is complex; need visibility into routing, projection, rehydration for debugging.

---

## Ops

| Element | Decision |
|---------|----------|
| Config format | YAML (`~/.sympal/config.yaml`) |
| Distribution | Single Go binary, manual install |
| CI/CD | Skip for V1 dogfooding |
| Backup | User responsibility (`~/.sympal/` folder) |

---

## Key Discussions

### Why Go + Deno?
- Go: Simple, LLM-friendly, single binary, great CLI ecosystem (Cobra, Bubble Tea)
- Deno: Built-in sandbox with deny-by-default permissions
- Research showed this combo recommended for similar use cases
- Alternative languages (Zig, Nim, Crystal) not viable due to small training corpus for AI-assisted dev

### Why SQLite for Knowledge Graph?
- Need to store entities + relationships for semantic projection
- SQLite with graph-like schema (nodes/edges tables) sufficient for V1 scale
- Upgrade path to dedicated graph DB if needed
- JSON files too limited; full graph DB overkill

### Why Single Logger?
- Privacy architecture has many moving parts (classifier, compiler, projection, rehydration)
- Need to debug routing decisions, entity recognition, projection accuracy
- Simpler than separate loggers for different components
- Supports P12 (transparency) — can audit what was sent

---

*Extracted: 2026-01-18*
