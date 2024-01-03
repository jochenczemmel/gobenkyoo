package books_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content"
	"github.com/jochenczemmel/gobenkyoo/content/books"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestLessonContains(t *testing.T) {

	wordList := []words.Card{
		{Meaning: "world", Nihongo: "世界"},
		{Meaning: "hello", Nihongo: "こんいてぃは"},
		{Meaning: "to see", Nihongo: "見る"},
	}
	kanjiList := []kanjis.Card{
		kanjis.NewBuilder('世').Build(),
		kanjis.NewBuilder('界').Build(),
		kanjis.NewBuilder('見').Build(),
	}

	lesson := books.Lesson{WordCards: wordList, KanjiCards: kanjiList}

	testCases := []struct {
		name string
		card content.Identifier
		want bool
	}{
		{
			name: "word in lesson",
			card: wordList[1],
			want: true,
		},
		{
			name: "word not in lesson",
			card: &words.Card{Meaning: "to differ", Nihongo: "違う"},
			want: false,
		},
		{
			name: "kanji in lesson",
			card: kanjiList[0],
			want: true,
		},
		{
			name: "kanji not in lesson",
			card: kanjis.NewBuilder('違').Build(),
			want: false,
		},
		{
			name: "not a content object",
			card: stringWithID("not found"),
			want: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			got := lesson.Contains(c.card)
			if got != c.want {
				t.Errorf("ERROR: got %v, want %v", got, c.want)
			}
		})
	}
}

// stringWithID enables creating an object that implements
// content.Identifier, but does not match the defined japanese
// content types.
type stringWithID string

// ID returns the string.
func (s stringWithID) ID() string {
	return string(s)
}
