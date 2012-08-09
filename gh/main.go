package main

import (
	"fmt"
	"github.com/injekt/github.go"
	"log"
)

func main() {
	user, err := github.GetUser("injekt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	repo, err := github.GetRepo("injekt/slop")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)

	gist, err := github.GetGist("3302678")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gist)
	for _, file := range gist.Files {
		fmt.Println(file)
	}
}
