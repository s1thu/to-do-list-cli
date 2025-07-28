# 📝 To-Do CLI Application in Go

A simple and testable **Command Line To-Do List App** written in Go using clean architecture. It supports adding, listing, completing, and deleting tasks directly from your terminal.

This app is built for learning and demonstrating:

- Go programming best practices
- Modular code structure
- Unit testing
- File-based persistence
- CLI design with Cobra

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
| CLI Parser   | Cobra                      |
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

## 📖 Installation & Usage

### Installation

1. **Clone the repository:**
   ```bash
   git clone git@github.com:s1thu/to-do-list-cli.git
   cd to-do-list-cli
   ```

2. **Build the application:**
   ```bash
   go build -o todo
   ```

3. **Run directly with Go:**
   ```bash
   go run main.go [command]
   ```

### Available Commands

#### 🆘 Get Help
```bash
# Show all available commands
todo --help

# Get help for a specific command
todo [command] --help
```

#### ➕ Add Tasks
```bash
# Add a new task
todo add "Buy groceries"
todo add "Complete project documentation"
todo add "Call dentist for appointment"
```

#### 📋 List Tasks
```bash
# Show all tasks
todo list
```
**Output example:**
```
[ ] 1: Buy groceries (Created: 2025-01-15 10:30) (Updated: 2025-01-15 10:30)
[x] 2: Complete project documentation (Created: 2025-01-15 09:15) (Updated: 2025-01-15 11:20)
[ ] 3: Call dentist for appointment (Created: 2025-01-15 14:45) (Updated: 2025-01-15 14:45)
```

#### ✅ Complete Tasks
```bash
# Mark task as completed by ID
todo complete 1
todo complete 3
```

#### ❌ Delete Tasks
```bash
# Delete task by ID
todo delete 2
todo delete 5
```

### Usage Examples

```bash
# Basic workflow
todo add "Learn Go programming"     # Add task
todo add "Build CLI application"    # Add another task
todo list                           # View all tasks
todo complete 1                     # Mark first task as done
todo delete 2                       # Remove second task
todo list                           # Check final state
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
