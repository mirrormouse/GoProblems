package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type arraysquareTest struct {
	in1 []int
	out []int
}

// 入力と期待出力の一覧
var arraysquareTests = []arraysquareTest{
	arraysquareTest{[]int{3, 2, 4}, []int{9, 4, 16}},
	arraysquareTest{[]int{0, -2, 5, 8}, []int{0, 4, 25, 64}},
	arraysquareTest{[]int{2, 2, 2, 2, 2}, []int{4, 4, 4, 4, 4}},
	arraysquareTest{[]int{1, 2, 3, -10}, []int{1, 4, 9, 100}},
}

func Test13(t *testing.T) {

	secret := true
	for i, st := range arraysquareTests {
		v := problems.Arraysquare(st.in1)
		if !compslice(v, st.out) {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%s, Your Answer:%s. \nGiven Argument is %s", slicetostring(st.out), slicetostring(v), slicetostring(st.in1))
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
