package gnum

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

func All[T any](list []T) bool {
	return true
}

func AllBy[T any](list []T, fn func(T) bool) bool {
	for _, v := range list {
		if fn(v) {
			continue
		} else {
			return false
		}
	}
	return true
}

func Any[T any](list []T) bool {
	return true
}

func AnyBy[T any](list []T, fn func(T) bool) bool {
	for _, v := range list {
		if fn(v) {
			return true
		}
		continue
	}
	return true
}

func At[T any](list []T, index int) T {
	if index < 0 {
		index = len(list) + index
	}
	return list[index]
}

func AtOrElse[T any](list []T, index int, byDefault T) T {
	if int(math.Abs(float64(index))) >= len(list) {
		return byDefault
	}
	return list[index]
}

func ChunkBy[T any](list []T, fn func(T) bool) T {
	panic("not ")
}

func ChunEvery[T any](list []T, fn func(T) bool) T {
	panic("not ")
}

func ChunkWhile[T any](list []T, fn func(T) bool) T {
	panic("not ")
}

func Concat[T any](list1 []T, list2 []T) []T {
	return append(list1, list2...)
}

func Count[T any](list []T) int {
	return len(list)
}

func CountBy[T any](list []T, fn func(T) bool) int {
	cnt := 0

	for _, v := range list {
		if fn(v) {
			cnt++
		}
	}

	return cnt
}

func CountUntil[T any](list []T, limit int) int {
	cnt := 0

	for range list {
		if cnt < limit {
			cnt++
		}
	}

	return cnt
}

func CountUntilBy[T any](list []T, limit int, fn func(T) bool) int {
	cnt := 0

	for _, v := range list {
		if cnt < limit && fn(v) {
			cnt++
		}
	}

	return cnt
}

func Dedup() {
	panic("not ")
}

func DedupBy() {
	panic("not ")
}

func Drop[T any](list []T, amount int) []T {
	if amount < len(list) {
		return list[amount:]

	}
	return []T{}
}

func DropEvery() {
	panic("not ")
}

func DropWhile() {
	panic("not ")
}

func Each[T any](list []T, fn func(T)) {
	for _, v := range list {
		fn(v)
	}
}

func Empty[T any](list []T) bool {
	return len(list) == 0
}

func Fetch[T any](list []T, index int) (T, error) {
	if int(math.Abs(float64(index))) >= len(list) {
		// todo default value should not list[0]
		return list[0], errors.New(":error")
	}

	return list[index], nil
}

func Find[T any](list []T) {

}

func Map[T, V any](list []T, fn func(T) V) []V {
	var result []V

	for _, v := range list {
		result = append(result, fn(v))
	}
	return result
}

func Min[T constraints.Ordered](list []T) T {
	result := list[0]

	for _, v := range list {
		if v < result {
			result = v
		}
	}
	return result
}

func Max[T constraints.Ordered](list []T) T {
	result := list[0]

	for _, v := range list {
		if v > result {
			result = v
		}
	}
	return result
}

func Filter[T any](list []T, fn func(T) bool) []T {
	var result []T

	for _, v := range list {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[T, V any](list []T, initValue V, fn func(x T, acc V) V) V {
	acc := initValue

	for _, v := range list {
		acc = fn(v, acc)
	}
	return acc
}

func Sort() {

}

func Sort2() {
	panic("unimplemented")
}

func Uniq[T comparable](list []T) []T {
	a := map[T]byte{}
	var result []T

	for _, v := range list {
		if _, flag := a[v]; flag {
			continue
		}
		a[v] = 0
		result = append(result, v)
	}

	return result
}

func Uniq_by() {
	panic("unimplemented")
}
