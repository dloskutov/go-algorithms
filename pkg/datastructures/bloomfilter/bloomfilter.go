package bloomfilter

var hashMods = []int{167, 193, 199}

type BloomFilter struct {
	bits     []bool // @TODO: convert to bits
	hashMods []int
}

func (b *BloomFilter) Insert(key string) {
	for i := range b.hashMods {
		index := hash(key, b.hashMods[i])
		b.bits[index] = true
	}
}

func (b *BloomFilter) Contains(key string) bool {
	for i := range b.hashMods {
		index := hash(key, b.hashMods[i])
		if !b.bits[index] {
			return false
		}
	}
	return true
}

func New() *BloomFilter {
	max := 0
	for i := range hashMods {
		if hashMods[i] > max {
			max = hashMods[i]
		}
	}
	return &BloomFilter{
		bits:     make([]bool, max),
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
