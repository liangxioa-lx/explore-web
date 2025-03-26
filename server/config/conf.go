package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Conf *Config

type (
	Config struct {
		DataBase   DataBase   `json:"database"`
		Server     Server     `json:"server"`
		FileUpload FileUpload `json:"fileUpload"`
		OSS        OSS        `json:"oss"`
		JWT        JWT        `json:"jwt"`
	}
	DataBase struct {
		Dsn         string `json:"dsn"`
		Userame     string `json:"username"`
		Password    string `json:"password"`
		MaxConnIdle int    `json:"maxConnIdle"`
	}
	Server struct {
		Address string `json:"address"`
		Port    string `json:"port"`
	}
	FileUpload struct {
		Directory string `json:"dir"`
	}
	OSS struct {
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		BucketName      string `json:"bucketName"`
		Endpoint        string `json:"endpoint"`
		ExpireTime      int64  `json:"expireTime"`
	}
	JWT struct {
		Secret string `json:"secret"`
	}
)

/*
支持热加载
减少有损服务发布
*/
func init() {
	parse()
}

func parse() {
	ft, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("config parse error")
		panic(err)
	}
	err = json.Unmarshal(ft, &Conf)
	if err != nil {
		fmt.Printf("config parse error")
		panic(err)
	}
}

// 热加载
func HotLoad() {
	parse()
	//for {
	//	parse()
	//	time.Sleep(10 * time.Minute)
	//}
}
