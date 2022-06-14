package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type FizzBuzz1Test struct {
	in1 int
	out string
}

// 入力と期待出力の一覧
var FizzBuzz1Tests = []FizzBuzz1Test{
	FizzBuzz1Test{7, "7"},
	FizzBuzz1Test{12, "Fizz"},
	FizzBuzz1Test{20, "Buzz"},
	FizzBuzz1Test{30, "FizzBuzz"},
	FizzBuzz1Test{59, "59"},
	FizzBuzz1Test{63, "Fizz"},
	FizzBuzz1Test{65, "Buzz"},
	FizzBuzz1Test{75, "FizzBuzz"},
}

func Test06(t *testing.T) {

	secret := true
	for i, st := range FizzBuzz1Tests {
		v := problems.FizzBuzzSingle(st.in1)
		if v != st.out {
			if i < 4 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%s, Your Answer:%s. \nGiven Argument is %d", st.out, v, st.in1)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
