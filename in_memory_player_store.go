package main

type InMemoryStore struct {
	store map[string]int
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
