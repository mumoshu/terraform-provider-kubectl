module github.com/mumoshu/terraform-provider-kubectl

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/mumoshu/shoal v0.2.18
	github.com/mumoshu/terraform-provider-eksctl v0.14.6
	github.com/pkg/profile v1.5.0
	github.com/rs/xid v1.2.1
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)

replace github.com/fishworks/gofish => github.com/mumoshu/gofish v0.13.1-0.20200908033248-ab2d494fb15c

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
