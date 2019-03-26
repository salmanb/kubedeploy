package main

import (
  "fmt"
  "github.com/salmanb/kubedeploy/parts/label"
  "github.com/salmanb/kubedeploy/parts/container"
  "io/ioutil"
  "log"
  "os"
)

func main() {
  ioreader, _ := os.Open("dply.toml")
  data, err := ioutil.ReadAll(ioreader)
  if (err != nil) {
    log.Fatal(err)
  }

a, b := label.MakeLabelMap(string(data))
if (b != nil) {
  log.Fatal(b)
}
for k, v := range(a) {
  fmt.Printf("%v => %v\n", k, v)
  fmt.Println("")
}

 c, err := container.MakeList(string(data))
 if (err != nil) {
   log.Fatal(err)
 }
 for i, _ := range c {
   fmt.Printf("Image: %v\n", c[i].Image)
   fmt.Printf("Image: %T\n", c[i].Image)
   fmt.Println("")
   fmt.Printf("Name: %v\n", c[i].Name)
   fmt.Printf("Name: %T\n", c[i].Name)
   fmt.Println("")
   fmt.Printf("Command: %v\n", c[i].Command)
   fmt.Printf("Command: %T\n", c[i].Command)
   fmt.Println("")
   fmt.Printf("Ports: %v\n", c[i].Ports)
   fmt.Printf("Ports: %T\n", c[i].Ports)
   fmt.Println("")
   fmt.Printf("Env: %v\n", c[i].Env)
   fmt.Printf("Env: %T\n", c[i].Env)
   fmt.Println("")
   fmt.Printf("Resources: %v\n", c[i].Resources)
   fmt.Printf("Resources: %T\n", c[i].Resources)
   fmt.Println("")
 }

}
