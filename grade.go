package main

import "fmt"

type Subject struct {
	Name   string
	Score  int
	Credit float64
	Grade  string
}

func main() {
	result := grade()
	fmt.Println(result)
}

func calculateGrade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 75 {
		return "B+"
	} else if score >= 70 {
		return "B"
	} else if score >= 65 {
		return "C+"
	} else if score >= 60 {
		return "C"
	} else if score >= 55 {
		return "D+"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}
func calculateCredit(score int, credit float64) float64 {
	if score >= 80 {
		return credit * 4.0
	} else if score >= 75 {
		return credit * 3.5
	} else if score >= 70 {
		return credit * 3.0
	} else if score >= 65 {
		return credit * 2.5
	} else if score >= 60 {
		return credit * 2.0
	} else if score >= 55 {
		return credit * 1.5
	} else if score >= 50 {
		return credit * 1.0
	} else {
		return 0.0
	}
}

func grade() string {

	fmt.Println("welcome to my calculateGrade")
	fmt.Print("Input number of subjects: ")
	var subject int
	fmt.Scan(&subject)
	var subjects []Subject
	var totalCredits float64
	var totalPoints float64

	for i := 0; i < subject; i++ {
		var name string
		var score int
		var credit float64
		fmt.Printf("Enter name of subject %d: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Enter score for subject %d: ", i+1)
		fmt.Scan(&score)
		fmt.Printf("Enter credit for subject %d: ", i+1)
		fmt.Scan(&credit)

		sub := Subject{
			Name:   name,
			Score:  score,
			Credit: credit,
			Grade:  calculateGrade(score),
		}
		subjects = append(subjects, sub)
		totalCredits += credit
		totalPoints += calculateCredit(score, credit)
	}

	// Display results
	fmt.Println("\n===== Results =====")
	for _, sub := range subjects {
		fmt.Printf("Subject: %s | Score: %d | Credit: %.0f | Grade: %s\n", sub.Name, sub.Score, sub.Credit, sub.Grade)
	}

	// Calculate GPA
	gpa := totalPoints / totalCredits
	fmt.Printf("\nTotal Credits: %.0f\n", totalCredits)
	fmt.Printf("GPA: %.2f\n", gpa)

	return "Process Completed Successfully"
}
