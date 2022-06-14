package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type HelloTest struct {
	out string
}

// 入力と期待出力の一覧
var HelloTests = []HelloTest{
	HelloTest{"Hello World"},
}

func Test00(t *testing.T) {

	secret := true
	for i, st := range HelloTests {
		v := extractStdout(t, problems.Hello)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%s, Your Answer:%s.", st.out, v)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
