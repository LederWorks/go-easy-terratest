package aks

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2017-08-31/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestAKSCluster(t *testing.T, terraformOptions *terraform.Options, tenantID string, clientID string, subscriptionID string, resourceGroupName string, aksClusterName string) {

	// Verify the AKS Cluster exists
	aksCli, err := getAKSClient(tenantID, clientID, subscriptionID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cluster, err := aksCli.Get(context.Background(), resourceGroupName, aksClusterName)
	if err != nil {
		fmt.Println("unable to find cluster by name:", aksClusterName)
		fmt.Println("check resource group and cluster name vars")
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println("Checking name")
	assert.Equal(t, aksClusterName, *cluster.Name)
}

func getAKSClient(tenantID string, clientID string, subscriptionID string) (containerservice.ManagedClustersClient, error) {
	aksClient := containerservice.NewManagedClustersClient(subscriptionID)
	authz := auth.NewClientCredentialsConfig(clientID, os.Getenv("ARM_CLIENT_SECRET"), tenantID)
	authzr, _ := authz.Authorizer()
	aksClient.Authorizer = authzr
	aksClient.PollingDuration = time.Hour * 1
	return aksClient, nil
}
