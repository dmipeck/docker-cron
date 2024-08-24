package schedule_test

import (
	"testing"

	"github.com/dmipeck/docker-cron/internal/labelconfig"
	"github.com/dmipeck/docker-cron/internal/schedule"
	"github.com/docker/docker/api/types"
)

func TestGetContainerSchedule(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		testCases := []struct {
			container *types.Container
			namespace labelconfig.Namespace
			expected  string
		}{
			{
				container: &types.Container{Labels: map[string]string{
					"my-namespace.schedule": "* * * * *",
				}},
				namespace: "my-namespace",
				expected:  "* * * * *",
			},
		}

		for i := range testCases {
			scheduleStr, err := schedule.GetContainerSchedule(
				testCases[i].namespace,
				testCases[i].container,
			)
			if err != nil {
				t.Errorf("expected no error, got :%s", err.Error())
			}
			if scheduleStr != testCases[i].expected {
				t.Errorf(
					"expected schedule string \"%s\", got: \"%s\"",
					testCases[i].expected,
					scheduleStr,
				)
			}
		}
	})

	t.Run("Error", func(t *testing.T) {
		testCases := []struct {
			container *types.Container
			namespace labelconfig.Namespace
			expected  error
		}{
			{
				container: &types.Container{Labels: map[string]string{
					"my-namespace.bad-value": "* * * * *",
				}},
				namespace: "my-namespace",
				expected:  schedule.ErrScheduleLabelNotFound,
			},
			{
				container: &types.Container{Labels: map[string]string{
					"not-my-namespace.schedule": "* * * * *",
				}},
				namespace: "my-namespace",
				expected:  schedule.ErrScheduleLabelNotFound,
			},
		}

		for i := range testCases {
			_, err := schedule.GetContainerSchedule(
				testCases[i].namespace,
				testCases[i].container,
			)
			if err != testCases[i].expected {
				t.Errorf(
					"expected error \"%s\", got: \"%s\"",
					testCases[i].expected.Error(),
					err.Error(),
				)
			}
		}
	})
}
