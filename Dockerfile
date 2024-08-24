# syntax=docker/dockerfile:1
FROM golang:1.23 AS build

WORKDIR /src

RUN \
    --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

FROM build AS build-cron

RUN \
    --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=.,target=. \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/bin/entrypoint cli/main.go

FROM scratch AS prod

COPY \
    --from=build-cron \
    --chown=root \
    --chmod=744 \
    /usr/bin/entrypoint /usr/bin/entrypoint

ENTRYPOINT [ "entrypoint" ]
