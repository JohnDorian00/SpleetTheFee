package testing

import (
	"fmt"
)

func Main() {
	fmt.Println("hello from testing polygon")
	defer func() {
		const size = 40
		printSeparator := func(size int) {
			for i := 0; i < size; i++ {
				fmt.Print("_")
				if i == size-1 {
					fmt.Println()
				}
			}
		}
		for i := 0; i < size/10; i++ {
			fmt.Println()
		}
		printSeparator(size)
		fmt.Println("testing is over")
		printSeparator(size)
	}()
	fmt.Println(sum(3, 5))
}

func sum(a, b int) int {
	return a + b
}
