module github.com/hashicorp/consul-terraform-sync

go 1.16

require (
	cloud.google.com/go v0.78.0 // indirect
	cloud.google.com/go/storage v1.13.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/PaloAltoNetworks/pango v0.5.1
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/aws/aws-sdk-go v1.37.19 // indirect
	github.com/deepmap/oapi-codegen v1.11.0
	github.com/fatih/color v1.10.0 // indirect
	github.com/getkin/kin-openapi v0.94.0
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-test/deep v1.0.7 // indirect
	github.com/google/uuid v1.3.0
	github.com/hashicorp/consul/api v1.13.0
	github.com/hashicorp/consul/sdk v0.9.0
	github.com/hashicorp/cronexpr v1.1.1
	github.com/hashicorp/go-bexpr v0.1.4
	github.com/hashicorp/go-checkpoint v0.5.0
	github.com/hashicorp/go-getter v1.6.1
	github.com/hashicorp/go-hclog v0.16.2
	github.com/hashicorp/go-retryablehttp v0.7.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2
	github.com/hashicorp/go-syslog v1.0.0
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.5.0
	github.com/hashicorp/hc-install v0.3.2
	github.com/hashicorp/hcat v0.2.1-0.20220519190242-5b1deea3fce6
	github.com/hashicorp/hcl v1.0.1-vault-2
	github.com/hashicorp/hcl/v2 v2.12.0
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform-exec v0.16.1
	github.com/hashicorp/terraform-json v0.14.0
	github.com/hashicorp/vault/api v1.7.2
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mitchellh/cli v1.1.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mitchellh/reflectwalk v1.0.2
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/posener/complete v1.2.3
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.1
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/zclconf/go-cty v1.10.0
	go.opencensus.io v0.22.6 // indirect
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93 // indirect
	google.golang.org/genproto v0.0.0-20210222212404-3e1e516060db // indirect

	// v3.0 has a CVE. Force dependencies github.com/deepmap/oapi-codegen to
	// use latest version of go-yaml
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
