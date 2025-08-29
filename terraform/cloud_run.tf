resource "google_cloud_run_service" "naqa_api" {
  name     = var.service_name
  location = var.region
  project  = var.project_id

  depends_on = [google_project_service.cloud_run_api]

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/${var.service_name}:latest"

        ports {
          container_port = 8080
        }

        env {
          name  = "ENVIRONMENT"
          value = "production"
        }

        env {
          name  = "APP_NAME"
          value = "Naqa API"
        }

        env {
          name  = "API_VERSION"
          value = "1.0.0"
        }

        env {
          name  = "MONGO_URI"
          value = var.mongo_uri
        }

        env {
          name  = "MONGO_DATABASE"
          value = var.mongo_database
        }

        resources {
          limits = {
            memory = "1Gi"
            cpu    = "1000m"
          }
        }
      }

      container_concurrency = 1000
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "100"
        "autoscaling.knative.dev/minScale" = "0"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_binding" "public_access" {
  location = google_cloud_run_service.naqa_api.location
  project  = google_cloud_run_service.naqa_api.project
  service  = google_cloud_run_service.naqa_api.name
  role     = "roles/run.invoker"
  members = [
    "allUsers",
  ]
}