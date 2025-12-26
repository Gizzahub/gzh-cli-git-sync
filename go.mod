module github.com/Gizzahub/gzh-cli-git-sync

go 1.25.1

require (
	github.com/gizzahub/gzh-cli-git v0.0.0-00010101000000-000000000000
	github.com/gizzahub/gzh-cli-gitforge v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.10.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/google/go-github/v66 v66.0.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/xanzy/go-gitlab v0.115.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	golang.org/x/time v0.3.0 // indirect
)

replace github.com/gizzahub/gzh-cli-git => ../gzh-cli-git

replace github.com/gizzahub/gzh-cli-gitforge => ../gzh-cli-gitforge
