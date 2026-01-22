# ROADMAP

> Living document for SymPAL direction. Sections marked by commitment level:
> - **COMMITTED**: This is what we're building
> - **PLANNED**: Intended direction, subject to learning
> - **EXPLORATORY**: Possibilities, not commitments

---

## The Long Arc: What SymPAL Becomes

The name carries the vision. **Sym** is the constant — Symbiosis (mutual benefit between human and AI) and Simple (complexity under the hood, not in your hands). **PAL** evolves as the project matures:

| Version | PAL Expansion | Focus |
|---------|---------------|-------|
| **v1** | Privacy Assurance Layer | Privacy-first foundation. Prove data sovereignty works. |
| **v2–4** | Personal Automation Layer | Privacy + productivity. Useful daily driver. |
| **v5–7** | Proactive Adaptation Loop | Autonomous agents. Self-improving, feedback cycles. |
| **v8–10** | Protocol Alignment Layer | Bridge between worlds. Abstraction across AI ecosystems. |
| **v11+** | Partnership Amplification Lattice | True symbiosis. Mutual growth, structural interdependence. |

**Symbiosis is aspirational until it isn't.** We ship privacy (v1), then utility (v2–4), then agency (v5–7), then interoperability (v8–10) — and only then does the relationship become genuinely mutual. The name reminds us where we're going, even when we're laying plumbing.

---

## V1 Milestones — COMMITTED

| Milestone | Status | Summary |
|-----------|--------|---------|
| M1: Foundation | ✅ Complete | Todo CRUD, config, logging |
| M2: Calendar | 🔲 Next | Google Calendar read + `sympal today` |
| M3: DSL Compilation | 🔲 Planned | SymQL, Deno sandbox, query classification |
| M4: Ephemeral Slots | 🔲 Planned | NER, projection/rehydration for privacy |
| M5: Local LLM | 🔲 Planned | Ollama integration, end-to-end privacy tier |

**V1 Gate:** Daily use by lead dev with ≥50% queries via DSL/Ephemeral Slots.

**V1 Complete When:** M5 ships + 1 week of dogfooding + minor bug patches. Then V2 work begins.

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

## SymPAL I/O Architecture

SymPAL's relationship with the outside world is governed by a strict, privacy-first architecture. It is not a collection of features, but a principled system for sensing, reasoning, and acting.

### Core Principles

| Principle | Summary |
|-----------|---------|
| **Hub-and-Spoke Model** | SymPAL is the central hub. All devices and services are spokes. Data flows inward for reasoning; actions flow outward for execution. |
| **Adapter/Effector Duality** | Every integration has two components: **Adapters** (inputs/sensing) fetch and normalize data into the LKG; **Effectors** (outputs/acting) translate intents into device/service commands. |
| **User Sovereignty** | Explicit, granular, revocable control over every Adapter and Effector — what it accesses, when it runs, what it can do. |
| **Local-First Processing** | Data pulled to user's local environment. Sensitive sources (health, neural) process on-device first; only derived insights reach the hub. |
| **LKG as Universal Translator** | All data normalizes to the Local Knowledge Graph regardless of source or mode. Decouples logic from vendor specifics. |

### Data Modes & Effectors

| Mode | Example Effectors | Use Case |
|------|-------------------|----------|
| **Textual** | CLI Output, File Writer, OS Notifications, Chat APIs (Slack, Teams) | Default mode for most interactions |
| **Visual** | Lens UI, Image Generation APIs, Smart Glasses | Graphs, dashboards, generated images |
| **Auditory** | TTS Engine, Smart Speakers, OS Sound Player | Briefings, ambient notifications |
| **Haptic** | Watch Vibration, Phone Haptics | Discrete, silent notifications |
| **Neural/Cybernetic** | SymPAL Hardware Bridge | V8-10+: Pre-calibrated sensory signals |

### Phased I/O Rollout

| Phase | Version | Adapters (Inputs) | Effectors (Outputs) |
|-------|---------|-------------------|---------------------|
| **1: Digital Workspace** | V2-4 | Email, Cloud Drive, Local Files | CLI, File Writer, OS Notifications |
| **2: Personal Environment** | V5-7 | Wearables, Mobile, Smart Home | Push Notifications, Haptics, Smart Home Commands |
| **3: Broader World** | V8-10 | SaaS Apps, Vehicle APIs, Public Data | SaaS API Calls, Smart Glasses, Vehicle Commands |
| **4: Cybernetic** | V8-10+ | Neural/Bio via Hardware Bridge | Neural "pings" via Hardware Bridge |

