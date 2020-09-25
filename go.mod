module github.com/elsaland/elsa

go 1.14

require (
	github.com/elsaland/quickjs v0.0.0-20200925155809-0246a93f1a32
	github.com/evanw/esbuild v0.7.3-0.20200919185132-ef34da4ee06e
	github.com/fatih/color v1.9.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/robertkrimen/otto v0.0.0-20191219234010-c382bd3c16ff
	github.com/spf13/afero v1.4.0
	github.com/spf13/cobra v1.0.0
	github.com/tdewolff/minify/v2 v2.9.5
)

// TODO(@qu4k): remove when qu4k/quickjs is merged into elsaland/quickjs
replace (
	github.com/elsaland/quickjs => ../quickjs
)
