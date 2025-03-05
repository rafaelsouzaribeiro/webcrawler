package usecase

import "github.com/PuerkitoBio/goquery"

func (u *SqlUsecase) GetElementsByTag(doc *goquery.Document, tag string, condition string, receive chan<- string) {
	doc.Find(tag).Each(func(index int, item *goquery.Selection) {
		href, exists := item.Attr("href")
		if exists {
			receive <- href
		}
	})
}
