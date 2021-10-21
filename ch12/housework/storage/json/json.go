package json

import (
	"encoding/json"
	"housework/storage"
	"io"
)

func NewStorage() *Storage {
	return &Storage{}
}

type Storage struct{}

func (s *Storage) Load(r io.Reader) ([]*storage.Chore, error) {
	var chores []*storage.Chore
	return chores, json.NewDecoder(r).Decode(&chores)
}

func (s *Storage) Flush(w io.Writer, chores []*storage.Chore) error {
	return json.NewEncoder(w).Encode(chores)
}
