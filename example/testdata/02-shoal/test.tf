provider "kubectl" {}

variable "kubeconfig" {
  type = string
}

resource "kubectl_ensure" "meta" {
  version = ">= 1.18.0, < 1.19.0"

  kubeconfig = var.kubeconfig

  namespace = "kube-system"
  resource = "configmap"
  name = "aws-auth"

  labels = {
    "key1" = "one"
    "key2" = "two"
  }

  annotations = {
    "key3" = "three"
    "key4" = "four"
  }
}
