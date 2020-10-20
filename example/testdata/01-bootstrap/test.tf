provider "kubectl" {}

variable "kubeconfig" {
  type = string
}

resource "kubectl_ensure" "meta" {
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
