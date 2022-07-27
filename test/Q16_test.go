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
	BoardgameTest{[][]string{{"A", "B", "O"}, {"O", "A", "A"}, {"B", "O", "O"}}, 1},
	BoardgameTest{[][]string{{"A", "B", "B", "A"}, {"O", "A", "A", "B"}, {"B", "B", "O", "A"}, {"B", "O", "O", "O"}}, -1},
	BoardgameTest{[][]string{{"O", "B", "B", "A", "A"}, {"O", "O", "A", "B", "O"}, {"B", "O", "O", "A", "A"}, {"B", "B", "O", "O", "B"}, {"O", "A", "A", "O", "O"}}, 0},
	BoardgameTest{[][]string{{"O", "B", "B", "B", "B"}, {"O", "O", "B", "B", "O"}, {"A", "O", "O", "A", "A"}, {"B", "B", "O", "O", "B"}, {"O", "A", "A", "O", "O"}}, -4},
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
