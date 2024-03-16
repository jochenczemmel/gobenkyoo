	
BINARY_BASE_NAME=gobenkyoo

# build linux binary
linux:
	GOOS=linux GOARCH=amd64 go build
	strip $(BINARY_BASE_NAME)

# build binaries
build: windows linux darwin darwinnew

# build windows binary
windows:
	GOOS=windows GOARCH=amd64 go build

darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_BASE_NAME).amd64

darwinnew:
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_BASE_NAME).arm64


.PHONY: build linux windows darwin darwinnew

