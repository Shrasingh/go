package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to a class on slices")
	var fruitList = []string{"apple", "banana", "grapes", "peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fruitList = append(fruitList, "mango", "banana")
	fmt.Println("fruit list is ", fruitList)
	fmt.Println("length of array is ", len(fruitList))
	fruitList = append(fruitList[1:3])
	fmt.Println("fruit list is ", fruitList)
	fmt.Println("length of array is ", len(fruitList))
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 900
	highScores[2] = 567
	highScores[3] = 789
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	highScores = append(highScores, 555, 777, 222)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	scores := make([]int, len(highScores))
	copy(scores, highScores)
	fmt.Println("scores ", scores)
	fmt.Println("highscores ", highScores)
	fmt.Println("are scores and highScores equal ", (scores == nil))
	fmt.Println("are scores and highScores equal ", (len(scores) == len(highScores)))
	// how to remove a value from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
