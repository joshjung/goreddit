package main

import (
  "log"
  "fmt"
  "github.com/joshjung/goreddit"
)

func main() {
  items, err := goreddit.Get("golang")
  if err != nil {
    log.Fatal(err)
  }

  for _, item := range items {
    fmt.Println(item)
  }
}