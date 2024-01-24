package object

import "strings"

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) Inspect() string {
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, pair.Key.Inspect()+":"+pair.Value.Inspect())
	}
	return "{" + strings.Join(pairs, ", ") + "}"
}

// Hashable interface
type Hashable interface {
	HashKey() HashKey
}
