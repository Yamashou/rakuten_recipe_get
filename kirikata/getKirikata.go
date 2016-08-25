package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func time(url string) []string {
	var i []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.txt01 ").Each(func(_ int, s *goquery.Selection) {
		i = append(i, s.Text())
	})
	return i
}

func main() {
	fmt.Print(time("http://park.ajinomoto.co.jp/recipe/corner/basic/vege_cutting"))
}
