# docker-cron

docker-cron is a tool for scheduling containers execution using labels

## Getting started

### 1. Start the docker-cron container

You can run the docker-cron container with the following:
```sh
docker run \
    --privileged \
    --restart="unless-stopped" \
    --mount=type=bind,source="/var/run/docker.sock",destination="/var/run/docker.sock" \
    "ghcr.io/dmipeck/docker-cron:latest"
```

The docker-cron container will check each minute for any containers that should be run based of of a cron expression added to the container labels

### 2. Create your scheduled container

You can run an example scheduled container with the following:
```sh
docker container create \
    --label="github.com/dmipeck/docker-cron.enabled=true" \
    --label="github.com/dmipeck/docker-cron.schedule=* * * * *" \
    "hello_world:latest"
```

This will run the `hello-world` container every minute

## Configuration

### Scheduled container labels

Scheduled containers require the following labels to be started by docker-cron:

| Label | Value |
| ----- | ----- |
| `github.com/dmipeck/docker-cron.enabled` | `"true"`|
| `github.com/dmipeck/docker-cron.schedule` | any valid cron expression, e.g `"* * * * *"` |

### Restarting scheduled containers that fail

docker-cron will only try to start the container once, regardless of if it succeeds. Use the [`--restart=on-failed`](https://docs.docker.com/reference/cli/docker/container/run/#restart), or [`--restart=on-failed:<max_retries>`](https://docs.docker.com/reference/cli/docker/container/run/#restart) flags to restart the container if there is an error.
