package search

import (
	"fmt"
	"strings"
)

// BinarySearch - binary search by key
type BinarySearch struct {
	keys   []string
	values []int64
}

// Put - BinarySearch put value by key
func (bs *BinarySearch) Put(key string, value int64) {
	if bs.keys == nil {
		bs.keys = make([]string, 0)
	}
	if bs.values == nil {
		bs.values = make([]int64, 0)
	}
	size := len(bs.keys)
	r := bs.rank(key)
	if r == size {
		bs.keys = append(bs.keys, key)
		bs.values = append(bs.values, value)
		return
	}
	if strings.Compare(bs.keys[r], key) == 0 {
		bs.values[r] = value
		return
	}
	bs.keys = append(bs.keys, key)
	bs.values = append(bs.values, value)
	for i := size; i > r; i-- {
		bs.keys[i] = bs.keys[i-1]
		bs.values[i] = bs.values[i-1]
	}
	bs.keys[r] = key
	bs.values[r] = value
}

// Get - search value by key
func (bs *BinarySearch) Get(key string) (int64, error) {
	r := bs.rank(key)
	if r < len(bs.keys) {
		return bs.values[r], nil
	}
	return 0, fmt.Errorf("there is no value associated with key: %s", key)
}

func (bs *BinarySearch) rank(key string) int {
	leftIndex := 0
	rightIndex := len(bs.keys)
	for leftIndex < rightIndex {
		middle := (leftIndex + rightIndex) / 2
		keyToCompare := bs.keys[middle]
		comparedValue := strings.Compare(key, keyToCompare)
		if comparedValue == 0 {
			return middle
		}
		if comparedValue < 0 {
			rightIndex = middle - 1
			continue
		}
		if comparedValue > 0 {
			leftIndex = middle + 1
			continue
		}
	}
	return leftIndex
}
