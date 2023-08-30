package words

// Option defines the interface for
// applying options to the extraction
type Option func(c *config)

// IncludeSymbols includes symbols in the extraction. E.g. "beer>food" => []{"beer", ">", "food"}
func IncludeSymbols() Option {
	return func(c *config) {
		c.includeSymbols = true
	}
}

// IncludePunctuation includes punctuation in extraction. E.g. "a.nested_path" => []{"a", ".", "nested", "-", "path"}
func IncludePunctuation() Option {
	return func(c *config) {
		c.includePunctuation = true
	}
}

// IncludeSpaces includes spaces in the extraction. E.g. "the  moon" => []{"the", "  ", "moon"}
func IncludeSpaces() Option {
	return func(c *config) {
		c.includeSpaces = true
	}
}

// AllowHyphenatedWords allows hyphenated words in the extraction.
// E.g. "a family-sized pizza" => []{"a", "family-sized", "pizza"}
func AllowHyphenatedWords() Option {
	return func(c *config) {
		c.allowHyphenatedWords = true
	}
}

// WithIgnoredRunes tells the extractor to ignore these runes
// when they are encountered, simply adding them to the output
// as the rune was of most recent rune kind.
// E.g. => WithIgnoredRunes('.') "Etc. and so on" becomes => []{"Etc.", "and", "so", "on"}
func WithIgnoredRunes(runes ...rune) Option {
	return func(c *config) {
		c.ignoredRunes = append(c.ignoredRunes, runes...)
	}
}

// WithIgnoredRuneKinds tells the extractor to ignore these rune kinds
// when they are encountered, simply adding them to the output
// as the rune was of most recent rune kind.
func WithIgnoredRuneKinds(runeKinds ...RuneKind) Option {
	return func(c *config) {
		c.ignoredRunesKinds = append(c.ignoredRunesKinds, runeKinds...)
	}
}
