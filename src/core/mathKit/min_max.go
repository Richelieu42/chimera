package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Integer | constraints.Float](numbers ...T) T {
	return mathutil.Min(numbers...)
}

func MinBy[T any](slice []T, comparator func(T, T) bool) T {
	return mathutil.MinBy(slice, comparator)
}

func Max[T constraints.Integer | constraints.Float](numbers ...T) T {
	return mathutil.Max(numbers...)
}

func MaxBy[T any](slice []T, comparator func(T, T) bool) T {
	return mathutil.MaxBy(slice, comparator)
}
