package container

import (
	"strings"

	"github.com/mesg-foundation/core/x/xstructhash"
)

const namespaceSeparator string = "-"

// CleanNamespace creates a namespace from a list of string.
func (c *DockerContainer) CleanNamespace(ss []string) string {
	ssWithPrefix := append([]string{c.config.Core.Name}, ss...)
	namespace := strings.Join(ssWithPrefix, namespaceSeparator)
	namespace = strings.Replace(namespace, " ", namespaceSeparator, -1)
	return namespace
}

// HashNamespace creates a namespace from a list of string.
func (c *DockerContainer) HashNamespace(ss []string) string {
	ssWithPrefix := append([]string{c.config.Core.Name}, ss...)
	return xstructhash.Hash(ssWithPrefix, 1)
}
