baidumap
========

baidumap 接口调用。支持GEO双向获取，IP定位城市。

# 安装/更新

```
go get -u github.com/shajiquan/baidumap
```

# 使用

```go
package main

import (
    "fmt"
    "github.com/shajiquan/baidumap"
)

func main() {

    var lat string = "40.069462"
    var lng string = "116.346364"

    // 从坐标到地址
    GEOToAddress, err := baidumap.GetAddressViaGEO(lat, lng)

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("坐标到地址：", GEOToAddress)
    fmt.Println("坐标到地址 - 地址", GEOToAddress.Result.AddressComponent)

    // 从地址到坐标
    address := "百度大厦"
    addressToGEO, err := baidumap.GetGeoViaAddress(address)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("从地址到坐标：", addressToGEO)
    latLngOut := fmt.Sprintf("LAT：%v，Lng：%v", addressToGEO.Result.Location.Lat, addressToGEO.Result.Location.Lng)
    fmt.Println("从地址到坐标：", latLngOut)

    // 从IP到地址
    ipAddress := "202.198.16.3"
    IPToAddress, err := baidumap.GetAddressViaIP(ipAddress)

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("从IP到地址：", IPToAddress)
    fmt.Println("从IP到地址 - 地址：", IPToAddress, IPToAddress.Content.Address)

}

```