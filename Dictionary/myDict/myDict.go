package myDict

import "errors"

//dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

}
