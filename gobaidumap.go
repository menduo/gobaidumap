package gobaidumap

// 百度地图API GEO ： http://developer.baidu.com/map/webservice-geocoding.htm
// 百度地图API IP ： http://developer.baidu.com/map/ip-location-api.htm
// 百度地图API GEO App Key 申请地址：http://lbsyun.baidu.com/apiconsole/key?application=key

// 注意，请到百度地图开发者中心申请自己的 App Key，下方的 key 是百度提供的，不保证永远有效。
// 你可以使用 NewBaiduMapClient("你的 app key") 的方式使用自己的 ak

const (
	// defaultAppKey 百度地图 defaultAppKey
	defaultAppKey string = "F454f8a5efe5e577997931cc01de3974" // baidu's

	// reqURLForGEO API URL for GetAddressViaGEO 通过 GEO 坐标信息获取地址
	reqURLForGEO string = "http://api.map.baidu.com/geocoder/v2/?ak="

	// reqURLForIP API URL for GetAddressViaIP 通过 IP 获取地址
	reqURLForIP string = "http://api.map.baidu.com/location/ip?ak="
)

// GetAddressViaGEO 通过 GEO 坐标信息获取地址
func GetAddressViaGEO(lat, lng string) (*StructGEOToAddress, error) {
	bc := NewBaiduMapClient(GetDefaultAK())
	return bc.GetAddressViaGEO(lat, lng)
}

// GetGeoViaAddress 通过地址获得 GEO 坐标
func GetGeoViaAddress(address string) (*StructAddressToGEO, error) {
	bc := NewBaiduMapClient(GetDefaultAK())
	return bc.GetGeoViaAddress(address)
}

// GetAddressViaIP 通过 IP 获取地址
func GetAddressViaIP(address string) (*StructIPToAddress, error) {
	bc := NewBaiduMapClient(GetDefaultAK())
	return bc.GetAddressViaIP(address)
}