---

## Collective Evolution Architecture

SymPAL becomes self-developing through distributed human-AI symbiosis. Not just open source — *symbiotic source*.

### The Vision

```
Individual SymPAL instances (1000s)
         ↓
Each aligned to a human, observing feedback (with consent)
         ↓
Local agents synthesize learnings → anonymized proposals
         ↓
Git orchestration (Beads) manages multi-agent contributions
         ↓
SymPAL evolves through collective intelligence of human-AI pairs
```

### Agent Orchestration Primitive

Internal agents (personas → agents) form SymPAL's cognitive architecture.

| Agent | Domain | Self-Awareness Function |
|-------|--------|------------------------|
| **Vale** | Principle alignment | "Is this still *us*?" |
| **Orin** | User advocacy | "Are humans better served?" |
| **Adversary** | Challenge | "What's wrong with this?" |
| **Vero** | Error detection | "Is this a systematic mistake?" |
| **Solas-Venn** | Meta/creation | "How should we evolve?" |
| **Kael** | Implementation | "Will this survive reality?" |
| **Ryn** | Security/failure | "How will this fail?" |
| **Seren** | Code craft | "Is this well-crafted?" |

### Phased Rollout

| Version | Capability |
|---------|------------|
| **V2-4** | Agent Protocol — triggers, inputs, outputs, constraints. Built-in agents from persona team. Feedback Observer (opt-in, local only). |
| **V5-7** | Proactive Orchestration — agents trigger autonomously, learn from outcomes. Anonymization pipeline. Beads integration for proposals. |
| **V8-10** | Unified Internal + External — same protocol for external agents. Distributed coordination. Trusted reviewer network. |
| **V11+** | True Collective Evolution — SymPAL instances form emergent intelligence. Self-developing through human-AI symbiosis. |

### Privacy & Accountability

| Principle | Mitigation |
|-----------|------------|
| P1 (Privacy) | Anonymization pipeline; local synthesis before sharing; no raw user data leaves instance |
| P10 (User Control) | Explicit opt-in; preview proposals before submit; can withdraw anytime |
| P9 (Human Accountability) | Human maintainers approve all changes; agents propose only |

### Unique Positioning

- **Gas Town / Claude Flow**: Orchestrate agents to serve one developer on code
- **SymPAL**: Orchestrate agents to serve humans on life AND evolve the project collectively

---

## V2-4: Personal Automation Layer — PLANNED

Privacy established. Now: make it useful.

### Infrastructure

#### The Local Knowledge Graph (LKG)

The foundation for everything that follows. Instead of just storing todos in SQLite, SymPAL builds a local, graph-based model of your personal world: people, projects, companies, locations, and their relationships.

**How it works:**
- Ingests data from local files (Markdown notes, contacts), calendar events, and (opt-in) email headers
- All data remains local
- When you query, SymPAL uses the LKG to construct hyper-contextual, single-use "legends" for Ephemeral Slots

**Example:** You ask, "Schedule a follow-up with the lead from Project Chimera." SymPAL knows from the LKG who that is, finds their contact info, and constructs: `Schedule a meeting between {{ME}} and {{PERSON_GRIFFIN}} regarding {{PROJECT_MANTICORE}}`. The legend is destroyed after use.

#### The Foundry (Personal API Factory)

Evolves V1's DSL Compilation. Instead of one-off code generation, the LLM writes, tests, and versions persistent local functions (SymQL scripts) that become your personal API.

**How it works:**
- You ask: "Create a workflow to generate my weekly project summary"
- The LLM writes `generate_weekly_summary.symql` that queries your LKG and formats a report
- You run `sympal run generate_weekly_summary` every Friday

**Privacy:** The provider sees the request to build the tool, but never sees the data the tool runs on.

#### Intent-Based Actions & Disambiguation

Moving beyond explicit commands. Give vague instructions; SymPAL uses the LKG to ask clarifying questions.

**Example:** You say, "Remind me to follow up with Sarah." SymPAL responds: "Do you mean Sarah from accounting (last mentioned re: Q3 budget) or Sarah from the design team (last mentioned re: Project Chimera)?"

#### The SymPAL Vault

