# Terraform Infrastructure

This folder contains Terraform configuration for provisioning the GCP resources
used by the portfolio application.  The first iteration provisions a single
Cloud Run service and an Artifact Registry repository for container images.

## Usage

1. Install and authenticate the Google Cloud SDK (`gcloud auth login`).
2. Create or select a GCP project and note its ID:
   ```sh
   gcloud projects create my-portfolio-demo
   gcloud config set project my-portfolio-demo
   ```
3. Enable billing for the project and make sure the APIs are available (the
   `google_project_service` resources will request them automatically).
4. Initialize Terraform.  If you don't have Terraform installed locally you can
   use the provided wrapper script which runs it in a Docker container:
   ```sh
   cd architecture/terraform
   # either
   terraform init
   terraform apply -var="project_id=$GOOGLE_CLOUD_PROJECT"
   # or via docker wrapper
   ./run.sh init
   ./run.sh apply -var="project_id=$GOOGLE_CLOUD_PROJECT"
   ```
5. The backend URL will be printed as an output; use it to configure the
   frontend or hit the service directly.

You can customize `region` and `repository_id` by passing additional variables.

### Extending the configuration

Additional modules (see `modules/`) can be added for storage buckets, IAM
policies, Cloud SQL instances, etc.  This layout encourages reuse across
different demos (microservices, self-healing, etc.).
