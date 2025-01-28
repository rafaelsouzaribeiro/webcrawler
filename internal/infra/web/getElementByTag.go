package web

import "github.com/PuerkitoBio/goquery"

func GetElementsByTag(doc *goquery.Document, tag string, receive chan<- string) {
	doc.Find(tag).Each(func(index int, item *goquery.Selection) {
		text := item.Text()
		receive <- text
	})
}
