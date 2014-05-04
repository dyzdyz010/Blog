package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Collection struct {
	Id       string `json:"id"`
	Title    string `json:"title" form:"title"`
	Subtitle string `json:"subtitle" form:"subtitle"`
	Author   string `json:"author"`
}

func CollectionsByUser(name string) ([]Collection, error) {
	result, err := db.Do("zsize", "blog_"+name+"_collection")
	if err != nil {
		return []Collection{}, err
	}
	size, _ := strconv.Atoi(result[1])
	result, err = db.Do("zscan", "blog_"+name+"_collection", "", "", "", size)
	if err != nil {
		return []Collection{}, err
	}
	status := result[0]
	if status != "ok" {
		return []Collection{}, errors.New(status)
	}

	cids := make([]string, 0)
	for i := 1; i < len(result); i += 2 {
		cids = append(cids, result[i])
	}

	result, err = db.Do("multi_hget", h_collection, cids)
	if err != nil {
		return []Collection{}, err
	}
	status = result[0]
	if status != "ok" {
		return []Collection{}, errors.New(status)
	}

	collections := []Collection{}
	for i := 2; i < len(result); i += 2 {
		c := Collection{}
		_ = json.Unmarshal([]byte(result[i]), &c)
		collections = append(collections, c)
	}

	return collections, nil
}

func CreateCollection(c Collection) (string, error) {
	c.Id = Hash(c.Title)
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
	result, err := db.Do("hset", h_collection, c.Id, string(cbytes))
	if err != nil {
		return "", err
	}
	status := result[0]
	if status != "ok" {
		return "", errors.New(status)
	}
	result, err = db.Do("zset", "blog_"+c.Author+"_collection", c.Id, time.Now().Unix())
	return c.Id, nil
}

func UpdateCollection(c Collection) (string, error) {
	oid := c.Id
	nid := Hash(c.Title)

	_, err := db.Do("hdel", h_collection, oid)
	if err != nil {
		return "", err
	}

	c.Id = nid

	cbytes, _ := json.Marshal(c)
	result, err := db.Do("hset", h_collection, c.Id, string(cbytes))
	if err != nil {
		return "", err
	}
	status := result[0]
	if status != "ok" {
		return "", errors.New(status)
	}

	return c.Id, nil
}

func CollectionById(id string) (Collection, error) {
	result, err := db.Do("hget", h_collection, id)
	if err != nil {
		panic(err)
		return Collection{}, err
	}
	fmt.Println(result)
	status := result[0]
	if status != "ok" {
		return Collection{}, errors.New(status)
	}

	c := Collection{}
	json.Unmarshal([]byte(result[1]), &c)
	return c, nil
}

func DeleteCollection(id string) error {
	result, err := db.Do("hdel", h_collection, id)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}
