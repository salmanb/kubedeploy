package container

import (
  "fmt"
  "github.com/BurntSushi/toml"
  apiv1 "k8s.io/api/core/v1"
  "strings"
)

type envFrom struct {
  Type  string `toml:"type,omitempty"`
  Name  string `toml:"name,omitempty"`
  Key   string `toml:"key,omitempty"`
}
type env struct {
  Name  string
  Value string
  EnvFrom envFrom `toml:"from,omitempty"`
}

type container struct {
  Image   string `toml:"image,omitempty"`
  Tag     string `toml:"tag,omitempty"`
  Name    string `toml:"name,omitempty"`
  Command string  `toml:"command,omitempty"`
  Port    []map[string]interface{} `toml:"port,omitempty"`
  EnvVar  []env `toml:"env,omitempty"`
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
    ports := c.makePortMapping(&c.Container[i])
    envVars := c.makeEnvMapping(&c.Container[i])
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

func (c *Container) makeEnvMapping(_container *container) []apiv1.EnvVar {
  envVars := make([]apiv1.EnvVar, 0)
  for i, _ := range _container.EnvVar {
    fmt.Printf("+++ EnvVar Name: %v +++\n", _container.EnvVar[i].Name)
    fmt.Printf("+++ EnvVar Value: %v +++\n", _container.EnvVar[i].Value)
    fmt.Printf("+++ EnvFrom Name: %v +++\n", _container.EnvVar[i].EnvFrom.Name)
    fmt.Printf("+++ EnvFrom Key: %v +++\n", _container.EnvVar[i].EnvFrom.Key)
    fmt.Printf("+++ EnvFrom Type: %v +++\n", _container.EnvVar[i].EnvFrom.Type)
    envVar := apiv1.EnvVar{
      Name: _container.EnvVar[i].Name,
      Value: _container.EnvVar[i].Value,
      ValueFrom: c.makeEnvVarSourceMapping(_container.EnvVar[i].EnvFrom.Type, _container.EnvVar[i].EnvFrom.Name, _container.EnvVar[i].EnvFrom.Key),
    }
    envVars = append(envVars, envVar)
  }
  return envVars
}

func (c *Container) makeEnvVarSourceMapping(envVarType, envVarName, envVarKey string) *apiv1.EnvVarSource {
  evs := &apiv1.EnvVarSource{}
    lor := apiv1.LocalObjectReference{
      Name: envVarName,
    }
  if (envVarType == "configmap") {
    cmks := &apiv1.ConfigMapKeySelector{}
    lor := apiv1.LocalObjectReference{
      Name: envVarName,
    }
    cmks.LocalObjectReference = lor
    cmks.Key = envVarKey
    evs.ConfigMapKeyRef = cmks
  } else if (envVarType == "secret") {
    sks := &apiv1.SecretKeySelector{}
    sks.LocalObjectReference = lor
    sks.Key = envVarKey
    evs.SecretKeyRef = sks
  }

  return evs
}

func (c *Container) makePortMapping(_container *container) []apiv1.ContainerPort {
  ports := make([]apiv1.ContainerPort, 0)
    for i, _ := range _container.Port {
      protocol := apiv1.ProtocolTCP
      if ( strings.ToLower(_container.Port[i]["protocol"].(string)) == "udp" ) {
        protocol = apiv1.ProtocolUDP
      } else if ( strings.ToLower(_container.Port[i]["protocol"].(string)) == "sctp" ) {
        protocol = apiv1.ProtocolSCTP
      }
      port := apiv1.ContainerPort{
        Name: _container.Port[i]["name"].(string),
        ContainerPort: int32(_container.Port[i]["portnum"].(int64)),
        Protocol: protocol,
      }
      ports = append(ports, port)
   }
   return ports
}
