package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}
func (p People) Less(i, j int) bool {
	if p[i].birthDay.Sub(p[j].birthDay) > 0 {
		return true
	}
	if p[i].birthDay.Sub(p[j].birthDay) < 0 {
		return false
	}
	if p[i].firstName < p[j].firstName {
		return true
	}
	if p[i].firstName > p[j].firstName {
		return false
	}
	return p[i].lastName < p[j].lastName
}
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ivanIvanov2Date, err := time.Parse("2006-Jan-02", "2003-Aug-05")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	artemIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := People{
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Ivan", "Ivanov", ivanIvanov2Date},
		{"Artem", "Ivanov", artemIvanovDate},
	}
	sort.Sort(p)
	for _, o := range p {
		fmt.Println(o.firstName, o.lastName, o.birthDay)
	}
}
