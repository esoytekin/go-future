package future

// FutureTask type for FutureTask
type FutureTask[V any] struct {
	result V
	error  error
	signal chan struct{}
}

// Get blocks current thread and waits for result
func (t *FutureTask[V]) Get() (V, error) {
	<-t.signal
	return t.result, t.error
}

// IsComplete check if future task is done
func (t *FutureTask[V]) IsComplete() bool {
	select {
	case <-t.signal:
		return true
	default:
	}

	return false
}

// HasError check if task finished with error
func (t *FutureTask[V]) HasError() bool {
	return t.error != nil
}

// NewFutureTask returns new *FutureTask instance
func Start[V any](callback func() (V, error)) *FutureTask[V] {
	f := new(FutureTask[V])
	f.signal = make(chan struct{})

	go func() {
		defer close(f.signal)
		f.result, f.error = callback()
	}()

	return f
}
