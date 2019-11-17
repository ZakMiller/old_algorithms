package hashtable

import (
	"fmt"
	"testing"
)

func TestKeyDoesntExist(t *testing.T) {
	ht := &HashTable{}
	val, err := ht.Get("key")
	verifyNotFound(val, err, t)
}

func TestSimpleRetrieve(t *testing.T) {
	ht := &HashTable{}
	ht.Add("key", "value")
	val, err := ht.Get("key")
	verifyFound(val, err, t, "value")
}

func TestKeyOverlapFirst(t *testing.T) {
	ht := &HashTable{}
	// a and k overlap
	ht.Add("a", "a value")
	ht.Add("k", "k value")
	val, err := ht.Get("a")
	verifyFound(val, err, t, "a value")

}

func TestKeyOverlapSecond(t *testing.T) {
	ht := &HashTable{}
	// a and k overlap
	ht.Add("a", "a value")
	ht.Add("k", "k value")
	val, err := ht.Get("k")
	verifyFound(val, err, t, "k value")
}

func TestKeyDoesntExistWithOtherKey(t *testing.T) {
	ht := &HashTable{}
	ht.Add("key2", "value")
	val, err := ht.Get("key")
	verifyNotFound(val, err, t)
}

func TestDelete(t *testing.T) {
	ht := &HashTable{}
	ht.Add("key", "value")
	ht.Delete("key")
	val, err := ht.Get("key")
	verifyNotFound(val, err, t)
}

func TestDeleteOther(t *testing.T) {
	ht := &HashTable{}
	ht.Add("key", "value")
	ht.Add("key2", "value2")
	ht.Delete("key")
	val, err := ht.Get("key2")
	verifyFound(val, err, t, "value2")
}

func verifyNotFound(val value, err error, t *testing.T) {
	if val != "" {
		t.Error("value found when not expecting")
	}
	if fmt.Sprint(err) != "key not found" {
		t.Error("incorrect error message")
	}
}

func verifyFound(val value, err error, t *testing.T, expected value) {
	if val != expected {
		t.Error("value not retrieved properly")
	}
	if err != nil {
		t.Error("error returned when it shouldn't be")
	}
}