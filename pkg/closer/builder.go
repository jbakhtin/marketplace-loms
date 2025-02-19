package closer

type Builder struct {
	closer Closer
}

func New() Builder {
	return Builder{}
}

func (b *Builder) Add(fn Fn) *Builder {
	b.closer.fns = append(b.closer.fns, fn)

	return b
}

func (b *Builder) Build() Closer {
	return b.closer
}
