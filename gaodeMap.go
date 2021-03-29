package gaodeMap


const (

	defaultAppKey string = ""

	reqURLForGEO string = "https://restapi.amap.com/v3/geocode/regeo?radius=1000&language=en&key="

	reqURLForRoute string = "https://restapi.amap.com/v3/direction/driving?"
)

func GetAddressViaGEO(lat, lng string) (*StructGEOToAddress, error) {
	ac := NewGaodeMapClient(GetDefaultAK())
	return ac.GetAddressViaGEO(lat, lng)
}

func GetRoute(lat1, lng1, lat2, lng2 string) (*StructRoute, error) {
	ac := NewGaodeMapClient(GetDefaultAK())
	return ac.GetRoute(lat1, lng1, lat2, lng2)
}
