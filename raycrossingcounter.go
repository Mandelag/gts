package gts

import (
	"github.com/twpayne/go-geom"
)

// RayCrossingCounter stores state ... TODO
type RayCrossingCounter struct {
	point            geom.Coord
	crossingCount    int
	isPointOnSegment bool
}

// NewRayCrossingCounter creates new RayCrossingCounter TODO
func NewRayCrossingCounter(point geom.Coord) *RayCrossingCounter {
	return &RayCrossingCounter{
		point: point,
	}
}

func (r *RayCrossingCounter) CountSegment(p1 geom.Coord, p2 geom.Coord) {
	// check for segment that is strictly to the left of the point
	if p1.X() < r.point.X() && p2.X() < r.point.X() {
		return
	}

	// check if the point is the vertex
	if p2.X() == r.point.X() && p2.Y() == r.point.Y() {
		r.isPointOnSegment = true
		return
	}

	// in the same horizontal line
	// if it is, it is ignored
	if p1.Y() == r.point.Y() && p2.Y() == r.point.Y() {
		minX, maxX := p1.X(), p2.X()
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		if minX <= r.point.X() && r.point.X() <= maxX {
			r.isPointOnSegment = true
		}
		return
	}

	// if in between y's
	if (p1.Y() <= r.point.Y() && r.point.Y() <= p2.Y()) || (p2.Y() <= r.point.Y() && r.point.Y() <= p1.Y()) {
		orientation := GetOrientation(p1, p2, r.point)
		if orientation == Collinear {
			r.isPointOnSegment = true
		}

		if p2.Y() < p1.Y() {
			orientation = orientation.Reverse()
		}

		if orientation == CounterClockwise {
			r.crossingCount++
		}
	}

}

func (r *RayCrossingCounter) IsOnSegment() bool {
	return r.isPointOnSegment
}

func (r *RayCrossingCounter) GetLocation() Location {
	if r.isPointOnSegment {
		return Boundary
	}
	if r.crossingCount%2 == 1 {
		return Interior
	}
	return Exterior
}
