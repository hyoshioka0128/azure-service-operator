// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=mongodbdatabasecollectionthroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={mongodbdatabasecollectionthroughputsettings/status,mongodbdatabasecollectionthroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210515.MongodbDatabaseCollectionThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
type MongodbDatabaseCollectionThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseCollectionThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *MongodbDatabaseCollectionThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *MongodbDatabaseCollectionThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseCollectionThroughputSetting{}

// AzureName returns the Azure name of the resource
func (setting *MongodbDatabaseCollectionThroughputSetting) AzureName() string {
	return setting.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting MongodbDatabaseCollectionThroughputSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (setting *MongodbDatabaseCollectionThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (setting *MongodbDatabaseCollectionThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *MongodbDatabaseCollectionThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (setting *MongodbDatabaseCollectionThroughputSetting) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *MongodbDatabaseCollectionThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (setting *MongodbDatabaseCollectionThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *MongodbDatabaseCollectionThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// Hub marks that this MongodbDatabaseCollectionThroughputSetting is the hub type for conversion
func (setting *MongodbDatabaseCollectionThroughputSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *MongodbDatabaseCollectionThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseCollectionThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210515.MongodbDatabaseCollectionThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
type MongodbDatabaseCollectionThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseCollectionThroughputSetting `json:"items"`
}

// Storage version of v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS
type DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS struct {
	Conditions  []conditions.Condition             `json:"conditions,omitempty"`
	Id          *string                            `json:"id,omitempty"`
	Location    *string                            `json:"location,omitempty"`
	Name        *string                            `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Type        *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS{}

// ConvertStatusFrom populates our DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS from the provided source
func (setting *DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(setting)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS
func (setting *DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(setting)
}

// Storage version of v1beta20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec
type DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Id              *string `json:"id,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Type        *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec from the provided source
func (setting *DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(setting)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec
func (setting *DatabaseAccountsMongodbDatabasesCollectionsThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(setting)
}

// Storage version of v1beta20210515.ThroughputSettingsResource
type ThroughputSettingsResource struct {
	AutoscaleSettings   *AutoscaleSettingsResource `json:"autoscaleSettings,omitempty"`
	MinimumThroughput   *string                    `json:"minimumThroughput,omitempty"`
	OfferReplacePending *string                    `json:"offerReplacePending,omitempty"`
	PropertyBag         genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Throughput          *int                       `json:"throughput,omitempty"`
}

// Storage version of v1beta20210515.ThroughputSettingsResource_STATUS
type ThroughputSettingsResource_STATUS struct {
	AutoscaleSettings   *AutoscaleSettingsResource_STATUS `json:"autoscaleSettings,omitempty"`
	MinimumThroughput   *string                           `json:"minimumThroughput,omitempty"`
	OfferReplacePending *string                           `json:"offerReplacePending,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Throughput          *int                              `json:"throughput,omitempty"`
}

// Storage version of v1beta20210515.AutoscaleSettingsResource
type AutoscaleSettingsResource struct {
	AutoUpgradePolicy   *AutoUpgradePolicyResource `json:"autoUpgradePolicy,omitempty"`
	MaxThroughput       *int                       `json:"maxThroughput,omitempty"`
	PropertyBag         genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	TargetMaxThroughput *int                       `json:"targetMaxThroughput,omitempty"`
}

// Storage version of v1beta20210515.AutoscaleSettingsResource_STATUS
type AutoscaleSettingsResource_STATUS struct {
	AutoUpgradePolicy   *AutoUpgradePolicyResource_STATUS `json:"autoUpgradePolicy,omitempty"`
	MaxThroughput       *int                              `json:"maxThroughput,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	TargetMaxThroughput *int                              `json:"targetMaxThroughput,omitempty"`
}

// Storage version of v1beta20210515.AutoUpgradePolicyResource
type AutoUpgradePolicyResource struct {
	PropertyBag      genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	ThroughputPolicy *ThroughputPolicyResource `json:"throughputPolicy,omitempty"`
}

// Storage version of v1beta20210515.AutoUpgradePolicyResource_STATUS
type AutoUpgradePolicyResource_STATUS struct {
	PropertyBag      genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ThroughputPolicy *ThroughputPolicyResource_STATUS `json:"throughputPolicy,omitempty"`
}

// Storage version of v1beta20210515.ThroughputPolicyResource
type ThroughputPolicyResource struct {
	IncrementPercent *int                   `json:"incrementPercent,omitempty"`
	IsEnabled        *bool                  `json:"isEnabled,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210515.ThroughputPolicyResource_STATUS
type ThroughputPolicyResource_STATUS struct {
	IncrementPercent *int                   `json:"incrementPercent,omitempty"`
	IsEnabled        *bool                  `json:"isEnabled,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseCollectionThroughputSetting{}, &MongodbDatabaseCollectionThroughputSettingList{})
}