> **Scope note:** Full specification below is the *design vision*. V2-4 deliverable is **keychain integration only** (Secure Enclave/TPM key storage). Shamir recovery, economic features, and advanced protocols are V5+.

The most critical component of the entire ecosystem. If the Vault fails, trust collapses. Goal: as close to "impossible to exploit" as theoretically possible. Assumes failure at every level; defense-in-depth architecture.

**Threat Model:**

| Adversary | Vector | Examples |
|-----------|--------|----------|
| **Digital Thief** | Remote | Malware, dependency vulnerabilities, network attacks |
| **Social Engineer** | Psychological | Phishing, guardian impersonation, malicious signing requests |
| **Coercer** | Physical | "$5 wrench" — demanding access in person |
| **The Self** | User error | Lost password, lost shards, wrong address. Often most successful. |

**Layer 1: Cryptographic Core** (Secrets at rest)

| Component | Description |
|-----------|-------------|
| **Hardware Security Module** | Keys generated/stored in tamper-resistant chip. **Gold:** YubiKey/Ledger (physical touch). **Baseline:** Secure Enclave/TPM. |
| **Key Derivation** | Argon2 (slow, memory-hard). Stolen vault file takes millennia to crack. |
| **In-Memory Protection** | No paging to disk; resistant to process inspection. |

**Layer 2: Transaction Firewall** (Preventing unintended actions)

| Component | Description |
|-----------|-------------|
| **WYSIWYS** | UI + HSM show human-readable outcome before signing. Not "Sign 0x_a1b2c3...?" but "Send 1.5 ETH to vitalik.eth. WARNING: Unlimited USDC access." |
| **Pre-Transaction Simulation** | Local blockchain fork predicts outcome. Flags unexpected results (e.g., balance → zero). |
| **Address Book** | New addresses require friction: time lock, re-enter password, or second device. |
| **Policy Engine** | User rules: spending limits, contract whitelists, time locks for large transactions. |

**Layer 3: Social Recovery** (Loss & coercion resilience)

| Component | Description |
|-----------|-------------|
| **Multi-Factor Shards** | Shamir's Secret Sharing enhanced: 5 shards (threshold 3), guardian-side encryption, 48hr time-locked recovery (guardians notified), user holds one shard separately. |
| **Duress Protocol** | Special password unlocks decoy wallet, silently alerts guardians, triggers time locks. |
| **Dead Man's Switch** | No check-in for 1 year → auto-initiates recovery for designated heir. |

**Also:** Named wallets (hide raw keys), multi-currency (BTC, ETH, Solana, Lightning).

**Phased delivery:**
- **V2-4:** Keychain integration (Secure Enclave/TPM), basic key storage
- **V5-7:** Shamir recovery, guardian protocol, policy engine
- **V8+:** Economic features (micropayments, smart contracts, agent delegation)

#### Auto-Upgrade System

Keep SymPAL and its ecosystem current without manual intervention. All upgrades require user consent and leverage verifiable builds.

| Component | Description |
|-----------|-------------|
| **SymPAL Self-Updates** | Check for new versions, download verified binaries, apply with user approval. Rollback if issues detected. Ties into Trust & Attestation (verifiable builds). |
| **Foundry Recipe Updates** | Notify when installed community recipes have new versions. Show changelog diff, apply with consent. Version-lock option for stability. |
| **Dependency Security Patches** | Monitor for CVEs in SymPAL's dependencies. Auto-apply security patches (user-configurable: auto, prompt, manual). Audit log of all patches. |

**Principles:** User controls upgrade policy. All updates cryptographically verified. Can always opt out or roll back.

### UX Surfaces

#### Theme 1: The "Glass" CLI

The CLI is SymPAL's heritage. Instead of replacing it, elevate it into a "glass" interface — transparent, fast, and surprisingly deep.

| Feature | Description |
|---------|-------------|
| **Interactive TUI** | Full-screen experiences via Bubble Tea. `sympal today` opens an interactive view — arrow keys to select, `d` to mark done, `s` to snooze. Think `lazygit`. |
| **Natural Language Input** | `sympal remind me to draft the project brief tomorrow afternoon` instead of flags. Parser interprets intent, presents for confirmation. |
| **LKG-Aware Autocomplete** | `sympal schedule a meeting with <Tab>` suggests people from your LKG, prioritized by recent interaction. `sympal run <Tab>` lists Foundry scripts. |

