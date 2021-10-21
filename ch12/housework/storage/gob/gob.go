package gob

import (
	"encoding/gob"
	"housework/storage"
	"io"
)

func NewStorage() *Storage {
	return &Storage{}
}

type Storage struct{}

func (s *Storage) Load(r io.Reader) ([]*storage.Chore, error) {
	var chores []*storage.Chore
	return chores, gob.NewDecoder(r).Decode(&chores)
}

func (s *Storage) Flush(w io.Writer, chores []*storage.Chore) error {
	return gob.NewEncoder(w).Encode(chores)
}
