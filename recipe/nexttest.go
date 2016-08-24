package main

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
)
func main() {
	var test []bool
	doc, _ := goquery.NewDocument("http://recipe.rakuten.co.jp/recipe/1150010609/")

	f:= doc.Find("li#step_box_li.stepBox")
	f.Each(func(_ int, s *goquery.Selection){
		t,_:=s.Html()
		if strings.Index(t, "img") != -1{
			test = append(test,true)
			}else{
				test = append(test,false)
			}
		})
		fmt.Print(test)
}
