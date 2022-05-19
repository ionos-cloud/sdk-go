package main

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"log"
	"net/http"
	"sync"
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

	ch := make(chan ionoscloud.DeleteStateChannel)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		apiClient.WaitForDeletionAsync(context.Background(), processRequestClusterDelete, resourceId, ch)
		for res := range ch {
			if res.Err != nil {
				log.Fatal(res.Err)
			}
			if res.Msg == http.StatusNotFound { // if statement will be true when the resource was deleted
				// put your code here(for example you can delete a k8s nodepool and after that a cluster)
				wg.Done()
			}
		}
	}()

	wg.Wait()
}
