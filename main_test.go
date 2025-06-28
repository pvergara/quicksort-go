package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestWhenIAddNumberToAQuicksortCollectionAndAskForSortTheSystemReturnsMeTheCollectionSorted(t *testing.T) {
	got := make([]int, 0)
	got = append(got, 10, 9, 3, 4, 5, 7, 6, 1, 4, 2)

	quickSortCollection := new(QuickSortCollection)
	quickSortCollection.addAll(got)
	var expected = quickSortCollection.sort()

	want := append([]int(nil), expected...)
	sort.Ints(want)
	if len(expected) != len(got) {
		t.Errorf("got and expedted has different sizes")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
	fmt.Println(expected)
	fmt.Println(want)
}
