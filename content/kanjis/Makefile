
TMP_FILE=coverage.txt
COVERAGE_FILE=coverage.html


# run unit test
test:
	go test

# run verbose unit test 
testv:
	go test -v

# run test coverage
cover:
	go test -cover -coverprofile $(TMP_FILE)
	go tool cover -func=$(TMP_FILE)
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)
	echo "firefox $(COVERAGE_FILE) &"
	echo "chromium-browser $(COVERAGE_FILE) &"

# package list for test coverage
PACKAGE_LIST:=github.com/jochenczemmel/gobenkyoo/content/kanjis
PACKAGE_LIST:=$(PACKAGE_LIST),github.com/jochenczemmel/gobenkyoo/content/kanjis/radicals

# run test coverage
coverall:
	go test -cover -coverprofile $(TMP_FILE) -coverpkg $(PACKAGE_LIST) ./...
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)
	echo "firefox $(COVERAGE_FILE) &"
	echo "chromium-browser $(COVERAGE_FILE) &"

.PHONY: test testv cover coverall

