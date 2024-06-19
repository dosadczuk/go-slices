package slices_test

import (
	"fmt"
	"net/netip"

	"github.com/dosadczuk/go-slices"
)

func ExampleAll() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := slices.All(values, func(num int) bool {
		return num > 5
	})
	fmt.Println(result)

	// Output:
	// false
}

func ExampleAny() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := slices.Any(values, func(num int) bool {
		return num > 0
	})
	fmt.Println(result)

	// Output:
	// true
}

func ExampleAssociate() {
	values := []int{1, 2, 3, 4, 5}

	result := slices.Associate(values, func(num int) (int, bool) {
		return num, num%2 == 0
	})
	fmt.Println(result)

	// Output:
	// map[1:false 2:true 3:false 4:true 5:false]
}

func ExampleContains() {
	values := []int{1, 2, 3, 4, 5}

	result := slices.Contains(values, 2)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleEqual() {
	v1 := []int{1, 2, 3, 4, 5}
	v2 := []int{1, 2, 3, 4, 5}

	result := slices.Equal(v1, v2)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleEqualFunc() {
	v1 := []fmt.Stringer{
		netip.AddrFrom4([4]byte{127, 0, 0, 1}),
		netip.AddrFrom4([4]byte{255, 0, 0, 0}),
	}
	v2 := []fmt.Stringer{
		netip.AddrFrom4([4]byte{127, 0, 0, 1}),
		netip.AddrFrom4([4]byte{255, 0, 0, 0}),
	}

	result := slices.EqualFunc(v1, v2, func(v1, v2 fmt.Stringer) bool {
		return v1.String() == v2.String()
	})
	fmt.Println(result)

	// Output:
	// true
}

func ExampleFilter() {
	values := []int{1, 2, 3, 4, 5}

	result := slices.Filter(values, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(result)

	// Output:
	// [2 4]
}

func ExampleFind() {
	values := []int{1, 2, 3, 4, 5}

	result, found := slices.Find(values, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(result, found)

	// Output:
	// 2 true
}

func ExampleFindIndex() {
	values := []int{1, 2, 3, 4, 5}

	result, found := slices.FindIndex(values, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(result, found)

	// Output:
	// 1 true
}

func ExampleFindLast() {
	values := []int{1, 2, 3, 4, 5}

	result, found := slices.FindLast(values, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(result, found)

	// Output:
	// 4 true
}

func ExampleFindLastIndex() {
	values := []int{1, 2, 3, 4, 5}

	result, found := slices.FindLastIndex(values, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(result, found)

	// Output:
	// 3 true
}

func ExampleFirst() {
	values := []int{1, 2, 3, 4, 5}

	result, found := slices.First(values)
	fmt.Println(result, found)

	// Output:
	// 1 true
}

func ExampleFlatten() {
	values := [][]int{{1}, {1, 2}, {2, 3}}

	result := slices.Flatten(values)
	fmt.Println(result)

	// Output:
	// [1 1 2 2 3]
}

func ExampleFold() {
	values := []string{"1", "2", "3", "4", "5"}

	result := slices.Fold(values, "order: ", func(acc string, val string) string {
		return acc + val
	})
	fmt.Println(result)

	// Output:
	// order: 12345
}

func ExampleFoldRight() {
	values := []string{"1", "2", "3", "4", "5"}

	result := slices.FoldRight(values, "order: ", func(acc string, val string) string {
		return acc + val
	})
	fmt.Println(result)

	// Output:
	// order: 54321
}

func ExampleForEach() {
	values := []int{1, 2, 3, 4, 5}

	slices.ForEach(values, func(idx int, num int) {
		fmt.Printf("%v:%v\n", idx, num)
	})

	// Output:
	// 0:1
	// 1:2
	// 2:3
	// 3:4
	// 4:5
}

func ExampleGroupBy() {
	values := []int{1, 2, 3, 4}

	result := slices.GroupBy(values, func(num int) int {
		return num % 2
	})
	fmt.Println(result)

	// Output:
	// map[0:[2 4] 1:[1 3]]
}

func ExampleIndexOf() {
	values := []int{1, 4, 3, 4, 5}

	result := slices.IndexOf(values, 4)
	fmt.Println(result)

	// Output:
	// 1
}

func ExampleLast() {
	values := []int{1, 2, 3, 4, 5}

	result, found := slices.Last(values)
	fmt.Println(result, found)

	// Output:
	// 5 true
}

func ExampleLastIndexOf() {
	values := []int{1, 4, 3, 4, 5}

	result := slices.LastIndexOf(values, 4)
	fmt.Println(result)

	// Output:
	// 3
}

func ExampleMap() {
	values := []int{1, 2, 3, 4, 5}

	result := slices.Map(values, func(num int) int {
		return 10 - num
	})
	fmt.Println(result)

	// Output:
	// [9 8 7 6 5]
}

func ExampleMax() {
	values := []int{3, 2, 1, 4, 5}

	result := slices.Max(values)
	fmt.Println(result)

	// Output:
	// 5
}

func ExampleMaxFunc() {
	values := []int{3, 2, 1, 4, 5}

	result := slices.MaxFunc(values, func(a, b int) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		} else {
			return +1
		}
	})
	fmt.Println(result)

	// Output:
	// 5
}

func ExampleMin() {
	values := []int{3, 2, 1, 4, 5}

	result := slices.Min(values)
	fmt.Println(result)

	// Output:
	// 1
}

func ExampleMinFunc() {
	values := []int{3, 2, 1, 4, 5}

	result := slices.MinFunc(values, func(a, b int) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		} else {
			return +1
		}
	})
	fmt.Println(result)

	// Output:
	// 1
}

func ExampleNone() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := slices.None(values, func(num int) bool {
		return num > 0
	})
	fmt.Println(result)

	// Output:
	// false
}

func ExamplePartition() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result1, result2 := slices.Partition(values, func(num int) bool {
		return num < 6
	})
	fmt.Println(result1, result2)

	// Output:
	// [1 2 3 4 5] [6 7 8 9 10]
}

func ExampleReduce() {
	values := []string{"1", "2", "3", "4", "5"}

	result := slices.Reduce(values, func(acc string, val string) string {
		return acc + val
	})
	fmt.Println(result)

	// Output:
	// 12345
}

func ExampleReduceRight() {
	values := []string{"1", "2", "3", "4", "5"}

	result := slices.ReduceRight(values, func(acc string, val string) string {
		return acc + val
	})
	fmt.Println(result)

	// Output:
	// 54321
}

func ExampleReverse() {
	values := []int{1, 2, 3, 4, 5}

	slices.Reverse(values)
	fmt.Println(values)

	// Output:
	// [5 4 3 2 1]
}

func ExampleReversed() {
	values := []int{1, 2, 3, 4, 5}

	result := slices.Reversed(values)
	fmt.Println(result)

	// Output:
	// [5 4 3 2 1]
}