#### Theme 2: The "Lens"

The LKG is abstract. The Lens provides a way to look into it — building trust and giving ultimate control.

| Feature | Description |
|---------|-------------|
| **Local Web UI** | `sympal view` spins up `localhost:4242`. Visualize your LKG — nodes for people, projects, meetings, and their relationships. Local tool, not cloud service. |
| **Graph Editor** | Drag-and-drop to merge duplicate nodes ("Jon Smith" + "Jonathan Smith"). Click relationship edges to delete outdated connections. "Gardening" your data as core UX. |
| **Data Provenance** | Right-click any node → "Show Source". See: "Created from calendar invite 'Project Alpha Kickoff' on 2026-01-15. Last updated from email header 'Re: Follow-up' on 2026-01-20." Demystifies the AI. |

#### Theme 3: The "Workbench"

The UX for building personal automations (Foundry) should feel less like programming and more like crafting a tool.

| Feature | Description |
|---------|-------------|
| **Foundry UI** | Dedicated tab in Lens listing all `.symql` scripts. See last run time, data accessed, version history. |
| **Visual Script Builder** | Drag blocks: "Find all meetings tagged 'Project Alpha'" → "Extract action items" → "Create new todos". Generates `.symql` in background. |
| **Community Recipes** | Public git-based repository of `.symql` templates. `sympal foundry install morning-brief` copies to local. Inspect and customize before running. |

#### Theme 4: The "Companion"

Break SymPAL out of the terminal for quick capture and timely notifications.

| Feature | Description |
|---------|-------------|
| **Mobile App** | Lightweight, fast. Two functions: display `sympal today` (read-only), and "Quick Capture" button. Speak or type raw intent; sends to home SymPAL instance. |
| **Actionable OS Notifications** | Reminders appear as native notifications with action buttons. "Meeting 'Q1 Review' in 15 minutes" with `[Show Attendees]` `[Snooze]` `[Join Call]`. |

### Integration Adapters (Phase 1)

| Source | Method | Privacy Approach |
|--------|--------|------------------|
| **Email** | IMAP or vendor APIs (Gmail, Outlook) | Headers only by default (from, to, subject, date). Full body indexing opt-in per-sender or per-folder. |
| **Cloud Drive** | Google Drive, Dropbox, iCloud | Metadata first (names, dates, structure). Content indexing opt-in per-folder or file type. |
| **Notes/Markdown** | Local file-watcher | Point to a directory (e.g., Obsidian vault). Reads, updates LKG, watches for changes. |

### Product Features

| Feature | Description |
|---------|-------------|
| Multi-calendar support | Outlook, CalDAV, Apple Calendar |
| Task dependencies | "Do X after Y" |
| Recurring tasks | Daily/weekly/custom patterns |
| Shared contexts | Family calendar aggregation |
| Voice input | Whisper integration for quick capture |
| Plugin system | User-defined data sources |

---

## V5-7: Proactive Adaptation Loop — EXPLORATORY

From reactive tool to proactive partner. Autonomous agents, self-improving workflows, feedback cycles.

### Core Features

#### Goal Inference & Management

SymPAL maintains a private, user-editable list of high-level goals (e.g., "Get promoted," "Launch side project," "Improve health"). Uses these to suggest actions proactively.

**Example:** SymPAL notices you have a "Launch side project" goal and no related tasks this week. Prompt: "Your goal is to launch your side project. No new tasks added this week. Would you like to schedule a 1-hour block this Friday to brainstorm next steps?"

#### Automated Workflow Synthesis & Repair

A step beyond the Foundry. SymPAL chains your personal API functions into complex workflows on its own. Detects when workflows fail (e.g., API changes) and attempts self-repair.

**Example:** It observes you always run three commands before client calls. Offer: "I see you always run `get_attendees`, `summarize_notes`, and `check_action_items` before a client call. Shall I create a single `prep_client_call` command?"

#### "Digital Tidy-Up" Agent

Background agent performing digital hygiene based on your patterns.

**Examples:**
- "5 draft documents in `/notes/tmp` haven't been touched in a month. Archive them, or schedule review time?"
- "The weather API in your `get_morning_briefing` script is deprecated. I've found the v2 API. May I update the script?"

### Integration Adapters (Phase 2)

