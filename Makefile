
TMP_FILE=testcoverage.txt
COVERAGE_FILE=testcoverage.html

# the first target is executed when make is called without target
develop: systemtest

# run all unit tests
testall: cover vulncheck lint

# vulnerability check
vulncheck:
	govulncheck ./...

# static code analysis
lint:
	golangci-lint run

# package list for test coverage
PACKAGE_LIST:=github.com/jochenczemmel/gobenkyoo/content
PACKAGE_LIST:=$(PACKAGE_LIST),github.com/jochenczemmel/gobenkyoo/content/words
PACKAGE_LIST:=$(PACKAGE_LIST),github.com/jochenczemmel/gobenkyoo/content/books
PACKAGE_LIST:=$(PACKAGE_LIST),github.com/jochenczemmel/gobenkyoo/content/kanjis
PACKAGE_LIST:=$(PACKAGE_LIST),github.com/jochenczemmel/gobenkyoo/content/kanjis/radicals

# run test coverage
cover:
	go test -cover -coverprofile $(TMP_FILE) -coverpkg $(PACKAGE_LIST) ./...
	# go test -cover -coverprofile $(TMP_FILE) ./...
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)
	echo "firefox $(COVERAGE_FILE) &"

# run unit test
test:
	go test ./...

# run verbose unit test 
testv:
	go test -v ./...

# run system test:
systemtest: build
	./systemtest/simple_correct_save.exp
	./systemtest/simple_correct_nosave.exp

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

