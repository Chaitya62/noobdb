package buffer

import (
	"fmt"
	"sync"
)

// Ephemeral stays in the volatile storage
// later could make a decision on using sync.Map vs map with locks
// ref: https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c
type PageTable struct {
	index map[uint32]int
	_lock sync.RWMutex
}

func (pt *PageTable) Init() {
	pt.index = make(map[uint32]int)
}

func (pt *PageTable) InsertOrUpdate(key uint32, val int) {
	pt._lock.Lock()
	defer pt._lock.Unlock()
	pt.index[key] = val
}

func (pt *PageTable) Get(key uint32) (int, bool) {
	fmt.Println("Getting Page", key)
	pt._lock.RLock()
	defer pt._lock.RUnlock()
	if val, ok := pt.index[key]; ok {
		return val, true
	}
	return -1, false
}
