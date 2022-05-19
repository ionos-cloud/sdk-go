package main

import (
	"context"
	"fmt"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
	"log"
	"sync"
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
	resourceId := "c376e371-51bb-45f2-9b24-3da0fc3a6cfc" // you can pass the resource id after you create a resource using resource.GetId()

	ch := make(chan ionoscloud.StateChannel)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		apiClient.WaitForStateAsync(context.Background(), processRequestCluster, resourceId, ch)
		for res := range ch {
			if res.Err != nil {
				log.Fatal(res.Err)
			}
			if res.Msg == ionoscloud.Available || res.Msg == ionoscloud.Active {
				fmt.Printf("%v ", "DONE Creating the Cluster")
				// put your code here(for example you can create a k8s nodepool after a cluster is created)
				wg.Done()
			}
		}
	}()

	wg.Wait()
}
