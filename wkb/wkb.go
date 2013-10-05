package wkb

import (
	"encoding/binary"
)

const (
	wkbXDR = 0
	wkbNDR = 1
)

const (
	wkbPoint                = 1
	wkbLineString           = 2
	wkbPolygon              = 3
	wkbTriangle             = 17
	wkbMultiPoint           = 4
	wkbMultiLineString      = 5
	wkbMultiPolygon         = 6
	wkbGeometryCollection   = 7
	wkbPolyhedralSurface    = 15
	wkbTIN                  = 16
	wkbPointZ               = 1001
	wkbLineStringZ          = 1002
	wkbPolygonZ             = 1003
	wkbTriangleZ            = 1017
	wkbMultiPointZ          = 1004
	wkbMultiLineStringZ     = 1005
	wkbMultiPolygonZ        = 1006
	wkbGeometryCollectionZ  = 1007
	wkbPolyhedralSurfaceZ   = 1015
	wkbTINZ                 = 1016
	wkbPointM               = 2001
	wkbLineStringM          = 2002
	wkbPolygonM             = 2003
	wkbTriangleM            = 2017
	wkbMultiPointM          = 2004
	wkbMultiLineStringM     = 2005
	wkbMultiPolygonM        = 2006
	wkbGeometryCollectionM  = 2007
	wkbPolyhedralSurfaceM   = 2015
	wkbTINM                 = 2016
	wkbPointZM              = 3001
	wkbLineStringZM         = 3002
	wkbPolygonZM            = 3003
	wkbTriangleZM           = 3017
	wkbMultiPointZM         = 3004
	wkbMultiLineStringZM    = 3005
	wkbMultiPolygonZM       = 3006
	wkbGeometryCollectionZM = 3007
	wkbPolyhedralSurfaceZM  = 3015
	wkbTINZM                = 3016
)

var (
	XDR = binary.BigEndian
	NDR = binary.LittleEndian
)
