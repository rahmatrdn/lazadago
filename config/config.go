package config

import (
	"errors"
	"fmt"
	"strings"
	"sort"

	"github.com/wjpxxx/letgo/encry"
	"github.com/wjpxxx/letgo/httpclient"
	"github.com/wjpxxx/letgo/lib"
)

//baseURL
var baseURL map[string]string=map[string]string{
	"th":"https://api.lazada.co.th/rest",
	"ph":"https://api.lazada.com.ph/rest",
	"sg":"https://api.lazada.sg/rest",
	"id":"https://api.lazada.co.id/rest",
	"my":"https://api.lazada.com.my/rest",
	"vn":"https://api.lazada.vn/rest",
	"all":"https://auth.lazada.com/oauth/authorize",
}
//Config
type Config struct{
	AppKey string `json:"app_key"`
	AccessToken string `json:"access_token"`
	AppSecret string `json:"app_secret"`
	Country string `json:"country"`
}

//String
func (c *Config) String() string {
	return lib.ObjectToString(c)
}
//GetApiUrl
func (c *Config)GetApiUrl(country string)string{
	if country!=""{
		return baseURL[country]
	}
	return baseURL[c.Country]
}


//GetCommonParam
func (c *Config) GetCommonParam() lib.InRow {
	ti := lib.TimeLong()
	param := lib.InRow{
		"app_key": c.AppKey,
		"timestamp":  ti,
		"access_token":c.AccessToken,
		"sign_method":"sha256",
	}
	return param
}

//HttpGet
func (c *Config) HttpGet(method string, data interface{}, out interface{}) error {
	return c.Http("GET", method, data, out)
}

//HttpPost
func (c *Config) HttpPost(method string, data interface{}, out interface{}) error {
	return c.Http("POST", method, data, out)
}

//HttpPostFile
func (c *Config) HttpPostFile(method string, data interface{}, out interface{}) error {
	return c.Http("POSTFILE", method, data, out)
}

//Http 请求
func (c *Config) Http(requestMethod, method string, data interface{}, out interface{}) error {
	param := c.GetCommonParam()
	inputParam:=data.(lib.InRow)
	allParam:=lib.MergeInRow(param,inputParam)
	allParam["sign"]=Sign(c.AppSecret,method,allParam)
	fullURL := fmt.Sprintf("%s?%s", c.GetApiUrl(""), method)
	ihttp := httpclient.New().WithTimeOut(120)
	var result *httpclient.HttpResponse
	if requestMethod == "GET" {
		result = ihttp.Get(fullURL, allParam)
	} else {
		result = ihttp.Post(fullURL, allParam)
	}
	//fmt.Println(result.Dump)
	if result.Err != "" {
		return errors.New(result.Err)
	}
	if result.Code != 200 {
		return errors.New("请求失败")
	}
	lib.StringToObject(result.Body(), out)
	return nil
}

//Sign
func Sign(appSecret,method string,param lib.InRow) string {
	tmpParam:=lib.InRow{}
	//排除掉文件上传
	for k,v:=range param{
		if k[0] != '@' {
			tmpParam[k]=v
		}
	}
	query := ""
	sortedKeys := make([]string, 0)
	for k, _ := range tmpParam {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	for _, k := range sortedKeys {
		if vs, ok := tmpParam[k].(string); ok {
			query +=fmt.Sprintf("%s%s",k,vs)
		}else if vs, ok := tmpParam[k].(int); ok {
			query +=fmt.Sprintf("%s%d",k,vs)
		}else if vs, ok := tmpParam[k].(int64); ok {
			query +=fmt.Sprintf("%s%d",k,vs)
		}else if vs, ok := tmpParam[k].(int32); ok {
			query +=fmt.Sprintf("%s%d",k,vs)
		}else if vs, ok := tmpParam[k].(float32); ok {
			query +=fmt.Sprintf("%s%f",k,vs)
		}else if vs, ok := tmpParam[k].(float64); ok {
			query +=fmt.Sprintf("%s%f",k,vs)
		}else if vs, ok := tmpParam[k].(bool); ok {
			query +=fmt.Sprintf("%s%t",k,vs)
		}
		
	}
	//fmt.Println(method+query)
	return strings.ToUpper(encry.HmacHex(method+query, appSecret))
}
