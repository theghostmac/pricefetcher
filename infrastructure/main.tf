provider "kubernetes" {
  config_path = "./kubeconfig.yaml"
}

resource "kubernetes_deployment" "pricefetcher_deployment" {
  metadata {
    name = "pricefetcher_deployment"
    labels = {
      app = "pricefetcher"
    }
  }

  spec {
    replicas = "var.replicas"

    selector {
      match_labels = {
        app = "pricefetcher"
      }
    }

    template {
      metadata {
        labels = {
          app = "pricefetcher"
        }
      }

      spec {
        container {
          name = "var.docker_image_url"
          name = "pricefetcher"

          port {
            container_port = 50051
          }
        }
      }
    }
  }
}