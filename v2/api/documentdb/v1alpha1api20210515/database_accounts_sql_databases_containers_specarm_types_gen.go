// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsSqlDatabasesContainers_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties to create and update Azure Cosmos DB container.
	Properties SqlContainerProperties_SpecARM `json:"properties"`
	Tags       map[string]string              `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainers_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (specarm DatabaseAccountsSqlDatabasesContainers_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm DatabaseAccountsSqlDatabasesContainers_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
func (specarm DatabaseAccountsSqlDatabasesContainers_SPECARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
}

type SqlContainerProperties_SpecARM struct {
	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_SpecARM `json:"options,omitempty"`

	//Resource: The standard JSON format of a container
	Resource SqlContainerResource_SpecARM `json:"resource"`
}

type SqlContainerResource_SpecARM struct {
	//AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	//ConflictResolutionPolicy: The conflict resolution policy for the container.
	ConflictResolutionPolicy *ConflictResolutionPolicy_SpecARM `json:"conflictResolutionPolicy,omitempty"`

	//DefaultTtl: Default time to live
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	//Id: Name of the Cosmos DB SQL container
	Id string `json:"id"`

	//IndexingPolicy: The configuration of the indexing policy. By default, the
	//indexing is automatic for all document paths within the container
	IndexingPolicy *IndexingPolicy_SpecARM `json:"indexingPolicy,omitempty"`

	//PartitionKey: The configuration of the partition key to be used for partitioning
	//data into multiple partitions
	PartitionKey *ContainerPartitionKey_SpecARM `json:"partitionKey,omitempty"`

	//UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness
	//constraints on documents in the collection in the Azure Cosmos DB service.
	UniqueKeyPolicy *UniqueKeyPolicy_SpecARM `json:"uniqueKeyPolicy,omitempty"`
}

type ConflictResolutionPolicy_SpecARM struct {
	//ConflictResolutionPath: The conflict resolution path in the case of
	//LastWriterWins mode.
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty"`

	//ConflictResolutionProcedure: The procedure to resolve conflicts in the case of
	//custom mode.
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty"`

	//Mode: Indicates the conflict resolution mode.
	Mode *ConflictResolutionPolicy_Mode_Spec `json:"mode,omitempty"`
}

type ContainerPartitionKey_SpecARM struct {
	//Kind: Indicates the kind of algorithm used for partitioning. For MultiHash,
	//multiple partition keys (upto three maximum) are supported for container create
	Kind *ContainerPartitionKey_Kind_Spec `json:"kind,omitempty"`

	//Paths: List of paths using which data within the container can be partitioned
	Paths []string `json:"paths,omitempty"`

	//Version: Indicates the version of the partition key definition
	Version *int `json:"version,omitempty"`
}

type IndexingPolicy_SpecARM struct {
	//Automatic: Indicates if the indexing policy is automatic
	Automatic *bool `json:"automatic,omitempty"`

	//CompositeIndexes: List of composite path list
	CompositeIndexes [][]CompositePath_SpecARM `json:"compositeIndexes,omitempty"`

	//ExcludedPaths: List of paths to exclude from indexing
	ExcludedPaths []ExcludedPath_SpecARM `json:"excludedPaths,omitempty"`

	//IncludedPaths: List of paths to include in the indexing
	IncludedPaths []IncludedPath_SpecARM `json:"includedPaths,omitempty"`

	//IndexingMode: Indicates the indexing mode.
	IndexingMode *IndexingPolicy_IndexingMode_Spec `json:"indexingMode,omitempty"`

	//SpatialIndexes: List of spatial specifics
	SpatialIndexes []SpatialSpec_SpecARM `json:"spatialIndexes,omitempty"`
}

type UniqueKeyPolicy_SpecARM struct {
	//UniqueKeys: List of unique keys on that enforces uniqueness constraint on
	//documents in the collection in the Azure Cosmos DB service.
	UniqueKeys []UniqueKey_SpecARM `json:"uniqueKeys,omitempty"`
}

type CompositePath_SpecARM struct {
	//Order: Sort order for composite paths.
	Order *CompositePath_Order_Spec `json:"order,omitempty"`

	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"Custom","LastWriterWins"}
