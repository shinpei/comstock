####
# 'go build' is annoying. just for type 'make' for build
####

release:
	python model/error-gen.py
	go build
all:
	python model/error-gen.py
	go build -tags=debug

clean:
	go clean
