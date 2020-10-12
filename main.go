package main

import "fmt"

func main() {
	printPattern(num)
}

func printPattern(num int) {
	if num%2 != 1 {
		fmt.Println("input harus berupa bilangan ganjil")
	} else {
		fmt.Println("====== Panjang ======")
		for i := 1; i < num+1; i++ {
			for j := 1; j < num+1; j++ {
				if j == 1 || j == num || i == (num+1)/2 {
					fmt.Print("*  ")
				} else {
					fmt.Print("=  ")
				}
			}
			fmt.Println("\n ")
		}
	}
}