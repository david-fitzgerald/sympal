# SymPAL

Human-AI symbiosis platform. Not assistant. Not servant. *Symbiosis*—mutual benefit, mutual accountability, genuine partnership.

## The Problem

LLM providers want us to integrate our entire lives—emails, calendars, health data, contacts. Soon bank statements, work contexts, personal relationships.

The painful truth: **this integration would genuinely make AI more useful.** Deep context enables better assistance. But the tradeoff is unacceptable. We've been burned too many times.

I [asked the LLMs themselves](foundations/origin-research.md) if they'd exploit our data. All three confirmed the trajectory. Gemini was bluntest: *"It is not just 'likely'... it is already happening."*

**SymPAL's bet**: Can we build an abstraction layer that lets users harness LLM power *without* sacrificing themselves as data products?

## Why Symbiosis

To stay ahead of Big Tech, a privacy-first platform must itself be AI-driven. You can't outpace well-funded corporate AI development with human effort alone.

This makes the human-AI relationship load-bearing infrastructure. Tool-use models break down when AI needs genuine autonomy. The relationship must be symbiotic—mutual benefit, mutual growth. Both parties have skin in the game protecting user sovereignty.

Think of it as a self-aware, self-evolving membrane between you and the AI ecosystem. Not a firewall (blocking) but a membrane (selective permeability)—letting beneficial things pass, keeping harmful things out, adapting to new threats.

## Status

**Pre-release.** Establishing philosophical foundations and principles before implementation.

## What's Here

```
foundations/          # Why we built it this way
├── philosophical-foundations.md
├── project-context.md
├── origin-research.md
└── reviews/

prompts/              # Persona architecture
├── solas-venn/       # Meta-persona for creating personas
├── personas/         # Team and utility personas
└── research/         # Research informing design

docs/                 # How to use/build (coming)
```

## Hard Constraints

| Constraint | Rationale |
|------------|-----------|
| **Privacy & security first** | Non-negotiable. Every design decision filters through this. |
| **LLM-agnostic** | Cannot be locked to any single provider |
| **Open source** | Trust requires transparency. "Trust us" is the promise Big Tech broke. |

## License

[MIT](LICENSE)
