package kanjis

// Detail holds a single reading with a list of meanings.
type Detail struct {
	Reading     string
	ReadingKana string
	Meanings    []string
}

/*
// NewDetail creates a new detail object with the given romaji
// reading and the meanings.
func NewDetail(reading string, meanings ...string) Detail {
	result := Detail{
		Reading: reading,
	}
	result.AddMeanings(meanings...)

	return result
}

// NewDetailWithKana creates a new detail object with the given romaji
// and kana reading and the meanings.
func NewDetailWithKana(reading, kana string, meanings ...string) Detail {
	result := NewDetail(reading, meanings...)
	result.ReadingKana = kana

	return result
}

// AddMeanings adds meanings if the do not yet exist.
func (d *Detail) AddMeanings(meanings ...string) {
	for _, m := range meanings {
		d.Meanings = append(d.Meanings, m)
	}
}
*/
