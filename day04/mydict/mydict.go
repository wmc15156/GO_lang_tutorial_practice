package mydict

import (
	"errors"
	"fmt"
)

// type Dictionary map[string]string

var errNotFound error = errors.New("Not Found")
var errExistError error = errors.New("always has word")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	data, exists := d[word]

	if exists == false {
		return "", errNotFound
	}
	return data, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	fmt.Println(err)
	if err != nil {
		d[word] = def
		return nil
	}
	return errExistError
}

// map에서 해당키값 지우는 방법
// delete(포인터리시버, 키값)
