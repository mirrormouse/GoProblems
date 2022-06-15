package problems

import (
	"fmt"
	"strconv"
)

func FizzBuzz2Single(n, a, b int) string {
	if n%(a*b) == 0 {
		return "FizzBuzz"
	} else if n%a == 0 {
		return "Fizz"
	} else if n%b == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}
}

func FizzBuzz2(n, a, b int) {

	for i := 1; i <= n; i++ {
		fmt.Printf(FizzBuzz2Single(i, a, b) + " ")
	}

}
