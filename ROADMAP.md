# ROADMAP

> Living document for SymPAL direction. V1 milestones are committed scope; V2+ are ideas that may evolve.

---

## V1 Milestones

| Milestone | Status | Summary |
|-----------|--------|---------|
| M1: Foundation | ✅ Complete | Todo CRUD, config, logging |
| M2: Calendar | 🔲 Next | Google Calendar read + `sympal today` |
| M3: DSL Compilation | 🔲 Planned | SymQL, Deno sandbox, query classification |
| M4: Ephemeral Slots | 🔲 Planned | NER, projection/rehydration for privacy |
| M5: Local LLM | 🔲 Planned | Ollama integration, end-to-end privacy tier |

**V1 Gate:** Daily use by lead dev with ≥50% queries via DSL/Ephemeral Slots.

See `foundations/implementation-plan.md` for detailed deliverables.

---

## Dogfood Feedback

Friction captured during real use. Informs future iterations.

### Todo Commands

| Issue | Notes | Priority |
|-------|-------|----------|
| `sympal todo` verbose | Alias (`st`) or natural language input | Medium |
| IDs don't reset | Consider dynamic display IDs or fuzzy match | Low |
| No modify command | `sympal todo edit [id] [content]` | Medium |
| No subcategories | Tags/projects for task grouping | Low |
| Incomplete logging | Add logging to list/done/delete | Low |

### General

*None yet — add as discovered*

---

## V2+ Ideas

Features beyond V1 scope. May be refined, merged, or dropped.

| Idea | Notes |
|------|-------|
| Multi-calendar support | Outlook, CalDAV, Apple Calendar |
| Task dependencies | "Do X after Y" |
| Recurring tasks | Daily/weekly/custom patterns |
| Mobile companion | Read-only view or simple input |
| Shared contexts | Family calendar aggregation |
| Voice input | Whisper integration for quick capture |
| Plugin system | User-defined data sources |

---

## Contributing Ideas

Dogfood → note friction → add to this doc → revisit during milestone planning.

---

*Last updated: 2026-01-21*
