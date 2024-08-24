package labelconfig_test

import (
	"testing"

	"github.com/dmipeck/docker-cron/internal/labelconfig"
)

func TestLabelKey(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		testCases := []struct {
			namespace labelconfig.Namespace
			key       string
			expected  string
		}{
			{
				namespace: labelconfig.Namespace("my-new-namespace"),
				key:       "enabled",
				expected:  "my-new-namespace.enabled",
			},
		}

		for i := range testCases {
			label := testCases[i].namespace.LabelKey(testCases[i].key)
			if label != testCases[i].expected {
				t.Errorf(
					"expected label string \"%s\", got: \"%s\"",
					testCases[i].expected,
					label,
				)
			}
		}
	})
}
