package twodtree

type TwoDTree struct{}

func (c *TwoDTree) Insert(x int, y int) {
	// @TODO: need to implement
}

func (c *TwoDTree) Remove(x int, y int) error {
	// @TODO: need to implement
	return nil
}

func (c *TwoDTree) Search(x int, y int) {
	// @TODO: need to implement
}

func (c *TwoDTree) NearestNeighbor(x int, y int) (int, int) {
	// @TODO: need to implement
	return 0, 0
}

func (c *TwoDTree) PointsInRegion(x1 int, y1 int, x2 int, y2 int) [][2]int {
	// @TODO: need to implement
	return [][2]int{}
}

func (c *TwoDTree) Size() int {
	// @TODO: need to implement
	return 0
}

func New(points [][2]int) *TwoDTree {
	c := &TwoDTree{}

	for i := range points {
		c.Insert(points[i][0], points[i][1])
	}

	return c
}
