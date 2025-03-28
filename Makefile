##
##编译solidity，并产生bin文件，abi文件，和相应的go文件 solc版本0.5.16

currentDir := `pwd`


SRC_CONTRACTJu := contracts
GO_OUT4Ju := ${SRC_CONTRACTJu}/generated
OUT := build
ALLOW_PATH := $(currentDir)

PACKAGE := generated

proj := "build"
.PHONY: default build clean registry bridgeBank setup

default: proposal punish validators

proposal:
	@solc --allow-paths $(ALLOW_PATH) --optimize --combined-json abi,bin,userdoc,devdoc  $(SRC_CONTRACTJu)/Proposal.sol -o $(SRC_CONTRACTJu)/ --overwrite
	@abigen --combined-json $(SRC_CONTRACTJu)/combined.json --pkg $(PACKAGE) --out $(GO_OUT4Ju)/proposal.go
	@rm $(SRC_CONTRACTJu)/combined.json

punish:
	@solc --allow-paths $(ALLOW_PATH) --optimize --combined-json abi,bin,userdoc,devdoc  $(SRC_CONTRACTJu)/Punish.sol -o $(SRC_CONTRACTJu)/ --overwrite
	@abigen --combined-json $(SRC_CONTRACTJu)/combined.json --pkg $(PACKAGE) --out $(GO_OUT4Ju)/punish.go
	@rm $(SRC_CONTRACTJu)/combined.json

validators:
	@#solc --allow-paths $(ALLOW_PATH) --optimize --combined-json abi,bin,userdoc,devdoc  $(SRC_CONTRACTJu)/Validators.sol -o $(SRC_CONTRACTJu)/ --overwrite
	@npx solcjs --optimize --combined-json abi --base-path $(ALLOW_PATH) $(SRC_CONTRACTJu)/Validators.sol > $(SRC_CONTRACTJu)/output.json
	@abigen --combined-json $(SRC_CONTRACTJu)/combined.json --pkg $(PACKAGE) --out $(GO_OUT4Ju)/validators.go
	@rm $(SRC_CONTRACTJu)/combined.json

build:
	@go build -o ${OUT}/congress

build_linux:
	@CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o ${OUT}/congress

# mac遇到依赖包用cgo的，需要开启cgo，先安装c交叉编译器 brew install FiloSottile/musl-cross/musl-cross
build_linux_cgo:
	@CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ CGO_ENABLED=1  GOOS=linux GOARCH=amd64 go build  -ldflags "-linkmode external -extldflags -static" -o ${OUT}/congress

clean:
	@rm -fr $(GO_OUT4Ju)/*


asmExample:
	@solc --asm --bin -o  tmp/ valset-bin=./tmp/valset-bin Valset.sol

