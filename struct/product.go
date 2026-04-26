package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Category string
	Discount float64
}

func main() {
	fmt.Print("Enter unit Product: ")
	var unit int
	fmt.Scan(&unit)

	var products []Product

	for i := 0; i < unit; i++ {
		var name string
		var price float64
		var category string
		var discount float64
		fmt.Printf("Enter name of product %d: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Enter price of product %d: ", i+1)
		fmt.Scan(&price)
		fmt.Printf("Enter category of product %d: ", i+1)
		fmt.Scan(&category)
		fmt.Printf("Enter discount of product %d: ", i+1)
		fmt.Scan(&discount)
		products = append(products, Product{
			Name:     name,
			Price:    price,
			Category: category,
			Discount: discount,
		})

	}
	fmt.Println("Products entered:")
	for _, product := range products {
		fmt.Printf("Name: %s, Price: %.2f, Category: %s, Discount: %.2f%%\n",
			product.Name, product.Price, product.Category, product.Discount)
	}

}
