package maps

// Dictionary represents a dicionary data structure
type Dictionary map[string]string

const (
	// ErrNotFound word not found
	ErrNotFound = ErrDictionary("We could not find a word you are looking for")
	// ErrWordExists word already exists
	ErrWordExists = ErrDictionary("You cannot add the word because it already exists")
)

// ErrDictionary represents a dictionary error
type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

// Search returns a word in an dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]
	if !exists {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add add a new word in a dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
