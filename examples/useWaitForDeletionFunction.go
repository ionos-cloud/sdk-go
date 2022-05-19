package main

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"log"
)

func processRequestClusterDelete(apiClient *ionoscloud.APIClient, resourceID string) (*ionoscloud.APIResponse, error) {
	apiRequest := apiClient.KubernetesApi.K8sFindByClusterId(context.Background(), resourceID)
	_, apiResp, err := apiRequest.Execute()
	if err != nil {
		return apiResp, fmt.Errorf("error occurred when executing the api get resource operation: %w", err)
	}

	return apiResp, nil
}

func main() {
	cfg := ionoscloud.NewConfigurationFromEnv()
	apiClient := ionoscloud.NewAPIClient(cfg)
	resourceId := "c376e371-51bb-45f2-9b24-3da0fc3a6cfc" // you can pass the resource id after you create a resource using resource.GetId()

	deleted, err := apiClient.WaitForDeletion(context.Background(), processRequestClusterDelete, resourceId)
	if err != nil {
		log.Fatal(err)
	}
	if deleted == true {
		// put your code here(for example you can delete a k8s nodepool and after that a cluster)
	}
}
