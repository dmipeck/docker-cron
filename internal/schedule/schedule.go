package schedule

import (
	"errors"

	"github.com/dmipeck/docker-cron/internal/labelconfig"
	"github.com/docker/docker/api/types"
)

var ErrScheduleLabelNotFound = errors.New("schedule label not found")

// GetContainerSchedule returns the cron schedule string for "container"
// e.g. "* * * * *"
func GetContainerSchedule(labelNamespace labelconfig.Namespace, container *types.Container) (string, error) {
	scheduleStr, ok := container.Labels[labelNamespace.LabelKey("schedule")]

	if !ok {
		return "", ErrScheduleLabelNotFound
	}

	return scheduleStr, nil
}
