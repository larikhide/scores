package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("cannot parse league: %v", err)
	}
	return league, err
}

func (l League) FindPlayer(name string) *Player {
	for i, p := range l {
		if name == p.Name {
			return &l[i]
		}
	}
	return nil
}
