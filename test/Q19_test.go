package problems

import (
	problems "problems/question"
	"strings"
	"testing"
)

// 入力と期待出力を定義
type NabeatsuTest struct {
	in  int
	out string
}

// 入力と期待出力の一覧
var NabeatsuTests = []NabeatsuTest{
	NabeatsuTest{5, "1 2 3! 4 5 6! 7 8 9! 10"},
}

func Test19(t *testing.T) {

	secret := true
	for i, st := range NabeatsuTests {
		f := func() func() {
			return func() { problems.Nabeatsu(st.in) }
		}
		v := strings.TrimSpace(extractStdout(t, f()))
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%s, Your Answer:%s. \nGiven Argument is %d", st.out, v, st.in)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
