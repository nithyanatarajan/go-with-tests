package language

// revive:disable:line-length-limit
const (
	ErrNotFound           = DictionaryErr("could not find the word you are looking for")
	ErrWordAlreadyPresent = DictionaryErr("cannot add word, word already present")
	ErrWordNotPresent     = DictionaryErr("cannot update word, word not present")
)

// revive:enable:line-length-limit

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	word, found := d[key]

	if !found {
		return word, ErrNotFound
	}

	return word, nil
}

func (d Dictionary) Add(key string, definition string) error {
	if _, found := d[key]; found {
		return ErrWordAlreadyPresent
	}

	d[key] = definition
	return nil
}

func (d Dictionary) Update(key string, definition string) error {
	if _, found := d[key]; !found {
		return ErrWordNotPresent
	}

	d[key] = definition
	return nil
}

func (d Dictionary) Delete(key string) error {
	if _, found := d[key]; !found {
		return ErrNotFound
	}

	delete(d, key)
	return nil
}
