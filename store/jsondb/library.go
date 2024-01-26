package jsondb

type Library struct {
	Name  string `json:"Name,omitempty"`
	Books []Book `json:"Books,omitempty"`
}

type Book struct {
	Title         string            `json:"Title,omitempty"`
	SeriesTitle   string            `json:"SeriesTitle,omitempty"`
	Volume        int               `json:"Volume,omitempty"`
	LessonNames   []string          `json:"LessonNames,omitempty"`
	LessonsByName map[string]Lesson `json:"LessonsByName,omitempty"`
}

type Lesson struct {
	Name       string      `json:"Name"`
	WordCards  []WordCard  `json:"WordCards,omitempty"`
	KanjiCards []KanjiCard `json:"KanjiCards,omitempty"`
}

type WordCard struct {
	ID          int    `json:"ID"`
	Nihongo     string `json:"Nihongo,omitempty"`
	Kana        string `json:"Kana,omitempty"`
	Romaji      string `json:"Romaji,omitempty"`
	Meaning     string `json:"Meaning"`
	Hint        string `json:"Hint,omitempty"`
	Explanation string `json:"Explanation,omitempty"`
	DictForm    string `json:"DictForm,omitempty"`
	TeForm      string `json:"TeForm,omitempty"`
	NaiForm     string `json:"NaiForm,omitempty"`
}

type KanjiCard struct {
	ID           int           `json:"ID"`
	Kanji        string        `json:"Kanji"`
	KanjiDetails []KanjiDetail `json:"KanjiDetails"`
}

type KanjiDetail struct {
	Reading     string   `json:"Reading"`
	ReadingKana string   `json:"ReadingKana,omitempty"`
	Meanings    []string `json:"Meanings"`
}
