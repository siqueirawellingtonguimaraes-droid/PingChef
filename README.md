# 🍳 PingChef

> Lightweight monitoring CLI built in Go for tracking HTTP services in real time.

---

# 🧠 About The Project

PingChef is currently in **early development stage**, focused on building a minimal but powerful CLI tool for service health-checking.

At this stage, the project does not yet include persistence or a web platform. The focus is on validating the core idea: a fast, concurrent and simple monitoring CLI.

The long-term vision is to evolve this into a complete monitoring ecosystem.

---

# 📌 Current Status

🚧 Early Stage / MVP

The project is under active development and currently supports only basic endpoint registration.

⚠️ Limitations at this stage:

- No persistence (endpoints are not saved yet)
- No SQLite integration implemented
- No web platform
- No historical monitoring data
- Only basic CLI command implemented

---

# ⚙️ Stack

- Go (Golang)
- Cobra CLI (structure in progress)
- net/http (health checks)
- Future: SQLite + Web dashboard

---

# 🧪 Current Feature

## ➕ Add Route

The only implemented feature so far is registering endpoints for future monitoring.

### Command

```bash
"go run .\cmd\main.go route add --name api --url https://api.example.com/health --interval 15"
```
---

# 🧠 Design Intent

Even in its early stage, PingChef is being designed with scalability and real-world infrastructure use cases in mind.

The architecture is being prepared for:

- modular CLI structure
- concurrent monitoring engine
- persistent storage layer
- future distributed/web architecture

---

# 🚀 Planned Features

## 📦 Persistence Layer (SQLite)

- Store endpoints locally
- Persist CLI configuration
- Enable restart-safe monitoring

---

## 🔄 Monitoring Engine

- Continuous HTTP health checks
- Goroutine-based workers per endpoint
- Event-based status updates
- Detection of state changes (UP / DOWN / RECOVERY)

---

## 🌐 Web Platform (Future)

- Dashboard for visualization
- Endpoint management interface
- Uptime history tracking
- Alerting system (Slack, Discord, Telegram)

---

# 🏗️ Architecture Vision

The system is being designed to evolve progressively:

```text
CLI → Monitoring Core → SQLite → Web Platform
```

Each layer builds on top of the previous one, keeping the system modular and extensible.

---

# 🧵 Engineering Focus

This project is intentionally being built to explore real-world backend engineering concepts such as:

- concurrency in Go (goroutines + channels)
- CLI architecture design
- system modularization
- infrastructure tooling
- monitoring systems fundamentals
- gradual system evolution

---

# 👨‍💻 Final Notes

PingChef is not just a CLI tool — it is the foundation of a monitoring ecosystem being built step-by-step.

The current stage focuses on validating the core interaction model before introducing persistence, distributed components and a web interface.