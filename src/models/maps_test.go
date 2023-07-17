package models

import (
	"reflect"
	"testing"
)

func TestMergeMaps_Scenario1(t *testing.T) {
	map1 := map[string]int{"apple": 2, "banana": 3}
	map2 := map[string]int{"banana": 4, "orange": 1}
	expectedMergedMap := map[string]int{"apple": 2, "banana": 7, "orange": 1}

	mergedMap := MergeMaps(map1, map2)

	if !reflect.DeepEqual(mergedMap, expectedMergedMap) {
		t.Errorf("MergeMaps() returned %v, expected %v", mergedMap, expectedMergedMap)
	}
}

func TestMergeMaps_Scenario2(t *testing.T) {
	map1 := map[string]int{"a": 10, "b": 5}
	map2 := map[string]int{"b": 3, "c": 7}
	expectedMergedMap := map[string]int{"a": 10, "b": 8, "c": 7}

	mergedMap := MergeMaps(map1, map2)

	if !reflect.DeepEqual(mergedMap, expectedMergedMap) {
		t.Errorf("MergeMaps() returned %v, expected %v", mergedMap, expectedMergedMap)
	}
}

func TestMergeMaps_Scenario3(t *testing.T) {
	map1 := map[string]int{}
	map2 := map[string]int{"a": 1, "b": 2, "c": 3}
	expectedMergedMap := map[string]int{"a": 1, "b": 2, "c": 3}

	mergedMap := MergeMaps(map1, map2)

	if !reflect.DeepEqual(mergedMap, expectedMergedMap) {
		t.Errorf("MergeMaps() returned %v, expected %v", mergedMap, expectedMergedMap)
	}
}

func TestMergeMaps_Scenario4(t *testing.T) {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{}
	expectedMergedMap := map[string]int{"a": 1, "b": 2, "c": 3}

	mergedMap := MergeMaps(map1, map2)

	if !reflect.DeepEqual(mergedMap, expectedMergedMap) {
		t.Errorf("MergeMaps() returned %v, expected %v", mergedMap, expectedMergedMap)
	}
}

func TestMergeMaps_Scenario5(t *testing.T) {
	map1 := map[string]int{}
	map2 := map[string]int{}
	expectedMergedMap := map[string]int{}

	mergedMap := MergeMaps(map1, map2)

	if !reflect.DeepEqual(mergedMap, expectedMergedMap) {
		t.Errorf("MergeMaps() returned %v, expected %v", mergedMap, expectedMergedMap)
	}
}

func TestCountOccurrences_Scenario1(t *testing.T) {
	inputStrings := []string{"apple", "banana", "apple", "orange", "banana", "banana"}
	expectedOccurrences := map[string]int{
		"apple":  2,
		"banana": 3,
		"orange": 1,
	}

	occurrences := CountOccurrences(inputStrings)

	if !reflect.DeepEqual(occurrences, expectedOccurrences) {
		t.Errorf("CountOccurrences() returned %v, expected %v", occurrences, expectedOccurrences)
	}
}

func TestCountOccurrences_Scenario2(t *testing.T) {
	inputStrings := []string{"apple", "apple", "apple", "apple"}
	expectedOccurrences := map[string]int{
		"apple": 4,
	}

	occurrences := CountOccurrences(inputStrings)

	if !reflect.DeepEqual(occurrences, expectedOccurrences) {
		t.Errorf("CountOccurrences() returned %v, expected %v", occurrences, expectedOccurrences)
	}
}

func TestCountOccurrences_Scenario3(t *testing.T) {
	inputStrings := []string{"", "", "", ""}
	expectedOccurrences := map[string]int{
		"": 4,
	}

	occurrences := CountOccurrences(inputStrings)

	if !reflect.DeepEqual(occurrences, expectedOccurrences) {
		t.Errorf("CountOccurrences() returned %v, expected %v", occurrences, expectedOccurrences)
	}
}

func TestCountOccurrences_Scenario4(t *testing.T) {
	inputStrings := []string{}
	expectedOccurrences := map[string]int{}

	occurrences := CountOccurrences(inputStrings)

	if !reflect.DeepEqual(occurrences, expectedOccurrences) {
		t.Errorf("CountOccurrences() returned %v, expected %v", occurrences, expectedOccurrences)
	}
}
