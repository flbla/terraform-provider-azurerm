package logicalnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogicalNetworkStatusProvisioningStatus struct {
	OperationId *string `json:"operationId,omitempty"`
	Status      *Status `json:"status,omitempty"`
}
