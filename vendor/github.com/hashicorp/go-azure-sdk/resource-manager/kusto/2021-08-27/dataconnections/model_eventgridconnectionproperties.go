package dataconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventGridConnectionProperties struct {
	BlobStorageEventType     *BlobStorageEventType `json:"blobStorageEventType,omitempty"`
	ConsumerGroup            string                `json:"consumerGroup"`
	DataFormat               *EventGridDataFormat  `json:"dataFormat,omitempty"`
	EventHubResourceId       string                `json:"eventHubResourceId"`
	IgnoreFirstRecord        *bool                 `json:"ignoreFirstRecord,omitempty"`
	MappingRuleName          *string               `json:"mappingRuleName,omitempty"`
	ProvisioningState        *ProvisioningState    `json:"provisioningState,omitempty"`
	StorageAccountResourceId string                `json:"storageAccountResourceId"`
	TableName                *string               `json:"tableName,omitempty"`
}
