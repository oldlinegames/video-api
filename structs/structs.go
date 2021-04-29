package structs

import (
	"sync"
)

type VideoQueue struct {
	Queue []string
	Mux sync.Mutex
}