package problems

import "fmt"

func FizzBuzz(n int) {

	for i := 1; i <= n; i++ {
		fmt.Printf(FizzBuzzSingle(i) + " ")
	}

}
