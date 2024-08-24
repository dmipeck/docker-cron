package main

import (
	"context"
	"fmt"

	"github.com/adhocore/gronx/pkg/tasker"
	"github.com/dmipeck/docker-cron/internal/labelconfig"
	"github.com/dmipeck/docker-cron/internal/schedule"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	namespace := labelconfig.Namespace("github.com/dmipeck/docker-cron")
	fmt.Printf("running docker-cron in namespace \"%s\"\n", namespace)

	taskr := tasker.New(tasker.Option{Verbose: true})
	taskr.Task("* * * * *", scheduleTask(namespace, cli))
	taskr.Run()
}

func scheduleTask(namespace labelconfig.Namespace, cli *client.Client) tasker.TaskFunc {
	return func(ctx context.Context) (int, error) {
		return schedule.StartScheduledContainers(ctx, namespace, cli)
	}
}
