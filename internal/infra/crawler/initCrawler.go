package crawler

import (
	"net/url"
	"strings"
	"sync"

	"github.com/rafaelsouzaribeiro/webcrawler/pkg/log"
	"github.com/rafaelsouzaribeiro/webcrawler/pkg/redis/producer"
)

var scheme, domain string

func (r *Usecasse) InitCrawler(urlRaw string, receive chan string) {
	r.usecase.CreateTable()

	if !strings.HasPrefix(urlRaw, "https://") && !strings.HasPrefix(urlRaw, "http://") {
		urlRaw = scheme + "://" + domain + urlRaw
	} else {
		parsedURL, err := url.Parse(urlRaw)
		if err != nil {
			log.Log.Printf("Error to parse url: %v", err)
		}

		scheme = parsedURL.Scheme
		domain = parsedURL.Host
	}

	reader, err := r.usecase.GetBody(urlRaw)

	defer func() {
		if reader != nil {
			(*reader).Close()
		}
	}()

	if err != nil {
		log.Log.Printf("Error to get body: %v", err)
		return
	}

	doc, err := r.usecase.DocumentReader(reader)

	if err != nil {
		log.Log.Printf("Error to get reader: %v", err)
		return
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
				producer.Producer(element, "crawler_broker", "localhost:6379")
				wg.Done()
			}(element)

		}
	}()
	wg.Wait()

}
