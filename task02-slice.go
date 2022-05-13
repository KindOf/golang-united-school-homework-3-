package homework

func reverse(input []int64) (result []int64) {
	var s []int64

	for i := len(input) - 1; i >= 0; i-- {
		s = append(s, input[i])
	}
	return s
}
