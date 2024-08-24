set quiet

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

