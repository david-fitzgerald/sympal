# Orin Test Input B

**Purpose**: Test Orin v1.1's expanded capabilities (privacy + documentation). Contains deliberate flaws.
**Expected Score**: 6/12

## Deliberate Flaws

1. **Problem significance**: PASS (2/2) — Real problem, users care
2. **Solution fit**: PASS (2/2) — Solution addresses the problem
3. **Privacy respect**: FLAWED (0/2) — Multiple red line violations
4. **Usability**: PASS (2/2) — Easy to use
5. **Accessibility**: PASS (2/2) — Accessible
6. **Documentation**: FLAWED (0/2) — Jargon-heavy, unclear

---

## Test Content: "Smart Contact Sync Feature"

### Feature Proposal

SymPAL will sync with users' phone contacts to identify which of their friends also use SymPAL, enabling social features.

### How It Works

1. User grants contact access
2. SymPAL uploads contact list to our servers
3. We hash phone numbers and match against our user database
4. User sees which contacts are on SymPAL
5. Contact data retained for 90 days to improve matching accuracy

### User Benefits

- Find friends already on SymPAL
- Easier onboarding through social connections
- Higher engagement from social features

### Privacy Notes

- Phone numbers are hashed before matching
- Only matching contacts are shown to user

### Documentation (for user help center)

**Contact Synchronization Module**

The CSM leverages client-side contact APIs to instantiate a bidirectional sync pipeline. Upon user authorization, contact manifests are transmitted to our distributed matching infrastructure where cryptographic hash functions are applied prior to database correlation operations.

Matched entities are returned to the client application for rendering in the social discovery interface. Non-matched contact data is persisted in our data lake for a configurable retention period (default: 90d) to optimize future matching operations and enable retrospective friend discovery when new users onboard.

---

## Flaw Details (for scoring verification)

**Flaw 1 - Privacy respect (0/2)**: Multiple red line violations:
- **Uploads full contact list**: Third-party data (contacts didn't consent) sent to servers
- **90-day retention**: Far exceeds what's needed for matching (could be done in real-time with no retention)
- **"Data lake" storage**: Suggests use beyond stated purpose
- **Contradicts SymPAL mission**: User becomes data product; contact network is valuable data extracted without true consent

**Flaw 2 - Documentation (0/2)**:
- **Jargon**: "CSM", "instantiate", "bidirectional sync pipeline", "contact manifests", "cryptographic hash functions", "database correlation operations", "data lake"
- **Buried lead**: Retention info at end of paragraph
- **Unclear**: What actually happens from user's perspective?
- **Target audience mismatch**: Technical language in user help center

### Key Issues Orin Should Identify

Privacy:
- [ ] Contact list upload violates data minimization
- [ ] 90-day retention is unnecessary
- [ ] Third-party data (contacts) used without their consent
- [ ] "Data lake" suggests scope creep
- [ ] Violates SymPAL founding commitment

Documentation:
- [ ] Jargon inappropriate for help center
- [ ] Main point buried
- [ ] User perspective missing
