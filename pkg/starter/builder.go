package starter

type Builder struct {
	starter Starter
}

func New() Builder {
	return Builder{}
}

func (b *Builder) Add(fn Func) *Builder {
	b.starter.fs = append(b.starter.fs, fn)
	return b
}

func (b *Builder) Build() Starter {
	return b.starter
}
