package gobaidumap

import (
	"testing"
)

// Test_GEOToAddress 测试 GEO To Address
func Test_GEOToAddress(t *testing.T) {
	lat := "40.069462"  //  北京市
	lng := "116.346364" //  北京市

	geoStr := lat + "," + lng
	// 从坐标到地址
	GEOToAddress, err := GetAddressViaGEO(lat, lng)

	if err != nil || GEOToAddress.Result.AddressComponent.City != "北京市" {
		t.Error("测试失败：坐标到地址测试失败，坐标：", geoStr, GEOToAddress.Result)
	} else {
		t.Log("测试通过，坐标：", geoStr, "城市:", GEOToAddress.Result.AddressComponent.City)
	}
}

// Test_GEOToAddress_2 测试 GEO to Address 2
func Test_GEOToAddress_2(t *testing.T) {
	lat := "s"
	lng := "116.346364"

	// 从坐标到地址
	GEOToAddress, err := GetAddressViaGEO(lat, lng)

	if err == nil || GEOToAddress.Result.AddressComponent.City != "" {
		t.Error("测试失败：坐标到地址测试失败，", err, GEOToAddress)
	} else {
		t.Log("测试通过")
	}
}

// 测试 IP 到地址
func Test_IPToAddress(t *testing.T) {
	ipAddress := "202.198.16.3"
	IPToAddress, err := GetAddressViaIP(ipAddress)

	if err != nil || IPToAddress.Content.AddressDetail.City != "长春市" {
		t.Error("测试失败", IPToAddress)
	} else {
		t.Log("测试通过")
	}
}

// 测试百度地图不支持国外IP
func Test_IPToAddress_2_outsea(t *testing.T) {

	ipAddress := "8.8.8.8"
	IPToAddress, err := GetAddressViaIP(ipAddress)
	if err == nil {
		t.Error("测试失败：百度地图居然支持国外 IP 了？", IPToAddress)
	} else {
		t.Log("测试通过")
	}
}

func Test_IPToAddress_3_local(t *testing.T) {
	// 测试百度地图不支持国外IP
	ipAddress := "127.0.0.1"
	IPToAddress, err := GetAddressViaIP(ipAddress)
	if err == nil {
		t.Error("测试失败", IPToAddress)
	} else {
		t.Log("测试通过")
	}
}
