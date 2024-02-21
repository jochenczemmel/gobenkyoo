
TMP_FILE=coverage.txt
COVERAGE_FILE=coverage.html


# run unit test
test:
	go test

# run verbose unit test 
testv:
	go test -v

# run unit test
inttest:
	go test -tags inttest

# run test coverage
cover:
	go test -tags inttest -cover -coverprofile $(TMP_FILE)
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)
	echo "firefox $(COVERAGE_FILE) &"
	echo "chromium-browser $(COVERAGE_FILE) &"


.PHONY: test testv cover

