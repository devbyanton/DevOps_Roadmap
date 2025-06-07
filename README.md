# 🛠️ DevOpsRoadmap

Welcome to **DevOpsRoadmap** — a curated, hands-on journey documenting my growth and exploration across the DevOps landscape. This repository serves as a journal of practical projects I’ve built to deepen my knowledge of infrastructure, automation, observability, and backend engineering with a DevOps mindset.

## 🚀 Milestone 1: HealthChecker – My First Go Project

The journey begins with [**GoHealthChecker**]([https://github.com/YourUsername/GOHealthChecker](https://github.com/devbyanton/GoHealthChecker)) – a simple yet complete health monitoring system written in Go.

### 🔍 What it does:
- **Agent**: Installed on Linux or Windows machines, the agent collects basic system health metrics (e.g., CPU, memory) and sends them to a central API.
- **API Server**: Exposes endpoints to receive, store, and query health reports.
- **Self-hosted Binary Downloads**: The server provides download links for precompiled agent binaries tailored for Ubuntu and Windows.
- **Systemd + Windows Service**: Scripts automate deployment and installation of the agent as a service on both platforms.

### 🧰 Technologies:
- Language: **Go**
- Packaging: Cross-platform Go builds
- Deployment: `systemd`, Windows Services
- HTTP & REST APIs
- Local Dev: `curl`, PowerShell, Bash

### 📝 Learnings:
- Go module organization
- Cross-compilation (Windows/Linux)
- Basic service orchestration
- HTTP handlers and file serving
- CLI tooling and automation scripts

---

## 🔁 Milestone 2: Python Rewrite (Coming Soon)

To strengthen my understanding and compare paradigms, the next milestone is a **Python implementation** of the same HealthChecker system.

### Planned Goals:
- Recreate agent and API in Python (possibly using FastAPI)
- Use `psutil` for system metrics
- Package agent with `pyinstaller` for cross-platform binaries
- Use `systemd` and `win32serviceutil` for background service setup
- Compare developer experience, performance, and deployment complexity with Go

---

## 🧭 Future Roadmap

This repository will grow to include:

- CI/CD pipelines (GitHub Actions, Jenkins)
- Dockerizing and container orchestration with Kubernetes
- Monitoring (Prometheus, Grafana)
- Configuration management (Ansible, Terraform)
- Cloud integrations (AWS/GCP)
- Secrets and credential management
- Event-driven automation

---

## 🙌 Contributions & Feedback

This is a personal learning journey, but suggestions, improvements, and feedback are always welcome!

---

> 📍 *Follow the commits and subfolders to trace each milestone in this DevOps adventure.*  
> 🧑‍💻 *Everything is built from scratch with learning in mind, not production polish.*

