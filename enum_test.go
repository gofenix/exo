package gnum

import (
	"fmt"
	"testing"
)

func TestChunk(t *testing.T) {
	rest := Chunk([]int{1, 2, 3, 5}, 1)
	fmt.Println(rest)
}

func TestChunkBy(t *testing.T) {
	a := []int{1, 2, 2, 3, 4, 4, 6, 7, 7}
	result := ChunkBy(a, func(x int) bool {
		return x%2 == 1
	})
	fmt.Println(result)
}

func TestAt(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(At(a, 1))
}

func TestDedup(t *testing.T) {
	a := []int{1, 1, 1, 2, 3, 4, 4, 7}
	res := Dedup(a)
	fmt.Println(res)
}

func TestDedupBy(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 4, 4, 8, 7, 5}
	res := DedupBy(a, func(x int) bool {
		return x%2 == 1
	})
	fmt.Println(res)
}

func TestSort(t *testing.T) {
	a := []int{1, 2, 3, 444, 3}
	res := Sort(a, "asc")
	fmt.Println(res)
	res = Sort(a, "desc")
	fmt.Println(res)
}

func TestUniq(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 4, 4, 8, 7, 5}
	res := Uniq(a)
	fmt.Println(res)
}
