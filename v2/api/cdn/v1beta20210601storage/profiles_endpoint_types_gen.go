// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cdn.azure.com,resources=profilesendpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cdn.azure.com,resources={profilesendpoints/status,profilesendpoints/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210601.ProfilesEndpoint
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/cdn.json
type ProfilesEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProfilesEndpoint_Spec   `json:"spec,omitempty"`
	Status            ProfilesEndpoint_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ProfilesEndpoint{}

// GetConditions returns the conditions of the resource
func (endpoint *ProfilesEndpoint) GetConditions() conditions.Conditions {
	return endpoint.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (endpoint *ProfilesEndpoint) SetConditions(conditions conditions.Conditions) {
	endpoint.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ProfilesEndpoint{}

// AzureName returns the Azure name of the resource
func (endpoint *ProfilesEndpoint) AzureName() string {
	return endpoint.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (endpoint ProfilesEndpoint) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (endpoint *ProfilesEndpoint) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (endpoint *ProfilesEndpoint) GetSpec() genruntime.ConvertibleSpec {
	return &endpoint.Spec
}

// GetStatus returns the status of this resource
func (endpoint *ProfilesEndpoint) GetStatus() genruntime.ConvertibleStatus {
	return &endpoint.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (endpoint *ProfilesEndpoint) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (endpoint *ProfilesEndpoint) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ProfilesEndpoint_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (endpoint *ProfilesEndpoint) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(endpoint.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  endpoint.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (endpoint *ProfilesEndpoint) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ProfilesEndpoint_STATUS); ok {
		endpoint.Status = *st
		return nil
	}

	// Convert status to required version
	var st ProfilesEndpoint_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	endpoint.Status = st
	return nil
}

// Hub marks that this ProfilesEndpoint is the hub type for conversion
func (endpoint *ProfilesEndpoint) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (endpoint *ProfilesEndpoint) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: endpoint.Spec.OriginalVersion,
		Kind:    "ProfilesEndpoint",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210601.ProfilesEndpoint
// Generator information:
// - Generated from: /cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/cdn.json
type ProfilesEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProfilesEndpoint `json:"items"`
}

// Storage version of v1beta20210601.ProfilesEndpoint_STATUS
type ProfilesEndpoint_STATUS struct {
	Conditions                       []conditions.Condition                                      `json:"conditions,omitempty"`
	ContentTypesToCompress           []string                                                    `json:"contentTypesToCompress,omitempty"`
	CustomDomains                    []CustomDomain_STATUS                                       `json:"customDomains,omitempty"`
	DefaultOriginGroup               *ResourceReference_STATUS                                   `json:"defaultOriginGroup,omitempty"`
	DeliveryPolicy                   *EndpointProperties_DeliveryPolicy_STATUS                   `json:"deliveryPolicy,omitempty"`
	GeoFilters                       []GeoFilter_STATUS                                          `json:"geoFilters,omitempty"`
	HostName                         *string                                                     `json:"hostName,omitempty"`
	Id                               *string                                                     `json:"id,omitempty"`
	IsCompressionEnabled             *bool                                                       `json:"isCompressionEnabled,omitempty"`
	IsHttpAllowed                    *bool                                                       `json:"isHttpAllowed,omitempty"`
	IsHttpsAllowed                   *bool                                                       `json:"isHttpsAllowed,omitempty"`
	Location                         *string                                                     `json:"location,omitempty"`
	Name                             *string                                                     `json:"name,omitempty"`
	OptimizationType                 *string                                                     `json:"optimizationType,omitempty"`
	OriginGroups                     []DeepCreatedOriginGroup_STATUS                             `json:"originGroups,omitempty"`
	OriginHostHeader                 *string                                                     `json:"originHostHeader,omitempty"`
	OriginPath                       *string                                                     `json:"originPath,omitempty"`
	Origins                          []DeepCreatedOrigin_STATUS                                  `json:"origins,omitempty"`
	ProbePath                        *string                                                     `json:"probePath,omitempty"`
	PropertyBag                      genruntime.PropertyBag                                      `json:"$propertyBag,omitempty"`
	ProvisioningState                *string                                                     `json:"provisioningState,omitempty"`
	QueryStringCachingBehavior       *string                                                     `json:"queryStringCachingBehavior,omitempty"`
	ResourceState                    *string                                                     `json:"resourceState,omitempty"`
	SystemData                       *SystemData_STATUS                                          `json:"systemData,omitempty"`
	Tags                             map[string]string                                           `json:"tags,omitempty"`
	Type                             *string                                                     `json:"type,omitempty"`
	UrlSigningKeys                   []UrlSigningKey_STATUS                                      `json:"urlSigningKeys,omitempty"`
	WebApplicationFirewallPolicyLink *EndpointProperties_WebApplicationFirewallPolicyLink_STATUS `json:"webApplicationFirewallPolicyLink,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ProfilesEndpoint_STATUS{}

// ConvertStatusFrom populates our ProfilesEndpoint_STATUS from the provided source
func (endpoint *ProfilesEndpoint_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(endpoint)
}

// ConvertStatusTo populates the provided destination from our ProfilesEndpoint_STATUS
func (endpoint *ProfilesEndpoint_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(endpoint)
}

// Storage version of v1beta20210601.ProfilesEndpoint_Spec
type ProfilesEndpoint_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName              string                             `json:"azureName,omitempty"`
	ContentTypesToCompress []string                           `json:"contentTypesToCompress,omitempty"`
	CustomDomains          []CustomDomain                     `json:"customDomains,omitempty"`
	DefaultOriginGroup     *ResourceReference                 `json:"defaultOriginGroup,omitempty"`
	DeliveryPolicy         *EndpointProperties_DeliveryPolicy `json:"deliveryPolicy,omitempty"`
	GeoFilters             []GeoFilter                        `json:"geoFilters,omitempty"`
	HostName               *string                            `json:"hostName,omitempty"`
	Id                     *string                            `json:"id,omitempty"`
	IsCompressionEnabled   *bool                              `json:"isCompressionEnabled,omitempty"`
	IsHttpAllowed          *bool                              `json:"isHttpAllowed,omitempty"`
	IsHttpsAllowed         *bool                              `json:"isHttpsAllowed,omitempty"`
	Location               *string                            `json:"location,omitempty"`
	OptimizationType       *string                            `json:"optimizationType,omitempty"`
	OriginGroups           []DeepCreatedOriginGroup           `json:"originGroups,omitempty"`
	OriginHostHeader       *string                            `json:"originHostHeader,omitempty"`
	OriginPath             *string                            `json:"originPath,omitempty"`
	OriginalVersion        string                             `json:"originalVersion,omitempty"`
	Origins                []DeepCreatedOrigin                `json:"origins,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                            *genruntime.KnownResourceReference                   `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	ProbePath                        *string                                              `json:"probePath,omitempty"`
	PropertyBag                      genruntime.PropertyBag                               `json:"$propertyBag,omitempty"`
	ProvisioningState                *string                                              `json:"provisioningState,omitempty"`
	QueryStringCachingBehavior       *string                                              `json:"queryStringCachingBehavior,omitempty"`
	ResourceState                    *string                                              `json:"resourceState,omitempty"`
	SystemData                       *SystemData                                          `json:"systemData,omitempty"`
	Tags                             map[string]string                                    `json:"tags,omitempty"`
	Type                             *string                                              `json:"type,omitempty"`
	UrlSigningKeys                   []UrlSigningKey                                      `json:"urlSigningKeys,omitempty"`
	WebApplicationFirewallPolicyLink *EndpointProperties_WebApplicationFirewallPolicyLink `json:"webApplicationFirewallPolicyLink,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ProfilesEndpoint_Spec{}

// ConvertSpecFrom populates our ProfilesEndpoint_Spec from the provided source
func (endpoint *ProfilesEndpoint_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(endpoint)
}

// ConvertSpecTo populates the provided destination from our ProfilesEndpoint_Spec
func (endpoint *ProfilesEndpoint_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(endpoint)
}

// Storage version of v1beta20210601.CustomDomain
type CustomDomain struct {
	CustomHttpsParameters           *CustomDomainHttpsParameters `json:"customHttpsParameters,omitempty"`
	CustomHttpsProvisioningState    *string                      `json:"customHttpsProvisioningState,omitempty"`
	CustomHttpsProvisioningSubstate *string                      `json:"customHttpsProvisioningSubstate,omitempty"`
	HostName                        *string                      `json:"hostName,omitempty"`
	Id                              *string                      `json:"id,omitempty"`
	Name                            *string                      `json:"name,omitempty"`
	PropertyBag                     genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                      `json:"provisioningState,omitempty"`
	ResourceState                   *string                      `json:"resourceState,omitempty"`
	SystemData                      *SystemData                  `json:"systemData,omitempty"`
	Type                            *string                      `json:"type,omitempty"`
	ValidationData                  *string                      `json:"validationData,omitempty"`
}

// Storage version of v1beta20210601.CustomDomain_STATUS
type CustomDomain_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
}

// Storage version of v1beta20210601.DeepCreatedOrigin
type DeepCreatedOrigin struct {
	Enabled                    *bool   `json:"enabled,omitempty"`
	HostName                   *string `json:"hostName,omitempty"`
	HttpPort                   *int    `json:"httpPort,omitempty"`
	HttpsPort                  *int    `json:"httpsPort,omitempty"`
	Name                       *string `json:"name,omitempty"`
	OriginHostHeader           *string `json:"originHostHeader,omitempty"`
	Priority                   *int    `json:"priority,omitempty"`
	PrivateEndpointStatus      *string `json:"privateEndpointStatus,omitempty"`
	PrivateLinkAlias           *string `json:"privateLinkAlias,omitempty"`
	PrivateLinkApprovalMessage *string `json:"privateLinkApprovalMessage,omitempty"`

	// PrivateLinkLocationReference: The location of the Private Link resource. Required only if 'privateLinkResourceId' is
	// populated
	PrivateLinkLocationReference *genruntime.ResourceReference `armReference:"PrivateLinkLocation" json:"privateLinkLocationReference,omitempty"`

	// PrivateLinkResourceReference: The Resource Id of the Private Link resource. Populating this optional field indicates
	// that this backend is 'Private'
	PrivateLinkResourceReference *genruntime.ResourceReference `armReference:"PrivateLinkResourceId" json:"privateLinkResourceReference,omitempty"`
	PropertyBag                  genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Weight                       *int                          `json:"weight,omitempty"`
}

// Storage version of v1beta20210601.DeepCreatedOriginGroup
type DeepCreatedOriginGroup struct {
	HealthProbeSettings                                   *HealthProbeParameters                       `json:"healthProbeSettings,omitempty"`
	Name                                                  *string                                      `json:"name,omitempty"`
	Origins                                               []ResourceReference                          `json:"origins,omitempty"`
	PropertyBag                                           genruntime.PropertyBag                       `json:"$propertyBag,omitempty"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParameters `json:"responseBasedOriginErrorDetectionSettings,omitempty"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                         `json:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes,omitempty"`
}

// Storage version of v1beta20210601.DeepCreatedOriginGroup_STATUS
type DeepCreatedOriginGroup_STATUS struct {
	HealthProbeSettings                                   *HealthProbeParameters_STATUS                       `json:"healthProbeSettings,omitempty"`
	Name                                                  *string                                             `json:"name,omitempty"`
	Origins                                               []ResourceReference_STATUS                          `json:"origins,omitempty"`
	PropertyBag                                           genruntime.PropertyBag                              `json:"$propertyBag,omitempty"`
	ResponseBasedOriginErrorDetectionSettings             *ResponseBasedOriginErrorDetectionParameters_STATUS `json:"responseBasedOriginErrorDetectionSettings,omitempty"`
	TrafficRestorationTimeToHealedOrNewEndpointsInMinutes *int                                                `json:"trafficRestorationTimeToHealedOrNewEndpointsInMinutes,omitempty"`
}

// Storage version of v1beta20210601.DeepCreatedOrigin_STATUS
type DeepCreatedOrigin_STATUS struct {
	Enabled                    *bool                  `json:"enabled,omitempty"`
	HostName                   *string                `json:"hostName,omitempty"`
	HttpPort                   *int                   `json:"httpPort,omitempty"`
	HttpsPort                  *int                   `json:"httpsPort,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	OriginHostHeader           *string                `json:"originHostHeader,omitempty"`
	Priority                   *int                   `json:"priority,omitempty"`
	PrivateEndpointStatus      *string                `json:"privateEndpointStatus,omitempty"`
	PrivateLinkAlias           *string                `json:"privateLinkAlias,omitempty"`
	PrivateLinkApprovalMessage *string                `json:"privateLinkApprovalMessage,omitempty"`
	PrivateLinkLocation        *string                `json:"privateLinkLocation,omitempty"`
	PrivateLinkResourceId      *string                `json:"privateLinkResourceId,omitempty"`
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Weight                     *int                   `json:"weight,omitempty"`
}

// Storage version of v1beta20210601.EndpointProperties_DeliveryPolicy
type EndpointProperties_DeliveryPolicy struct {
	Description *string                `json:"description,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rules       []DeliveryRule         `json:"rules,omitempty"`
}

// Storage version of v1beta20210601.EndpointProperties_DeliveryPolicy_STATUS
type EndpointProperties_DeliveryPolicy_STATUS struct {
	Description *string                `json:"description,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rules       []DeliveryRule_STATUS  `json:"rules,omitempty"`
}

// Storage version of v1beta20210601.EndpointProperties_WebApplicationFirewallPolicyLink
type EndpointProperties_WebApplicationFirewallPolicyLink struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.EndpointProperties_WebApplicationFirewallPolicyLink_STATUS
type EndpointProperties_WebApplicationFirewallPolicyLink_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.GeoFilter
type GeoFilter struct {
	Action       *string                `json:"action,omitempty"`
	CountryCodes []string               `json:"countryCodes,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RelativePath *string                `json:"relativePath,omitempty"`
}

// Storage version of v1beta20210601.GeoFilter_STATUS
type GeoFilter_STATUS struct {
	Action       *string                `json:"action,omitempty"`
	CountryCodes []string               `json:"countryCodes,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RelativePath *string                `json:"relativePath,omitempty"`
}

// Storage version of v1beta20210601.ResourceReference
type ResourceReference struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1beta20210601.ResourceReference_STATUS
type ResourceReference_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.UrlSigningKey
type UrlSigningKey struct {
	KeyId               *string                       `json:"keyId,omitempty"`
	KeySourceParameters *KeyVaultSigningKeyParameters `json:"keySourceParameters,omitempty"`
	PropertyBag         genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.UrlSigningKey_STATUS
type UrlSigningKey_STATUS struct {
	KeyId               *string                              `json:"keyId,omitempty"`
	KeySourceParameters *KeyVaultSigningKeyParameters_STATUS `json:"keySourceParameters,omitempty"`
	PropertyBag         genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.CustomDomainHttpsParameters
type CustomDomainHttpsParameters struct {
	CertificateSource *string                `json:"certificateSource,omitempty"`
	MinimumTlsVersion *string                `json:"minimumTlsVersion,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProtocolType      *string                `json:"protocolType,omitempty"`
}

// Storage version of v1beta20210601.DeliveryRule
type DeliveryRule struct {
	Actions     []DeliveryRuleAction    `json:"actions,omitempty"`
	Conditions  []DeliveryRuleCondition `json:"conditions,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Order       *int                    `json:"order,omitempty"`
	PropertyBag genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.DeliveryRule_STATUS
type DeliveryRule_STATUS struct {
	Actions     []DeliveryRuleAction_STATUS    `json:"actions,omitempty"`
	Conditions  []DeliveryRuleCondition_STATUS `json:"conditions,omitempty"`
	Name        *string                        `json:"name,omitempty"`
	Order       *int                           `json:"order,omitempty"`
	PropertyBag genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.HealthProbeParameters
type HealthProbeParameters struct {
	ProbeIntervalInSeconds *int                   `json:"probeIntervalInSeconds,omitempty"`
	ProbePath              *string                `json:"probePath,omitempty"`
	ProbeProtocol          *string                `json:"probeProtocol,omitempty"`
	ProbeRequestType       *string                `json:"probeRequestType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.HealthProbeParameters_STATUS
type HealthProbeParameters_STATUS struct {
	ProbeIntervalInSeconds *int                   `json:"probeIntervalInSeconds,omitempty"`
	ProbePath              *string                `json:"probePath,omitempty"`
	ProbeProtocol          *string                `json:"probeProtocol,omitempty"`
	ProbeRequestType       *string                `json:"probeRequestType,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.KeyVaultSigningKeyParameters
type KeyVaultSigningKeyParameters struct {
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceGroupName *string                `json:"resourceGroupName,omitempty"`
	SecretName        *string                `json:"secretName,omitempty"`
	SecretVersion     *string                `json:"secretVersion,omitempty"`
	SubscriptionId    *string                `json:"subscriptionId,omitempty"`
	TypeName          *string                `json:"typeName,omitempty"`
	VaultName         *string                `json:"vaultName,omitempty"`
}

// Storage version of v1beta20210601.KeyVaultSigningKeyParameters_STATUS
type KeyVaultSigningKeyParameters_STATUS struct {
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceGroupName *string                `json:"resourceGroupName,omitempty"`
	SecretName        *string                `json:"secretName,omitempty"`
	SecretVersion     *string                `json:"secretVersion,omitempty"`
	SubscriptionId    *string                `json:"subscriptionId,omitempty"`
	TypeName          *string                `json:"typeName,omitempty"`
	VaultName         *string                `json:"vaultName,omitempty"`
}

// Storage version of v1beta20210601.ResponseBasedOriginErrorDetectionParameters
type ResponseBasedOriginErrorDetectionParameters struct {
	HttpErrorRanges                          []HttpErrorRangeParameters `json:"httpErrorRanges,omitempty"`
	PropertyBag                              genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	ResponseBasedDetectedErrorTypes          *string                    `json:"responseBasedDetectedErrorTypes,omitempty"`
	ResponseBasedFailoverThresholdPercentage *int                       `json:"responseBasedFailoverThresholdPercentage,omitempty"`
}

// Storage version of v1beta20210601.ResponseBasedOriginErrorDetectionParameters_STATUS
type ResponseBasedOriginErrorDetectionParameters_STATUS struct {
	HttpErrorRanges                          []HttpErrorRangeParameters_STATUS `json:"httpErrorRanges,omitempty"`
	PropertyBag                              genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	ResponseBasedDetectedErrorTypes          *string                           `json:"responseBasedDetectedErrorTypes,omitempty"`
	ResponseBasedFailoverThresholdPercentage *int                              `json:"responseBasedFailoverThresholdPercentage,omitempty"`
}

// Storage version of v1beta20210601.DeliveryRuleAction
type DeliveryRuleAction struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.DeliveryRuleAction_STATUS
type DeliveryRuleAction_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.DeliveryRuleCondition
type DeliveryRuleCondition struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.DeliveryRuleCondition_STATUS
type DeliveryRuleCondition_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.HttpErrorRangeParameters
type HttpErrorRangeParameters struct {
	Begin       *int                   `json:"begin,omitempty"`
	End         *int                   `json:"end,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20210601.HttpErrorRangeParameters_STATUS
type HttpErrorRangeParameters_STATUS struct {
	Begin       *int                   `json:"begin,omitempty"`
	End         *int                   `json:"end,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ProfilesEndpoint{}, &ProfilesEndpointList{})
}
