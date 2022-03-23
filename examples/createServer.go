package main

import (
	"context"
	"fmt"
	"log"
	"time"

	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

// apiService - a wrapper to have access to the computeEngine client and context
type apiService struct {
	computeEngine *ionoscloud.APIClient
	ctx           context.Context
}

func main() {
	err := createDataCenterAndServer()
	if err != nil {
		log.Fatal("err : ", err)
	}
}

// createDataCenterAndServer - creates a server and datacenter and waits for the resources to finish being created
func createDataCenterAndServer() error {

	as := newApiService()

	var cancelFunc func()
	// If creation takes more than 10 minutes, cancel and release resources
	duration := 10 * time.Minute
	as.ctx, cancelFunc = context.WithTimeout(as.ctx, duration)

	// Ensure the ctx is canceled.
	// See ctx package for more information, https://golang.org/pkg/context/
	if cancelFunc != nil {
		defer cancelFunc()
	}

	dc, err := as.createDcAndWaitUntilDone()
	if err != nil {
		return fmt.Errorf(
			"error creating datacenter (%w)", err)
	}
	dcId := *dc.Id
	_, err = as.createServerAndWaitUntilDone(dcId)
	if err != nil {
		return fmt.Errorf(
			"error creating server (%w)", err)
	}
	return nil
}

// newApiService - creates apiClient(used to make requests to IONOS Cloud API) and context needed for operations
func newApiService() apiService {

	// NewConfigurationFromEnv looks for the following variables in the environment: IONOS_USERNAME, IONOS_PASSWORD, IONOS_TOKEN, IONOS_API_URL
	// You can either export IONOS_USERNAME and IONOS_PASSWORD, or IONOS_TOKEN
	// IONOS_API_URL - should be set only if you with to overwrite the default ionoscloud.DefaultIonosServerUrl
	cfg := ionoscloud.NewConfigurationFromEnv()

	// Uncomment line if you want to see request/response packets
	// We recommend you only set this field for debugging purposes.
	// Disable it in your production environments because it can log sensitive data.
	// It logs the full request and response without encryption, even for an HTTPS call.
	// Verbose request and response logging can also significantly impact your application's performance.
	//cfg.Debug = true

	computeClient := ionoscloud.NewAPIClient(cfg)

	ctx := context.Background()

	return apiService{
		computeEngine: computeClient,
		ctx:           ctx,
	}

}

// createDcAndWaitUntilDone - creates datacenter and waits until provisioning is successful
// return - datacenter object created, or error
func (as apiService) createDcAndWaitUntilDone() (*ionoscloud.Datacenter, error) {

	datacenterName := "testDatacenter"
	description := "this is the test datacenter"
	loc := "de/txl" // other location values: de/fra", "us/las", "us/ewr", "de/txl", gb/lhr", "es/vit"
	dc, apiResponse, err := as.createDatacenter(datacenterName, description, loc)
	if err != nil {
		return nil, fmt.Errorf(
			"error creating data center (%w)", err)
	}
	// gets the Location Header value, where Request ID is stored, to interrogate the request status
	requestPath := getRequestPath(apiResponse)
	if requestPath == "" {
		return nil, fmt.Errorf("error getting location from header for datacenter")
	}

	// Waits for the datacenter creation to finish. Polls until it receives an answer that
	// provisioning is successful
	err = as.waitForRequestToBeDone(requestPath)
	if err != nil {
		return nil, fmt.Errorf("error while waiting for datacenter creation to finish (%w)", err)
	}
	return &dc, nil
}

// createServerAndWaitUntilDone - creates server and waits until provisioning is successful
// return - server object created, or error
func (as apiService) createServerAndWaitUntilDone(dcId string) (*ionoscloud.Server, error) {
	serverName := "testServer"
	var cores int32 = 1
	var ram int32 = 1024         // must be a multiple of 1024 : 1024, 2048, 3072, 4096, etc
	cpuFamily := "INTEL_SKYLAKE" // other valid values: AMD_OPTERON,INTEL_XEON, INTEL_SKYLAKE.
	// cpu_family should be in sync with what is available in the location you choose for the datacenter
	server, apiResponse, err := as.createServer(dcId, serverName, cores, ram, cpuFamily)
	if err != nil {
		return nil, fmt.Errorf(
			"error creating server (%w)", err)
	}
	// The initial response from the cloud is a HTTP/2.0 202 Accepted - after this response, the IONOS Cloud API
	// starts to actually create the server
	// Gets path to interrogate server creation status
	requestPath := getRequestPath(apiResponse)
	if requestPath == "" {
		return nil, fmt.Errorf("error getting server path")
	}
	// Waits for the server creation to finish. It takes some time to create
	// a compute resource, so we poll until provisioning is successful
	err = as.waitForRequestToBeDone(requestPath)
	if err != nil {
		return nil, fmt.Errorf("error while waiting for server creation to finish (%w)", err)
	}
	return &server, nil
}

func (as apiService) createDatacenter(name, description, location string) (ionoscloud.Datacenter, *ionoscloud.APIResponse, error) {
	// The required parameter for datacenter creation is: 'location'.
	// Creates the datacenter structure and populates it
	dc := ionoscloud.Datacenter{
		Properties: &ionoscloud.DatacenterProperties{
			Name:        &name,
			Description: &description,
			Location:    &location,
		},
	}
	// The computeClient has access to all the resources in the ionos compute ecosystem. First we get the DatacenterApi.
	// The datacenter is the basic building block in which to create your infrastructure.
	// Builder pattern is used, to allow for easier creation and cleaner code.
	// In this case, the order is DatacentersPost -> Datacenter (loads datacenter structure) -> Execute
	// The final step that actually sends the request is 'execute'.
	return as.computeEngine.DataCentersApi.DatacentersPost(as.ctx).Datacenter(dc).Execute()
}

func (as apiService) createServer(datacenterId, name string, cores, ram int32, cpuFamily string) (ionoscloud.Server, *ionoscloud.APIResponse, error) {
	// Required parameters for server creation: 'cores' and 'ram'.
	// Creates the server structure and populates it
	server := ionoscloud.Server{
		Properties: &ionoscloud.ServerProperties{
			Name:      &name,
			Cores:     &cores,
			Ram:       &ram,
			CpuFamily: &cpuFamily,
		},
	}

	return as.computeEngine.ServersApi.DatacentersServersPost(as.ctx, datacenterId).Server(server).Execute()
}

// waitForRequestToBeDone - polls until the request is 'Done', or
// until the context timeout expires
func (as apiService) waitForRequestToBeDone(path string) error {
	// Polls until context timeout expires
	_, err := as.computeEngine.WaitForRequest(as.ctx, path)
	if err != nil {
		return fmt.Errorf("error waiting for status for %s : (%w)", path, err)
	}
	log.Printf("resource created for path %s", path)
	return nil
}

// getRequestPath - returns location header value which is the path
// used to poll the request for readiness
func getRequestPath(resp *ionoscloud.APIResponse) string {
	if resp != nil {
		return resp.Header.Get("location")
	}
	return ""
}
