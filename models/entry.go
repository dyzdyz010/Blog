package models

import (
	"encoding/json"
	// "errors"
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

func EntriesByCollection(cid string) ([]Entry, error) {
	c, err := CollectionById(cid)
	if err != nil {
		panic(err)
	}

	size, err := zsize(zname(c.Title, "entry"))
	if err != nil {
		return nil, err
	}

	result, err := zscan(zname(c.Title, "entry"), "", "", "", size)
	if err != nil {
		panic(err)
		return nil, err
	}

	eids := make([]string, 0)
	for i := 0; i < len(result); i += 2 {
		eids = append(eids, result[i])
	}

	result, err = multi_hget(h_entry, eids)
	if err != nil {
		return nil, err
	}

	entries := []Entry{}
	for i := len(result) - 1; i > 0; i -= 2 {
		c := Entry{}
		_ = json.Unmarshal([]byte(result[i]), &c)
		t, _ := time.Parse(time.RFC3339, c.Date)
		c.Date = t.Format(time.ANSIC)
		entries = append(entries, c)
	}

	return entries, nil
}

func EntryById(id string) (*Entry, error) {
	eStr, err := hget(h_entry, id)
	if err != nil {
		panic(err)
		return nil, err
	}

	entry := &Entry{}
	json.Unmarshal([]byte(eStr), entry)
	t, _ := time.Parse(time.RFC3339, entry.Date)
	entry.Date = t.Format(time.ANSIC)

	return entry, nil
}

func UpdateEntry(e Entry) error {
	fmt.Println(e)
	t := time.Now()
	e.Date = t.Format(time.RFC3339)

	ebytes, _ := json.Marshal(e)
	err := hset(h_entry, e.Id, string(ebytes))
	if err != nil {
		return err
	}
	return nil
}

func DeleteEntry(id string) error {
	e, err := EntryById(id)
	if err != nil {
		return err
	}

	err = hdel(h_entry, id)
	if err != nil {
		return err
	}

	err = zdel(zname(e.Collection, "entry"), id)
	if err != nil {
		return err
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
	// result, err := db.Do("hset", h_entry, e.Id, string(ebytes))
	// if err != nil {
	// 	return "", err
	// }
	// status := result[0]
	// if status != "ok" {
	// 	return "", errors.New(status)
	// }
	err := hset(h_entry, e.Id, string(ebytes))
	if err != nil {
		return "", err
	}

	// result, err = db.Do("zset", "blog_"+e.Author+"_entry", e.Id, time.Now().Unix())
	// if err != nil {
	// 	return "", err
	// }
	// status = result[0]
	// if status != "ok" {
	// 	return "", errors.New(status)
	// }
	score := time.Now().Unix()
	err = zset(zname(e.Author, "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	// result, err = db.Do("zset", "blog_"+e.Collection+"_entry", e.Id, time.Now().Unix())
	// if err != nil {
	// 	return "", err
	// }
	// status = result[0]
	// if status != "ok" {
	// 	return "", errors.New(status)
	// }
	err = zset(zname(e.Collection, "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	return e.Id, nil
}
