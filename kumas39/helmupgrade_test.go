package kumas39

import (
	"os"
	"testing"

	helm "github.com/gruntwork-io/terratest/modules/helm"
	k8s "github.com/gruntwork-io/terratest/modules/k8s"
)

func TestHelmUpgrade(t *testing.T) {

	opt := k8s.KubectlOptions{ContextName: os.Getenv("ContextName"), ConfigPath: os.Getenv("ConfigPath"), Namespace: os.Getenv("Namespace")}
	helmoptions := helm.Options{KubectlOptions: &opt}
	Upgrade(t, &helmoptions, os.Getenv("ChartPath"), os.Getenv("ChartName"))

}
