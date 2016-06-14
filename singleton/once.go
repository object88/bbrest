package singleton

import (
	"sync"
	"sync/atomic"
)

// Once is a singleton function wrapper
// See http://marcio.io/2015/07/singleton-pattern-in-go/
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do ensures that for a given `Once` instance, the function
// is only invoked once.
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
