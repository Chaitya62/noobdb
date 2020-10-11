package buffer

import (
	"sync"
)

type Replacer interface {
	GetNextVictim() (int, bool)
	PinPage(frame_id int)
	UnPinPage(frame_id int)
	Size() uint32
}

type ClockReplacer struct {
	_frames     []bool
	_frames_ref []bool
	_size       uint32
	_capacity   uint32
	size_lock   sync.Mutex
}

func NewClockReplacer(buffer_size uint32) Replacer {
	cr := &ClockReplacer{}
	cr._frames = make([]bool, buffer_size)
	cr._frames_ref = make([]bool, buffer_size)
	cr._size = 0
	cr._capacity = buffer_size

	for i := uint32(0); i < buffer_size; i++ {
		cr._frames[i] = false
		cr._frames_ref[i] = false
	}

	return cr
}

func (cr *ClockReplacer) GetNextVictim() (int, bool) {
	if cr._size == 0 {
		return -1, false
	}

	cr.size_lock.Lock()
	defer cr.size_lock.Unlock()

	var frame_id int

	var i = 0
	for {
		if !cr._frames[i] {
			i = (i + 1) % int(cr._capacity)
			continue
		}

		if !cr._frames_ref[i] {
			frame_id = i
			break
		}

		cr._frames_ref[i] = false

		i = (i + 1) % int(cr._capacity)

	}

	return frame_id, true
}

func (cr *ClockReplacer) PinPage(frame_id int) {
	cr.size_lock.Lock()
	defer cr.size_lock.Unlock()
	cr._frames[frame_id] = false
	cr._size--
}

func (cr *ClockReplacer) UnPinPage(frame_id int) {
	cr.size_lock.Lock()
	defer cr.size_lock.Unlock()
	cr._frames[frame_id] = true
	cr._frames_ref[frame_id] = true
	cr._size++
}

func (cr *ClockReplacer) Size() uint32 {
	cr.size_lock.Lock()
	defer cr.size_lock.Unlock()
	return cr._size
}
