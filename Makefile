
TMP_FILE=testcoverage.txt
COVERAGE_FILE=testcoverage.html

all: cover vulncheck lint

# vulnerability check
vulncheck:
	govulncheck ./...

# static code analysis
lint:
	golangci-lint run

# run unit test
test:
	go test ./...

# run verbose unit test 
testv:
	go test -v ./...

# run test coverage
cover:
	go test -cover -coverprofile $(TMP_FILE) ./...
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)
	echo "firefox $(COVERAGE_FILE) &"

.PHONY: all vulncheck lint test testv cover

