package main

import (
	"fmt"

	"github.com/zhenfeng-zhu/gnum"
)

func main() {
	l := []string{"a", "b", "c"}
	a := gnum.At([]string{"a", "b", "c"}, -1)
	fmt.Println(a)

	gnum.Each(l, func(x string) { fmt.Println(x) })

	b := gnum.Map(l, func(x string) int {
		return 0
	})
	fmt.Println(b)

	m := gnum.Min(l)
	fmt.Println(m)

	fmt.Println(gnum.Filter(l, func(x string) bool {
		if x == "a" {
			return true
		}
		return false
	}))

	fmt.Println(
		gnum.Reduce(
			l,
			"",
			func(x, acc string) string {
				return x + acc
			},
		),
	)

	fmt.Println(
		gnum.Uniq([]int{1, 2, 2, 3, 4, 1, 1}),
	)
}
