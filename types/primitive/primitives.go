package primitive

import (
	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/copyable"
	"golang.org/x/exp/constraints"
)

// Primitive is a type that can be used to represent a primitive type.
type Primitive[T constraints.Ordered] struct {
	Value T
}

// Of returns a new primitive type with the given value.
func Of[T constraints.Ordered](value T) Primitive[T] {
	return Primitive[T]{Value: value}
}

// OfList returns a new primitive type with the given values.
func OfList[T constraints.Ordered](value ...T) []Primitive[T] {
	l := make([]Primitive[T], len(value))
	for i, v := range value {
		l[i] = Primitive[T]{v}
	}
	return l
}

var _ comparable.Comparable[Primitive[int]] = Primitive[int]{}

// CompareTo compares this primitive type with the given one.
func (p Primitive[T]) CompareTo(other Primitive[T]) comparable.Compared {
	if p.Value < other.Value {
		return comparable.Less
	}
	if p.Value > other.Value {
		return comparable.Greater
	}
	return comparable.Equal
}

var _ copyable.Clonable[Primitive[int]] = Primitive[int]{}

// Clone returns a new primitive type with the same value as this one.
func (p Primitive[T]) Clone() Primitive[T] {
	return Primitive[T]{p.Value}
}
