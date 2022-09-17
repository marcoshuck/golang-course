package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	const (
		key   = "Test"
		value = "This is a definition"

		anotherKey   = "Cat"
		anotherValue = "Purrrrrr"
	)

	d := NewDictionary(
		DictionaryEntry{
			Key:   key,
			Value: value,
		},
		DictionaryEntry{
			Key:   anotherKey,
			Value: anotherValue,
		},
	)

	definition := d.Get(key)

	t.Logf("Key: %s, Value: %s", key, definition)

	if definition != value {
		t.Fail()
	}

	definition = d.Get(anotherKey)

	t.Logf("Key: %s, Value: %s", anotherKey, definition)

	if definition != anotherValue {
		t.Fail()
	}

	result := d.Get("")
	if result != "" {
		t.Fail()
	}
}

type Dictionary map[string]string

func (d Dictionary) Get(key string) string {
	return d[key]
}

type DictionaryEntry struct {
	Key   string
	Value string
}

func NewDictionary(entries ...DictionaryEntry) Dictionary {
	m := make(Dictionary, len(entries))

	for _, e := range entries {
		m[e.Key] = e.Value
	}

	return m
}
