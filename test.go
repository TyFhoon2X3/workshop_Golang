package main

import "fmt"

func setPeplo() {
	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)

	var names []string
	var grades []int

	for i := 0; i < size; i++ {
		var name string
		var grade int
		fmt.Printf("\nEnter name of subject %d: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Enter grade for subject %d: ", i+1)
		fmt.Scan(&grade)

		names = append(names, name)
		grades = append(grades, grade)
	}
	fmt.Printf("Names in the array: %d\n", len(names))
	for i := 0; i < len(names); i++ {
		var result string

		if grades[i] >= 80 {
			result = "A"
		} else if grades[i] >= 60 {
			result = "B"
		} else if grades[i] >= 40 {
			result = "C"
		} else {
			result = "D"
		}

		// พิมพ์ทุกอย่างออกมาในบรรทัดเดียวให้สวยงาม
		fmt.Printf("Subject %d: %-10s | Grade: %d | Result: %s\n", i+1, names[i], grades[i], result)
	}
}
