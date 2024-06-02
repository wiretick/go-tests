package maps

import "errors"

type Dictionary map[string]string

var ErrUnknownTerm = errors.New("Unknown search term")
var ErrAlreadyExists = errors.New("Term already exists")

func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]

	if !ok {
		return "", ErrUnknownTerm
	}

	return result, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)
	if err != ErrUnknownTerm {
		return ErrAlreadyExists
	}

	d[term] = definition
	return nil
}
