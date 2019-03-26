package container

import (
  apiv1 "k8s.io/api/core/v1"
)

type EnvFrom struct {
  Type  string `toml:"type,omitempty"`
  Name  string `toml:"name,omitempty"`
  Key   string `toml:"key,omitempty"`
}
type Env struct {
  Name  string
  Value string
  EnvFrom EnvFrom `toml:"from,omitempty"`
}

func (c *Container) MakeEnvMapping(_container *container) []apiv1.EnvVar {
  envVars := make([]apiv1.EnvVar, 0)
  for i, _ := range _container.EnvVar {
    envVar := apiv1.EnvVar{
      Name: _container.EnvVar[i].Name,
      Value: _container.EnvVar[i].Value,
      ValueFrom: c.MakeEnvVarSourceMapping(_container.EnvVar[i].EnvFrom.Type, _container.EnvVar[i].EnvFrom.Name, _container.EnvVar[i].EnvFrom.Key),
    }
    envVars = append(envVars, envVar)
  }
  return envVars
}

func (c *Container) MakeEnvVarSourceMapping(envVarType, envVarName, envVarKey string) *apiv1.EnvVarSource {
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
