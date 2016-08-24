package main

import(
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func materialQuantity(url string) []string {
	var quantity []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("img#step_image.processImage").Each(func(_ int, s *goquery.Selection) {
		t,_ := s.Attr("src")
		// fmt.Println(t)
		quantity = append(quantity, t)
	})
	return quantity
}

func main() {
	fmt.Println(materialQuantity("http://recipe.rakuten.co.jp/recipe/1150010609/"))
}
