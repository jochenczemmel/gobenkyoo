package ui

type ExamController interface {
	NextCard() bool
	Question() ([]string, []string)
	JudgeCard(correct bool)
	Answer() ([]string, []string)

	ExamIndex() (int, int)
	CurrentExamInfos() (string, string, int)
	CurrentCardID() string

	Save() error // SaveCurrentBox() error

	// ExamMode() string
}

type LearnController interface {
	Levels() []string
	Modes() []string
	BookTitles() []string
	BoxLevels(booktitle, boxtitle, mode string, cumulative bool) []int
	BoxTitlesForBook(booktitle string) []string

	// TODO: add cumulative?
	StartExam(mode string, level int) ExamController
	// TODO: remove cumulative?
	SetCurrentBox(booktitle, boxtitle string, cumulative bool) error
	ResetCurrentBox()

	// IsKanji2NativeMode(mode string) bool
	// SaveCurrentBox() error
}
