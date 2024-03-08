
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



PACKAGE_LIST=

# PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/content
PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/content/words
PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/content/kanjis
PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/content/kanjis/radicals
PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/content/books

# PACKAGE_LIST:=github.com/jochenczemmel/gobenkyoo/cfg

# PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/app
PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/app/learn
# PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/app/search


# PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/store
PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/store/jsondb
PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/store/csvimport

# PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/ui
# PACKAGE_LIST+=github.com/jochenczemmel/gobenkyoo/ui/cli

# run test coverage
cover:
	go test -cover -coverprofile $(TMP_FILE) $(PACKAGE_LIST) 
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

