package words_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestWordCard(t *testing.T) {
	const id1 = "id1"
	word := words.NewBuilder(id1).
		SetContent(
			"分かります",
			"わかります",
			"wakarimasu",
			"verstehen, begreifen").
		SetHint("kann auch 'wissen' bedeuten").
		SetExplanation("oft in hiragana geschrieben").
		SetVerbForms(
			"わかる",
			"わかって",
			"わからない").
		SetContentType("Verb").
		Build()

	cases := []struct {
		name           string
		variable, want string
	}{
		{"Identifier", word.Identifier, id1},
		{"Nihongo", word.Nihongo, "分かります"},
		{"Kana", word.Kana, "わかります"},
		{"Romaji", word.Romaji, "wakarimasu"},
		{"Meaning", word.Meaning, "verstehen, begreifen"},
		{"Hint", word.Hint, "kann auch 'wissen' bedeuten"},
		{"Explanation", word.Explanation, "oft in hiragana geschrieben"},
		{"DictForm", word.DictForm, "わかる"},
		{"TeForm", word.TeForm, "わかって"},
		{"NaiForm", word.NaiForm, "わからない"},
		{"ContentType", word.ContentType, "Verb"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.variable
			if got != c.want {
				t.Errorf("ERROR: got %q, want %q", got, c.want)
			}
		})
	}

	t.Run("ID()", func(t *testing.T) {
		got := word.ID()
		if got != id1 {
			t.Errorf("ERROR: got %q, want %q", got, id1)
		}
	})
}
