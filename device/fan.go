package device

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int, 1)
	var sum int
	for _, c := range chs {
		sum += <-c
	}
	out <- sum
	return out
}
