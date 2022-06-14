package problems

import (
	problems "problems/question"
	"strings"
	"testing"
)

// 入力と期待出力を定義
type Mod4Test struct {
	in  int
	out string
}

// 入力と期待出力の一覧
var Mod4Tests = []Mod4Test{
	Mod4Test{7, "three"},
	Mod4Test{12, "zero"},
	Mod4Test{5, "one"},
	Mod4Test{14, "two"},
}

func Test04(t *testing.T) {

	for i, st := range Mod4Tests {
		v := strings.ToLower(problems.Mod4(st.in))
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is %d", st.out, v, st.in)
			} else {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
