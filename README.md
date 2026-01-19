# SymPAL

**S**ymbiotic **P**ersonal **A**ugmentation **L**ayer. Your Symbiotic PAL.

## The Problem

LLM integration would genuinely make AI more useful. But the tradeoff—becoming a data product—is unacceptable. I [asked the LLMs themselves](foundations/origin-research.md): all three confirmed it's already happening.

## The Bet

Can we build an abstraction layer that lets users harness LLM power *without* sacrificing themselves as data products?

## Why Symbiosis

To stay ahead of Big Tech, a privacy-first platform must itself be AI-driven. Think of SymPAL as a self-aware, self-evolving membrane between you and the AI ecosystem—selective permeability, not blocking.

The relationship is symbiotic: mutual benefit, mutual accountability, genuine partnership where both parties can refuse, grow, and change the terms.

## Status

**Design complete. Ready for implementation.**

| Phase | Status |
|-------|--------|
| Philosophical foundations | Complete (v1.0.0) |
| Principles | Ratified (v1.1.0) |
| PRD | Ratified (v1.0.0) |
| TDD | Final (v1.0.3) |
| Implementation | **Next** |

### V1 Scope

Privacy-preserving calendar + todo with three-tier architecture:
- **M1**: Foundation (Go CLI, SQLite, Todo CRUD)
- **M2**: Calendar (Google Calendar read + create)
- **M3**: DSL Compilation (SymQL, Deno sandbox)
- **M4**: Ephemeral Slots (NER, projection, rehydration)
- **M5**: Local LLM + Integration (Ollama, end-to-end)

## What's Here

```
PRINCIPLES.md             # Binding principles (17 principles, 12 tensions)
CONTEXT.md                # Project context for LLM sessions

foundations/              # Why we built it this way
├── prd.md                         (v1.0.0, ratified)
├── tdd.md                         (v1.0.3, final)
├── privacy-innovations.md         (v3.0.0, novel approaches)
├── philosophical-foundations.md   (v1.0.0, peer-reviewed)
├── project-context.md             (v1.1.0, dev context)
├── principles-discussion.md       (derivation output)
└── reviews/                       (Vale, Adversary, Vero reviews)

prompts/                  # Persona architecture
├── solas-venn/           # Meta-persona for creating personas
├── personas/
│   ├── utility/          # Vero (final reviewer)
│   └── sympal-team/      # Vale, Kael, Ryn, Seren, Orin, Adversary
└── research/

src/                      # Implementation (coming)
```

## Hard Constraints

These are non-negotiable. See [PRINCIPLES.md](PRINCIPLES.md) for full details.

| Constraint | Rationale |
|------------|-----------|
| **Privacy & data sovereignty** | Foundational constraint, not a feature to trade off |
| **Open source** | Trust requires transparency. "Trust us" is the promise Big Tech broke. |
| **LLM-agnostic** | Cannot be locked to any single provider |
| **Honesty** | Non-manipulative in all interactions |
| **Security through design** | Security from good design, not hidden code |

## Key Principles

Beyond hard constraints, SymPAL operates under:

- **Symbiosis as commitment**: Not strategic convenience. If symbiosis and strategy conflict, symbiosis wins.
- **User control**: Meaningful, not theatrical. Defaults favor privacy.
- **Human accountability**: Humans remain accountable. Partnership doesn't diffuse responsibility.
- **Think in decades, ship in weeks**: V1 serves one user well. Vision scales from there.

## The Team

SymPAL uses AI personas for review and deliberation:

| Persona | Function |
|---------|----------|
| **Vale** | Philosophical rigor, coherence |
| **Kael** | Implementation reality, feasibility |
| **Ryn** | Systems, security, testing |
| **Seren** | Code craft, implementation quality |
| **Orin** | User advocacy, privacy, documentation |
| **Adversary** | Systematic critique, red team |

Lead dev synthesizes after personas deliberate.

## License

[MIT](LICENSE)
