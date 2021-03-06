package devices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// DeviceConnect ...
	DeviceConnect AccessRights = "DeviceConnect"
	// RegistryRead ...
	RegistryRead AccessRights = "RegistryRead"
	// RegistryReadDeviceConnect ...
	RegistryReadDeviceConnect AccessRights = "RegistryRead, DeviceConnect"
	// RegistryReadRegistryWrite ...
	RegistryReadRegistryWrite AccessRights = "RegistryRead, RegistryWrite"
	// RegistryReadRegistryWriteDeviceConnect ...
	RegistryReadRegistryWriteDeviceConnect AccessRights = "RegistryRead, RegistryWrite, DeviceConnect"
	// RegistryReadRegistryWriteServiceConnect ...
	RegistryReadRegistryWriteServiceConnect AccessRights = "RegistryRead, RegistryWrite, ServiceConnect"
	// RegistryReadRegistryWriteServiceConnectDeviceConnect ...
	RegistryReadRegistryWriteServiceConnectDeviceConnect AccessRights = "RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect"
	// RegistryReadServiceConnect ...
	RegistryReadServiceConnect AccessRights = "RegistryRead, ServiceConnect"
	// RegistryReadServiceConnectDeviceConnect ...
	RegistryReadServiceConnectDeviceConnect AccessRights = "RegistryRead, ServiceConnect, DeviceConnect"
	// RegistryWrite ...
	RegistryWrite AccessRights = "RegistryWrite"
	// RegistryWriteDeviceConnect ...
	RegistryWriteDeviceConnect AccessRights = "RegistryWrite, DeviceConnect"
	// RegistryWriteServiceConnect ...
	RegistryWriteServiceConnect AccessRights = "RegistryWrite, ServiceConnect"
	// RegistryWriteServiceConnectDeviceConnect ...
	RegistryWriteServiceConnectDeviceConnect AccessRights = "RegistryWrite, ServiceConnect, DeviceConnect"
	// ServiceConnect ...
	ServiceConnect AccessRights = "ServiceConnect"
	// ServiceConnectDeviceConnect ...
	ServiceConnectDeviceConnect AccessRights = "ServiceConnect, DeviceConnect"
)

// PossibleAccessRightsValues returns an array of possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{DeviceConnect, RegistryRead, RegistryReadDeviceConnect, RegistryReadRegistryWrite, RegistryReadRegistryWriteDeviceConnect, RegistryReadRegistryWriteServiceConnect, RegistryReadRegistryWriteServiceConnectDeviceConnect, RegistryReadServiceConnect, RegistryReadServiceConnectDeviceConnect, RegistryWrite, RegistryWriteDeviceConnect, RegistryWriteServiceConnect, RegistryWriteServiceConnectDeviceConnect, ServiceConnect, ServiceConnectDeviceConnect}
}

// Capabilities enumerates the values for capabilities.
type Capabilities string

const (
	// DeviceManagement ...
	DeviceManagement Capabilities = "DeviceManagement"
	// None ...
	None Capabilities = "None"
)

// PossibleCapabilitiesValues returns an array of possible values for the Capabilities const type.
func PossibleCapabilitiesValues() []Capabilities {
	return []Capabilities{DeviceManagement, None}
}

// EndpointHealthStatus enumerates the values for endpoint health status.
type EndpointHealthStatus string

const (
	// Dead ...
	Dead EndpointHealthStatus = "dead"
	// Healthy ...
	Healthy EndpointHealthStatus = "healthy"
	// Unhealthy ...
	Unhealthy EndpointHealthStatus = "unhealthy"
	// Unknown ...
	Unknown EndpointHealthStatus = "unknown"
)

// PossibleEndpointHealthStatusValues returns an array of possible values for the EndpointHealthStatus const type.
func PossibleEndpointHealthStatusValues() []EndpointHealthStatus {
	return []EndpointHealthStatus{Dead, Healthy, Unhealthy, Unknown}
}

// IotHubNameUnavailabilityReason enumerates the values for iot hub name unavailability reason.
type IotHubNameUnavailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists IotHubNameUnavailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid IotHubNameUnavailabilityReason = "Invalid"
)

