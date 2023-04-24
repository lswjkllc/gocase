package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("-----------------------------------------------------")
	get1()

	fmt.Println("-----------------------------------------------------")
	get2()

	fmt.Println("-----------------------------------------------------")
	get3()

	fmt.Println("-----------------------------------------------------")
	get4()

	fmt.Println("-----------------------------------------------------")
	get5()
}

// 基本的 GET 请求
func get1() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
}

// 带参数的 GET 请求
func get2() {
	resp, err := http.Get("http://httpbin.org/get?name=zhaofan&age=23")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 把一些参数做成变量而不是直接放到 URL 中的 GET 请求
func get3() {
	params := url.Values{}
	uri, err := url.Parse("http://httpbin.org/get")
	if err != nil {
		return
	}
	params.Set("name", "zhaofan")
	params.Set("age", "23")
	//如果参数中有中文参数,这个方法会进行URLEncode
	uri.RawQuery = params.Encode()
	urlPath := uri.String()
	fmt.Println(urlPath) // https://httpbin.org/get?age=23&name=zhaofan

	resp, _ := http.Get(urlPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 解析 JSON 类型的返回结果的 GET 请求
func get4() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var res Result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)
}

type Result struct {
	Args    string            `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

// 添加请求头的 GET 请求
func get5() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	req.Header.Add("name", "zhaofan")
	req.Header.Add("age", "3")

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
