package device

import (
	"simulator/client"
)

// measuring devices struct
type Meter struct {
	client client.Client
	value  int
	topic  string
}
