package main

import (
  "encoding/json"
  "fmt"
  "os"
  "strings"

  "github.com/polera/gonameparts"
)

func main() {
  parts, _ := json.Marshal(gonameparts.Parse(strings.Join(os.Args[1:], " ")))

  fmt.Printf("%v\n", string(parts))
}
