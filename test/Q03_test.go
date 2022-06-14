package problems

import (
	"testing"

	problems "problems/question"
)

// 入力と期待出力を定義
type SquareSumTest struct {
	in, out int
}

// 入力と期待出力の一覧
var SquareSumTests = []SquareSumTest{
	SquareSumTest{3, 14},
	SquareSumTest{5, 55},
	SquareSumTest{7, 140},
	SquareSumTest{11, 506},
}

func Test03(t *testing.T) {

	secret := true
	for i, st := range SquareSumTests {
		v := problems.SquareSum(st.in)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %d", st.out, v, st.in)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
