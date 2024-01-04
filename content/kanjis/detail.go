package kanjis

// detail holds a single reading with a list of meanings.
type detail struct {
	reading        string
	readingKana    string
	meanings       []string
	uniqueMeanings map[string]bool
}

// newDetail creates a new detail object with the given romaji
// reading and the meanings.
func newDetail(reading string, meanings ...string) detail {
	result := detail{
		reading:        reading,
		uniqueMeanings: map[string]bool{},
	}
	result.addMeanings(meanings...)

	return result
}

// newDetailWithlKana creates a new detail object with the given romaji
// and kana reading and the meanings.
func newDetailWithlKana(reading, kana string, meanings ...string) detail {
	result := newDetail(reading, meanings...)
	result.readingKana = kana

	return result
}

// addMeanings adds meanings if the do not yet exist.
func (d *detail) addMeanings(meanings ...string) {
	for _, m := range meanings {
		if !d.uniqueMeanings[m] {
			d.uniqueMeanings[m] = true
			d.meanings = append(d.meanings, m)
		}
	}
}
