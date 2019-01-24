# gobaidumap

百度地图接口调用 golang 版。支持GEO、地址双向获取，IP获取地址。

外国 IP 什么的，百度不支持。

练习 golang 时写的，见笑啦！

感谢 [@zzdboy](https://github.com/zzdboy) 的测试的反馈！

# 安装/更新

```
go get -u github.com/menduo/gobaidumap
```

# 注意
请到百度地图开发者中心申请自己的 App Key。代码里的常量 ak 中保证永久有效，且配额经常被大家耗光（应该是大家测试着玩呢？）。

简单说，你自己去申请一个就万事大吉啦！可以通过环境变量或者 BaiduMapClient 构造方式自定义。

# 使用

```go
package main

import (
	"fmt"
	"github.com/menduo/gobaidumap"
)

func main() {

	lat := "40.069462"
	lng := "116.346364"

	bc := gobaidumap.NewBaiduMapClient(gobaidumap.GetDefaultAK())

	// 从坐标到地址
	GEOToAddress, err := bc.GetAddressViaGEO(lat, lng)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("坐标到地址：", GEOToAddress)
		fmt.Println("坐标到地址 - 地址", GEOToAddress.Result.AddressComponent)
		fmt.Println("\n")
	}

	// 从地址到坐标
	address := "百度大厦"
	addressToGEO, err := gobaidumap.GetGeoViaAddress(address)
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("从地址到坐标 - All", addressToGEO)
		fmt.Println("从地址到坐标 - Lat", addressToGEO.Result.Location.Lat)
		fmt.Println("从地址到坐标 - Lng", addressToGEO.Result.Location.Lng)
		fmt.Println("\n")
	}

	// 从IP到地址
	ipAddress := "202.198.16.3"
	IPToAddress, err := gobaidumap.GetAddressViaIP(ipAddress)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("从IP到地址：", IPToAddress)
		fmt.Println("从IP到地址 - 地址：", IPToAddress, IPToAddress.Content.Address)
		fmt.Println("\n")
	}

	// 从IP到地址
	ipAddress = "8.8.8.8"
	IPToAddress, err = gobaidumap.GetAddressViaIP(ipAddress)

	if err != nil {
		fmt.Println("从IP到地址，err !=nil：", err)
		fmt.Println("\n")
	} else {
		fmt.Println("从IP到地址：", IPToAddress)
		fmt.Println("从IP到地址 - 地址：", IPToAddress, IPToAddress.Content.Address)
		fmt.Println("\n")
	}
}


```

# CHANGELOG

## 2019-01-24

- 调整 代码结构
- 新增 `gobaidumap.NewBaiduMapClient`（创建  `BaiduMapClient` ）。
- 新增 自定义 ak 支持（环境变量，或 client 构造传参 ）
- 调整 默认 ak 的获取方式，优先从环境变量读取，没有从取默认 ak（这个 ak 的资源配额已经被大家耗光了……）


# 联系 &  反馈


shimenduo at gmail dot com

[https://github.com/menduo/gobaidumap/issues](https://github.com/menduo/gobaidumap/issues)



