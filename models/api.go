package models

import (
	"errors"
	"strconv"
)

func hsize(name string) (int, error) {
	result, err := db.Do("hsize", name)
	if err != nil {
		return 0, err
	}
	status := result[0]
	if status != "ok" {
		return 0, errors.New(status)
	}

	size, _ := strconv.Atoi(result[1])

	return size, nil
}

func hget(name, key string) (string, error) {
	result, err := db.Do("hget", name, key)
	if err != nil {
		return "", err
	}
	status := result[0]
	if status != "ok" {
		return "", errors.New(status)
	}

	return result[1], nil
}

func multi_hget(name string, keys []string) ([]string, error) {
	result, err := db.Do("multi_hget", name, keys)
	if err != nil {
		return nil, err
	}
	status := result[0]
	if status != "ok" {
		return nil, errors.New(status)
	}

	return result[1:], nil
}

func hset(name, key, value string) error {
	result, err := db.Do("hset", name, key, value)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func hdel(name, key string) error {
	result, err := db.Do("hdel", name, key)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func zsize(name string) (int, error) {
	result, err := db.Do("zsize", name)
	if err != nil {
		return 0, err
	}
	status := result[0]
	if status != "ok" {
		return 0, errors.New(status)
	}

	size, _ := strconv.Atoi(result[1])

	return size, nil
}

func zset(name, key, score string) error {
	result, err := db.Do("zset", name, key, score)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func zdel(name, key) error {
	result, err = db.Do("zdel", name, key)
	if err != nil {
		return err
	}
	status = result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func zscan(name, key_start, score_start, score_end string, limit int) ([]string, error) {
	result, err := db.Do("zscan", key_start, score_start, score_end, limit)
	if err != nil {
		return nil, err
	}
	status := result[0]
	if status != "ok" {
		return nil, errors.New(status)
	}

	return result[1:], nil
}
