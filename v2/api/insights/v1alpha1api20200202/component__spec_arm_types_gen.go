// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Component_SpecARM struct {
	AzureName string `json:"azureName"`

	//Etag: Resource etag
	Etag *string `json:"etag,omitempty"`

	//Kind: The kind of application that this component refers to, used to customize
	//UI. This value is a freeform string, values should typically be one of the
	//following: web, ios, other, store, java, phone.
	Kind string `json:"kind"`

	//Location: Resource location
	Location string `json:"location"`
	Name     string `json:"name"`

	//Properties: Properties that define an Application Insights component resource.
	Properties *ApplicationInsightsComponentPropertiesARM `json:"properties,omitempty"`

	//Tags: Resource tags
	Tags *v1.JSON `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Component_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-02-02"
func (component Component_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (component Component_SpecARM) GetName() string {
	return component.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/components"
func (component Component_SpecARM) GetType() string {
	return "Microsoft.Insights/components"
}

type ApplicationInsightsComponentPropertiesARM struct {
	//Application_Type: Type of application being monitored.
	Application_Type ApplicationInsightsComponentPropertiesApplication_Type `json:"Application_Type"`

	//DisableIpMasking: Disable IP masking.
	DisableIpMasking *bool `json:"DisableIpMasking,omitempty"`

	//DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"DisableLocalAuth,omitempty"`

	//Flow_Type: Used by the Application Insights system to determine what kind of
	//flow this component was created by. This is to be set to 'Bluefield' when
	//creating/updating a component via the REST API.
	Flow_Type *ApplicationInsightsComponentPropertiesFlow_Type `json:"Flow_Type,omitempty"`

	//ForceCustomerStorageForProfiler: Force users to create their own storage account
	//for profiler and debugger.
	ForceCustomerStorageForProfiler *bool `json:"ForceCustomerStorageForProfiler,omitempty"`

	//HockeyAppId: The unique application ID created when a new application is added
	//to HockeyApp, used for communications with HockeyApp.
	HockeyAppId *string `json:"HockeyAppId,omitempty"`

	//ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"ImmediatePurgeDataOn30Days,omitempty"`

	//IngestionMode: Indicates the flow of the ingestion.
	IngestionMode *ApplicationInsightsComponentPropertiesIngestionMode `json:"IngestionMode,omitempty"`

	//PublicNetworkAccessForIngestion: The network access type for accessing
	//Application Insights ingestion.
	PublicNetworkAccessForIngestion *PublicNetworkAccessType `json:"publicNetworkAccessForIngestion,omitempty"`

	//PublicNetworkAccessForQuery: The network access type for accessing Application
	//Insights query.
	PublicNetworkAccessForQuery *PublicNetworkAccessType `json:"publicNetworkAccessForQuery,omitempty"`

	//Request_Source: Describes what tool created this Application Insights component.
	//Customers using this API should set this to the default 'rest'.
	Request_Source *ApplicationInsightsComponentPropertiesRequest_Source `json:"Request_Source,omitempty"`

	//RetentionInDays: Retention period in days.
	RetentionInDays *int `json:"RetentionInDays,omitempty"`

	//SamplingPercentage: Percentage of the data produced by the application being
	//monitored that is being sampled for Application Insights telemetry.
	SamplingPercentage *float64 `json:"SamplingPercentage,omitempty"`

	//WorkspaceResourceId: Resource Id of the log analytics workspace which the data
	//will be ingested to. This property is required to create an application with
	//this API version. Applications from older versions will not have this property.
	WorkspaceResourceId *string `json:"WorkspaceResourceId,omitempty"`
}
