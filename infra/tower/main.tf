module "backend_service" {
  source             = "../services/backend"
  project_config     = var.project_config
  environment_config = var.environment_config
  service_config     = {
    tag  = var.application_config.backend_tag
    mode = var.application_config.mode
  }
}
