package maps

const ( // Makes sure the message does not get changed later
	ErrUnknownTerm   = DictionaryErr("Unknown search term")
	ErrAlreadyExists = DictionaryErr("Term already exists")
)

type DictionaryErr string

// This makes our custom error compatible with the error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

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
