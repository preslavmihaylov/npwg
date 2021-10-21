package server

import (
	"context"
	"fmt"
	pbgen "housework/idl"
	"sync"
)

type Rosie struct {
	mu     sync.Mutex
	chores []*pbgen.Chore

	*pbgen.UnimplementedRobotMaidServer
}

func (r *Rosie) Add(ctx context.Context, chores *pbgen.Chores) (*pbgen.Response, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.chores = append(r.chores, chores.Chores...)

	return &pbgen.Response{Message: "ok"}, nil
}

func (r *Rosie) Complete(ctx context.Context, req *pbgen.CompleteRequest) (*pbgen.Response, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.chores == nil || req.ChoreNumber < 1 || int(req.ChoreNumber) > len(r.chores) {
		return nil, fmt.Errorf("chore %d not found", req.ChoreNumber)
	}

	r.chores[req.ChoreNumber-1].Complete = true

	return &pbgen.Response{Message: "ok"}, nil
}

func (r *Rosie) List(ctx context.Context, req *pbgen.Empty) (*pbgen.Chores, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.chores == nil {
		r.chores = make([]*pbgen.Chore, 0)
	}

	return &pbgen.Chores{Chores: r.chores}, nil
}
