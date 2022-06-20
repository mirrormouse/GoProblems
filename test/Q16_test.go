package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type BoardgameTest struct {
	in1 [][]string
	out int
}

// 入力と期待出力の一覧
var BoardgameTests = []BoardgameTest{
	BoardgameTest{[][]string{{"A"}, {"O"}, {"B"}}, 0},
	BoardgameTest{[][]string{{"A"}, {"O"}, {"B"}}, 0},
	BoardgameTest{[][]string{{"A"}, {"O"}, {"B"}}, 0},
}

func Test16(t *testing.T) {

	secret := true
	for i, st := range BoardgameTests {
		v := problems.Boardgame(st.in1)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %s", st.out, v, slice2strtostring(st.in1))
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
