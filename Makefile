####
# 'go build' is annoying. just for type 'make' for build
####

COMSTOCK_DIR="github.com/shinpei/comstock"
.PHONY: all
all:
	python model/error-gen.py
	go build -tags debug

.PHONY: release
release:
	go test ${COMSTOCK_DIR}/engine
	go test ${COMSTOCK_DIR}/parser
	python model/error-gen.py
	go build

clean:
	go clean
