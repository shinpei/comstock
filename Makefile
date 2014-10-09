####
# 'go build' is annoying. just for type 'make' for build
####

.PHONY: all
all:
	python model/error-gen.py
	go build -tags debug

.PHONY: release
release:
	python model/error-gen.py
	go build

clean:
	go clean
