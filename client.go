package gaodeMap

import (
	"fmt"
	"errors"
)

type GaodeMapClient struct {
	ak string
}

func NewGaodeMapClient(ak string) *GaodeMapClient {
	return &GaodeMapClient{ak: ak}
}

func (ac *GaodeMapClient) GetAk() string {
	return ac.ak
}

func (ac *GaodeMapClient) SetAk(ak string) {
	ac.ak = ak
}

func (ac *GaodeMapClient) GetRoute(lat1, lng1, lat2, lng2 string) (*StructRoute, error) {
	res := new(StructRoute)
	parameter := fmt.Sprintf("origin=%s,%s&destination=%s,%s&key=%s", lng1, lat1, lng2, lat2, ac.GetAk())
	reqURL := fmt.Sprintf("%s%s", reqURLForRoute, parameter)

	res2, err := requestGaode("GetRoute", reqURL)
	if err != nil {
		return res, err
	}
	if res2.(*StructRoute).Status != "1" {
		message := fmt.Sprintf("error：%s", res2.(*StructRoute).Message)
		return res, errors.New(message)
	}
	res3 := res2.(*StructRoute)
	return res3, nil
}

func (ac *GaodeMapClient) GetAddressViaGEO(lat, lng string) (*StructGEOToAddress, error) {
	res := new(StructGEOToAddress)

	parameter := fmt.Sprintf("%s&output=JSON&location=%s,%s", ac.GetAk(), lng, lat)
	reqURL := fmt.Sprintf("%s%s", reqURLForGEO, parameter)

	res2, err := requestGaode("GetAddressViaGEO", reqURL)
	if err != nil {
		return res, err
	}
	if res2.(*StructGEOToAddress).Status != "1" {
		message := fmt.Sprintf("require failed")
		return res, errors.New(message)
	}
	res3 := res2.(*StructGEOToAddress)
	return res3, nil
}
