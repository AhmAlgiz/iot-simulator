package device

import (
	"math/rand"
)

func GenerateBaseTemp(base int) <-chan int {
	out := make(chan int, 1)
	out <- base + rand.Intn(3) - 2
	return out
}

func GenerateCondTemp(status bool) <-chan int {
	out := make(chan int, 1)
	out <- -1
	return out
}

func GenerateHeatTemp(status bool) <-chan int {
	out := make(chan int, 1)
	out <- 1
	return out
}

func GenerateBaseHum(base int) <-chan int {
	out := make(chan int, 1)
	out <- base + rand.Intn(10) - 5
	return out
}

func GenerateCondHum(base, point int) <-chan int {
	out := make(chan int, 1)
	delta := base - point
	if delta > 0 {
		out <- -1
	} else if delta == 0 {
		out <- 0
	} else {
		out <- 1
	}
	return out
}
