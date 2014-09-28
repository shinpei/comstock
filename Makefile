####
# 'go build' is annoying. just for type 'make' for build
####

all:
	python model/error-gen.py
	go build

clean:
	go clean
