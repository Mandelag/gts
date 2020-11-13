package gts

import (
	"github.com/twpayne/go-geom"
)

// LocatePointInLinearRing checks whether a point is inside a linear ring or not.
// Ring orientation doesn't matter
// Does not check for bounding box.
//
// This algorithm based on JTS implementation:
// https://github.com/locationtech/jts/blob/master/modules/core/src/main/java/org/locationtech/jts/algorithm/RayCrossingCounter.java
//
// This class is uses (heavily coupled) to "github.com/twpayne/go-geom" module for its geometry object model.
func LocatePointInLinearRing(point *geom.Coord, ring *geom.LinearRing) Location {
	counter := NewRayCrossingCounter(*point)

	var p1 geom.Coord
	var p2 geom.Coord
	for i := 1; i < ring.NumCoords(); i++ {
		p1 = ring.Coord(i - 1)
		p2 = ring.Coord(i)
		counter.CountSegment(p1, p2)
		if counter.IsOnSegment() {
			return counter.GetLocation()
		}
	}
	return counter.GetLocation()
}
