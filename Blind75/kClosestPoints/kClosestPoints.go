package kClosestPoints

import (
	"sort"
)

type Distance struct {
	distance int
	pos      int
}
type Distances []Distance

func KClosest(points [][]int, K int) [][]int {
	if len(points) == K {
		return points
	}

	var distances Distances
	for i := 0; i < len(points); i++ {
		dist := Distance{
			distance: points[i][0]*points[i][0] + points[i][1]*points[i][1],
			pos:      i,
		}
		distances = append(distances, dist)
	}
	sort.Sort(distances)

	var re [][]int
	for i := 0; i < K; i++ {
		re = append(re, points[distances[i].pos])
	}
	return re
}

func (d Distances) Len() int {
	return len(d)
}
func (d Distances) Less(i, j int) bool {
	return d[i].distance < d[j].distance
}
func (d Distances) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
