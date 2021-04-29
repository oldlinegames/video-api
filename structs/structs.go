package structs

import (
	"sync"
)

type VideoQueue struct {
	Queue []string
	Mux sync.Mutex
}

type VideoUpload struct {
	Videos []string `json:"videos"`
}