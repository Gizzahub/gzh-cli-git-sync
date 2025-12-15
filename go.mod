module github.com/Gizzahub/gzh-cli-git-sync

go 1.25.1

require (
	github.com/gizzahub/gzh-cli-git v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.10.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	golang.org/x/sync v0.18.0 // indirect
)

replace github.com/gizzahub/gzh-cli-git => ../gzh-cli-git
