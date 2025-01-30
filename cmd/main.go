package main

import "github.com/rafaelsouzaribeiro/webcrawler/internal/infra/crawler"

func main() {
	receive := make(chan string)

	crawler.InitCrawler("https://www.google.com", receive)

	for element := range receive {
		println(element)
	}

}
