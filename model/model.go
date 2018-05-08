package model

import (
	"sync"
)

var (
	DB    map[string][]interface{}
	mutex *sync.Mutex
)

func init() {
	mutex = new(sync.Mutex)
	mutex.Lock()
	DB = make(map[string][]interface{})
	DB["server"] = make([]interface{}, 0, 10)
	DB["channel"] = make([]interface{}, 0, 10)
	mutex.Unlock()
}