type ConflictResolutionPolicy_Mode_Spec string

const (
	ConflictResolutionPolicy_Mode_SpecCustom         = ConflictResolutionPolicy_Mode_Spec("Custom")
	ConflictResolutionPolicy_Mode_SpecLastWriterWins = ConflictResolutionPolicy_Mode_Spec("LastWriterWins")
)

// +kubebuilder:validation:Enum={"Hash","MultiHash","Range"}
type ContainerPartitionKey_Kind_Spec string

const (
	ContainerPartitionKey_Kind_SpecHash      = ContainerPartitionKey_Kind_Spec("Hash")
	ContainerPartitionKey_Kind_SpecMultiHash = ContainerPartitionKey_Kind_Spec("MultiHash")
	ContainerPartitionKey_Kind_SpecRange     = ContainerPartitionKey_Kind_Spec("Range")
)

type ExcludedPath_SpecARM struct {
	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`
}

type IncludedPath_SpecARM struct {
	//Indexes: List of indexes for this path
	Indexes []Indexes_SpecARM `json:"indexes,omitempty"`

	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"consistent","lazy","none"}
type IndexingPolicy_IndexingMode_Spec string

const (
	IndexingPolicy_IndexingMode_SpecConsistent = IndexingPolicy_IndexingMode_Spec("consistent")
	IndexingPolicy_IndexingMode_SpecLazy       = IndexingPolicy_IndexingMode_Spec("lazy")
	IndexingPolicy_IndexingMode_SpecNone       = IndexingPolicy_IndexingMode_Spec("none")
)

type SpatialSpec_SpecARM struct {
	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`

	//Types: List of path's spatial type
	Types []SpatialType_Spec `json:"types,omitempty"`
}

type UniqueKey_SpecARM struct {
	//Paths: List of paths must be unique for each document in the Azure Cosmos DB
	//service
	Paths []string `json:"paths,omitempty"`
}

// +kubebuilder:validation:Enum={"ascending","descending"}
type CompositePath_Order_Spec string

const (
	CompositePath_Order_SpecAscending  = CompositePath_Order_Spec("ascending")
	CompositePath_Order_SpecDescending = CompositePath_Order_Spec("descending")
)

type Indexes_SpecARM struct {
	//DataType: The datatype for which the indexing behavior is applied to.
	DataType *Indexes_DataType_Spec `json:"dataType,omitempty"`

	//Kind: Indicates the type of index.
	Kind *Indexes_Kind_Spec `json:"kind,omitempty"`

	//Precision: The precision of the index. -1 is maximum precision.
	Precision *int `json:"precision,omitempty"`
}

// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Point","Polygon"}
type SpatialType_Spec string

const (
	SpatialType_SpecLineString   = SpatialType_Spec("LineString")
	SpatialType_SpecMultiPolygon = SpatialType_Spec("MultiPolygon")
	SpatialType_SpecPoint        = SpatialType_Spec("Point")
	SpatialType_SpecPolygon      = SpatialType_Spec("Polygon")
)

// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Number","Point","Polygon","String"}
type Indexes_DataType_Spec string

const (
	Indexes_DataType_SpecLineString   = Indexes_DataType_Spec("LineString")
	Indexes_DataType_SpecMultiPolygon = Indexes_DataType_Spec("MultiPolygon")
	Indexes_DataType_SpecNumber       = Indexes_DataType_Spec("Number")
	Indexes_DataType_SpecPoint        = Indexes_DataType_Spec("Point")
	Indexes_DataType_SpecPolygon      = Indexes_DataType_Spec("Polygon")
	Indexes_DataType_SpecString       = Indexes_DataType_Spec("String")
)

// +kubebuilder:validation:Enum={"Hash","Range","Spatial"}
type Indexes_Kind_Spec string

const (
	Indexes_Kind_SpecHash    = Indexes_Kind_Spec("Hash")
	Indexes_Kind_SpecRange   = Indexes_Kind_Spec("Range")
	Indexes_Kind_SpecSpatial = Indexes_Kind_Spec("Spatial")
)
