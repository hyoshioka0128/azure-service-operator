// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101preview

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=servers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={servers/status,servers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/Servers.json
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Server_Spec   `json:"spec,omitempty"`
	Status            Server_STATUS `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/Servers.json
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-11-01-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-11-01-preview")

type Server_STATUS struct {
	v1alpha1.ResourceStatus `json:",inline,omitempty"`
	AtProvider              ServerObservation `json:"atProvider,omitempty"`
}

type Server_Spec struct {
	v1alpha1.ResourceSpec `json:",inline,omitempty"`
	ForProvider           ServerParameters `json:"forProvider,omitempty"`
}

type ServerObservation struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: The Azure Active Directory identity of the server.
	Administrators *ServerExternalAdministrator_STATUS `json:"administrators,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of the server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity_STATUS `json:"identity,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// Kind: Kind of sql server. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *ServerProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// State: The state of the server.
	State *string `json:"state,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`

	// WorkspaceFeature: Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature *ServerProperties_WorkspaceFeature_STATUS `json:"workspaceFeature,omitempty"`
}

type ServerParameters struct {
	// AdministratorLogin: Administrator username for the server. Once created it cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	// Administrators: The Azure Active Directory identity of the server.
	Administrators *ServerExternalAdministrator `json:"administrators,omitempty"`
	AzureName      string                       `json:"azureName,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of the server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	// KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	// Kind: Kind of sql server. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	// +kubebuilder:validation:Required
	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`

	// Name: Resource name.
	Name string `json:"name,omitempty"`

	// PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this server.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess       *ServerProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
	ResourceGroupName         string                                `json:"resourceGroupName,omitempty"`
	ResourceGroupNameRef      *v1alpha1.Reference                   `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector *v1alpha1.Selector                    `json:"resourceGroupNameSelector,omitempty"`

	// State: The state of the server.
	State *string `json:"state,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Version: The version of the server.
	Version *string `json:"version,omitempty"`

	// WorkspaceFeature: Whether or not existing server has a workspace created and if it allows connection from workspace
	WorkspaceFeature *ServerProperties_WorkspaceFeature `json:"workspaceFeature,omitempty"`
}

type ResourceIdentity struct {
	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// TenantId: The Azure Active Directory tenant id.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type *ResourceIdentity_Type `json:"type,omitempty"`
}

type ResourceIdentity_STATUS struct {
	// PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The Azure Active Directory tenant id.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active
	// Directory principal for the resource.
	Type *ResourceIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: The resource ids of the user assigned identities to use
	UserAssignedIdentities map[string]UserIdentity_STATUS `json:"userAssignedIdentities,omitempty"`
}

type ServerExternalAdministrator struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministrator_AdministratorType `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministrator_PrincipalType `json:"principalType,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

type ServerExternalAdministrator_STATUS struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministrator_AdministratorType_STATUS `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministrator_PrincipalType_STATUS `json:"principalType,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

type ServerPrivateEndpointConnection struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Properties: Private endpoint connection properties
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`
}

type ServerPrivateEndpointConnection_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Properties: Private endpoint connection properties
	Properties *PrivateEndpointConnectionProperties_STATUS `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerProperties_PublicNetworkAccess string

const (
	ServerProperties_PublicNetworkAccess_Disabled = ServerProperties_PublicNetworkAccess("Disabled")
	ServerProperties_PublicNetworkAccess_Enabled  = ServerProperties_PublicNetworkAccess("Enabled")
)

type ServerProperties_PublicNetworkAccess_STATUS string

const (
	ServerProperties_PublicNetworkAccess_Disabled_STATUS = ServerProperties_PublicNetworkAccess_STATUS("Disabled")
	ServerProperties_PublicNetworkAccess_Enabled_STATUS  = ServerProperties_PublicNetworkAccess_STATUS("Enabled")
)

// +kubebuilder:validation:Enum={"Connected","Disconnected"}
type ServerProperties_WorkspaceFeature string

