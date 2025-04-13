# gRPC Exam Service 📚

![Banner](public/banner_grpc.png)

A fun, modular gRPC service for fetching exam results using different gRPC communication types:
- Unary RPC
- Server Streaming
- Client Streaming
- Bidirectional Streaming

Cleanly structured for scalability and learning.

---

## 🌐 Project Structure

```bash
.
├── client/
│   ├── clients/
│   │   ├── unary.go
│   │   ├── server_stream.go
│   │   ├── client_stream.go
│   │   ├── bi_stream.go
│   └── main.go
├── proto/
│   ├── exam.proto
│   └── generated/exampb/
│       ├── exam.pb.go
│       └── exam_grpc.pb.go
├── server/
│   ├── main.go
│   └── servers/
│       ├── unary.go
│       ├── server_stream.go
│       ├── client_stream.go
│       ├── bi_stream.go
│       └── exam_service_server.go
├── utils/
│   └── logger.go
├── go.mod
├── go.sum
└── Makefile
```

---

## 📚 Features

| Type               | File                                  | Description                                |
|--------------------|---------------------------------------|--------------------------------------------|
| Unary              | `unary.go`                            | Fetch a single exam result                 |
| Server Streaming   | `server_stream.go`                    | Fetch multiple results for a student       |
| Client Streaming   | `client_stream.go`                    | Send multiple requests, get summary        |
| Bidirectional      | `bi_stream.go`                        | Live querying of student exam results      |

---

## ⚙️ Usage

### 1. Generate gRPC code
```bash
make proto -B
```

### 2. Run the server
```bash
go run server/main.go
```

### 3. Run the client
```bash
# Unary RPC
go run client/main.go unary

# Server Streaming
go run client/main.go server

# Client Streaming
go run client/main.go client

# Bidirectional Streaming
go run client/main.go bi
```

---

## 🔧 Tools Used
- gRPC & Protocol Buffers
- Go Modules
- slog (structured logging)

---

## 🚀 Quick Example (Bi-Directional)
```bash
Enter student_id and exam_id (or 'exit'): 123 math101
🎓 John Doe | Math 101: 95/100 (A+)

Enter student_id and exam_id (or 'exit'): exit
👋 Session ended.
```

---

## ✨ Have fun learning gRPC with this mini project!

