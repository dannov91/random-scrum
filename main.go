package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Group ...
type Group struct {
	Name    string
	Members []string
}

// Groups ...
type Groups []Group

func main() {
	groups := Groups{
		{
			Name:    "Dev",
			Members: []string{"Capi", "Moslo", "Danier", "Grego", "Mañe"},
		},
		{
			Name:    "QA",
			Members: []string{"Mario", "Hillary"},
		},
		{
			Name:    "Design",
			Members: []string{"Nicole"},
		},
	}

	fmt.Println("Hey People, good morning!")
	fmt.Println("--")

	rand.Seed(time.Now().UnixNano())

	for _, group := range groups {
		rand.Shuffle(len(group.Members), func(i, j int) { group.Members[i], group.Members[j] = group.Members[j], group.Members[i] })
	}

	rand.Shuffle(len(groups), func(i, j int) { groups[i], groups[j] = groups[j], groups[i] })

	for _, group := range groups {
		for _, member := range group.Members {
			fmt.Println(member)
		}
	}
}
