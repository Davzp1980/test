package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	switch {
	case a >= 0 && a <= 10:
		for i := 0; i <= 10; i++ {
			if i == a {
				switch i {
				case 0:
					fmt.Println("Zero")
				case 1:
					fmt.Println("One")
				case 2:
					fmt.Println("Two")
				case 3:
					fmt.Println("Three")
				case 4:
					fmt.Println("Four")
				case 5:
					fmt.Println("Five")
				case 6:
					fmt.Println("Six")
				case 7:
					fmt.Println("Seven")
				case 8:
					fmt.Println("Eight")
				case 9:
					fmt.Println("Nine")
				case 10:
					fmt.Println("Ten")
				}
			}
		}
		fallthrough
	case b >= 0 && b <= 10:
		for i := 0; i <= 10; i++ {
			if i == b {
				switch i {
				case 0:
					fmt.Println("Zero")
				case 1:
					fmt.Println("One")
				case 2:
					fmt.Println("Two")
				case 3:
					fmt.Println("Three")
				case 4:
					fmt.Println("Four")
				case 5:
					fmt.Println("Five")
				case 6:
					fmt.Println("Six")
				case 7:
					fmt.Println("Seven")
				case 8:
					fmt.Println("Eight")
				case 9:
					fmt.Println("Nine")
				case 10:
					fmt.Println("Ten")
				}
			}
		}
		fallthrough

	case c >= 0 && c <= 10:
		for i := 0; i <= 10; i++ {
			if i == c {
				switch i {
				case 0:
					fmt.Println("Zero")
				case 1:
					fmt.Println("One")
				case 2:
					fmt.Println("Two")
				case 3:
					fmt.Println("Three")
				case 4:
					fmt.Println("Four")
				case 5:
					fmt.Println("Five")
				case 6:
					fmt.Println("Six")
				case 7:
					fmt.Println("Seven")
				case 8:
					fmt.Println("Eight")
				case 9:
					fmt.Println("Nine")
				case 10:
					fmt.Println("Ten")
				}
			}
		}
	default:
		fmt.Println("Input error")
	}

}