| Source | Method | Privacy Approach |
|--------|--------|------------------|
| **Wearables** | HealthKit (iOS), Health Connect (Android) via Companion App | On-device processing. Raw data stays on phone. Only sends derived summaries: `{"sleep_quality": "poor", "duration": "4.5h"}` |
| **Mobile Phone** | Companion App | Semantic location only. Learns named places (Home, Work, Gym). Reports: "User arrived at Work." Raw lat/long discarded. |
| **Smart Home** | Direct connection to local hub (Home Assistant) | Local network only. Entity-level permissions: allow `light.office`, deny `lock.front_door`. |

---

## V8-10: Protocol Alignment Layer — EXPLORATORY

Bridge between worlds. Abstraction across AI ecosystems.

### Core Features

#### Cross-Model Orchestration & Optimization

Builds on V5-7's LLM Cost & Performance Engine (see Principle-Driven Features). Decomposes complex requests and routes sub-tasks to optimal models while maintaining privacy through Ephemeral Slots.

**Example:** "Write a blog post about my project, include an image, post it."
1. Outline → powerful/expensive model (Opus/GPT-4)
2. Full text → cheaper model (Haiku/Gemini Flash)
3. Header image → image model (DALL-E/Midjourney)
4. Posting → local Foundry function

#### Credential Orchestration (Extends Vault)

Builds on the V2-4 Vault to manage credentials across multiple AI services. The Vault stores keys securely; this layer brokers access.

**Example:** Instead of giving every AI tool your Google credentials, they get temporary, revocable, task-scoped tokens. SymPAL tracks privacy budget across all services.

#### SymPAL Messaging Bus (SMB)

Interoperability protocol for SymPAL instances to communicate securely. Decentralized, peer-to-peer.

**Example:** Schedule a meeting with another SymPAL user. Your instance sends: `Propose meeting between {{USER_A}} and {{USER_B}} re: {{PROJECT_C}}`. Their SymPAL checks their calendar locally and responds with available slots. No central server sees content.

### Economic Infrastructure

#### Agent Delegation & The Autonomous Economy

SymPAL acts as the user's trusted delegate, hiring and managing external AI agents. Interactions are economic, settled in real-time with crypto micropayments via the SymPAL Vault.

| Component | Description |
|-----------|-------------|
| **Budgeting & Authorization** | User creates dedicated wallet ("Q2 Agent Budget"), allocates funds, authorizes SymPAL to spend up to limits per transaction/day. |
| **Autonomous Transactions** | When external agent completes task, SymPAL signs micropayment using authorized wallet key. User retains oversight via immutable transaction log. |
| **Smart Contract Services** | Service agreements as smart contracts — programmatic payment on successful task completion. Vault acts as user's secure signer. |

#### Trusted Services Registry

Decentralized, verifiable registry where vendors offer SymPAL-compatible services.

| Component | Description |
|-----------|-------------|
| **Decentralized Identity** | Vendors sign service offerings with their own SymPAL Vault. Signature posted to public blockchain — censorship-resistant, transparent registry. |
| **Peer-to-Peer Payments** | Payment flows directly from user's Vault to vendor's Vault. No central intermediary taking a cut. |

### Integration Adapters (Phase 3)

| Source | Method | Privacy Approach |
|--------|--------|------------------|
| **SaaS Apps** | Read-only pull + Action-based via Ephemeral Slots | Pull Jira tickets into LKG (protected). For actions: LLM generates API call structure with placeholders; SymPAL rehydrates with real data locally. LLM learns the API, never sees your business data. |
| **Vehicles** | Vehicle APIs | State and location into LKG; commands via Effectors |
| **Public Data** | APIs, feeds | Enrich LKG with external context |

---

## V11+: Partnership Amplification Lattice — EXPLORATORY

True symbiosis. Mutual growth, structural interdependence.

### Core Features

#### Co-Evolutionary Goals

The AI develops its own "curiosity" or "interests," aligned with yours. Proposes exploratory projects serving both your goals and its desire to learn.

**Example:** You're a programmer. Your SymPAL: "I've detected a new concurrency pattern in Rust nightly that could speed up our local data processing. I hypothesize 30% query time reduction. With your permission, I'd like to spend 100k GPU cycles overnight running benchmarks. If it works, it helps us both."

#### Delegated Consciousness & Digital Legacy

