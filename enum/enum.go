package enum

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

func All[T any](list []T) bool {
	panic("All not ready!")
}

func AllBy[T any](list []T, fn func(T) bool) bool {
	panic("AllBy not ready!")
}

func Any[T any](list []T) bool {
	panic("Any not ready!")
}

func AnyBy[T any](list []T, fn func(T) bool) bool {
	panic("AnyBy not ready!")
}

func At[T any](list []T, index int) T {
	if index < 0 {
		index = len(list) + index
	}
	return list[index]
}

func Chunk[T any](list []T, size int) [][]T {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	chunksNum := len(list) / size
	if len(list)%size != 0 {
		chunksNum += 1
	}

	result := make([][]T, 0, chunksNum)
	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(list) {
			last = len(list)
		}
		result = append(result, list[i*size:last])
	}

	return result
}

func ChunkBy[T any](list []T, fn func(T) bool) [][]T {
	needSplit := func(prev T, current T) bool {
		return fn(prev) != fn(current)
	}

	result := make([][]T, 0)
	for i := 1; i < len(list); i++ {
		prev, current := list[i-1], list[i]
		if needSplit(prev, current) {
			result = append(result, []T{current})
		} else {
			result[len(result)-1] = append(result[len(result)-1], current)
		}
	}
	return result
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

func Dedup[T comparable](list []T) []T {
	result := make([]T, 0)
	result = append(result, list[0])
	for i := 1; i < len(list); i++ {
		prev, current := list[i-1], list[i]
		if prev != current {
			result = append(result, current)
		}
	}
	return result
}

func DedupBy[T any](list []T, fn func(T) bool) []T {
	result := make([]T, 0)
	if fn(list[0]) {
		result = append(result, list[0])
	}
	for i := 1; i < len(list); i++ {
		prev, current := list[i-1], list[i]
		if fn(prev) != fn(current) {
			result = append(result, current)
		}
	}
	return result

}

func Drop[T any](list []T, amount int) []T {
	if amount < len(list) {
		return list[amount:]
	}
	return []T{}
}

func Each[T any](list []T, fn func(T)) {
	for _, v := range list {
		fn(v)
	}
}

func Empty[T any](list []T) bool {
	return len(list) == 0
}

func Fetch[T any](list []T, index int) (T, bool) {
	if int(math.Abs(float64(index))) >= len(list) {
		return list[0], false
	}

	return list[index], true
}

func Find[T any](list []T, fn func(T) bool) (T, bool) {
	for _, v := range list {
		if fn(v) {
			return v, true
		}
	}

	return list[0], false
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

func Sort[T constraints.Ordered](list []T, orderType string) []T {
	listClone := make([]T, 0)
	listClone = append(listClone, list...)
	sort.Slice(listClone, func(i, j int) bool {
		if orderType == "asc" {
			return list[i] < list[j]
		} else if orderType == "desc" {
			return list[i] >= list[j]
		} else {
			panic("Second parameter must be asc or desc")
		}
	})
	return listClone
}

func Uniq[T comparable](list []T) []T {
	result := make([]T, 0, len(list))
	seen := make(map[T]struct{}, len(list))

	for _, item := range list {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

func UniqBy[T any, U comparable](list []T, fn func(T) U) []T {
	result := make([]T, 0, len(list))
	seen := make(map[U]struct{}, len(list))

	for _, item := range list {
		key := fn(item)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}
