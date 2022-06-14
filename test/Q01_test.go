package problems

import (
	"testing"

	problems "problems/question"
)

// 入力と期待出力を定義
type AveTest struct {
	in1 int
	in2 int
	out int
}

// 入力と期待出力の一覧
var AveTests = []AveTest{
	AveTest{1, 3, 2},
	AveTest{5, 0, 2},
	AveTest{5, -7, -1},
	AveTest{-2, -11, -6},
}

func Test01(t *testing.T) {

	for i, st := range AveTests {
		v := problems.Average(st.in1, st.in2)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %d and %d", st.out, v, st.in1, st.in2)
			} else {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
