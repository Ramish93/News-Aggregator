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
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}


func main() {
	var s SitemapIndex
	var n News
	news_map := make(map[string] NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/wp-stat/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations{
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		fmt.Println(n)
		for idx,_ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		
		}
	}

}

// http.HandleFunc("/", index_handler)
// fmt.Println("server going live on port 3000")
// http.ListenAndServe(":3000", nil)