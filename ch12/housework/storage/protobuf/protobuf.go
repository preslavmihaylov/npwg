package protobuf

import (
	"io"
	"io/ioutil"

	"housework/storage"

	pbgen "housework/idl"

	"google.golang.org/protobuf/proto"
)

func NewStorage() *Storage {
	return &Storage{}
}

type Storage struct{}

func (s *Storage) Load(r io.Reader) ([]*storage.Chore, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var chores pbgen.Chores
	err = proto.Unmarshal(b, &chores)
	return mapFromProtobuf(&chores), err
}

func (s *Storage) Flush(w io.Writer, chores []*storage.Chore) error {
	pbChores := mapToProtobuf(chores)
	b, err := proto.Marshal(pbChores)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func mapFromProtobuf(chores *pbgen.Chores) []*storage.Chore {
	res := []*storage.Chore{}
	for _, v := range chores.Chores {
		res = append(res, &storage.Chore{
			Complete:    v.Complete,
			Description: v.Description,
		})
	}

	return res
}

func mapToProtobuf(chores []*storage.Chore) *pbgen.Chores {
	res := &pbgen.Chores{}
	for _, v := range chores {
		res.Chores = append(res.Chores, &pbgen.Chore{
			Complete:    v.Complete,
			Description: v.Description,
		})
	}

	return res
}
