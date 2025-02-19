package starter

import "context"

type Func func(ctx context.Context) error

type Starter struct {
	fs []Func
}

func (s *Starter) Start(ctx context.Context) error {
	return nil
}
