package main

import (
	"context"
	"fmt"

	ores "github.com/protsack-stephan/mediawiki-ores-client"
)

func main() {
	client := ores.NewClient()

	score, err := client.Damaging().ScoreOne(context.Background(), "enwiki", 1)

	fmt.Println(score, err)
}
