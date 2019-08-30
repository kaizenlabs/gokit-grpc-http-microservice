package service

import (
	"context"
	"errors"
	"strings"

	gl "github.com/drhodes/golorem"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

type Service interface {
	// generate a word with a least min letters and at most max letters
	Lorem(ctx context.Context, requestType string, min, max int) (string, error)
}

type LoremService struct {
}

// Lorem create a new lorem construct
func (s LoremService) Lorem(ctx context.Context, requestType string, min, max int) (string, error) {
	var result string
	var err error

	if strings.EqualFold(requestType, "Word") {
		result = gl.Word(min, max)
	} else if strings.EqualFold(requestType, "Sentence") {
		result = gl.Sentence(min, max)
	} else if strings.EqualFold(requestType, "Paragraph") {
		result = gl.Paragraph(min, max)
	} else {
		err = ErrRequestTypeNotFound
	}

	return result, err
}
