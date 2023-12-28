package learn

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestMakeWordCard(t *testing.T) {

	inputCardVerb := &words.Card{
		Nihongo:     "習います",
		Kana:        "ならいます",
		Romaji:      "naraimasu",
		Meaning:     "to learn",
		DictForm:    "習う",
		TeForm:      "習って",
		NaiForm:     "習わない",
		Hint:        "learn from somebody else",
		Explanation: "to study is benkyoo (勉強)",
	}

	cand := []struct {
		mode  string
		input *words.Card
		want  *Card
	}{
		{
			mode:  Native2Japanese,
			input: inputCardVerb,
			want: &Card{
				Question: "to learn",
				Hint:     "learn from somebody else",
				Answer: []string{
					"習います",
					"ならいます",
					"naraimasu",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCardVerb,
			},
		},
		{
			mode:  Japanese2Native,
			input: inputCardVerb,
			want: &Card{
				Question: "習います",
				Hint:     "learn from somebody else",
				Answer: []string{
					"to learn",
					"ならいます",
					"naraimasu",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCardVerb,
			},
		},
		{
			mode:  Native2Kana,
			input: inputCardVerb,
			want: &Card{
				Question: "to learn",
				Hint:     "learn from somebody else",
				Answer: []string{
					"ならいます",
					"naraimasu",
					"習います",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCardVerb,
			},
		},
		{
			mode:  Kana2Native,
			input: inputCardVerb,
			want: &Card{
				Question: "ならいます",
				Hint:     "learn from somebody else",
				Answer: []string{
					"to learn",
					"naraimasu",
					"習います",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCardVerb,
			},
		},
		{
			mode:  "invalid",
			input: inputCardVerb,
			want:  &Card{},
		},
	}

	for _, c := range cand {
		t.Run(c.mode, func(t *testing.T) {
			got := makeWordCard(c.mode, c.input)
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}
