	
BINARY_BASE_NAME=gobenkyoo

# build linux binary
build:
	GOOS=linux GOARCH=amd64 go build
	strip $(BINARY_BASE_NAME)

# build binaries
build_all: build_windows build

# build windows binary
build_windows:
	GOOS=windows GOARCH=amd64 go build

# call executable like in the system tests
systemtest: build
	./gobenkyoo -dbtype json -ui learn


.PHONY: build_all build build_windows systemtest

