package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "go is neat")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":3000", nil)
}
