package baidumap

type StructGEOToAddress struct {
	Status int64              `json:"status"`
	Result *GEOToAddressInner `json:"result"`
}

type GEOToAddressInner struct {
	Business          string                 `json:"business"`
	Formatted_Address string                 `json:"formatted_address"`
	AddressComponent  *GEOToAddressComponent `json:"addressComponent"`
	Pois              []*GEOToAddressPois    `json:"pois"`
}

type GEOToAddressComponent struct {
	City     string `json:"city"`
	District string `json:"district"`
	Province string `json:"province"`
	Street   string `json:"street"`
}

type GEOToAddressPois struct {
	Addr     string                   `json:"addr"`
	Cp       string                   `json:"cp"`
	Distance string                   `json:"distance"`
	Name     string                   `json:"name"`
	PoiType  string                   `json:"poiType"`
	Point    []*GEOToAddressPoisPoint `json:"point"`
	Tel      string                   `json:"tel"`
	Uid      string                   `json:"uid"`
	Zip      string                   `json:"zip"`
}

type GEOToAddressPoisPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Address to GEO
type StructAddressToGEO struct {
	Status int64              `json:"status"`
	Result *AddressToGEOInner `json:"result"`
}

type AddressToGEOInner struct {
	Location   *AddressToGEOInnerGeo `json:"location"`
	Precise    int64                 `json:"precise" `
	Confidence int64                 `json:"confidence" `
	Level      string                `json:"level" `
}

type AddressToGEOInnerGeo struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type StructIPToAddress struct {
	Address string                         `json:"address"`
	Content *IPToAddressResultInnerContent `json:"content"`
	Status  int64                          `json:"status"`
}

type IPToAddressResultInnerContent struct {
	Address        string                                      `json:"address"`
	Address_Detail *IPToAddressResultInnerContentAddressDetail `json:"address_detail"`
}

type IPToAddressResultInnerContentAddressDetail struct {
	City          string                                           `json:"city"`
	City_Code     int64                                            `json:"city_code"`
	District      string                                           `json:"district"`
	Province      string                                           `json:"province"`
	Street        string                                           `json:"street"`
	Street_Number string                                           `json:"street_number"`
	Point         *IPToAddressResultInnerContentAddressDetailPoint `json:"point"`
}

type IPToAddressResultInnerContentAddressDetailPoint struct {
	X string `json:"x"`
	Y string `json:"y"`
}
