package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type StringTest struct {
	in1 string
	in2 string
	out string
}

// 入力と期待出力の一覧
var StringTests = []StringTest{
	StringTest{"Hello", "World", "HelloWorld!"},
	StringTest{"Hello", "Go", "HelloGo!"},
	StringTest{"Go", "Program", "GoProgram!"},
	StringTest{"Keep", "Practice", "KeepPractice!"},
}

func Test05(t *testing.T) {

	secret := true
	for i, st := range StringTests {
		v := problems.StringConcat(st.in1, st.in2)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%s, Your Answer:%s. \nGiven Argument is %s and %s", st.out, v, st.in1, st.in2)
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
