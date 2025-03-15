package log

import (
	"flag"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func CreateFileLog() {

	flag.Parse()
	var file, err1 = os.Create("crawler.log")

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
