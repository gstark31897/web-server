package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

var server_file_path = "/home/octalus/file-server"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
    fmt.Println("Serving Request")
}

func main() {
    files, _ := ioutil.ReadDir(server_file_path)
    for _, f := range files {
            fmt.Println(f.Name())
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
