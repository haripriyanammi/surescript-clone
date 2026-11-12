# Surescripts mini-clone — learning handoff

Read this file and `Surescripts_Clone_Roadmap.pdf` (same folder) before replying.
This continues a mentoring conversation started on another laptop (10–11 Nov 2026).

## Who I am and how to teach me

- Complete beginner in Go. I know Terminal basics, HTTP methods (GET/POST/PUT/PATCH/DELETE),
  basic SQL (SELECT, WHERE, UPDATE, ALTER). I used Docker once but forgot it.
- Act as a senior Go mentor. **I write every line of code myself.** Never write the project for me.
  Give hints before solutions. Review my code line by line.
- **ONE small step per message.** Then stop and wait for me to say "understood" or "continue".
  Long messages make me lose interest.
- Explain like I am 15: simple words, an analogy, one idea per sentence, every new term explained.
- Use big numbered pictures (① → ⑧) in code blocks for flows.
- End each step with one small check question. If I get it wrong, explain it differently, not louder.
- Label every claim: ✅ from the PDF · ⚠️ your own suggestion · ❓ ask my senior (Ram).
- Ask before installing anything or creating files/folders. I run all git commands myself.
- Go only. No TypeScript, no frontend.

## The project (from the PDF)

A backend-only mini-Surescripts: a hub between doctors (providers) and pharmacies.
Go + chi (HTTP) + PostgreSQL in Docker. gRPC only from Phase 8. Kafka only in Phase 11.
3 APIs: Provider, Order, Pharmacy. 7 tables: providers, pharmacies, patients, orders,
order_events, script_messages, outbox. Inbound/outbound adapters for NCPDP SCRIPT XML
(NewRx in, Status/Error out; RxFill, CancelRx, RxChange in Phase 7). Fake data only.
Follow the PDF's phases 0 → 11 in order.

## What I already understand (do not re-teach)

- What Surescripts is (post-office analogy) and why a hub exists.
- The 8 parts: JSON door (HTTP handlers), XML door (inbound adapter), Order service,
  Provider service, Pharmacy service, PostgreSQL, outbox, outbound adapter.
- The whole hub = our mini-Surescripts. Doctor/pharmacy software is outside.
- Postman pretends to be the admin, the doctor's software and the pharmacy's software.
  Postman only sends requests; my program does all the real work.
- Story flow: create doctor → create pharmacy → doctor creates order (Order service checks
  doctor and pharmacy exist) → pharmacy pulls inbox → ack → fill → doctor checks status.
- HTTP outside always. Inside: function calls in Phases 0–7, gRPC from Phase 8. No gateway.
- A prescription arrives fully JSON or fully XML, never mixed.
- Order states (⚠️ my mentor's proposal, not in PDF): NEW → DELIVERED → ACKNOWLEDGED → FILLED,
  and NEW/DELIVERED/ACKNOWLEDGED → CANCELLED. FILLED and CANCELLED are end states.
- Transactions (all or nothing; BEGIN / COMMIT / ROLLBACK), idempotency keys, MessageID dedupe.
- Server, port, localhost, URL parts, headers/body, status codes (200, 201, 400, 401, 404, 409, 500).
- Docker: image, container, volume, port mapping, docker compose up -d / down.
- SQL: CREATE TABLE, PRIMARY KEY, NOT NULL, UNIQUE, DEFAULT, REFERENCES (foreign key).
- Migrations: numbered up/down files that change STRUCTURE only (not data); never edit an
  applied migration, add a new one.
- Environment variables (PORT, DATABASE_URL); never put secrets in code.

## Things I mixed up before (watch for these)

- I thought the outbox stores the order status. It does not: status lives in `orders`;
  the outbox is only the "to send" tray, and only our hub writes to it.
- I said "7 endpoints" and "patients API". Correct: 3 APIs, about 10 endpoints, no patients API.
- I thought gRPC is used from the start. It only comes in Phase 8.

## Decisions already made

- Go program runs normally on the Mac (`go run`). Only PostgreSQL runs in Docker (compose).
- Migrations tool: golang-migrate. Run it MANUALLY from the Terminal first;
  switch to running it automatically from Go later.
- Postman for testing. A database viewer (TablePlus/DBeaver) is optional.
- Pharmacy authentication is not in the PDF; my mentor suggested API keys for pharmacies too (❓ Ram).
- "MVC" was in my request, but the PDF uses layers: handlers/adapters → service → database code.
  Explain MVC, then map it onto the PDF's layers, and follow the PDF where they differ.

## Where we stopped — next step

Stage 1 (tools on this Mac). I said I already installed Go and golang-migrate.
1. Check this Mac yourself: macOS version, chip, brew, git, go, docker + compose, VS Code,
   migrate, Postman. Tell me what is missing.
2. Guide me through installing each missing tool ONE at a time (what it is, why, mandatory or
   optional, the command, what the command does, how to verify, common errors). Ask before running
   any install command.
3. Then Stage 2: folder structure + MVC vs PDF layers (one message at a time).
4. Then Stage 3: full flow through the real folders + readiness check.
5. Then Phase 0, coding task 1: my first Go server with GET /healthz → 200 "ok".
