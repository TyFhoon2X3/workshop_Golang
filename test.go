package main 

import "fmt"

func main() { 
    test()
}

func test() {
    i := 1
    for i <= 10 {
        if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
			
		}else {
			fmt.Printf("%d is odd\n", i)
		}

        i++
    }
}