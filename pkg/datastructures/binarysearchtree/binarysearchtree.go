package binarysearchtree

type BinarySearchTree struct{}

func (bst *BinarySearchTree) Put(key int, value interface{}) error {
	// @TODO: need to implement
	return nil
}

func (bst *BinarySearchTree) Search(key int) (interface{}, error) {
	// @TODO: need to implement
	return nil, nil
}

func (bst *BinarySearchTree) Remove(key int) error {
	// @TODO: need to implement
	return nil
}

func New(data map[int]interface{}) (*BinarySearchTree, error) {
	bst := &BinarySearchTree{}
	for key, value := range data {
		err := bst.Put(key, value)
		if err != nil {
			return nil, err
		}
	}
	return bst, nil
}
