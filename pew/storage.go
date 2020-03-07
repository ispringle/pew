package pew

import "sync"

type AlertStorage struct {
	sync.Mutex
	AlertMessages []*Message
}
