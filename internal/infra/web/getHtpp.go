package web

import (
	"io"
	"log"
	"net/http"
)

func GetBody(url string) (*io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Fatalf("Erro ao acessar a página: %d %s", res.StatusCode, res.Status)
	}

	return &res.Body, nil
}
