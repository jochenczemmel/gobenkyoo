
TMP_FILE=testcoverage.txt
COVERAGE_FILE=testcoverage.html


# run acceptance test
acctest:
	go test -tags acctest

# run acceptance test test coverage
acccover:
	go test -tags acctest -cover -coverprofile $(TMP_FILE)
	go tool cover -html=$(TMP_FILE) -o $(COVERAGE_FILE)

.PHONY: acctest acccover

