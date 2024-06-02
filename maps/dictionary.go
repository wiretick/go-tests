package maps

import "errors"

type Dictionary map[string]string

var ErrUnknownTerm = errors.New("Unknown search term")

func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]

	if !ok {
		return "", ErrUnknownTerm
	}

	return result, nil
}

func (d Dictionary) Add(term, definition string) error {
	d[term] = definition
	return nil
}
