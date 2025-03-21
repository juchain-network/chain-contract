##
##编译solidity，并产生bin文件，abi文件，和相应的go文件 solc版本0.7.6

currentDir := `pwd`


SRC_CONTRACTJu := contracts
GO_OUT4Ju := ${SRC_CONTRACTJu}/generated

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
	@solc --allow-paths $(ALLOW_PATH) --optimize --combined-json abi,bin,userdoc,devdoc  $(SRC_CONTRACTJu)/Validators.sol -o $(SRC_CONTRACTJu)/ --overwrite
	@abigen --combined-json $(SRC_CONTRACTJu)/combined.json --pkg $(PACKAGE) --out $(GO_OUT4Ju)/validators.go
	@rm $(SRC_CONTRACTJu)/combined.json




clean:
	@rm -fr $(GO_OUT)/*


asmExample:
	@solc --asm --bin -o  tmp/ valset-bin=./tmp/valset-bin Valset.sol

