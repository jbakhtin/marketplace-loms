package closer

import "context"

type Fn func(ctx context.Context) error

type Closer struct {
	fns []Fn
}

func (c *Closer) Close(ctx context.Context) error {
	for _, fn := range c.fns {
		_ = fn(ctx)
	}

	return nil
}
