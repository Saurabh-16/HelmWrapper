package kumas39

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/gruntwork-cli/errors"
	"github.com/gruntwork-io/terratest/modules/files"
	helm "github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
)

// Upgrade will upgrade the selected helm chart with the provided options under the given release name. This will fail
// the test if there is an error.
func Upgrade(t *testing.T, options *helm.Options, chart string, releaseName string) {
	require.NoError(t, UpgradeE(t, options, chart, releaseName))
}

// UpgradeE will upgrade the selected helm chart with the provided options under the given release name.
func UpgradeE(t *testing.T, options *helm.Options, chart string, releaseName string) error {
	// If the chart refers to a path, convert to absolute path. Otherwise, pass straight through as it may be a remote
	// chart.
	if files.FileExists(chart) {
		absChartDir, err := filepath.Abs(chart)
		if err != nil {
			return errors.WithStackTrace(err)
		}
		chart = absChartDir
	}

	var err error
	args := []string{}
	args, err = getValuesArgsE(t, options, args...)
	if err != nil {
		return err
	}

	args = append(args, releaseName, chart)
	_, err = helm.RunHelmCommandAndGetOutputE(t, options, "upgrade", args...)
	return err
}
