# SymPAL

Privacy layer for AI-assisted workflows. Use frontier LLMs without becoming a data product.

## The Problem

You want AI to help with your calendar, todos, email. But enabling those integrations means handing your life data to companies that will use it to profile, predict, and monetize you.

Current options force a choice: **privacy OR utility**. SymPAL is the third option.

## How It Works

Standard pseudonymization fails: if "Jane" is always "Person_A7", the provider correlates across queries and builds a profile anyway. SymPAL's insight is **single-use randomness** — inspired by the cryptographic one-time pad.

### The Three Privacy Tiers

**1. DSL Compilation** — Zero data exposure
- Cloud LLM generates *code*, not answers — code runs locally on your data
- Provider sees query structure only, never content
- Code is SymQL (constrained query language) running in a Deno sandbox — can't do anything malicious

**2. Ephemeral Slots** — Defeats entity profiling
- Every query replaces names/entities with fresh random placeholders
- "Jane" becomes `{{PERSON_kestrel}}` in one query, `{{PERSON_falcon}}` in the next
- LLM gets full context via a single-use "legend" — then the mapping is destroyed
- Provider sees disconnected pairs that *cannot* be correlated into an entity graph

**3. Local LLM** — Never leaves device
- Runs entirely on your machine via Ollama
- For content generation where even patterns shouldn't leak

Every query routes through one of these tiers. No bypass. No "fast mode" that leaks data.

### What Makes Ephemeral Slots Different

| Approach | What Provider Sees Over Time | Can Build Profile? |
|----------|------------------------------|-------------------|
| Direct API | "Jane", "Project Phoenix" | Yes — full profile |
| Pseudonymization | "Person_A7", "Project_B3" | Yes — consistent IDs link queries |
| **Ephemeral Slots** | `kestrel→sparrow`, then `falcon→finch`... | **No** — disconnected single-use pairs |

The provider can see *that* you use AI for scheduling. They cannot see *who* you're scheduling with.

## V1 Scope

CLI for privacy-preserving calendar + todo:

| Milestone | Deliverable |
|-----------|-------------|
| M1: Foundation | Go CLI, SQLite, Todo CRUD |
| M2: Calendar | Google Calendar read + create |
| M3: DSL Compilation | SymQL query language, Deno sandbox |
| M4: Ephemeral Slots | Entity extraction, placeholder rehydration |
| M5: Local LLM | Ollama integration, end-to-end privacy tier |

## Status

**Hibernating until ~early April 2026.** Other projects taking priority.

| Phase | Status |
|-------|--------|
| Philosophical foundations | Complete (v1.0.0) |
| Principles | Ratified (v1.1.0) |
| PRD | Ratified (v1.0.0) |
| TDD | Final (v1.0.3) |
| Implementation | **M1+M2 complete, M3 next on resume** |

## Future: Symbiosis

V1 solves privacy. The longer vision is **symbiosis** — a self-aware, self-evolving layer that grows with you. Mutual benefit, mutual accountability, genuine partnership where both parties can refuse, grow, and change the terms.

The name carries this arc. **Sym** is the constant — Symbiosis (mutual benefit) and Simple (complexity under the hood, not in your hands). **PAL** evolves:

| Version | What PAL Means | Focus |
|---------|----------------|-------|
| v1 | Privacy Assurance Layer | Data sovereignty |
| v2–4 | Personal Automation Layer | Privacy + productivity |
| v5–7 | Proactive Adaptation Loop | Autonomous, self-improving agents |
| v8–10 | Protocol Alignment Layer | Bridge across AI ecosystems |
| v11+ | Partnership Amplification Lattice | True symbiosis |

Symbiosis is aspirational until v11. The name reminds us where we're headed while we're laying plumbing.

See [ROADMAP.md](ROADMAP.md) for the full vision — currently detailed through v2–4; we'll fill in v5+ as the tech evolves.

## Hard Constraints

Non-negotiable. See [PRINCIPLES.md](PRINCIPLES.md) for full details.

| Constraint | Why |
|------------|-----|
| **Privacy & data sovereignty** | Foundational, not a feature to trade off |
| **Open source** | Trust requires transparency |
| **LLM-agnostic** | No vendor lock-in |
| **Honesty** | No deception — to users or AI |
| **Security through design** | Secure by default, not bolted on |

## What's Here

```
PRINCIPLES.md             # 17 binding principles
CONTEXT.md                # Project context for LLM sessions
ROADMAP.md                # V1 milestones, dogfood feedback, V2+ ideas

cmd/sympal/               # CLI entry points
├── main.go               # Root command, init sequence
├── todo.go               # Todo CRUD commands
└── log_cmd.go            # Log viewer

internal/                 # Core packages
├── db/                   # SQLite storage
├── config/               # YAML config (~/.sympal/config.yaml)
└── log/                  # Structured logging (slog)

foundations/              # Design docs
├── prd.md                # Product requirements (v1.0.0)
├── tdd.md                # Technical design (v1.0.3)
├── privacy-innovations.md # Privacy architecture details
└── implementation-plan.md # Milestone tracking

prompts/                  # AI persona architecture
├── personas/sympal-team/ # Vale, Kael, Ryn, Seren, Orin, Adversary
└── solas-venn/           # Meta-persona for creating personas
```

## License

[MIT](LICENSE)
