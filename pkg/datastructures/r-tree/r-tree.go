package rtree

type RTree struct{}

func (c *RTree) Insert(x int, y int) {
	// @TODO: need to implement
}

func (c *RTree) Remove(x int, y int) error {
	// @TODO: need to implement
	return nil
}

func (c *RTree) Search(x int, y int) {
	// @TODO: need to implement
}

func (c *RTree) NearestNeighbor(x int, y int) (int, int) {
	// @TODO: need to implement
	return 0, 0
}

func (c *RTree) PointsInRegion(x1 int, y1 int, x2 int, y2 int) [][2]int {
	// @TODO: need to implement
	return [][2]int{}
}

func (c *RTree) Size() int {
	// @TODO: need to implement
	return 0
}

func New(points [][2]int) *RTree {
	c := &RTree{}

	for i := range points {
		c.Insert(points[i][0], points[i][1])
	}

	return c
}
