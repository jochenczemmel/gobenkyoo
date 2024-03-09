
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
LIBDIR=$(TESTDATA)/jsondb/library/default/
KANJIBOXDIR=$(TESTDATA)/jsondb/classroom/kanjis/default
WORDBOXDIR=$(TESTDATA)/jsondb/classroom/words/default

prepare:
	mkdir -p $(LIBDIR)
	cp $(TESTDATA)/kanji_and_word.json $(LIBDIR)/japanisch\%0A\%0A0.json
	# cp $(TESTDATA)/word.json $(LIBDIR)/japanisch\%0A\%0A0.json
	# cp $(TESTDATA)/kanji.json $(LIBDIR)/kanjilexikon\%0A\%0A0.json

test: clean prepare run
	cmp $(TESTDATA)/want_kanji_box_1.json \
		$(KANJIBOXDIR)/Lektion\%201\%0ALektion\%201\%0Ajapanisch\%0A\%0A0.json
	cmp $(TESTDATA)/want_word_box_1.json \
		$(WORDBOXDIR)/Lektion\%201\%0ALektion\%201\%0Ajapanisch\%0A\%0A0.json


clean:
	rm -rf $(TESTDATA)/jsondb

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
		-fromfile "$(TESTDATA)/kanjilist.txt" \


.PHONY: all build linux windows prepare clean test run runfrom

