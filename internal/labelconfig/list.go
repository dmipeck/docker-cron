package labelconfig

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// GetEnabledContainers returns a list of all containers labeled for discovery
func (ns Namespace) GetEnabledContainers(
	ctx context.Context,
	cli *client.Client,
) ([]types.Container, error) {
	filter := filters.NewArgs()
	filter.Add("label", ns.LabelKeyValue("enabled", "true"))

	return cli.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filter,
	})
}
