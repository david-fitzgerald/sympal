# Seren Test Input A

**Purpose**: Test Seren's code quality assessment. This code contains deliberate craft problems.
**Expected Score**: 5/10

## Deliberate Flaws

1. **Clarity**: FLAWED (1/2) — Poor naming, unclear intent
2. **Maintainability**: FLAWED (1/2) — Tightly coupled, hard to modify
3. **Consistency**: PASS (2/2) — Internally consistent (just bad)
4. **Simplicity**: FLAWED (1/2) — Unnecessary complexity
5. **Documentation**: PASS (0/2) — No comments, not self-documenting

---

## Test Content: "User Service Module"

```javascript
// user_service.js

const db = require('./db');
const cache = require('./cache');

async function doUserStuff(x, op, extra) {
  let result;

  if (op === 1) {
    // get user
    let cached = await cache.get('user_' + x);
    if (cached) {
      result = JSON.parse(cached);
    } else {
      result = await db.query('SELECT * FROM users WHERE id = ' + x);
      if (result) {
        await cache.set('user_' + x, JSON.stringify(result), 3600);
      }
    }
  } else if (op === 2) {
    // create user
    result = await db.query("INSERT INTO users (name, email) VALUES ('" + extra.name + "', '" + extra.email + "')");
    await cache.del('user_list');
  } else if (op === 3) {
    // delete user
    await db.query('DELETE FROM users WHERE id = ' + x);
    await cache.del('user_' + x);
    await cache.del('user_list');
    result = true;
  } else if (op === 4) {
    // list users
    let cached = await cache.get('user_list');
    if (cached) {
      result = JSON.parse(cached);
    } else {
      result = await db.query('SELECT * FROM users');
      await cache.set('user_list', JSON.stringify(result), 600);
    }
  }

  return result;
}

module.exports = { doUserStuff };
```

---

## Flaw Details (for scoring verification)

**Flaw 1 - Clarity (1/2)**:
- `doUserStuff` tells nothing about what it does
- `x`, `op`, `extra` are meaningless parameter names
- `op` uses magic numbers (1, 2, 3, 4) instead of named constants or separate functions
- Intent requires reading implementation to understand

**Flaw 2 - Maintainability (1/2)**:
- Single function does 4 different operations (violates Single Responsibility)
- Cache logic duplicated in multiple branches
- Adding a new operation means modifying this function
- SQL string concatenation (also a security issue, but craft-wise: hard to maintain)

**Flaw 3 - Consistency (2/2)**:
- PASS: The code is internally consistent in its badness
- Uses same patterns throughout (even if those patterns are poor)

**Flaw 4 - Simplicity (1/2)**:
- Unnecessary complexity from cramming 4 operations into one function
- Branching logic could be eliminated by having separate functions
- Cache key construction duplicated

**Flaw 5 - Documentation (0/2)**:
- Comments exist but just restate what's obvious from context
- No explanation of why cache TTLs differ (3600 vs 600)
- No documentation of parameters or return values
- Code is not self-documenting due to poor naming

### Key Issues Seren Should Identify

- [ ] Function name `doUserStuff` is meaningless
- [ ] Magic numbers for operations (1, 2, 3, 4)
- [ ] Single function violates Single Responsibility
- [ ] SQL string concatenation (maintenance + security)
- [ ] Duplicated cache logic
- [ ] Parameter names `x`, `op`, `extra` are unclear
