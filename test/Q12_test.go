package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type arraysumTest struct {
	in1 [5]int
	out int
}

// 入力と期待出力の一覧
var arraysumTests = []arraysumTest{
	arraysumTest{[5]int{3, 2, 4, 0, 0}, 9},
	arraysumTest{[5]int{0, -2, 5, 8, 0}, 11},
	arraysumTest{[5]int{2, 2, 2, 2, 2}, 10},
}

func array5tostring(s [5]int) string {
	t := "["
	for i := 0; i < 5; i++ {
		t += string(s[i])
		if i > 0 {
			t += ","
		}
	}
	t += "]"
	return t
}

func Test12(t *testing.T) {

	secret := true
	for i, st := range arraysumTests {
		v := problems.Arraysum(st.in1)
		if v != st.out {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%d, Your Answer:%d. \nGiven Argument is ", st.out, v, array5tostring(st.in1))

				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
