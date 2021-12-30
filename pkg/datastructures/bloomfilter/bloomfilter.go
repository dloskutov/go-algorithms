package bloomfilter

const countOfBitsInByte = 8

var hashMods = []int{167, 193, 199}

type BloomFilter struct {
	max      int
	bits     []byte
	hashMods []int
}

func (b *BloomFilter) Insert(key string) {
	for i := range b.hashMods {
		index := hash(key, b.hashMods[i])
		b.setBit(index)
	}
}

func (b *BloomFilter) Contains(key string) bool {
	for i := range b.hashMods {
		index := hash(key, b.hashMods[i])
		if !b.isSet(index) {
			return false
		}
	}
	return true
}

func (b *BloomFilter) setBit(index int) {
	bytePosition := index / countOfBitsInByte
	offset := index - bytePosition*countOfBitsInByte
	b.bits[bytePosition] = b.bits[bytePosition] | (1 << offset)
}

func (b *BloomFilter) isSet(index int) bool {
	bytePosition := index / countOfBitsInByte
	offset := index - bytePosition*countOfBitsInByte
	var value byte = 1 << offset
	return b.bits[bytePosition]&value == value
}

func New() *BloomFilter {
	max := 0
	for i := range hashMods {
		if hashMods[i] > max {
			max = hashMods[i]
		}
	}
	return &BloomFilter{
		max:      max,
		bits:     make([]byte, max/countOfBitsInByte+1),
		hashMods: hashMods,
	}
}

func hash(key string, hashMod int) int {
	value := 0
	for index, r := range key {
		value += int(r) + int(int(r)/(index+1))
	}
	return value % hashMod
}
