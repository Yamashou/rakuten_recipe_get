package recipe

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func image(url string) string {
	var imgURL string
	doc, _ := goquery.NewDocument(url)
	doc.Find("li#step_box_li.stepBox > div.stepPhoto > span > img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		fmt.Println(url)
		imgURL = url
	})
	return imgURL
}

// func main() {
// 	url := "http://recipe.rakuten.co.jp/recipe/1280001430/"
// 	fmt.Println(image(url))
// }
