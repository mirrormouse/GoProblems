package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type Average2Test struct {
	in1 int
	in2 int
	out float32
}

// 入力と期待出力の一覧
var Average2Tests = []Average2Test{
	Average2Test{5, 3, 4},
	Average2Test{6, 9, 7.5},
	Average2Test{-8, 3, -2.5},
	Average2Test{-4, 8, 2},
	Average2Test{-3, -10, -6.5},
}

func Test20(t *testing.T) {

	secret := true
	for i, st := range Average2Tests {
		v := problems.Average2(st.in1, st.in2)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%f, Your Answer:%f. \nGiven Argument is %d and %d", st.out, v, st.in1, st.in2)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
