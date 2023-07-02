package main

import (
	"fmt"
	"strings"
)

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	s := upperstring("Learning Go!")
	s = s + "!"
	anothers := "another string"
	s = s + upperstring(anothers)
	fmt.Println(s)
	fmt.Println(s.Upper())
}
