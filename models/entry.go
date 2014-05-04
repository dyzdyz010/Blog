package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Entry struct {
	Id         string `json:"id"`
	Title      string `json:"title" form:"title"`
	Subtitle   string `json:"subtitle" form:"subtitle"`
	Author     string `json:"author"`
	Date       string `json:"date"`
	Collection string `json:"collection form:"collection"`
	Content    string `json:"content" form:"content"`
	Likes      int    `json:"likes"`
	Status     string `json:"status" form:"status"`
}

func PublishedEntries() (entries []Entry) {
	entries = []Entry{}

	result, err := db.Do("hscan", h_entry, "", "", 10)
	if err != nil {
		panic(err)
		return
	}

	status := result[0]
	if status != "ok" {
		return
	} else {
		for i := 2; i < len(result); i += 2 {
			entryStr := result[i]
			entry := Entry{}
			json.Unmarshal([]byte(entryStr), &entry)
			t, _ := time.Parse(time.RFC3339, entry.Date)
			entry.Date = t.Format(time.ANSIC)
			entries = append(entries, entry)
		}
	}
	return entries
}

func EntryById(id string) *Entry {
	// fmt.Println(id)
	result, err := db.Do("hget", h_entry, id)
	if err != nil {
		panic(err)
		return nil
	}
	status := result[0]
	if status != "ok" {
		return nil
	}

	entry := &Entry{}
	json.Unmarshal([]byte(result[1]), entry)
	t, _ := time.Parse(time.RFC3339, entry.Date)
	entry.Date = t.Format(time.ANSIC)

	return entry
}

func UpdateEntry(e Entry) (string, error) {
	oid := e.Id
	nid := Hash(e.Title)

	_, err := db.Do("hdel", h_entry, oid)
	if err != nil {
		return "", err
	}

	e.Id = nid
	t := time.Now()
	e.Date = t.Format(time.RFC3339)

	ebytes, _ := json.Marshal(e)
	result, err := db.Do("hset", h_entry, e.Id, string(ebytes))
	if err != nil {
		return "", err
	}
	status := result[0]
	if status != "ok" {
		return "", errors.New(status)
	}

	return e.Id, nil
}

func DeleteEntry(id string) error {
	result, err := db.Do("hdel", h_entry, id)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func PostNewEntry(e Entry) (string, error) {
	fmt.Println(e)
	e.Id = Hash(e.Title)
	t := time.Now()
	e.Date = t.Format(time.RFC3339)
	e.Likes = 0
	e.Status = "published"

	ebytes, _ := json.Marshal(e)
	result, err := db.Do("hset", h_entry, e.Id, string(ebytes))
	if err != nil {
		return "", err
	}
	status := result[0]
	if status != "ok" {
		return "", errors.New(status)
	}

	return e.Id, nil
}
