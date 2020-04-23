package algs

import (
	"log"
)

// use the list structure
var q1 = ListQuestion{
	"Find the second greatest number",
	[]TestIntList{
		TestIntList{"t1", []int{2, 3, 1, 0, 9}, 3},
		TestIntList{"t2", []int{3, 2, 1, 4, 9}, 4},
		TestIntList{"all_equal", []int{3, 3, 3, 3, 3, 3}, 3},
		//TestIntList{"all_equal", []int{3, 3, 3, 3, 3, 3}, MinInt},
		TestIntList{"desc", []int{6, 5, 4, 3, 2, 1}, 5},
		TestIntList{"asc", []int{1, 2, 3, 4, 5}, 4},
		//TestIntList{"1", []int{0}, MinInt},
		//TestIntList{"nil", nil, MinInt},
	},
	nil,
}

// https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Find2ndGreatest(l []int) int {
	var max2, max int = MinInt, MinInt

	if !(l != nil && len(l) > 1) {
		log.Fatal("list should len > 1")
	}
	for _, i := range l {
		log.Println(i)
		if i > max {
			max2 = max
			max = i
		} else if i > max2 {
			max2 = i
		}
	}
	// ! for the last one
	//return maxs[maxIdx]
	return max2
}