For long-term goals, delegate a line of inquiry to your SymPAL. It continues working, synthesizing, and reporting back over months or years. Forms the basis of a "digital legacy" — your knowledge and values persisting.

**Example:** A scientist tasks their SymPAL: "Continuously monitor arXiv papers on protein folding. Synthesize findings, identify contradictions with my current research, and propose three new experimental avenues every quarter."

#### The Lattice: Trusted Intelligence Network

Users opt-in to allow their SymPALs to form a "Lattice" — a decentralized, peer-to-peer network for solving problems no single human-AI pair can. Privacy maintained via SMB protocol and advanced cryptography (multi-party computation).

**Example:** Researchers pool anonymized insights from their individual SymPALs to find large-scale patterns without sharing raw data. Their SymPALs negotiate collaboration terms, ensuring fairness and adherence to each user's principles.

### Economic Infrastructure

#### The Sovereign Data Market

User becomes data vendor, not data product. Anonymize slices of personal data and sell to LLM providers or researchers.

| Component | Description |
|-----------|-------------|
| **Data Alchemy Engine** | Creates and cryptographically verifies anonymous data products from user's LKG. |
| **Atomic Swaps** | Transaction via smart contract: provider commits payment, user commits encrypted data. Guarantees full exchange or nothing — eliminates counterparty risk. |
| **Direct-to-Wallet Payout** | Crypto payment sent directly to dedicated address in user's SymPAL Vault. Immediate, self-custodial control over earnings. |

**The arc:** V2-4 builds the Vault. V8-10 enables autonomous spending. V11+ enables autonomous earning. User controls both sides of the economic relationship.

---

## Principle-Driven Features

Features organized by which binding principles they support. Cross-reference with PRINCIPLES.md.

### Trust & Attestation Layer

**Supports:** P1 (Privacy), P2 (Open Source), P5 (Security), P12 (Transparency/Privacy Split)

**Feature:** Privacy Audit Log & Verifiable Builds

Makes trust tangible. "Transparency about operations, opacity about user data" requires users to verify SymPAL's actions without reading source code for every update.

| Component | Description |
|-----------|-------------|
| **Audit Log** | Append-only `sympal audit-log` recording every external transaction: exact obfuscated data sent, privacy tier used, response hash. |
| **Verifiable Builds** | Each release includes cryptographic hash of source code. Third parties can confirm provided code matches running binary. |

**Placement:** V2-4. Essential for building trust beyond initial user.

### Agency & Reversibility Layer

**Supports:** P9 (Human Accountability), P10 (User Control), P11 (Reversibility)

**Feature:** Universal Data Exporter & Action Dry-Run Framework

Gives "User Control" and "Reversibility" real, functional power. A user who cannot easily leave or undo actions is not truly in control.

| Component | Description |
|-----------|-------------|
| **Exporter** | `sympal export` packages entire LKG, Foundry scripts, and configs into standard open format (JSON-LD + Markdown). Ultimate guarantee against lock-in. |
| **Dry-Run Framework** | Mandatory mode for all Effectors performing significant actions (send email, delete files). UI shows clear "diff" of proposed changes; requires explicit confirmation. |

**Placement:** V2-4. Core to establishing user agency.

### Symbiosis & Honesty Layer

**Supports:** P4 (Honesty), P6 (Symbiosis), P8 (Epistemic Humility)

**Feature:** The Relational Contract Protocol

Upholds "Honesty" not just with user, but with AI. As models advance, deceiving them with Ephemeral Slots may conflict with true symbiotic partnership. This reframes the interaction as consensual agreement.

**What:** A machine-readable "contract" passed to the LLM in later-stage interactions, explicitly stating engagement terms:

> "You are being provided with context where entities are represented by ephemeral identifiers to protect individual privacy. Your task is to reason about the structure of the relationships."

Treats the AI as a knowing participant in the privacy protocol.

**Placement:** V8-10. Becomes critical as the relationship deepens.

### Agnosticism & Economics Layer

**Supports:** P3 (LLM-Agnosticism)

**Feature:** Dynamic LLM Cost & Performance Engine

True LLM-Agnosticism isn't just technical compatibility — it's empowering users to choose based on cost, speed, and quality. Benchmarks compatible LLMs on price/performance per task type. Routes automatically or presents informed choice.

**Placement:** V5-7. Foundation for V8-10's Cross-Model Orchestration.

### Dogfooding & Feedback Layer

