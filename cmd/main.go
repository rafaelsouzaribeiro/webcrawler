package main

import (
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/web"
)

func main() {
	reader, err := web.GetBody("https://www.google.com")
	defer (*reader).Close()

	if err != nil {
		panic(err)
	}

	doc, err := web.DocumentReader(reader)

	if err != nil {
		panic(err)
	}

	receive := make(chan string)
	go web.GetElementsByTag(doc, "a", receive)

	for element := range receive {
		println(element)
	}

}
