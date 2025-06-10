


````markdown
# 🦂 ScorpionDB

ScorpionDB is a high-performance, fault-tolerant, in-memory key-value datastore optimized for geospatial data operations. Designed with speed and resilience in mind, it supports write-ahead logging, TTL-based expiration, and efficient batched disk persistence — all accessible through a clean gRPC API interface.

---

## 🚀 Features

- **In-Memory Speed**: Fast reads and writes using an efficient key-based in-memory store.
- **Geospatial Support**: Stores spatial points with optional metadata for advanced location-aware use cases.
- **WAL (Write-Ahead Logging)**: Ensures durability and crash recovery by logging all mutating operations before execution.
- **Crash Recovery & Rehydration**: Rebuilds memory state on startup using WAL replay.
- **TTL (Time-to-Live)**: Auto-expiry of keys based on configurable lifespan.
- **Batched Flushing**: Periodically writes in-memory updates to persistent storage in optimized batches.
- **gRPC API**: Schema-defined service layer enabling fast, typed interaction with clients.
- **Protobuf for Wire Protocol**: Ensures compact, efficient, and structured communication.

---

## 📦 Current Architecture

```plaintext
[API Layer - gRPC TS Client (Monorepo)]
          |
      [gRPC Server]
          |
    ----------------
   |                |
[Engine Layer]    [TTL Manager]
   |                |
[WAL System]     [Flush Scheduler]
   |
[Persistent Disk Storage]
````

---

## 🛠 Components (Built so far)

| Component           | Status         | Description                                                     |
| ------------------- | -------------- | --------------------------------------------------------------- |
| `store.go`          | ✅ Done         | In-memory data structure operations (Add, Get, Update, Delete). |
| `wal.go`            | ✅ Done         | Write-ahead logging for all mutating operations.                |
| `recovery.go`       | ✅ Done         | WAL-based replay on boot for crash consistency.                 |
| `ttl.go`            | ✅ Done         | Background goroutine checking and removing expired keys.        |
| `flusher.go`        | ✅ Done         | Periodic batched write-to-disk mechanism.                       |
| `engine.go`         | ✅ Done         | Core coordinator for internal logic execution.                  |
| `api.go`            | ✅ Done         | gRPC server implementation handling client calls.               |
| `proto/store.proto` | ✅ Done         | Protobuf schema defining RPC services and messages.             |
| `TS Client (API)`   | 🏗 In Progress | TypeScript-based gRPC client for user-friendly API usage.       |

---

## 📁 Monorepo File Structure

```bash
scorpiondb/
│
├── engine/             # Core Go backend
│   ├── store.go
│   ├── wal.go
│   ├── ttl.go
│   ├── flusher.go
│   ├── recovery.go
│   └── engine.go
│
├── proto/              # gRPC Protobuf Definitions
│   └── store.proto
│
├── api/                # gRPC Server Binding Layer
│   └── api.go
│
├── ts-client/          # gRPC Client in TypeScript (API Layer)
│   ├── index.ts
│   └── generated/      # Generated gRPC client from .proto
│
├── main.go             # Entry point
└── README.md
```

---

## 📡 gRPC API Overview

```protobuf
service GeoStore {
  rpc AddPoint(AddRequest) returns (AddResponse);
  rpc GetPoint(GetRequest) returns (GetResponse);
  rpc UpdatePoint(UpdateRequest) returns (UpdateResponse);
  rpc DeletePoint(DeleteRequest) returns (DeleteResponse);
}
```

---

## 🛣️ Roadmap

* [x] WAL logging and recovery
* [x] TTL-based key expiry
* [x] Batch flushing to disk
* [x] Core Engine and In-memory logic
* [x] gRPC interface (Go)
* [x] Protobuf schema
* [ ] TypeScript API Client (gRPC-Web or Node)
* [ ] Snapshot + WAL truncation mechanism (for large datasets)
* [ ] Indexing/Spatial querying (Future scope)

---

## 🧪 Running Locally

> Requirements: Go 1.20+, `protoc`, `buf`, Node.js (for TS client)

```bash
# From root
go run main.go
```

TS Client setup instructions coming soon.

---

## 📖 License

This project is licensed under the MIT License.

---

## 👤 Author

**Shreevesh Kumar**
*Backend Developer | Cloud + DevOps Enthusiast | Builder of fast, real-time systems*

---

```

---

