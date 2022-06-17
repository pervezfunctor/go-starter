package main

import "golang.org/x/exp/constraints"

func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SliceEqual[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func AppendElem[T any](x []T, y T) []T {
	var z []T
	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]T, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y

	return z
}

func BSContains[T constraints.Ordered](s []T, v T) bool {
	if len(s) == 0 {
		return false
	}

	mid := len(s) / 2

	if v < s[mid] {
		return BSContains(s[:mid], v)
	}

	if v > s[mid] {
		return BSContains(s[mid+1:], v)
	}

	return true
}

func bsIndex[T constraints.Ordered](s []T, v T) (index int, found bool) {
	lo, hi, mid := 0, len(s), 0

	for lo < hi {
		mid = lo + (hi-lo)/2

		if v < s[mid] {
			hi = mid
		} else if v > s[mid] {
			lo = mid + 1
		} else {
			return mid, true
		}
	}

	return hi, false
}

func BSIndex[T constraints.Ordered](s []T, v T) int {
	index, _ := bsIndex(s, v)
	return index
}

func IsSorted[T constraints.Ordered](s []T) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}

	return true
}

func InsertSorted[T constraints.Ordered](sorted []T, v T) {
	for i, e := range sorted {
		if e > v {
			copy(sorted[i+1:], sorted[i:])
			sorted[i] = v
			break
		}
	}
}

func InsertionSort[T constraints.Ordered](s []T) {
	for i, v := range s {
		InsertSorted(s[:i+1], v)
	}
}

func Merge[T constraints.Ordered](fst, snd []T) []T {
	merged := make([]T, len(fst)+len(snd))

	for i, j, k := 0, 0, 0; i < len(fst) && j < len(snd); k++ {
		if fst[i] <= snd[j] {
			merged[k] = fst[i]
			i++
		} else {
			merged[k] = fst[j]
			j++
		}
	}

	return merged
}

func MergeSort[T constraints.Ordered](s *[]T) {
	if len(*s) < 20 {
		InsertionSort(*s)
		return
	}

	lst := *s
	mid := len(lst) / 2
	fst, snd := lst[:mid], lst[mid+1:]

	MergeSort(&fst)
	MergeSort(&snd)
	*s = Merge(fst, snd)
}
