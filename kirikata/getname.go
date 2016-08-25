package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func time(url string) []string {
	var i []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.f_l > h4 > img ").Each(func(_ int, s *goquery.Selection) {
		f, _ := s.Attr("alt")
		i = append(i, f)
	})
	return i
}

func main() {
	fmt.Print(time("http://park.ajinomoto.co.jp/recipe/corner/basic/vege_cutting"))
}
