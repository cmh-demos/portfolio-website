# Cloud Infrastructure Monitoring Dashboard

## Project Overview
A web-based dashboard that monitors and visualizes GCP infrastructure health (e.g., Compute Engine instances, Cloud Storage buckets, Cloud Functions). It pulls real-time data via GCP APIs, displays metrics like CPU usage, costs, and alerts, and includes automation for scaling or backups.

## Why it fits your skills
- **Cloud & Infrastructure**: Leverages GCP services (Compute Engine, Cloud Monitoring, Cloud Functions, Cloud Storage) that you're expert in.
- **Programming**: Python backend (using google-cloud libraries for GCP integration), JavaScript frontend for interactive charts (e.g., with Chart.js).
- **Automation & DevOps**: Terraform for infrastructure as code (IaC) to provision resources; Ansible for configuration; Docker for containerization; CI/CD pipeline (e.g., GitHub Actions or Cloud Build/Cloud Deploy) for automated deployment.
- **Monitoring**: Integrates Prometheus/Grafana-like features using Cloud Monitoring.

## Key Features
1. **Real-Time Metrics Display**: Dashboard with graphs for Compute Engine CPU/memory, Cloud Storage usage, and Cloud Function invocations.
2. **Alerts & Automation**: Email/SMS notifications (via SendGrid or Pub/Sub triggers) for anomalies; auto-scaling triggers.
3. **Interactive UI**: Filter by region/service; responsive design with purple theme.
4. **Security**: IAM roles, encrypted data, and best practices from your experience.
5. **Deployment**: Fully automated setup on GCP (e.g., GKE cluster or Cloud Run endpoints), accessible via a public URL.
6. **Plugin Architecture**: Design the UI and backend to accept custom widgets or data sources so others can extend the dashboard.
7. **Cost Optimization Module**: Analyze Cloud Monitoring logs to suggest rightsizing or identify waste for reduced bills.
8. **Offline Support**: Use a Service Worker so the dashboard continues working during outages or limited connectivity.
9. **Kubernetes Monitoring**: Add support for EKS/GKE clusters to demonstrate multi‑cloud readiness.

## Tech Stack
- **Backend**: Python (Flask/FastAPI) with google-cloud libraries.
- **Frontend**: HTML/CSS/JS (building on your current site) + Chart.js for visualizations.
- **Infrastructure**: GCP (Compute Engine/GKE or Cloud Run, Cloud SQL or BigQuery for any data, Cloud Monitoring).
- **Tools**: Terraform (IaC), Docker (containers), Git (version control), Ansible (config management).

## Implementation Steps
1. **Planning & Setup**: Define GCP resources; set up a Git repo.
2. **Backend Development**: Python API to fetch GCP data.
3. **Frontend Integration**: Build the dashboard UI.
4. **Automation**: Write Terraform scripts for deployment; add CI/CD.
5. **Testing & Deployment**: Test locally with Docker, deploy to GCP.
6. **Integration**: Link it to your portfolio site (e.g., as a "Projects" section with a demo link).

## Why This Project?
- It's scalable and demonstrates end-to-end DevOps (from code to cloud).
- You can host it on GCP for free/low-cost (using free tier).
- It directly ties into your resume achievements (e.g., maintaining 99.99% SLA, automation pipelines).
- Estimated time: 2-4 weeks, depending on depth.
- Add a note about demonstrating specific challenges (e.g., troubleshooting a runaway cost spike) to tie back to your resume.