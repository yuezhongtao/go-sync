package main

import (
	"go-sync"
)

func main() {
	l := go_sync.Mutex{}

	l.Lock()
	l.Unlock()

	l.Lock()
	l.Unlock()

}
