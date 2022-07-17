# go-future

go-future is an implementation that blocks current thread only when response is needed like Future object in Java.

## Installation

`go get github.com/esoytekin/go-future`

## Usage

```go

import "fmt"

func main() {
	// create new instance
	ft := future.NewFutureTask(func() (string, error) {
		// callback implemantation details
		return "", nil
	})

	// get response
	response, err := ft.Get()

	// check err
	if err != nil {
		panic(err)
	}

	// check if completed
	isComplete := ft.IsComplete()
	fmt.Println(isComplete == true)

	// check if completed with error
	hasError := ft.HasError()
	fmt.Println(hasError == false)

	fmt.Print(response)

}
```
