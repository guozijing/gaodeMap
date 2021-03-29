package gaodeMap

// get the route
type StructRoute struct {
	Status string `json:"status"`
	Message string `json:"info"`
	Result *StructResult `json:"route"`
}

type StructResult struct {
	Origin string `json:"orign"`
	Destination string `json:"destination"`
	Routes []*StructRoutes `json:"paths"`
}

type StructRoutes struct {
	Distance string `json:"distance"`
	Duration string `json:"duration"`
	Toll string `json:"tolls"`
	Steps []*StructSteps `json:"steps"`
}

type StructSteps struct {
	Path string `json:"polyline"`
}