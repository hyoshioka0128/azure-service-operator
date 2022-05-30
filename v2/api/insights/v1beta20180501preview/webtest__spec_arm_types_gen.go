// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180501preview

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Webtest_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Kind: The kind of WebTest that this web test watches. Choices are ping and multistep.
	Kind *Webtest_Spec_Kind `json:"kind,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Metadata describing a web test for an Azure resource.
	Properties *WebTestPropertiesARM `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags *v1.JSON `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Webtest_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01-preview"
func (webtest Webtest_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (webtest *Webtest_SpecARM) GetName() string {
	return webtest.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (webtest *Webtest_SpecARM) GetType() string {
	return ""
}

type WebTestPropertiesARM struct {
	// Configuration: An XML configuration specification for a WebTest.
	Configuration *WebTestProperties_ConfigurationARM `json:"Configuration,omitempty"`

	// Description: User defined description for this WebTest.
	Description *string `json:"Description,omitempty"`

	// Enabled: Is the test actively being monitored.
	Enabled *bool `json:"Enabled,omitempty"`

	// Frequency: Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency *int `json:"Frequency,omitempty"`

	// Kind: The kind of web test this is, valid choices are ping, multistep, basic, and standard.
	Kind *WebTestProperties_Kind `json:"Kind,omitempty"`

	// Locations: A list of where to physically run the tests from to give global coverage for accessibility of your
	// application.
	Locations []WebTestGeolocationARM `json:"Locations,omitempty"`

	// Name: User defined name if this WebTest.
	Name *string `json:"Name,omitempty"`

	// Request: The collection of request properties
	Request *WebTestProperties_RequestARM `json:"Request,omitempty"`

	// RetryEnabled: Allow for retries should this WebTest fail.
	RetryEnabled *bool `json:"RetryEnabled,omitempty"`

	// SyntheticMonitorId: Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorId *string `json:"SyntheticMonitorId,omitempty"`

	// Timeout: Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout *int `json:"Timeout,omitempty"`

	// ValidationRules: The collection of validation rule properties
	ValidationRules *WebTestProperties_ValidationRulesARM `json:"ValidationRules,omitempty"`
}

// +kubebuilder:validation:Enum={"multistep","ping"}
type Webtest_Spec_Kind string

const (
	Webtest_Spec_Kind_Multistep = Webtest_Spec_Kind("multistep")
	Webtest_Spec_Kind_Ping      = Webtest_Spec_Kind("ping")
)

type WebTestGeolocationARM struct {
	// Id: Location ID for the WebTest to run from.
	Id *string `json:"Id,omitempty"`
}

type WebTestProperties_ConfigurationARM struct {
	// WebTest: The XML specification of a WebTest to run against an application.
	WebTest *string `json:"WebTest,omitempty"`
}

type WebTestProperties_RequestARM struct {
	// FollowRedirects: Follow redirects for this web test.
	FollowRedirects *bool `json:"FollowRedirects,omitempty"`

	// Headers: List of headers and their values to add to the WebTest call.
	Headers []HeaderFieldARM `json:"Headers,omitempty"`

	// HttpVerb: Http verb to use for this web test.
	HttpVerb *string `json:"HttpVerb,omitempty"`

	// ParseDependentRequests: Parse Dependent request for this WebTest.
	ParseDependentRequests *bool `json:"ParseDependentRequests,omitempty"`

	// RequestBody: Base64 encoded string body to send with this web test.
	RequestBody *string `json:"RequestBody,omitempty"`

	// RequestUrl: Url location to test.
	RequestUrl *string `json:"RequestUrl,omitempty"`
}

type WebTestProperties_ValidationRulesARM struct {
	// ContentValidation: The collection of content validation properties
	ContentValidation *WebTestProperties_ValidationRules_ContentValidationARM `json:"ContentValidation,omitempty"`

	// ExpectedHttpStatusCode: Validate that the WebTest returns the http status code provided.
	ExpectedHttpStatusCode *int `json:"ExpectedHttpStatusCode,omitempty"`

	// IgnoreHttpsStatusCode: When set, validation will ignore the status code.
	IgnoreHttpsStatusCode *bool `json:"IgnoreHttpsStatusCode,omitempty"`

	// SSLCertRemainingLifetimeCheck: A number of days to check still remain before the the existing SSL cert expires.  Value
	// must be positive and the SSLCheck must be set to true.
	SSLCertRemainingLifetimeCheck *int `json:"SSLCertRemainingLifetimeCheck,omitempty"`

	// SSLCheck: Checks to see if the SSL cert is still valid.
	SSLCheck *bool `json:"SSLCheck,omitempty"`
}

type HeaderFieldARM struct {
	// Key: The name of the header.
	Key *string `json:"key,omitempty"`

	// Value: The value of the header.
	Value *string `json:"value,omitempty"`
}

type WebTestProperties_ValidationRules_ContentValidationARM struct {
	// ContentMatch: Content to look for in the return of the WebTest.  Must not be null or empty.
	ContentMatch *string `json:"ContentMatch,omitempty"`

	// IgnoreCase: When set, this value makes the ContentMatch validation case insensitive.
	IgnoreCase *bool `json:"IgnoreCase,omitempty"`

	// PassIfTextFound: When true, validation will pass if there is a match for the ContentMatch string.  If false, validation
	// will fail if there is a match
	PassIfTextFound *bool `json:"PassIfTextFound,omitempty"`
}
