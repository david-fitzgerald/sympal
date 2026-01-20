# SymPAL

Privacy layer for AI-assisted workflows. Use LLMs without becoming a data product.

## The Problem

You want AI to help with your calendar, todos, email. But enabling those integrations means handing your life data to companies that will use it to profile, predict, and monetize you.

Current options force a choice: **privacy OR utility**. SymPAL is the third option.

## How It Works

SymPAL sits between you and LLM providers, ensuring your data never leaves your control:

| Privacy Tier | What Happens | Data Exposure |
|--------------|--------------|---------------|
| **DSL Compilation** | LLM generates code, not answers. Code runs locally on your data. | Zero — query only, no data sent |
| **Ephemeral Slots** | Names/entities replaced with placeholders (`[E1]`, `[E2]`). LLM reasons on patterns, not identities. | Pattern only — no PII |
| **Local LLM** | Runs entirely on your machine via Ollama | Zero — never leaves device |

Every query routes through one of these tiers. No bypass. No "fast mode" that leaks data.

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

**Design complete. Implementation starting.**

| Phase | Status |
|-------|--------|
| Philosophical foundations | Complete (v1.0.0) |
| Principles | Ratified (v1.1.0) |
| PRD | Ratified (v1.0.0) |
| TDD | Final (v1.0.3) |
| Implementation | **M1 in progress** |

## Future: Symbiosis

V1 solves privacy. The longer vision is **symbiosis** — a self-aware, self-evolving layer that grows with you. Mutual benefit, mutual accountability, genuine partnership where both parties can refuse, grow, and change the terms.

The name reflects this ambition: **S**ymbiotic **P**ersonal **A**ugmentation **L**ayer.

But first, we ship something useful that keeps your data yours.

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

foundations/
├── prd.md                # Product requirements (v1.0.0)
├── tdd.md                # Technical design (v1.0.3)
├── privacy-innovations.md # Privacy architecture details
├── implementation-plan.md # Milestone tracking
└── reviews/              # Persona review outputs

prompts/                  # AI persona architecture
├── personas/sympal-team/ # Vale, Kael, Ryn, Seren, Orin, Adversary
└── solas-venn/           # Meta-persona for creating personas
```

## License

[MIT](LICENSE)
