package clusters

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureScaleType string

const (
	AzureScaleTypeAutomatic AzureScaleType = "automatic"
	AzureScaleTypeManual    AzureScaleType = "manual"
	AzureScaleTypeNone      AzureScaleType = "none"
)

func PossibleValuesForAzureScaleType() []string {
	return []string{
		string(AzureScaleTypeAutomatic),
		string(AzureScaleTypeManual),
		string(AzureScaleTypeNone),
	}
}

func parseAzureScaleType(input string) (*AzureScaleType, error) {
	vals := map[string]AzureScaleType{
		"automatic": AzureScaleTypeAutomatic,
		"manual":    AzureScaleTypeManual,
		"none":      AzureScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureScaleType(input)
	return &out, nil
}

type AzureSkuName string

const (
	AzureSkuNameDevNoSLAStandardDOneOneVTwo             AzureSkuName = "Dev(No SLA)_Standard_D11_v2"
	AzureSkuNameDevNoSLAStandardETwoaVFour              AzureSkuName = "Dev(No SLA)_Standard_E2a_v4"
	AzureSkuNameStandardDOneFourVTwo                    AzureSkuName = "Standard_D14_v2"
	AzureSkuNameStandardDOneOneVTwo                     AzureSkuName = "Standard_D11_v2"
	AzureSkuNameStandardDOneThreeVTwo                   AzureSkuName = "Standard_D13_v2"
	AzureSkuNameStandardDOneTwoVTwo                     AzureSkuName = "Standard_D12_v2"
	AzureSkuNameStandardDSOneFourVTwoPositiveFourTBPS   AzureSkuName = "Standard_DS14_v2+4TB_PS"
	AzureSkuNameStandardDSOneFourVTwoPositiveThreeTBPS  AzureSkuName = "Standard_DS14_v2+3TB_PS"
	AzureSkuNameStandardDSOneThreeVTwoPositiveOneTBPS   AzureSkuName = "Standard_DS13_v2+1TB_PS"
	AzureSkuNameStandardDSOneThreeVTwoPositiveTwoTBPS   AzureSkuName = "Standard_DS13_v2+2TB_PS"
	AzureSkuNameStandardEEightZeroidsVFour              AzureSkuName = "Standard_E80ids_v4"
	AzureSkuNameStandardEEightaVFour                    AzureSkuName = "Standard_E8a_v4"
	AzureSkuNameStandardEEightasVFourPositiveOneTBPS    AzureSkuName = "Standard_E8as_v4+1TB_PS"
	AzureSkuNameStandardEEightasVFourPositiveTwoTBPS    AzureSkuName = "Standard_E8as_v4+2TB_PS"
	AzureSkuNameStandardEFouraVFour                     AzureSkuName = "Standard_E4a_v4"
	AzureSkuNameStandardEOneSixaVFour                   AzureSkuName = "Standard_E16a_v4"
	AzureSkuNameStandardEOneSixasVFourPositiveFourTBPS  AzureSkuName = "Standard_E16as_v4+4TB_PS"
	AzureSkuNameStandardEOneSixasVFourPositiveThreeTBPS AzureSkuName = "Standard_E16as_v4+3TB_PS"
	AzureSkuNameStandardESixFouriVThree                 AzureSkuName = "Standard_E64i_v3"
	AzureSkuNameStandardETwoaVFour                      AzureSkuName = "Standard_E2a_v4"
	AzureSkuNameStandardLEights                         AzureSkuName = "Standard_L8s"
	AzureSkuNameStandardLEightsVTwo                     AzureSkuName = "Standard_L8s_v2"
	AzureSkuNameStandardLFours                          AzureSkuName = "Standard_L4s"
	AzureSkuNameStandardLOneSixs                        AzureSkuName = "Standard_L16s"
	AzureSkuNameStandardLOneSixsVTwo                    AzureSkuName = "Standard_L16s_v2"
)

func PossibleValuesForAzureSkuName() []string {
	return []string{
		string(AzureSkuNameDevNoSLAStandardDOneOneVTwo),
		string(AzureSkuNameDevNoSLAStandardETwoaVFour),
		string(AzureSkuNameStandardDOneFourVTwo),
		string(AzureSkuNameStandardDOneOneVTwo),
		string(AzureSkuNameStandardDOneThreeVTwo),
		string(AzureSkuNameStandardDOneTwoVTwo),
		string(AzureSkuNameStandardDSOneFourVTwoPositiveFourTBPS),
		string(AzureSkuNameStandardDSOneFourVTwoPositiveThreeTBPS),
		string(AzureSkuNameStandardDSOneThreeVTwoPositiveOneTBPS),
		string(AzureSkuNameStandardDSOneThreeVTwoPositiveTwoTBPS),
		string(AzureSkuNameStandardEEightZeroidsVFour),
		string(AzureSkuNameStandardEEightaVFour),
		string(AzureSkuNameStandardEEightasVFourPositiveOneTBPS),
		string(AzureSkuNameStandardEEightasVFourPositiveTwoTBPS),
		string(AzureSkuNameStandardEFouraVFour),
		string(AzureSkuNameStandardEOneSixaVFour),
		string(AzureSkuNameStandardEOneSixasVFourPositiveFourTBPS),
		string(AzureSkuNameStandardEOneSixasVFourPositiveThreeTBPS),
		string(AzureSkuNameStandardESixFouriVThree),
		string(AzureSkuNameStandardETwoaVFour),
		string(AzureSkuNameStandardLEights),
		string(AzureSkuNameStandardLEightsVTwo),
		string(AzureSkuNameStandardLFours),
		string(AzureSkuNameStandardLOneSixs),
		string(AzureSkuNameStandardLOneSixsVTwo),
	}
}

func parseAzureSkuName(input string) (*AzureSkuName, error) {
	vals := map[string]AzureSkuName{
		"dev(no sla)_standard_d11_v2": AzureSkuNameDevNoSLAStandardDOneOneVTwo,
		"dev(no sla)_standard_e2a_v4": AzureSkuNameDevNoSLAStandardETwoaVFour,
		"standard_d14_v2":             AzureSkuNameStandardDOneFourVTwo,
		"standard_d11_v2":             AzureSkuNameStandardDOneOneVTwo,
		"standard_d13_v2":             AzureSkuNameStandardDOneThreeVTwo,
		"standard_d12_v2":             AzureSkuNameStandardDOneTwoVTwo,
		"standard_ds14_v2+4tb_ps":     AzureSkuNameStandardDSOneFourVTwoPositiveFourTBPS,
		"standard_ds14_v2+3tb_ps":     AzureSkuNameStandardDSOneFourVTwoPositiveThreeTBPS,
		"standard_ds13_v2+1tb_ps":     AzureSkuNameStandardDSOneThreeVTwoPositiveOneTBPS,
		"standard_ds13_v2+2tb_ps":     AzureSkuNameStandardDSOneThreeVTwoPositiveTwoTBPS,
		"standard_e80ids_v4":          AzureSkuNameStandardEEightZeroidsVFour,
		"standard_e8a_v4":             AzureSkuNameStandardEEightaVFour,
		"standard_e8as_v4+1tb_ps":     AzureSkuNameStandardEEightasVFourPositiveOneTBPS,
		"standard_e8as_v4+2tb_ps":     AzureSkuNameStandardEEightasVFourPositiveTwoTBPS,
		"standard_e4a_v4":             AzureSkuNameStandardEFouraVFour,
		"standard_e16a_v4":            AzureSkuNameStandardEOneSixaVFour,
		"standard_e16as_v4+4tb_ps":    AzureSkuNameStandardEOneSixasVFourPositiveFourTBPS,
		"standard_e16as_v4+3tb_ps":    AzureSkuNameStandardEOneSixasVFourPositiveThreeTBPS,
		"standard_e64i_v3":            AzureSkuNameStandardESixFouriVThree,
		"standard_e2a_v4":             AzureSkuNameStandardETwoaVFour,
		"standard_l8s":                AzureSkuNameStandardLEights,
		"standard_l8s_v2":             AzureSkuNameStandardLEightsVTwo,
		"standard_l4s":                AzureSkuNameStandardLFours,
		"standard_l16s":               AzureSkuNameStandardLOneSixs,
		"standard_l16s_v2":            AzureSkuNameStandardLOneSixsVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureSkuName(input)
	return &out, nil
}

type AzureSkuTier string

const (
	AzureSkuTierBasic    AzureSkuTier = "Basic"
	AzureSkuTierStandard AzureSkuTier = "Standard"
)

func PossibleValuesForAzureSkuTier() []string {
	return []string{
		string(AzureSkuTierBasic),
		string(AzureSkuTierStandard),
	}
}

func parseAzureSkuTier(input string) (*AzureSkuTier, error) {
	vals := map[string]AzureSkuTier{
		"basic":    AzureSkuTierBasic,
		"standard": AzureSkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureSkuTier(input)
	return &out, nil
}

type ClusterNetworkAccessFlag string

const (
	ClusterNetworkAccessFlagDisabled ClusterNetworkAccessFlag = "Disabled"
	ClusterNetworkAccessFlagEnabled  ClusterNetworkAccessFlag = "Enabled"
)

func PossibleValuesForClusterNetworkAccessFlag() []string {
	return []string{
		string(ClusterNetworkAccessFlagDisabled),
		string(ClusterNetworkAccessFlagEnabled),
	}
}

func parseClusterNetworkAccessFlag(input string) (*ClusterNetworkAccessFlag, error) {
	vals := map[string]ClusterNetworkAccessFlag{
		"disabled": ClusterNetworkAccessFlagDisabled,
		"enabled":  ClusterNetworkAccessFlagEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterNetworkAccessFlag(input)
	return &out, nil
}

type ClusterType string

const (
	ClusterTypeMicrosoftPointKustoClusters ClusterType = "Microsoft.Kusto/clusters"
)

func PossibleValuesForClusterType() []string {
	return []string{
		string(ClusterTypeMicrosoftPointKustoClusters),
	}
}

func parseClusterType(input string) (*ClusterType, error) {
	vals := map[string]ClusterType{
		"microsoft.kusto/clusters": ClusterTypeMicrosoftPointKustoClusters,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterType(input)
	return &out, nil
}

type EngineType string

const (
	EngineTypeVThree EngineType = "V3"
	EngineTypeVTwo   EngineType = "V2"
)

func PossibleValuesForEngineType() []string {
	return []string{
		string(EngineTypeVThree),
		string(EngineTypeVTwo),
	}
}

func parseEngineType(input string) (*EngineType, error) {
	vals := map[string]EngineType{
		"v3": EngineTypeVThree,
		"v2": EngineTypeVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EngineType(input)
	return &out, nil
}

type LanguageExtensionName string

const (
	LanguageExtensionNamePYTHON LanguageExtensionName = "PYTHON"
	LanguageExtensionNameR      LanguageExtensionName = "R"
)

func PossibleValuesForLanguageExtensionName() []string {
	return []string{
		string(LanguageExtensionNamePYTHON),
		string(LanguageExtensionNameR),
	}
}

func parseLanguageExtensionName(input string) (*LanguageExtensionName, error) {
	vals := map[string]LanguageExtensionName{
		"python": LanguageExtensionNamePYTHON,
		"r":      LanguageExtensionNameR,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageExtensionName(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"moving":    ProvisioningStateMoving,
		"running":   ProvisioningStateRunning,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

func parsePublicNetworkAccess(input string) (*PublicNetworkAccess, error) {
	vals := map[string]PublicNetworkAccess{
		"disabled": PublicNetworkAccessDisabled,
		"enabled":  PublicNetworkAccessEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicNetworkAccess(input)
	return &out, nil
}

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
		string(ReasonInvalid),
	}
}

func parseReason(input string) (*Reason, error) {
	vals := map[string]Reason{
		"alreadyexists": ReasonAlreadyExists,
		"invalid":       ReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Reason(input)
	return &out, nil
}

type State string

const (
	StateCreating    State = "Creating"
	StateDeleted     State = "Deleted"
	StateDeleting    State = "Deleting"
	StateRunning     State = "Running"
	StateStarting    State = "Starting"
	StateStopped     State = "Stopped"
	StateStopping    State = "Stopping"
	StateUnavailable State = "Unavailable"
	StateUpdating    State = "Updating"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateCreating),
		string(StateDeleted),
		string(StateDeleting),
		string(StateRunning),
		string(StateStarting),
		string(StateStopped),
		string(StateStopping),
		string(StateUnavailable),
		string(StateUpdating),
	}
}

func parseState(input string) (*State, error) {
	vals := map[string]State{
		"creating":    StateCreating,
		"deleted":     StateDeleted,
		"deleting":    StateDeleting,
		"running":     StateRunning,
		"starting":    StateStarting,
		"stopped":     StateStopped,
		"stopping":    StateStopping,
		"unavailable": StateUnavailable,
		"updating":    StateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := State(input)
	return &out, nil
}
