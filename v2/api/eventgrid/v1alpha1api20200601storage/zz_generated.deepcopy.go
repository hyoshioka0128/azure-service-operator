//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedFilter_Spec) DeepCopyInto(out *AdvancedFilter_Spec) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.OperatorType != nil {
		in, out := &in.OperatorType, &out.OperatorType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedFilter_Spec.
func (in *AdvancedFilter_Spec) DeepCopy() *AdvancedFilter_Spec {
	if in == nil {
		return nil
	}
	out := new(AdvancedFilter_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdvancedFilter_Status) DeepCopyInto(out *AdvancedFilter_Status) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.OperatorType != nil {
		in, out := &in.OperatorType, &out.OperatorType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdvancedFilter_Status.
func (in *AdvancedFilter_Status) DeepCopy() *AdvancedFilter_Status {
	if in == nil {
		return nil
	}
	out := new(AdvancedFilter_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeadLetterDestination_Spec) DeepCopyInto(out *DeadLetterDestination_Spec) {
	*out = *in
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeadLetterDestination_Spec.
func (in *DeadLetterDestination_Spec) DeepCopy() *DeadLetterDestination_Spec {
	if in == nil {
		return nil
	}
	out := new(DeadLetterDestination_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeadLetterDestination_Status) DeepCopyInto(out *DeadLetterDestination_Status) {
	*out = *in
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeadLetterDestination_Status.
func (in *DeadLetterDestination_Status) DeepCopy() *DeadLetterDestination_Status {
	if in == nil {
		return nil
	}
	out := new(DeadLetterDestination_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain_Status) DeepCopyInto(out *Domain_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.InboundIpRules != nil {
		in, out := &in.InboundIpRules, &out.InboundIpRules
		*out = make([]InboundIpRule_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InputSchema != nil {
		in, out := &in.InputSchema, &out.InputSchema
		*out = new(string)
		**out = **in
	}
	if in.InputSchemaMapping != nil {
		in, out := &in.InputSchemaMapping, &out.InputSchemaMapping
		*out = new(InputSchemaMapping_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MetricResourceId != nil {
		in, out := &in.MetricResourceId, &out.MetricResourceId
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_Status_Domain_SubResourceEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain_Status.
func (in *Domain_Status) DeepCopy() *Domain_Status {
	if in == nil {
		return nil
	}
	out := new(Domain_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domains_SPEC) DeepCopyInto(out *Domains_SPEC) {
	*out = *in
	if in.InboundIpRules != nil {
		in, out := &in.InboundIpRules, &out.InboundIpRules
		*out = make([]InboundIpRule_Spec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InputSchema != nil {
		in, out := &in.InputSchema, &out.InputSchema
		*out = new(string)
		**out = **in
	}
	if in.InputSchemaMapping != nil {
		in, out := &in.InputSchemaMapping, &out.InputSchemaMapping
		*out = new(InputSchemaMapping_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	out.Owner = in.Owner
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domains_SPEC.
func (in *Domains_SPEC) DeepCopy() *Domains_SPEC {
	if in == nil {
		return nil
	}
	out := new(Domains_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscription) DeepCopyInto(out *EventSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscription.
func (in *EventSubscription) DeepCopy() *EventSubscription {
	if in == nil {
		return nil
	}
	out := new(EventSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionDestination_Spec) DeepCopyInto(out *EventSubscriptionDestination_Spec) {
	*out = *in
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionDestination_Spec.
func (in *EventSubscriptionDestination_Spec) DeepCopy() *EventSubscriptionDestination_Spec {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionDestination_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionDestination_Status) DeepCopyInto(out *EventSubscriptionDestination_Status) {
	*out = *in
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionDestination_Status.
func (in *EventSubscriptionDestination_Status) DeepCopy() *EventSubscriptionDestination_Status {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionDestination_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionFilter_Spec) DeepCopyInto(out *EventSubscriptionFilter_Spec) {
	*out = *in
	if in.AdvancedFilters != nil {
		in, out := &in.AdvancedFilters, &out.AdvancedFilters
		*out = make([]AdvancedFilter_Spec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludedEventTypes != nil {
		in, out := &in.IncludedEventTypes, &out.IncludedEventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IsSubjectCaseSensitive != nil {
		in, out := &in.IsSubjectCaseSensitive, &out.IsSubjectCaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubjectBeginsWith != nil {
		in, out := &in.SubjectBeginsWith, &out.SubjectBeginsWith
		*out = new(string)
		**out = **in
	}
	if in.SubjectEndsWith != nil {
		in, out := &in.SubjectEndsWith, &out.SubjectEndsWith
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionFilter_Spec.
func (in *EventSubscriptionFilter_Spec) DeepCopy() *EventSubscriptionFilter_Spec {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionFilter_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionFilter_Status) DeepCopyInto(out *EventSubscriptionFilter_Status) {
	*out = *in
	if in.AdvancedFilters != nil {
		in, out := &in.AdvancedFilters, &out.AdvancedFilters
		*out = make([]AdvancedFilter_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludedEventTypes != nil {
		in, out := &in.IncludedEventTypes, &out.IncludedEventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IsSubjectCaseSensitive != nil {
		in, out := &in.IsSubjectCaseSensitive, &out.IsSubjectCaseSensitive
		*out = new(bool)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubjectBeginsWith != nil {
		in, out := &in.SubjectBeginsWith, &out.SubjectBeginsWith
		*out = new(string)
		**out = **in
	}
	if in.SubjectEndsWith != nil {
		in, out := &in.SubjectEndsWith, &out.SubjectEndsWith
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionFilter_Status.
func (in *EventSubscriptionFilter_Status) DeepCopy() *EventSubscriptionFilter_Status {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionFilter_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptionList) DeepCopyInto(out *EventSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptionList.
func (in *EventSubscriptionList) DeepCopy() *EventSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscription_Status) DeepCopyInto(out *EventSubscription_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeadLetterDestination != nil {
		in, out := &in.DeadLetterDestination, &out.DeadLetterDestination
		*out = new(DeadLetterDestination_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(EventSubscriptionDestination_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.EventDeliverySchema != nil {
		in, out := &in.EventDeliverySchema, &out.EventDeliverySchema
		*out = new(string)
		**out = **in
	}
	if in.ExpirationTimeUtc != nil {
		in, out := &in.ExpirationTimeUtc, &out.ExpirationTimeUtc
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(EventSubscriptionFilter_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.RetryPolicy != nil {
		in, out := &in.RetryPolicy, &out.RetryPolicy
		*out = new(RetryPolicy_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscription_Status.
func (in *EventSubscription_Status) DeepCopy() *EventSubscription_Status {
	if in == nil {
		return nil
	}
	out := new(EventSubscription_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSubscriptions_SPEC) DeepCopyInto(out *EventSubscriptions_SPEC) {
	*out = *in
	if in.DeadLetterDestination != nil {
		in, out := &in.DeadLetterDestination, &out.DeadLetterDestination
		*out = new(DeadLetterDestination_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(EventSubscriptionDestination_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.EventDeliverySchema != nil {
		in, out := &in.EventDeliverySchema, &out.EventDeliverySchema
		*out = new(string)
		**out = **in
	}
	if in.ExpirationTimeUtc != nil {
		in, out := &in.ExpirationTimeUtc, &out.ExpirationTimeUtc
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(EventSubscriptionFilter_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Owner = in.Owner
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RetryPolicy != nil {
		in, out := &in.RetryPolicy, &out.RetryPolicy
		*out = new(RetryPolicy_Spec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSubscriptions_SPEC.
func (in *EventSubscriptions_SPEC) DeepCopy() *EventSubscriptions_SPEC {
	if in == nil {
		return nil
	}
	out := new(EventSubscriptions_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboundIpRule_Spec) DeepCopyInto(out *InboundIpRule_Spec) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.IpMask != nil {
		in, out := &in.IpMask, &out.IpMask
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboundIpRule_Spec.
func (in *InboundIpRule_Spec) DeepCopy() *InboundIpRule_Spec {
	if in == nil {
		return nil
	}
	out := new(InboundIpRule_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InboundIpRule_Status) DeepCopyInto(out *InboundIpRule_Status) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.IpMask != nil {
		in, out := &in.IpMask, &out.IpMask
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InboundIpRule_Status.
func (in *InboundIpRule_Status) DeepCopy() *InboundIpRule_Status {
	if in == nil {
		return nil
	}
	out := new(InboundIpRule_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputSchemaMapping_Spec) DeepCopyInto(out *InputSchemaMapping_Spec) {
	*out = *in
	if in.InputSchemaMappingType != nil {
		in, out := &in.InputSchemaMappingType, &out.InputSchemaMappingType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputSchemaMapping_Spec.
func (in *InputSchemaMapping_Spec) DeepCopy() *InputSchemaMapping_Spec {
	if in == nil {
		return nil
	}
	out := new(InputSchemaMapping_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputSchemaMapping_Status) DeepCopyInto(out *InputSchemaMapping_Status) {
	*out = *in
	if in.InputSchemaMappingType != nil {
		in, out := &in.InputSchemaMappingType, &out.InputSchemaMappingType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputSchemaMapping_Status.
func (in *InputSchemaMapping_Status) DeepCopy() *InputSchemaMapping_Status {
	if in == nil {
		return nil
	}
	out := new(InputSchemaMapping_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_Status_Domain_SubResourceEmbedded) DeepCopyInto(out *PrivateEndpointConnection_Status_Domain_SubResourceEmbedded) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_Status_Domain_SubResourceEmbedded.
func (in *PrivateEndpointConnection_Status_Domain_SubResourceEmbedded) DeepCopy() *PrivateEndpointConnection_Status_Domain_SubResourceEmbedded {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_Status_Domain_SubResourceEmbedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_Status_Topic_SubResourceEmbedded) DeepCopyInto(out *PrivateEndpointConnection_Status_Topic_SubResourceEmbedded) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_Status_Topic_SubResourceEmbedded.
func (in *PrivateEndpointConnection_Status_Topic_SubResourceEmbedded) DeepCopy() *PrivateEndpointConnection_Status_Topic_SubResourceEmbedded {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_Status_Topic_SubResourceEmbedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryPolicy_Spec) DeepCopyInto(out *RetryPolicy_Spec) {
	*out = *in
	if in.EventTimeToLiveInMinutes != nil {
		in, out := &in.EventTimeToLiveInMinutes, &out.EventTimeToLiveInMinutes
		*out = new(int)
		**out = **in
	}
	if in.MaxDeliveryAttempts != nil {
		in, out := &in.MaxDeliveryAttempts, &out.MaxDeliveryAttempts
		*out = new(int)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicy_Spec.
func (in *RetryPolicy_Spec) DeepCopy() *RetryPolicy_Spec {
	if in == nil {
		return nil
	}
	out := new(RetryPolicy_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryPolicy_Status) DeepCopyInto(out *RetryPolicy_Status) {
	*out = *in
	if in.EventTimeToLiveInMinutes != nil {
		in, out := &in.EventTimeToLiveInMinutes, &out.EventTimeToLiveInMinutes
		*out = new(int)
		**out = **in
	}
	if in.MaxDeliveryAttempts != nil {
		in, out := &in.MaxDeliveryAttempts, &out.MaxDeliveryAttempts
		*out = new(int)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryPolicy_Status.
func (in *RetryPolicy_Status) DeepCopy() *RetryPolicy_Status {
	if in == nil {
		return nil
	}
	out := new(RetryPolicy_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_Status) DeepCopyInto(out *SystemData_Status) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_Status.
func (in *SystemData_Status) DeepCopy() *SystemData_Status {
	if in == nil {
		return nil
	}
	out := new(SystemData_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topic) DeepCopyInto(out *Topic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topic.
func (in *Topic) DeepCopy() *Topic {
	if in == nil {
		return nil
	}
	out := new(Topic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Topic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicList) DeepCopyInto(out *TopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Topic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicList.
func (in *TopicList) DeepCopy() *TopicList {
	if in == nil {
		return nil
	}
	out := new(TopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topic_Status) DeepCopyInto(out *Topic_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.InboundIpRules != nil {
		in, out := &in.InboundIpRules, &out.InboundIpRules
		*out = make([]InboundIpRule_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InputSchema != nil {
		in, out := &in.InputSchema, &out.InputSchema
		*out = new(string)
		**out = **in
	}
	if in.InputSchemaMapping != nil {
		in, out := &in.InputSchemaMapping, &out.InputSchemaMapping
		*out = new(InputSchemaMapping_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MetricResourceId != nil {
		in, out := &in.MetricResourceId, &out.MetricResourceId
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_Status_Topic_SubResourceEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topic_Status.
func (in *Topic_Status) DeepCopy() *Topic_Status {
	if in == nil {
		return nil
	}
	out := new(Topic_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topics_SPEC) DeepCopyInto(out *Topics_SPEC) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	out.Owner = in.Owner
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topics_SPEC.
func (in *Topics_SPEC) DeepCopy() *Topics_SPEC {
	if in == nil {
		return nil
	}
	out := new(Topics_SPEC)
	in.DeepCopyInto(out)
	return out
}