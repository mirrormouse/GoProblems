package problems

func Mypow(x, a int) int {
	res := 1
	for i := 0; i < a; i++ {
		res *= x
	}
	return res
}
