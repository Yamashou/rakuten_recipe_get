package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type yaku struct {
	Num         string `json:"Num"`
	Name        string `json:"name"`
	Explanation string `json:"explanation"`
}

func getName(url string) []string {
	var i []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("tr > td.cook_text02 ").Each(func(_ int, s *goquery.Selection) {
		i = append(i, s.Text())
	})
	return i
}
func getEx(url string) []string {
	var i []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("table >tbody> tr > td ").Each(func(_ int, s *goquery.Selection) {
		i = append(i, s.Text())
	})
	return i
}

func main() {
	var y []yaku
	var g  []string["1","2","3","4","5"]

	name := getName("http://www.shokurepe.com/kiso/cook01.html")
	ex := getEx("http://www.shokurepe.com/kiso/cook01.html")
	for i, s := range name {
		y = append(y, yaku{
			Num:         i,
			Name:        s,
			Explanation: ex[i],
		})
	}
	bytes, _ := json.Marshal(y)
	ioutil.WriteFile("./tes.json", bytes, os.ModePerm)
}
