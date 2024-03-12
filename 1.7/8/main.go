package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("not found")

func getValue(m map[string]string, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

type languages map[string]string

func (l languages) describe(lang string) (string, error) {
	descr, err := getValue(l, lang)
	if err != nil {
		return "", fmt.Errorf("error describing %s: %w", lang, err)
	}
	return descr, nil
}

var langs languages = languages{
	"go":     "is awesome",
	"python": "is everywhere",
	"php":    "just is",
}

func main() {
	descr, err := langs.describe("java")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(descr)

}
