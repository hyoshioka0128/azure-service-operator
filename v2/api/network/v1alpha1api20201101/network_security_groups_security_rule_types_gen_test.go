// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import (
	"encoding/json"
	alpha20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
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
	var hub v20201101s.NetworkSecurityGroupsSecurityRule
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
	var other alpha20201101s.NetworkSecurityGroupsSecurityRule
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

	// Check for a match
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
// NetworkSecurityGroupsSecurityRuleGenerator()
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
	gens["Spec"] = NetworkSecurityGroupsSecurityRule_SpecGenerator()
	gens["Status"] = NetworkSecurityGroupsSecurityRule_STATUSGenerator()
}

func Test_NetworkSecurityGroupsSecurityRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroupsSecurityRule_STATUS to NetworkSecurityGroupsSecurityRule_STATUS via AssignPropertiesToNetworkSecurityGroupsSecurityRule_STATUS & AssignPropertiesFromNetworkSecurityGroupsSecurityRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule_STATUS, NetworkSecurityGroupsSecurityRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule_STATUS tests if a specific instance of NetworkSecurityGroupsSecurityRule_STATUS can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule_STATUS(subject NetworkSecurityGroupsSecurityRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.NetworkSecurityGroupsSecurityRule_STATUS
	err := copied.AssignPropertiesToNetworkSecurityGroupsSecurityRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroupsSecurityRule_STATUS
	err = actual.AssignPropertiesFromNetworkSecurityGroupsSecurityRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroupsSecurityRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_STATUS, NetworkSecurityGroupsSecurityRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_STATUS runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_STATUS(subject NetworkSecurityGroupsSecurityRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule_STATUS
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

// Generator of NetworkSecurityGroupsSecurityRule_STATUS instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRule_STATUSGenerator()
var networkSecurityGroupsSecurityRule_STATUSGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRule_STATUSGenerator returns a generator of NetworkSecurityGroupsSecurityRule_STATUS instances for property testing.
// We first initialize networkSecurityGroupsSecurityRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRule_STATUSGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRule_STATUSGenerator != nil {
		return networkSecurityGroupsSecurityRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(generators)
	networkSecurityGroupsSecurityRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(generators)
	networkSecurityGroupsSecurityRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_STATUS{}), generators)

	return networkSecurityGroupsSecurityRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_Allow_STATUS, SecurityRuleAccess_Deny_STATUS))
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_Inbound_STATUS, SecurityRuleDirection_Outbound_STATUS))
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_Ah_STATUS,
		SecurityRulePropertiesFormat_Protocol_Esp_STATUS,
		SecurityRulePropertiesFormat_Protocol_Icmp_STATUS,
		SecurityRulePropertiesFormat_Protocol_Star_STATUS,
		SecurityRulePropertiesFormat_Protocol_Tcp_STATUS,
		SecurityRulePropertiesFormat_Protocol_Udp_STATUS))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_Deleting_STATUS,
		ProvisioningState_Failed_STATUS,
		ProvisioningState_Succeeded_STATUS,
		ProvisioningState_Updating_STATUS))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator())
}

func Test_NetworkSecurityGroupsSecurityRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NetworkSecurityGroupsSecurityRule_Spec to NetworkSecurityGroupsSecurityRule_Spec via AssignPropertiesToNetworkSecurityGroupsSecurityRule_Spec & AssignPropertiesFromNetworkSecurityGroupsSecurityRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule_Spec, NetworkSecurityGroupsSecurityRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule_Spec tests if a specific instance of NetworkSecurityGroupsSecurityRule_Spec can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForNetworkSecurityGroupsSecurityRule_Spec(subject NetworkSecurityGroupsSecurityRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.NetworkSecurityGroupsSecurityRule_Spec
	err := copied.AssignPropertiesToNetworkSecurityGroupsSecurityRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NetworkSecurityGroupsSecurityRule_Spec
	err = actual.AssignPropertiesFromNetworkSecurityGroupsSecurityRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NetworkSecurityGroupsSecurityRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec, NetworkSecurityGroupsSecurityRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec(subject NetworkSecurityGroupsSecurityRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule_Spec
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

// Generator of NetworkSecurityGroupsSecurityRule_Spec instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRule_SpecGenerator()
var networkSecurityGroupsSecurityRule_SpecGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRule_SpecGenerator returns a generator of NetworkSecurityGroupsSecurityRule_Spec instances for property testing.
// We first initialize networkSecurityGroupsSecurityRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRule_SpecGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRule_SpecGenerator != nil {
		return networkSecurityGroupsSecurityRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	networkSecurityGroupsSecurityRule_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	networkSecurityGroupsSecurityRule_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_Spec{}), generators)

	return networkSecurityGroupsSecurityRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.OneConstOf(SecurityRuleAccess_Allow, SecurityRuleAccess_Deny))
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.OneConstOf(SecurityRuleDirection_Inbound, SecurityRuleDirection_Outbound))
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(
		SecurityRulePropertiesFormat_Protocol_Ah,
		SecurityRulePropertiesFormat_Protocol_Esp,
		SecurityRulePropertiesFormat_Protocol_Icmp,
		SecurityRulePropertiesFormat_Protocol_Star,
		SecurityRulePropertiesFormat_Protocol_Tcp,
		SecurityRulePropertiesFormat_Protocol_Udp))
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpecGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpecGenerator())
}

func Test_ApplicationSecurityGroupSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationSecurityGroupSpec to ApplicationSecurityGroupSpec via AssignPropertiesToApplicationSecurityGroupSpec & AssignPropertiesFromApplicationSecurityGroupSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationSecurityGroupSpec, ApplicationSecurityGroupSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationSecurityGroupSpec tests if a specific instance of ApplicationSecurityGroupSpec can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForApplicationSecurityGroupSpec(subject ApplicationSecurityGroupSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.ApplicationSecurityGroupSpec
	err := copied.AssignPropertiesToApplicationSecurityGroupSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationSecurityGroupSpec
	err = actual.AssignPropertiesFromApplicationSecurityGroupSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationSecurityGroupSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec, ApplicationSecurityGroupSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec runs a test to see if a specific instance of ApplicationSecurityGroupSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec(subject ApplicationSecurityGroupSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupSpec
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

// Generator of ApplicationSecurityGroupSpec instances for property testing - lazily instantiated by
// ApplicationSecurityGroupSpecGenerator()
var applicationSecurityGroupSpecGenerator gopter.Gen

// ApplicationSecurityGroupSpecGenerator returns a generator of ApplicationSecurityGroupSpec instances for property testing.
func ApplicationSecurityGroupSpecGenerator() gopter.Gen {
	if applicationSecurityGroupSpecGenerator != nil {
		return applicationSecurityGroupSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec(generators)
	applicationSecurityGroupSpecGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupSpec{}), generators)

	return applicationSecurityGroupSpecGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupSpec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded to ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via AssignPropertiesToApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded & AssignPropertiesFromApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded, ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded tests if a specific instance of ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded can be assigned to v1alpha1api20201101storage and back losslessly
func RunPropertyAssignmentTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20201101s.ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err := copied.AssignPropertiesToApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err = actual.AssignPropertiesFromApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded, ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator()
var applicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(generators)
	applicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
