package starter

import "context"

type Func func(ctx context.Context) error

type Starter struct {
	fns []Func
}

func (s *Starter) Start(ctx context.Context) error {
	for _, fn := range s.fns {
		return fn(ctx)
	}
	return nil
}
