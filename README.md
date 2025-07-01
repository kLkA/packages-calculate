# Backend Task - Stage 1 Interview

## 📌 Project Description

This project is a demo API to calculate optimal set of packs to serve order amount for shipping

Key requirements:
1. Be written with Golang
2. UI for demo usage
3. Dynamic pack sizes
4. API Endpoint to calculate packs
5. Unit tests for algorithm


The goal was to demonstrate familiarity with backend development, code structure, and testability and readiness to write production-ready code.

---

## 🛠 How It Was Implemented

- **Language:** Go 1.24
- **Frameworks:** Gin router framework to simplify and speedup implementation
- **Architecture:** cmd,internal folders as go standard. Code arch is as per SOLID/DDD principles
- **Libraries:**
1. `oklog/ulid` - Universally Unique Lexicographically Sortable Identifier that is better for database keys, event logs, or where insert order or timestamp is valuable.
2. `mockery` - for generating mocks
3. `testify` - for asserting test cases
4. `zerolog` - for convenient logging in dev/prod modes
---


## ▶️ How to Run

`make start` or `make start-build` to start

`make stop` to stop docker compose

`make logs` to see app logs

`make docker-test` - run tests

## Tests coverage

Unit tests
```
ok  	homework/internal/pack/usecase	1.083s	coverage: 97.4% of statements
```