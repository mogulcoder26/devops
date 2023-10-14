package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    PerformGetReq()
}

type KQuote struct {
    Quote string 
}

func PerformGetReq() {
    const url string = "https://api.kanye.rest"
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    var result KQuote
    decoder := json.NewDecoder(res.Body)
    errx := decoder.Decode(&result)
    if errx != nil {
        panic(err)
    }
    fmt.Println(result.Quote)
}