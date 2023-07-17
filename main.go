package main

import (
	"fmt"
	"sale-hire/projects/golang-journey/go-basic-tasks/src/models"
)

var prl = fmt.Println

func main() {
	inputSlice := []int{1, 2, 2, 3, 4, 4, 4, 5, 6, 7, 7, 10, 32, 44, 44, 55}
	uniqueSlice, repeated := models.RemoveDuplicates(inputSlice)
	prl(uniqueSlice)
	prl(repeated)

	inputStrings := []string{"apple", "banana", "apple", "orange", "banana", "banana"}
	occurrences := models.CountOccurrences(inputStrings)
	fmt.Println(occurrences)

	map1 := map[string]int{"apple": 2, "banana": 3}
	map2 := map[string]int{"banana": 4, "orange": 1}
	mergedMap := models.MergeMaps(map1, map2)
	fmt.Println(mergedMap)

	rect := models.Rectangle{Width: 5.0, Height: 3.0}
	fmt.Println("Rectangle Width:", rect.Width, "Height:", rect.Height)
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())

	circle := models.Circle{Radius: 2.5}
	fmt.Println("Circle Radius:", circle.Radius)
	fmt.Println("Area:", circle.Area())
	fmt.Println("Circumference:", circle.Perimeter())

	prl("In the interface section")
	models.PrintShapeDetails(rect)
	models.PrintShapeDetails(circle)
}
