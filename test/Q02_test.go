package problems

import (
	"testing"

	problems "problems/question"
)

// 入力と期待出力を定義
type MaxTest struct {
	in1, in2, in3, out int
}

// 入力と期待出力の一覧
var MaxTests = []MaxTest{
	MaxTest{3, 9, 7, 9},
	MaxTest{2, 2, 2, 2},
	MaxTest{0, -5, 3, 3},
	MaxTest{-4, -1, -4, -1},
}

func Test02(t *testing.T) {

	for i, st := range MaxTests {
		v := problems.Max(st.in1, st.in2, st.in3)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %d , %d and %d", st.out, v, st.in1, st.in2, st.in3)
			} else {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
