output "backend_url" {
  description = "URL of the deployed backend Cloud Run service."
  value       = google_cloud_run_service.backend.status[0].url
}
