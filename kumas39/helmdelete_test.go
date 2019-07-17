package kumas39

import (
	"os"
	"testing"

	helm "github.com/gruntwork-io/terratest/modules/helm"
	k8s "github.com/gruntwork-io/terratest/modules/k8s"
)

func TestHelmDelete(t *testing.T) {

	opt := k8s.KubectlOptions{ContextName: os.Getenv("ContextName"), ConfigPath: os.Getenv("ConfigPath"), Namespace: os.Getenv("Namespace")}
	helmopt := helm.Options{KubectlOptions: &opt}
	helm.Delete(t, &helmopt, os.Getenv("ChartName"), true)
}
