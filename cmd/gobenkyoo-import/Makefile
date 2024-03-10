	
BINARY_BASE_NAME=gobenkyoo-import

# build all binaries
all: windows linux


# build linux binary
linux:
	GOOS=linux GOARCH=amd64 go build
	strip $(BINARY_BASE_NAME)

# alias on linux
build: linux

# build windows binary
windows:
	GOOS=windows GOARCH=amd64 go build


TESTDATA=testdata
TESTKANJIFILE=$(TESTDATA)/kanjis1.csv
TESTWORDFILE=$(TESTDATA)/words1.csv
TESTWANTFILE=$(TESTDATA)/want.json

# run a local test
inttest: clean runword runkanji
	cmp $(TESTWANTFILE) $(TESTDATA)/jsondb/library/default/japanisch\%0A\%0A0.json

clean:
	rm -rf $(TESTDATA)/jsondb

runkanji:
	./$(BINARY_BASE_NAME) \
		-book "japanisch" \
		-lesson "Lektion 1" \
		-type "kanji" \
		-file "$(TESTKANJIFILE)" \
		-header \
		-sep ";" \
		-fieldsep "/" \
		-fields "KANJI,,,READING,MEANINGS,HINT,EXPLANATION" \
		-configdir "$(TESTDATA)" \
		-nominify	

runword:
	./$(BINARY_BASE_NAME) \
		-book "japanisch" \
		-lesson "Lektion 1" \
		-type "word" \
		-file "$(TESTWORDFILE)" \
		-header \
		-sep "," \
		-fields "KANA,NIHONGO,ROMAJI,MEANING,HINT,EXPLANATION,DICTFORM,TEFORM,NAIFORM" \
		-configdir "$(TESTDATA)" \
		-nominify	


.PHONY: all build linux windows inttest runword runkanji

