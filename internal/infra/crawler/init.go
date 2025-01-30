package crawler

import "strings"

func InitCrawler(url string, receive chan string) {

	reader, err := GetBody(url)

	defer func() {
		if reader != nil {
			(*reader).Close()
		}
	}()

	if err != nil {
		panic(err)
	}

	doc, err := DocumentReader(reader)

	if err != nil {
		panic(err)
	}

	go GetElementsByTag(doc, "a", "href", receive)

	go func() {
		for element := range receive {
			if strings.HasPrefix(element, "http://") || strings.HasPrefix(element, "https://") {
				InitCrawler(element, receive)
			}
		}
	}()

}
