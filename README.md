# terraform-provider-kubectl

A terraform provider to run various `kubectl` operations across multiple clusters in a single `terraform apply`.

Use-cases:

- Integrate `kubectl` run(s) into Terraform-managed DAG of resources
  - Annotate or label K8s resources created by other Terraform resource like `eksctl_cluster` and `eksctl`

## Prerequisites

- Kubectl (If you prefer not letting the provider to install it)

## Installation

**For Terraform 0.12:**

Install the `terraform-provider-kubectl` binary under `.terraform/plugins/${OS}_${ARCH}`, so that the binary is at e.g. `${WORKSPACE}/.terraform/plugins/darwin_amd64/terraform-provider-kubectl`.

**For Terraform 0.13 and later:**

The provider is [available at Terraform Registry](https://registry.terraform.io/providers/mumoshu/kubectl/latest?pollNotifications=true) so you can just add the following to your tf file for installation:

```
terraform {
  required_providers {
    kubectl = {
      source = "mumoshu/kubectl"
      version = "VERSION"
    }
  }
}
```

Please replace `VERSION` with the version number of the provider without the `v` prefix, like `0.1.0`. 

## Examples

There is nothing to configure for the provider, so you firstly declare the provider like:

```
provider "kubectl" {}
```

The only supported resource is `kubectl_ensure`.

```hcl
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
```

See [`the labels and annotations example`](./example/testdata/01-bootstrap) for more details.

## Advanced Features

- [Declarative binary version management](#declarative-binary-version-management)

## Declarative binary version management

`terraform-provider-kubectl` has a built-in package manager called [shoal](https://github.com/mumoshu/shoal).
With that, you can specify the following `kubectl_ensure` attributes to let the provider install the executable binaries on demand:

- `version` for installing `kubectl`

`version` uses the Go runtime and [go-git](https://github.com/go-git/go-git) so it should work without any dependency.

With the below example, the provider installs the latest version of `kubectl` v1.18.x so that you don't need to install them beforehand.
This should be handy when you're trying to use this provider on Terraform Cloud, whose runtime environment is [not available for customization by the user](https://www.terraform.io/docs/cloud/run/run-environment.html).    

```hcl-terraform
resource "kubectl_ensure" "meta" {
  version = ">= 1.18.0, < 1.19.0"

  // snip
```

Please see [this example](./example/testdata/02-shoal) for more details.

## Develop
If you wish to build this yourself, follow the instructions:

	cd terraform-provider-kubectl
	go build

## Alternatives

This project has been published in October 2020. At that time, there was already an existing, well-known, and awesome project called [gavinbunney/terraform-provider-kubectl](https://github.com/gavinbunney/terraform-provider-kubectl).

The reason I created this was I wanted to patch K8s resources with labels and annotations, so that I can adopt `aws-auth` configmap created by `eksctl` to be managed by `helm`.

I may gradually add other features like gabinbunney's kubectl provider cover, like applying K8s manifests. But if you think this provider isn't what you want,
I definitely recommend checking out [gabinbuney's](https://github.com/gavinbunney/terraform-provider-kubectl) as an alternative to this provider!
