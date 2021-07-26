package main

import "testing"

func Test_maxIndex(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		IndexWant  int
		MaxElementWant int
	}{
		{"All of numbers of positive", []int{1, 4, 25, 25, 100}, 100, 4},
		{"All of numbers is negative", []int{-1, -40, -10, -1, -5}, 0, -1},
		{"All of numbers is positive and negative", []int{-1,5,-20, 2}, 5, 1},
	}

	for _, test := range tests {
		index, element := maxIndex(test.slice)
		if element != test.MaxElementWant || index != test.IndexWant {
			t.Errorf("maxIndex test %s gotindex %d IndexWant %d gotElement %d Elementwant %d", test.name, element, test.MaxElementWant, index, test.IndexWant)
		}
	}
}
