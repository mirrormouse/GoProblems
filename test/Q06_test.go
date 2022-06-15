package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type PowTest struct {
	in1 int
	in2 int
	out int
}

// 入力と期待出力の一覧
var PowTests = []PowTest{
	PowTest{2, 3, 8},
	PowTest{2, 0, 1},
	PowTest{-3, 3, -27},
	PowTest{-5, 4, 625},
	PowTest{-7, 0, 1},
	PowTest{-1, 13, -1},
	PowTest{0, 20, 0},
}

func Test06(t *testing.T) {

	secret := true
	for i, st := range PowTests {
		v := problems.Mypow(st.in1, st.in2)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %d and %d", st.out, v, st.in1, st.in2)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
