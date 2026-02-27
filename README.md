# Portfolio Website

This is a portfolio website for C. Michael Harrington, showcasing his skills and experience as a DevOps Engineer.

## Usage

Bring up both frontend and backend locally using Docker Compose:

```bash
cd /home/mike/junk/my_website
docker compose up --build
```

Frontend will be available at `http://localhost/` and the API at
`http://localhost:8080`.  The backend exposes `/projects` and
`/projects/{slug}` as a placeholder framework for the demo projects.


## Files

- `Resume_Harrington_2025.docx.txt`: Original resume file.

## Technologies Used

## Infrastructure & Deployment

The application is designed to run entirely in Google Cloud.  Two support
mechanisms are provided:

* **Terraform** – located under `architecture/terraform`; it provisions an
  Artifact Registry repository and a Cloud Run service for the backend.  See
  `architecture/terraform/README.md` for usage instructions.
* **CI/CD** – deployment to GCP happens **exclusively** via the
  GitHub Actions workflow (`.github/workflows/cloudbuild.yml`).  On each push
  to `main`, the action builds the backend Docker image, pushes it to Artifact
  Registry and deploys to Cloud Run.  No manual `gcloud`/`cloudbuild` invocations
  are permitted for production deployments (they are useful only for local
  experimentation).


## Troubleshooting

## Next Steps

1. **Install CLI tools**
   - `gcloud` (Google Cloud SDK) and `terraform` (see https://learn.hashicorp.com/tutorials/terraform/install-cli).  On Debian/Ubuntu you can use `sudo snap install terraform`.
1. **Create/configure GCP project**
   ```sh
   gcloud projects create my-portfolio-demo
   gcloud config set project my-portfolio-demo
   # enable billing
   ```
   The Terraform configuration will also activate required APIs automatically.
1. **Run Terraform**
   ```sh
   cd architecture/terraform
   terraform init
   terraform apply -var="project_id=$GOOGLE_CLOUD_PROJECT"
   ```
   After the apply completes note the `backend_url` output.
1. **Set up Git/GitHub integration**
   - Push this repo to GitHub.
   - Add repository secrets `GCP_CREDENTIALS` (service account JSON) and
     `GCP_PROJECT_ID`.
   - The workflow at `.github/workflows/cloudbuild.yml` will trigger on every
     push to `main` and is the *only* supported deployment channel.
1. **Trigger a build**
   - Push a commit to `main` and wait for the action to run.
   - Confirm the backend receives traffic at the URL from Terraform output.
   - Manual `gcloud`/`cloudbuild` commands are permitted only for local proof‑
     of‑concept work and are **not** used for production deployments.
1. **Iterate**
   - Expand Terraform with additional modules (databases, VPC, etc.).
   - Update `cloudbuild.yaml` to build frontend or multiple services.
   - Follow the microservices CI/CD project outline in `/projects/microservices_project.md`.

## Requirements from User

1. Hosted on Google Cloud
2. Prioritize free tier useage at all times
3. Must be containerized for all development and testing
4. Prefer Go over Python
