package array

type NearestNeighborContainer struct{}

func (c *NearestNeighborContainer) Insert(x int, y int) {
	// @TODO: need to implement
}

func (c *NearestNeighborContainer) Remove(x int, y int) error {
	// @TODO: need to implement
	return nil
}

func (c *NearestNeighborContainer) Search(x int, y int) {
	// @TODO: need to implement
}

func (c *NearestNeighborContainer) NearestNeighbor(x int, y int) (int, int) {
	// @TODO: need to implement
	return 0, 0
}

func (c *NearestNeighborContainer) PointsInRegion(x1 int, y1 int, x2 int, y2 int) [][2]int {
	// @TODO: need to implement
	return [][2]int{}
}

func (c *NearestNeighborContainer) Size() int {
	// @TODO: need to implement
	return 0
}

func New(points [][2]int) *NearestNeighborContainer {
	c := &NearestNeighborContainer{}

	for i := range points {
		c.Insert(points[i][0], points[i][1])
	}

	return c
}
