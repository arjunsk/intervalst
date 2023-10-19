package main

import (
	"fmt"
	"github.com/arjunsk/intervalst"
)

func main() {
	ist := intervalst.NewIntervalTree()
	ist.Insert(intervalst.NewInt64Interval(1, 17), "pasta")
	ist.Insert(intervalst.NewInt64Interval(4, 18), "oregano")
	ist.Insert(intervalst.NewInt64Interval(10, 19), "cheese")
	//ist.Insert(intervalst.NewInt64Interval(14, 23), "chilly")
	ist.Insert(intervalst.NewInt64Interval(14, 23), "milk")

	// fetch all the ingredients that are in my inventory at 15:00
	// if any of the ingredients is chilly, I can't make the recipe
	var result []string
	cantMake := false
	ist.Visit(intervalst.NewInt64Point(15), func(v *intervalst.IntervalValue) bool {
		result = append(result, v.Val.(string))
		if v.Val.(string) == "chilly" {
			// Here we can do early termination of the search, since we know we can't make the recipe
			cantMake = true
			return false
		}
		return true
	})

	if cantMake {
		fmt.Println("Sorry, I can't make the recipe")
	} else {
		fmt.Println("I can make the recipe with these ingredients:", result)
	}
}
