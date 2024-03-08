
TMP_FILE=coverage.txt
COVERAGE_FILE=coverage.html

# the first target is executed when make is called without target
# develop: systemtest

# run all unit tests
testall: vulncheck vet lint test

# vulnerability check
vulncheck:
	govulncheck ./...

# static code analysis
# https://pkg.go.dev/cmd/vet
vet:
	go vet ./...

# static code analysis
# https://golangci-lint.run/
lint:
	golangci-lint run


# run test coverage
cover:
	go test -cover -coverprofile $(TMP_FILE) ./...
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)
	echo "firefox $(COVERAGE_FILE) &"
	echo "chromium-browser $(COVERAGE_FILE) &"

# run unit test
test:
	go test ./...

# run verbose unit test 
testv:
	go test -v ./...

# run system test:
# systemtest: build
# 	./systemtest/simple_correct_save.exp
# 	./systemtest/simple_correct_nosave.exp

# run system test verbose:
systemtestv: build
	./systemtest/simple_correct_save.exp -d
	./systemtest/simple_correct_nosave.exp -d


# main executable
DIR_GOBENKYOO=cmd/gobenkyoo

# build executables
build:
	cd $(DIR_GOBENKYOO) && make build

.PHONY: testall vulncheck lint cover test testv systemtest

