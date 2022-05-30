// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20180501preview

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Deprecated version of Webtest_STATUS. Use v1beta20180501preview.Webtest_STATUS instead
type Webtest_STATUSARM struct {
	Id         *string                      `json:"id,omitempty"`
	Kind       *Webtest_Kind_STATUS         `json:"kind,omitempty"`
	Location   *string                      `json:"location,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *WebTestProperties_STATUSARM `json:"properties,omitempty"`
	Tags       *v1.JSON                     `json:"tags,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}

// Deprecated version of WebTestProperties_STATUS. Use v1beta20180501preview.WebTestProperties_STATUS instead
type WebTestProperties_STATUSARM struct {
	Configuration      *WebTestProperties_Configuration_STATUSARM   `json:"Configuration,omitempty"`
	Description        *string                                      `json:"Description,omitempty"`
	Enabled            *bool                                        `json:"Enabled,omitempty"`
	Frequency          *int                                         `json:"Frequency,omitempty"`
	Kind               *WebTestProperties_Kind_STATUS               `json:"Kind,omitempty"`
	Locations          []WebTestGeolocation_STATUSARM               `json:"Locations,omitempty"`
	Name               *string                                      `json:"Name,omitempty"`
	ProvisioningState  *string                                      `json:"provisioningState,omitempty"`
	Request            *WebTestProperties_Request_STATUSARM         `json:"Request,omitempty"`
	RetryEnabled       *bool                                        `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                                      `json:"SyntheticMonitorId,omitempty"`
	Timeout            *int                                         `json:"Timeout,omitempty"`
	ValidationRules    *WebTestProperties_ValidationRules_STATUSARM `json:"ValidationRules,omitempty"`
}

// Deprecated version of Webtest_Kind_STATUS. Use v1beta20180501preview.Webtest_Kind_STATUS instead
type Webtest_Kind_STATUS string

const (
	Webtest_Kind_Multistep_STATUS = Webtest_Kind_STATUS("multistep")
	Webtest_Kind_Ping_STATUS      = Webtest_Kind_STATUS("ping")
)

// Deprecated version of WebTestGeolocation_STATUS. Use v1beta20180501preview.WebTestGeolocation_STATUS instead
type WebTestGeolocation_STATUSARM struct {
	Id *string `json:"Id,omitempty"`
}

// Deprecated version of WebTestProperties_Configuration_STATUS. Use v1beta20180501preview.WebTestProperties_Configuration_STATUS instead
type WebTestProperties_Configuration_STATUSARM struct {
	WebTest *string `json:"WebTest,omitempty"`
}

// Deprecated version of WebTestProperties_Request_STATUS. Use v1beta20180501preview.WebTestProperties_Request_STATUS instead
type WebTestProperties_Request_STATUSARM struct {
	FollowRedirects        *bool                   `json:"FollowRedirects,omitempty"`
	Headers                []HeaderField_STATUSARM `json:"Headers,omitempty"`
	HttpVerb               *string                 `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool                   `json:"ParseDependentRequests,omitempty"`
	RequestBody            *string                 `json:"RequestBody,omitempty"`
	RequestUrl             *string                 `json:"RequestUrl,omitempty"`
}

// Deprecated version of WebTestProperties_ValidationRules_STATUS. Use v1beta20180501preview.WebTestProperties_ValidationRules_STATUS instead
type WebTestProperties_ValidationRules_STATUSARM struct {
	ContentValidation             *WebTestProperties_ValidationRules_ContentValidation_STATUSARM `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                                           `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                                          `json:"IgnoreHttpsStatusCode,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                                           `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                                          `json:"SSLCheck,omitempty"`
}

// Deprecated version of HeaderField_STATUS. Use v1beta20180501preview.HeaderField_STATUS instead
type HeaderField_STATUSARM struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Deprecated version of WebTestProperties_ValidationRules_ContentValidation_STATUS. Use v1beta20180501preview.WebTestProperties_ValidationRules_ContentValidation_STATUS instead
type WebTestProperties_ValidationRules_ContentValidation_STATUSARM struct {
	ContentMatch    *string `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool   `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool   `json:"PassIfTextFound,omitempty"`
}