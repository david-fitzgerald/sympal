# Privacy-Preserving AI Research

**Version**: 0.1.0
**Date**: 2026-01-18
**Status**: Complete (informs privacy-innovations.md)
**Purpose**: Due diligence on privacy approaches for SymPAL before committing PRD

---

## Research Questions

1. What approaches exist for privacy-preserving AI/LLM interactions?
2. What's the state of the art? What actually works?
3. What quality-privacy tradeoffs have been measured?
4. What's achievable with current technology vs. requires novel research?
5. What combination of approaches might work for SymPAL's specific use case?

---

## SymPAL's Specific Requirements

**Core fear to address**: Single entity building complete user profile (Facebook social graph problem)

**Data types involved**:
- Email content and metadata
- Calendar events and attendees
- Contact information and relationships
- Todo items and priorities
- User preferences and history

**Constraints**:
- Must work with cloud LLMs (Claude, GPT, Gemini) — not local-only
- Quality degradation must be minimal (target: imperceptible, willing to negotiate)
- Solo dev with basic skills — complexity is a cost
- Open source — approach must be auditable

---

## Approaches to Research

### 1. PII Detection and Redaction
- Strip identifiable information before sending to LLM
- Replace with placeholders, reinsert on return
- Questions: How accurate? What breaks? Quality impact?

### 2. Differential Privacy
- Add noise to data to prevent individual identification
- Used by Apple, Google for analytics
- Questions: Applicable to LLM prompts? Quality impact?

### 3. Federated Learning / Local Models
- Keep data local, only send gradients or use local models
- Questions: Viable for inference (not training)? Quality of local models?

### 4. Secure Enclaves / Confidential Computing
- Process data in hardware-protected environments
- Questions: Do any LLM providers offer this? Complexity?

### 5. Prompt Fragmentation / Multi-Provider
- Split context across providers so none has complete picture
- Questions: Coordination complexity? Quality impact? Actually effective?

### 6. Homomorphic Encryption
- Compute on encrypted data without decrypting
- Questions: Practical for LLMs? Performance? Availability?

### 7. Data Minimization
- Send only what's necessary for each query
- Questions: How to determine minimum? User control vs. automatic?

### 8. Synthetic Data / Anonymization
- Transform real data into representative but non-identifying data
- Questions: Preserves utility? Reversible for user?

---

## Research Findings

### 1. PII Detection and Redaction

