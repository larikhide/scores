package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var score int

	for _, v := range f.GetLeague() {
		if name == v.Name {
			score = v.Wins
			break
		}
	}

	return score
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}
