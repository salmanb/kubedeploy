package container

import (
  "github.com/BurntSushi/toml"
  apiv1 "k8s.io/api/core/v1"
  "strings"
)

type container struct {
  Image   string `toml:"image"`
  Tag     string `toml:"tag"`
  Name    string `toml:"name"`
  Port    []map[string]interface{}
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
  clist := make([]apiv1.Container, 0)
  for i, _ := range c.Container {
    ports := make([]apiv1.ContainerPort, 0)
    for j, _ := range c.Container[i].Port {
      protocol := apiv1.ProtocolTCP
      if ( strings.ToLower(c.Container[i].Port[j]["protocol"].(string)) == "udp" ) {
        protocol = apiv1.ProtocolUDP
      } else if ( strings.ToLower(c.Container[i].Port[j]["protocol"].(string)) == "sctp" ) {
        protocol = apiv1.ProtocolSCTP
      }
      port := apiv1.ContainerPort{
        Name: c.Container[i].Port[j]["name"].(string),
        ContainerPort: int32(c.Container[i].Port[j]["portnum"].(int64)),
        Protocol: protocol,
      }
      ports = append(ports, port)
    }
    container := apiv1.Container{
      Name:  c.Container[i].Name,
      Image: c.Container[i].Image + ":" + c.Container[i].Tag,
      Ports: ports,
    }
    clist = append(clist, container)
  }

  return clist, nil
}
