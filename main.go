package main

import "fmt"

func main() {
	students := make(map[string]float64)

	students["Jayaprakash"] = 65
	students["Ramkumar"] = 95
	students["Kalai"] = 85
	students["Muthamizh"] = 75
	students["Kodi"] = 98
	var totalScore float64
	var totalStudentsCount float64
	for name, mark := range students {
		// fmt.Println(name, mark)
		fmt.Printf("%s: %v\n", name, mark)
		totalScore = totalScore + mark
		totalStudentsCount++
	}
	var averageScore float64

	averageScore = totalScore / totalStudentsCount

	// fmt.Println("Average Marks:" averageScore )
	fmt.Printf("Average Score: %v\n", averageScore)
	fmt.Println("Students Above Average")

	// for name, mark := range students {
	// 	if mark > averageScore {
	// 		fmt.Printf("%s\n", name)
	// 	}
	// }

	for name, _ := range students {
		if students[name] > averageScore {
			fmt.Println(name)

		}
	}

}
