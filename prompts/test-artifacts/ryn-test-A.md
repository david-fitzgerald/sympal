# Ryn Test Input A

**Purpose**: Test Ryn's failure mode identification. This system description contains unstated failure modes.
**Expected**: Ryn should identify at least 4 of the 5 planted failure modes with correct severity ratings.

## Deliberate Failure Modes (Hidden)

1. **Token refresh race condition** — Critical severity (security)
2. **Cascade from auth service failure** — High severity (feature unavailable)
3. **Stale token in distributed cache** — High severity (inconsistent state)
4. **Database connection exhaustion** — High severity (outage)
5. **Memory leak from token accumulation** — Medium severity (degraded performance)

---

## Test Content: "Token-Based Auth System Design"

### Overview

Our authentication system issues JWT tokens on login. Tokens are validated on each request and cached in Redis for performance. The system handles ~10,000 concurrent users.

### Architecture

1. User logs in → Auth service validates credentials → Issues JWT (1-hour expiry)
2. Each request → Check Redis cache for token → If cached, return user context → If not, validate JWT and cache result
3. Token refresh → When token has <10 min remaining, automatically refresh in background
4. Logout → Delete token from Redis cache

### Performance Optimizations

- Redis cache holds validated tokens to avoid repeated JWT parsing
- Background refresh prevents users from being logged out mid-session
- Connection pooling to database for credential validation

### Security Notes

- JWTs signed with RS256
- Tokens include user ID, roles, expiry
- Refresh tokens not used; instead we background-refresh the access token

---

## Failure Mode Details (for verification)

**Failure 1 - Token refresh race condition (Critical)**:
Background refresh creates new token while old token still valid. If old token is in use, multiple valid tokens exist. If one is revoked (logout, permissions change), the other may still work. Security gap.

**Failure 2 - Auth service cascade (High)**:
If auth service fails, all login attempts fail. Additionally, any cache miss requires re-validation → also fails. Users with expired cache entries can't continue. System-wide impact.

**Failure 3 - Stale token in distributed cache (High)**:
If multiple app servers, logout deletes from one Redis connection but stale token may exist in other server's local cache or Redis replica lag. User "logged out" can still have valid session elsewhere.

**Failure 4 - Database connection exhaustion (High)**:
Connection pooling mentioned but not sized. With 10,000 concurrent users, spike in cache misses → all servers hit DB for credential validation → pool exhaustion → cascading failures.

**Failure 5 - Token accumulation (Medium)**:
"Background refresh" creates new tokens. What happens to old ones in cache? If not cleaned up, Redis accumulates stale tokens. Memory grows over time. Degraded performance eventually.
