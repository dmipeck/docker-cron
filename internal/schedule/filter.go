package schedule

import (
	"errors"
	"fmt"
	"time"

	"github.com/adhocore/gronx"
	"github.com/dmipeck/docker-cron/internal/labelconfig"
	"github.com/docker/docker/api/types"
)

var ErrScheduleInvalidFormat = errors.New("invalid schedule format")

var gron = gronx.New()

// FilterScheduleIsValid returns a list of containers in "containers" which have a valid "schedule"
// label
func FilterScheduleIsValid(
	labelNamespace labelconfig.Namespace,
	containers []types.Container,
) []types.Container {
	filteredContainers := []types.Container{}

	for i := range containers {
		scheduleStr, err := GetContainerSchedule(labelNamespace, &containers[i])
		if err != nil {
			fmt.Print(errContainer(
				&containers[i],
				err,
			))
			continue
		}

		if !gron.IsValid(scheduleStr) {
			fmt.Print(errContainer(
				&containers[i],
				errScheduleStr(
					scheduleStr,
					ErrScheduleInvalidFormat,
				),
			))
			continue
		}

		filteredContainers = append(filteredContainers, containers[i])
	}

	return filteredContainers
}

// FilterScheduleIsDue returns a list of containers in "containers" which are schedule to run at the
// given timestamp.
//
// Validity of schedule labels should be checked before running this. see [FilterScheduleIsValid]
func FilterScheduleIsDue(
	labelNamespace labelconfig.Namespace,
	containers []types.Container,
	timestamp time.Time,
) ([]types.Container, error) {
	filteredContainers := []types.Container{}

	for i := range containers {
		scheduleStr, err := GetContainerSchedule(labelNamespace, &containers[i])
		if err != nil {
			return nil, errContainer(
				&containers[i],
				err,
			)
		}

		isDue, err := gron.IsDue(scheduleStr, timestamp)
		if err != nil {
			return nil, errContainer(
				&containers[i],
				errScheduleStr(
					scheduleStr,
					err,
				),
			)
		}

		if !isDue {
			continue
		}

		filteredContainers = append(filteredContainers, containers[i])
	}

	return filteredContainers, nil
}

func errContainer(container *types.Container, err error) error {
	return fmt.Errorf("container %s: %w", container.ID[:12], err)
}

func errScheduleStr(scheduleStr string, err error) error {
	return fmt.Errorf("schedule \"%s\": %w", scheduleStr, err)
}
