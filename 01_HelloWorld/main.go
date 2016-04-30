package main

import "fmt"

func main() {
	fmt.Println("Hello World!!")
	//fmt.Println(42)
	//fmt.Printf("%d - %b - %x\n",24,24,24)
	//fmt.Printf("%d \t %b \t %#X \n",42,42,42)

	for i := 0; i <= 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q     \n", i, i, i, i)
	}

}
