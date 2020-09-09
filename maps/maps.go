package maps

// Dictionary represents a dicionary data structure
type Dictionary map[string]string

const (
	// ErrNotFound word not found
	ErrNotFound = ErrDictionary("We could not find a word you are looking for")
	// ErrWordExists word already exists
	ErrWordExists = ErrDictionary("You cannot add the word because it already exists")
	// ErrWordNonexistent word not exist
	ErrWordNonexistent = ErrDictionary("the word could not be updated because it does not exist")
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

// Update receives a new definition and update in a dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordNonexistent
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete remove word from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
