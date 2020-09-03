package main

import (
	"fmt"
	"octavius/internal/pkg/kubernetes"
	"octavius/internal/pkg/kubernetes/config"
)

func main() {
	var kubeConfig = config.KubernetesConfig{
		KubeConfig:                   config.Config().KubeConfig,
		KubeContext:                  config.Config().KubeContext,
		DefaultNamespace:             config.Config().DefaultNamespace,
		KubeServiceAccountName:       config.Config().KubeServiceAccountName,
		JobPodAnnotations:            config.Config().JobPodAnnotations,
		KubeJobActiveDeadlineSeconds: config.Config().KubeJobActiveDeadlineSeconds,
		KubeJobRetries:               config.Config().KubeJobRetries,
		KubeWaitForResourcePollCount: config.Config().KubeWaitForResourcePollCount,
	}

	var newClient, err = kubernetes.NewKubernetesClient(kubeConfig)

	if err != nil {
		fmt.Println("ERROR! Unable to create new clientSet: ", err)
	} else {
		fmt.Printf("new Kube client created: %+v\n", newClient)
	}

	env := map[string]string{
		"name": "akshay",
	}

	executionName, _ := newClient.ExecuteJob("jaiminrathod98765/i1", env)
	fmt.Println(executionName)
	fmt.Println(newClient.JobExecutionStatus(executionName))

}
