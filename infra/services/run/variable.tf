variable "project_config" {
  type = object({
    project_id      = string
    region          = string
    docker_repo_url = string
  })
}

variable "environment_config" {
  type = object({
    name   = string
    domain = string
  })
}

variable "service_config" {
  type = object({
    name                        = string
    tag                         = string
    environment_variable_prefix = string
    mode                        = string
  })
}
