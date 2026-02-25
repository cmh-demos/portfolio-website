# Shared Reference Architecture

This folder contains the consolidated templates, modules, and documentation
that support all of the individual project exercises found in the root of the
workspace.  The goal is to provide a common foundation so new projects can be
spun up quickly without duplicating effort.

## Structure

- `backend/` – Python FastAPI service template with Dockerfile and utilities.
- `frontend/` – React component library and sample app.
- `terraform/` – Reusable Terraform modules (GCP storage, functions, IAM, etc.).
- `ci_cd/` – GitHub Actions workflows and scripts for linting, testing, and
deployment.
- `docs/` – Architecture diagrams, usage guides, onboarding notes.

Refer to the subdirectories for usage details and examples.
