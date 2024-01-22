package jsondb

type LearnCard struct {
	ID          int      `json:"ID"`
	LessonID    LessonID `json:"LessonID,omitempty"`
	Question    string   `json:"Question"`
	Hint        string   `json:"Hint,omitempty"`
	Answer      string   `json:"Answer"`
	MoreAnswers []string `json:"MoreAnswers,omitempty"`
	Explanation string   `json:"Explanation,omitempty"`
}

type LessonID struct {
	LessonTitle string `json:"Title,omitempty"`
	BookTitle   string `json:"BookTitle,omitempty"`
	SeriesTitle string `json:"SeriesTitle,omitempty"`
	Volume      int    `json:"Volume,omitempty"`
}
