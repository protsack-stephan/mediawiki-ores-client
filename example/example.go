package main

import (
	"context"
	"fmt"
	"log"

	ores "github.com/protsack-stephan/mediawiki-ores-client"
)

func main() {
	client := ores.NewClient()

	score, err := client.Damaging().ScoreOne(context.Background(), "enwiki", 1)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(score)

	scores, err := client.Damaging().ScoreMany(context.Background(), "enwiki", 1, 112)

	if err != nil {
		log.Panic(err)
	}

	for revID, score := range scores {
		fmt.Println(revID, score)
	}
}
