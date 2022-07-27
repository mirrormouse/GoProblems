package problems

import (
	problems "problems/question"
	"testing"
)

// 入力と期待出力を定義
type SelectstringTest struct {
	in1 []string
	in2 []string
	out []int
}

// 入力と期待出力の一覧
var SelectstringTests = []SelectstringTest{
	SelectstringTest{[]string{"apple", "orange", "apple", "apple", "orange"}, []string{"apple", "orange"}, []int{3, 2}},
	SelectstringTest{[]string{"one", "one", "two", "three", "four", "two", "three", "two"}, []string{"two", "four"}, []int{3, 1}},
	SelectstringTest{[]string{"Monday", "Monday", "Tuesday", "Friday", "Saturday", "Thursday", "Monday", "Monday"}, []string{"Monday", "Friday", "Sunday"}, []int{4, 1, 0}},
	SelectstringTest{[]string{"ab", "bc", "ac", "ab", "ab", "cd", "ac", "ac", "ab"}, []string{"cd", "ab", "cb"}, []int{1, 4, 0}},
}

func Test17(t *testing.T) {

	secret := true
	for i, st := range SelectstringTests {
		v := problems.Selectstring(st.in1, st.in2)
		if compslice(v, st.out) {
			if i < 3 { //３つ目まではオープンケースとして誤りを通知
				t.Errorf("Expected Answer:%s, Your Answer:%d. \nGiven Argument is %s, %s", slicetostring(st.out), v, slicestrtostring(st.in1), slicestrtostring(st.in2))
				secret = false
			} else if secret {
				t.Errorf("Secret Case Failed")
			}
		}
	}
}
