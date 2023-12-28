
TMP_FILE=testcoverage.txt
COVERAGE_FILE=testcoverage.html


# run acceptance test
# this should not fail any more during development!
test:
	go test -tags acctest

# run acceptance test verbose
testv:
	go test -v -tags acctest

# run current acceptance test verbose
# this fails during development!
# when it passes, change build tag to acctest
current:
	go test -v -tags current

# run acceptance test test coverage
cover:
	go test -tags acctest -cover -coverprofile $(TMP_FILE)
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)

.PHONY: test testv cover current

