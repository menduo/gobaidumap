// Created by @menduo @ 2019-01-24
package gobaidumap

import (
	"fmt"
	"errors"
)

// BaiduMapClient
type BaiduMapClient struct {
	ak string
}

// NewBaiduMapClient 新建一个
func NewBaiduMapClient(ak string) *BaiduMapClient {
	return &BaiduMapClient{ak: ak}
}

// GetAk 获取 ak
func (bc *BaiduMapClient) GetAk() string {
	return bc.ak
}

// SetAk 设置 ak
func (bc *BaiduMapClient) SetAk(ak string) {
	bc.ak = ak
}

// GetAddressViaGEO 通过 GEO 坐标信息获取地址
func (bc *BaiduMapClient) GetAddressViaGEO(lat, lng string) (*StructGEOToAddress, error) {
	res := new(StructGEOToAddress)

	parameter := fmt.Sprintf("&location=%s,%s&output=json&pois=0", lat, lng)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForGEO, bc.GetAk(), parameter)

	res2, err := requestBaidu("GetAddressViaGEO", reqURL)
	if err != nil {
		return res, err
	}

	if res2.(*StructGEOToAddress).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructGEOToAddress).Msg)
		return res, errors.New(message)
	}

	res3 := res2.(*StructGEOToAddress)
	return res3, nil
}

// GetGeoViaAddress 通过地址获得 GEO 坐标
func (bc *BaiduMapClient) GetGeoViaAddress(address string) (*StructAddressToGEO, error) {
	res := new(StructAddressToGEO)

	parameter := fmt.Sprintf("&address=%s&output=json&pois=1", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForGEO, bc.GetAk(), parameter)

	res2, err := requestBaidu("GetGeoViaAddress", reqURL)
	if err != nil {
		return res, err
	}

	if res2.(*StructAddressToGEO).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructAddressToGEO).Msg)
		return res, errors.New(message)
	}

	res3 := res2.(*StructAddressToGEO)
	return res3, nil
}

// GetAddressViaIP 通过 IP 获取地址
func (bc *BaiduMapClient) GetAddressViaIP(address string) (*StructIPToAddress, error) {
	res := new(StructIPToAddress)

	parameter := fmt.Sprintf("&ip=%s&output=json&pois=0", address)
	reqURL := fmt.Sprintf("%s%s%s", reqURLForIP, bc.GetAk(), parameter)

	res2, err := requestBaidu("GetAddressViaIP", reqURL)
	if err != nil {
		return res, err
	}

	if res2.(*StructIPToAddress).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructIPToAddress).Message)
		return res, errors.New(message)
	}

	res3 := res2.(*StructIPToAddress)
	return res3, nil

}
