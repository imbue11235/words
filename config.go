package words

type config struct {
	includeSymbols bool
	includePunctuation bool
	includeSpaces bool
	allowHyphenatedWords bool
}

func newDefaultConfig() *config {
	return &config{
		includeSymbols: false,
		includePunctuation: false,
		includeSpaces: false,
		allowHyphenatedWords: false,
	}
}

func (c *config) apply(options ...Option) {
	for _, opt := range options {
		opt.apply(c)
	}
}