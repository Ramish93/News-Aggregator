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
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}


func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/wp-stat/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)

	for _, Location := range s.Locations{
		fmt.Printf("\n %s", Location)
	}

	// http.HandleFunc("/", index_handler)
	// fmt.Println("server going live on port 3000")
	// http.ListenAndServe(":3000", nil)
}
