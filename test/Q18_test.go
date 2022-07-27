package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type LCMTest struct {
	in1 int
	in2 int
	out int
}

// 入力と期待出力の一覧
var LCMTests = []LCMTest{
	LCMTest{5, 6, 30},
	LCMTest{8, 12, 24},
	LCMTest{7, 21, 21},
	LCMTest{1, 13, 13},
	LCMTest{16, 16, 16},
}

func Test18(t *testing.T) {

	secret := true
	for i, st := range LCMTests {
		v := problems.LCM(st.in1, st.in2)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %d, %d", st.out, v, st.in1, st.in2)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
