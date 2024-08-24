set quiet

image_name := 'ghcr.io/dmipeck/docker-cron'
image_tag := 'latest'

coverage_file := '.go-coverage.txt'

[private]
default:
    just --list

# run the docker-cron cli
run:
    go run cli/main.go

# run tests for all packages in the docker-cron module
test:
    go clean -testcache
    go test ./...

# run tests for all packages in the docker-cron module and reports coverage results
test_coverage:
    go clean -testcache
    go test -coverprofile='{{coverage_file}}' ./...
    echo ''
    go tool cover -func='{{coverage_file}}'

# create the docker-cron container image
docker_build:
    docker buildx build --tag {{image_name}}:{{image_tag}} {{justfile_directory()}}

# create and run the docker-cron container
docker_run:
    docker run \
        --rm \
        --privileged \
        --mount=type=bind,source="/var/run/docker.sock",destination="/var/run/docker.sock" \
        {{image_name}}:{{image_tag}}

# create a scheduled container that logs a timestamp to "logging.log"
docker_run_logger schedule='* * * * *' delay='0':
    touch logger.log
    docker container create \
        --label="github.com/dmipeck/docker-cron.enabled=true" \
        --label="github.com/dmipeck/docker-cron.schedule={{schedule}}" \
        --mount=type=bind,source="{{justfile_directory()}}/logger.log",destination="/var/log/logger.log" \
        alpine:3.20 \
        ash -c 'sleep "{{delay}}" && printf "%s %s\n" "$(cat /etc/hostname)" "$(date +"%F %T")" >> /var/log/logger.log'
