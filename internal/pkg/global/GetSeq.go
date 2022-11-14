package global

import (
	"sync"
)

var WG sync.WaitGroup
var Mut sync.Mutex

func GetSeq() int64 {
	Mut.Lock()
	if sequence == 1048576 {
		sequence = 0
	}

	seq := sequence

	sequence++
	Mut.Unlock()
	return seq
}
