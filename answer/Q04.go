package problems

func Mod4(n int) string {
	switch n % 4 {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	default:
		return "zero"
	}
}
