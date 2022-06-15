package problems

import (
	problems "problems/question"
	"strings"
	"testing"
)

// 入力と期待出力を定義
type FizzBuzz2Test struct {
	in1 int
	in2 int
	in3 int
	out string
}

// 入力と期待出力の一覧
var FizzBuzz2Tests = []FizzBuzz2Test{
	FizzBuzz2Test{7, 2, 3, "1 Fizz Buzz Fizz 5 FizzBuzz 7"},
	FizzBuzz2Test{12, 2, 6, "1 Fizz 3 Fizz 5 Fizz 7 Fizz 9 Fizz 11 FizzBuzz"},
	FizzBuzz2Test{12, 4, 6, "1 2 3 Fizz 5 Buzz 7 Fizz 9 10 11 Fizz"},
	FizzBuzz2Test{24, 4, 6, "1 2 3 Fizz 5 Buzz 7 Fizz 9 10 11 Fizz 13 14 15 Fizz 17 Buzz 19 Fizz 21 22 23 FizzBuzz"},
	FizzBuzz2Test{50, 18, 24, "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 Fizz 19 20 21 22 23 Buzz 25 26 27 28 29 30 31 32 33 34 35 Fizz 37 38 39 40 41 42 43 44 45 46 47 Buzz 49 50"},
}

func Test09(t *testing.T) {

	secret := true
	for i, st := range FizzBuzz2Tests {
		f := func() func() {
			return func() { problems.FizzBuzz2(st.in1, st.in2, st.in3) }
		}
		v := strings.TrimSpace(extractStdout(t, f()))
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%s, Your Answer:%s. \nGiven Argument is %d , %d and %d", st.out, v, st.in1, st.in2, st.in3)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
