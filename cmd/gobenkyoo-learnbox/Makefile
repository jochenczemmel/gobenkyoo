
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

# execute integration tests
inttest: test1 test2

# prepare test 1
prepare1:
	mkdir -p $(LIBDIR)
	cp $(TESTDATA)/kanji_and_word.json $(LIBDIR)/japanisch\%0A\%0A0.json

# prepare test 2
prepare2:
	mkdir -p $(LIBDIR)
	cp $(TESTDATA)/kanji.json $(LIBDIR)/kanjilexikon\%0A\%0A0.json
	# cp $(TESTDATA)/word.json $(LIBDIR)/japanisch\%0A\%0A0.json


# test1: create box from lesson
test1: clean prepare1 run
	cmp $(TESTDATA)/want_kanji_box_1.json \
		$(KANJIBOXDIR)/Lektion\%201\%0ALektion\%201\%0Ajapanisch\%0A\%0A0.json
	cmp $(TESTDATA)/want_word_box_1.json \
		$(WORDBOXDIR)/Lektion\%201\%0ALektion\%201\%0Ajapanisch\%0A\%0A0.json

# test2: create box from kanji list and kanji book
test2: clean prepare2 runfrom
	cmp $(TESTDATA)/want_kanji_box_2.json \
		$(KANJIBOXDIR)/lesson\%201\%0Alesson\%201\%0Aminna\%20sho\%201\%0Aminna\%20no\%20nihongo\%0A1.json

# remove temporary files
clean:
	rm -rf $(TESTDATA)/jsondb

# execute box creation
run:
	./$(BINARY_BASE_NAME) \
		-configdir "$(TESTDATA)" \
		-book "japanisch" \
		-lesson "Lektion 1" \

# execute box creation
runfrom:
	./$(BINARY_BASE_NAME) \
		-configdir "$(TESTDATA)" \
		-book "minna sho 1" \
		-series "minna no nihongo" \
		-volume 1 \
		-lesson "lesson 1" \
		-frombook "kanjilexikon" \
		-fromfile "$(TESTDATA)/kanjilist.txt" \


.PHONY: all build linux windows inttest test1 test2 \
	prepare1 prepare2 clean run runfrom

