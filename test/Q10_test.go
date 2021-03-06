package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type factTest struct {
	in1 int
	out int
}

// 入力と期待出力の一覧
var factTests = []factTest{
	factTest{3, 6},
	factTest{5, 120},
	factTest{7, 5040},
	factTest{13, 6227020800},
}

func Test10(t *testing.T) {

	secret := true
	for i, st := range factTests {
		v := problems.Fact(st.in1)
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
