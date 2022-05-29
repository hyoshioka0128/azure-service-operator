// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20180501preview

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Deprecated version of Webtest_Spec. Use v1beta20180501preview.Webtest_Spec instead
type Webtest_SpecARM struct {
	AzureName  string                `json:"azureName,omitempty"`
	Id         *string               `json:"id,omitempty"`
	Kind       *Webtest_Spec_Kind    `json:"kind,omitempty"`
	Location   *string               `json:"location,omitempty"`
	Name       string                `json:"name,omitempty"`
	Properties *WebTestPropertiesARM `json:"properties,omitempty"`
	Tags       *v1.JSON              `json:"tags,omitempty"`
	Type       *string               `json:"type,omitempty"`
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

// Deprecated version of WebTestProperties. Use v1beta20180501preview.WebTestProperties instead
type WebTestPropertiesARM struct {
	Configuration      *WebTestProperties_ConfigurationARM   `json:"Configuration,omitempty"`
	Description        *string                               `json:"Description,omitempty"`
	Enabled            *bool                                 `json:"Enabled,omitempty"`
	Frequency          *int                                  `json:"Frequency,omitempty"`
	Kind               *WebTestProperties_Kind               `json:"Kind,omitempty"`
	Locations          []WebTestGeolocationARM               `json:"Locations,omitempty"`
	Name               *string                               `json:"Name,omitempty"`
	ProvisioningState  *string                               `json:"provisioningState,omitempty"`
	Request            *WebTestProperties_RequestARM         `json:"Request,omitempty"`
	RetryEnabled       *bool                                 `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                               `json:"SyntheticMonitorId,omitempty"`
	Timeout            *int                                  `json:"Timeout,omitempty"`
	ValidationRules    *WebTestProperties_ValidationRulesARM `json:"ValidationRules,omitempty"`
}

// Deprecated version of Webtest_Spec_Kind. Use v1beta20180501preview.Webtest_Spec_Kind instead
// +kubebuilder:validation:Enum={"multistep","ping"}
type Webtest_Spec_Kind string

const (
	Webtest_Spec_Kind_Multistep = Webtest_Spec_Kind("multistep")
	Webtest_Spec_Kind_Ping      = Webtest_Spec_Kind("ping")
)

// Deprecated version of WebTestGeolocation. Use v1beta20180501preview.WebTestGeolocation instead
type WebTestGeolocationARM struct {
	Id *string `json:"Id,omitempty"`
}

// Deprecated version of WebTestProperties_Configuration. Use v1beta20180501preview.WebTestProperties_Configuration instead
type WebTestProperties_ConfigurationARM struct {
	WebTest *string `json:"WebTest,omitempty"`
}

// Deprecated version of WebTestProperties_Request. Use v1beta20180501preview.WebTestProperties_Request instead
type WebTestProperties_RequestARM struct {
	FollowRedirects        *bool            `json:"FollowRedirects,omitempty"`
	Headers                []HeaderFieldARM `json:"Headers,omitempty"`
	HttpVerb               *string          `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool            `json:"ParseDependentRequests,omitempty"`
	RequestBody            *string          `json:"RequestBody,omitempty"`
	RequestUrl             *string          `json:"RequestUrl,omitempty"`
}

// Deprecated version of WebTestProperties_ValidationRules. Use v1beta20180501preview.WebTestProperties_ValidationRules instead
type WebTestProperties_ValidationRulesARM struct {
	ContentValidation             *WebTestProperties_ValidationRules_ContentValidationARM `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                                    `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                                   `json:"IgnoreHttpsStatusCode,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                                    `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                                   `json:"SSLCheck,omitempty"`
}

// Deprecated version of HeaderField. Use v1beta20180501preview.HeaderField instead
type HeaderFieldARM struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Deprecated version of WebTestProperties_ValidationRules_ContentValidation. Use v1beta20180501preview.WebTestProperties_ValidationRules_ContentValidation instead
type WebTestProperties_ValidationRules_ContentValidationARM struct {
	ContentMatch    *string `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool   `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool   `json:"PassIfTextFound,omitempty"`
}
