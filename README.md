# ORES client for GO

Score one revision:
```go
client := ores.NewClient()

score, err := client.Damaging().ScoreOne(context.Background(), "enwiki", 1)

fmt.Println(score, err)
```

Score many revisions:
```go
client := ores.NewClient()

scores, err := client.Damaging().ScoreMany(context.Background(), "enwiki", 1, 112)

if err != nil {
  log.Panic(err)
}

for revID, score := range scores {
  fmt.Println(revID, score)
}
```

If you need pass custom `base URL` or another options you can use `ClientBuilder`:
```go
client := ores.NewBuilder().
  URL("https://ores.wikimedia.org/v3/scores").
  HTTPClient(&http.Client{}).
  Build()

score, err := client.Damaging().ScoreOne(context.Background(), "enwiki", 1)

fmt.Println(score, err)
```

### *Note that right now we are supporting `damaging` model only. For more information look at [https://ores.wikimedia.org/v3/#/](https://ores.wikimedia.org/v3/#/) and [https://ores.wikimedia.org/](https://ores.wikimedia.org/).