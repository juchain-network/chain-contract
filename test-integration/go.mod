module juchain.org/chain/tools/ci

go 1.24.11

replace juchain.org/chain/tools => ../tools

require (
	github.com/ethereum/go-ethereum v1.16.8
	gopkg.in/yaml.v3 v3.0.1
	juchain.org/chain/tools v0.0.0-00010101000000-000000000000
)

replace juchain.org/chain => ../
