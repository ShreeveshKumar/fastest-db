type GeoFeature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	Time Timestamp `json:"timestamp"`
	GeoHash string `json:"geohash`
}

type Properties struct {
	Name         string  `json:"name"`
	Crop         string  `json:"crop"`
	AreaHectares float64 `json:"area_hectares"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"` 
}


type Timestamp struct{
	Time time.Time `json:"time"`
}

