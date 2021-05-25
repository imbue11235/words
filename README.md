# Words [![Test Status](https://github.com/imbue11235/words/workflows/Go/badge.svg)](https://github.com/imbue11235/words/actions?query=workflow%3A%Go%22) [![codecov](https://codecov.io/gh/imbue11235/words/branch/main/graph/badge.svg?token=XTJ42655U1)](https://codecov.io/gh/imbue11235/words)
Go package `words` provides capabilities for extracting words from a string, by a collection of rules.

## Rules

1. Invalid UTF8-strings will not be split
2. Hyphenated words will be treated as individual words unless disabled. E.g. `"small-town" => []{"small", "town"}`
3. If the character is a space, punctuation or symbol, it will be voided, unless disabled. E.g. `"my_string  here" => []{"my", "string", "here"}`
4. Characters of same type in sequence, will be put together.
5. If the current character is a lowercase, and the last character of the previous word was uppercase, the uppercase letter will be moved to the lowercase string. E.g. `"YAMLParser" => []{"YAML", "Parser"}`

## Installation

```sh
$ go get github.com/imbue11235/words
```

## Usage

### Basic usage

```go
words.Extract("Do you prefer camelCase to snake_case?") 
// => []string{"Do", "you", "prefer", "camel", "case", "to", "snake", "case")

words.Extract("YAMLParser")
// => []string{"YAML", "Parser"}

words.Extract("Bose QC35")
// => []string{"Bose", "QC", "35"}
```

### With options

To further customize the extraction, options can be passed to the extract-method.

#### Punctuation

To include [punctuation](https://en.wikipedia.org/wiki/General_Punctuation)

```go
words.Extract("So, now punctuation will be included.", words.IncludePunctuation())
// => []string{"So", ",", "now", "punctuation", "will", "be", "included", "."}
```

#### Spaces

To include spaces

```go
words.Extract("So   many   spaces", words.IncludeSpaces())
// => []string{"So", "   ", "many", "   ", "spaces"}
```

#### Symbols

To include symbols

```go
words.Extract("Some>String", words.IncludeSymbols())
// => []string{"Some", ">", "String"}
```

#### Hyphenated words 

To allow hyphenated words

```go
words.Extract("An anti-clockwise direction", words.AllowHyphenatedWords())
// => []string{"An", "anti-clockwise", "direction"}
```

#### Multiple options

To use multiple options at the same time

```go
words.Extract("Using multiple options!" words.IncludeSpaces(), words.IncludePunctuation())
// => []string{"Using", " ", "multiple", " ", "options", "!"}
```