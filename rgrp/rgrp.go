package rgrp

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestResourceGroupExists(t *testing.T, terraformOptions *terraform.Options, subscriptionID string, resourceGroupName string) {
	// Verify the resource group exists
	exists := azure.ResourceGroupExists(t, resourceGroupName, subscriptionID)
	assert.True(t, exists, "Resource group does not exist")
}
