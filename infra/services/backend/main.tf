module "run_backend_service_rpcpublic" {
  source             = "../run"
  project_config     = var.project_config
  environment_config = var.environment_config
  service_config     = {
    name                        = "be-rpcpublic"
    tag                         = var.service_config.tag
    environment_variable_prefix = "TOWER_"
    mode                        = var.service_config.mode
  }
}
