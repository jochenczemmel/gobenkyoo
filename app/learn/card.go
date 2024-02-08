package learn

// Card provides the information that is needed for learning drills.
// Cards can be created from kanji or word content.
// The learning mode determines which information is put in the
// Question, Answer, and MoreAnswers.
// The Hint and Explanation are not depending on the learning mode.
type Card struct {
	ID          CardID   // unique identifier
	Question    string   // what is presented first
	Hint        string   // additional information for the question
	Answer      string   // the correct answer
	MoreAnswers []string // additional information (verb forms, kanji readings, ...)
	Explanation string   // further information
}

// emptyCard is returned from several functions and methods.
// It ensures that MoreAnswers is not nil.
var emptyCard = Card{
	MoreAnswers: []string{},
}
