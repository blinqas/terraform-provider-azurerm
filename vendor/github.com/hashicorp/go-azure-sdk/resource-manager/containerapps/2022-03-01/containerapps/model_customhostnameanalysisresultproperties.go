package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomHostnameAnalysisResultProperties struct {
	ARecords                            *[]string                  `json:"aRecords,omitempty"`
	AlternateCNameRecords               *[]string                  `json:"alternateCNameRecords,omitempty"`
	AlternateTxtRecords                 *[]string                  `json:"alternateTxtRecords,omitempty"`
	CNameRecords                        *[]string                  `json:"cNameRecords,omitempty"`
	ConflictingContainerAppResourceId   *string                    `json:"conflictingContainerAppResourceId,omitempty"`
	CustomDomainVerificationFailureInfo *DefaultErrorResponse      `json:"customDomainVerificationFailureInfo,omitempty"`
	CustomDomainVerificationTest        *DnsVerificationTestResult `json:"customDomainVerificationTest,omitempty"`
	HasConflictOnManagedEnvironment     *bool                      `json:"hasConflictOnManagedEnvironment,omitempty"`
	HostName                            *string                    `json:"hostName,omitempty"`
	IsHostnameAlreadyVerified           *bool                      `json:"isHostnameAlreadyVerified,omitempty"`
	TxtRecords                          *[]string                  `json:"txtRecords,omitempty"`
}
