package dictionary

type Dictionary map[string]string
type DictErr string

func (d DictErr) Error() string {
	return string(d)
}

const (
	ErrNotFound         = DictErr("could not find the word in dictionary.")
	ErrWordExists       = DictErr("word already exists in dictionary.")
	ErrWordDoesNotExist = DictErr("Can't Update, word doesn't not exist")
)

// Search the dictionary, if word not in the dictionary return error
func (d Dictionary) Search(w string) (string, error) {
	definition, ok := d[w]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add the word in the dictionary, if word exists, returns error
func (d Dictionary) Add(w, def string) error {
	_, err := d.Search(w)
	switch err {
	case ErrNotFound:
		d[w] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update the dictionary, if word doesn't exist, returns error
func (d Dictionary) Update(w, def string) error {
	_, err := d.Search(w)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[w] = def
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(w string) {
	delete(d, w)
}
