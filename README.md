# go-auth-service

[![CI](https://github.com/Eyob49/go-auth-service/actions/workflows/ci.yml/badge.svg)](https://github.com/Eyob49/go-auth-service/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Eyob49/go-auth-service)](https://goreportcard.com/report/github.com/Eyob49/go-auth-service)
[![Go version](https://img.shields.io/badge/go-1.26.1-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Simple Go authentication service providing register, login and a protected profile endpoint.

Features
- Register and login with JWT-based auth
- Protected /profile endpoint
- /health liveness endpoint
- Graceful shutdown and DB init

API Endpoints
- POST /register  — register new user (JSON body)
- POST /login     — login and receive JWT
- GET  /profile   — protected; requires Authorization: Bearer <token>
- GET  /health    — returns 200 OK when healthy

Quickstart
1. Copy .env.example to .env and update values (do not commit secrets):

   copy .env.example .env

2. Run locally:

   go run ./cmd

3. Or build a binary:

   go build -o bin/auth ./cmd
   ./bin/auth

Environment
Required env vars:
- DATABASE_URL (e.g. postgres://user:pass@localhost:5432/authdb)
- JWT_SECRET
- PORT (optional; default 8080)

Tests
Run unit tests (none currently):

go test ./... -v

CI
A GitHub Actions workflow is included at `.github/workflows/ci.yml` to run fmt, vet, tests and go mod tidy.

Contributing
- Open issues or PRs. Add tests for new behavior.

License
MIT — see LICENSE
