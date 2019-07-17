package kumas39

import (
	"os"
	"testing"

	helm "github.com/gruntwork-io/terratest/modules/helm"
	k8s "github.com/gruntwork-io/terratest/modules/k8s"
)

func TestHelmInstall(t *testing.T) {

	opt := k8s.KubectlOptions{ContextName: os.Getenv("contextname"), ConfigPath: os.Getenv("configpath"), Namespace: os.Getenv("namespace")}
	helmoptions := helm.Options{KubectlOptions: &opt}
	helm.Install(t, &helmoptions, os.Getenv("ChartPath"), os.Getenv("ChartName"))

}
