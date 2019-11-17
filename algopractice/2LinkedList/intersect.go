package linkedlist

func getLengthAndFinalValue(ll *LinkedList) (length int, finalAddress *Node) {
	node := ll.node
	length = 1
	for node.next != nil {
		length++
		node = node.next
	}
	finalAddress = node
	return
}
func doesIntersect(one *LinkedList, two *LinkedList) (intersects bool, lengthOne int, lengthTwo int) {
	lengthOne, addressOne := getLengthAndFinalValue(one)
	lengthTwo, addressTwo := getLengthAndFinalValue(two)

	intersects = addressOne == addressTwo
	return
}

func getIntersectionPoint(longer *LinkedList, shorter *LinkedList, difference int) *Node {
	firstNode := longer.node
	secondNode := shorter.node
	for difference > 0 {
		firstNode = firstNode.next
		difference--
	}

	for firstNode != nil {
		if firstNode == secondNode {
			return firstNode
		}
		firstNode = firstNode.next
		secondNode = secondNode.next
	}
	return nil
}

func GetIntersection(one *LinkedList, two *LinkedList) *Node {
	intersects, lengthOne, lengthTwo := doesIntersect(one, two)
	if !intersects {
		return nil
	}
	difference := lengthOne - lengthTwo
	if difference > 0 {
		return getIntersectionPoint(one, two, difference)
	} else {
		return getIntersectionPoint(two, one, difference)
	}
}
