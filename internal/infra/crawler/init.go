package crawler

func InitCrawler(url string, receive chan string) {

	reader, err := GetBody(url)
	defer (*reader).Close()

	if err != nil {
		panic(err)
	}

	doc, err := DocumentReader(reader)

	if err != nil {
		panic(err)
	}

	go GetElementsByTag(doc, "a", "href", receive)
}
