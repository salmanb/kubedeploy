package container

import (
  "fmt"
  "github.com/BurntSushi/toml"
  apiv1 "k8s.io/api/core/v1"
  "strings"
)

type container struct {
  Image   string `toml:"image,omitempty"`
  Tag     string `toml:"tag,omitempty"`
  Name    string `toml:"name,omitempty"`
  Command string  `toml:"command,omitempty"`
  Port    []map[string]interface{} `toml:"port,omitempty"`
  EnvVar  []Env `toml:"env,omitempty"`
  Resources RequestedResource `toml:"resource,omitempty"`
}

type Container struct {
  Container []container `toml:"container"`
}

func MakeList(s string) ([]apiv1.Container, error) {
  var c Container
  _, err := toml.Decode(s, &c)
  if (err != nil) {
    return nil, err
  }
  fmt.Printf(" ===== %v =====\n", c)
  return c.makeContainerList(), nil
}

func (c *Container) makeContainerList() []apiv1.Container {
  clist := make([]apiv1.Container, 0)
  for i, _ := range c.Container {
    ports := c.MakePortMapping(&c.Container[i])
    envVars := c.MakeEnvMapping(&c.Container[i])
    rreqs := c.MakeResourceMapping(&c.Container[i])
    fmt.Printf("--- RReqs: %v ---\n", rreqs)
    container := apiv1.Container{
      Name:  c.Container[i].Name,
      Image: c.Container[i].Image + ":" + c.Container[i].Tag,
      Ports: ports,
      Command: strings.Split(c.Container[i].Command, " "),
      Env: envVars,
    }
    clist = append(clist, container)
  }

  return clist
}

