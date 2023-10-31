
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

PACKAGE_LIST=github.com/jochenczemmel/gobenkyoo/content,github.com/jochenczemmel/gobenkyoo/content/words,github.com/jochenczemmel/gobenkyoo/content/books,github.com/jochenczemmel/gobenkyoo/content/kanjis,github.com/jochenczemmel/gobenkyoo/content/kanjis/radicals

# run test coverage
cover:
	go test -cover -coverprofile $(TMP_FILE) -coverpkg $(PACKAGE_LIST) ./...
	# go test -cover -coverprofile $(TMP_FILE) ./...
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)
	echo "firefox $(COVERAGE_FILE) &"


.PHONY: all vulncheck lint test testv cover

