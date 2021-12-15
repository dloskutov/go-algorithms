package disjointset

type DisjoinSet struct {
	sets map[int]int
}

func (d *DisjoinSet) FindPartition(elem int) int {
	partition := d.sets[elem]
	for partition != d.sets[partition] {
		partition = d.sets[partition]
		d.sets[elem] = partition
	}
	return partition
}

func (d *DisjoinSet) Merge(first int, second int) {
	d.sets[first] = d.sets[second]
}

func (d *DisjoinSet) AreDisjoint(first int, second int) bool {
	return d.FindPartition(first) != d.FindPartition(second)
}

func (d *DisjoinSet) Add(elem int) bool {
	_, ok := d.sets[elem]
	if !ok {
		d.sets[elem] = elem
		return true
	}
	return false
}

func New(data []int) *DisjoinSet {
	d := &DisjoinSet{
		sets: make(map[int]int),
	}
	for i := range data {
		_ = d.Add(data[i])
	}
	return d
}
