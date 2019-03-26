package container

import (
  apiv1 "k8s.io/api/core/v1"
  "strings"
)

func (c *Container) MakePortMapping(_container *container) []apiv1.ContainerPort {
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
