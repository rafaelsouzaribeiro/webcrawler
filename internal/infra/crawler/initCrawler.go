package crawler

import (
	"strings"
)

func (r *Usecasse) InitCrawler(url string, receive chan string) {
	reader, err := r.usecase.GetBody(url)

	defer func() {
		if reader != nil {
			(*reader).Close()
		}
	}()

	if err != nil {
		panic(err)
	}

	doc, err := r.usecase.DocumentReader(reader)

	if err != nil {
		panic(err)
	}

	go r.usecase.GetElementsByTag(doc, "a", "href", receive)

	go func() {
		for element := range receive {
			if strings.HasPrefix(element, "http://") || strings.HasPrefix(element, "https://") {
				r.usecase.Repository.InsertVisitedPage(element)
				r.InitCrawler(element, receive)
			}
		}
	}()

}
