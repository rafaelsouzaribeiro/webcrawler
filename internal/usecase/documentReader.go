package usecase

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func (u *SqlUsecase) DocumentReader(body *io.ReadCloser) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(*body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
