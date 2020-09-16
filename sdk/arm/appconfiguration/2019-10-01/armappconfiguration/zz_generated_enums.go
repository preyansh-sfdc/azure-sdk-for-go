// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

// ConfigurationResourceType - The resource type to check for name availability.
type ConfigurationResourceType string

const (
	ConfigurationResourceTypeMicrosoftAppConfigurationStores ConfigurationResourceType = "Microsoft.AppConfiguration/configurationStores"
)

func PossibleConfigurationResourceTypeValues() []ConfigurationResourceType {
	return []ConfigurationResourceType{
		ConfigurationResourceTypeMicrosoftAppConfigurationStores,
	}
}

func (c ConfigurationResourceType) ToPtr() *ConfigurationResourceType {
	return &c
}

// IDentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
type IDentityType string

const (
	IDentityTypeNone                       IDentityType = "None"
	IDentityTypeSystemAssigned             IDentityType = "SystemAssigned"
	IDentityTypeSystemAssignedUserAssigned IDentityType = "SystemAssigned, UserAssigned"
	IDentityTypeUserAssigned               IDentityType = "UserAssigned"
)

func PossibleIDentityTypeValues() []IDentityType {
	return []IDentityType{
		IDentityTypeNone,
		IDentityTypeSystemAssigned,
		IDentityTypeSystemAssignedUserAssigned,
		IDentityTypeUserAssigned,
	}
}

func (c IDentityType) ToPtr() *IDentityType {
	return &c
}

// ProvisioningState - The provisioning state of the configuration store.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}