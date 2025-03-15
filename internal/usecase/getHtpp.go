package usecase

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (u *SqlUsecase) GetBody(url string) (*io.ReadCloser, error) {

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	return &res.Body, nil
}
