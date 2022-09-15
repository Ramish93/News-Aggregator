package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "go is neat")
}

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}


func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/wp-stat/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

	// http.HandleFunc("/", index_handler)
	// fmt.Println("server going live on port 3000")
	// http.ListenAndServe(":3000", nil)
}
