// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
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

func Test_NetworkSecurityGroupsSecurityRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroupsSecurityRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNetworkSecurityGroupsSecurityRule tests if a specific instance of NetworkSecurityGroupsSecurityRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1alpha1api20201101storage.NetworkSecurityGroupsSecurityRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NetworkSecurityGroupsSecurityRule
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroupsSecurityRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroupsSecurityRule to NetworkSecurityGroupsSecurityRule via AssignPropertiesToNetworkSecurityGroupsSecurityRule & AssignPropertiesFromNetworkSecurityGroupsSecurityRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule tests if a specific instance of NetworkSecurityGroupsSecurityRule can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.NetworkSecurityGroupsSecurityRule
	err := copied.AssignPropertiesToNetworkSecurityGroupsSecurityRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroupsSecurityRule
	err = actual.AssignPropertiesFromNetworkSecurityGroupsSecurityRule(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroupsSecurityRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule
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

// Generator of NetworkSecurityGroupsSecurityRule instances for property testing - lazily instantiated by
//NetworkSecurityGroupsSecurityRuleGenerator()
var networkSecurityGroupsSecurityRuleGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRuleGenerator returns a generator of NetworkSecurityGroupsSecurityRule instances for property testing.
func NetworkSecurityGroupsSecurityRuleGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRuleGenerator != nil {
		return networkSecurityGroupsSecurityRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(generators)
	networkSecurityGroupsSecurityRuleGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule{}), generators)

	return networkSecurityGroupsSecurityRuleGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NetworkSecurityGroupsSecurityRulesSPECGenerator()
	gens["Status"] = SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()
}

func Test_NetworkSecurityGroupsSecurityRules_SPEC_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroupsSecurityRules_SPEC to NetworkSecurityGroupsSecurityRules_SPEC via AssignPropertiesToNetworkSecurityGroupsSecurityRulesSPEC & AssignPropertiesFromNetworkSecurityGroupsSecurityRulesSPEC returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRulesSPEC, NetworkSecurityGroupsSecurityRulesSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRulesSPEC tests if a specific instance of NetworkSecurityGroupsSecurityRules_SPEC can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRulesSPEC(subject NetworkSecurityGroupsSecurityRules_SPEC) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.NetworkSecurityGroupsSecurityRules_SPEC
	err := copied.AssignPropertiesToNetworkSecurityGroupsSecurityRulesSPEC(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroupsSecurityRules_SPEC
	err = actual.AssignPropertiesFromNetworkSecurityGroupsSecurityRulesSPEC(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroupsSecurityRules_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRules_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSPEC, NetworkSecurityGroupsSecurityRulesSPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSPEC runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRules_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRulesSPEC(subject NetworkSecurityGroupsSecurityRules_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRules_SPEC
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

// Generator of NetworkSecurityGroupsSecurityRules_SPEC instances for property testing - lazily instantiated by
//NetworkSecurityGroupsSecurityRulesSPECGenerator()
var networkSecurityGroupsSecurityRulesSPECGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRulesSPECGenerator returns a generator of NetworkSecurityGroupsSecurityRules_SPEC instances for property testing.
// We first initialize networkSecurityGroupsSecurityRulesSPECGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRulesSPECGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRulesSPECGenerator != nil {
		return networkSecurityGroupsSecurityRulesSPECGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSPEC(generators)
	networkSecurityGroupsSecurityRulesSPECGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRules_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSPEC(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSPEC(generators)
	networkSecurityGroupsSecurityRulesSPECGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRules_SPEC{}), generators)

	return networkSecurityGroupsSecurityRulesSPECGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSPEC(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_SpecAllow, SecurityRuleAccess_SpecDeny))
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_SpecInbound, SecurityRuleDirection_SpecOutbound))
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormatSpecProtocolAh,
		SecurityRulePropertiesFormatSpecProtocolEsp,
		SecurityRulePropertiesFormatSpecProtocolIcmp,
		SecurityRulePropertiesFormatSpecProtocolStar,
		SecurityRulePropertiesFormatSpecProtocolTcp,
		SecurityRulePropertiesFormatSpecProtocolUdp))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRulesSPEC(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpecGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpecGenerator())
}

func Test_SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded to SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded & AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded, SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded tests if a specific instance of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(subject SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err := copied.AssignPropertiesToSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err = actual.AssignPropertiesFromSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded, SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded runs a test to see if a specific instance of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(subject SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
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

// Generator of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing
//- lazily instantiated by SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()
var securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator gopter.Gen

// SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator returns a generator of SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing.
// We first initialize securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator() gopter.Gen {
	if securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator != nil {
		return securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	AddRelatedPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	return securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_StatusAllow, SecurityRuleAccess_StatusDeny))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_StatusInbound, SecurityRuleDirection_StatusOutbound))
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormatStatusProtocolAh,
		SecurityRulePropertiesFormatStatusProtocolEsp,
		SecurityRulePropertiesFormatStatusProtocolIcmp,
		SecurityRulePropertiesFormatStatusProtocolStar,
		SecurityRulePropertiesFormatStatusProtocolTcp,
		SecurityRulePropertiesFormatStatusProtocolUdp))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator())
}

func Test_ApplicationSecurityGroup_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationSecurityGroup_Spec to ApplicationSecurityGroup_Spec via AssignPropertiesToApplicationSecurityGroupSpec & AssignPropertiesFromApplicationSecurityGroupSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationSecurityGroupSpec, ApplicationSecurityGroupSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationSecurityGroupSpec tests if a specific instance of ApplicationSecurityGroup_Spec can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForApplicationSecurityGroupSpec(subject ApplicationSecurityGroup_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.ApplicationSecurityGroup_Spec
	err := copied.AssignPropertiesToApplicationSecurityGroupSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationSecurityGroup_Spec
	err = actual.AssignPropertiesFromApplicationSecurityGroupSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationSecurityGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec, ApplicationSecurityGroupSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec runs a test to see if a specific instance of ApplicationSecurityGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec(subject ApplicationSecurityGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_Spec
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

// Generator of ApplicationSecurityGroup_Spec instances for property testing - lazily instantiated by
//ApplicationSecurityGroupSpecGenerator()
var applicationSecurityGroupSpecGenerator gopter.Gen

// ApplicationSecurityGroupSpecGenerator returns a generator of ApplicationSecurityGroup_Spec instances for property testing.
func ApplicationSecurityGroupSpecGenerator() gopter.Gen {
	if applicationSecurityGroupSpecGenerator != nil {
		return applicationSecurityGroupSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec(generators)
	applicationSecurityGroupSpecGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_Spec{}), generators)

	return applicationSecurityGroupSpecGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded to ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via AssignPropertiesToApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded & AssignPropertiesFromApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded, ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded tests if a specific instance of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(subject ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20201101storage.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err := copied.AssignPropertiesToApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err = actual.AssignPropertiesFromApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded, ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(subject ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for
//property testing - lazily instantiated by
//ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator()
var applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(generators)
	applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}