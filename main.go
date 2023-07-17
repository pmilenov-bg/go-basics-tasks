package main

import (
	"fmt"
	"math"
)

var prl = fmt.Println

func removeDuplicates(input []int) ([]int, []int) {
	uniqueSet := make(map[int]bool)
	repeated := make(map[int]bool)
	result, result1 := []int{}, []int{}

	for _, num := range input {
		if !uniqueSet[num] {
			uniqueSet[num] = true
			result = append(result, num)
		} else {
			if !repeated[num] {
				repeated[num] = true
				result1 = append(result1, num)
			}
		}
	}

	return result, result1
}

func countOccurrences(input []string) map[string]int {
	occurrences := make(map[string]int)

	for _, str := range input {
		occurrences[str]++
	}

	return occurrences
}

func mergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] += value
	}

	return mergedMap
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	inputSlice := []int{1, 2, 2, 3, 4, 4, 4, 5, 6, 7, 7, 10, 32, 44, 44, 55}
	uniqueSlice, repeated := removeDuplicates(inputSlice)
	prl(uniqueSlice)
	prl(repeated)

	inputStrings := []string{"apple", "banana", "apple", "orange", "banana", "banana"}
	occurrences := countOccurrences(inputStrings)
	fmt.Println(occurrences)

	map1 := map[string]int{"apple": 2, "banana": 3}
	map2 := map[string]int{"banana": 4, "orange": 1}
	mergedMap := mergeMaps(map1, map2)
	fmt.Println(mergedMap)

	rect := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Println("Rectangle Width:", rect.Width, "Height:", rect.Height)
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())

	circle := Circle{Radius: 2.5}
	fmt.Println("Circle Radius:", circle.Radius)
	fmt.Println("Area:", circle.Area())
	fmt.Println("Circumference:", circle.Perimeter())

	prl("In the interface section")
	PrintShapeDetails(rect)
	PrintShapeDetails(circle)
}
