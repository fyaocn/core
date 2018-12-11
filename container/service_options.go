package container

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
)

// ServiceOptions is a simplify version of swarm.ServiceSpec.
type ServiceOptions struct {
	Image     string
	Namespace []string
	Ports     []Port
	Mounts    []Mount
	Env       []string // TODO: should be transform to  map[string]string and use the func mapToEnv
	Args      []string
	Command   string
	Networks  []Network
	Labels    map[string]string
}

// Network keeps the network info for service.
type Network struct {
	// ID of the docker network.
	ID string

	// Alias is an optional attribute to name this service in the
	// network and be able to access to it using this name.
	Alias string
}

// Port is a simplify version of swarm.PortConfig.
type Port struct {
	Target    uint32
	Published uint32
}

// Mount is a generic interface of mount.Mount.
type Mount interface {
	swarmMounts(*DockerContainer) mount.Mount
	source(*DockerContainer) string
}

// Bind is a bind version of mount.Mount.
type Bind struct {
	Source string
	Target string
}

func (m *Bind) swarmMounts(_ *DockerContainer) mount.Mount {
	return mount.Mount{
		Source: m.Source,
		Target: m.Target,
		Type:   mount.TypeBind,
	}
}

func (m *Bind) source(_ *DockerContainer) string {
	return m.Source
}

// Volume is a volume version of mount.Mount.
type Volume struct {
	Namespace []string
	Target    string
}

func (m *Volume) swarmMounts(c *DockerContainer) mount.Mount {
	return mount.Mount{
		Source: m.source(c),
		Target: m.Target,
		Type:   mount.TypeVolume,
	}
}

func (m *Volume) source(c *DockerContainer) string {
	return c.HashNamespace(m.Namespace)
}

func (options *ServiceOptions) toSwarmServiceSpec(c *DockerContainer) swarm.ServiceSpec {
	namespace := c.Namespace(options.Namespace)
	return swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: c.HashNamespace(options.Namespace),
			Labels: mergeLabels(options.Labels, map[string]string{
				"com.docker.stack.namespace": namespace,
				"com.docker.stack.image":     options.Image,
			}),
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: &swarm.ContainerSpec{
				Image: options.Image,
				Labels: map[string]string{
					"com.docker.stack.namespace": namespace,
				},
				Env:     options.Env,
				Args:    options.Args,
				Command: strings.Fields(options.Command),
				Mounts:  options.swarmMounts(c, false),
			},
			Networks: options.swarmNetworks(),
		},
		EndpointSpec: &swarm.EndpointSpec{
			Ports: options.swarmPorts(),
		},
	}
}

func (options *ServiceOptions) swarmPorts() []swarm.PortConfig {
	ports := make([]swarm.PortConfig, len(options.Ports))
	for i, p := range options.Ports {
		ports[i] = swarm.PortConfig{
			Protocol:      swarm.PortConfigProtocolTCP,
			PublishMode:   swarm.PortConfigPublishModeIngress,
			TargetPort:    p.Target,
			PublishedPort: p.Published,
		}
	}
	return ports
}

func (options *ServiceOptions) swarmMounts(c *DockerContainer, force bool) []mount.Mount {
	// TOFIX: hack to prevent mount when in CircleCI (Mount in CircleCI doesn't work). Should use CircleCi with machine to fix this.
	circleCI, errCircle := strconv.ParseBool(os.Getenv("CIRCLECI"))
	if !force && errCircle == nil && circleCI {
		return nil
	}
	mounts := make([]mount.Mount, len(options.Mounts))
	for i, m := range options.Mounts {
		mounts[i] = m.swarmMounts(c)
	}
	return mounts
}

// swarmNetworks creates all necessary network attachment configurations for service.
// each network will be attached based on their networkID and an alias can be used to
// identify service in the network.
// aliases will make services accessible from other containers inside the same network.
func (options *ServiceOptions) swarmNetworks() (networks []swarm.NetworkAttachmentConfig) {
	networks = make([]swarm.NetworkAttachmentConfig, len(options.Networks))
	for i, network := range options.Networks {
		cfg := swarm.NetworkAttachmentConfig{
			Target: network.ID,
		}
		if network.Alias != "" {
			cfg.Aliases = []string{network.Alias}
		}
		networks[i] = cfg
	}
	return networks
}

func mergeLabels(l1 map[string]string, l2 map[string]string) map[string]string {
	if l1 == nil {
		l1 = make(map[string]string)
	}
	for k, v := range l2 {
		l1[k] = v
	}
	return l1
}

// MapToEnv transform a map of key value to a array of env string.
// env vars sorted by names to get an accurate order while testing, otherwise
// comparing a string slice with different orders will fail.
func MapToEnv(data map[string]string) []string {
	env := make([]string, 0, len(data))
	for key, value := range data {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	sort.Strings(env)
	return env
}
