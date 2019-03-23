package label

import (
  "fmt"
  "github.com/BurntSushi/toml"
)

type label struct {
  Key   string 
  Value string
}

type Label struct {
  Label []label `toml:"label"`
}

func MakeLabelMap(s string) (map[string]string, error) {
  var l Label
  lmap := make(map[string]string)
  _, err := toml.Decode(s, &l)
  if (err != nil) {
    fmt.Println("Error Occurred")
    return nil, err
  }

  for i, _ := range(l.Label) {
    lmap[l.Label[i].Key] = l.Label[i].Value
  }

  return lmap, nil
}

