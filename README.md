# ORES client for GO

Score one or many revisions:
```go
client := ores.NewClient()
res, err := client.ScoreMany(context.Background(), "enwiki", []ores.Model{ores.ModelArticleQuality, ores.ModelGoodFaith, ores.ModelDamaging}, 1, 122)

if err != nil {
  log.Panic(err)
}

for _, scores := range res.Scores {
  fmt.Println(scores.Articlequality.Score)
  fmt.Println(scores.Goodfaith.Score)
  fmt.Println(scores.Damaging.Score)
}
```

If you need pass custom `base URL` or another options you can use `ClientBuilder`:
```go
client := ores.NewBuilder().
  URL("https://ores.wikimedia.org/v3/scores").
  HTTPClient(&http.Client{}).
  Build()
```

### *Note that right now we are supporting `damaging` model only. For more information look at [https://ores.wikimedia.org/v3/#/](https://ores.wikimedia.org/v3/#/) and [https://ores.wikimedia.org/](https://ores.wikimedia.org/).