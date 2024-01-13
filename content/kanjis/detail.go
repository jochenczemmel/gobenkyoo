package kanjis

// Detail holds a single reading with a list of meanings.
type Detail struct {
	Reading        string
	ReadingKana    string
	Meanings       []string
	uniqueMeanings map[string]bool
}

// newDetail creates a new detail object with the given romaji
// reading and the meanings.
func newDetail(reading string, meanings ...string) Detail {
	result := Detail{
		Reading:        reading,
		uniqueMeanings: map[string]bool{},
	}
	result.addMeanings(meanings...)

	return result
}

// newDetailWithlKana creates a new detail object with the given romaji
// and kana reading and the meanings.
func newDetailWithlKana(reading, kana string, meanings ...string) Detail {
	result := newDetail(reading, meanings...)
	result.ReadingKana = kana

	return result
}

// addMeanings adds meanings if the do not yet exist.
func (d *Detail) addMeanings(meanings ...string) {
	for _, m := range meanings {
		if !d.uniqueMeanings[m] {
			d.uniqueMeanings[m] = true
			d.Meanings = append(d.Meanings, m)
		}
	}
}
