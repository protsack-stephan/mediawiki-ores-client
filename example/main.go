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
	res, err := client.ScoreMany(ctx, "enwiki", []ores.Model{ores.ModelArticleQuality, ores.ModelGoodFaith, ores.ModelDamaging}, 1, 122)

	if err != nil {
		log.Panic(err)
	}

	for _, scores := range res.Scores {
		fmt.Println(scores.Articlequality.Score)
		fmt.Println(scores.Goodfaith.Score)
		fmt.Println(scores.Damaging.Score)
	}
}
