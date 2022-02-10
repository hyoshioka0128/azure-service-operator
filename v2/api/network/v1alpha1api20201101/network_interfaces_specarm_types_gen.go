// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkInterfaces_SPECARM struct {
	AzureName string `json:"azureName"`

	//ExtendedLocation: The extended location of the network interface.
	ExtendedLocation *ExtendedLocation_SpecARM `json:"extendedLocation,omitempty"`
	Id               *string                   `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties of the network interface.
	Properties *NetworkInterfacePropertiesFormat_SpecARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkInterfaces_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (specarm NetworkInterfaces_SPECARM) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (specarm NetworkInterfaces_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm NetworkInterfaces_SPECARM) GetType() string {
	return ""
}

type NetworkInterfacePropertiesFormat_SpecARM struct {
	//DnsSettings: The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettings_SpecARM `json:"dnsSettings,omitempty"`

	//EnableAcceleratedNetworking: If the network interface is accelerated networking
	//enabled.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`

	//EnableIPForwarding: Indicates whether IP forwarding is enabled on this network
	//interface.
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`

	//IpConfigurations: A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfiguration_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"ipConfigurations,omitempty"`

	//MigrationPhase: Migration phase of Network Interface resource.
	MigrationPhase *NetworkInterfacePropertiesFormatSpecMigrationPhase `json:"migrationPhase,omitempty"`

	//NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroup_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"networkSecurityGroup,omitempty"`

	//NicType: Type of Network Interface resource.
	NicType *NetworkInterfacePropertiesFormatSpecNicType `json:"nicType,omitempty"`

	//PrivateLinkService: Privatelinkservice of the network interface resource.
	PrivateLinkService *PrivateLinkService_SpecARM `json:"privateLinkService,omitempty"`
}

type NetworkInterfaceDnsSettings_SpecARM struct {
	//DnsServers: List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch
	//to azure provided DNS resolution. 'AzureProvidedDNS' value cannot be combined
	//with other IPs, it must be the only value in dnsServers collection.
	DnsServers []string `json:"dnsServers,omitempty"`

	//InternalDnsNameLabel: Relative DNS name for this NIC used for internal
	//communications between VMs in the same virtual network.
	InternalDnsNameLabel *string `json:"internalDnsNameLabel,omitempty"`
}

type NetworkInterfaceIPConfiguration_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Network interface IP configuration properties.
	Properties *NetworkInterfaceIPConfigurationPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type NetworkSecurityGroup_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Properties: Properties of the network security group.
	Properties *NetworkSecurityGroupPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type PrivateLinkService_SpecARM struct {
	//ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation_SpecARM `json:"extendedLocation,omitempty"`
	Id               *string                   `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Properties: Properties of the private link service.
	Properties *PrivateLinkServiceProperties_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type NetworkInterfaceIPConfigurationPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	//ApplicationGatewayBackendAddressPools: The reference to
	//ApplicationGatewayBackendAddressPool resource.
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool_SpecARM `json:"applicationGatewayBackendAddressPools,omitempty"`

	//ApplicationSecurityGroups: Application security groups in which the IP
	//configuration is included.
	ApplicationSecurityGroups []ApplicationSecurityGroup_SpecARM `json:"applicationSecurityGroups,omitempty"`

	//LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool
	//resource.
	LoadBalancerBackendAddressPools []BackendAddressPool_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"loadBalancerBackendAddressPools,omitempty"`

	//LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.
	LoadBalancerInboundNatRules []InboundNatRule_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"loadBalancerInboundNatRules,omitempty"`

	//Primary: Whether this is a primary customer address on the network interface.
	Primary *bool `json:"primary,omitempty"`

	//PrivateIPAddress: Private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	//PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6.
	//Default is IPv4.
	PrivateIPAddressVersion *IPVersion_Spec `json:"privateIPAddressVersion,omitempty"`

	//PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_Spec `json:"privateIPAllocationMethod,omitempty"`

	//PublicIPAddress: Public IP address bound to the IP configuration.
	PublicIPAddress *PublicIPAddress_SpecARM `json:"publicIPAddress,omitempty"`

	//Subnet: Subnet bound to the IP configuration.
	Subnet *Subnet_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"subnet,omitempty"`

	//VirtualNetworkTaps: The reference to Virtual Network Taps.
	VirtualNetworkTaps []VirtualNetworkTap_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"virtualNetworkTaps,omitempty"`
}

type NetworkSecurityGroupPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	//SecurityRules: A collection of security rules of the network security group.
	SecurityRules []SecurityRule_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"securityRules,omitempty"`
}

