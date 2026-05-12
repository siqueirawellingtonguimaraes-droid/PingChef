# 🍳 PingChef

> Lightweight Go-based CLI for real-time HTTP service monitoring.

---

# 🧠 About the Project

PingChef is a lightweight CLI tool for service health checks.

At this stage, the project includes endpoint management and event persistence, along with a basic health-checking engine.

The long-term vision is to evolve this into a complete monitoring ecosystem.

---

# 📌 Current Status

* Persistence layer (endpoints and events stored in a database)
* SQLite integration implemented
* Basic CLI commands available (events and endpoints)
* No web platform yet

---

# ⚙️ Stack

* Go (Golang)
* Cobra CLI (in progress)
* net/http (health checks)
* Future: Web dashboard and dedicated production database

---

# 🧪 Current Feature

## ➕ Add Route

### Command

```bash
go run .\cmd\main.go route add --name api --url https://api.example.com/health --interval 15
```

### Flags

| Flag         | Type     | Description                                 | Default  |
| ------------ | -------- | ------------------------------------------- | -------- |
| `--name`     | `string` | Friendly name used to identify the endpoint | required |
| `--url`      | `string` | HTTP endpoint to be monitored               | required |
| `--interval` | `int`    | Interval between health checks in seconds   | `10`     |

---

## ➕ List Routes

### Command

```bash
go run .\cmd\main.go route list
```

---

## ➕ Run All Endpoints

### Command

```bash
go run .\cmd\main.go event run-all
```

---

# 🧠 Design Intent

Even in its early stage, PingChef is designed with scalability and real-world infrastructure use cases in mind.

The architecture is being prepared for:

* Modular CLI structure
* Concurrent monitoring engine
* Persistent storage layer
* Future distributed/web architecture

---

# 🚀 Planned Features

## 🌐 Web Platform (Future)

* Dashboard for visualization
* Endpoint management interface
* Uptime history tracking
* Alerting system
* Remote endpoint persistence


---

# 🧵 Engineering Focus

This project is intentionally built to explore real-world backend engineering concepts such as:

* Concurrency in Go (goroutines + channels)
* CLI architecture design
* System modularization
* Infrastructure tooling
* Monitoring systems fundamentals
* Gradual system evolution

---

# 👨‍💻 Final Notes

PingChef is not just a CLI tool — it is the foundation of a monitoring ecosystem built step by step.

The current stage focuses on validating the core interaction model before introducing persistence, distributed components, and a web interface.
