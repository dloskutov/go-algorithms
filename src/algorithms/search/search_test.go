package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSequentialSearch(t *testing.T) {
	sequentialSearch := new(SequentialSearch)

	value, err := sequentialSearch.Get("non-exist key")
	assert.Equal(t, int64(0), value)
	assert.Error(t, err)

	sequentialSearch.Put("first", 1)
	sequentialSearch.Put("second", 2)
	sequentialSearch.Put("third", 3)

	value, err = sequentialSearch.Get("second")
	assert.Equal(t, int64(2), value)
	assert.Nil(t, err)

	sequentialSearch.Put("second", 20)
	value, err = sequentialSearch.Get("second")
	assert.Equal(t, int64(20), value)
	assert.Nil(t, err)
}

func TestBinarySearch(t *testing.T) {
	binarySearch := new(BinarySearch)

	value, err := binarySearch.Get("non-exist key")
	assert.Equal(t, int64(0), value)
	assert.Error(t, err)

	binarySearch.Put("first", 1)
	binarySearch.Put("second", 2)
	binarySearch.Put("third", 3)
	binarySearch.Put("alpha", 45)

	value, err = binarySearch.Get("second")
	assert.Equal(t, int64(2), value)
	assert.Nil(t, err)

	binarySearch.Put("second", 20)
	value, err = binarySearch.Get("second")
	assert.Equal(t, int64(20), value)
	assert.Nil(t, err)

	value, err = binarySearch.Get("alpha")
	assert.Equal(t, int64(45), value)
	assert.Nil(t, err)
}

func TestBinarySearchTree(t *testing.T) {
	binarySearchTree := new(BinarySearchTree)

	value, err := binarySearchTree.Get("non-exist key")
	assert.Equal(t, int64(0), value)
	assert.Error(t, err)

	binarySearchTree.Put("g", 40)
	binarySearchTree.Put("h", 80)
	binarySearchTree.Put("a", 10)
	binarySearchTree.Put("c", 20)
	binarySearchTree.Put("e", 30)
	binarySearchTree.Put("b", 50)
	binarySearchTree.Put("d", 60)
	binarySearchTree.Put("f", 70)

	value, err = binarySearchTree.Get("c")
	assert.Equal(t, int64(20), value)
	assert.Nil(t, err)

	binarySearchTree.Put("g", 100)
	value, err = binarySearchTree.Get("g")
	assert.Equal(t, int64(100), value)
	assert.Nil(t, err)

	value, err = binarySearchTree.Get("f")
	assert.Equal(t, int64(70), value)
	assert.Nil(t, err)
}

func TestTwoThreeSearchTree(t *testing.T) {
	twoThreeSearchTree := new(TwoThreeSearchTree)

	value, err := twoThreeSearchTree.Get("non-exist key")
	assert.Equal(t, int64(0), value)
	assert.Error(t, err)

	arr := []struct {
		key   string
		value int64
	}{
		{
			key:   "b",
			value: 4,
		},
		{
			key:   "a",
			value: 8,
		},
		{
			key:   "r",
			value: 10,
		},
		{
			key:   "n",
			value: 7,
		},
		{
			key:   "o",
			value: 3,
		},
		{
			key:   "i",
			value: 9,
		},
		{
			key:   "x",
			value: 2,
		},
		{
			key:   "s",
			value: 86,
		},
		{
			key:   "y",
			value: 31,
		},
		{
			key:   "m",
			value: 0,
		},
		{
			key:   "df",
			value: 3,
		},
		{
			key:   "l",
			value: 82,
		},
		{
			key:   "o",
			value: 3,
		},
		{
			key:   "i",
			value: 9,
		},
		{
			key:   "x",
			value: 2,
		},
		{
			key:   "s",
			value: 86,
		},
		{
			key:   "p",
			value: 3,
		},
		{
			key:   "y",
			value: 9,
		},
		{
			key:   "v",
			value: 2,
		},
		{
			key:   "c",
			value: 86,
		},
		{
			key:   "q",
			value: 3,
		},
	}

	for _, item := range arr {
		twoThreeSearchTree.Put(item.key, item.value)
	}

	value, err = twoThreeSearchTree.Get("r")
	assert.Equal(t, int64(10), value)
	assert.Nil(t, err)

	twoThreeSearchTree.Put("g", 100)
	value, err = twoThreeSearchTree.Get("g")
	assert.Equal(t, int64(100), value)
	assert.Nil(t, err)

	value, err = twoThreeSearchTree.Get("a")
	assert.Equal(t, int64(8), value)
	assert.Nil(t, err)

	value, err = twoThreeSearchTree.Get("v")
	assert.Equal(t, int64(2), value)
	assert.Nil(t, err)

	value, err = twoThreeSearchTree.Get("c")
	assert.Equal(t, int64(86), value)
	assert.Nil(t, err)

	value, err = twoThreeSearchTree.Get("p")
	assert.Equal(t, int64(3), value)
	assert.Nil(t, err)
}

