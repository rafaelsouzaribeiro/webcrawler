package crawler

import (
	"strings"
)

func (r *Usecasse) InitCrawler(url string, receive chan string) {
	r.usecase.CreateTable()
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
				if visited, _ := r.usecase.IsPageVisited(element); visited {
					continue
				}
				r.usecase.InsertVisitedPage(element)
				r.InitCrawler(element, receive)
			}
		}
	}()

}
