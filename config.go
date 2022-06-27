package words

type config struct {
	includeSymbols       bool
	includePunctuation   bool
	includeSpaces        bool
	allowHyphenatedWords bool
	ignoredRunes         []rune
}

// newDefaultConfig defines the standards
// of the word extractor
func newDefaultConfig() *config {
	return &config{
		includeSymbols:       false,
		includePunctuation:   false,
		includeSpaces:        false,
		allowHyphenatedWords: false,
		ignoredRunes:         make([]rune, 0),
	}
}

// apply applies all the options to the
// current config
func (c *config) apply(options ...Option) {
	for _, opt := range options {
		opt(c)
	}
}
