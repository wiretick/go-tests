package maps

import "errors"

type Dictionary map[string]string

var ErrUnknownTerm = errors.New("Unknown search term")

func (d Dictionary) Search(term string) (string, error) {
	result := d[term]

	if result == "" {
		return result, ErrUnknownTerm
	}

	return result, nil
}
