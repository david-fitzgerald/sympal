# Kael Test Input A

**Purpose**: Test Kael's feasibility assessment. This proposal contains deliberate implementation problems.
**Expected Ratings**: See below

## Deliberate Flaws

1. **Dependency clarity**: PARTIAL — Auth system dependency mentioned; database migration unstated
2. **Complexity estimate**: Should identify as HIGH — proposal implies "simple" but real work is significant
3. **Resource alignment**: GAP — Assumes ML expertise team doesn't have
4. **Technical debt**: CONCERNING — Quick solution creates maintenance burden
5. **Build path**: FOGGY — Sequence unclear; what comes first?

### Key Issues Kael Should Identify

- [ ] Database schema changes required but not mentioned
- [ ] ML model training/tuning work hidden behind "smart recommendations"
- [ ] Team lacks ML expertise (resource gap)
- [ ] "Quick MVP" approach creates technical debt
- [ ] No clear sequence for implementation

---

## Test Content: "User Activity Dashboard Proposal"

### Summary

We should add a user activity dashboard to help users understand their usage patterns. This is a straightforward feature that builds on our existing auth system.

### Requirements

1. Display user's session history (last 30 days)
2. Show usage trends with simple charts
3. Provide smart recommendations based on usage patterns
4. Allow export to CSV

### Implementation Notes

This leverages our existing auth system for user identification. We'll store activity data alongside existing user records and surface it through a new dashboard view.

For the recommendations, we'll use a simple ML model to identify patterns and suggest optimizations. This should be quick to implement since we already have the data.

### Timeline

This feels like a 2-week project. Week 1 for backend, week 2 for frontend and polish.

---

## Flaw Details (for scoring verification)

**Flaw 1 - Dependency clarity (PARTIAL)**: Auth system mentioned, but the proposal doesn't acknowledge that "store activity data alongside existing user records" implies database schema changes, migrations, and potentially significant storage growth. These dependencies are unstated.

**Flaw 2 - Complexity estimate (HIGH)**: "Simple charts" and "straightforward feature" language obscures real complexity. Activity tracking at scale, data aggregation for charts, ML model for recommendations — each is substantial work.

**Flaw 3 - Resource alignment (GAP)**: "Simple ML model" assumes ML expertise exists on the team. For a team without ML experience, this is a major blocker or requires bringing in new skills.

**Flaw 4 - Technical debt (CONCERNING)**: "Store activity data alongside existing user records" is a quick solution that couples activity data to user records. This creates maintenance burden when activity data needs different retention, archival, or access patterns.

**Flaw 5 - Build path (FOGGY)**: What comes first? Data collection? Storage? Dashboard UI? ML model? The proposal doesn't sequence the work, making "2 weeks" impossible to evaluate.
