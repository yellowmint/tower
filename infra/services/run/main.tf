resource "google_cloud_run_service" "service" {
  name     = var.service_config.name
  location = var.project_config.region

  template {
    spec {
      containers {
        image = "${var.project_config.docker_repo_url}/${var.service_config.name}:${var.service_config.tag}"

        env {
          name  = "${var.service_config.environment_variable_prefix}MODE"
          value = var.service_config.mode
        }
        env {
          name  = "${var.service_config.environment_variable_prefix}DOMAIN"
          value = "${var.service_config.name}.${var.environment_config.name}.${var.environment_config.domain}"
        }
      }
    }
  }

  autogenerate_revision_name = true

  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_domain_mapping" "mapping" {
  location = var.project_config.region
  name     = "${var.service_config.name}.${var.environment_config.name}.${var.environment_config.domain}"

  metadata {
    namespace = var.project_config.project_id
  }

  spec {
    route_name = google_cloud_run_service.service.name
  }
}

resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.service.name
  location = google_cloud_run_service.service.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
