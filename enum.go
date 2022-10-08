package gnum

import (
	"golang.org/x/exp/constraints"
)

func All() {
	panic("unimplemented")

}

func Any() {
	panic("unimplemented")
}

func At[T any](list []T, index int) T {
	if index < 0 {
		index = len(list) + index
	}
	return list[index]
}

func Each[T any](list []T, fn func(T)) {
	for _, v := range list {
		fn(v)
	}
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

	for _, v := range list {
		a[v] = 0
	}

	var result []T
	for k := range a {
		result = append(result, k)
	}

	return result
}

func Uniq_by() {
	panic("unimplemented")
}
