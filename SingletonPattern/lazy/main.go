package lazy

import "sync"

var once *sync.Once

type singleton struct {
}

var single *singleton

func GetInstance() *singleton {
	once.Do(func() {
		single = new(singleton)
	})
	return single
}
