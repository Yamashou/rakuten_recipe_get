package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func title(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("h1").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}
func main(){
	fmt.Println(title("http://recipe.rakuten.co.jp/recipe/1150010609/"))
}
