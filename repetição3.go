package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%02 == 1 {
			fmt.Println(i)
		}
	}

}
