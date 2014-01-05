package baidumap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 百度地图API GEO ： http://developer.baidu.com/map/webservice-geocoding.htm
// 百度地图API IP ： http://developer.baidu.com/map/ip-location-api.htm
// 百度地图API GEO App Key 申请地址：http://lbsyun.baidu.com/apiconsole/key?application=key

// 注意，请到百度地图开发者中心申请自己的 App Key，下方的 key 是百度提供的，不保证永远有效。

const (
	AppKey       string = "F454f8a5efe5e577997931cc01de3974" // baidu's
	reqURLForGEO string = "http://api.map.baidu.com/geocoder/v2/?ak="
	reqURLForIP  string = "http://api.map.baidu.com/location/ip?ak="
)

var (
	httpClient = http.Client{}
)

func GetAddressViaGEO(lat, lng string) (*StructGEOToAddress, error) {
	res := new(StructGEOToAddress)

	parameter := fmt.Sprintf("&location=%s,%s&output=json&pois=0", lat, lng)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForGEO, AppKey, parameter)

	res2, err := requestBaidu("GetAddressViaGEO", reqURL)
	if err != nil {
		return res, err
	}
	res3 := res2.(*StructGEOToAddress)
	return res3, nil
}

func GetGeoViaAddress(address string) (*StructAddressToGEO, error) {
	res := new(StructAddressToGEO)

	parameter := fmt.Sprintf("&address=%s&output=json&pois=1", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForGEO, AppKey, parameter)

	res2, err := requestBaidu("GetGeoViaAddress", reqURL)
	if err != nil {
		return res, err
	}
	res3 := res2.(*StructAddressToGEO)
	return res3, nil
}

func GetAddressViaIP(address string) (*StructIPToAddress, error) {
	res := new(StructIPToAddress)

	parameter := fmt.Sprintf("&ip=%s&output=json&pois=0", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForIP, AppKey, parameter)

	res2, err := requestBaidu("GetAddressViaIP", reqURL)
	if err != nil {
		return res, err
	}
	res3 := res2.(*StructIPToAddress)
	return res3, nil

}

func requestBaidu(reqType, reqURL string) (interface{}, error) {
	res, err := getResStruct(reqType)
	if err != nil {
		return res, err
	}

	resp, err := httpClient.Get(reqURL)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {

		err := json.Unmarshal(bytes, &res)

		if err != nil {
			return res, err
		}
	} else {
		return res, errors.New("请求百度API失败，状态码不等于200")
	}

	return res, nil
}

func getResStruct(reqType string) (interface{}, error) {
	var res interface{}

	if reqType == "GetAddressViaIP" {
		return new(StructIPToAddress), nil
	}

	if reqType == "GetGeoViaAddress" {
		return new(StructAddressToGEO), nil
	}

	if reqType == "GetAddressViaGEO" {
		return new(StructGEOToAddress), nil
	}
	return res, errors.New("结构体请求错误")
}
