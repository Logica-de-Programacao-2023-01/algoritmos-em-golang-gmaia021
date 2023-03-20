package main

import "fmt"

func main() {
	for i := 0; i < 21; i++ {
		if i%02 == 0 {
			fmt.Println(i)
		}
	}

}
