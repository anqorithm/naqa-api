resource "google_project_service" "cloud_run_api" {
  project = var.project_id
  service = "run.googleapis.com"
}

resource "google_project_service" "container_registry_api" {
  project = var.project_id
  service = "containerregistry.googleapis.com"
}

resource "google_project_service" "cloud_build_api" {
  project = var.project_id
  service = "cloudbuild.googleapis.com"
}