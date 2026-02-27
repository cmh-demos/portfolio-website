# Edge Compute IoT Platform

## Project Overview
Build a distributed IoT platform that processes sensor data at the edge on devices such as Raspberry Pi or other SBCs. Rather than shipping every reading to the cloud, edge nodes perform initial filtering, rule evaluation, or even simple ML inference before forwarding summaries to a central server. The system should demonstrate offline capability, over‑the‑air updates, and secure synchronization with cloud services.

## Why it fits your skills
- Leverages your hardware and scripting experience along with cloud knowledge.
- Shows your ability to design low‑latency, fault‑tolerant architectures familiar from aerospace telemetry systems.
- Combines Python/Go development with containerization, networking, and Google Cloud IoT Core or MQTT expertise.

## Key Features
1. **Edge Agent**: A lightweight containerized service that runs on an edge device, collects sensor data (e.g., temperature, GPS), and applies configurable filters/alerts.
2. **Central Dashboard**: Aggregates edge summaries in a web UI, enabling map views, historical graphs, and device status.
3. **Offline First**: Nodes continue capturing and queuing data when disconnected; sync automatically upon reconnection.
4. **OTA Updates**: Secure mechanism to push new edge software/images via the central server.
5. **Security**: Mutual TLS or IAM‑based authentication between devices and cloud; encrypted local storage.
6. **Extensibility**: Plugin system for new sensor types and processing rules.

## Tech Stack
- **Edge**: Python/Go, Docker, MQTT (Eclipse Mosquitto or Cloud IoT), SQLite for local buffering.
- **Central**: Node.js or Python backend, React frontend, PostgreSQL or Firestore for metadata.
- **Infrastructure**: Cloud IoT Core / IoT Edge, Cloud Storage for firmware blobs, Terraform for IaC.
- **Tools**: Git, CI/CD pipeline to build and publish edge images.

## Implementation Steps
1. Prototype edge agent on a Pi with a sample sensor (e.g., temperature).
2. Build central server API and dashboard; connect via MQTT broker.
3. Add offline buffering and sync logic.
4. Implement OTA update mechanism using signed images in Cloud Storage.
5. Harden security and add extensibility points.
6. Test end‑to‑end with multiple devices; deploy a demo environment.

## Why This Project?
- Bridges physical computing and cloud, broadening your portfolio beyond pure software.
- Demonstrates real‑world problem solving in constrained environments, similar to airborne systems.
- Offers opportunities to showcase telemetry, networking, and edge‑computing know‑how.