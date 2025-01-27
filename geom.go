/*
Package geom holds geometry objects and functions to operate on them.
They can be encoded and decoded by other packages in this repository.*/
package geom

import proj "github.com/fernando-maureira/geom/proj"

// Geom is an interface for generic geometry types.
type Geom interface {
	Bounds() *Bounds
	Similar(Geom, float64) bool
	Transform(proj.Transformer) (Geom, error)

	// Len returns the total number of points in the geometry
	Len() int

	// Points returns an iterator that returns the points in the
	// geometry.
	Points() func() Point
}

// Linear is an interface for types that are linear in nature.
type Linear interface {
	Geom
	Length() float64

	// Clip returns the part of the line that falls within the given polygon.
	Clip(Polygonal) Linear

	Simplify(tolerance float64) Geom

	// Within determines whether this geometry is within the Polygonal geometry.
	// Points that lie on the edge of the polygon are considered within.
	Within(Polygonal) WithinStatus

	// Distance calculates the shortest distance to the given Point.
	Distance(Point) float64
}

// Polygonal is an interface for types that are polygonal in nature.
type Polygonal interface {
	Geom
	Polygons() []Polygon
	Intersection(Polygonal) Polygonal
	Union(Polygonal) Polygonal
	XOr(Polygonal) Polygonal
	Difference(Polygonal) Polygonal
	Area() float64
	Simplify(tolerance float64) Geom
	Centroid() Point
}

// PointLike is an interface for types that are pointlike in nature.
type PointLike interface {
	Geom
	//On(l Linear, tolerance float64) bool

	// Within determines whether this geometry is within the Polygonal geometry.
	Within(Polygonal) WithinStatus
}
