package recipe

import(
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func materialQuantity(url string) []string {
	var quantity []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > ul > li > p.amount").Each(func(_ int, s *goquery.Selection) {
		quantity = append(quantity, s.Text())
	})
	return quantity
}

func main() {
	fmt.Println(materialQuantity("http://recipe.rakuten.co.jp/recipe/1280001430/"))
}
