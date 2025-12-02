package main

import (
	"encoding/json"
)

type repoRecord struct {
	Id       int
	Language string
	Stars    int
	Watchers int
}

func (r repoRecord) Stringify() (string, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

type point struct {
	Language        string
	Stars           int
	Watchers        int
	UniqueRepoCount int
	Year            int
	Quarter         int
}
