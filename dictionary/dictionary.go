package dictionary

// Dictionary is an interface that is used to look up words
type Dictionary interface {
	DefineWord(word string) []string
}

