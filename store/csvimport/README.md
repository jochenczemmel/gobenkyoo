# package csvimport: test data

The cvs import contains some test files and
corresponding card structs in the test code.

## files

The imported csv files are located in the directory "testdata".
The expected cards lists are defined in kanji\_test.go and
word\_test.go

### kanjis1.csv

This file contains kanji data with hint and explanation,
but without kana readings. It also contains some fields that
are not used in the import.

Fields are separated by a semicolon.

Each kanji is defined on one line.
Multiple readings or Meanings are separated by a slash.

* The data is imported using the format "format1".
* The expected card list is defined in the variable "kanji1Cards".

Another test imports only two fields from this file:

* Format: "minimalFormat1"
* Expected card list: "kanji1CardsMinimal"


### kanjis1noheader.csv

This file has the same content as kanjis1.csv, but without a
header line.

* Format: "format1".
* Expected card list: "kanji1CardsMinimal"


### kanjis1kana.csv

This file has the same content as kanjis1.csv with additional
readings in kana.

* Format: "format1kana".
* Expected card list: "kanji1CardsKana"

### kanjis2.csv

This file contains kanji data without hint and explanation,
but with kana readings. It also contains some fields that
are not used in the import.

Fields are separated by a semicolon.

A kanji can be defined on multiple lines.

Each line has a single reading.
Multiple meanings are separated by a semicolon.

* Format: "format2".
* Expected card list: "kanji2Cards"

### words1.csv

This file contains words data with all fields.
The fields are separated by a comma.

* Format: "format".
* Expected card list: "word1Cards"

Another test imports only two fields from this file:

* Format: "minimalFormat".
* Expected card list: "word1CardsMinimal"


### words1noheader.csv

This file has the same content as words1.csv, but without a
header line.

* Format: "format".
* Expected card list: "word1Cards"

