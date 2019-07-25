package maps

//Errors
const (
	ErrNotFound  = DictionaryErr("word not found")
	ErrDuplicate = DictionaryErr("duplicate key")
)

// DictionaryErr thing
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search for key
func Search(dic map[string]string, key string) string {
	return dic[key]
}

// Dictionary custom type
type Dictionary map[string]string

// Search dic
func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return word, nil
}

// Add to dictionary
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrDuplicate
	default:
		return err

	}

	return nil
}

// Update dictionary
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

// Delete from dictionary
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
