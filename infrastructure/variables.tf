variable "docker_image_url" {
  description = "URL of the Docker image to deploy"
}

variable "replicas" {
  description = "Number of replicas for the deployment"
  default = 3
}