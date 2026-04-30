package main

import "fmt"

func main() {
    choice := 0

    // ใช้ for ครอบไว้เพื่อให้โปรแกรมวนกลับมาเริ่มใหม่เสมอ
    for {
        fmt.Println("\n--- Main Menu ---")
        fmt.Println("1: Square")
        fmt.Println("2: Triangle")
        fmt.Println("3: Exit Program") // เพิ่มทางเลือกสำหรับปิดโปรแกรม
        fmt.Print("Select your choice: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            fmt.Println("Drawing Square...")
            for i := 1; i <= 5; i++ {
                for j := 1; j <= 5; j++ {
                    fmt.Print("* ")
                }
                fmt.Println()
            }
            // เมื่อจบงานใน case 1 จะวนกลับไปที่ต้น loop for ทันที

        case 2:
            fmt.Println("Drawing Triangle...")
            for i := 1; i <= 5; i++ {
                for j := 1; j <= i; j++ {
                    fmt.Print("* ")
                }
                fmt.Println()
            }

        case 3:
            fmt.Println("Exiting... Bye!")
            return // คำสั่งนี้จะทำให้หลุดออกจากฟังก์ชัน main และจบโปรแกรม

        default:
            // ถ้ากดเลขอื่นที่ไม่ใช่ 1, 2, 3
            fmt.Println("-------------------------------")
            fmt.Println("Invalid choice! Please try again.")
            fmt.Println("-------------------------------")
            // ไม่ต้องทำอะไรเพิ่ม ลูป for จะวนกลับไปถามใหม่ให้เอง
        }
    }
}