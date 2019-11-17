package stacksandqueues

import (
	"testing"
)

func threeDogs() *animalQueue {
	aq := &animalQueue{}
	aq.enqueue(dog{})
	aq.enqueue(dog{})
	aq.enqueue(dog{})
	return aq
}

func dequeue2() *animalQueue {
	aq := &animalQueue{}
	aq.enqueue(dog{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.dequeueAny()
	aq.dequeueAny()
	return aq
}

func dequeueCat() *animalQueue {
	aq := &animalQueue{}
	aq.enqueue(dog{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.dequeueCat()
	return aq
}

func dequeueCats() *animalQueue {
	aq := &animalQueue{}
	aq.enqueue(dog{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.dequeueCat()
	aq.dequeueCat()
	return aq
}

func dequeueCoupleDogs() *animalQueue {
	aq := &animalQueue{}
	aq.enqueue(dog{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.enqueue(dog{})
	aq.enqueue(cat{})
	aq.dequeueDog()
	aq.dequeueDog()
	return aq
}

func TestAnimalQueue(t *testing.T) {
	tests := []struct {
		aq       *animalQueue
		expected string
	}{
		{threeDogs(), "dog dog dog "},
		{dequeue2(), "cat dog cat "},
		{dequeueCat(), "cat dog dog dog "},
		{dequeueCats(), "dog dog dog "},
		{dequeueCoupleDogs(), "cat dog cat "},
	}

	for _, test := range tests {
		actual := test.aq.String()
		if actual != test.expected {
			t.Errorf("animal queue error actual %v expected %v", actual, test.expected)
		}
	}

}
