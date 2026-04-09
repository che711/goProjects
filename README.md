# Go Projects Collection 🚀

This repository contains a collection of small to medium-sized projects written in **Go (Golang)**.
The goal of this repository is to practice core Go concepts, system design, and real-world development patterns.

Each directory represents a standalone project with its own logic, structure, and learning objectives.

---

## 📦 Projects Overview

| Project                 | Description                                       |
| ----------------------- | ------------------------------------------------- |
| `01-todo-list`          | CLI-based task manager                            |
| `02-backend-api`        | REST API service (calculator or similar)          |
| `03-web-scraper`        | Tool for scraping websites / detecting dead links |
| `04-url-shortener`      | URL shortener with backend + web interface        |
| `05-currency-converter` | CLI currency converter using external API         |

---

## 🧠 Goals

* Learn Go fundamentals (goroutines, channels, interfaces)
* Practice building CLI and web applications
* Understand REST API design
* Work with external APIs
* Improve project structure and modularization

---

## 🛠️ Tech Stack

* Go (Golang)
* Standard library
* Optional:

  * HTTP frameworks (`net/http`, `chi`, `gin`)
  * Third-party APIs
  * CLI libraries

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/che711/goProjects.git
cd goProjects
```

### 2. Initialize Go module (if needed)

```bash
go mod init github.com/che711/goProjects
go mod tidy
```

### 3. Run a project

Navigate to a project folder:

```bash
cd 01-todo-list
go run main.go
```

---

## 📂 Project Structure

```
goProjects/
│
├── 01-todo-list/
├── 02-backend-api/
├── 03-web-scraper/
├── 04-url-shortener/
├── 05-currency-converter/
│
└── README.md
```

Each project may contain:

* `main.go` — entry point
* `internal/` or `pkg/` — core logic
* `api/` — handlers (for backend services)
* `cmd/` — CLI commands

---

## 🧪 Ideas for Improvements

* Add tests (`go test ./...`)
* Dockerize services
* Add CI/CD (GitHub Actions)
* Improve logging & configuration
* Add database integration (PostgreSQL, Redis)

---

## 📖 Learning Resources

* https://go.dev/doc/
* https://gobyexample.com/
* https://pkg.go.dev/

---

## 🤝 Contributing

Feel free to fork this repository and experiment with your own ideas.

---

## 📄 License

This project is licensed under the MIT License.
