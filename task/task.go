package task

import (
	"github.com/reugn/async"
	"github.com/samber/mo"
)

func Async[V any](fn func() mo.Result[V]) async.Future[V] {
	promise := async.NewPromise[V]()
	go func() {
		res := fn()
		if res.IsOk() {
			promise.Success(res.MustGet())
		}
		if res.IsError() {
			promise.Failure(res.Error())
		}
	}()
	return promise.Future()
}

func Await[T any](task async.Future[T]) mo.Result[T] {
	res, err := task.Join()
	return mo.TupleToResult(res, err)
}
