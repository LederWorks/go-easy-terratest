package loadbalancer

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestLoadBalancerExists(t *testing.T, terraformOptions *terraform.Options, subscriptionID string, resourceGroupName string, loadBalancerName string) {
	// Run `terraform output` to get the values of output variables.
	_, err := azure.LoadBalancerExistsE(loadBalancerName, resourceGroupName, subscriptionID)

	require.Error(t, err)
}
