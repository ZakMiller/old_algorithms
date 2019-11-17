package hashtable

import (
	"errors"
	"hash/fnv"
)

type HashTable [10]*Node

type key string
type value string

type Node struct {
	key key
	value value
	next *Node
}

func hash(s key) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % 10
}

func (ht *HashTable) Add(k key, v value) {
	h := hash(k)
	ht.assign(h, k, v)
}

func (ht *HashTable) Delete(k key) {
	h := hash(k)
	current := ht[h]
	var prev *Node = nil
	var next *Node = nil
	for current != nil {
		if current.key == k {
			if prev == nil && next == nil {
				ht[h] = nil
			} else if prev == nil {
				ht[h] = next
			} else if next == nil {
				prev.next = nil
			} else {
				prev.next = next
			}
			return
		}
		prev = current
		current = current.next
		next = current.next
	}
}

func (ht *HashTable) assign(i uint32, k key, v value) {
	current := ht[i]
	if current == nil {
		ht[i] = &Node{k, v, nil}
	} else {
		next := current
		for next.next != nil {
			next = next.next
		}
		next.next = &Node{k, v, nil}
	}
}

func (ht *HashTable) Get(k key) (value, error) {
	h := hash(k)
	n := ht[h]
	for n != nil {
		if n.key == k {
			return n.value, nil
		} else {
			n = n.next
		}
	}
	return "", errors.New("key not found")
}