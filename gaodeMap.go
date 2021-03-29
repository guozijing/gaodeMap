package gaodeMap


const (

	defaultAppKey string = ""

	reqURLForRoute string = "https://restapi.amap.com/v3/direction/driving?"
)

func GetRoute(lat1, lng1, lat2, lng2 string) (*StructRoute, error) {
	ac := NewGaodeMapClient(GetDefaultAK())
	return ac.GetRoute(lat1, lng1, lat2, lng2)
}
