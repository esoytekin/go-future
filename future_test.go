package future

import (
	"errors"
	"testing"
	"time"
)

func TestFutureTask(t *testing.T) {
	ft := NewFutureTask(func() (bool, error) {
		return true, nil
	})

	response, err := ft.Get()

	if err != nil {
		t.Error("unexpected error")
	}

	if !response {
		t.Error("unexpected response")
	}

	if !ft.IsComplete() {
		t.Error("unexpected completion status")
	}

	if ft.HasError() {
		t.Error("unexpected error in future")
	}

}

func TestFutureTaskWithBlockingResponse(t *testing.T) {
	var futureItems []*FutureTask[int]
	for i := 0; i < 5; i++ {
		ft := func(i int) *FutureTask[int] {
			return NewFutureTask(func() (int, error) {
				time.Sleep(5 * time.Second)
				if i == 0 {
					return 0, errors.New("got 0")
				}
				return i, nil
			})
		}(i)

		futureItems = append(futureItems, ft)
	}

	for i, fi := range futureItems {
		response, err := fi.Get()

		if err != nil {
			if response != 0 {
				t.Errorf("expected 0 indexed item to have error, got %d", response)
			}

			if !fi.HasError() {
				t.Error("should had error")
			}
		}

		if response != i {
			t.Errorf("expected %d got %d", i, response)
		}
	}
}
