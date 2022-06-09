package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSIndex(t *testing.T) {
	fmt.Print("foo")
	assert := assert.New(t)

	assert.Equal(0, BSIndex([]int{}, 10))
	assert.Equal(0, BSIndex([]int{10}, 10))
	assert.Equal(1, BSIndex([]int{5, 10}, 10))
	assert.Equal(1, BSIndex([]int{5, 10, 15}, 10))
	assert.Equal(-1, BSIndex([]int{5, 15}, 10))
	assert.Equal(-2, BSIndex([]int{5, 9, 15}, 10))

	assert.Equal(4, BSIndex([]int{5, 6, 7, 8, 10, 20}, 10))
	assert.Equal(3, BSIndex([]int{5, 6, 7, 10, 20}, 10))
	assert.Equal(2, BSIndex([]int{5, 6, 10, 20}, 10))
	assert.Equal(1, BSIndex([]int{6, 10, 20}, 10))
	assert.Equal(0, BSIndex([]int{10, 20}, 10))
	assert.Equal(4, BSIndex([]int{5, 6, 7, 8, 10}, 10))

	assert.Equal(-5, BSIndex([]int{5, 6, 7, 8, 10, 20}, 11))
	assert.Equal(-4, BSIndex([]int{5, 6, 7, 10, 20}, 11))
	assert.Equal(-3, BSIndex([]int{5, 6, 10, 20}, 11))
	assert.Equal(-2, BSIndex([]int{6, 10, 20}, 11))
	assert.Equal(-1, BSIndex([]int{10, 20}, 11))
	assert.Equal(-5, BSIndex([]int{5, 6, 7, 8, 10}, 11))
}

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)

	assert.False(BSContains([]int{}, 10))
	assert.True(BSContains([]int{10}, 10))
	assert.True(BSContains([]int{5, 10}, 10))
	assert.True(BSContains([]int{5, 10, 15}, 10))
	assert.False(BSContains([]int{5, 15}, 10))
	assert.False(BSContains([]int{5, 9, 15}, 10))

	assert.True(BSContains([]int{5, 6, 7, 8, 10, 20}, 10))
	assert.True(BSContains([]int{5, 6, 7, 10, 20}, 10))
	assert.True(BSContains([]int{5, 6, 10, 20}, 10))
	assert.True(BSContains([]int{6, 10, 20}, 10))
	assert.True(BSContains([]int{10, 20}, 10))
	assert.True(BSContains([]int{5, 6, 7, 8, 10}, 10))

	assert.False(BSContains([]int{5, 6, 7, 8, 10, 20}, 11))
	assert.False(BSContains([]int{5, 6, 7, 10, 20}, 11))
	assert.False(BSContains([]int{5, 6, 10, 20}, 11))
	assert.False(BSContains([]int{6, 10, 20}, 11))
	assert.False(BSContains([]int{10, 20}, 11))
	assert.False(BSContains([]int{5, 6, 7, 8, 10}, 11))
}

func TestIsSorted(t *testing.T) {
	assert := assert.New(t)

	s := []int{5, 6, 7, 8, 10, 20}
	assert.True(IsSorted(s))
	assert.True(IsSorted(s[:1]))
	assert.True(IsSorted(s[:2]))
	assert.True(IsSorted(s[:3]))
	assert.True(IsSorted(s[:4]))
	assert.True(IsSorted(s[:5]))
	assert.True(IsSorted(s[:6]))

	s = append(s, 1)
	assert.False(IsSorted(s))

	s = []int{2, 1}
	assert.False(IsSorted(s))
}

func TestInsertionSort(t *testing.T) {
	assert := assert.New(t)

	orig := []int{5, 18, 18, 6, 100, 20}

	s := (orig[:])
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[:1]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[:2]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[:3]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[:4]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[:5]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[:6]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[1:]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[2:]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[3:]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[4:]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[5:]
	InsertionSort(s)
	assert.True(IsSorted(s))

	s = orig[6:]
	InsertionSort(s)
	assert.True(IsSorted(s))
}

func TestMergeSort(t *testing.T) {
	assert := assert.New(t)

	orig := []int{5, 18, 18, 6, 100, 20}

	s := (orig[:])
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[:1]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[:2]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[:3]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[:4]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[:5]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[:6]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[1:]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[2:]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[3:]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[4:]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[5:]
	MergeSort(&s)
	assert.True(IsSorted(s))

	s = orig[6:]
	MergeSort(&s)
	assert.True(IsSorted(s))
}
