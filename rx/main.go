package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/reactivex/rxgo"
)

func main() {
	// func1()
	// func2()
	func4()
}

func func1() {
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}
}

func func2() {
	observable := rxgo.Just(1, 2, errors.New("joker error"), 4, 5)()
	<-observable.ForEach(func(v interface{}) {
		fmt.Println("received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})
}

func func3() {
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
		next <- rxgo.Error(errors.New("unknown"))
		next <- rxgo.Of(4)
		next <- rxgo.Of(5)
	}})

	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}

func func4() {
	observable := rxgo.Just(1, 2, 3, 4)()

	observable = observable.BufferWithCount(3)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
