
BINARY_BASE_NAME=gobenkyoo-learnbox

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

IMPORT_NAME=gobenkyoo-import
TESTKANJIFILE=../$(IMPORT_NAME)/$(TESTDATA)/kanjis1.csv
TESTWORDFILE=../$(IMPORT_NAME)/$(TESTDATA)/words1.csv

BINARY_IMPORT_NAME=../$(IMPORT_NAME)/$(IMPORT_NAME)

prepare:
	./$(BINARY_IMPORT_NAME) \
		-book "kanjilexikon" \
		-lesson "all" \
		-type "kanji" \
		-file "$(TESTKANJIFILE)" \
		-header \
		-sep ";" \
		-fieldsep "/" \
		-fields "KANJI,,,READING,MEANINGS,HINT,EXPLANATION" \
		-configdir "$(TESTDATA)"
	./$(BINARY_IMPORT_NAME) \
		-book "japanisch" \
		-lesson "Lektion 1" \
		-type "word" \
		-file "$(TESTWORDFILE)" \
		-header \
		-sep "," \
		-fields "KANA,NIHONGO,ROMAJI,MEANING,HINT,EXPLANATION,DICTFORM,TEFORM,NAIFORM" \
		-configdir "$(TESTDATA)" \

run:
	./$(BINARY_BASE_NAME) \
		-configdir "$(TESTDATA)" \
		-book "japanisch" \
		-lesson "Lektion 1" \

runfrom:
	./$(BINARY_BASE_NAME) \
		-configdir "$(TESTDATA)" \
		-book "minna 1" \
		-lesson "Kanjis Lektion 1" \
		-frombook "kanjilexikon" \
		-fromfile "$(TESTDATA)/kanjis.txt" \


.PHONY: all build linux windows prepare run runfrom

