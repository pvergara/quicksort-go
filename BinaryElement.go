package main

type BinaryElement struct {
	element int
	left    *BinaryElement
	right   *BinaryElement
}

func (e *BinaryElement) add(element int) {
	if e.element == 0 {
		e.element = element
	} else if e.element > element {
		if e.left == nil {
			e.left = new(BinaryElement)
		}
		e.left.add(element)
	} else {
		if e.right == nil {
			e.right = new(BinaryElement)
		}
		e.right.add(element)
	}
}

func (e *BinaryElement) binarySort() []int {
	result := make([]int, 0)

	if e.left == nil {
		result = append(result, e.element)
		if e.right != nil {
			result = append(result, e.right.binarySort()...)
			e.right = nil
		}
		return result
	} else {
		result = append(result, e.left.binarySort()...)
		e.left = nil
		result = append(result, e.element)
		if e.right != nil {
			result = append(result, e.right.binarySort()...)
			e.right = nil
		}

	}

	return result
}
