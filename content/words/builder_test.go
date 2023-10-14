package words_test

import (
	"testing"

	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestWordCard(t *testing.T) {

	word := words.NewBuilder("id1").
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
		name       string
		callGetter func() string
		want       string
	}{
		{"ID", word.ID, "id1"},
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
			got := c.callGetter()
			if got != c.want {
				t.Errorf("ERROR: got %q, want %q", got, c.want)
			}
		})
	}
}
