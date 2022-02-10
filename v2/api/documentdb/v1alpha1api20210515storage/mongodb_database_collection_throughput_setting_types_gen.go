// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

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
//Storage version of v1alpha1api20210515.MongodbDatabaseCollectionThroughputSetting
type MongodbDatabaseCollectionThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC `json:"spec,omitempty"`
	Status            ThroughputSettingsUpdateParameters_Status                          `json:"status,omitempty"`
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
	return "2021-05-15"
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
	return &ThroughputSettingsUpdateParameters_Status{}
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
	if st, ok := status.(*ThroughputSettingsUpdateParameters_Status); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsUpdateParameters_Status
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
//Storage version of v1alpha1api20210515.MongodbDatabaseCollectionThroughputSetting
type MongodbDatabaseCollectionThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseCollectionThroughputSetting `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC
type DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource_Spec  `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC from the provided source
func (spec *DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC
func (spec *DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210515.ThroughputSettingsUpdateParameters_Status
type ThroughputSettingsUpdateParameters_Status struct {
	Conditions  []conditions.Condition             `json:"conditions,omitempty"`
	Id          *string                            `json:"id,omitempty"`
	Location    *string                            `json:"location,omitempty"`
	Name        *string                            `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource_Status `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Type        *string                            `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ThroughputSettingsUpdateParameters_Status{}

// ConvertStatusFrom populates our ThroughputSettingsUpdateParameters_Status from the provided source
func (parameters *ThroughputSettingsUpdateParameters_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(parameters)
}

// ConvertStatusTo populates the provided destination from our ThroughputSettingsUpdateParameters_Status
func (parameters *ThroughputSettingsUpdateParameters_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(parameters)
}

//Storage version of v1alpha1api20210515.ThroughputSettingsResource_Spec
type ThroughputSettingsResource_Spec struct {
	AutoscaleSettings *AutoscaleSettingsResource_Spec `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	Throughput        *int                            `json:"throughput,omitempty"`
}

//Storage version of v1alpha1api20210515.ThroughputSettingsResource_Status
type ThroughputSettingsResource_Status struct {
	AutoscaleSettings   *AutoscaleSettingsResource_Status `json:"autoscaleSettings,omitempty"`
	MinimumThroughput   *string                           `json:"minimumThroughput,omitempty"`
	OfferReplacePending *string                           `json:"offerReplacePending,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Throughput          *int                              `json:"throughput,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettingsResource_Spec
type AutoscaleSettingsResource_Spec struct {
	AutoUpgradePolicy *AutoUpgradePolicyResource_Spec `json:"autoUpgradePolicy,omitempty"`
	MaxThroughput     *int                            `json:"maxThroughput,omitempty"`
	PropertyBag       genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettingsResource_Status
type AutoscaleSettingsResource_Status struct {
	AutoUpgradePolicy   *AutoUpgradePolicyResource_Status `json:"autoUpgradePolicy,omitempty"`
	MaxThroughput       *int                              `json:"maxThroughput,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	TargetMaxThroughput *int                              `json:"targetMaxThroughput,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoUpgradePolicyResource_Spec
type AutoUpgradePolicyResource_Spec struct {
	PropertyBag      genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	ThroughputPolicy *ThroughputPolicyResource_Spec `json:"throughputPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoUpgradePolicyResource_Status
type AutoUpgradePolicyResource_Status struct {
	PropertyBag      genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ThroughputPolicy *ThroughputPolicyResource_Status `json:"throughputPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.ThroughputPolicyResource_Spec
type ThroughputPolicyResource_Spec struct {
	IncrementPercent *int                   `json:"incrementPercent,omitempty"`
	IsEnabled        *bool                  `json:"isEnabled,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ThroughputPolicyResource_Status
type ThroughputPolicyResource_Status struct {
	IncrementPercent *int                   `json:"incrementPercent,omitempty"`
	IsEnabled        *bool                  `json:"isEnabled,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseCollectionThroughputSetting{}, &MongodbDatabaseCollectionThroughputSettingList{})
}
