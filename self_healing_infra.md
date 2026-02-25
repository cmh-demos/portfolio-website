# Self‑Healing Infrastructure Demo

## Project Overview
Assemble a small cloud environment that detects its own failures and automatically remediates them. For example, if a Compute Engine VM becomes unresponsive, a watchdog Cloud Function could replace it or roll back to a previous healthy configuration. The demo should exercise monitoring, alerting, and automated remediation together.

## Why it fits your skills
- Aligns with your SRE experience and incident‑response automation.
- Highlights creativity in automating recovery and reducing manual toil.
- Connects to your past roles where you built automation for launch campaigns and maintenance windows.

## Key Features
1. **Health Detection**: Cloud Monitoring alerts for instance/cluster health, latency, or service errors.
2. **Remediation Actions**: Cloud Functions or Ansible playbooks that recreate instances, roll back deployments, or restart services.
3. **Fallback Logic**: Maintain a small fleet of standby resources or snapshots for quick recovery.
4. **Audit Trail**: Log all detections and actions to a centralized store (e.g., Cloud Storage/Elasticsearch).
5. **Chaos Injection**: Optional script to deliberately break components to test the healing logic.
6. **Dashboard**: Visualize incidents and remediation outcomes over time.

## Tech Stack
- **Monitoring**: Cloud Monitoring, Prometheus.
- **Automation**: Cloud Functions, Ansible, Terraform.
- **Logging**: Cloud Monitoring Logs or ELK stack.
- **Dashboard**: Grafana or simple React app.

## Implementation Steps
1. Define failure scenarios and corresponding remediation steps.
2. Write monitoring rules and alarm configurations.
5. Develop Cloud Functions/Ansible remediation handlers.
4. Build a simple dashboard displaying health and actions.
5. Conduct chaos experiments to validate healing.
6. Package everything with Terraform and document the process.

## Why This Project?
- Demonstrates mature operational thinking and reduces perceived risk for potential employers.
- Offers a compelling narrative: "the system heals itself while I sleep."