package main

import (
	"fmt"
)
import "SimpleRTree"

func main() {
	points := []float64{99, 99, 0.0, 0.0, 1.0, 1.0, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10} // array of two points 0, 0 and 1, 1

	fp := SimpleRTree.FlatPoints(points)
	r := SimpleRTree.New().Load(fp)

	closestX, closestY, distanceSquared := r.FindNearestPoint(1.0, 3.0)
	fmt.Println(closestX, closestY, distanceSquared)
}
