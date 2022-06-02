// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Vault_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Vault_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVault_SpecARM, Vault_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVault_SpecARM runs a test to see if a specific instance of Vault_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVault_SpecARM(subject Vault_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Vault_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Vault_SpecARM instances for property testing - lazily instantiated by Vault_SpecARMGenerator()
var vault_SpecARMGenerator gopter.Gen

// Vault_SpecARMGenerator returns a generator of Vault_SpecARM instances for property testing.
// We first initialize vault_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Vault_SpecARMGenerator() gopter.Gen {
	if vault_SpecARMGenerator != nil {
		return vault_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVault_SpecARM(generators)
	vault_SpecARMGenerator = gen.Struct(reflect.TypeOf(Vault_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVault_SpecARM(generators)
	AddRelatedPropertyGeneratorsForVault_SpecARM(generators)
	vault_SpecARMGenerator = gen.Struct(reflect.TypeOf(Vault_SpecARM{}), generators)

	return vault_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForVault_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVault_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVault_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVault_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VaultPropertiesARMGenerator())
}

func Test_VaultPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VaultPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVaultPropertiesARM, VaultPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVaultPropertiesARM runs a test to see if a specific instance of VaultPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVaultPropertiesARM(subject VaultPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VaultPropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VaultPropertiesARM instances for property testing - lazily instantiated by VaultPropertiesARMGenerator()
var vaultPropertiesARMGenerator gopter.Gen

// VaultPropertiesARMGenerator returns a generator of VaultPropertiesARM instances for property testing.
// We first initialize vaultPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VaultPropertiesARMGenerator() gopter.Gen {
	if vaultPropertiesARMGenerator != nil {
		return vaultPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVaultPropertiesARM(generators)
	vaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(VaultPropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVaultPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForVaultPropertiesARM(generators)
	vaultPropertiesARMGenerator = gen.Struct(reflect.TypeOf(VaultPropertiesARM{}), generators)

	return vaultPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForVaultPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(VaultProperties_CreateMode_Default, VaultProperties_CreateMode_Recover))
	gens["EnablePurgeProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableRbacAuthorization"] = gen.PtrOf(gen.Bool())
	gens["EnableSoftDelete"] = gen.PtrOf(gen.Bool())
	gens["EnabledForDeployment"] = gen.PtrOf(gen.Bool())
	gens["EnabledForDiskEncryption"] = gen.PtrOf(gen.Bool())
	gens["EnabledForTemplateDeployment"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(VaultProperties_ProvisioningState_RegisteringDns, VaultProperties_ProvisioningState_Succeeded))
	gens["SoftDeleteRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["VaultUri"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVaultPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVaultPropertiesARM(gens map[string]gopter.Gen) {
	gens["AccessPolicies"] = gen.SliceOf(AccessPolicyEntryARMGenerator())
	gens["NetworkAcls"] = gen.PtrOf(NetworkRuleSetARMGenerator())
	gens["Sku"] = gen.PtrOf(SkuARMGenerator())
}

func Test_AccessPolicyEntryARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AccessPolicyEntryARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAccessPolicyEntryARM, AccessPolicyEntryARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAccessPolicyEntryARM runs a test to see if a specific instance of AccessPolicyEntryARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAccessPolicyEntryARM(subject AccessPolicyEntryARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AccessPolicyEntryARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AccessPolicyEntryARM instances for property testing - lazily instantiated by
// AccessPolicyEntryARMGenerator()
var accessPolicyEntryARMGenerator gopter.Gen

// AccessPolicyEntryARMGenerator returns a generator of AccessPolicyEntryARM instances for property testing.
// We first initialize accessPolicyEntryARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AccessPolicyEntryARMGenerator() gopter.Gen {
	if accessPolicyEntryARMGenerator != nil {
		return accessPolicyEntryARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAccessPolicyEntryARM(generators)
	accessPolicyEntryARMGenerator = gen.Struct(reflect.TypeOf(AccessPolicyEntryARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAccessPolicyEntryARM(generators)
	AddRelatedPropertyGeneratorsForAccessPolicyEntryARM(generators)
	accessPolicyEntryARMGenerator = gen.Struct(reflect.TypeOf(AccessPolicyEntryARM{}), generators)

	return accessPolicyEntryARMGenerator
}

// AddIndependentPropertyGeneratorsForAccessPolicyEntryARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAccessPolicyEntryARM(gens map[string]gopter.Gen) {
	gens["ApplicationId"] = gen.PtrOf(gen.AlphaString())
	gens["ObjectId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAccessPolicyEntryARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAccessPolicyEntryARM(gens map[string]gopter.Gen) {
	gens["Permissions"] = gen.PtrOf(PermissionsARMGenerator())
}

func Test_NetworkRuleSetARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkRuleSetARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkRuleSetARM, NetworkRuleSetARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkRuleSetARM runs a test to see if a specific instance of NetworkRuleSetARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkRuleSetARM(subject NetworkRuleSetARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkRuleSetARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NetworkRuleSetARM instances for property testing - lazily instantiated by NetworkRuleSetARMGenerator()
var networkRuleSetARMGenerator gopter.Gen

// NetworkRuleSetARMGenerator returns a generator of NetworkRuleSetARM instances for property testing.
// We first initialize networkRuleSetARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkRuleSetARMGenerator() gopter.Gen {
	if networkRuleSetARMGenerator != nil {
		return networkRuleSetARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSetARM(generators)
	networkRuleSetARMGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSetARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkRuleSetARM(generators)
	AddRelatedPropertyGeneratorsForNetworkRuleSetARM(generators)
	networkRuleSetARMGenerator = gen.Struct(reflect.TypeOf(NetworkRuleSetARM{}), generators)

	return networkRuleSetARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkRuleSetARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkRuleSetARM(gens map[string]gopter.Gen) {
	gens["Bypass"] = gen.PtrOf(gen.OneConstOf(NetworkRuleSet_Bypass_AzureServices, NetworkRuleSet_Bypass_None))
	gens["DefaultAction"] = gen.PtrOf(gen.OneConstOf(NetworkRuleSet_DefaultAction_Allow, NetworkRuleSet_DefaultAction_Deny))
}

// AddRelatedPropertyGeneratorsForNetworkRuleSetARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkRuleSetARM(gens map[string]gopter.Gen) {
	gens["IpRules"] = gen.SliceOf(IPRuleARMGenerator())
	gens["VirtualNetworkRules"] = gen.SliceOf(VirtualNetworkRuleARMGenerator())
}

func Test_SkuARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SkuARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuARM, SkuARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuARM runs a test to see if a specific instance of SkuARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuARM(subject SkuARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SkuARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SkuARM instances for property testing - lazily instantiated by SkuARMGenerator()
var skuARMGenerator gopter.Gen

// SkuARMGenerator returns a generator of SkuARM instances for property testing.
func SkuARMGenerator() gopter.Gen {
	if skuARMGenerator != nil {
		return skuARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuARM(generators)
	skuARMGenerator = gen.Struct(reflect.TypeOf(SkuARM{}), generators)

	return skuARMGenerator
}

// AddIndependentPropertyGeneratorsForSkuARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuARM(gens map[string]gopter.Gen) {
	gens["Family"] = gen.PtrOf(gen.OneConstOf(Sku_Family_A))
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_Premium, Sku_Name_Standard))
}

func Test_IPRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IPRuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIPRuleARM, IPRuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIPRuleARM runs a test to see if a specific instance of IPRuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIPRuleARM(subject IPRuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IPRuleARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IPRuleARM instances for property testing - lazily instantiated by IPRuleARMGenerator()
var ipRuleARMGenerator gopter.Gen

// IPRuleARMGenerator returns a generator of IPRuleARM instances for property testing.
func IPRuleARMGenerator() gopter.Gen {
	if ipRuleARMGenerator != nil {
		return ipRuleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIPRuleARM(generators)
	ipRuleARMGenerator = gen.Struct(reflect.TypeOf(IPRuleARM{}), generators)

	return ipRuleARMGenerator
}

// AddIndependentPropertyGeneratorsForIPRuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIPRuleARM(gens map[string]gopter.Gen) {
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_PermissionsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PermissionsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPermissionsARM, PermissionsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPermissionsARM runs a test to see if a specific instance of PermissionsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPermissionsARM(subject PermissionsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PermissionsARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PermissionsARM instances for property testing - lazily instantiated by PermissionsARMGenerator()
var permissionsARMGenerator gopter.Gen

// PermissionsARMGenerator returns a generator of PermissionsARM instances for property testing.
func PermissionsARMGenerator() gopter.Gen {
	if permissionsARMGenerator != nil {
		return permissionsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPermissionsARM(generators)
	permissionsARMGenerator = gen.Struct(reflect.TypeOf(PermissionsARM{}), generators)

	return permissionsARMGenerator
}

// AddIndependentPropertyGeneratorsForPermissionsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPermissionsARM(gens map[string]gopter.Gen) {
	gens["Certificates"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Certificates_Backup,
		Permissions_Certificates_Create,
		Permissions_Certificates_Delete,
		Permissions_Certificates_Deleteissuers,
		Permissions_Certificates_Get,
		Permissions_Certificates_Getissuers,
		Permissions_Certificates_Import,
		Permissions_Certificates_List,
		Permissions_Certificates_Listissuers,
		Permissions_Certificates_Managecontacts,
		Permissions_Certificates_Manageissuers,
		Permissions_Certificates_Purge,
		Permissions_Certificates_Recover,
		Permissions_Certificates_Restore,
		Permissions_Certificates_Setissuers,
		Permissions_Certificates_Update))
	gens["Keys"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Keys_Backup,
		Permissions_Keys_Create,
		Permissions_Keys_Decrypt,
		Permissions_Keys_Delete,
		Permissions_Keys_Encrypt,
		Permissions_Keys_Get,
		Permissions_Keys_Import,
		Permissions_Keys_List,
		Permissions_Keys_Purge,
		Permissions_Keys_Recover,
		Permissions_Keys_Release,
		Permissions_Keys_Restore,
		Permissions_Keys_Sign,
		Permissions_Keys_UnwrapKey,
		Permissions_Keys_Update,
		Permissions_Keys_Verify,
		Permissions_Keys_WrapKey))
	gens["Secrets"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Secrets_Backup,
		Permissions_Secrets_Delete,
		Permissions_Secrets_Get,
		Permissions_Secrets_List,
		Permissions_Secrets_Purge,
		Permissions_Secrets_Recover,
		Permissions_Secrets_Restore,
		Permissions_Secrets_Set))
	gens["Storage"] = gen.SliceOf(gen.OneConstOf(
		Permissions_Storage_Backup,
		Permissions_Storage_Delete,
		Permissions_Storage_Deletesas,
		Permissions_Storage_Get,
		Permissions_Storage_Getsas,
		Permissions_Storage_List,
		Permissions_Storage_Listsas,
		Permissions_Storage_Purge,
		Permissions_Storage_Recover,
		Permissions_Storage_Regeneratekey,
		Permissions_Storage_Restore,
		Permissions_Storage_Set,
		Permissions_Storage_Setsas,
		Permissions_Storage_Update))
}

func Test_VirtualNetworkRuleARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkRuleARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkRuleARM, VirtualNetworkRuleARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkRuleARM runs a test to see if a specific instance of VirtualNetworkRuleARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkRuleARM(subject VirtualNetworkRuleARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkRuleARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of VirtualNetworkRuleARM instances for property testing - lazily instantiated by
// VirtualNetworkRuleARMGenerator()
var virtualNetworkRuleARMGenerator gopter.Gen

// VirtualNetworkRuleARMGenerator returns a generator of VirtualNetworkRuleARM instances for property testing.
func VirtualNetworkRuleARMGenerator() gopter.Gen {
	if virtualNetworkRuleARMGenerator != nil {
		return virtualNetworkRuleARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkRuleARM(generators)
	virtualNetworkRuleARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkRuleARM{}), generators)

	return virtualNetworkRuleARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkRuleARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkRuleARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IgnoreMissingVnetServiceEndpoint"] = gen.PtrOf(gen.Bool())
}