// PossibleIotHubNameUnavailabilityReasonValues returns an array of possible values for the IotHubNameUnavailabilityReason const type.
func PossibleIotHubNameUnavailabilityReasonValues() []IotHubNameUnavailabilityReason {
	return []IotHubNameUnavailabilityReason{AlreadyExists, Invalid}
}

// IotHubScaleType enumerates the values for iot hub scale type.
type IotHubScaleType string

const (
	// IotHubScaleTypeAutomatic ...
	IotHubScaleTypeAutomatic IotHubScaleType = "Automatic"
	// IotHubScaleTypeManual ...
	IotHubScaleTypeManual IotHubScaleType = "Manual"
	// IotHubScaleTypeNone ...
	IotHubScaleTypeNone IotHubScaleType = "None"
)

// PossibleIotHubScaleTypeValues returns an array of possible values for the IotHubScaleType const type.
func PossibleIotHubScaleTypeValues() []IotHubScaleType {
	return []IotHubScaleType{IotHubScaleTypeAutomatic, IotHubScaleTypeManual, IotHubScaleTypeNone}
}

// IotHubSku enumerates the values for iot hub sku.
type IotHubSku string

const (
	// B1 ...
	B1 IotHubSku = "B1"
	// B2 ...
	B2 IotHubSku = "B2"
	// B3 ...
	B3 IotHubSku = "B3"
	// F1 ...
	F1 IotHubSku = "F1"
	// S1 ...
	S1 IotHubSku = "S1"
	// S2 ...
	S2 IotHubSku = "S2"
	// S3 ...
	S3 IotHubSku = "S3"
)

// PossibleIotHubSkuValues returns an array of possible values for the IotHubSku const type.
func PossibleIotHubSkuValues() []IotHubSku {
	return []IotHubSku{B1, B2, B3, F1, S1, S2, S3}
}

// IotHubSkuTier enumerates the values for iot hub sku tier.
type IotHubSkuTier string

const (
	// Basic ...
	Basic IotHubSkuTier = "Basic"
	// Free ...
	Free IotHubSkuTier = "Free"
	// Standard ...
	Standard IotHubSkuTier = "Standard"
)

// PossibleIotHubSkuTierValues returns an array of possible values for the IotHubSkuTier const type.
func PossibleIotHubSkuTierValues() []IotHubSkuTier {
	return []IotHubSkuTier{Basic, Free, Standard}
}

// IPFilterActionType enumerates the values for ip filter action type.
type IPFilterActionType string

const (
	// Accept ...
	Accept IPFilterActionType = "Accept"
	// Reject ...
	Reject IPFilterActionType = "Reject"
)

// PossibleIPFilterActionTypeValues returns an array of possible values for the IPFilterActionType const type.
func PossibleIPFilterActionTypeValues() []IPFilterActionType {
	return []IPFilterActionType{Accept, Reject}
}

// JobStatus enumerates the values for job status.
type JobStatus string

const (
	// JobStatusCancelled ...
	JobStatusCancelled JobStatus = "cancelled"
	// JobStatusCompleted ...
	JobStatusCompleted JobStatus = "completed"
	// JobStatusEnqueued ...
	JobStatusEnqueued JobStatus = "enqueued"
	// JobStatusFailed ...
	JobStatusFailed JobStatus = "failed"
	// JobStatusRunning ...
	JobStatusRunning JobStatus = "running"
	// JobStatusUnknown ...
	JobStatusUnknown JobStatus = "unknown"
)

// PossibleJobStatusValues returns an array of possible values for the JobStatus const type.
func PossibleJobStatusValues() []JobStatus {
	return []JobStatus{JobStatusCancelled, JobStatusCompleted, JobStatusEnqueued, JobStatusFailed, JobStatusRunning, JobStatusUnknown}
}

// JobType enumerates the values for job type.
type JobType string

