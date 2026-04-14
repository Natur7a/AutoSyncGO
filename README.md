# AutoSyncGO
## Overview

This project is a lightweight data synchronization system that replicates changes between two SQL Server databases using an incremental, event-driven approach.

The goal is to build a practical backend system that handles real-world concerns such as data consistency, retries, and failure handling — not just basic CRUD operations.

---

## Problem

In distributed systems, data often needs to be synchronized across multiple databases. Full replication is inefficient, and manual syncing is error-prone.

This project solves that by:

* Tracking only changed data (incremental sync)
* Publishing changes in a Pub/Sub style
* Applying updates reliably to a target database

---

## Scope (v1)

To ensure completion, this version is intentionally constrained:

* Source: SQL Server
* Target: SQL Server
* Single table sync (initially)
* Incremental sync using timestamp or version column
* Simple conflict resolution: **Last Write Wins**
* Basic retry mechanism
* Structured logging

---

## Architecture

### Components

1. **Change Detector**

   * Queries source DB for updated rows since last sync

2. **Publisher**

   * Sends detected changes into a channel (in-memory queue)

3. **Consumer**

   * Reads changes and applies them to the target DB

4. **Sync Engine**

   * Orchestrates detection → publish → consume loop

---

## Data Flow

1. Detect new/updated records:

   ```sql
   SELECT * FROM Table WHERE UpdatedAt > @lastSync
   ```

2. Publish changes into queue

3. Consumer processes each change:

   * Insert or update in target DB

4. Update last sync timestamp

---

## Tech Stack

* Language: Go
* Database: SQL Server
* DB Driver: `database/sql` + SQL Server driver
* Concurrency: goroutines + channels
* Logging: standard `log` (upgrade later if needed)

---

## Core Concepts

### Incremental Sync

Only sync records that changed since last run using:

* `UpdatedAt` timestamp OR
* `RowVersion` (preferred for accuracy)

---

### Conflict Resolution

**Last Write Wins**

* The most recent update (based on timestamp/version) overrides previous data

---

### Retry Mechanism

* Retry failed operations
* Basic exponential backoff
* Prevents transient failures from breaking sync

---

### Logging

Track:

* Sync start/end
* Number of records processed
* Errors and retries

---

## Project Structure

```
/cmd
    main.go

/internal
    /sync
        engine.go
        detector.go
        publisher.go
        consumer.go

    /db
        connection.go

    /model
        change.go
```

---

## Example Flow (Simplified)

```go
changes := detectChanges(lastSync)
publish(changes, queue)
consume(queue, targetDB)
```

---

## Future Improvements

* Multi-table support
* Batch processing
* Idempotency handling (avoid duplicate writes)
* Dead-letter queue for failed records
* Metrics dashboard
* External message broker (Kafka / RabbitMQ)

---

## Why This Project Matters

This project demonstrates:

* Handling real-world data synchronization problems
* Designing event-driven systems
* Managing failure, retries, and consistency
* Writing concurrent systems in Go

---

## Constraints & Philosophy

* Keep v1 simple and shippable
* Avoid over-engineering
* Focus on correctness over features
* Finish before expanding

---

## Status

🚧 In Development

---

## Notes

This project is intentionally scoped to ensure completion.
Future iterations may expand functionality once a stable core is built.
