package main

import (
	"fmt"
	// "gopkg.in/mgo.v2"
	"bytes"
	"regexp"
)

type User struct {
	Title string
}

func main() {
	// hz := regexp.MustCompile("[\\p{Han}]")
	s := "如果有一天"
	fmt.Println(splitHan(s))
}
func splitHan(han string) string {
	bf := bytes.Buffer{}
	hz := regexp.MustCompile("[\\p{Han}]")
	for _, r := range han {
		str := string(r)
		if hz.MatchString(str) {
			bf.WriteString(str)
			bf.WriteString(" ")
			continue
		}
		bf.WriteString(str)
	}
	return bf.String()
}

