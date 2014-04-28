package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
)

type Collection struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle"`
	Author   string   `json:"author"`
	Entries  []string `json:"entries"`
}

func CreateCollection(c Collection) error {
	hash := md5.New()
	hash.Write([]byte(c.Title))
	c.Id = hex.EncodeToString(hash.Sum(nil))

	_, err := CollectionById(c.Id)
	if err.Error() != "not_found" || err == nil {
		if err != nil {
			return err
		} else {
			return errors.New("collection exists")
		}
	}

	cbytes, _ := json.Marshal(c)
	result, err := db.Do("hset", "collection", c.Id, string(cbytes))
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}
	return nil
}

func CollectionById(id string) (*Collection, error) {
	result, err := db.Do("hget", "collection", id)
	if err != nil {
		panic(err)
		return nil, err
	}
	fmt.Println(result)
	status := result[0]
	if status != "ok" {
		return nil, errors.New(status)
	}

	c := &Collection{}
	json.Unmarshal([]byte(result[1]), c)
	return c, nil
}
