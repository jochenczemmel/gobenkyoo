package learn

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jochenczemmel/gobenkyoo/content/words"
)

func TestMakeWordCard(t *testing.T) {

	// verb card with verb forms filled
	// hint and explanation filled
	inputCard1 := &words.Card{
		Nihongo:     "習います",
		Kana:        "ならいます",
		Romaji:      "naraimasu",
		Meaning:     "to learn",
		DictForm:    "習う",
		TeForm:      "習って",
		NaiForm:     "習わない",
		Hint:        "from somebody",
		Explanation: "to study is benkyoo (勉強)",
	}

	// noun card, verb forms empty
	// hint and explanation empty
	inputCard2 := &words.Card{
		Nihongo: "世界",
		Kana:    "せかい",
		Romaji:  "sekai",
		Meaning: "world",
	}

	cand := []struct {
		mode  string
		input *words.Card
		want  *Card
	}{
		{
			mode:  Native2Japanese,
			input: inputCard2,
			want: &Card{
				Question: "world",
				Answer:   "世界",
				MoreAnswers: []string{
					"せかい",
					"sekai",
				},
				WordCard: inputCard2,
			},
		},
		{
			mode:  Native2Japanese,
			input: inputCard1,
			want: &Card{
				Question: "to learn",
				Hint:     "from somebody",
				Answer:   "習います",
				MoreAnswers: []string{
					"ならいます",
					"naraimasu",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCard1,
			},
		},
		{
			mode:  Japanese2Native,
			input: inputCard1,
			want: &Card{
				Question: "習います",
				Hint:     "from somebody",
				Answer:   "to learn",
				MoreAnswers: []string{
					"ならいます",
					"naraimasu",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCard1,
			},
		},
		{
			mode:  Native2Kana,
			input: inputCard1,
			want: &Card{
				Question: "to learn",
				Hint:     "from somebody",
				Answer:   "ならいます",
				MoreAnswers: []string{
					"naraimasu",
					"習います",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCard1,
			},
		},
		{
			mode:  Kana2Native,
			input: inputCard1,
			want: &Card{
				Question: "ならいます",
				Hint:     "from somebody",
				Answer:   "to learn",
				MoreAnswers: []string{
					"naraimasu",
					"習います",
					"習う",
					"習って",
					"習わない",
				},
				Explanation: "to study is benkyoo (勉強)",
				WordCard:    inputCard1,
			},
		},
		{
			mode:  "invalid",
			input: inputCard1,
			want:  &Card{},
		},
	}

	for _, c := range cand {
		t.Run(c.mode+" "+c.input.Meaning, func(t *testing.T) {
			got := makeWordCard(c.mode, c.input)
			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("ERROR: got- want+\n%s", diff)
			}
		})
	}
}
