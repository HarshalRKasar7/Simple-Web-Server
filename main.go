package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

const (
    port = ":8080"
)

func main() {
    http.HandleFunc("/", homeHandler)
    
    fmt.Println("Server is running on port", port)
    fmt.Println("Link to index page : http://127.0.0.1:8080")
    log.Fatal(http.ListenAndServe(port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Read the contents of index.html
    htmlContent, err := ioutil.ReadFile("index.html")
    if err != nil {
        http.Error(w, "Error reading index.html: " + err.Error(), http.StatusInternalServerError)
        return
    }

    // Set Content-Type header
    w.Header().Set("Content-Type", "text/html")

    // Write HTML content
    w.Write(htmlContent)
}
