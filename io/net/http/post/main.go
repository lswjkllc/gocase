package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("-----------------------------------------------------")
	post1()

	fmt.Println("-----------------------------------------------------")
	post2()

	fmt.Println("-----------------------------------------------------")
	post3()

	fmt.Println("-----------------------------------------------------")
	post4()
}

func post1() {
	urlValues := url.Values{}
	urlValues.Add("name", "zhaofan")
	urlValues.Add("age", "22")
	resp, _ := http.PostForm("http://httpbin.org/post", urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func post2() {
	urlValues := url.Values{
		"name": {"zhaofan"},
		"age":  {"23"},
	}
	reqBody := urlValues.Encode()
	resp, _ := http.Post("http://httpbin.org/post", "text/html", strings.NewReader(reqBody))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func post3() {
	client := &http.Client{}

	data := map[string]interface{}{
		"name": "zhaofan",
		"age":  "23",
	}
	bytesData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://httpbin.org/post", bytes.NewReader(bytesData))

	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func post4() {
	data := map[string]interface{}{
		"name": "zhaofan",
		"age":  "23",
	}
	bytesData, _ := json.Marshal(data)

	resp, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(bytesData))

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
