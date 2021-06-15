package main

import (
	"context"
	"fmt"
	"log"

	ores "github.com/protsack-stephan/mediawiki-ores-client"
)

func main() {
	client := ores.NewClient()
	ctx := context.Background()
	res, err := client.ScoreMany(ctx, "enwiki", []ores.Model{ores.ModelArticleQuality, ores.ModelGoodFaith}, 1, 122)

	if err != nil {
		log.Panic(err)
	}

	for _, score := range res.Scores {
		fmt.Println(score)
	}
}
