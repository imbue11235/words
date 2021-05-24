package words

// Option defines the interface for
// applying options to the extraction
type Option interface {
	apply(c *config)
}

// Include symbols option
type includeSymbols bool

// IncludeSymbols includes symbols in the extraction. E.g. "beer>food" => []{"beer", ">", "food"}
func IncludeSymbols() includeSymbols {
	return includeSymbols(true)
}

// IgnoreSymbols ignores symbols in the extraction. E.g. "beer>food" => []{"beer", "food"}
func IgnoreSymbols() includeSymbols {
	return includeSymbols(false)
}

func (i includeSymbols) apply(c *config) {
	c.includeSymbols = bool(i)
}

// Include punctuation option
type includePunctuation bool

// IncludePunctuation includes punctuation in extraction. E.g. "a.nested_path" => []{"a", ".", "nested", "-", "path"}
func IncludePunctuation() includePunctuation {
	return includePunctuation(true)
}

// IgnorePunctuation ignores punctuation in extraction. E.g. "a.nested_path" => []{"a", "nested", "path"}
func IgnorePunctuation() includePunctuation {
	return includePunctuation(false)
}

func (i includePunctuation) apply(c *config) {
	c.includePunctuation = bool(i)
}

// Include spaces option
type includeSpaces bool

// IncludeSpaces includes spaces in the extraction. E.g. "the  moon" => []{"the", "  ", "moon"}
func IncludeSpaces() includeSpaces {
	return includeSpaces(true)
}

// Ignore spaces includes spaces in the extraction. E.g. "the  moon" => []{"the", "moon"}
func IgnoreSpaces() includeSpaces {
	return includeSpaces(false)
}

func (i includeSpaces) apply(c *config) {
	c.includeSpaces = bool(i)
}

// Allow hyphenated words option
type allowHyphenatedWords bool

// Allow hyphenated words allows hyphenated words in the extraction.
// E.g. "a family-sized pizza" => []{"a", "family-sized", "pizza"}
func AllowHyphenatedWords() allowHyphenatedWords {
	return allowHyphenatedWords(true)
}

func (a allowHyphenatedWords) apply(c *config) {
	c.allowHyphenatedWords = true
}