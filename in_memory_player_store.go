package main

type InMemoryStore struct {
	store map[string]int
}

// GetLeague implements playerStore.
func (i *InMemoryStore) GetLeague() League {
	var league []Player

	for name, score := range i.store {
		league = append(league, Player{name, score})
	}

	return league
}

// GetPlayerScore implements playerStore.
func (i *InMemoryStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin implements playerStore.
func (i *InMemoryStore) RecordWin(name string) {
	i.store[name]++
}

func NewInMemoryStore() playerStore {
	return &InMemoryStore{map[string]int{}}
}