const (
	// JobTypeBackup ...
	JobTypeBackup JobType = "backup"
	// JobTypeExport ...
	JobTypeExport JobType = "export"
	// JobTypeFactoryResetDevice ...
	JobTypeFactoryResetDevice JobType = "factoryResetDevice"
	// JobTypeFirmwareUpdate ...
	JobTypeFirmwareUpdate JobType = "firmwareUpdate"
	// JobTypeImport ...
	JobTypeImport JobType = "import"
	// JobTypeReadDeviceProperties ...
	JobTypeReadDeviceProperties JobType = "readDeviceProperties"
	// JobTypeRebootDevice ...
	JobTypeRebootDevice JobType = "rebootDevice"
	// JobTypeUnknown ...
	JobTypeUnknown JobType = "unknown"
	// JobTypeUpdateDeviceConfiguration ...
	JobTypeUpdateDeviceConfiguration JobType = "updateDeviceConfiguration"
	// JobTypeWriteDeviceProperties ...
	JobTypeWriteDeviceProperties JobType = "writeDeviceProperties"
)

// PossibleJobTypeValues returns an array of possible values for the JobType const type.
func PossibleJobTypeValues() []JobType {
	return []JobType{JobTypeBackup, JobTypeExport, JobTypeFactoryResetDevice, JobTypeFirmwareUpdate, JobTypeImport, JobTypeReadDeviceProperties, JobTypeRebootDevice, JobTypeUnknown, JobTypeUpdateDeviceConfiguration, JobTypeWriteDeviceProperties}
}

// OperationMonitoringLevel enumerates the values for operation monitoring level.
type OperationMonitoringLevel string

const (
	// OperationMonitoringLevelError ...
	OperationMonitoringLevelError OperationMonitoringLevel = "Error"
	// OperationMonitoringLevelErrorInformation ...
	OperationMonitoringLevelErrorInformation OperationMonitoringLevel = "Error, Information"
	// OperationMonitoringLevelInformation ...
	OperationMonitoringLevelInformation OperationMonitoringLevel = "Information"
	// OperationMonitoringLevelNone ...
	OperationMonitoringLevelNone OperationMonitoringLevel = "None"
)

// PossibleOperationMonitoringLevelValues returns an array of possible values for the OperationMonitoringLevel const type.
func PossibleOperationMonitoringLevelValues() []OperationMonitoringLevel {
	return []OperationMonitoringLevel{OperationMonitoringLevelError, OperationMonitoringLevelErrorInformation, OperationMonitoringLevelInformation, OperationMonitoringLevelNone}
}

// RouteErrorSeverity enumerates the values for route error severity.
type RouteErrorSeverity string

const (
	// Error ...
	Error RouteErrorSeverity = "error"
	// Warning ...
	Warning RouteErrorSeverity = "warning"
)

// PossibleRouteErrorSeverityValues returns an array of possible values for the RouteErrorSeverity const type.
func PossibleRouteErrorSeverityValues() []RouteErrorSeverity {
	return []RouteErrorSeverity{Error, Warning}
}

// RoutingSource enumerates the values for routing source.
type RoutingSource string

const (
	// RoutingSourceDeviceJobLifecycleEvents ...
	RoutingSourceDeviceJobLifecycleEvents RoutingSource = "DeviceJobLifecycleEvents"
	// RoutingSourceDeviceLifecycleEvents ...
	RoutingSourceDeviceLifecycleEvents RoutingSource = "DeviceLifecycleEvents"
	// RoutingSourceDeviceMessages ...
	RoutingSourceDeviceMessages RoutingSource = "DeviceMessages"
	// RoutingSourceInvalid ...
	RoutingSourceInvalid RoutingSource = "Invalid"
	// RoutingSourceTwinChangeEvents ...
	RoutingSourceTwinChangeEvents RoutingSource = "TwinChangeEvents"
)

// PossibleRoutingSourceValues returns an array of possible values for the RoutingSource const type.
func PossibleRoutingSourceValues() []RoutingSource {
	return []RoutingSource{RoutingSourceDeviceJobLifecycleEvents, RoutingSourceDeviceLifecycleEvents, RoutingSourceDeviceMessages, RoutingSourceInvalid, RoutingSourceTwinChangeEvents}
}

// TestResultStatus enumerates the values for test result status.
type TestResultStatus string

const (
	// False ...
	False TestResultStatus = "false"
	// True ...
	True TestResultStatus = "true"
	// Undefined ...
	Undefined TestResultStatus = "undefined"
)

// PossibleTestResultStatusValues returns an array of possible values for the TestResultStatus const type.
func PossibleTestResultStatusValues() []TestResultStatus {
	return []TestResultStatus{False, True, Undefined}
}
