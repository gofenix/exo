package task

import (
	"fmt"
	"testing"
	"time"

	"github.com/samber/mo"
)

func TestAsync(t *testing.T) {
	asyncBlock()
}

func asyncBlock() {
	fn := func() mo.Result[string] {
		fmt.Println("this is before")
		time.Sleep(1 * time.Second)
		fmt.Println("this is after")
		return mo.Ok("hello world")
	}

	t := Async(fn)
	fmt.Println("this is task: ", t)

	res := Await(t)
	fmt.Println("this is result: ", res)

	time.Sleep(10 * time.Second)
}
