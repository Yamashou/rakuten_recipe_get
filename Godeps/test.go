package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type process struct {
	Image     string `json:"image"`
	Operation string `json:"operation"`
}

type material struct {
	Name     []string `json:"name"`
	Quantity []string `json:"quantity"`
}

//メインのjson
type mesi struct {
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	MemberNum   string    `json:"membernum"`
	Explanation string    `json:"explanation"`
	Material    material  `json:"material"`
	Process     []process `json:"process"`
}

//料理名（タイトル）を返す。
func title(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("h1").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

//何人前かを返す関数
func people(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > div > h3 > span > span").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

//材料の量をstringの配列で返す
func materialQuantity(url string) []string {
	var quantity []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > ul > li > p.amount").Each(func(_ int, s *goquery.Selection) {
		quantity = append(quantity, s.Text())
	})
	return quantity
}

//手順をstringの配列で返す
func procedure(url string) []string {
	var procedure []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("li#step_box_li.stepBox > p.stepMemo").Each(func(_ int, s *goquery.Selection) {
		procedure = append(procedure, s.Text())
	})
	return procedure
}

//完成時の写真のurlをstringで返す。
func image(url string) string {
	var imgURL string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > span > img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		// fmt.Println(url)
		imgURL = url
	})
	return imgURL
}

//調理時のコツについてをstringで返す。
func kotu(url string) string {
	var kotu string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.howtoPointBox.last > p").Each(func(_ int, s *goquery.Selection) {
		kotu = s.Text()
	})
	return kotu
}

//材料の名前をstringの配列で返す
func mat(url string) []string {
	var materials []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > ul > li > a#material_link.name").Each(func(_ int, s *goquery.Selection) {
		materials = append(materials, s.Text())
	})
	return materials
}

//説明文を返す。
func exp(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find(" div >div > div > div > p.summary").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

//各pointを１.もう一度作りたい　２.簡単だった　３.節約できた
// func point(url string) []int64 {
// 	var points []string
// 	doc, _ := goquery.NewDocument(url)
// 	doc.Find("dl > dd > div > p.num ").Each(func(_ int, s *goquery.Selection) {
// 		points = append(points, s.Text())
// 	})
// 	return points
// }

func main() {
	var recipe mesi
	var mat material
	var proc [30]process
	s := 0
	url := "http://recipe.rakuten.co.jp/recipe/1150010609/"

	recipe.Name = title(url)
	recipe.Image = image(url)
	recipe.MemberNum = people(url)
	resipe.Explanation = exp(url)

	mat.Quantity = materialQuantity(url)
	mat.Name = mat(url)
	resipe.Material = mat
	for i := range procedure(url) {
		proc[s].Operation = append(proc.Operation, i)
		s++
	}
	recipe.Process = proc
	bytes, _ := json.Marshal(recipe)
	ioutil.WriteFile("./test.json", bytes, os.ModePerm)
}
