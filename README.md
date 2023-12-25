# gobenkyoo
Application to learn japanese

# code directory structure
The directories contain the following parts of the application:

## content
Contains the 'business objects', the content of the learning process.

### words
Contains the code for the vocabulary content.

### kanjis
Contains the code for the kanji content.

#### radicals
Contains the code for handling of kanji radicals.

### books
Contains the code to bundle content in books and lessons.

## app
Contains the application logic.

### learn
Contains the code for the learning process.

### search
Contains the code for searching in the content.

## storage
Contains the code for importing and storing content and learning state.

## ui
Contains the code for the user interfaces.

## cmd
Contains directories for executables.

## systemtest
Contains some expect scripts for complete system test.

wikipedia: https://en.wikipedia.org/wiki/Expect

maybe use go-libraries?

* https://pkg.go.dev/github.com/google/goexpect (https://github.com/google/goexpect)
* https://pkg.go.dev/github.com/Netflix/go-expect (https://github.com/Netflix/go-expect)
* https://pkg.go.dev/github.com/ryandgoulding/goexpect


