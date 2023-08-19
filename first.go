package main

import "fmt"

func main() {
	var x int
	for i := 0; i < 3; i++ {
		fmt.Scanln(&x)
		switch x {
		case 1:
			fmt.Println("Один")
		case 2:
			fmt.Println("Два")
		case 3:
			fmt.Println("Три")
		case 4:
			fmt.Println("Четыре")
		case 5:
			fmt.Println("Пять")
		case 6:
			fmt.Println("Шесть")
		case 7:
			fmt.Println("Семь")
		case 8:
			fmt.Println("Восемь")
		case 9:
			fmt.Println("Девять")
		case 10:
			fmt.Println("Десять")
		case 0:
			fmt.Println("Ноль")
		}
	}
}
