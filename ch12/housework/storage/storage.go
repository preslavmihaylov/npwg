package storage

import "io"

type Chore struct {
	Complete    bool
	Description string
}

type Storage interface {
	Load(r io.Reader) ([]*Chore, error)
	Flush(w io.Writer, chores []*Chore) error
}
