//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20210401previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyEntry) DeepCopyInto(out *AccessPolicyEntry) {
	*out = *in
	if in.ApplicationId != nil {
		in, out := &in.ApplicationId, &out.ApplicationId
		*out = new(string)
		**out = **in
	}
	if in.ObjectId != nil {
		in, out := &in.ObjectId, &out.ObjectId
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(Permissions)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyEntry.
func (in *AccessPolicyEntry) DeepCopy() *AccessPolicyEntry {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyEntry_STATUS) DeepCopyInto(out *AccessPolicyEntry_STATUS) {
	*out = *in
	if in.ApplicationId != nil {
		in, out := &in.ApplicationId, &out.ApplicationId
		*out = new(string)
		**out = **in
	}
	if in.ObjectId != nil {
		in, out := &in.ObjectId, &out.ObjectId
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(Permissions_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyEntry_STATUS.
func (in *AccessPolicyEntry_STATUS) DeepCopy() *AccessPolicyEntry_STATUS {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyEntry_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRule) DeepCopyInto(out *IPRule) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRule.
func (in *IPRule) DeepCopy() *IPRule {
	if in == nil {
		return nil
	}
	out := new(IPRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPRule_STATUS) DeepCopyInto(out *IPRule_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPRule_STATUS.
func (in *IPRule_STATUS) DeepCopy() *IPRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(IPRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRuleSet) DeepCopyInto(out *NetworkRuleSet) {
	*out = *in
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(string)
		**out = **in
	}
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IPRule, len(*in))
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
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRuleSet.
func (in *NetworkRuleSet) DeepCopy() *NetworkRuleSet {
	if in == nil {
		return nil
	}
	out := new(NetworkRuleSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRuleSet_STATUS) DeepCopyInto(out *NetworkRuleSet_STATUS) {
	*out = *in
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(string)
		**out = **in
	}
	if in.DefaultAction != nil {
		in, out := &in.DefaultAction, &out.DefaultAction
		*out = new(string)
		**out = **in
	}
	if in.IpRules != nil {
		in, out := &in.IpRules, &out.IpRules
		*out = make([]IPRule_STATUS, len(*in))
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
	if in.VirtualNetworkRules != nil {
		in, out := &in.VirtualNetworkRules, &out.VirtualNetworkRules
		*out = make([]VirtualNetworkRule_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRuleSet_STATUS.
func (in *NetworkRuleSet_STATUS) DeepCopy() *NetworkRuleSet_STATUS {
	if in == nil {
		return nil
	}
	out := new(NetworkRuleSet_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permissions) DeepCopyInto(out *Permissions) {
	*out = *in
	if in.Certificates != nil {
		in, out := &in.Certificates, &out.Certificates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permissions.
func (in *Permissions) DeepCopy() *Permissions {
	if in == nil {
		return nil
	}
	out := new(Permissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permissions_STATUS) DeepCopyInto(out *Permissions_STATUS) {
	*out = *in
	if in.Certificates != nil {
		in, out := &in.Certificates, &out.Certificates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permissions_STATUS.
func (in *Permissions_STATUS) DeepCopy() *Permissions_STATUS {
	if in == nil {
		return nil
	}
	out := new(Permissions_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint) DeepCopyInto(out *PrivateEndpoint) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint.
func (in *PrivateEndpoint) DeepCopy() *PrivateEndpoint {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnectionItem) DeepCopyInto(out *PrivateEndpointConnectionItem) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionState)
		(*in).DeepCopyInto(*out)
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnectionItem.
func (in *PrivateEndpointConnectionItem) DeepCopy() *PrivateEndpointConnectionItem {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnectionItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnectionItem_STATUS) DeepCopyInto(out *PrivateEndpointConnectionItem_STATUS) {
	*out = *in
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(PrivateEndpoint_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateLinkServiceConnectionState != nil {
		in, out := &in.PrivateLinkServiceConnectionState, &out.PrivateLinkServiceConnectionState
		*out = new(PrivateLinkServiceConnectionState_STATUS)
		(*in).DeepCopyInto(*out)
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnectionItem_STATUS.
func (in *PrivateEndpointConnectionItem_STATUS) DeepCopy() *PrivateEndpointConnectionItem_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnectionItem_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint_STATUS) DeepCopyInto(out *PrivateEndpoint_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint_STATUS.
func (in *PrivateEndpoint_STATUS) DeepCopy() *PrivateEndpoint_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionState) DeepCopyInto(out *PrivateLinkServiceConnectionState) {
	*out = *in
	if in.ActionsRequired != nil {
		in, out := &in.ActionsRequired, &out.ActionsRequired
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionState.
func (in *PrivateLinkServiceConnectionState) DeepCopy() *PrivateLinkServiceConnectionState {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateLinkServiceConnectionState_STATUS) DeepCopyInto(out *PrivateLinkServiceConnectionState_STATUS) {
	*out = *in
	if in.ActionsRequired != nil {
		in, out := &in.ActionsRequired, &out.ActionsRequired
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateLinkServiceConnectionState_STATUS.
func (in *PrivateLinkServiceConnectionState_STATUS) DeepCopy() *PrivateLinkServiceConnectionState_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrivateLinkServiceConnectionState_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku.
func (in *Sku) DeepCopy() *Sku {
	if in == nil {
		return nil
	}
	out := new(Sku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS) DeepCopyInto(out *Sku_STATUS) {
	*out = *in
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS.
func (in *Sku_STATUS) DeepCopy() *Sku_STATUS {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS) DeepCopyInto(out *SystemData_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vault) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultList) DeepCopyInto(out *VaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vault, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultList.
func (in *VaultList) DeepCopy() *VaultList {
	if in == nil {
		return nil
	}
	out := new(VaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultProperties) DeepCopyInto(out *VaultProperties) {
	*out = *in
	if in.AccessPolicies != nil {
		in, out := &in.AccessPolicies, &out.AccessPolicies
		*out = make([]AccessPolicyEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.EnablePurgeProtection != nil {
		in, out := &in.EnablePurgeProtection, &out.EnablePurgeProtection
		*out = new(bool)
		**out = **in
	}
	if in.EnableRbacAuthorization != nil {
		in, out := &in.EnableRbacAuthorization, &out.EnableRbacAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.EnableSoftDelete != nil {
		in, out := &in.EnableSoftDelete, &out.EnableSoftDelete
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDeployment != nil {
		in, out := &in.EnabledForDeployment, &out.EnabledForDeployment
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDiskEncryption != nil {
		in, out := &in.EnabledForDiskEncryption, &out.EnabledForDiskEncryption
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForTemplateDeployment != nil {
		in, out := &in.EnabledForTemplateDeployment, &out.EnabledForTemplateDeployment
		*out = new(bool)
		**out = **in
	}
	if in.HsmPoolResourceId != nil {
		in, out := &in.HsmPoolResourceId, &out.HsmPoolResourceId
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = new(NetworkRuleSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnectionItem, len(*in))
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
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.SoftDeleteRetentionInDays != nil {
		in, out := &in.SoftDeleteRetentionInDays, &out.SoftDeleteRetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.VaultUri != nil {
		in, out := &in.VaultUri, &out.VaultUri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultProperties.
func (in *VaultProperties) DeepCopy() *VaultProperties {
	if in == nil {
		return nil
	}
	out := new(VaultProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultProperties_STATUS) DeepCopyInto(out *VaultProperties_STATUS) {
	*out = *in
	if in.AccessPolicies != nil {
		in, out := &in.AccessPolicies, &out.AccessPolicies
		*out = make([]AccessPolicyEntry_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.EnablePurgeProtection != nil {
		in, out := &in.EnablePurgeProtection, &out.EnablePurgeProtection
		*out = new(bool)
		**out = **in
	}
	if in.EnableRbacAuthorization != nil {
		in, out := &in.EnableRbacAuthorization, &out.EnableRbacAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.EnableSoftDelete != nil {
		in, out := &in.EnableSoftDelete, &out.EnableSoftDelete
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDeployment != nil {
		in, out := &in.EnabledForDeployment, &out.EnabledForDeployment
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForDiskEncryption != nil {
		in, out := &in.EnabledForDiskEncryption, &out.EnabledForDiskEncryption
		*out = new(bool)
		**out = **in
	}
	if in.EnabledForTemplateDeployment != nil {
		in, out := &in.EnabledForTemplateDeployment, &out.EnabledForTemplateDeployment
		*out = new(bool)
		**out = **in
	}
	if in.HsmPoolResourceId != nil {
		in, out := &in.HsmPoolResourceId, &out.HsmPoolResourceId
		*out = new(string)
		**out = **in
	}
	if in.NetworkAcls != nil {
		in, out := &in.NetworkAcls, &out.NetworkAcls
		*out = new(NetworkRuleSet_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnectionItem_STATUS, len(*in))
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
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SoftDeleteRetentionInDays != nil {
		in, out := &in.SoftDeleteRetentionInDays, &out.SoftDeleteRetentionInDays
		*out = new(int)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.VaultUri != nil {
		in, out := &in.VaultUri, &out.VaultUri
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultProperties_STATUS.
func (in *VaultProperties_STATUS) DeepCopy() *VaultProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(VaultProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault_STATUS) DeepCopyInto(out *Vault_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
		*out = new(VaultProperties_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault_STATUS.
func (in *Vault_STATUS) DeepCopy() *Vault_STATUS {
	if in == nil {
		return nil
	}
	out := new(Vault_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault_Spec) DeepCopyInto(out *Vault_Spec) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(VaultProperties)
		(*in).DeepCopyInto(*out)
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault_Spec.
func (in *Vault_Spec) DeepCopy() *Vault_Spec {
	if in == nil {
		return nil
	}
	out := new(Vault_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRule) DeepCopyInto(out *VirtualNetworkRule) {
	*out = *in
	if in.IgnoreMissingVnetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVnetServiceEndpoint, &out.IgnoreMissingVnetServiceEndpoint
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
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRule.
func (in *VirtualNetworkRule) DeepCopy() *VirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRule_STATUS) DeepCopyInto(out *VirtualNetworkRule_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.IgnoreMissingVnetServiceEndpoint != nil {
		in, out := &in.IgnoreMissingVnetServiceEndpoint, &out.IgnoreMissingVnetServiceEndpoint
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRule_STATUS.
func (in *VirtualNetworkRule_STATUS) DeepCopy() *VirtualNetworkRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRule_STATUS)
	in.DeepCopyInto(out)
	return out
}
