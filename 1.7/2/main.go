package main

import (
	"fmt"
	"strings"
)

type error interface {
	Error() string
}

type lookupError struct {
	src    string
	substr string
}

func (e lookupError) Error() string {
	return fmt.Sprintf("'%s' not found in '%s'", e.substr, e.src)

}

func index0f(src string, substr string) (int, error) {
	idx := strings.Index(src, substr)
	if idx == -1 {
		//создаем и возращаем ошибку типа lookupError
		return -1, lookupError{src, substr}
	}
	return idx, nil

}

func main() {
	src := "go is awesome"
	for _, substr := range []string{"go", "js"} {
		if res, err := index0f(src, substr); err != nil {
			fmt.Printf("index0f(%#v, %#v) failed: %v\n", src, substr, err)
		} else {
			fmt.Printf("index0f(%#v, %#v) = %v\n", src, substr, res)
		}
	}
	_, err := index0f(src, "js")
	if err, ok := err.(lookupError); ok {
		fmt.Println("err.src:", err.src)
		fmt.Println("err.substr:", err.substr)
	}

}