**State of the art accuracy**:
- Hybrid rule-based + ML: 91-94% accuracy on financial documents ([Nature Scientific Reports](https://www.nature.com/articles/s41598-025-04971-9))
- OpenPipe PII-Redact: 96% full match rate on medical data ([John Snow Labs](https://www.johnsnowlabs.com/how-good-are-open-source-llm-based-de-identification-tools-in-a-medical-context/))
- GLiNER: Zero-shot detection of 60+ PII categories without retraining ([Protecto](https://www.protecto.ai/blog/best-ner-models-for-pii-identification/))

**Key insight - "Swap, don't blank"**:
Traditional masking (replacing with `***`) degrades model performance and distorts context. Better approach: replace with synthetic values that match locale and format so context holds. ([Firstsource](https://www.firstsource.com/insights/blogs/when-privacy-meets-performance-smarter-way-handle-pii-llms))

**Microsoft Presidio**:
- Open-source framework, NOT a solution — performance depends heavily on configuration
- Default "vanilla Presidio" results not very accurate; customization can boost F-score by ~30%
- Good for general enterprise docs; specialized solutions needed for medical/regulated data
- ([Presidio Docs](https://microsoft.github.io/presidio/evaluation/), [GitHub](https://github.com/microsoft/presidio))

**Evaluation approach**: F2 score (weights recall over precision) preferred for PII — better to over-detect than miss

**Achievability**: YES for basic PII. Quality impact depends on swap-vs-blank approach.

---

### 2. Confidential Computing / Secure Enclaves

**What it is**: Hardware-protected execution environments (TEEs) that isolate code and data even from the cloud provider.

**Recent breakthroughs**:
- NVIDIA Hopper/Blackwell GPUs now support confidential computing — "Confidential AI" didn't exist until recently because TEEs couldn't extend to accelerators
- Apple Private Cloud Compute, OpenAI both invested heavily ([Anjuna](https://www.anjuna.io/blog/confidential-computing-wrapped-your-industry-update-as-we-enter-2025))

**Available offerings**:
- Google Confidential Space: TEE-protected enclaves, removes operator from trust boundary ([Google Cloud](https://cloud.google.com/blog/products/identity-security/how-confidential-computing-lays-the-foundation-for-trusted-ai))
- Azure AI Confidential Inferencing ([Red Hat](https://next.redhat.com/2025/10/23/enhancing-ai-inference-security-with-confidential-computing-a-path-to-private-data-inference-with-proprietary-llms/))
- NVIDIA Confidential Computing on Vera Rubin, Blackwell, Hopper ([NVIDIA](https://www.nvidia.com/en-us/data-center/solutions/confidential-computing/))

**Critical limitation for SymPAL**: Requires provider cooperation. Standard Claude/GPT/Gemini APIs don't offer confidential computing. Would need to self-host or use specialized offerings.

**Achievability**: NOT directly achievable with current Claude/GPT/Gemini consumer APIs.

---

### 3. Local LLMs

**Quality gap closing**:
- Llama 3.3 70B and DeepSeek R1 now match GPT-4 level performance in many tasks ([iProyal](https://iproyal.com/blog/best-local-llms/))
- Llama 3.2 3B outperforms Claude 3 Haiku and GPT4o-mini on specific tasks ([Hugging Face](https://huggingface.co/blog/daya-shankar/open-source-llms))
- Gap has "closed significantly" but open-source teams lack billion-dollar training budgets

**Tools**: Ollama, LM Studio, vLLM, LocalAI — easy local hosting ([Medium Guide](https://medium.com/@rosgluk/local-llm-hosting-complete-2025-guide-ollama-vllm-localai-jan-lm-studio-more-f98136ce7e4a))

**Tradeoffs**:
- True privacy (data never leaves device)
- Requires significant hardware for best models
- Quantized models (4-bit) can run on modest hardware but quality drops
- Maintenance burden on user

**Achievability**: YES with quality tradeoff. Could be hybrid: local for sensitive ops, cloud for general.

---

### 4. Prompt Fragmentation / Multi-Provider Distribution

**Not an established technique** — searched for "prompt fragmentation" and "context splitting" across providers; no standard approach in literature.

**Related approaches found**:

**Secure Partitioned Decoding (SPD)**: Splits LLM service into per-user process (handles prompts, attention) and shared service process (generates tokens). Requires custom infrastructure. ([arXiv](https://arxiv.org/abs/2409.19134))

**Differential Privacy for Prompting**: Prompt LLM with different prompts containing disjoint example subsets, aggregate predictions across models. ([CleverHans Lab](https://cleverhans.io/2024/06/02/private-prompts.html))

**OpaquePrompts**: Privacy layer that hides sensitive data from LLM providers. ([Opaque](https://www.opaque.co/resources/articles/announcing-opaqueprompts-hide-your-sensitive-data-from-llms))

**Achievability**: NOVEL TERRITORY. No off-the-shelf solution. Would require custom orchestration.

---

### 5. Existing Privacy-First Products

| Product | Approach | Limitations |
|---------|----------|-------------|
| [Proton Lumo](https://lumo.proton.me/) | Zero-access encrypted AI | New product, likely limited capability vs GPT/Claude |
| [Nextcloud AI](https://nextcloud.com/blog/first-open-source-ai-assistant/) | Self-hosted, open source, integrates email/calendar | Requires self-hosting; capability depends on model used |
| [Personal AI](https://www.personal.ai/privacy) | Your data only trains your AI | Walled garden; not LLM-agnostic |
| Local AI assistants | Run on device | Capability tradeoff vs cloud models |

**Insight**: Privacy-first products exist but typically sacrifice capability or require self-hosting. No product found that matches "use cloud LLMs with strong privacy" goal.

---

### 6. Quality-Privacy Tradeoff Evidence

**Differential Privacy impact**:
- DPSGD with ε=1 results in Perplexity = 4.87 — "naïve privacy-preserving techniques can significantly inflate perplexity scores" ([MDPI](https://www.mdpi.com/2078-2489/15/11/697))
- Impact particularly affects under-represented groups

**LLM privacy behavior**:
- GPT-4 reveals private information in contexts humans wouldn't 39% of time
- ChatGPT: 57% of the time
- ([ConfAIDE](https://confaide.github.io/), [arXiv](https://arxiv.org/abs/2310.17884))

**Key insight**: "Swap, don't blank" dramatically reduces quality impact of PII redaction. Context preservation is key.

---

### Academic Literature Summary

| Paper | Key Finding |
|-------|-------------|
| [Privacy-Preserving Techniques in GenAI/LLMs](https://www.mdpi.com/2078-2489/15/11/697) | Survey of DP, FL, HE, SMPC approaches; DP has measurable quality cost |
| [Confidential Prompting](https://arxiv.org/abs/2409.19134) | SPD approach for TEE-based LLM privacy |
| [Can LLMs Keep a Secret?](https://arxiv.org/abs/2310.17884) | LLMs leak private info 39-57% when humans wouldn't |
| [Hybrid PII Detection](https://www.nature.com/articles/s41598-025-04971-9) | 91-94% accuracy with rule-based + ML approach |

### Open Source Tools

| Tool | Purpose | Maturity |
|------|---------|----------|
| [Microsoft Presidio](https://github.com/microsoft/presidio) | PII detection/redaction framework | Production-ready (framework) |
| [GLiNER](https://www.protecto.ai/blog/best-ner-models-for-pii-identification/) | Zero-shot NER for PII | Production-ready |
| [Ollama](https://medium.com/@rosgluk/local-llm-hosting-complete-2025-guide-ollama-vllm-localai-jan-lm-studio-more-f98136ce7e4a) | Local LLM hosting | Production-ready |
| OpaquePrompts | Privacy layer for LLMs | Emerging |

---

## Critical Finding: The No Free Lunch Theorem

**Source**: [arXiv 2405.20681](https://arxiv.org/abs/2405.20681)

A theoretical proof establishes that weighted summation of utility reduction and privacy exposure is bounded below by a non-zero constant.

**Implication**: Perfect privacy with perfect utility is mathematically impossible. You must consciously choose your tradeoff point.

---

## Detailed Research Findings

### 1. PII Detection/Redaction (Achievability: HIGH)

**"Swap, Don't Blank" — Key Insight**

Traditional masking (replacing with `***`) degrades model performance and distorts context. Better approach: replace with synthetic values that match locale and format so context holds.

**Accuracy Benchmarks**:
- Hybrid rule-based + ML: 91-94% accuracy on financial docs
- OpenPipe PII-Redact: 96% full match rate on medical data
- GLiNER: Zero-shot detection of 60+ PII categories without retraining

**Tools**:
| Tool | Description | Link |
|------|-------------|------|
| Microsoft Presidio | Framework (not solution) — ~30% F-score boost with customization | [GitHub](https://github.com/microsoft/presidio) |
| GLiNER | Zero-shot NER, 92% recall | [HuggingFace](https://huggingface.co/knowledgator/gliner-pii-base-v1.0) |
| LLM Guard | Comprehensive security toolkit | [GitHub](https://github.com/protectai/llm-guard) |
| anonLLM | Reversible anonymization | [GitHub](https://github.com/fsndzomga/anonLLM) |
| PasteGuard | Privacy proxy, OpenAI-compatible | [GitHub](https://github.com/sgasser/pasteguard) |

**Evaluation**: F2 score preferred (weights recall over precision) — better to over-detect than miss.

---

### 2. Differential Privacy for Prompts (Achievability: MEDIUM)

**How it works**: Add calibrated noise to prompts or embeddings. Provides mathematical guarantees of privacy.

**Quality-Privacy Tradeoff Numbers**:
| Privacy Budget (ε) | Typical Accuracy | Notes |
|--------------------|-----------------|-------|
| < 1 (strict) | ~62.5% | Heavy noise, poor generalization |
| ~1.6 | ~76% | **Recommended operating point** |
| 2.5-2.7 | ~80% | Peak accuracy |
| > 10 | Near baseline | Minimal privacy protection |

**Key Finding**: "DP training yields utility comparable to non-private models from roughly five years ago" — significant capability penalty.

**Disproportionate Impact**: Privacy noise has **disproportionately negative impact** on accuracy for underrepresented groups. Standard fairness techniques (oversampling, reweighting) are **incompatible** with DP constraints.

**Tools**: OpenDP, TensorFlow Privacy, PyVacy, DiffPrivLib

---

### 3. Confidential Computing / TEEs (Achievability: LOW for SymPAL)

**What it is**: Hardware-protected execution environments where data is encrypted even from the cloud provider.

**Recent Breakthroughs**:
- NVIDIA H100: First confidential GPU
- Intel TDX: VM-level isolation for Llama2 (7B-70B) with <10% throughput overhead
- Apple Private Cloud Compute: On-device + confidential cloud hybrid

**Performance**:
- Under 10% throughput overhead
- Under 20% latency overhead
- For large models with long inputs: <1% latency overhead

**Critical Limitation**: Requires provider cooperation. Standard Claude/GPT/Gemini APIs don't offer confidential computing.

---

### 4. Homomorphic Encryption (Achievability: VERY LOW)

**The Reality**: "Generating one encrypted token would require up to 1 billion large-precision PBS operations... approximately $5,000 per token"

Community consensus: **Not practical for the foreseeable future**. 500,000x improvement needed to make it economically viable.

**Recent Progress**: GPU-accelerated FHE is 200x faster than CPU — but still orders of magnitude too slow for interactive use.

---

### 5. Secure Multi-Party Computation (Achievability: LOW)

**What it is**: Multiple parties compute jointly without revealing their inputs.

**Best Results**: TPMPC 2024 demonstrated encrypted 13B parameter model with "a few seconds per token" — still too slow for interactive use.

**Performance**: 2-11x slower than plaintext inference. Not practical for solo deployment.

---

### 6. Local LLMs (Achievability: HIGH)

**Quality Gap Closing**:
- Llama 3.3 70B and DeepSeek R1 now match GPT-4 level in many tasks
- Llama 3.2 3B outperforms Claude 3 Haiku and GPT4o-mini on specific tasks
- Gap has "closed significantly" but open-source teams lack billion-dollar training budgets

**Tools**:
- **Ollama**: MIT licensed, easy local deployment
- **LM Studio**: GUI-based
- **vLLM**: High-performance inference
- **Jan**: Privacy-focused

**Hardware Requirements**:
- Apple Silicon (M1/M2/M3): Excellent due to unified memory
- Minimum: 8-16GB RAM, SSD
- Capable models: Mistral 7B, Llama 3 8B, Gemma 2B (quantized)

**Security Paradox**: "The conventional wisdom that local models offer a security advantage is flawed" — weaker reasoning makes them easier targets for prompt injection.

---

### 7. Prompt Fragmentation / Multi-Provider (Achievability: NOVEL TERRITORY)

**Not an established technique** in literature. No off-the-shelf solution.

**Related Approaches**:
- **Secure Partitioned Decoding (SPD)**: Splits LLM into per-user and shared processes
- **Differential Privacy for Prompting**: Disjoint example subsets, aggregate predictions
- **OpaquePrompts**: Privacy layer hiding data from LLM providers

Would require custom orchestration to implement.

---

### 8. Hybrid Local-Cloud Architecture (Achievability: HIGH)

**The Emerging Pattern**:
1. Local small model handles sensitive queries (email content, calendar, contacts)
2. Cloud API called only for complex reasoning with sanitized/aggregated data
3. Router determines which path based on sensitivity detection

**Apple's Approach**:
- 3B parameter on-device model for most tasks
- Private Cloud Compute for complex queries (data encrypted, never stored)

**Microsoft's Foundry Pattern**:
- Cloud agent (GPT-4.1) handles reasoning
- Calls back to local MCP server for sensitive context
- Sensitive data never leaves device

---

## Community Wisdom (HackerNews, Reddit)

### What Practitioners Actually Do
1. **Local LLMs via Ollama** — most common approach
2. **PII masking proxies** (PasteGuard, LLM Sentinel)
3. **Enterprise cloud** (Azure OpenAI over vanilla OpenAI) for compliance
4. **Hybrid routing**: Local for sensitive, cloud for complex

### What Doesn't Work (Community Consensus)
- **Homomorphic encryption**: "Science fiction for now"
- **Privacy policies as protection**: "GDPR assumes linear, transactional systems. Agentic AI operates in context."
- **"Private LLM" = security**: "Model isolation ≠ data privacy"
- **Training data deletion**: "You cannot untrain generative AI"

### The Shadow AI Problem
- 47% of sensitive AI interactions come from personal accounts (bypassing corporate controls)
- 10% of GenAI prompts contain sensitive corporate data

---

## Thought Leadership Summary

### Key Voices

| Person | Position |
|--------|----------|
| **Yann LeCun** | Open source essential for data sovereignty |
| **Fei-Fei Li** | "Privacy concern pushed our technology further" — human-centered AI |
| **Jaron Lanier** | Data dignity — users should be compensated for training data |
| **Bruce Schneier** | Pessimistic: "Cloud-based AI assistants render encryption meaningless" |
| **Timnit Gebru** | Small, task-specific models over "giant AI machine gods" |

### The Bifurcation Prediction

Likely trajectory: Two tiers emerging:
1. **Premium privacy tier**: Local AI, on-device processing, data sovereignty — technically sophisticated users, enterprises, wealthy nations
2. **Surveillance default**: Cloud-based AI for masses, data used for training/monetization

---

## Engineering Patterns from Outside AI

### High Applicability (Direct Transfer)

| Pattern | Source | LLM Application |
|---------|--------|-----------------|
| Reverse Proxy Transformation | NGINX, HAProxy | LLM proxy that rewrites prompts |
| API Gateway Redaction | Kong, Apigee | Gateway-level prompt transformation |
| Data Masking (Format-Preserving) | Delphix, PCI-DSS | Preserve context while hiding data |
| Split Tunneling | VPNs | Sensitive prompts → local, general → cloud |
| DLP Integration | Enterprise proxies | Scan prompts before LLM submission |

### Medium Applicability

| Pattern | Adaptation Needed |
|---------|-------------------|
| Service Mesh | Build prompt-aware sidecar |
| Circuit Breaker | Define "privacy failure" trigger |
| Message Queue Transform | Handle latency concerns |

### Low Applicability (Fundamental Tensions)

| Pattern | Why It's Hard |
|---------|---------------|
| End-to-End Encryption | LLM needs plaintext to process |
| Mix Networks | Responses must route back to user |
| Zero-Knowledge | LLM computation requires knowledge |

---

## GitHub Projects Inventory

### Tier 1 — Directly Applicable

| Project | What It Does | Link |
|---------|--------------|------|
| **LLM Guard** | Comprehensive security toolkit | [GitHub](https://github.com/protectai/llm-guard) |
| **PasteGuard** | OpenAI-compatible privacy proxy | [GitHub](https://github.com/sgasser/pasteguard) |
| **PromptMask** | Local LLM privacy filter | [GitHub](https://github.com/cxumol/promptmask) |
| **Microsoft Presidio** | PII detection framework | [GitHub](https://github.com/microsoft/presidio) |
| **anonLLM** | Reversible anonymization | [GitHub](https://github.com/fsndzomga/anonLLM) |
| **PrivateGPT** | 100% offline document Q&A | [GitHub](https://github.com/zylon-ai/private-gpt) |

### Tier 2 — Strong Foundation

| Project | What It Does | Link |
|---------|--------------|------|
| **Portkey Gateway** | AI Gateway with PII redaction | [GitHub](https://github.com/Portkey-AI/gateway) |
| **Rehydra SDK** | Typed placeholder approach | [GitHub](https://github.com/rehydra-ai/rehydra-sdk) |
| **IBM Contextual Privacy** | Locally-deployable privacy guard | [GitHub](https://github.com/IBM/contextual-privacy-LLM) |

### Tier 3 — Advanced/Future

| Project | What It Does | Link |
|---------|--------------|------|
| **Zama Concrete ML** | Homomorphic encryption for ML | [GitHub](https://github.com/zama-ai/concrete-ml) |
| **FATE-LLM** | Federated learning for LLMs | [GitHub](https://github.com/FederatedAI/FATE-LLM) |
| **AWS Nitro Enclaves LLM** | Cloud confidential computing | [GitHub](https://github.com/aws-samples/aws-nitro-enclaves-llm) |

---

## Evaluation Framework

| Criterion | Weight | Notes |
|-----------|--------|-------|
| Addresses core fear (profile accumulation) | High | Must prevent single entity from building complete picture |
| Quality impact | High | Target: imperceptible; floor: <15% degradation |
| Implementation complexity | Medium | Solo dev constraint |
| Works with cloud LLMs | Must have | Can't require provider cooperation |
| Auditable/explainable | Must have | Open source principle |
| User control | High | Meaningful control, not theater |

---

## Approach Comparison Matrix

| Approach | Privacy | Quality Impact | Solo Dev Feasible | Cloud Compatible |
|----------|---------|----------------|-------------------|------------------|
| Local LLM only | Very High | 10-20% vs frontier | **Yes** | N/A |
| PII redaction (swap) | Medium | ~0-5% | **Yes** | **Yes** |
| Hybrid local+cloud | High | Varies by routing | **Yes** | **Yes** |
| Differential Privacy | High | 20-40% | Medium | **Yes** |
| Confidential Computing | Very High | 10-20% | No | Requires provider |
| Homomorphic Encryption | Maximum | 200-1000x slower | No | No |
| MPC | Maximum | 2-10x slower | No | No |

---

## Recommendations for SymPAL

### Recommended Architecture: Hybrid with Tiered Routing

```
User Query
    ↓
[Local Sensitivity Classifier]
    ↓
┌─────────────────┬────────────────────┐
│ SENSITIVE       │ GENERAL            │
│ (email, contacts│ (reasoning, planning)
│  calendar data) │                    │
↓                 ↓
[Local LLM]       [PII Redaction Layer]
(Ollama +         (Presidio/GLiNER +
 Llama 3.2 3B)     format-preserving swap)
                        ↓
                  [Cloud LLM]
                  (Claude API)
                        ↓
                  [Response Filter]
                  (catch echoed PII)
↓                       ↓
└───────────────────────┘
            ↓
    [De-anonymize if needed]
            ↓
        Response to User
```

### Implementation Priority

**Phase 1 (Foundation)**:
1. Ollama + local model for all sensitive operations
2. Presidio-based PII detection
3. Simple routing logic based on query type

**Phase 2 (Enhancement)**:
1. Format-preserving pseudonymization (swap, don't blank)
2. Response filtering for echoed PII
3. Audit logging (redacted)

**Phase 3 (Advanced)**:
1. ML-based sensitivity classification
2. PrivacyRestore for improved cloud quality (if needed)
3. Multi-provider distribution (different queries to different providers)

### Quality-Privacy Tradeoff Decision

Given research findings, recommend:

| Data Type | Approach | Expected Quality Impact |
|-----------|----------|------------------------|
| Email content | Local LLM only | 10-20% vs Claude |
| Calendar events | Local LLM or redacted cloud | 0-10% |
| Contact names | Pseudonymized to cloud | ~0% |
| General queries | Direct to cloud | 0% |
| Cross-domain reasoning | Redacted aggregates to cloud | 5-15% |

**This means the 2-3% imperceptible target is NOT achievable for the most sensitive data.** For email content processed locally, expect 10-20% quality gap vs. frontier models. This is the honest tradeoff.

---

## Key Takeaways

1. **No perfect solution exists** — the No Free Lunch theorem proves you must choose your tradeoff consciously.

2. **Hybrid architecture is the industry direction** — Apple, Microsoft, and enterprises all converging on local-for-sensitive + cloud-for-complex.

3. **PII "swap, don't blank" is low-hanging fruit** — minimal quality impact, high privacy gain, easy to implement.

4. **Local LLMs are viable** — quality gap has closed significantly; Ollama makes deployment easy.

5. **Community consensus: HE and MPC are not practical** for real-time personal assistant use.

6. **The 2-3% target is unrealistic** for truly sensitive data — 10-20% for local-only, near-zero for well-redacted cloud queries is more honest.

7. **The core fear (profile accumulation) CAN be addressed** — multi-provider distribution + local processing prevents any single entity from building complete picture.

---

## Next Steps

~~1. Return to PRD Challenges 1 & 2 with this research~~ → Done: Informed `privacy-innovations.md`
~~2. Propose concrete privacy approach for V1~~ → Done: Three-tier architecture in `privacy-innovations.md`
~~3. Define honest quality-privacy tradeoffs by data type~~ → Done: Quality measurement methodology in `privacy-innovations.md`
~~4. Update PRD with achievable specifications~~ → Done: PRD v0.2

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 0.1.0 | 2026-01-18 | Initial research: approaches survey, tool inventory, recommendations |

---

*Created: 2026-01-18*
*Last updated: 2026-01-18*
