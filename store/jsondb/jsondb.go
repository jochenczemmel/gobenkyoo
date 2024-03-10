// Package jsondb provides storing data as json files.
package jsondb

// DB provides loading and storing book libraries and
// learn classrooms in json format.
type DB struct {
	baseDir string
}

func New(dir string) DB {
	return DB{baseDir: dir}
}

var Minify bool

const BaseDir = "jsondb"

const (
	// directory name for library files.
	libraryPath = "library"
	// directory name for classroom files.
	classroomPath = "classroom"
	// directory name for kanji box files.
	kanjiPath = "kanjis"
	// directory name for word box files.
	wordPath = "words"
	// extension for json files.
	jsonExtension = ".json"
	// permissions for creating files.
	defaultFilePermissions = 0750
	// option for reading all files in a directory.
	readAllFiles = -1
)
