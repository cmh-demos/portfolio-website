# GCP Security Compliance Auditor

## Project Overview
Develop a Python-based tool (available as a CLI script, web app, or Cloud Function) that scans GCP resources for security compliance against standards like CIS GCP Foundations Benchmark, NIST, or custom rules. It checks configurations (e.g., IAM policies, encryption, VPC settings), generates detailed reports with recommendations, applies automated fixes where safe, and sends alerts via email/SMS. Designed for DevOps teams to maintain security posture.

## Why it fits your skills
- **Cloud Security**: Deep expertise in GCP security (IAM, encryption, SSO) from your SRE roles (e.g., implementing vaults and decommissioning shared passwords).
- **Programming & Automation**: Python with google-cloud libraries (e.g., google-api-python-client) for GCP API interactions; BASH for CLI wrappers.
- **DevOps Tools**: Ansible for applying configuration fixes; Terraform for setting up auditor infrastructure.
- **Monitoring & Alerts**: Integrates with Cloud Logging/Cloud Monitoring and Pub/Sub or Cloud Functions for notifications.

## Key Features
1. **Comprehensive Scans**: Audit Compute Engine, Cloud Storage, Cloud SQL, IAM, VPC, and more for vulnerabilities (e.g., open ports, weak passwords).
2. **Custom Rules Engine**: Define rules via JSON/YAML for flexibility.
3. **Reporting & Recommendations**: Generate HTML/PDF reports with severity levels and fix scripts.
4. **Automated Fixes**: Safe, reversible changes (e.g., enabling encryption) with approval workflows.
5. **Alerts & Scheduling**: Real-time alerts for critical issues; cron/Cloud Functions for periodic scans.
6. **Web Dashboard**: Optional UI for viewing reports and managing scans.
7. **Multi-Cloud Plugin**: Architect the rules engine so new modules for Azure/GCP can be plugged in easily.
8. **Compliance-as-Code**: Integrate with GitHub Actions to scan infrastructure pull requests and generate report comments.
9. **Drift Detection ML**: Use simple machine‑learning models to spot unusual configuration drift over time.
10. **Policy-as-Code Integration**: Show examples of using OPA or GCP Policy Controller/Cloud Asset Inventory alongside the auditor.

## Tech Stack
- **Core**: Python (with boto3, Click for CLI); libraries like Policy Sentry for IAM analysis.
- **Frontend (Optional)**: JavaScript (React) for web interface.
- **Infrastructure**: Cloud Functions (serverless scans), Cloud Storage (report storage), Pub/Sub/Cloud Monitoring (alerts).
- **Tools**: Ansible (fixes), Terraform (IaC), Git (version control), Docker (packaging).

## Implementation Steps
1. **Planning & Setup**: Research compliance frameworks; set up GCP test environment with Terraform.
2. **Core Auditor Development**: Build Python scanner with boto3; define initial rules.
3. **Reporting & Fixes**: Add report generation and automated fix logic.
4. **CLI/Web Interface**: Create command-line tool and optional web app.
5. **Testing & Deployment**: Test on mock GCP resources; deploy as Cloud Function or container.
6. **Integration**: Add monitoring; include in portfolio with sample reports.

## Why This Project?
- Highlights your security focus and automation skills, aligning with achievements like improving security posture and implementing monitoring.
- Practical for real-world use in regulated industries; demonstrates proactive compliance.
- Low-cost to run (Cloud Functions-based); estimated time: 2-3 weeks.
- Can be open-sourced or used as a consulting tool.
- Consider highlighting a concrete example, such as automatically remediating an insecure Cloud Storage bucket, to show the value you brought to an operation.