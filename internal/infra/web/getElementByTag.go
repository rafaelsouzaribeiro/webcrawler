package web

import "github.com/PuerkitoBio/goquery"

func GetElementsByTag(doc *goquery.Document, tag string, condition string, receive chan<- string) {
	doc.Find(tag).Each(func(index int, item *goquery.Selection) {
		switch condition {
		case "text":
			text := item.Text()
			receive <- text
		case "href":
			href, exists := item.Attr("href")
			if exists {
				receive <- href
			}
		}
	})
}
