package models

import (
	"encoding/json"
	"errors"
	// "fmt"
)

type Author struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func AuthorByName(name string) (*Author, error) {
	result, err := db.Do("hget", h_author, name)
	if err != nil {
		return nil, err
	}
	status := result[0]
	if status != "ok" {
		return nil, errors.New(status)
	}
	// fmt.Println(result)
	author := &Author{}
	err = json.Unmarshal([]byte(result[1]), author)
	if err != nil {
		return nil, err
	}

	return author, nil
}
