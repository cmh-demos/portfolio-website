variable "project_id" {
  description = "The GCP project ID where resources will be created."
  type        = string
}

variable "region" {
  description = "GCP region for resources (Cloud Run, Artifact Registry)."
  type        = string
  default     = "us-central1"
}

variable "repository_id" {
  description = "Artifact Registry repository name used for Docker images."
  type        = string
  default     = "my-repo"
}