const (
	ServerProperties_WorkspaceFeature_Connected    = ServerProperties_WorkspaceFeature("Connected")
	ServerProperties_WorkspaceFeature_Disconnected = ServerProperties_WorkspaceFeature("Disconnected")
)

type ServerProperties_WorkspaceFeature_STATUS string

const (
	ServerProperties_WorkspaceFeature_Connected_STATUS    = ServerProperties_WorkspaceFeature_STATUS("Connected")
	ServerProperties_WorkspaceFeature_Disconnected_STATUS = ServerProperties_WorkspaceFeature_STATUS("Disconnected")
)

type PrivateEndpointConnectionProperties struct {
	// PrivateEndpoint: Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: State of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProperties_ProvisioningState `json:"provisioningState,omitempty"`
}

type PrivateEndpointConnectionProperties_STATUS struct {
	// PrivateEndpoint: Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty_STATUS `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty_STATUS `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: State of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ResourceIdentity_Type string

const (
	ResourceIdentity_Type_None                       = ResourceIdentity_Type("None")
	ResourceIdentity_Type_SystemAssigned             = ResourceIdentity_Type("SystemAssigned")
	ResourceIdentity_Type_SystemAssignedUserAssigned = ResourceIdentity_Type("SystemAssigned,UserAssigned")
	ResourceIdentity_Type_UserAssigned               = ResourceIdentity_Type("UserAssigned")
)

type ResourceIdentity_Type_STATUS string

const (
	ResourceIdentity_Type_None_STATUS                       = ResourceIdentity_Type_STATUS("None")
	ResourceIdentity_Type_SystemAssigned_STATUS             = ResourceIdentity_Type_STATUS("SystemAssigned")
	ResourceIdentity_Type_SystemAssignedUserAssigned_STATUS = ResourceIdentity_Type_STATUS("SystemAssigned,UserAssigned")
	ResourceIdentity_Type_UserAssigned_STATUS               = ResourceIdentity_Type_STATUS("UserAssigned")
)

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type ServerExternalAdministrator_AdministratorType string

const ServerExternalAdministrator_AdministratorType_ActiveDirectory = ServerExternalAdministrator_AdministratorType("ActiveDirectory")

type ServerExternalAdministrator_AdministratorType_STATUS string

const ServerExternalAdministrator_AdministratorType_ActiveDirectory_STATUS = ServerExternalAdministrator_AdministratorType_STATUS("ActiveDirectory")

// +kubebuilder:validation:Enum={"Application","Group","User"}
type ServerExternalAdministrator_PrincipalType string

const (
	ServerExternalAdministrator_PrincipalType_Application = ServerExternalAdministrator_PrincipalType("Application")
	ServerExternalAdministrator_PrincipalType_Group       = ServerExternalAdministrator_PrincipalType("Group")
	ServerExternalAdministrator_PrincipalType_User        = ServerExternalAdministrator_PrincipalType("User")
)

type ServerExternalAdministrator_PrincipalType_STATUS string

const (
	ServerExternalAdministrator_PrincipalType_Application_STATUS = ServerExternalAdministrator_PrincipalType_STATUS("Application")
	ServerExternalAdministrator_PrincipalType_Group_STATUS       = ServerExternalAdministrator_PrincipalType_STATUS("Group")
	ServerExternalAdministrator_PrincipalType_User_STATUS        = ServerExternalAdministrator_PrincipalType_STATUS("User")
)

type UserIdentity_STATUS struct {
	// ClientId: The Azure Active Directory client id.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`
}

// +kubebuilder:validation:Enum={"Approving","Dropping","Failed","Ready","Rejecting"}
type PrivateEndpointConnectionProperties_ProvisioningState string

const (
	PrivateEndpointConnectionProperties_ProvisioningState_Approving = PrivateEndpointConnectionProperties_ProvisioningState("Approving")
	PrivateEndpointConnectionProperties_ProvisioningState_Dropping  = PrivateEndpointConnectionProperties_ProvisioningState("Dropping")
	PrivateEndpointConnectionProperties_ProvisioningState_Failed    = PrivateEndpointConnectionProperties_ProvisioningState("Failed")
	PrivateEndpointConnectionProperties_ProvisioningState_Ready     = PrivateEndpointConnectionProperties_ProvisioningState("Ready")
	PrivateEndpointConnectionProperties_ProvisioningState_Rejecting = PrivateEndpointConnectionProperties_ProvisioningState("Rejecting")
)

type PrivateEndpointConnectionProperties_ProvisioningState_STATUS string

const (
	PrivateEndpointConnectionProperties_ProvisioningState_Approving_STATUS = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Approving")
	PrivateEndpointConnectionProperties_ProvisioningState_Dropping_STATUS  = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Dropping")
	PrivateEndpointConnectionProperties_ProvisioningState_Failed_STATUS    = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Failed")
	PrivateEndpointConnectionProperties_ProvisioningState_Ready_STATUS     = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Ready")
	PrivateEndpointConnectionProperties_ProvisioningState_Rejecting_STATUS = PrivateEndpointConnectionProperties_ProvisioningState_STATUS("Rejecting")
)

type PrivateEndpointProperty struct {
	// Id: Resource id of the private endpoint.
	Id *string `json:"id,omitempty"`
}

type PrivateEndpointProperty_STATUS struct {
	// Id: Resource id of the private endpoint.
	Id *string `json:"id,omitempty"`
}

type PrivateLinkServiceConnectionStateProperty struct {
	// ActionsRequired: The actions required for private link service connection.
	ActionsRequired *PrivateLinkServiceConnectionStateProperty_ActionsRequired `json:"actionsRequired,omitempty"`

	// +kubebuilder:validation:Required
	// Description: The private link service connection description.
	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	// Status: The private link service connection status.
	Status *PrivateLinkServiceConnectionStateProperty_Status `json:"status,omitempty"`
}

type PrivateLinkServiceConnectionStateProperty_STATUS struct {
	// ActionsRequired: The actions required for private link service connection.
	ActionsRequired *PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS `json:"actionsRequired,omitempty"`

	// Description: The private link service connection description.
	Description *string `json:"description,omitempty"`

	// Status: The private link service connection status.
	Status *PrivateLinkServiceConnectionStateProperty_Status_STATUS `json:"status,omitempty"`
}

// +kubebuilder:validation:Enum={"None"}
type PrivateLinkServiceConnectionStateProperty_ActionsRequired string

const PrivateLinkServiceConnectionStateProperty_ActionsRequired_None = PrivateLinkServiceConnectionStateProperty_ActionsRequired("None")

type PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS string

const PrivateLinkServiceConnectionStateProperty_ActionsRequired_None_STATUS = PrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS("None")

// +kubebuilder:validation:Enum={"Approved","Disconnected","Pending","Rejected"}
type PrivateLinkServiceConnectionStateProperty_Status string

const (
	PrivateLinkServiceConnectionStateProperty_Status_Approved     = PrivateLinkServiceConnectionStateProperty_Status("Approved")
	PrivateLinkServiceConnectionStateProperty_Status_Disconnected = PrivateLinkServiceConnectionStateProperty_Status("Disconnected")
	PrivateLinkServiceConnectionStateProperty_Status_Pending      = PrivateLinkServiceConnectionStateProperty_Status("Pending")
	PrivateLinkServiceConnectionStateProperty_Status_Rejected     = PrivateLinkServiceConnectionStateProperty_Status("Rejected")
)

type PrivateLinkServiceConnectionStateProperty_Status_STATUS string

const (
	PrivateLinkServiceConnectionStateProperty_Status_Approved_STATUS     = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Approved")
	PrivateLinkServiceConnectionStateProperty_Status_Disconnected_STATUS = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Disconnected")
	PrivateLinkServiceConnectionStateProperty_Status_Pending_STATUS      = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Pending")
	PrivateLinkServiceConnectionStateProperty_Status_Rejected_STATUS     = PrivateLinkServiceConnectionStateProperty_Status_STATUS("Rejected")
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
