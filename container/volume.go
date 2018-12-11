package container

import (
	"context"
)

// DeleteVolume deletes a Docker Volume by name.
func (c *DockerContainer) DeleteVolume(mount Mount) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.callTimeout)
	defer cancel()
	return c.client.VolumeRemove(ctx, mount.source(c), false)
}
