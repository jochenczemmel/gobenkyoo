package learn

type Card struct {
	Identity    string
	Question    string
	Hint        string
	Answer      string
	MoreAnswers []string
	Explanation string
}

// emptyCard is returned from several functions and methods,
// it avoids a nil value für MoreAnswers.
var emptyCard = Card{
	MoreAnswers: []string{},
}
