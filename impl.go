package geojson

type BaseGeoJson struct {
	crs CRSObject
}

func (g BaseGeoJson) GetCRS() CRSObject {
	return g.crs
}

type GeoJsonPosition struct {
	BaseGeoJson
	x float64
	y float64
}

func (g GeoJsonPosition) GetX() float64 {
	return g.x
}

func (g GeoJsonPosition) GetY() float64 {
	return g.y
}

type GeoJsonPoint struct {
	BaseGeoJson
	coordinates Position
}

func (g GeoJsonPoint) GetType() GeoJsonType {
	return TYPE_POINT
}

func (g GeoJsonPoint) GetCoordinates() Position {
	return g.coordinates
}

type GeoJsonMultiPoint struct {
	GeoJSON
	coordinates []Position
}

func (g GeoJsonMultiPoint) GetType() GeoJsonType {
	return TYPE_MULTIPOINT
}

func (g GeoJsonMultiPoint) GetCoordinates() []Position {
	return g.coordinates
}

type GeoJsonLineString struct {
	BaseGeoJson
	coordinates []Position
}

func (g GeoJsonLineString) GetCoordinates() []Position {
	return g.coordinates
}

type GeoJsonMultiLineString struct {
	BaseGeoJson
	coordinates []LineString
}

func (g GeoJsonMultiLineString) GetCoordinates() []LineString {
	return g.coordinates
}

type GeoJsonPolygon struct {
	BaseGeoJson
	coordinates []LineString
}

func (g GeoJsonPolygon) GetCoordinates() []LineString {
	return g.coordinates
}

type GeoJsonMultiPolygon struct {
	BaseGeoJson
	coordinates []Polygon
}

func (g GeoJsonMultiPolygon) GetCoordinates() []Polygon {
	return g.coordinates
}

type GeoJsonGeometryCollection struct {
	BaseGeoJson
	geometries []GeoJSON
}

func (g GeoJsonGeometryCollection) GetGeometries() []GeoJSON {
	return g.geometries
}

type GeoJsonFeature struct {
	BaseGeoJson
	id         string
	properties map[string]interface{}
	geometry   GeoJSON
}

func (g GeoJsonFeature) GetId() string {
	return g.id
}

func (g GeoJsonFeature) GetProperties() map[string]interface{} {
	return g.properties
}

func (g GeoJsonFeature) GetGeometry() GeoJSON {
	return g.geometry
}

type GeoJsonFeatureCollection struct {
	BaseGeoJson
	features []Feature
}

func (g GeoJsonFeatureCollection) GetFeatures() []Feature {
	return g.features
}