func TestRedBlackBinarySearchTree(t *testing.T) {
	tree := new(RedBlackBinarySearchTree)

	value, err := tree.Get("non-exist key")
	assert.Equal(t, int64(0), value)
	assert.Error(t, err)

	arr := []struct {
		key   string
		value int64
	}{
		{
			key:   "b",
			value: 4,
		},
		{
			key:   "a",
			value: 8,
		},
		{
			key:   "r",
			value: 10,
		},
		{
			key:   "n",
			value: 7,
		},
		{
			key:   "o",
			value: 3,
		},
		{
			key:   "i",
			value: 9,
		},
		{
			key:   "x",
			value: 2,
		},
		{
			key:   "s",
			value: 86,
		},
		{
			key:   "y",
			value: 31,
		},
		{
			key:   "m",
			value: 0,
		},
		{
			key:   "df",
			value: 3,
		},
		{
			key:   "l",
			value: 82,
		},
		{
			key:   "o",
			value: 3,
		},
		{
			key:   "i",
			value: 9,
		},
		{
			key:   "x",
			value: 2,
		},
		{
			key:   "s",
			value: 86,
		},
		{
			key:   "p",
			value: 3,
		},
		{
			key:   "y",
			value: 9,
		},
		{
			key:   "v",
			value: 2,
		},
		{
			key:   "c",
			value: 86,
		},
		{
			key:   "q",
			value: 3,
		},
		{
			key:   "we",
			value: 48,
		},
		{
			key:   "eb",
			value: 82,
		},
		{
			key:   "ktr",
			value: 103,
		},
		{
			key:   "aasn",
			value: 73,
		},
		{
			key:   "qo",
			value: 32,
		},
		{
			key:   "mni",
			value: 94,
		},
		{
			key:   "opx",
			value: 23,
		},
		{
			key:   "zs",
			value: 86,
		},
		{
			key:   "ty",
			value: 331,
		},
		{
			key:   "prm",
			value: 32,
		},
		{
			key:   "adf",
			value: 3,
		},
		{
			key:   "wl",
			value: 898,
		},
		{
			key:   "owr",
			value: 33,
		},
		{
			key:   "ji",
			value: 9,
		},
		{
			key:   "sx",
			value: 2,
		},
		{
			key:   "ts",
			value: 86,
		},
		{
			key:   "up",
			value: 3,
		},
		{
			key:   "xy",
			value: 9,
		},
		{
			key:   "vnv",
			value: 2,
		},
		{
			key:   "pc",
			value: 86,
		},
	}

	for _, item := range arr {
		tree.Put(item.key, item.value)
	}

	value, err = tree.Get("r")
	assert.Equal(t, int64(10), value)
	assert.Nil(t, err)

	tree.Put("g", 100)
	value, err = tree.Get("g")
	assert.Equal(t, int64(100), value)
	assert.Nil(t, err)

	value, err = tree.Get("a")
	assert.Equal(t, int64(8), value)
	assert.Nil(t, err)

	value, err = tree.Get("v")
	assert.Equal(t, int64(2), value)
	assert.Nil(t, err)

	value, err = tree.Get("c")
	assert.Equal(t, int64(86), value)
	assert.Nil(t, err)

	value, err = tree.Get("p")
	assert.Equal(t, int64(3), value)
	assert.Nil(t, err)

	value, err = tree.Get("vnv")
	assert.Equal(t, int64(2), value)
	assert.Nil(t, err)
}