type PrivateLinkServiceProperties_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	//AutoApproval: The auto-approval list of the private link service.
	AutoApproval *ResourceSet_SpecARM `json:"autoApproval,omitempty"`

	//EnableProxyProtocol: Whether the private link service is enabled for proxy
	//protocol or not.
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty"`

	//Fqdns: The list of Fqdn.
	Fqdns []string `json:"fqdns,omitempty"`

	//IpConfigurations: An array of private link service IP configurations.
	IpConfigurations []PrivateLinkServiceIpConfiguration_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"ipConfigurations,omitempty"`

	//LoadBalancerFrontendIpConfigurations: An array of references to the load
	//balancer IP configurations.
	LoadBalancerFrontendIpConfigurations []FrontendIPConfiguration_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"loadBalancerFrontendIpConfigurations,omitempty"`

	//Visibility: The visibility list of the private link service.
	Visibility *ResourceSet_SpecARM `json:"visibility,omitempty"`
}

type ApplicationGatewayBackendAddressPool_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Name: Name of the backend address pool that is unique within an Application
	//Gateway.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the application gateway backend address pool.
	Properties *ApplicationGatewayBackendAddressPoolPropertiesFormat_SpecARM `json:"properties,omitempty"`
}

type BackendAddressPool_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type FrontendIPConfiguration_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within the set of frontend IP
	//configurations used by the load balancer. This name can be used to access the
	//resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the load balancer probe.
	Properties *FrontendIPConfigurationPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Zones: A list of availability zones denoting the IP allocated for the resource
	//needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type InboundNatRule_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type PrivateLinkServiceIpConfiguration_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`

	//Name: The name of private link service ip configuration.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the private link service ip configuration.
	Properties *PrivateLinkServiceIpConfigurationProperties_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`
}

type ResourceSet_SpecARM struct {
	//Subscriptions: The list of subscriptions.
	Subscriptions []string `json:"subscriptions,omitempty"`
}

type SecurityRule_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type Subnet_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

type VirtualNetworkTap_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Properties: Virtual Network Tap Properties.
	Properties *VirtualNetworkTapPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

type ApplicationGatewayBackendAddressPoolPropertiesFormat_SpecARM struct {
	//BackendAddresses: Backend addresses.
	BackendAddresses []ApplicationGatewayBackendAddress_SpecARM `json:"backendAddresses,omitempty"`
}

type FrontendIPConfigurationPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	//PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	//PrivateIPAddressVersion: Whether the specific ipconfiguration is IPv4 or IPv6.
	//Default is taken as IPv4.
	PrivateIPAddressVersion *IPVersion_Spec `json:"privateIPAddressVersion,omitempty"`

	//PrivateIPAllocationMethod: The Private IP allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_Spec `json:"privateIPAllocationMethod,omitempty"`

	//PublicIPAddress: The reference to the Public IP resource.
	PublicIPAddress *PublicIPAddress_SpecARM `json:"publicIPAddress,omitempty"`

	//PublicIPPrefix: The reference to the Public IP Prefix resource.
	PublicIPPrefix *SubResource_SpecARM `json:"publicIPPrefix,omitempty"`

	//Subnet: The reference to the subnet resource.
	Subnet *Subnet_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

type PrivateLinkServiceIpConfigurationProperties_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	//Primary: Whether the ip configuration is primary or not.
	Primary *bool `json:"primary,omitempty"`

	//PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	//PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6.
	//Default is IPv4.
	PrivateIPAddressVersion *IPVersion_Spec `json:"privateIPAddressVersion,omitempty"`

	//PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_Spec `json:"privateIPAllocationMethod,omitempty"`

	//Subnet: The reference to the subnet resource.
	Subnet *Subnet_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

type VirtualNetworkTapPropertiesFormat_Spec_NetworkInterface_SubResourceEmbeddedARM struct {
	//DestinationLoadBalancerFrontEndIPConfiguration: The reference to the private IP
	//address on the internal Load Balancer that will receive the tap.
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfiguration_Spec_NetworkInterface_SubResourceEmbeddedARM `json:"destinationLoadBalancerFrontEndIPConfiguration,omitempty"`

	//DestinationPort: The VXLAN destination port that will receive the tapped traffic.
	DestinationPort *int `json:"destinationPort,omitempty"`
}

type ApplicationGatewayBackendAddress_SpecARM struct {
	//Fqdn: Fully qualified domain name (FQDN).
	Fqdn *string `json:"fqdn,omitempty"`

	//IpAddress: IP address.
	IpAddress *string `json:"ipAddress,omitempty"`
}
