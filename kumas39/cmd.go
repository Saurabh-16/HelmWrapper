package kumas39

import (
	"testing"

	"github.com/gruntwork-io/gruntwork-cli/errors"
	helm "github.com/gruntwork-io/terratest/modules/helm"
)

// getValuesArgsE computes the args to pass in for setting values
func getValuesArgsE(t *testing.T, options *helm.Options, args ...string) ([]string, error) {
	args = append(args, formatSetValuesAsArgs(options.SetValues, "--set")...)
	args = append(args, formatSetValuesAsArgs(options.SetStrValues, "--set-string")...)

	valuesFilesArgs, err := formatValuesFilesAsArgsE(t, options.ValuesFiles)
	if err != nil {
		return args, errors.WithStackTrace(err)
	}
	args = append(args, valuesFilesArgs...)

	setFilesArgs, err := formatSetFilesAsArgsE(t, options.SetFiles)
	if err != nil {
		return args, errors.WithStackTrace(err)
	}
	args = append(args, setFilesArgs...)
	return args, nil
}
