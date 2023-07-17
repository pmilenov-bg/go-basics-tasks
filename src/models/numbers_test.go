package models

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates_Scenario1(t *testing.T) {
	inputSlice := []int{1, 2, 2, 3, 4, 4, 4, 5, 6, 7, 7, 10, 32, 44, 44, 55}
	expectedUniqueSlice := []int{1, 2, 3, 4, 5, 6, 7, 10, 32, 44, 55}
	expectedRepeated := []int{2, 4, 7, 44}

	uniqueSlice, repeated := RemoveDuplicates(inputSlice)

	if !reflect.DeepEqual(uniqueSlice, expectedUniqueSlice) {
		t.Errorf("RemoveDuplicates() returned unique slice %v, expected %v", uniqueSlice, expectedUniqueSlice)
	}

	if !reflect.DeepEqual(repeated, expectedRepeated) {
		t.Errorf("RemoveDuplicates() returned repeated slice %v, expected %v", repeated, expectedRepeated)
	}
}

func TestRemoveDuplicates_Scenario2(t *testing.T) {
	inputSlice := []int{1, 1, 1, 1, 1}
	expectedUniqueSlice := []int{1}
	expectedRepeated := []int{1}

	uniqueSlice, repeated := RemoveDuplicates(inputSlice)

	if !reflect.DeepEqual(uniqueSlice, expectedUniqueSlice) {
		t.Errorf("RemoveDuplicates() returned unique slice %v, expected %v", uniqueSlice, expectedUniqueSlice)
	}

	if !reflect.DeepEqual(repeated, expectedRepeated) {
		t.Errorf("RemoveDuplicates() returned repeated slice %v, expected %v", repeated, expectedRepeated)
	}
}

func TestRemoveDuplicates_Scenario3(t *testing.T) {
	inputSlice := []int{}
	expectedUniqueSlice := []int{}
	expectedRepeated := []int{}

	uniqueSlice, repeated := RemoveDuplicates(inputSlice)

	if !reflect.DeepEqual(uniqueSlice, expectedUniqueSlice) {
		t.Errorf("RemoveDuplicates() returned unique slice %v, expected %v", uniqueSlice, expectedUniqueSlice)
	}

	if !reflect.DeepEqual(repeated, expectedRepeated) {
		t.Errorf("RemoveDuplicates() returned repeated slice %v, expected %v", repeated, expectedRepeated)
	}
}