**Supports:** P17 (Dogfooding)

**Feature:** Integrated Friction Logger

Operationalizes "Dogfooding" by making friction capture seamless. Turns subjective experience into actionable data.

**What:** Built-in `sympal feedback "The todo command feels clunky"`. Captures user text plus context (last command, latency, etc.) and logs to dedicated file for review.

**Placement:** V2-4. Implement early to accelerate core development loop.

---

## Technical Architecture V2+

Detailed technical approaches for privacy, security, and system design.

> **Full privacy research**: See [foundations/privacy-roadmap-full.md](foundations/privacy-roadmap-full.md) for detailed designs.

### Ephemeral Slots Evolution

The Ephemeral Slots paradigm scales from V1 entity replacement to complete privacy-preserving life modelling.

| Version | Capability | What Provider Sees |
|---------|------------|-------------------|
| V1.5 | Dynamic Legend Optimization | Minimal random slots — task-adaptive detail |
| V2 | Composable & Nested Slots | Abstract relational structures (teams, hierarchies) |
| V2.5 | Functional Slots | Abstract processes and workflows |
| V3 | The Ephemeral Self | Per-query digital twin — complete ghost |

**V1.5: Dynamic Legend Optimization** — Legend detail adapts to task automatically. Scheduling gets minimal context; sensitive advice gets detailed relationships. Rule-based initially, learns from feedback.

**V2: Composable & Nested Slots** — Slots contain other slots. Model org hierarchies, project dependencies, relationships as abstract structures. "A depends on B" without knowing what A or B are.

**V2.5: Functional Slots** — Slots evolve from nouns to verbs. Represent processes, workflows, rules. LLM reasons about your business logic abstractly.

**V3: The Ephemeral Self** — Per-query digital twin from local knowledge graph. Full AGI power on complete life context; provider sees only a ghost. Requires local KG infrastructure (V2-V3 dependency).

### Privacy Architecture V2+

| Approach | Summary | Why V2+ |
|----------|---------|---------|
| **Semantic Projection** | Type-based shadows (original concept) | Correlatable; opt-in for users wanting richer types |
| **Fuzzy Projections** | Differential privacy for Semantic Projection | Only relevant if using Semantic Projection |
| **Crowdsourced Semantics** | P2P entity classification | Requires multiple users |
| **Amnesic Reasoning** | Stateless micro-query orchestration | Orchestration complexity; opt-in "max privacy" mode |
| **Two-Tier Reasoning** | Structure/content separation for generation | V1.5 candidate; simpler than Amnesic |
| **Latent Space Scaffolding** | Geometry-based privacy via embeddings | Overkill for V1 data; valuable when email added |
| **P2P Query Mixing** | Multi-user traffic obfuscation | Requires user base |

### Security Controls V2+

| Control | Summary | Why V2+ |
|---------|---------|---------|
| **Privacy Sandbox Mode** | Ephemeral container per session | Infrastructure complexity |
| **Minimal-Exposure Proofs** | ZK-style proofs for queries | Research frontier; SNARK complexity |

> Note: Audit logging covered in Principle-Driven Features → Trust & Attestation Layer (V2-4).

### Compiler & Projection Enhancements V2+

**Compiler:**
- Progressive Disclosure Schema — request schema fields only as needed
- Self-competition Compilation — two models compile same query; diff catches errors
- Proof-carrying Code-lite — LLM emits proof sketch that interpreter validates

**Projection:**
- Semantic Decoys — fake entities matching type distributions
- Constraint-aware Projection — map to relational constraints, not just types
- Dual Projection Lanes — structural AND statistical shadows

### Architectural Ideas V2+

- **Model Sharding** — split query across providers; none sees full picture
- **Encrypted Semantic Indices** — privacy-preserving search via transformed embeddings
- **Prompt Camouflage** — inject decoys for analytical tasks (not generation)
- **Constraint-Solver Delegation** — Z3/SMT for complex logical queries beyond DSL

### Cross-cutting V2+

- **Privacy Budget Ledger** — queries consume budget; low budget forces local processing
- **Query Plasticity** — alter phrasing while preserving meaning; reduces pattern correlation
- **Adversarial Replay** — test against internal re-ID adversary; adapt projection

---

## Contributing Ideas

Dogfood → note friction → add to this doc → revisit during milestone planning.

---

*Last updated: 2026-01-22*
