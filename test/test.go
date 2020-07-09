package main

import (
	"fmt"
)
import "SimpleRTree"

func main() {
	points := []float64{99, 5, 0.0, 0.0, 1.0, 1.0, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 1.2, 1.2, 10, 10, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10} // array of two points 0, 0 and 1, 1
	fmt.Println(len(points))                                                                                                                                                                     //lines := []float64{0,0,12,12}

	fp := SimpleRTree.FlatPoints(points)
	r := SimpleRTree.New().Load(fp)
	//r := SimpleRTree.Reload
	closestX, closestY, distanceSquared := r.FindNearestPoint(1.0, 3.0)
	fmt.Println(closestX, closestY, distanceSquared)
}
