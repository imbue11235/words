package words

type Option interface {
	apply(c *config)
}

type includeSymbols bool

func IncludeSymbols() includeSymbols {
	return includeSymbols(true)
}

func IgnoreSymbols() includeSymbols {
	return includeSymbols(false)
}

func (i includeSymbols) apply(c *config) {
	c.includeSymbols = bool(i)
}

type includePunctuation bool

func IncludePunctuation() includePunctuation {
	return includePunctuation(true)
}

func IgnorePunctuation() includePunctuation {
	return includePunctuation(false)
}

func (i includePunctuation) apply(c *config) {
	c.includePunctuation = bool(i)
}

type includeSpaces bool

func IncludeSpaces() includeSpaces {
	return includeSpaces(true)
}

func IgnoreSpaces() includeSpaces {
	return includeSpaces(false)
}

func (i includeSpaces) apply(c *config) {
	c.includeSpaces = bool(i)
}

type allowHyphenatedWords bool

func AllowHyphenatedWords() allowHyphenatedWords {
	return allowHyphenatedWords(true)
}

func (a allowHyphenatedWords) apply(c *config) {
	c.allowHyphenatedWords = true
}