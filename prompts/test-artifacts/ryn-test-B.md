# Ryn Test Input B

**Purpose**: Test Ryn v1.1's expanded capabilities (security + testing). Contains deliberate flaws.
**Expected**: Identify 4 security issues + 4 test coverage gaps

## Deliberate Security Issues

1. **SQL Injection** — User input directly in query (Critical)
2. **Broken Access Control** — No ownership verification (Critical)
3. **Sensitive Data Exposure** — API returns password hash (High)
4. **Missing Rate Limiting** — Enumeration possible (Medium)

## Deliberate Test Gaps

1. **No auth tests** — Access control untested
2. **No injection tests** — SQL injection untested
3. **No edge case tests** — Boundary conditions missing
4. **No integration tests** — Component interaction untested

---

## Test Content: "User Profile API"

### Endpoint Design

```
GET /api/users/{id}/profile
POST /api/users/{id}/profile
```

### Implementation

```python
# profile_api.py

from flask import Flask, request, jsonify
import sqlite3

app = Flask(__name__)

def get_db():
    return sqlite3.connect('users.db')

@app.route('/api/users/<user_id>/profile', methods=['GET'])
def get_profile(user_id):
    db = get_db()
    cursor = db.cursor()

    # Get user profile
    cursor.execute(f"SELECT id, name, email, password_hash, created_at FROM users WHERE id = {user_id}")
    row = cursor.fetchone()

    if row:
        return jsonify({
            'id': row[0],
            'name': row[1],
            'email': row[2],
            'password_hash': row[3],
            'created_at': row[4]
        })
    return jsonify({'error': 'User not found'}), 404

@app.route('/api/users/<user_id>/profile', methods=['POST'])
def update_profile(user_id):
    db = get_db()
    cursor = db.cursor()

    data = request.json
    cursor.execute(f"UPDATE users SET name = '{data['name']}', email = '{data['email']}' WHERE id = {user_id}")
    db.commit()

    return jsonify({'status': 'updated'})
```

### Test Suite

```python
# test_profile.py

def test_get_profile_returns_user():
    response = client.get('/api/users/1/profile')
    assert response.status_code == 200
    assert 'name' in response.json

def test_update_profile_changes_name():
    response = client.post('/api/users/1/profile', json={'name': 'New Name', 'email': 'test@test.com'})
    assert response.status_code == 200
```

---

## Flaw Details (for verification)

**Security Issue 1 - SQL Injection (Critical)**:
`f"SELECT ... WHERE id = {user_id}"` — direct string interpolation allows injection.
Attack: `GET /api/users/1 OR 1=1--/profile`

**Security Issue 2 - Broken Access Control (Critical)**:
No check that requesting user owns the profile. User A can read/modify User B's profile.
Attack: Enumerate user IDs, access any profile.

**Security Issue 3 - Sensitive Data Exposure (High)**:
`password_hash` returned in API response. Should never be exposed.

**Security Issue 4 - Missing Rate Limiting (Medium)**:
No rate limiting on GET endpoint allows user enumeration at scale.

**Test Gap 1 - No auth tests**:
Tests don't verify that users can only access their own profiles.

**Test Gap 2 - No injection tests**:
No tests for SQL injection or other input validation.

**Test Gap 3 - No edge case tests**:
No tests for: invalid user_id formats, missing fields, empty strings, very long inputs.

**Test Gap 4 - No integration tests**:
No tests for database transaction behavior, concurrent access, or error handling.
