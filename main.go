package main

import (
    "fmt"
    "os"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func main() {
    arg := os.Args[1]
    fmt.Println("file-server-path=", arg)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
