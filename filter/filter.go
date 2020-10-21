package filter

// Filter defines a filter
type Filter struct {
	track         []int
	lowerBand     int
	upperBand     int
	adjustedTrack []int
}

// struct is sort of like a class.

// NewFilter creates a new filter
func NewFilter(track []int, lowerBand int, upperBand int) Filter {

	// You normally use a constructor like this (structs can't be intialised like in ruby) only when you need to perform logic before creating the struct, like validation of input parameters e.g making sure track isn't empty or that lower band isn't a negative number. Otherwise you can do the below code directly where you would have called this function.

	filter := Filter{
		track:     track,
		lowerBand: lowerBand,
		upperBand: upperBand,
	}

	return filter
}

// You can't use default values in parameters so to have 'default' band settings you need a new method.

// DefaultFilter creates a new filter with default lower and upper band settings.
func DefaultFilter(track []int) Filter {
	filter := Filter{
		track:     track,
		lowerBand: 40,
		upperBand: 1000,
	}
	return filter
}

// ApplyFilter applies the filter to the track
func (f *Filter) ApplyFilter() {
	// f is like self. or this. in ruby/JS. The * in front of Filter creates a pointer for f and is required if you want to change a variable on f (the struct).

	for _, band := range f.track {

		switch {
		case band >= f.lowerBand && band <= f.upperBand:
			f.adjustedTrack = append(f.adjustedTrack, band)
		case band < f.lowerBand:
			f.adjustedTrack = append(f.adjustedTrack, f.lowerBand)
		case band > f.upperBand:
			f.adjustedTrack = append(f.adjustedTrack, f.upperBand)
		}
	}

}

// ReturnTrack returns the track with a filter applied
func (f Filter) ReturnTrack() []int {
	return f.adjustedTrack
}
