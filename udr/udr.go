// package udr

// import (
// 	"context"
// 	"testing"

// 	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
// 	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
// 	"github.com/gruntwork-io/terratest/modules/terraform"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRouteTableExists(t *testing.T, terraformOptions *terraform.Options, subscriptionID string, resourceGroupName string, routeTableName string) {
// 	// Authorize
// 	cred, err := azidentity.NewDefaultAzureCredential(nil)
// 	if err != nil {
// 	}
// 	//Set up the client API, nil= no options feeded
// 	netClient := armnetwork.NewRouteTablesClient(subscriptionID, cred, nil)
// 	// Verify the UDR Exists
// 	res, err := netClient.Get(context.Background(), resourceGroupName, routeTableName, &armnetwork.RouteTablesClientGetOptions{Expand: nil})
// 	assert.True(t, res.ID != nil, "Route table does not exist")
// }