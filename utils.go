// Created by @menduo @ 2019-01-24
package gobaidumap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"errors"
)

// GetDefaultAK 获取默认的 ak
func GetDefaultAK() string {
	ak := defaultAppKey // baidu's
	if value, has := os.LookupEnv("GOBAIDUMAP_AK"); has {
		ak = value
	}
	return ak
}

// requestBaidu 构建 HTTP 请求
func requestBaidu(reqType, reqURL string) (interface{}, error) {

	res, err := getResStruct(reqType)
	if err != nil {
		return res, err
	}
	httpClient := http.Client{}
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

// getResStruct 处理百度 API 返回数据，映射到结构体中
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
