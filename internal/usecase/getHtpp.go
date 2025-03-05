package usecase

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

func (u *SqlUsecase) GetBody(url string, tlsCond bool) (*io.ReadCloser, error) {
	var res *http.Response
	var err error

	if tlsCond {
		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
		res, err = httpClient.Get(url)

	} else {
		res, err = http.Get(url)
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		log.Fatalf("Erro ao acessar a página: %d %s", res.StatusCode, res.Status)
	}

	return &res.Body, nil
}
