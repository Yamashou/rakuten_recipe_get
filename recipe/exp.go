package recipe

import "github.com/PuerkitoBio/goquery"

func exp(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find(" div >div > div > div > p.summary").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

// func main() {
// 	fmt.Print(people("http://recipe.rakuten.co.jp/recipe/1150010609/"))
// }
