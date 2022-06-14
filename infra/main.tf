terraform {
  required_version = ">= 1.2.2"

  required_providers {
    google = ">= 4.24.0"
  }
}

variable "project_id" {
  default = "ad-tower"
}

variable "region" {
  default = "us-central1"
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_project_service" "api_run" {
  service            = "run.googleapis.com"
  disable_on_destroy = true
}

resource "google_project_service" "api_artifactregistry" {
  service            = "artifactregistry.googleapis.com"
  disable_on_destroy = true
}

resource "google_artifact_registry_repository" "docker_services_repo" {
  provider = google-beta
  project  = var.project_id
  location = var.region

  repository_id = "services"
  format        = "DOCKER"

  depends_on = [google_project_service.api_artifactregistry]
}

module "tower_alpha" {
  source         = "./tower"
  project_config = {
    project_id               = var.project_id
    region                   = var.region
    docker_repo_url = "${var.region}-docker.pkg.dev/${var.project_id}/${google_artifact_registry_repository.docker_services_repo.repository_id}"
  }
  environment_config = {
    name   = "alpha"
    domain = "tower.decoct.dev"
  }
  application_config = {
    mode        = "prod"
    backend_tag = "0.0.1-main-build-2"
  }

  depends_on = [google_project_service.api_run, google_artifact_registry_repository.docker_services_repo]
}
