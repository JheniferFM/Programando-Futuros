package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "pong")
    })

    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
