package homework

func average(input [6]float32) (result float32) {
	var sum float32
	var l int

	for _, n := range input {
		l++
		sum += n
	}
	return sum / float32(l)
}
