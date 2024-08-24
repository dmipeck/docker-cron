package schedule

import (
	"context"
	"fmt"
	"time"

	"github.com/dmipeck/docker-cron/internal/labelconfig"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const TimeFormat = "2006-01-02 15:04:05"

func StartScheduledContainers(
	ctx context.Context,
	labelNamespace labelconfig.Namespace,
	cli *client.Client,
) (int, error) {
	timestamp := time.Now()
	containers, err := labelNamespace.GetEnabledContainers(ctx, cli)
	if err != nil {
		return 0, err
	}

	enabledCount := len(containers)
	fmt.Printf("found %d containers total\n", enabledCount)

	if enabledCount == 0 {
		fmt.Printf("skipping: no enabled containers\n")
		return 0, nil
	}

	containers = FilterScheduleIsValid(labelNamespace, containers)

	validCount := len(containers)
	fmt.Printf("found %d containers with invalid schedules\n", validCount)

	if validCount == 0 {
		fmt.Printf("skipping: no valid containers\n")
		return 0, nil
	}

	containers, err = FilterScheduleIsDue(labelNamespace, containers, timestamp)
	if err != nil {
		return 0, err
	}
	totalCount := len(containers)

	fmt.Printf(
		"found %d containers scheduled for %s\n",
		totalCount,
		timestamp.Format(TimeFormat),
	)

	if totalCount == 0 {
		fmt.Printf("skipping: no scheduled containers\n")
		return 0, nil
	}

	for i := range containers {
		fmt.Printf("starting container %s\n", containers[i].ID[:12])

		err := cli.ContainerStart(ctx, containers[i].ID, container.StartOptions{})
		if err != nil {
			fmt.Print(fmt.Errorf("starting container %s: %w", containers[i].ID[:12], err))
			continue
		}
	}

	return 0, err
}
