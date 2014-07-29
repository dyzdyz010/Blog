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

func AllEntries() ([]Entry, error) {
	size, err := zsize(zname("all", "entry"))
	if err != nil {
		return nil, err
	}

	result, err := zscan(zname("all", "entry"), "", "", "", size)
	if err != nil {
		return nil, err
	}

	eids := make([]string, 0)
	for i := 0; i < len(result); i += 2 {
		eids = append(eids, result[i])
	}
	if len(eids) == 0 {
		return nil, nil
	}

	result, err = multi_hget(h_entry, eids)
	if err != nil {
		return nil, err
	}

	entries := []Entry{}
	for i := 1; i < len(result); i += 2 {
		entryStr := result[i]
		entry := Entry{}
		json.Unmarshal([]byte(entryStr), &entry)
		t, _ := time.Parse(time.RFC3339, entry.Date)
		entry.Date = t.Format(time.ANSIC)
		entries = append(entries, entry)
	}

	return entries, nil
}

func PublishedEntries(dir, id string) ([]Entry, bool, bool, error) {
	havePrev := false
	haveNext := false

	score_start := ""
	score_end := ""
	if dir != "" {
		score, err := zget(zname("published", "entry"), id)
		if err != nil {
			return nil, havePrev, haveNext, err
		}
		if dir == "next" {
			score_start = score
		} else {
			score_end = score
		}
	}

	result, err := zscan(zname("published", "entry"), "", score_start, score_end, page_size)

	eids := make([]string, 0)
	escores := make([]string, 0)
	for i := 0; i < len(result); i += 2 {
		eids = append(eids, result[i])
		escores = append(eids, result[i+1])
	}
	if len(eids) == 0 {
		return nil, havePrev, haveNext, nil
	}

	result, err = zscan(zname("published", "entry"), "", escores[len(eids)-1], "", page_size)
	fmt.Println(result)
	if err == nil {
		haveNext = true
	} else {
		if err.Error() != "not_found" {
			return nil, havePrev, haveNext, err
		}
	}
	_, err = zscan(zname("published", "entry"), "", "", escores[0], page_size)
	if err == nil {
		havePrev = true
	} else {
		if err.Error() != "not_found" {
			return nil, havePrev, haveNext, err
		}
	}

	result, err = multi_hget(h_entry, eids)
	if err != nil {
		return nil, havePrev, haveNext, err
	}

	entries := []Entry{}
	for i := 1; i < len(result); i += 2 {
		entryStr := result[i]
		entry := Entry{}
		json.Unmarshal([]byte(entryStr), &entry)
		t, _ := time.Parse(time.RFC3339, entry.Date)
		entry.Date = t.Format(time.ANSIC)
		entries = append(entries, entry)
	}
	return entries, havePrev, haveNext, nil
}

func EntriesByCollection(cid, dir, eid string) ([]Entry, error) {
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
		e := Entry{}
		_ = json.Unmarshal([]byte(result[i]), &e)
		t, _ := time.Parse(time.RFC3339, e.Date)
		e.Date = t.Format(time.ANSIC)
		if e.Status == "published" {
			entries = append(entries, e)
		}
	}

	return entries, nil
}

func EntryById(id string) (*Entry, error) {
	eStr, err := hget(h_entry, id)
	if err != nil {
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

	eStr, err := hget(h_entry, e.Id)
	if err != nil {
		return err
	}

	oldEntry := &Entry{}
	json.Unmarshal([]byte(eStr), oldEntry)

	ebytes, _ := json.Marshal(e)
	err = hset(h_entry, e.Id, string(ebytes))
	if err != nil {
		return err
	}

	if oldEntry.Collection != e.Collection && oldEntry.Collection != "none" {
		err = zdel(zname(oldEntry.Collection, "entry"), e.Id)
		if err != nil {
			return err
		}
	}

	if e.Collection != "none" {
		zset(zname(e.Collection, "entry"), e.Id, t.Unix())
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

	// Delete entry from All Index
	err = zdel(zname("all", "entry"), e.Id)
	if err != nil {
		return err
	}

	// Delete entry from published/draft Index
	err = zdel(zname(e.Status, "entry"), e.Id)
	if err != nil {
		return err
	}

	// Delete entry from User Index
	err = zdel(zname(e.Author, "entry"), e.Id)
	if err != nil {
		return err
	}

	// Delete entry from Collection Index
	err = zdel(zname(e.Collection, "entry"), e.Id)
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
	// e.Status = "published"

	ebytes, _ := json.Marshal(e)
	err := hset(h_entry, e.Id, string(ebytes))
	if err != nil {
		return "", err
	}

	score := t.Unix()

	// Add entry to All Index
	err = zset(zname("all", "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	// Add entry to published/draft Index
	err = zset(zname(e.Status, "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	// Add entry to User Index
	err = zset(zname(e.Author, "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	// Add entry to Collection Index
	err = zset(zname(e.Collection, "entry"), e.Id, score)
	if err != nil {
		return "", err
	}

	return e.Id, nil
}
