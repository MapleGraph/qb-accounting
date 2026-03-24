# QB Accounting

Enterprise-grade accounting module for the QueueBuster ERP ecosystem.

## Overview

The Accounting module is the **system of record for financial recognition and financial controls**. It owns the chart of accounts, fiscal years, accounting periods, journal engine, posting engine, source-document-to-ledger traceability, AR/AP open-item accounting, and financial reporting.

## Tech Stack

- **Language**: Go 1.25.1
- **Module**: `qb-accounting`
- **Framework**: Gin
- **Databases**: PostgreSQL, Redis, ClickHouse
- **Core**: [qb-core](https://github.com/MapleGraph/qb-core)
- **Docs**: Swagger / OpenAPI

## Getting Started

### Prerequisites

- Go 1.25+
- PostgreSQL 15+
- Redis 7+

### Setup

```bash
cp .env.example .env    # Configure environment
make db-setup           # Initialize database schema
make run                # Start the server (port 8086)
```

### Development

```bash
make dev                # Start with Docker dependencies
make swagger            # Regenerate Swagger docs
make test               # Run tests
make lint               # Run linter
```

## API Documentation

Swagger UI is available at: `http://localhost:8086/api/public/accounting/swagger/index.html`

## Architecture

```
cmd/                    # Entry point
internal/
  config/               # Configuration loading
  dto/                  # Request/response DTOs
  helpers/              # Environment helpers
  models/               # Database models
  proto/                # Generated protobuf code
  repository/
    postgres/           # PostgreSQL repositories
    remote/             # gRPC service clients
  services/             # Business logic
  transport/
    http/
      handlers/         # HTTP handlers with Swagger annotations
      routes/           # Route definitions
  utils/                # Response helpers, JSONB utilities
proto/                  # Protobuf source files
schema/                 # SQL schema
deployment/             # Docker files
docs/                   # Generated Swagger docs
```

## Key Entities

| Entity | Description |
|--------|-------------|
| Books | Accounting book master (primary, statutory, management) |
| Fiscal Years | Book-specific fiscal year lifecycle |
| Accounting Periods | Monthly period open/lock/close control |
| Account Groups | Hierarchical chart-of-accounts grouping |
| Accounts | Postable and header accounts with control flags |
| Voucher Sequences | Voucher numbering configuration |
| Journal Batches | Optional journal grouping layer |
| Journals | Core immutable accounting events |
| Journal Lines | Debit/credit entries with dimensions |
| Posting Rule Versions | Versioned automated posting rules |
| Posting Requests | Idempotent inbound posting envelopes |
| Posting Request Snapshots | Immutable source snapshots at posting time |
| Journal Source Links | Source-to-journal traceability |
| Open Items | Unified AR/AP open-item ledger |
| Open Item Allocations | Settlement and application flows |
| Open Item Adjustments | Write-offs and balance adjustments |
