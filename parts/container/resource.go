package container

//  [container.resource.request]
//    cpu = "500m"
//   memory = "1Gi"
//    storage = "5Gi"
//    ephemeralstorage = "10Gi"
//  [container.resource.limit]
//    cpu = "1024m"
//    memory = "2Gi"
//    storage = "10Gi"
//    ephemeralstorage = "20Gi"

import (
  apiv1 "k8s.io/api/core/v1"
  "k8s.io/apimachinery/pkg/api/resource"
)

type RsRequest struct {
  Cpu string `toml:"cpu,omitempty"`
  Memory  string `toml:"memory,omitempty"`
  Storage string  `toml:"storage,omitempty"`
  EphemeralStorage  string `toml:"ephemeralstorage,omitempty"`
}

type RsLimit struct {
  Cpu string `toml:"cpu,omitempty"`
  Memory  string `toml:"memory,omitempty"`
  Storage string  `toml:"storage,omitempty"`
  EphemeralStorage  string `toml:"ephemeralstorage,omitempty"`
}

type RequestedResource struct {
  Requested RsRequest `toml:"request,omitempty"`
  Limited RsLimit `toml:"limit,omitempty"`
}

func (c *Container) MakeResourceMapping(_container *container) apiv1.ResourceRequirements {
  rreqs := apiv1.ResourceRequirements{}
  rlimitlist := make(map[apiv1.ResourceName]resource.Quantity)
  rreqlist := make(map[apiv1.ResourceName]resource.Quantity)

  //requested resources
  if (_container.Resources.Requested.Cpu == "") {
    rreqlist[apiv1.ResourceCPU] = resource.MustParse("500m")
  } else {
    rreqlist[apiv1.ResourceCPU] = resource.MustParse(_container.Resources.Requested.Cpu)
  }

  if (_container.Resources.Requested.Memory == "") {
    rreqlist[apiv1.ResourceCPU] = resource.MustParse("500m")
  } else {
    rreqlist[apiv1.ResourceMemory] = resource.MustParse(_container.Resources.Requested.Memory)
  }

  if (_container.Resources.Requested.Storage == "") {
    rreqlist[apiv1.ResourceCPU] = resource.MustParse("1Gi")
  } else {
    rreqlist[apiv1.ResourceStorage] = resource.MustParse(_container.Resources.Requested.Storage)
  }

  if (_container.Resources.Requested.EphemeralStorage == "") {
    rreqlist[apiv1.ResourceEphemeralStorage] = resource.MustParse("1Gi")
  } else {
    rreqlist[apiv1.ResourceEphemeralStorage] = resource.MustParse( _container.Resources.Requested.EphemeralStorage)
  }

  // set limits on resources
  if (_container.Resources.Limited.Cpu == "") {
    rlimitlist[apiv1.ResourceCPU] = resource.MustParse("500m")
  } else {
    rreqlist[apiv1.ResourceCPU] = resource.MustParse(_container.Resources.Limited.Cpu)
  }

  if (_container.Resources.Limited.Memory == "") {
    rlimitlist[apiv1.ResourceCPU] = resource.MustParse("500m")
  } else {
    rlimitlist[apiv1.ResourceMemory] = resource.MustParse(_container.Resources.Limited.Memory)
  }

  if (_container.Resources.Limited.Storage == "") {
    rlimitlist[apiv1.ResourceCPU] = resource.MustParse("1Gi")
  } else {
    rlimitlist[apiv1.ResourceStorage] = resource.MustParse(_container.Resources.Limited.Storage)
  }

  if (_container.Resources.Limited.EphemeralStorage == "") {
    rlimitlist[apiv1.ResourceEphemeralStorage] = resource.MustParse("1Gi")
  } else {
    rlimitlist[apiv1.ResourceEphemeralStorage] = resource.MustParse( _container.Resources.Limited.EphemeralStorage)
  }

  rreqs.Limits = rlimitlist

  return rreqs
}
