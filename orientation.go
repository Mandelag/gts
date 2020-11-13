package gts

import (
	"log"

	"github.com/twpayne/go-geom"
)

// Orientation represents an orientation
type Orientation float64

// Reverse the orientation
func (o Orientation) Reverse() Orientation {
	return o * -1
}

const (
	// Clockwise orientation
	Clockwise Orientation = -1

	// CounterClockwise orientation
	CounterClockwise Orientation = 1

	// Collinear orientation
	Collinear Orientation = 0

	// Right is an alias for Clockwise orientation
	Right Orientation = Clockwise

	// Left is an alias for CounterClockwise orientatation
	Left Orientation = CounterClockwise

	// Straight is an alias for Collinear orientation
	Straight Orientation = Collinear
)

func GetOrientation(p1, p2, p3 geom.Coord) Orientation {
	det, detSum := SumDeterminant(p1.X(), p1.Y(), p2.X(), p2.Y(), p3.X(), p3.Y())

	warnBound := detSum * doublePrecissionSafeƐ
	if -det < warnBound || det < warnBound {
		// warning not precise calculation
		// we can only warn because we haven't provided with better alternative
		log.Println("float64 calculation for orientation is not precise")
	}

	if det > 0 {
		return CounterClockwise
	} else if det < 0 {
		return Clockwise
	}

	return Collinear
}

const doublePrecissionSafeƐ float64 = 1e-15

// SumDeterminant get determinant and its sum
// the sum can be used to specify detect precission threshold
// determined by b1 to b2
func SumDeterminant(ax, ay, bx, by, cx, cy float64) (float64, float64) {
	var detSum float64

	detLeft := (ax - cx) * (by - cy)
	detRight := (ay - cy) * (bx - cx)
	det := detLeft - detRight

	if detLeft > 0 {
		if detRight <= 0 {
			return det, detSum
		}
		detSum = detLeft + detRight
	} else if detLeft < 0 {
		if detRight >= 0 {
			return det, detSum
		}
		detSum = -detLeft - detRight
	}

	return det, detSum
}
