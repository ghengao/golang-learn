package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "Heng GAO"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	s1 = strings.Title(module)
	s2 = strings.ToUpper(author)
	return
}
