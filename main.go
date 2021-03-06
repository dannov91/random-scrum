package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/danoviedo91/scrum-rand/models"
	"gopkg.in/yaml.v2"
)

func main() {
	var source models.Source

	yamlPath := os.Args[1]
	yamlFile, err := ioutil.ReadFile(yamlPath)

	if err != nil {
		log.Fatalf("Cannot read yaml file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &source)
	if err != nil {
		log.Fatalf("Cannot unmarshal data: %v", err)
	}

	// This is to get different groups/members random order each time it is run
	rand.Seed(time.Now().UnixNano())

	// Shuffle members
	for groupIndex, group := range source.Groups {
		rand.Shuffle(len(group.Members), func(i, j int) {
			source.Groups[groupIndex].Members[i], source.Groups[groupIndex].Members[j] = source.Groups[groupIndex].Members[j], source.Groups[groupIndex].Members[i]
		})
	}

	// Shuffle groups
	rand.Shuffle(len(source.Groups), func(i, j int) { source.Groups[i], source.Groups[j] = source.Groups[j], source.Groups[i] })

	message := fmt.Sprintf("%v\n", source.Message)
	for _, group := range source.Groups {
		for _, member := range group.Members {
			message = fmt.Sprintf("%v%v\n", message, member)
		}
	}

	clipboard.WriteAll(message)
	fmt.Println("Success! Copied to clipboard.")
}
