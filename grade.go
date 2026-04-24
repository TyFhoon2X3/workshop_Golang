package main

import "fmt"

func main() {
	result := grade()
	fmt.Println(result)
}

func calculateGrade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"

	}
}
func calculateCredit(score int, credit int) int {
	if score >= 80 {
		return credit * 4
	} else if score >= 70 {
		return credit * 3
	} else if score >= 60 {
		return credit * 2
	} else if score >= 50 {
		return credit * 1
	} else {
		return 0

	}

}

func grade() string {

	fmt.Println("welcome to my calculateGrade")
	fmt.Print("Input number of subjects: ")
	var subject int
	fmt.Scan(&subject)
	names := []string{}
	scores := []int{}
	GradeType := []string{}
	creadits := []int{}
	totalCredits := []int{}

	for i := 0; i < subject; i++ {
		var name string
		var score int
		var credit int
		fmt.Print("Input subject name: ")
		fmt.Scan(&name)
		fmt.Printf("Input score for %s: ", name)
		fmt.Scan(&score)
		fmt.Printf("Input credit for %s: ", name)
		fmt.Scan(&credit)

		charGrade := calculateGrade(score)
		totalCredit := calculateCredit(score, credit)
		names = append(names, name)
		scores = append(scores, score)
		GradeType = append(GradeType, charGrade)
		totalCredits = append(totalCredits, totalCredit)
		creadits = append(creadits, credit)

		

	}

	for i := 0; i < subject; i++ {
		fmt.Printf("Subject %d: %-10s | Score: %d | Credit: %d | Grade: %s | Total Credit: %d\n", i+1, names[i], scores[i], creadits[i], GradeType[i], totalCredits[i])
	}	

	return "Process Completed Successfully"

}
