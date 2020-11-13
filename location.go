package gts

// Location represents relationship between two geometry
// TODO streamline docs to align with JTS
type Location string

const (
	// Interior implies the geometry is inside other geometry
	Interior Location = "interior"

	// Exterior implies the geometry is outside other geometry
	Exterior Location = "exterior"

	// Boundary implies the geometry is at other geometry
	Boundary Location = "boundary"

	// None should be the default value for Location
	None Location = "none"
)
