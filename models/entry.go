package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Entry struct {
	Id         string `json:"id"`
	Title      string `json:"title" form:"title"`
	Subtitle   string `json:"subtitle" form:"subtitle"`
	Author     string `json:"author"`
	Date       string `json:"date"`
	Collection string `json:"collection"`
	Content    string `json:"content" form:"content"`
	Likes      int    `json:"likes"`
	Status     string `json:"status" form:"status"`
}

func PublishedEntries() (entries []Entry) {
	entries = []Entry{}

	result, err := db.Do("hscan", "entry", "", "", 10)
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
	fmt.Println(id)
	result, err := db.Do("hget", "entry", id)
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
