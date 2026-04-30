package main

import (
	"fmt"
)

type Subject struct {
	Name   string
	Score  int
	Credit float64
	Grade  float64
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
	var subjects []Subject
	var subject int
	var totalCredits float64
	var totalPoints float64

	fmt.Println("welcome to my calculateGrade")
	fmt.Print("Input number of subjects: ")
	fmt.Scan(&subject)
	for i := 0; i < subject; i++ {
		var name string
		fmt.Printf("Input name of subject %d: ", i+1)
		fmt.Scan(&name)

		var score int
		fmt.Printf("Input score of subject %d: ", i+1)
		fmt.Scan(&score)

		var credit float64
		fmt.Printf("Input credit of subject %d: ", i+1)
		fmt.Scan(&credit)

		subjects = append(subjects, Subject{
			Name:   name,
			Score:  score,
			Credit: credit,
			Grade:  calculateCredit(score, credit),
		})

		totalCredits = totalCredits + credit
		totalPoints = totalPoints + calculateCredit(score, credit)

	}

	for i := 0; i < subject; i++ {
		fmt.Printf("Subject: %s, Score: %d, Credit: %.2f, Grade: %.2f\n", subjects[i].Name, subjects[i].Score, subjects[i].Credit, subjects[i].Grade)

	}

	gpa := totalPoints / totalCredits
	fmt.Printf("Your GPA is: %.2f\n", gpa)

	if gpa >= 3.5 {
		return "Congratulations! You have an excellent GPA!"
	} else if gpa >= 3.0 {
		return "Great job! You have a good GPA!"
	} else if gpa >= 2.0 {
		return "Not bad! You have a decent GPA!"
	} else {
		return "You need to work harder! Your GPA is below average."
	}

}
