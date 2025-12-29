package main

import "fmt"

type QuickSort struct{}

func (qs QuickSort) Sort(data []int) {
	fmt.Println("Sorting using quicksort")
}

type MergeSort struct{}

func (ms MergeSort) Sort(data []int) {
	fmt.Println("Sorting using merge sort")
}

type SortStrategy interface {
	Sort([]int)
}

type Sorter struct {
	strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(data []int) {
	s.strategy.Sort(data)
}

func main() {
	sorter := Sorter{}

	sorter.SetStrategy(QuickSort{})
	sorter.Sort([]int{3, 1, 4, 1, 5})

	sorter.SetStrategy(MergeSort{})
	sorter.Sort([]int{3, 1, 4, 1, 5})
}
