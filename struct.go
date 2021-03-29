package gaodeMap

// get the route
type StructRoute struct {
	Status int64 `json:"status"`
	Message string `json:"info"`
	Result *StructResult `json:"route"`
}

type StructResult struct {
	Origin string `json:"orign"`
	Destination string `json:"destination"`
	Routes []*StructRoutes `json:"paths"`
}

type StructRoutes struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Toll int64 `json:"tolls"`
		Steps []*StructSteps `json:"steps"`
}

type StructSteps struct {
	Path string `json:"polyline"`
}

// StructGEOToAddress
type StructGEOToAddress struct {
	Status int64 `json:"status"`
	Regeocode struct {
		AddressComponent struct {
			Country  string `json:"country"`
		} `json:"addressComponent"`
	} `json:"regeocode"`
}
