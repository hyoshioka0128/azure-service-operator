// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualNetworkGateway_Spec. Use v1beta20201101.VirtualNetworkGateway_Spec instead
type VirtualNetworkGateway_SpecARM struct {
	AzureName        string                                    `json:"azureName,omitempty"`
	Etag             *string                                   `json:"etag,omitempty"`
	ExtendedLocation *ExtendedLocationARM                      `json:"extendedLocation,omitempty"`
	Id               *string                                   `json:"id,omitempty"`
	Location         *string                                   `json:"location,omitempty"`
	Name             string                                    `json:"name,omitempty"`
	Properties       *VirtualNetworkGatewayPropertiesFormatARM `json:"properties,omitempty"`
	Tags             map[string]string                         `json:"tags,omitempty"`
	Type             *string                                   `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworkGateway_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (gateway VirtualNetworkGateway_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (gateway *VirtualNetworkGateway_SpecARM) GetName() string {
	return gateway.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (gateway *VirtualNetworkGateway_SpecARM) GetType() string {
	return ""
}

// Deprecated version of VirtualNetworkGatewayPropertiesFormat. Use v1beta20201101.VirtualNetworkGatewayPropertiesFormat instead
type VirtualNetworkGatewayPropertiesFormatARM struct {
	ActiveActive                   *bool                                                       `json:"activeActive,omitempty"`
	BgpSettings                    *BgpSettingsARM                                             `json:"bgpSettings,omitempty"`
	CustomRoutes                   *AddressSpaceARM                                            `json:"customRoutes,omitempty"`
	EnableBgp                      *bool                                                       `json:"enableBgp,omitempty"`
	EnableDnsForwarding            *bool                                                       `json:"enableDnsForwarding,omitempty"`
	EnablePrivateIpAddress         *bool                                                       `json:"enablePrivateIpAddress,omitempty"`
	GatewayDefaultSite             *SubResourceARM                                             `json:"gatewayDefaultSite,omitempty"`
	GatewayType                    *VirtualNetworkGatewayPropertiesFormat_GatewayType          `json:"gatewayType,omitempty"`
	InboundDnsForwardingEndpoint   *string                                                     `json:"inboundDnsForwardingEndpoint,omitempty"`
	IpConfigurations               []VirtualNetworkGatewayIPConfigurationARM                   `json:"ipConfigurations,omitempty"`
	ProvisioningState              *ProvisioningState                                          `json:"provisioningState,omitempty"`
	ResourceGuid                   *string                                                     `json:"resourceGuid,omitempty"`
	Sku                            *VirtualNetworkGatewaySkuARM                                `json:"sku,omitempty"`
	VNetExtendedLocationResourceId *string                                                     `json:"vNetExtendedLocationResourceId,omitempty"`
	VpnClientConfiguration         *VpnClientConfigurationARM                                  `json:"vpnClientConfiguration,omitempty"`
	VpnGatewayGeneration           *VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration `json:"vpnGatewayGeneration,omitempty"`
	VpnType                        *VirtualNetworkGatewayPropertiesFormat_VpnType              `json:"vpnType,omitempty"`
}

// Deprecated version of AddressSpace. Use v1beta20201101.AddressSpace instead
type AddressSpaceARM struct {
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

// Deprecated version of BgpSettings. Use v1beta20201101.BgpSettings instead
type BgpSettingsARM struct {
	Asn                 *uint32                               `json:"asn,omitempty"`
	BgpPeeringAddress   *string                               `json:"bgpPeeringAddress,omitempty"`
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddressARM `json:"bgpPeeringAddresses,omitempty"`
	PeerWeight          *int                                  `json:"peerWeight,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayIPConfiguration. Use v1beta20201101.VirtualNetworkGatewayIPConfiguration instead
type VirtualNetworkGatewayIPConfigurationARM struct {
	Etag       *string                                                  `json:"etag,omitempty"`
	Id         *string                                                  `json:"id,omitempty"`
	Name       *string                                                  `json:"name,omitempty"`
	Properties *VirtualNetworkGatewayIPConfigurationPropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualNetworkGatewaySku. Use v1beta20201101.VirtualNetworkGatewaySku instead
type VirtualNetworkGatewaySkuARM struct {
	Capacity *int                           `json:"capacity,omitempty"`
	Name     *VirtualNetworkGatewaySku_Name `json:"name,omitempty"`
	Tier     *VirtualNetworkGatewaySku_Tier `json:"tier,omitempty"`
}

// Deprecated version of VpnClientConfiguration. Use v1beta20201101.VpnClientConfiguration instead
type VpnClientConfigurationARM struct {
	AadAudience                  *string                                         `json:"aadAudience,omitempty"`
	AadIssuer                    *string                                         `json:"aadIssuer,omitempty"`
	AadTenant                    *string                                         `json:"aadTenant,omitempty"`
	RadiusServerAddress          *string                                         `json:"radiusServerAddress,omitempty"`
	RadiusServerSecret           *string                                         `json:"radiusServerSecret,omitempty"`
	RadiusServers                []RadiusServerARM                               `json:"radiusServers,omitempty"`
	VpnAuthenticationTypes       []VpnClientConfiguration_VpnAuthenticationTypes `json:"vpnAuthenticationTypes,omitempty"`
	VpnClientAddressPool         *AddressSpaceARM                                `json:"vpnClientAddressPool,omitempty"`
	VpnClientIpsecPolicies       []IpsecPolicyARM                                `json:"vpnClientIpsecPolicies,omitempty"`
	VpnClientProtocols           []VpnClientConfiguration_VpnClientProtocols     `json:"vpnClientProtocols,omitempty"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificateARM                `json:"vpnClientRevokedCertificates,omitempty"`
	VpnClientRootCertificates    []VpnClientRootCertificateARM                   `json:"vpnClientRootCertificates,omitempty"`
}

// Deprecated version of IPConfigurationBgpPeeringAddress. Use v1beta20201101.IPConfigurationBgpPeeringAddress instead
type IPConfigurationBgpPeeringAddressARM struct {
	CustomBgpIpAddresses  []string `json:"customBgpIpAddresses,omitempty"`
	DefaultBgpIpAddresses []string `json:"defaultBgpIpAddresses,omitempty"`
	IpconfigurationId     *string  `json:"ipconfigurationId,omitempty"`
	TunnelIpAddresses     []string `json:"tunnelIpAddresses,omitempty"`
}

// Deprecated version of IpsecPolicy. Use v1beta20201101.IpsecPolicy instead
type IpsecPolicyARM struct {
	DhGroup             *DhGroup         `json:"dhGroup,omitempty"`
	IkeEncryption       *IkeEncryption   `json:"ikeEncryption,omitempty"`
	IkeIntegrity        *IkeIntegrity    `json:"ikeIntegrity,omitempty"`
	IpsecEncryption     *IpsecEncryption `json:"ipsecEncryption,omitempty"`
	IpsecIntegrity      *IpsecIntegrity  `json:"ipsecIntegrity,omitempty"`
	PfsGroup            *PfsGroup        `json:"pfsGroup,omitempty"`
	SaDataSizeKilobytes *int             `json:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   *int             `json:"saLifeTimeSeconds,omitempty"`
}

// Deprecated version of RadiusServer. Use v1beta20201101.RadiusServer instead
type RadiusServerARM struct {
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`
	RadiusServerScore   *int    `json:"radiusServerScore,omitempty"`
	RadiusServerSecret  *string `json:"radiusServerSecret,omitempty"`
}

// Deprecated version of VirtualNetworkGatewayIPConfigurationPropertiesFormat. Use v1beta20201101.VirtualNetworkGatewayIPConfigurationPropertiesFormat instead
type VirtualNetworkGatewayIPConfigurationPropertiesFormatARM struct {
	PrivateIPAddress          *string             `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *ProvisioningState  `json:"provisioningState,omitempty"`
	PublicIPAddress           *SubResourceARM     `json:"publicIPAddress,omitempty"`
	Subnet                    *SubResourceARM     `json:"subnet,omitempty"`
}

// Deprecated version of VpnClientRevokedCertificate. Use v1beta20201101.VpnClientRevokedCertificate instead
type VpnClientRevokedCertificateARM struct {
	Etag       *string                                         `json:"etag,omitempty"`
	Id         *string                                         `json:"id,omitempty"`
	Name       *string                                         `json:"name,omitempty"`
	Properties *VpnClientRevokedCertificatePropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of VpnClientRootCertificate. Use v1beta20201101.VpnClientRootCertificate instead
type VpnClientRootCertificateARM struct {
	Etag       *string                                      `json:"etag,omitempty"`
	Id         *string                                      `json:"id,omitempty"`
	Name       *string                                      `json:"name,omitempty"`
	Properties *VpnClientRootCertificatePropertiesFormatARM `json:"properties,omitempty"`
}

// Deprecated version of VpnClientRevokedCertificatePropertiesFormat. Use v1beta20201101.VpnClientRevokedCertificatePropertiesFormat instead
type VpnClientRevokedCertificatePropertiesFormatARM struct {
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Thumbprint        *string            `json:"thumbprint,omitempty"`
}

// Deprecated version of VpnClientRootCertificatePropertiesFormat. Use v1beta20201101.VpnClientRootCertificatePropertiesFormat instead
type VpnClientRootCertificatePropertiesFormatARM struct {
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	PublicCertData    *string            `json:"publicCertData,omitempty"`
}