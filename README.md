# Introduction 
GoLang Terratest Packages

# Examples
https://github.com/gruntwork-io/terratest/tree/master/examples/azure
https://github.com/Azure/azure-sdk-for-go/tree/main/services

# Modules Required Parameters
Each module required parameters can be fed from terraform outputs.

```
parameterName := terraform.Output(t, terraformOptions, "parameter_name")
```

## aks
  - tenantID
  - clientID
  - subscriptionID
  - resourceGroupName
  - aksClusterName

## loadbalancer
  - subscriptionID
  - resourceGroupName
  - loadBalancerName

## rgrp 
  - subscriptionID
  - resourceGroupName
