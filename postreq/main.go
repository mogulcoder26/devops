package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformPostReq()
}

func PerformPostReq() {
	const url = "https://xyz.dev"

	requestBody := strings.NewReader(`
	{
		"json":"it is",
		"number":2,
		"lang":"goLang",
	}`)

	res, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	content,_ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))

}

func formData () {
	
	data := make([string]string)

}
