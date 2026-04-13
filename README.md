# 🚀 DevOps Platform

A full-featured DevOps platform demonstrating modern infrastructure practices using **Docker, Kubernetes, CI/CD, Monitoring, and Blue-Green deployments**.

---

# 📌 Overview

This project simulates a production-like environment:

* Containerized Go application
* Kubernetes deployment with **Ingress**
* **Blue-Green deployment strategy**
* **Monitoring & Observability** (Prometheus + Grafana)
* **Alerting system**
* **CI/CD pipeline (GitHub Actions)**
* Infrastructure as Code (Kustomize, YAML)

---

# 🏗️ Architecture

```
User → Ingress (NGINX) → Service → Pods (Blue/Green)
                           ↓
                    Prometheus → Grafana
```

---

# ⚙️ Tech Stack

* Go (HTTP service)
* Docker
* Kubernetes (kind)
* Kustomize (dev/prod environments)
* Prometheus (metrics)
* Grafana (visualization)
* GitHub Actions (CI/CD)
* Ansible (provisioning)

---

# 📂 Project Structure

```
.
├── go-service/          # Go application
├── docker/              # Dockerfile
├── k8s/                 # Kubernetes manifests
│   ├── base/
│   ├── blue-green/
│   └── overlays/
├── monitoring/
│   ├── grafana/
│   └── prometheus/
├── scripts/             # Automation scripts
├── ansible/             # Infrastructure provisioning
```

---

# 🚀 Features

## ✅ Application

* REST service written in Go
* Exposes `/metrics` endpoint

## ✅ Kubernetes

* Deployment + Service + Ingress
* Liveness & Readiness probes
* Namespace isolation

## ✅ Blue-Green Deployment

* Two environments: `blue` and `green`
* Traffic switching via Service selector
* Zero-downtime deployment

## ✅ Monitoring

* Prometheus collects metrics
* Grafana dashboard included
* Custom metrics:

  * `http_requests_total`
  * `http_request_duration_seconds`

## ✅ Alerting

* Alert rules for error rate
* Detects incidents automatically

## ✅ CI/CD

* Automated build & deploy via GitHub Actions

---

# 📊 Metrics

Example:

```
http_requests_total{method="GET", path="/", status="200"}
http_request_duration_seconds
```

---

# 🚨 Alerts

Example rule:

* High error rate (HTTP 500)
* Triggered if errors detected within time window

---

# 🛠️ Run Locally

## 1. Build Docker image

```
docker build -t devops-platform:latest -f docker/app/Dockerfile .
```

## 2. Load into kind

```
kind load docker-image devops-platform:latest --name devops-cluster
```

## 3. Deploy

```
./scripts/deploy.sh
```

## 4. Access app

```
curl http://devops.local
```

---

# 🔄 Blue-Green Switch

```
./scripts/blue-green-switch.sh
```

---

# 📈 Monitoring

* Open Grafana
* Use included dashboard
* Explore metrics:

```
rate(http_requests_total[1m])
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[1m]))
```

---

# 💡 What This Project Demonstrates

* Real-world DevOps workflow
* Infrastructure as Code
* Observability & monitoring
* Deployment strategies
* Debugging and incident handling

---

# 🎯 Author

DevOps Engineer in progress 🚀
Focused on Kubernetes, CI/CD, and production systems.

---

# ⭐ Notes

This project is designed to simulate a **real production environment** and demonstrate practical DevOps skills.

