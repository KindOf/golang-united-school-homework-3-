package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var l int

	for _, n := range input {
		if n > 0 {
			l++
			sum += n
		}
	}

	return sum / float32(l)
}
