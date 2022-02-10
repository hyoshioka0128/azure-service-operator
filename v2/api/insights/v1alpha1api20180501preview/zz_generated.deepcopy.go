//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20180501preview

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField_Spec) DeepCopyInto(out *HeaderField_Spec) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField_Spec.
func (in *HeaderField_Spec) DeepCopy() *HeaderField_Spec {
	if in == nil {
		return nil
	}
	out := new(HeaderField_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField_SpecARM) DeepCopyInto(out *HeaderField_SpecARM) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField_SpecARM.
func (in *HeaderField_SpecARM) DeepCopy() *HeaderField_SpecARM {
	if in == nil {
		return nil
	}
	out := new(HeaderField_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField_Status) DeepCopyInto(out *HeaderField_Status) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField_Status.
func (in *HeaderField_Status) DeepCopy() *HeaderField_Status {
	if in == nil {
		return nil
	}
	out := new(HeaderField_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField_StatusARM) DeepCopyInto(out *HeaderField_StatusARM) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField_StatusARM.
func (in *HeaderField_StatusARM) DeepCopy() *HeaderField_StatusARM {
	if in == nil {
		return nil
	}
	out := new(HeaderField_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation_Spec) DeepCopyInto(out *WebTestGeolocation_Spec) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation_Spec.
func (in *WebTestGeolocation_Spec) DeepCopy() *WebTestGeolocation_Spec {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation_SpecARM) DeepCopyInto(out *WebTestGeolocation_SpecARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation_SpecARM.
func (in *WebTestGeolocation_SpecARM) DeepCopy() *WebTestGeolocation_SpecARM {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation_Status) DeepCopyInto(out *WebTestGeolocation_Status) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation_Status.
func (in *WebTestGeolocation_Status) DeepCopy() *WebTestGeolocation_Status {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation_StatusARM) DeepCopyInto(out *WebTestGeolocation_StatusARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation_StatusARM.
func (in *WebTestGeolocation_StatusARM) DeepCopy() *WebTestGeolocation_StatusARM {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_SpecARM) DeepCopyInto(out *WebTestProperties_SpecARM) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_Spec_ConfigurationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation_SpecARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_Spec_RequestARM)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_Spec_ValidationRulesARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_SpecARM.
func (in *WebTestProperties_SpecARM) DeepCopy() *WebTestProperties_SpecARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_Configuration) DeepCopyInto(out *WebTestProperties_Spec_Configuration) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_Configuration.
func (in *WebTestProperties_Spec_Configuration) DeepCopy() *WebTestProperties_Spec_Configuration {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_ConfigurationARM) DeepCopyInto(out *WebTestProperties_Spec_ConfigurationARM) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_ConfigurationARM.
func (in *WebTestProperties_Spec_ConfigurationARM) DeepCopy() *WebTestProperties_Spec_ConfigurationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_ConfigurationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_Request) DeepCopyInto(out *WebTestProperties_Spec_Request) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField_Spec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_Request.
func (in *WebTestProperties_Spec_Request) DeepCopy() *WebTestProperties_Spec_Request {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_RequestARM) DeepCopyInto(out *WebTestProperties_Spec_RequestARM) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField_SpecARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_RequestARM.
func (in *WebTestProperties_Spec_RequestARM) DeepCopy() *WebTestProperties_Spec_RequestARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_RequestARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_ValidationRules) DeepCopyInto(out *WebTestProperties_Spec_ValidationRules) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_Spec_ValidationRules_ContentValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_ValidationRules.
func (in *WebTestProperties_Spec_ValidationRules) DeepCopy() *WebTestProperties_Spec_ValidationRules {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_ValidationRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_ValidationRulesARM) DeepCopyInto(out *WebTestProperties_Spec_ValidationRulesARM) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_Spec_ValidationRules_ContentValidationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_ValidationRulesARM.
func (in *WebTestProperties_Spec_ValidationRulesARM) DeepCopy() *WebTestProperties_Spec_ValidationRulesARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_ValidationRulesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_ValidationRules_ContentValidation) DeepCopyInto(out *WebTestProperties_Spec_ValidationRules_ContentValidation) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_ValidationRules_ContentValidation.
func (in *WebTestProperties_Spec_ValidationRules_ContentValidation) DeepCopy() *WebTestProperties_Spec_ValidationRules_ContentValidation {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_ValidationRules_ContentValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Spec_ValidationRules_ContentValidationARM) DeepCopyInto(out *WebTestProperties_Spec_ValidationRules_ContentValidationARM) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Spec_ValidationRules_ContentValidationARM.
func (in *WebTestProperties_Spec_ValidationRules_ContentValidationARM) DeepCopy() *WebTestProperties_Spec_ValidationRules_ContentValidationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Spec_ValidationRules_ContentValidationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_StatusARM) DeepCopyInto(out *WebTestProperties_StatusARM) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_Status_ConfigurationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation_StatusARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_Status_RequestARM)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_Status_ValidationRulesARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_StatusARM.
func (in *WebTestProperties_StatusARM) DeepCopy() *WebTestProperties_StatusARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_Configuration) DeepCopyInto(out *WebTestProperties_Status_Configuration) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_Configuration.
func (in *WebTestProperties_Status_Configuration) DeepCopy() *WebTestProperties_Status_Configuration {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_ConfigurationARM) DeepCopyInto(out *WebTestProperties_Status_ConfigurationARM) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_ConfigurationARM.
func (in *WebTestProperties_Status_ConfigurationARM) DeepCopy() *WebTestProperties_Status_ConfigurationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_ConfigurationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_Request) DeepCopyInto(out *WebTestProperties_Status_Request) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_Request.
func (in *WebTestProperties_Status_Request) DeepCopy() *WebTestProperties_Status_Request {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_RequestARM) DeepCopyInto(out *WebTestProperties_Status_RequestARM) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField_StatusARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_RequestARM.
func (in *WebTestProperties_Status_RequestARM) DeepCopy() *WebTestProperties_Status_RequestARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_RequestARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_ValidationRules) DeepCopyInto(out *WebTestProperties_Status_ValidationRules) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_Status_ValidationRules_ContentValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_ValidationRules.
func (in *WebTestProperties_Status_ValidationRules) DeepCopy() *WebTestProperties_Status_ValidationRules {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_ValidationRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_ValidationRulesARM) DeepCopyInto(out *WebTestProperties_Status_ValidationRulesARM) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_Status_ValidationRules_ContentValidationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_ValidationRulesARM.
func (in *WebTestProperties_Status_ValidationRulesARM) DeepCopy() *WebTestProperties_Status_ValidationRulesARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_ValidationRulesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_ValidationRules_ContentValidation) DeepCopyInto(out *WebTestProperties_Status_ValidationRules_ContentValidation) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_ValidationRules_ContentValidation.
func (in *WebTestProperties_Status_ValidationRules_ContentValidation) DeepCopy() *WebTestProperties_Status_ValidationRules_ContentValidation {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_ValidationRules_ContentValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Status_ValidationRules_ContentValidationARM) DeepCopyInto(out *WebTestProperties_Status_ValidationRules_ContentValidationARM) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Status_ValidationRules_ContentValidationARM.
func (in *WebTestProperties_Status_ValidationRules_ContentValidationARM) DeepCopy() *WebTestProperties_Status_ValidationRules_ContentValidationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Status_ValidationRules_ContentValidationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTest_Status) DeepCopyInto(out *WebTest_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_Status_Configuration)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestPropertiesStatusKind)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertiesName != nil {
		in, out := &in.PropertiesName, &out.PropertiesName
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_Status_Request)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_Status_ValidationRules)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTest_Status.
func (in *WebTest_Status) DeepCopy() *WebTest_Status {
	if in == nil {
		return nil
	}
	out := new(WebTest_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTest_StatusARM) DeepCopyInto(out *WebTest_StatusARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(WebTestProperties_StatusARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTest_StatusARM.
func (in *WebTest_StatusARM) DeepCopy() *WebTest_StatusARM {
	if in == nil {
		return nil
	}
	out := new(WebTest_StatusARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest) DeepCopyInto(out *Webtest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest.
func (in *Webtest) DeepCopy() *Webtest {
	if in == nil {
		return nil
	}
	out := new(Webtest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Webtest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebtestList) DeepCopyInto(out *WebtestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Webtest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebtestList.
func (in *WebtestList) DeepCopy() *WebtestList {
	if in == nil {
		return nil
	}
	out := new(WebtestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebtestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtests_SPEC) DeepCopyInto(out *Webtests_SPEC) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_Spec_Configuration)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestPropertiesSpecKind)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation_Spec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	out.Owner = in.Owner
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_Spec_Request)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_Spec_ValidationRules)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtests_SPEC.
func (in *Webtests_SPEC) DeepCopy() *Webtests_SPEC {
	if in == nil {
		return nil
	}
	out := new(Webtests_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtests_SPECARM) DeepCopyInto(out *Webtests_SPECARM) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(WebTestProperties_SpecARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtests_SPECARM.
func (in *Webtests_SPECARM) DeepCopy() *Webtests_SPECARM {
	if in == nil {
		return nil
	}
	out := new(Webtests_SPECARM)
	in.DeepCopyInto(out)
	return out
}