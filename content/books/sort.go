package books

// bySeriesVolumeTitle provides sorting of a slice of book ids
// according to series title, volume and book title.
// type bySeriesVolumeTitle []Book
type bySeriesVolumeTitle []ID

// Len implements sort.Interface.
func (b bySeriesVolumeTitle) Len() int { return len(b) }

// Swap implements sort.Interface.
func (b bySeriesVolumeTitle) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements sort.Interface.
// It defines the sort order as described above.
func (b bySeriesVolumeTitle) Less(i, j int) bool {
	if b[i].SeriesTitle < b[j].SeriesTitle {
		return true
	}
	if b[i].SeriesTitle > b[j].SeriesTitle {
		return false
	}
	if b[i].Volume < b[j].Volume {
		return true
	}
	if b[i].Volume > b[j].Volume {
		return false
	}

	return b[i].Title < b[j].Title
}
