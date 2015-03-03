package main

import (
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
)

var file_path string

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")

    files, _ := ioutil.ReadDir(file_path)
    for _, f := range files {
        fmt.Println(f.Name())
    }
}

func main() {
    file_path := os.Args[1]
    fmt.Println("file-server-path=", file_path)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
