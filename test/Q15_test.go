package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type evensplitTest struct {
	in1 []int
	out []int
}

// 入力と期待出力の一覧
var evensplitTests = []evensplitTest{
	evensplitTest{[]int{3, 2, 4}, []int{3, 1, 1, 2, 2}},
	evensplitTest{[]int{0, -2, 5, 8}, []int{0, 0, -1, -1, 5, 4, 4}},
	evensplitTest{[]int{2, 2, 2, 2, 2}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	evensplitTest{[]int{1, 2, 3, -12, 0}, []int{1, 1, 1, 3, -6, -6, 0, 0}},
}

func Test15(t *testing.T) {

	secret := true
	for i, st := range evensplitTests {
		v := problems.Evensplit(st.in1)
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
