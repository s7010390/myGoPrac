package cal

func SumOffirst(x int) int {
	if x == 1 {
		return 1
	} else {
		return x + SumOffirst(x-1)
	}
}
