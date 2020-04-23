package algs

import (
	"fmt"

	"github.com/pennz/DataViz/lists/arraylist"
	"github.com/pennz/DataViz/trees/binaryheap"
)

type WantIntList struct {
	name   string
	input  []int
	output interface{}
}

type Solver interface {
	Solve() // solve with your algs
}

type ListQuestion struct { // testing related
	dscpt string // should instead use testing
	wants []WantIntList
	alg   interface{}
}

func (q *ListQuestion) Init() {
	fmt.Println("Test")
	return
}

// Check if your algs is right
func (q *ListQuestion) Check() error {
	return nil
}

func (q *ListQuestion) Solve() {
	fmt.Println("Test")
	return
}

func CreateListQuestion(values ...interface{}) {
	l := arraylist.New()
	l.Add(values)
	return
}

func TreeGetInfo() {
	var bh *binaryheap.Heap = binaryheap.NewWithIntComparator()
	fmt.Println(bh)
}
