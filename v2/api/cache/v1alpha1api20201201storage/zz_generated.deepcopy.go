//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20201201storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Redis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCreateParameters_Status) DeepCopyInto(out *RedisCreateParameters_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableNonSslPort != nil {
		in, out := &in.EnableNonSslPort, &out.EnableNonSslPort
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
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
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisCreateProperties_Status_RedisConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(int)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(int)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticIP != nil {
		in, out := &in.StaticIP, &out.StaticIP
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
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
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCreateParameters_Status.
func (in *RedisCreateParameters_Status) DeepCopy() *RedisCreateParameters_Status {
	if in == nil {
		return nil
	}
	out := new(RedisCreateParameters_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCreateProperties_Spec_RedisConfiguration) DeepCopyInto(out *RedisCreateProperties_Spec_RedisConfiguration) {
	*out = *in
	if in.AofStorageConnectionString0 != nil {
		in, out := &in.AofStorageConnectionString0, &out.AofStorageConnectionString0
		*out = new(string)
		**out = **in
	}
	if in.AofStorageConnectionString1 != nil {
		in, out := &in.AofStorageConnectionString1, &out.AofStorageConnectionString1
		*out = new(string)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
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
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(string)
		**out = **in
	}
	if in.RdbStorageConnectionString != nil {
		in, out := &in.RdbStorageConnectionString, &out.RdbStorageConnectionString
		*out = new(string)
		**out = **in
	}
	if in.additionalProperties != nil {
		in, out := &in.additionalProperties, &out.additionalProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCreateProperties_Spec_RedisConfiguration.
func (in *RedisCreateProperties_Spec_RedisConfiguration) DeepCopy() *RedisCreateProperties_Spec_RedisConfiguration {
	if in == nil {
		return nil
	}
	out := new(RedisCreateProperties_Spec_RedisConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCreateProperties_Status_RedisConfiguration) DeepCopyInto(out *RedisCreateProperties_Status_RedisConfiguration) {
	*out = *in
	if in.AofStorageConnectionString0 != nil {
		in, out := &in.AofStorageConnectionString0, &out.AofStorageConnectionString0
		*out = new(string)
		**out = **in
	}
	if in.AofStorageConnectionString1 != nil {
		in, out := &in.AofStorageConnectionString1, &out.AofStorageConnectionString1
		*out = new(string)
		**out = **in
	}
	if in.Maxclients != nil {
		in, out := &in.Maxclients, &out.Maxclients
		*out = new(string)
		**out = **in
	}
	if in.MaxfragmentationmemoryReserved != nil {
		in, out := &in.MaxfragmentationmemoryReserved, &out.MaxfragmentationmemoryReserved
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryDelta != nil {
		in, out := &in.MaxmemoryDelta, &out.MaxmemoryDelta
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryPolicy != nil {
		in, out := &in.MaxmemoryPolicy, &out.MaxmemoryPolicy
		*out = new(string)
		**out = **in
	}
	if in.MaxmemoryReserved != nil {
		in, out := &in.MaxmemoryReserved, &out.MaxmemoryReserved
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
	if in.RdbBackupEnabled != nil {
		in, out := &in.RdbBackupEnabled, &out.RdbBackupEnabled
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupFrequency != nil {
		in, out := &in.RdbBackupFrequency, &out.RdbBackupFrequency
		*out = new(string)
		**out = **in
	}
	if in.RdbBackupMaxSnapshotCount != nil {
		in, out := &in.RdbBackupMaxSnapshotCount, &out.RdbBackupMaxSnapshotCount
		*out = new(string)
		**out = **in
	}
	if in.RdbStorageConnectionString != nil {
		in, out := &in.RdbStorageConnectionString, &out.RdbStorageConnectionString
		*out = new(string)
		**out = **in
	}
	if in.additionalProperties != nil {
		in, out := &in.additionalProperties, &out.additionalProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCreateProperties_Status_RedisConfiguration.
func (in *RedisCreateProperties_Status_RedisConfiguration) DeepCopy() *RedisCreateProperties_Status_RedisConfiguration {
	if in == nil {
		return nil
	}
	out := new(RedisCreateProperties_Status_RedisConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRule) DeepCopyInto(out *RedisFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRule.
func (in *RedisFirewallRule) DeepCopy() *RedisFirewallRule {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRuleList) DeepCopyInto(out *RedisFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRuleList.
func (in *RedisFirewallRuleList) DeepCopy() *RedisFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRule_Status) DeepCopyInto(out *RedisFirewallRule_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndIP != nil {
		in, out := &in.EndIP, &out.EndIP
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
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
	if in.StartIP != nil {
		in, out := &in.StartIP, &out.StartIP
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRule_Status.
func (in *RedisFirewallRule_Status) DeepCopy() *RedisFirewallRule_Status {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRule_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisFirewallRules_SPEC) DeepCopyInto(out *RedisFirewallRules_SPEC) {
	*out = *in
	if in.EndIP != nil {
		in, out := &in.EndIP, &out.EndIP
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
	if in.StartIP != nil {
		in, out := &in.StartIP, &out.StartIP
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisFirewallRules_SPEC.
func (in *RedisFirewallRules_SPEC) DeepCopy() *RedisFirewallRules_SPEC {
	if in == nil {
		return nil
	}
	out := new(RedisFirewallRules_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServer) DeepCopyInto(out *RedisLinkedServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServer.
func (in *RedisLinkedServer) DeepCopy() *RedisLinkedServer {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisLinkedServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerCreateParameters_Status) DeepCopyInto(out *RedisLinkedServerCreateParameters_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedRedisCacheId != nil {
		in, out := &in.LinkedRedisCacheId, &out.LinkedRedisCacheId
		*out = new(string)
		**out = **in
	}
	if in.LinkedRedisCacheLocation != nil {
		in, out := &in.LinkedRedisCacheLocation, &out.LinkedRedisCacheLocation
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
	if in.ServerRole != nil {
		in, out := &in.ServerRole, &out.ServerRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerCreateParameters_Status.
func (in *RedisLinkedServerCreateParameters_Status) DeepCopy() *RedisLinkedServerCreateParameters_Status {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerCreateParameters_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServerList) DeepCopyInto(out *RedisLinkedServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisLinkedServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServerList.
func (in *RedisLinkedServerList) DeepCopy() *RedisLinkedServerList {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisLinkedServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisLinkedServers_SPEC) DeepCopyInto(out *RedisLinkedServers_SPEC) {
	*out = *in
	if in.LinkedRedisCacheLocation != nil {
		in, out := &in.LinkedRedisCacheLocation, &out.LinkedRedisCacheLocation
		*out = new(string)
		**out = **in
	}
	out.LinkedRedisCacheReference = in.LinkedRedisCacheReference
	out.Owner = in.Owner
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServerRole != nil {
		in, out := &in.ServerRole, &out.ServerRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisLinkedServers_SPEC.
func (in *RedisLinkedServers_SPEC) DeepCopy() *RedisLinkedServers_SPEC {
	if in == nil {
		return nil
	}
	out := new(RedisLinkedServers_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisList) DeepCopyInto(out *RedisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Redis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisList.
func (in *RedisList) DeepCopy() *RedisList {
	if in == nil {
		return nil
	}
	out := new(RedisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchSchedule) DeepCopyInto(out *RedisPatchSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchSchedule.
func (in *RedisPatchSchedule) DeepCopy() *RedisPatchSchedule {
	if in == nil {
		return nil
	}
	out := new(RedisPatchSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisPatchSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchScheduleList) DeepCopyInto(out *RedisPatchScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisPatchSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchScheduleList.
func (in *RedisPatchScheduleList) DeepCopy() *RedisPatchScheduleList {
	if in == nil {
		return nil
	}
	out := new(RedisPatchScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisPatchScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchSchedule_Status) DeepCopyInto(out *RedisPatchSchedule_Status) {
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
	if in.ScheduleEntries != nil {
		in, out := &in.ScheduleEntries, &out.ScheduleEntries
		*out = make([]ScheduleEntry_Status, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchSchedule_Status.
func (in *RedisPatchSchedule_Status) DeepCopy() *RedisPatchSchedule_Status {
	if in == nil {
		return nil
	}
	out := new(RedisPatchSchedule_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisPatchSchedules_SPEC) DeepCopyInto(out *RedisPatchSchedules_SPEC) {
	*out = *in
	out.Owner = in.Owner
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ScheduleEntries != nil {
		in, out := &in.ScheduleEntries, &out.ScheduleEntries
		*out = make([]ScheduleEntry_Spec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisPatchSchedules_SPEC.
func (in *RedisPatchSchedules_SPEC) DeepCopy() *RedisPatchSchedules_SPEC {
	if in == nil {
		return nil
	}
	out := new(RedisPatchSchedules_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis_SPEC) DeepCopyInto(out *Redis_SPEC) {
	*out = *in
	if in.EnableNonSslPort != nil {
		in, out := &in.EnableNonSslPort, &out.EnableNonSslPort
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
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
	if in.RedisConfiguration != nil {
		in, out := &in.RedisConfiguration, &out.RedisConfiguration
		*out = new(RedisCreateProperties_Spec_RedisConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ReplicasPerMaster != nil {
		in, out := &in.ReplicasPerMaster, &out.ReplicasPerMaster
		*out = new(int)
		**out = **in
	}
	if in.ReplicasPerPrimary != nil {
		in, out := &in.ReplicasPerPrimary, &out.ReplicasPerPrimary
		*out = new(int)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticIP != nil {
		in, out := &in.StaticIP, &out.StaticIP
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
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
	if in.TenantSettings != nil {
		in, out := &in.TenantSettings, &out.TenantSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis_SPEC.
func (in *Redis_SPEC) DeepCopy() *Redis_SPEC {
	if in == nil {
		return nil
	}
	out := new(Redis_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleEntry_Spec) DeepCopyInto(out *ScheduleEntry_Spec) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
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
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleEntry_Spec.
func (in *ScheduleEntry_Spec) DeepCopy() *ScheduleEntry_Spec {
	if in == nil {
		return nil
	}
	out := new(ScheduleEntry_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleEntry_Status) DeepCopyInto(out *ScheduleEntry_Status) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
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
	if in.StartHourUtc != nil {
		in, out := &in.StartHourUtc, &out.StartHourUtc
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleEntry_Status.
func (in *ScheduleEntry_Status) DeepCopy() *ScheduleEntry_Status {
	if in == nil {
		return nil
	}
	out := new(ScheduleEntry_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_Spec) DeepCopyInto(out *Sku_Spec) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_Spec.
func (in *Sku_Spec) DeepCopy() *Sku_Spec {
	if in == nil {
		return nil
	}
	out := new(Sku_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_Status) DeepCopyInto(out *Sku_Status) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_Status.
func (in *Sku_Status) DeepCopy() *Sku_Status {
	if in == nil {
		return nil
	}
	out := new(Sku_Status)
	in.DeepCopyInto(out)
	return out
}
