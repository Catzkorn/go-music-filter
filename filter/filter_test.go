package filter_test

import (
	"testing"

	"github.com/Catzkorn/go-music-filter/filter"
)

// To test we can use $go test in the directory the code/test is in, or go test ./... from the top level to run all tests.

// Testing convention is `func NameOfTestInPascalCase(t *testing.T){testing code in here}`

func TestLowerBandFilter(t *testing.T) {
	track := filter.NewFilter([]int{15}, 20, 50)
	track.ApplyFilter()
	editedTrack := track.ReturnTrack()

	// Testing in go looks more like normal code where you do logical checks with if statements. We test for length here because the second if would panic if there wasn't a [0] element in the slice. Also, arrays are called slices in Go.
	if len(editedTrack) != 1 {
		t.Fatalf("Length of track was not one, got: %d, want: %d.", len(editedTrack), 1) // the %d is a stand in for an integer, which is defined after the string - like SQL injection vulnerability protection! :D
	}
	if editedTrack[0] != 20 {
		t.Fatalf("Frequency was not corrected to lower band limit, got %d, want: %d", editedTrack[0], 20)
	}

	// t.FatalF is a type of error, there is also t.ErrorF.
}

func TestBandWithinRangeFilter(t *testing.T) {
	track := filter.NewFilter([]int{45}, 20, 50)
	track.ApplyFilter()
	editedTrack := track.ReturnTrack()
	if len(editedTrack) != 1 {
		t.Fatalf("Length of track was not one, got: %d, want: %d.", len(editedTrack), 1)
	}
	if editedTrack[0] != 45 {
		t.Fatalf("Frequency was changed unexpectedly, got %d, want: %d", editedTrack[0], 45)
	}
}

func TestUpperBandFilter(t *testing.T) {
	track := filter.NewFilter([]int{1000}, 20, 50)
	track.ApplyFilter()
	editedTrack := track.ReturnTrack()
	if len(editedTrack) != 1 {
		t.Fatalf("Length of track was not one, got: %d, want: %d.", len(editedTrack), 1)
	}
	if editedTrack[0] != 50 {
		t.Fatalf("Frequency was not corrected to upper band limit, got %d, want: %d", editedTrack[0], 50)
		// In go, you write your expected arguments (if you want to)
	}
}

func TestMultipleFrequencyInputs(t *testing.T) {
	track := filter.NewFilter([]int{60, 10, 45, 60, 1500}, 20, 50)
	track.ApplyFilter()
	editedTrack := track.ReturnTrack()
	if len(editedTrack) != 5 {
		t.Fatalf("Length of track was not one, got: %d, want: %d.", len(editedTrack), 5)
	}
	want := []int{50, 20, 45, 50, 50}
	for i, band := range want {
		if band != editedTrack[i] {
			t.Errorf("Filter was not correctly applied to track, at index %d got %d, want: %d", i, editedTrack[i], band)
		}
		// This compares each element of want which is the expected output (in the loop each element is named band) against each element editedTrack, and raises an error if it is not what is expected.
	}
}

func TestFrequenciesOnLimit(t *testing.T) {
	track := filter.NewFilter([]int{20, 10, 50, 55, 1500}, 20, 50)
	track.ApplyFilter()
	editedTrack := track.ReturnTrack()
	if len(editedTrack) != 5 {
		t.Fatalf("Length of track was not one, got: %d, want: %d.", len(editedTrack), 5)
	}
	want := []int{20, 20, 50, 50, 50}
	for i, band := range want {
		if band != editedTrack[i] {
			t.Errorf("Filter was not correctly applied to track, at index %d got %d, want: %d", i, editedTrack[i], band)
		}
	}
}

func TestDefaultFilters(t *testing.T) {
	track := filter.DefaultFilter([]int{20, 10, 50, 55, 1500})
	track.ApplyFilter()
	editedTrack := track.ReturnTrack()
	if len(editedTrack) != 5 {
		t.Fatalf("Length of track was not one, got: %d, want: %d.", len(editedTrack), 5)
	}
	want := []int{40, 40, 50, 55, 1000}
	for i, band := range want {
		if band != editedTrack[i] {
			t.Errorf("Filter was not correctly applied to track, at index %d got %d, want: %d", i, editedTrack[i], band)
		}
	}
}

// In my ruby code, I would have tested that the methods would raise an error if the array contained a nil value. Because Go has static typing, it won't even let the program compile and it is not necessary to test for.
