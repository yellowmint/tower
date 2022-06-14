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

variable "application_config" {
  type = object({
    mode        = string
    backend_tag = string
  })
}
