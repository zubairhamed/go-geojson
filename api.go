package geojson

type GeoJsonType string

const (
	TYPE_POINT              GeoJsonType = "Point"
	TYPE_MULTIPOINT         GeoJsonType = "MultiPoint"
	TYPE_LINESTRING         GeoJsonType = "LineString"
	TYPE_MULTILINESTRING    GeoJsonType = "MultiLineString"
	TYPE_POLYGON            GeoJsonType = "Polygon"
	TYPE_MULTIPOLYGON       GeoJsonType = "MultiPolygon"
	TYPE_GEOMETRYCOLLECTION GeoJsonType = "GeometryCollection"
	TYPE_FEATURE            GeoJsonType = "Feature"
	TYPE_FEATURECOLLECTION  GeoJsonType = "FeatureCollection"
)

type CRSObject interface {
}

type Position interface {
	GetX() float64
	GetY() float64
}

type GeoJSON interface {
	GetType() GeoJsonType
	GetCRS() CRSObject
}

type Point interface {
	GeoJSON
	GetCoordinates() Position
}

type MultiPoint interface {
	GeoJSON
	GetCoordinates() []Position
}

type LineString interface {
	GeoJSON
	GetCoordinates() []Position
}

type MultiLineString interface {
	GeoJSON
	GetCoordinates() []LineString
}

type Polygon interface {
	GeoJSON
	GetCoordinates() []LineString
}

type MultiPolygon interface {
	GeoJSON
	GetCoordinates() []Polygon
}

type GeometryCollection interface {
	GeoJSON
	GetGeometries() []GeoJSON
}

type Feature interface {
	GeoJSON
	GetId() string
	GetProperties() map[string]interface{}
	GetGeometry() GeoJSON
}

type FeatureCollection interface {
	GeoJSON
	GetFeatures() []Feature
}
