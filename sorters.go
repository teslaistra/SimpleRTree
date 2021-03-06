package SimpleRTree

import "fmt"

type xSorter struct {
	n                      *rNode
	points                 FlatPoints
	start, end, bucketSize int
}

func (s xSorter) Less(i, j int) bool {
	x1, _ := s.points.GetPointAt(i + s.start)
	x2, _ := s.points.GetPointAt(j + s.start)
	return x1 < x2
}

func (s xSorter) Swap(i, j int) {
	s.points.Swap(i+s.start, j+s.start)
}

func (s xSorter) Len() int {
	return s.end - s.start
}

func (s xSorter) Sort(buffer []int) {
	fmt.Println("from sroters X buffer is", buffer)
	bucketsX(s, s.bucketSize, buffer)
}

type ySorter struct {
	n                      *rNode
	points                 FlatPoints
	start, end, bucketSize int
}

func (s ySorter) Less(i, j int) bool {
	_, y1 := s.points.GetPointAt(i + s.start)
	_, y2 := s.points.GetPointAt(j + s.start)
	return y1 < y2
}

func (s ySorter) Swap(i, j int) {
	s.points.Swap(i+s.start, j+s.start)
}

func (s ySorter) Len() int {
	return s.end - s.start
}
func (s ySorter) Sort(buffer []int) {
	fmt.Println("from sroters y buffer is", buffer)

	// we already do the shifting on the sort functions
	bucketsY(s, s.bucketSize, buffer)
}
