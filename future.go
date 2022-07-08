package future

// FutureTask type for FutureTask
type FutureTask struct {
	result interface{}
	error  error
	signal chan struct{}
}

// Get blocks current thread and waits for result
func (t *FutureTask) Get() (interface{}, error) {
	<-t.signal
	return t.result, t.error
}

// IsComplete check if future task is done
func (t *FutureTask) IsComplete() bool {
	select {
	case <-t.signal:
		return true
	default:
	}

	return false
}

// HasError check if task finished with error
func (t *FutureTask) HasError() bool {
	return t.error != nil
}

// NewFutureTask returns new *FutureTask instance
func NewFutureTask(callback func() (interface{}, error)) *FutureTask {
	f := new(FutureTask)
	f.signal = make(chan struct{})

	go func() {
		defer close(f.signal)
		f.result, f.error = callback()
	}()

	return f
}
