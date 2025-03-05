package crawler

import (
	"strings"
	"sync"

	"github.com/rafaelsouzaribeiro/webcrawler/pkg/redis/producer"
)

func (r *Usecasse) InitCrawler(url string, receive chan string) {
	r.usecase.CreateTable()
	var cond bool = false
	if strings.HasPrefix(url, "https://") {
		cond = true
	}
	reader, err := r.usecase.GetBody(url, cond)

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
	var wg sync.WaitGroup

	go func() {
		for element := range receive {
			if visited, _ := r.usecase.IsPageVisited(element); visited {
				continue
			}
			wg.Add(1)
			go func(element string) {
				r.usecase.InsertVisitedPage(element)
				producer.Producer(element, "crawler", "localhost:6379")
				wg.Done()
			}(element)

		}
	}()
	wg.Wait()

}
