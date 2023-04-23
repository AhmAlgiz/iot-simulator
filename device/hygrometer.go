package device

import "math/rand"

type Hygrometer struct {
	Device
	hum int
}

func (h *Hygrometer) GetValue() int {

}

func (h *Hygrometer) GenerateValue() int {

}

func generateBaseHum(base int) <-chan int {
	out := make(chan int, 1)
	out <- base + rand.Intn(10) - 5
	return out
}

func generateCondHum(base, point int) <-chan int {
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
