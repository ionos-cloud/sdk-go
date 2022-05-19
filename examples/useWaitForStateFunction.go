package main

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"log"
)

func processRequestCluster(apiClient *ionoscloud.APIClient, resourceID string) (ionoscloud.ResourceHandler, error) {
	apiRequest := apiClient.KubernetesApi.K8sFindByClusterId(context.Background(), resourceID)
	dc, _, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}
	return &dc, nil
}

func main() {
	cfg := ionoscloud.NewConfigurationFromEnv()
	apiClient := ionoscloud.NewAPIClient(cfg)
	resourceId := "c376e371-51bb-45f2-9b24-3da0fc3a6cfc"

	active, err := apiClient.WaitForState(context.Background(), processRequestCluster, resourceId)
	if err != nil {
		log.Fatal(err)
	}
	if active == true {
		// put your code here(for example you can create a k8s nodepool after a cluster is created)
	}
}
