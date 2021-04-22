package main

import (
    "github.com/micro/go-micro/v2/web"
    "net/http"
    "log"
)

func main() {
    service := web.NewService(
        web.Address(":8080"), // http 端口
    )
    service.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello Go-Micro world"))
    })

    if err := service.Run(); err != nil {
        log.Println(err.Error())
    }
}

