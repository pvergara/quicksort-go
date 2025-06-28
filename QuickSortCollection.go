package main

type QuickSortCollection struct {
	collection *BinaryElement
}

func (c *QuickSortCollection) addAll(got []int) {
	for _, element := range got {
		c.add(element)
	}
}

func (c *QuickSortCollection) add(element int) {
	if c.collection == nil {
		c.collection = new(BinaryElement)
	}
	c.collection.add(element)
}

func (c *QuickSortCollection) sort() []int {
	return c.collection.binarySort()
}
