package bloomfilter

type BloomFilter struct {
	bits []byte // @TODO: convert to bits
}

func (b *BloomFilter) Insert(key string) {
	// @TODO: need to implement
}

func (b *BloomFilter) Contains() bool {
	// @TODO: need to implement
	return false
}

func New(size int) *BloomFilter {
	return &BloomFilter{
		bits: make([]byte, size),
	}
}
