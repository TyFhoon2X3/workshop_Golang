package main

import "fmt"

type member struct {
	name         string
	subname      string
	age          int
	saraly       float64
	saralyYearly float64
}

func main() {
	setPeplo()
}

func saralyYears(salary float64) float64 {
	return salary * 12
}

func setPeplo() {
	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)

	var members []member

	for i := 0; i < size; i++ {
		var name string
		var subname string
		var age int
		var saraly float64
		fmt.Println()
		fmt.Printf("Enter name of person %d: ", i+1)
		fmt.Scan(&name, &subname)
		fmt.Printf("Enter age for person %d: ", i+1)
		fmt.Scan(&age)
		fmt.Printf("Enter salary for person %d: ", i+1)
		fmt.Scan(&saraly)

		members = append(members, member{
			name:         name,
			subname:      subname,
			age:          age,
			saraly:       saraly,
			saralyYearly: saralyYears(saraly),
		})
		


		// m := member{
		// 	name:         name,
		// 	subname:      subname,
		// 	age:          age,
		// 	saraly:       saraly,
		// 	saralyYearly: saralyYears(saraly),
		// }
		// members = append(members, m)
	}

	for i, m := range members {
		fmt.Printf("Person %d: %s %s | age: %d | salary: %.2f | salary_per_year: %.2f\n", i+1, m.name, m.subname, m.age, m.saraly, m.saralyYearly)
	}
}
