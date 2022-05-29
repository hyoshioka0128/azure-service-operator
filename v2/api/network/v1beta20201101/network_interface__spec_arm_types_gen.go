// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkInterface_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`

	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of the network interface.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the network interface.
	Properties *NetworkInterfacePropertiesFormatARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkInterface_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (networkInterface NetworkInterface_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (networkInterface *NetworkInterface_SpecARM) GetName() string {
	return networkInterface.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (networkInterface *NetworkInterface_SpecARM) GetType() string {
	return ""
}

type NetworkInterfacePropertiesFormatARM struct {
	// DnsSettings: The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettingsARM `json:"dnsSettings,omitempty"`

	// DscpConfiguration: A reference to the dscp configuration to which the network interface is linked.
	DscpConfiguration *SubResourceARM `json:"dscpConfiguration,omitempty"`

	// EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`

	// EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`

	// HostedWorkloads: A list of references to linked BareMetal resources.
	HostedWorkloads []string `json:"hostedWorkloads,omitempty"`

	// IpConfigurations: A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfiguration_NetworkInterface_SubResourceEmbeddedARM `json:"ipConfigurations,omitempty"`

	// MacAddress: The MAC address of the network interface.
	MacAddress *string `json:"macAddress,omitempty"`

	// MigrationPhase: Migration phase of Network Interface resource.
	MigrationPhase *NetworkInterfacePropertiesFormat_MigrationPhase `json:"migrationPhase,omitempty"`

	// NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupSpec_NetworkInterface_SubResourceEmbeddedARM `json:"networkSecurityGroup,omitempty"`

	// NicType: Type of Network Interface resource.
	NicType *NetworkInterfacePropertiesFormat_NicType `json:"nicType,omitempty"`

	// Primary: Whether this is a primary network interface on a virtual machine.
	Primary *bool `json:"primary,omitempty"`

	// PrivateEndpoint: A reference to the private endpoint to which the network interface is linked.
	PrivateEndpoint *PrivateEndpointSpec_NetworkInterface_SubResourceEmbeddedARM `json:"privateEndpoint,omitempty"`

	// PrivateLinkService: Privatelinkservice of the network interface resource.
	PrivateLinkService *PrivateLinkServiceSpecARM `json:"privateLinkService,omitempty"`

	// ProvisioningState: The provisioning state of the network interface resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// ResourceGuid: The resource GUID property of the network interface resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// TapConfigurations: A list of TapConfigurations of the network interface.
	TapConfigurations []NetworkInterfaceTapConfiguration_NetworkInterface_SubResourceEmbeddedARM `json:"tapConfigurations,omitempty"`

	// VirtualMachine: The reference to a virtual machine.
	VirtualMachine *SubResourceARM `json:"virtualMachine,omitempty"`
}

type NetworkInterfaceDnsSettingsARM struct {
	// AppliedDnsServers: If the VM that uses this NIC is part of an Availability Set, then this list will have the union of
	// all DNS servers from all NICs that are part of the Availability Set. This property is what is configured on each of
	// those VMs.
	AppliedDnsServers []string `json:"appliedDnsServers,omitempty"`

	// DnsServers: List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution.
	// 'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
	DnsServers []string `json:"dnsServers,omitempty"`

	// InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
	// network.
	InternalDnsNameLabel *string `json:"internalDnsNameLabel,omitempty"`

	// InternalDomainNameSuffix: Even if internalDnsNameLabel is not specified, a DNS entry is created for the primary NIC of
	// the VM. This DNS name can be constructed by concatenating the VM name with the value of internalDomainNameSuffix.
	InternalDomainNameSuffix *string `json:"internalDomainNameSuffix,omitempty"`

	// InternalFqdn: Fully qualified DNS name supporting internal communications between VMs in the same virtual network.
	InternalFqdn *string `json:"internalFqdn,omitempty"`
}

type NetworkInterfaceIPConfiguration_NetworkInterface_SubResourceEmbeddedARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Network interface IP configuration properties.
	Properties *NetworkInterfaceIPConfigurationPropertiesFormat_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type NetworkInterfaceTapConfiguration_NetworkInterface_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type NetworkSecurityGroupSpec_NetworkInterface_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type PrivateEndpointSpec_NetworkInterface_SubResourceEmbeddedARM struct {
	// ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type PrivateLinkServiceSpecARM struct {
	// ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type SubResourceARM struct {
	Id *string `json:"id,omitempty"`
}

type NetworkInterfaceIPConfigurationPropertiesFormat_NetworkInterface_SubResourceEmbeddedARM struct {
	// ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool_NetworkInterface_SubResourceEmbeddedARM `json:"applicationGatewayBackendAddressPools,omitempty"`

	// ApplicationSecurityGroups: Application security groups in which the IP configuration is included.
	ApplicationSecurityGroups []ApplicationSecurityGroupSpec_NetworkInterface_SubResourceEmbeddedARM `json:"applicationSecurityGroups,omitempty"`

	// LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.
	LoadBalancerBackendAddressPools []BackendAddressPool_NetworkInterface_SubResourceEmbeddedARM `json:"loadBalancerBackendAddressPools,omitempty"`

	// LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.
	LoadBalancerInboundNatRules []InboundNatRule_NetworkInterface_SubResourceEmbeddedARM `json:"loadBalancerInboundNatRules,omitempty"`

	// Primary: Whether this is a primary customer address on the network interface.
	Primary *bool `json:"primary,omitempty"`

	// PrivateIPAddress: Private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.
	PrivateIPAddressVersion *IPVersion `json:"privateIPAddressVersion,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	// PrivateLinkConnectionProperties: PrivateLinkConnection properties for the network interface.
	PrivateLinkConnectionProperties *NetworkInterfaceIPConfigurationPrivateLinkConnectionPropertiesARM `json:"privateLinkConnectionProperties,omitempty"`

	// ProvisioningState: The provisioning state of the network interface IP configuration.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// PublicIPAddress: Public IP address bound to the IP configuration.
	PublicIPAddress *PublicIPAddressSpec_NetworkInterface_SubResourceEmbeddedARM `json:"publicIPAddress,omitempty"`

	// Subnet: Subnet bound to the IP configuration.
	Subnet *Subnet_NetworkInterface_SubResourceEmbeddedARM `json:"subnet,omitempty"`

	// VirtualNetworkTaps: The reference to Virtual Network Taps.
	VirtualNetworkTaps []VirtualNetworkTapSpec_NetworkInterface_SubResourceEmbeddedARM `json:"virtualNetworkTaps,omitempty"`
}

type ApplicationGatewayBackendAddressPool_NetworkInterface_SubResourceEmbeddedARM struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Name of the backend address pool that is unique within an Application Gateway.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the application gateway backend address pool.
	Properties *ApplicationGatewayBackendAddressPoolPropertiesFormat_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type ApplicationSecurityGroupSpec_NetworkInterface_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type BackendAddressPool_NetworkInterface_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type InboundNatRule_NetworkInterface_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type NetworkInterfaceIPConfigurationPrivateLinkConnectionPropertiesARM struct {
	// Fqdns: List of FQDNs for current private link connection.
	Fqdns []string `json:"fqdns,omitempty"`

	// GroupId: The group ID for current private link connection.
	GroupId *string `json:"groupId,omitempty"`

	// RequiredMemberName: The required member name for current private link connection.
	RequiredMemberName *string `json:"requiredMemberName,omitempty"`
}

type PublicIPAddressSpec_NetworkInterface_SubResourceEmbeddedARM struct {
	// ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Sku: The public IP address SKU.
	Sku *PublicIPAddressSkuARM `json:"sku,omitempty"`

	// Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type Subnet_NetworkInterface_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type VirtualNetworkTapSpec_NetworkInterface_SubResourceEmbeddedARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type ApplicationGatewayBackendAddressPoolPropertiesFormat_NetworkInterface_SubResourceEmbeddedARM struct {
	// BackendAddresses: Backend addresses.
	BackendAddresses []ApplicationGatewayBackendAddressARM `json:"backendAddresses,omitempty"`

	// ProvisioningState: The provisioning state of the backend address pool resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}

type ApplicationGatewayBackendAddressARM struct {
	// Fqdn: Fully qualified domain name (FQDN).
	Fqdn *string `json:"fqdn,omitempty"`

	// IpAddress: IP address.
	IpAddress *string `json:"ipAddress,omitempty"`
}
