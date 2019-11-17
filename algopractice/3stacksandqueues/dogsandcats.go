package stacksandqueues

import (
	"fmt"
	"strings"
)

type animal interface {
	isAnimal() bool
}

type dog struct {
}

func (d dog) isAnimal() bool {
	return true
}

type cat struct {
}

func (c cat) isAnimal() bool {
	return true
}

type animalNode struct {
	previous           *animalNode
	next               *animalNode
	previousSameAnimal *animalNode
	value              animal
}

type animalQueue struct {
	first    *animalNode
	last     *animalNode
	firstDog *animalNode
	firstCat *animalNode
	lastDog  *animalNode
	lastCat  *animalNode
}

func (aq *animalQueue) String() string {
	b := strings.Builder{}
	node := aq.first
	for node != nil {
		switch node.value.(type) {
		case dog:
			b.WriteString("dog")
		case cat:
			b.WriteString("cat")
		}
		b.WriteString(" ")
		node = node.next
	}
	return b.String()

}

func (aq *animalQueue) enqueue(a animal) {
	next := &animalNode{nil, aq.first, nil, a}
	switch a.(type) {
	case dog:
		if aq.firstDog == nil {
			aq.firstDog, aq.lastDog = next, next
		} else {
			aq.firstDog.previousSameAnimal = next
			aq.firstDog = next
		}

	case cat:
		if aq.firstCat == nil {
			aq.firstCat, aq.lastCat = next, next
		} else {
			aq.firstCat.previousSameAnimal = next
			aq.firstCat = next
		}
	}
	if aq.first == nil {
		aq.first, aq.last = next, next
	} else {
		aq.first.previous = next
		aq.first = next
	}
}

func (aq *animalQueue) dequeueAny() (animal, bool) {
	if aq.last == nil {
		return nil, false
	}
	node := aq.last
	if node.previous != nil {
		node.previous.next = node.next
	} else {
		aq.first = node.next
	}
	if node.next != nil {
		node.next.previous = node.previous
	} else {
		aq.last = node.previous
	}

	switch node.value.(type) {
	case dog:
		aq.lastDog = aq.lastDog.previousSameAnimal
	case cat:
		aq.lastCat = aq.lastCat.previousSameAnimal
	}
	return node.value, true
}

func (aq *animalQueue) dequeueCat() (animal, bool) {
	before := aq.String()
	if aq.lastCat == nil {
		return nil, false
	}
	node := aq.lastCat
	value := node.value
	if node == aq.last {
		aq.last = aq.last.previous
	}
	aq.lastCat = aq.lastCat.previousSameAnimal
	if node.previous != nil {
		node.previous.next = node.next
	} else {
		aq.first = node.next
	}
	if node.next != nil {
		node.next.previous = node.previous
	} else {
		aq.last = node.previous
	}

	after := aq.String()
	fmt.Printf("DEQUEUE CAT: before: %v after: %v\n", before, after)
	return value, true
}

func (aq *animalQueue) dequeueDog() (animal, bool) {
	before := aq.String()
	if aq.lastDog == nil {
		return nil, false
	}
	node := aq.lastDog
	value := node.value
	if node == aq.last {
		aq.last = aq.last.previous
	}
	aq.lastDog = aq.lastDog.previousSameAnimal
	if node.previous != nil {
		node.previous.next = node.next
	} else {
		aq.first = node.next
	}
	if node.next != nil {
		node.next.previous = node.previous
	} else {
		aq.last = node.previous
	}

	after := aq.String()
	fmt.Printf("DEQUEUE DOG: before: %v after: %v\n", before, after)
	return value, true
}
