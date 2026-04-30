	package main

	import (
		"fmt"
		"os"
		"os/exec"
		"runtime"
	)

	func CallClear() {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			// ต้องสั่งผ่าน cmd /c เพื่อเรียกคำสั่งภายในอย่าง cls
			cmd = exec.Command("cmd", "/c", "cls")
		} else {
			cmd = exec.Command("clear")
		}
		cmd.Stdout = os.Stdout
		cmd.Run()
	}


	

	type Node struct {
		data int
		next *Node
	}

	func main() {
		const (
			ColorReset  = "\033[0m" // ล้างค่าสี กลับมาเป็นสีปกติ
			ColorRed    = "\033[31m"
			ColorGreen  = "\033[32m"
			ColorYellow = "\033[33m"
			ColorBlue   = "\033[34m"
			ColorPurple = "\033[35m"
			ColorCyan   = "\033[36m"
		)
		var head *Node
		var number int
		for {
			// CallClear() // เคลียร์ก่อนเริ่มเมนูใหม่เสมอ
			fmt.Println(ColorCyan + "--- LINKED LIST MANAGER ---" + ColorReset)
			fmt.Println(ColorYellow + "1: Add Node")
			fmt.Println("2: Find Node")
			fmt.Println("3: Show All Nodes")
			fmt.Println("4: Delete Nodes")
			fmt.Println("5: Chang Nodes")
			fmt.Println("6: Insert betwed")
			fmt.Println(ColorRed + "7: Exit" + ColorReset)
			fmt.Print("Select option: ")
			fmt.Scan(&number)

			switch number {
			case 1:
				fmt.Print(ColorBlue + "enter node do u want : " + ColorReset)
				var n int
				fmt.Scan(&n)

				for i := 0; i < n; i++ {
					var value int
					fmt.Printf(ColorBlue+"enter value[%d] : "+ColorReset, i+1)
					fmt.Scan(&value)
					Newnode := &Node{data: value}
					if head == nil {
						head = Newnode
					} else {
						currency := head
						for currency.next != nil {
							currency = currency.next
						}
						currency.next = Newnode

					}
				}

			case 2:
				var target int
				fmt.Print("Enter value to search: ")
				fmt.Scan(&target)

				run := head
				Tig := false
				Index := 1
				for run != nil {
					if target == run.data {
						fmt.Printf(ColorGreen+"Found %d at index %d\n"+ColorReset, target, Index)
						Tig = true
					}
					run = run.next
					Index++
				}
				if !Tig {
					fmt.Println(ColorRed + "ไม่เจอจ้าาาาา" + ColorReset)
				}
				// --- จุดสำคัญ: ใส่ Pause ---
				fmt.Print("\nPress Enter to continue...")
				fmt.Scanln() // ล้างบัฟเฟอร์
				fmt.Scanln() // รอรับ Enter
			case 3:
				if head == nil {
					fmt.Println(ColorRed + "List is empty!" + ColorReset)
				} else {
					fmt.Print(ColorPurple + "Current List: " + ColorReset)
					run := head
					for run != nil {
						fmt.Printf("%d => ", run.data)
						run = run.next
					}
					fmt.Println(ColorRed + "NULL" + ColorReset)
				}
				fmt.Print("\nPress Enter to continue...")
				fmt.Scanln()
				fmt.Scanln()
			case 4:
				if head == nil {
					fmt.Println(ColorRed + "List ว่างเปล่า ลบไม่ได้จ้า!" + ColorReset)
				} else {
					var target int
					fmt.Print("ใส่เลขที่ต้องการลบ: ")
					fmt.Scan(&target)

					// กรณีที่ 1: ลบตัวแรก (Head)
					if target == head.data {
						head = head.next
						fmt.Println(ColorGreen + "ลบตัวแรกเรียบร้อย!" + ColorReset)
					} else {
						n1 := head
						n2 := head.next
						status := false

						for n2 != nil {
							if n2.data == target {

								n1.next = n2.next
								status = true
							}
							n1 = n2
							n2 = n2.next
						}
						if !status {
							fmt.Print("หาม่าย")
						}
						fmt.Println(ColorGreen + "ลบเรียบร้อย!" + ColorReset)
					}
				}
				fmt.Print("\nPress Enter to continue...")
				fmt.Scanln()
				fmt.Scanln()
			case 5:
				if head == nil {
					fmt.Println(ColorRed + "ไม่มีอะไรให้แก้" + ColorReset)
				} else {
					var target int
					var changNumber int
					fmt.Print("ต้องการแก้ที่เลขอะไร: ")
					fmt.Scan(&target)
					fmt.Print("เลขที่ต้องการแก้: ")
					fmt.Scan(&changNumber)

					if target == head.data {
						head.data = changNumber
						fmt.Println(ColorGreen + "เปลี่ยนเรียบร้อย!" + ColorReset)
					} else {
						run := head
						status := false

						for run != nil {
							if run.data == target {
								run.data = changNumber
								status = true
								fmt.Println(ColorGreen + "เปลี่ยนเรียบร้อย!" + ColorReset)
								break

							}
							run = run.next

						}
						if !status {
							fmt.Println(ColorRed + "ไม่พบเลขที่ต้องการแก้ไข" + ColorReset)
						}

					}
				}
				fmt.Print("\nPress Enter to continue...")
				fmt.Scanln()
				fmt.Scanln()
			case 6: // แทรกหลังตัวเลขที่กำหนด
				var target, value int
				fmt.Print("ต้องการแทรกหลังเลขอะไร: ")
				fmt.Scan(&target)
				fmt.Print("ใส่เลขใหม่ที่ต้องการเพิ่ม: ")
				fmt.Scan(&value)

				newNode := &Node{data: value}
				run := head
				status := false

				for run != nil {
					if run.data == target {
					
						newNode.next = run.next
						run.next = newNode    
				
						status = true
						break
					}
					run = run.next
				}

				if status {
					fmt.Println(ColorGreen + "แทรกเรียบร้อย!" + ColorReset)
				} else {
					fmt.Println(ColorRed + "ไม่เจอเลขเป้าหมาย" + ColorReset)
				}
			case 7:
				fmt.Println(ColorCyan + "Goodbye!" + ColorReset)
				return
			}
		}
	}
