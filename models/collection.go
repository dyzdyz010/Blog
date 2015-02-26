package models

import (
	"encoding/json"
	"errors"
	"fmt"
	// "strconv"
	"time"
)

type Collection struct {
	Id       string `json:"id"`
	Title    string `json:"title" form:"title"`
	Subtitle string `json:"subtitle" form:"subtitle"`
	Author   string `json:"author"`
	Date     string `json:"date"`
}

func AllCollections() ([]Collection, error) {
	size, err := hsize(h_collection)
	if err != nil {
		return []Collection{}, err
	}

	result, err := hscan(h_collection, "", "", size)
	if err != nil {
		return []Collection{}, err
	}

	collections := []Collection{}
	for i := 1; i < len(result); i += 2 {
		c := Collection{}
		_ = json.Unmarshal([]byte(result[i]), &c)
		t, _ := time.Parse(time.RFC3339, c.Date)
		c.Date = t.Format(time.ANSIC)
		collections = append(collections, c)
	}

	return collections, nil
}

func CollectionsByUser(name string) ([]Collection, error) {
	zsetName := zname(name, "collection")
	size, err := zsize(zsetName)
	if err != nil {
		return []Collection{}, err
	}
	if size == 0 {
		return []Collection{}, err
	}

	result, err := zscan(zsetName, "", "", "", size)
	if err != nil {
		return []Collection{}, err
	}

	cids := make([]string, 0)
	for i := 0; i < len(result); i += 2 {
		cids = append(cids, result[i])
	}

	result, err = multi_hget(h_collection, cids)
	if err != nil {
		return []Collection{}, err
	}

	collections := []Collection{}
	for i := len(result) - 1; i > 0; i -= 2 {
		c := Collection{}
		_ = json.Unmarshal([]byte(result[i]), &c)
		t, _ := time.Parse(time.RFC3339, c.Date)
		c.Date = t.Format(time.ANSIC)
		collections = append(collections, c)
	}

	return collections, nil
}

func CreateCollection(c Collection) (string, error) {
	c.Id = Hash(c.Title)
	c.Date = time.Now().Format(time.RFC3339)
	fmt.Println("Create collection: ", c)

	_, err := CollectionById(c.Id)
	fmt.Println("Check duplication: ", err)
	if err.Error() != "not_found" || err == nil {
		if err != nil {
			return "", err
		} else {
			return "", errors.New("collection exists")
		}
	}

	cbytes, _ := json.Marshal(c)
	err = hset(h_collection, c.Id, string(cbytes))
	if err != nil {
		return "", err
	}
	err = zset(zname(c.Author, "collection"), c.Id, time.Now().Unix())
	if err != nil {
		return "", err
	}

	return c.Id, nil
}

func UpdateCollection(c Collection) (string, error) {
	cbytes, _ := json.Marshal(c)

	err := hset(h_collection, c.Id, string(cbytes))
	if err != nil {
		return "", err
	}

	return c.Id, nil
}

func CollectionById(id string) (Collection, error) {
	result, err := hget(h_collection, id)
	if err != nil {
		return Collection{}, err
	}

	c := Collection{}
	json.Unmarshal([]byte(result), &c)
	return c, nil
}

func DeleteCollection(id string) error {
	err := hdel(h_collection, id)
	if err != nil {
		return err
	}

	return nil
}
