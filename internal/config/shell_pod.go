package config

import (
	"github.com/openqt/osc/internal/client"
	v1 "k8s.io/api/core/v1"
)

const defaultDockerShellImage = "busybox:1.31"

// Limits represents resource limits.
type Limits map[v1.ResourceName]string

// ShellPod represents k9s shell configuration.
type ShellPod struct {
	Image     string `json:"Image"`
	Namespace string `json:"namespace"`
	Limits    Limits `json:"resources,omitempty"`
}

// NewShellPod returns a new instance.
func NewShellPod() *ShellPod {
	return &ShellPod{
		Image:     defaultDockerShellImage,
		Namespace: "default",
		Limits:    defaultLimits(),
	}
}

// Validate validates the configuration.
func (s *ShellPod) Validate(client.Connection, KubeSettings) {
	if s.Image == "" {
		s.Image = defaultDockerShellImage
	}
	if len(s.Limits) == 0 {
		s.Limits = defaultLimits()
	}
}

func defaultLimits() Limits {
	return Limits{
		v1.ResourceCPU:    "100m",
		v1.ResourceMemory: "100Mi",
	}
}
