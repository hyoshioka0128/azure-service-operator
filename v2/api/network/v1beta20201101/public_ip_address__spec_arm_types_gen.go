// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PublicIPAddress_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name string `json:"name,omitempty"`

	// Properties: Public IP address properties.
	Properties *PublicIPAddressPropertiesFormatARM `json:"properties,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSkuARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PublicIPAddress_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (address PublicIPAddress_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (address *PublicIPAddress_SpecARM) GetName() string {
	return address.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (address *PublicIPAddress_SpecARM) GetType() string {
	return ""
}

type PublicIPAddressPropertiesFormatARM struct {
	// DdosSettings: The DDoS protection custom policy associated with the public IP address.
	DdosSettings *DdosSettingsARM `json:"ddosSettings,omitempty"`

	// DnsSettings: The FQDN of the DNS record associated with the public IP address.
	DnsSettings *PublicIPAddressDnsSettingsARM `json:"dnsSettings,omitempty"`

	// IdleTimeoutInMinutes: The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// IpAddress: The IP address associated with the public IP address resource.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpConfiguration: The IP configuration associated with the public IP address.
	IpConfiguration *IPConfiguration_PublicIPAddress_SubResourceEmbeddedARM `json:"ipConfiguration,omitempty"`

	// IpTags: The list of tags associated with the public IP address.
	IpTags []IpTagARM `json:"ipTags,omitempty"`

	// LinkedPublicIPAddress: The linked public IP address of the public IP address resource.
	LinkedPublicIPAddress *PublicIPAddressSpec_PublicIPAddress_SubResourceEmbeddedARM `json:"linkedPublicIPAddress,omitempty"`

	// MigrationPhase: Migration phase of Public IP Address.
	MigrationPhase *PublicIPAddressPropertiesFormat_MigrationPhase `json:"migrationPhase,omitempty"`

	// NatGateway: The NatGateway for the Public IP address.
	NatGateway *NatGatewaySpec_PublicIPAddress_SubResourceEmbeddedARM `json:"natGateway,omitempty"`

	// ProvisioningState: The provisioning state of the public IP address resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// PublicIPAddressVersion: The public IP address version.
	PublicIPAddressVersion *IPVersion `json:"publicIPAddressVersion,omitempty"`

	// PublicIPAllocationMethod: The public IP address allocation method.
	PublicIPAllocationMethod *IPAllocationMethod `json:"publicIPAllocationMethod,omitempty"`

	// PublicIPPrefix: The Public IP Prefix this Public IP Address should be allocated from.
	PublicIPPrefix *SubResourceARM `json:"publicIPPrefix,omitempty"`

	// ResourceGuid: The resource GUID property of the public IP address resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// ServicePublicIPAddress: The service public IP address of the public IP address resource.
	ServicePublicIPAddress *PublicIPAddressSpec_PublicIPAddress_SubResourceEmbeddedARM `json:"servicePublicIPAddress,omitempty"`
}

type PublicIPAddressSkuARM struct {
	// Name: Name of a public IP address SKU.
	Name *PublicIPAddressSku_Name `json:"name,omitempty"`

	// Tier: Tier of a public IP address SKU.
	Tier *PublicIPAddressSku_Tier `json:"tier,omitempty"`
}

type DdosSettingsARM struct {
	// DdosCustomPolicy: The DDoS custom policy associated with the public IP.
	DdosCustomPolicy *SubResourceARM `json:"ddosCustomPolicy,omitempty"`

	// ProtectedIP: Enables DDoS protection on the public IP.
	ProtectedIP *bool `json:"protectedIP,omitempty"`

	// ProtectionCoverage: The DDoS protection policy customizability of the public IP. Only standard coverage will have the
	// ability to be customized.
	ProtectionCoverage *DdosSettings_ProtectionCoverage `json:"protectionCoverage,omitempty"`
}

type IPConfiguration_PublicIPAddress_SubResourceEmbeddedARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the IP configuration.
	Properties *IPConfigurationPropertiesFormat_PublicIPAddress_SubResourceEmbeddedARM `json:"properties,omitempty"`
}

type IpTagARM struct {
	// IpTagType: The IP tag type. Example: FirstPartyUsage.
	IpTagType *string `json:"ipTagType,omitempty"`

	// Tag: The value of the IP tag associated with the public IP. Example: SQL.
	Tag *string `json:"tag,omitempty"`
}

type NatGatewaySpec_PublicIPAddress_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Sku: The nat gateway SKU.
	Sku *NatGatewaySkuARM `json:"sku,omitempty"`

	// Zones: A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `json:"zones,omitempty"`
}

type PublicIPAddressDnsSettingsARM struct {
	// DomainNameLabel: The domain name label. The concatenation of the domain name label and the regionalized DNS zone make up
	// the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS
	// record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`

	// Fqdn: The Fully Qualified Domain Name of the A DNS record associated with the public IP. This is the concatenation of
	// the domainNameLabel and the regionalized DNS zone.
	Fqdn *string `json:"fqdn,omitempty"`

	// ReverseFqdn: The reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If
	// the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain
	// to the reverse FQDN.
	ReverseFqdn *string `json:"reverseFqdn,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","Standard"}
type PublicIPAddressSku_Name string

const (
	PublicIPAddressSku_Name_Basic    = PublicIPAddressSku_Name("Basic")
	PublicIPAddressSku_Name_Standard = PublicIPAddressSku_Name("Standard")
)

// +kubebuilder:validation:Enum={"Global","Regional"}
type PublicIPAddressSku_Tier string

const (
	PublicIPAddressSku_Tier_Global   = PublicIPAddressSku_Tier("Global")
	PublicIPAddressSku_Tier_Regional = PublicIPAddressSku_Tier("Regional")
)

type PublicIPAddressSpec_PublicIPAddress_SubResourceEmbeddedARM struct {
	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSkuARM `json:"sku,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type IPConfigurationPropertiesFormat_PublicIPAddress_SubResourceEmbeddedARM struct {
	// PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	// ProvisioningState: The provisioning state of the IP configuration resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// PublicIPAddress: The reference to the public IP resource.
	PublicIPAddress *PublicIPAddressSpec_PublicIPAddress_SubResourceEmbeddedARM `json:"publicIPAddress,omitempty"`

	// Subnet: The reference to the subnet resource.
	Subnet *Subnet_PublicIPAddress_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

type NatGatewaySkuARM struct {
	// Name: Name of Nat Gateway SKU.
	Name *NatGatewaySku_Name `json:"name,omitempty"`
}

type Subnet_PublicIPAddress_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}
