
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

.PHONY: test testv cover maketestdata


TESTDATA=testdata

maketestdata:
	../cmd/gobenkyoo-import/gobenkyoo-import \
		-book "minna kyu 1" \
		-series "minna" \
		-volume 1 \
		-lesson "lesson 1" \
		-type "kanji" \
		-file "$(TESTDATA)/kanjis1.csv" \
		-header \
		-sep ";" \
		-fieldsep "/" \
		-fields "KANJI,,,READING,MEANINGS,HINT,EXPLANATION" \
		-configdir "$(TESTDATA)"
	../cmd/gobenkyoo-import/gobenkyoo-import \
		-book "minna kyu 1" \
		-series "minna" \
		-volume 1 \
		-lesson "lesson 2" \
		-type "kanji" \
		-file "$(TESTDATA)/kanjis2.csv" \
		-header \
		-sep ";" \
		-fieldsep "/" \
		-fields "KANJI,,,READING,MEANINGS,HINT,EXPLANATION" \
		-configdir "$(TESTDATA)"
	../cmd/gobenkyoo-import/gobenkyoo-import \
		-book "minna kyu 1" \
		-series "minna" \
		-volume 1 \
		-lesson "lesson 1" \
		-type "word" \
		-file "$(TESTDATA)/words1.csv" \
		-header \
		-sep "," \
		-fields "KANA,NIHONGO,ROMAJI,MEANING,HINT,EXPLANATION,DICTFORM,TEFORM,NAIFORM" \
		-configdir "$(TESTDATA)" \

