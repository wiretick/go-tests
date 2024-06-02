package maps

import "errors"

type Dictionary map[string]string

var (
	ErrUnknownTerm   = errors.New("Unknown search term")
	ErrAlreadyExists = errors.New("Term already exists")
)

func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]

	if !ok {
		return "", ErrUnknownTerm
	}

	return result, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrUnknownTerm: // we want this error, so continue

	// All unknown cases, should return error
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	// We are good at this point
	d[term] = definition
	return nil
}
