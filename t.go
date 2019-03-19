package main

import (
  "fmt"
_  "github.com/salmanb/kubedeploy/parts/label"
  "github.com/salmanb/kubedeploy/parts/container"
  "io/ioutil"
  "log"
  "os"
)

func main() {
  ioreader, _ := os.Open("dply.toml")
//  a, b := label.MakeLabelMap(ioreader)
//  if (b != nil) {
//    log.Fatal(b)
//  }
//  for k, v := range(a) {
//    fmt.Printf("%v => %v\n", k, v)
//  }

  data, err := ioutil.ReadAll(ioreader)
  if (err != nil) {
    log.Fatal(err)
  }
  c, err := container.MakeList(string(data))
  if (err != nil) {
    log.Fatal(err)
  }
  for i, _ := range c {
    fmt.Printf("%v\n", c[i].Name)
  }
}
