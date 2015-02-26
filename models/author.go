package models

import (
	"encoding/json"
	// "errors"
	"fmt"
)

type Author struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func AuthorByName(name string) (*Author, error) {
	result, err := hget(h_author, name)
	if err != nil {
		return nil, err
	}
	// fmt.Println(result)
	author := &Author{}
	err = json.Unmarshal([]byte(result), author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func AddAuthor(name, password string) (*Author, error) {
	author := &Author{}
	author.Name = name
	author.Password = Hash(password)
	abytes, err := json.Marshal(author)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(abytes))
	err = hset(h_author, author.Name, string(abytes))
	if err != nil {
		return nil, err
	}

	return author, nil
}
