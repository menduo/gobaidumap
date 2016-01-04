package gobaidumap

// 为百度 API 返回的 JSON 创建结构体

// StructGEOToAddress GEO到地址
type StructGEOToAddress struct {
	// Status 状态码
	Status int64 `json:"status"`
	// Result
	Result *struct {
		Business string `json:"business"`
		// FormattedAddress 格式化地址
		FormattedAddress string `json:"formatted_address"`
		AddressComponent *struct {
			City     string `json:"city"`
			District string `json:"district"`
			Province string `json:"province"`
			Street   string `json:"street"`
		} `json:"addressComponent"`
		Pois []*geoToAddressPois `json:"pois"`
	} `json:"result"`
	Msg string `json:"msg"`
}

type geoToAddressPois struct {
	Addr     string `json:"addr"`
	Cp       string `json:"cp"`
	Distance string `json:"distance"`
	Name     string `json:"name"`
	PoiType  string `json:"poiType"`
	Point    []*struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"point"`
	Tel string `json:"tel"`
	UID string `json:"uid"`
	Zip string `json:"zip"`
}

// StructAddressToGEO 地址到 GEO 坐标 结构体
type StructAddressToGEO struct {
	Status int64 `json:"status"`
	Result *struct {
		Location *struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		Precise    int64  `json:"precise" `
		Confidence int64  `json:"confidence" `
		Level      string `json:"level" `
	} `json:"result"`
	Msg string `json:"msg"`
}

// StructIPToAddress IP 到 地址
type StructIPToAddress struct {
	Address string `json:"address"`
	Content *struct {
		Address       string `json:"address"`
		AddressDetail *struct {
			City         string `json:"city"`
			CityCode     int64  `json:"city_code"`
			District     string `json:"district"`
			Province     string `json:"province"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
			Point        *struct {
				X string `json:"x"`
				Y string `json:"y"`
			} `json:"point"`
		} `json:"address_detail"`
	} `json:"content"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}
