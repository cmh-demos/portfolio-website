terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

# Enable the APIs we'll need for Cloud Run and Artifact Registry
resource "google_project_service" "run" {
  service = "run.googleapis.com"
}
resource "google_project_service" "artifact" {
  service = "artifactregistry.googleapis.com"
}
resource "google_project_service" "iam" {
  service = "iam.googleapis.com"
}

# Artifact Registry repository for our Docker images
resource "google_artifact_registry_repository" "repo" {
  provider      = google
  location      = var.region
  repository_id = var.repository_id
  format        = "DOCKER"
}

# Cloud Run service for the backend
resource "google_cloud_run_service" "backend" {
  name     = "backend"
  location = var.region

  template {
    spec {
      containers {
        image = "${google_artifact_registry_repository.repo.location}-docker.pkg.dev/${var.project_id}/${var.repository_id}/backend:latest"
      }
    }
  }

  autogenerate_revision_name = true
}

# Allow public, unauthenticated access (for demo purposes)
resource "google_cloud_run_service_iam_member" "public_invoker" {
  service = google_cloud_run_service.backend.name
  location = google_cloud_run_service.backend.location
  role    = "roles/run.invoker"
  member  = "allUsers"
}
