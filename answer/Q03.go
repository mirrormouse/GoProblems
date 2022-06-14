package problems

func Square(x int) int {
	return x * x
}

func SquareSum(n int) int {
	var a int
	a = 0
	for i := 1; i < n+1; i++ {
		a += Square(i)
	}
	return a

}
