# 📝 To-Do CLI Application in Go

A simple and testable **Command Line To-Do List App** written in Go using clean architecture. It supports adding, listing, completing, and deleting tasks directly from your terminal.

This app is built for learning and demonstrating:

- Go programming best practices
- Modular code structure
- Unit testing
- File-based persistence
- CLI design with Cobra (planned)

---

## 🚀 Features

- ✅ Add new tasks
- 📋 List all tasks (sorted by ID)
- ✅ Mark tasks as completed
- ❌ Delete tasks
- 💾 Persist tasks to local JSON file
- 🧪 Unit tested using Go’s `testing` package

---

## 🧱 Tech Stack

| Component    | Tech                       |
| ------------ | -------------------------- |
| Language     | Go 1.21+                   |
| CLI Parser   | Cobra _(coming soon)_      |
| File Storage | Local JSON File            |
| Testing      | Built-in `testing` package |
| CI/CD        | GitHub Actions (Go CI)     |

---

## 📁 Project Structure

```
to_do_cli/
├── cmd/                 # (optional) Cobra CLI commands
├── todo/                # Core business logic and models
│   ├── task.go
│   ├── manager.go
│   ├── storage.go
│   ├── *_test.go
├── main.go              # Entry point (delegates to cmd or logic)
├── tasks.json           # File for persisting task data
├── go.mod               # Go module definition
├── README.md            # Project overview and usage
└── .github/workflows/   # GitHub Actions CI pipeline
```

---

## 🧪 Running Tests

```bash
go test ./...
```

Or run with verbose output:

```bash
go test -v ./...
```

---

## 🤝 Contributing

Feel free to fork and contribute! This project is open for improvements, experimentation, and learning.

---

## 📄 License

MIT
