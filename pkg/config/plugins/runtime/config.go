package runtime

import (
	"github.com/pkg/errors"

	"github.com/Kong/kuma/pkg/config/core"
	"github.com/Kong/kuma/pkg/config/plugins/runtime/k8s"
)

func DefaultRuntimeConfig() *RuntimeConfig {
	return &RuntimeConfig{
		Kubernetes: k8s.DefaultKubernetesRuntimeConfig(),
	}
}

// Environment-specific configuration
type RuntimeConfig struct {
	// Kubernetes-specific configuration
	Kubernetes *k8s.KubernetesRuntimeConfig `yaml:"kubernetes"`
}

func (c *RuntimeConfig) Validate(env core.EnvironmentType) error {
	switch env {
	case core.KubernetesEnvironment:
		if err := c.Kubernetes.Validate(); err != nil {
			return errors.Wrap(err, "Kubernetes validation failed")
		}
	case core.UniversalEnvironment:
	default:
		return errors.Errorf("unknown environment type %q", env)
	}
	return nil
}
