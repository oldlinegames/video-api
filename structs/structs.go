package structs

import (
	"sync"
	//this syncs it.
)

type VideoQueue struct {
	Queue []string
	//this queues it.
	Mux sync.Mutex
}

type VideoUpload struct {
	Films []string `json:"videos"`
	//this films it.
}