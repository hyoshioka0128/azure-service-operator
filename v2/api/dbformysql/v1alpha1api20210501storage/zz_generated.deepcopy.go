//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1api20210501storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup_Spec) DeepCopyInto(out *Backup_Spec) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int)
		**out = **in
	}
	if in.GeoRedundantBackup != nil {
		in, out := &in.GeoRedundantBackup, &out.GeoRedundantBackup
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup_Spec.
func (in *Backup_Spec) DeepCopy() *Backup_Spec {
	if in == nil {
		return nil
	}
	out := new(Backup_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup_Status) DeepCopyInto(out *Backup_Status) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int)
		**out = **in
	}
	if in.EarliestRestoreDate != nil {
		in, out := &in.EarliestRestoreDate, &out.EarliestRestoreDate
		*out = new(string)
		**out = **in
	}
	if in.GeoRedundantBackup != nil {
		in, out := &in.GeoRedundantBackup, &out.GeoRedundantBackup
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup_Status.
func (in *Backup_Status) DeepCopy() *Backup_Status {
	if in == nil {
		return nil
	}
	out := new(Backup_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database_Status) DeepCopyInto(out *Database_Status) {
	*out = *in
	if in.Charset != nil {
		in, out := &in.Charset, &out.Charset
		*out = new(string)
		**out = **in
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
		*out = new(string)
		**out = **in
	}
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
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database_Status.
func (in *Database_Status) DeepCopy() *Database_Status {
	if in == nil {
		return nil
	}
	out := new(Database_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRule_Status) DeepCopyInto(out *FirewallRule_Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndIpAddress != nil {
		in, out := &in.EndIpAddress, &out.EndIpAddress
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
	if in.StartIpAddress != nil {
		in, out := &in.StartIpAddress, &out.StartIpAddress
		*out = new(string)
		**out = **in
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRule_Status.
func (in *FirewallRule_Status) DeepCopy() *FirewallRule_Status {
	if in == nil {
		return nil
	}
	out := new(FirewallRule_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServer) DeepCopyInto(out *FlexibleServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServer.
func (in *FlexibleServer) DeepCopy() *FlexibleServer {
	if in == nil {
		return nil
	}
	out := new(FlexibleServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServerList) DeepCopyInto(out *FlexibleServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServerList.
func (in *FlexibleServerList) DeepCopy() *FlexibleServerList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersDatabase) DeepCopyInto(out *FlexibleServersDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersDatabase.
func (in *FlexibleServersDatabase) DeepCopy() *FlexibleServersDatabase {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersDatabaseList) DeepCopyInto(out *FlexibleServersDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServersDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersDatabaseList.
func (in *FlexibleServersDatabaseList) DeepCopy() *FlexibleServersDatabaseList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersDatabases_SPEC) DeepCopyInto(out *FlexibleServersDatabases_SPEC) {
	*out = *in
	if in.Charset != nil {
		in, out := &in.Charset, &out.Charset
		*out = new(string)
		**out = **in
	}
	if in.Collation != nil {
		in, out := &in.Collation, &out.Collation
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersDatabases_SPEC.
func (in *FlexibleServersDatabases_SPEC) DeepCopy() *FlexibleServersDatabases_SPEC {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersDatabases_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersFirewallRule) DeepCopyInto(out *FlexibleServersFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersFirewallRule.
func (in *FlexibleServersFirewallRule) DeepCopy() *FlexibleServersFirewallRule {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersFirewallRuleList) DeepCopyInto(out *FlexibleServersFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlexibleServersFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersFirewallRuleList.
func (in *FlexibleServersFirewallRuleList) DeepCopy() *FlexibleServersFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlexibleServersFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServersFirewallRules_SPEC) DeepCopyInto(out *FlexibleServersFirewallRules_SPEC) {
	*out = *in
	if in.EndIpAddress != nil {
		in, out := &in.EndIpAddress, &out.EndIpAddress
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
	if in.StartIpAddress != nil {
		in, out := &in.StartIpAddress, &out.StartIpAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServersFirewallRules_SPEC.
func (in *FlexibleServersFirewallRules_SPEC) DeepCopy() *FlexibleServersFirewallRules_SPEC {
	if in == nil {
		return nil
	}
	out := new(FlexibleServersFirewallRules_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlexibleServers_SPEC) DeepCopyInto(out *FlexibleServers_SPEC) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AdministratorLoginPassword != nil {
		in, out := &in.AdministratorLoginPassword, &out.AdministratorLoginPassword
		*out = new(genruntime.SecretReference)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(Backup_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailability_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindow_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(Network_Spec)
		(*in).DeepCopyInto(*out)
	}
	out.Owner = in.Owner
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(string)
		**out = **in
	}
	if in.RestorePointInTime != nil {
		in, out := &in.RestorePointInTime, &out.RestorePointInTime
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceServerResourceId != nil {
		in, out := &in.SourceServerResourceId, &out.SourceServerResourceId
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage_Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlexibleServers_SPEC.
func (in *FlexibleServers_SPEC) DeepCopy() *FlexibleServers_SPEC {
	if in == nil {
		return nil
	}
	out := new(FlexibleServers_SPEC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailability_Spec) DeepCopyInto(out *HighAvailability_Spec) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
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
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailability_Spec.
func (in *HighAvailability_Spec) DeepCopy() *HighAvailability_Spec {
	if in == nil {
		return nil
	}
	out := new(HighAvailability_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HighAvailability_Status) DeepCopyInto(out *HighAvailability_Status) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
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
	if in.StandbyAvailabilityZone != nil {
		in, out := &in.StandbyAvailabilityZone, &out.StandbyAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HighAvailability_Status.
func (in *HighAvailability_Status) DeepCopy() *HighAvailability_Status {
	if in == nil {
		return nil
	}
	out := new(HighAvailability_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindow_Spec) DeepCopyInto(out *MaintenanceWindow_Spec) {
	*out = *in
	if in.CustomWindow != nil {
		in, out := &in.CustomWindow, &out.CustomWindow
		*out = new(string)
		**out = **in
	}
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
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
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(int)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindow_Spec.
func (in *MaintenanceWindow_Spec) DeepCopy() *MaintenanceWindow_Spec {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindow_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindow_Status) DeepCopyInto(out *MaintenanceWindow_Status) {
	*out = *in
	if in.CustomWindow != nil {
		in, out := &in.CustomWindow, &out.CustomWindow
		*out = new(string)
		**out = **in
	}
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
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
	if in.StartHour != nil {
		in, out := &in.StartHour, &out.StartHour
		*out = new(int)
		**out = **in
	}
	if in.StartMinute != nil {
		in, out := &in.StartMinute, &out.StartMinute
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindow_Status.
func (in *MaintenanceWindow_Status) DeepCopy() *MaintenanceWindow_Status {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindow_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network_Spec) DeepCopyInto(out *Network_Spec) {
	*out = *in
	if in.DelegatedSubnetResourceReference != nil {
		in, out := &in.DelegatedSubnetResourceReference, &out.DelegatedSubnetResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.PrivateDnsZoneResourceReference != nil {
		in, out := &in.PrivateDnsZoneResourceReference, &out.PrivateDnsZoneResourceReference
		*out = new(genruntime.ResourceReference)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network_Spec.
func (in *Network_Spec) DeepCopy() *Network_Spec {
	if in == nil {
		return nil
	}
	out := new(Network_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network_Status) DeepCopyInto(out *Network_Status) {
	*out = *in
	if in.DelegatedSubnetResourceId != nil {
		in, out := &in.DelegatedSubnetResourceId, &out.DelegatedSubnetResourceId
		*out = new(string)
		**out = **in
	}
	if in.PrivateDnsZoneResourceId != nil {
		in, out := &in.PrivateDnsZoneResourceId, &out.PrivateDnsZoneResourceId
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network_Status.
func (in *Network_Status) DeepCopy() *Network_Status {
	if in == nil {
		return nil
	}
	out := new(Network_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server_Status) DeepCopyInto(out *Server_Status) {
	*out = *in
	if in.AdministratorLogin != nil {
		in, out := &in.AdministratorLogin, &out.AdministratorLogin
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(Backup_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateMode != nil {
		in, out := &in.CreateMode, &out.CreateMode
		*out = new(string)
		**out = **in
	}
	if in.FullyQualifiedDomainName != nil {
		in, out := &in.FullyQualifiedDomainName, &out.FullyQualifiedDomainName
		*out = new(string)
		**out = **in
	}
	if in.HighAvailability != nil {
		in, out := &in.HighAvailability, &out.HighAvailability
		*out = new(HighAvailability_Status)
		(*in).DeepCopyInto(*out)
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
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindow_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(Network_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReplicaCapacity != nil {
		in, out := &in.ReplicaCapacity, &out.ReplicaCapacity
		*out = new(int)
		**out = **in
	}
	if in.ReplicationRole != nil {
		in, out := &in.ReplicationRole, &out.ReplicationRole
		*out = new(string)
		**out = **in
	}
	if in.RestorePointInTime != nil {
		in, out := &in.RestorePointInTime, &out.RestorePointInTime
		*out = new(string)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_Status)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceServerResourceId != nil {
		in, out := &in.SourceServerResourceId, &out.SourceServerResourceId
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage_Status)
		(*in).DeepCopyInto(*out)
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server_Status.
func (in *Server_Status) DeepCopy() *Server_Status {
	if in == nil {
		return nil
	}
	out := new(Server_Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_Spec) DeepCopyInto(out *Sku_Spec) {
	*out = *in
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
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
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
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage_Spec) DeepCopyInto(out *Storage_Spec) {
	*out = *in
	if in.AutoGrow != nil {
		in, out := &in.AutoGrow, &out.AutoGrow
		*out = new(string)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
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
	if in.StorageSizeGB != nil {
		in, out := &in.StorageSizeGB, &out.StorageSizeGB
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage_Spec.
func (in *Storage_Spec) DeepCopy() *Storage_Spec {
	if in == nil {
		return nil
	}
	out := new(Storage_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage_Status) DeepCopyInto(out *Storage_Status) {
	*out = *in
	if in.AutoGrow != nil {
		in, out := &in.AutoGrow, &out.AutoGrow
		*out = new(string)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
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
	if in.StorageSizeGB != nil {
		in, out := &in.StorageSizeGB, &out.StorageSizeGB
		*out = new(int)
		**out = **in
	}
	if in.StorageSku != nil {
		in, out := &in.StorageSku, &out.StorageSku
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage_Status.
func (in *Storage_Status) DeepCopy() *Storage_Status {
	if in == nil {
		return nil
	}
	out := new(Storage_Status)
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
