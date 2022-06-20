package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type FibTest struct {
	in1 int
	out int
}

// 入力と期待出力の一覧
var FibTests = []FibTest{
	FibTest{2, 2},
	FibTest{5, 8},
	FibTest{7, 21},
	FibTest{9, 89},
}

func Test11(t *testing.T) {

	secret := true
	for i, st := range FibTests {
		v := problems.Fib(st.in1)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %d", st.out, v, st.in1)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
