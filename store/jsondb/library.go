package jsondb

type Library struct {
	Name  string `json:"name,omitempty"`
	Books []Book `json:"books,omitempty"`
}

type Book struct {
	Title         string            `json:"title,omitempty"`
	SeriesTitle   string            `json:"seriesTitle,omitempty"`
	Volume        int               `json:"volume,omitempty"`
	LessonNames   []string          `json:"lessonNames,omitempty"`
	LessonsByName map[string]Lesson `json:"lessonsByName,omitempty"`
}

type Lesson struct {
	Name       string      `json:"name"`
	WordCards  []WordCard  `json:"wordCards,omitempty"`
	KanjiCards []KanjiCard `json:"kanjiCards,omitempty"`
}

type WordCard struct {
	ID          int    `json:"id"`
	Nihongo     string `json:"nihongo,omitempty"`
	Kana        string `json:"kana,omitempty"`
	Romaji      string `json:"romaji,omitempty"`
	Meaning     string `json:"meaning"`
	Hint        string `json:"hint,omitempty"`
	Explanation string `json:"explanation,omitempty"`
	DictForm    string `json:"dictForm,omitempty"`
	TeForm      string `json:"teForm,omitempty"`
	NaiForm     string `json:"naiForm,omitempty"`
}

type KanjiCard struct {
	ID           int           `json:"id"`
	Kanji        string        `json:"kanji"`
	KanjiDetails []KanjiDetail `json:"kanjiDetails"`
}

type KanjiDetail struct {
	Reading     string   `json:"reading"`
	ReadingKana string   `json:"readingKana,omitempty"`
	Meanings    []string `json:"meanings"`
}
