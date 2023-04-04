package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if ok {
		return val, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word string, def string) error {
	_, alreadyPresent := d[word]
	if alreadyPresent {
		return ErrWordExists
	}
	d[word] = def
	return nil

}
