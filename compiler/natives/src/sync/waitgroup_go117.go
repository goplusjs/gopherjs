//go:build js && !go1.18
// +build js,!go1.18

package sync

type WaitGroup struct {
	counter int
	ch      chan struct{}

	state1 [3]uint32
	sema   uint32
}

func (wg *WaitGroup) Add(delta int) {
	wg.counter += delta
	if wg.counter < 0 {
		panic("sync: negative WaitGroup counter")
	}
	if wg.counter > 0 && wg.ch == nil {
		wg.ch = make(chan struct{})
	}
	if wg.counter == 0 && wg.ch != nil {
		close(wg.ch)
		wg.ch = nil
	}
}

func (wg *WaitGroup) Wait() {
	if wg.counter > 0 {
		<-wg.ch
	}
}
