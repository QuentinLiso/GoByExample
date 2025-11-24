package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:\t", s.Contains("test", "es"))
	p("Count:\t", s.Count("test", "t"))
	p("HasPrefix:\t", s.HasPrefix("test", "te"))
	p("HasSuffix:\t", s.HasSuffix("test", "st"))
	p("Index:\t", s.Index("test", "e"))
	p("Join:\t", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:\t", s.Repeat("a", 5))
	p("Replace:\t", s.Replace("foo", "o", "0", -1))
	p("Replace:\t", s.Replace("foo", "o", "0", 1))
	p("Split:\t", s.Split("a-b-c-d-e", "-"))
	p("ToLower:\t", s.ToLower("TEST"))
	p("ToUpper:\t", s.ToUpper("test"))
}
