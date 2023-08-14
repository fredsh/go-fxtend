package fx

type OptionEnhancer[T any] func(T) T

func OptionBuilder[T any](initial func() T, opts ...OptionEnhancer[T]) T {
	options := initial()
	for _, o := range opts {
		options = o(options)
	}
	return options
}

type OptionBuilderFluent[T any] struct {
	options T
}

func NewConfigBuilderFluent[T any](defaultConfig T) OptionBuilderFluent[T] {
	return OptionBuilderFluent[T]{
		options: defaultConfig,
	}
}

func (b *OptionBuilderFluent[T]) Config() T {
	return b.options
}

func (b *OptionBuilderFluent[T]) With(o OptionEnhancer[T]) *OptionBuilderFluent[T] {
	b.options = o(b.options)
	return b
}
