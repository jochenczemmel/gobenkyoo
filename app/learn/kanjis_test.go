package learn

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jochenczemmel/gobenkyoo/content/kanjis"
)

func TestMakeKanjiCard(t *testing.T) {

	kb := kanjis.NewBuilder('人')
	kb.AddDetailsWithKana("JIN", "ジン", "Mensch", "Person", "Leute")
	kb.AddDetailsWithKana("NIN", "ニン", "Mensch")
	kb.AddDetailsWithKana("hito", "ひと", "Mensch", "Person", "Leute")
	inputCard := kb.Build()

	testCases := []struct {
		mode  string
		input *kanjis.Card
		want  *Card
	}{
		{
			mode:  Kanji2Native,
			input: inputCard,
			want: &Card{
				Question: "人",
				Answer:   "Mensch, Person, Leute",
				MoreAnswers: []string{
					"JIN, NIN, hito",
					"ジン, ニン, ひと",
				},
				KanjiCard: inputCard,
			},
		},
		/*
			{
				mode:  "invalid",
				input: inputCard,
				want:  &Card{},
			},
		*/
	}

	for _, c := range testCases {
		t.Run(c.mode+" "+c.input.Kanji(), func(t *testing.T) {
			got := makeKanjiCard(c.mode, c.input)
			if diff := cmp.Diff(got, c.want,
				cmpopts.IgnoreUnexported(kanjis.Card{})); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}
