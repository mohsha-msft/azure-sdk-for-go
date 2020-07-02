// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package insights

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ApplicationType = original.ApplicationType

const (
	Other ApplicationType = original.Other
	Web   ApplicationType = original.Web
)

type CategoryType = original.CategoryType

const (
	CategoryTypePerformance CategoryType = original.CategoryTypePerformance
	CategoryTypeRetention   CategoryType = original.CategoryTypeRetention
	CategoryTypeTSG         CategoryType = original.CategoryTypeTSG
	CategoryTypeWorkbook    CategoryType = original.CategoryTypeWorkbook
)

type FavoriteSourceType = original.FavoriteSourceType

const (
	Events       FavoriteSourceType = original.Events
	Funnel       FavoriteSourceType = original.Funnel
	Impact       FavoriteSourceType = original.Impact
	Notebook     FavoriteSourceType = original.Notebook
	Retention    FavoriteSourceType = original.Retention
	Segmentation FavoriteSourceType = original.Segmentation
	Sessions     FavoriteSourceType = original.Sessions
	Userflows    FavoriteSourceType = original.Userflows
)

type FavoriteType = original.FavoriteType

const (
	Shared FavoriteType = original.Shared
	User   FavoriteType = original.User
)

type FlowType = original.FlowType

const (
	Bluefield FlowType = original.Bluefield
)

type IngestionMode = original.IngestionMode

const (
	ApplicationInsights                       IngestionMode = original.ApplicationInsights
	ApplicationInsightsWithDiagnosticSettings IngestionMode = original.ApplicationInsightsWithDiagnosticSettings
	LogAnalytics                              IngestionMode = original.LogAnalytics
)

type ItemScope = original.ItemScope

const (
	ItemScopeShared ItemScope = original.ItemScopeShared
	ItemScopeUser   ItemScope = original.ItemScopeUser
)

type ItemScopePath = original.ItemScopePath

const (
	AnalyticsItems   ItemScopePath = original.AnalyticsItems
	MyanalyticsItems ItemScopePath = original.MyanalyticsItems
)

type ItemType = original.ItemType

const (
	Folder   ItemType = original.Folder
	Function ItemType = original.Function
	Query    ItemType = original.Query
	Recent   ItemType = original.Recent
)

type ItemTypeParameter = original.ItemTypeParameter

const (
	ItemTypeParameterFolder   ItemTypeParameter = original.ItemTypeParameterFolder
	ItemTypeParameterFunction ItemTypeParameter = original.ItemTypeParameterFunction
	ItemTypeParameterNone     ItemTypeParameter = original.ItemTypeParameterNone
	ItemTypeParameterQuery    ItemTypeParameter = original.ItemTypeParameterQuery
	ItemTypeParameterRecent   ItemTypeParameter = original.ItemTypeParameterRecent
)

type PurgeState = original.PurgeState

const (
	Completed PurgeState = original.Completed
	Pending   PurgeState = original.Pending
)

type RequestSource = original.RequestSource

const (
	Rest RequestSource = original.Rest
)

type SharedTypeKind = original.SharedTypeKind

const (
	SharedTypeKindShared SharedTypeKind = original.SharedTypeKindShared
	SharedTypeKindUser   SharedTypeKind = original.SharedTypeKindUser
)

type WebTestKind = original.WebTestKind

const (
	Multistep WebTestKind = original.Multistep
	Ping      WebTestKind = original.Ping
)

type APIKeyRequest = original.APIKeyRequest
type APIKeysClient = original.APIKeysClient
type AnalyticsItemsClient = original.AnalyticsItemsClient
type Annotation = original.Annotation
type AnnotationError = original.AnnotationError
type AnnotationsClient = original.AnnotationsClient
type AnnotationsListResult = original.AnnotationsListResult
type ApplicationInsightsComponent = original.ApplicationInsightsComponent
type ApplicationInsightsComponentAPIKey = original.ApplicationInsightsComponentAPIKey
type ApplicationInsightsComponentAPIKeyListResult = original.ApplicationInsightsComponentAPIKeyListResult
type ApplicationInsightsComponentAnalyticsItem = original.ApplicationInsightsComponentAnalyticsItem
type ApplicationInsightsComponentAnalyticsItemProperties = original.ApplicationInsightsComponentAnalyticsItemProperties
type ApplicationInsightsComponentAvailableFeatures = original.ApplicationInsightsComponentAvailableFeatures
type ApplicationInsightsComponentBillingFeatures = original.ApplicationInsightsComponentBillingFeatures
type ApplicationInsightsComponentDataVolumeCap = original.ApplicationInsightsComponentDataVolumeCap
type ApplicationInsightsComponentExportConfiguration = original.ApplicationInsightsComponentExportConfiguration
type ApplicationInsightsComponentExportRequest = original.ApplicationInsightsComponentExportRequest
type ApplicationInsightsComponentFavorite = original.ApplicationInsightsComponentFavorite
type ApplicationInsightsComponentFeature = original.ApplicationInsightsComponentFeature
type ApplicationInsightsComponentFeatureCapabilities = original.ApplicationInsightsComponentFeatureCapabilities
type ApplicationInsightsComponentFeatureCapability = original.ApplicationInsightsComponentFeatureCapability
type ApplicationInsightsComponentListResult = original.ApplicationInsightsComponentListResult
type ApplicationInsightsComponentListResultIterator = original.ApplicationInsightsComponentListResultIterator
type ApplicationInsightsComponentListResultPage = original.ApplicationInsightsComponentListResultPage
type ApplicationInsightsComponentProactiveDetectionConfiguration = original.ApplicationInsightsComponentProactiveDetectionConfiguration
type ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions = original.ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions
type ApplicationInsightsComponentProperties = original.ApplicationInsightsComponentProperties
type ApplicationInsightsComponentQuotaStatus = original.ApplicationInsightsComponentQuotaStatus
type ApplicationInsightsComponentWebTestLocation = original.ApplicationInsightsComponentWebTestLocation
type ApplicationInsightsWebTestLocationsListResult = original.ApplicationInsightsWebTestLocationsListResult
type BaseClient = original.BaseClient
type ComponentAvailableFeaturesClient = original.ComponentAvailableFeaturesClient
type ComponentCurrentBillingFeaturesClient = original.ComponentCurrentBillingFeaturesClient
type ComponentFeatureCapabilitiesClient = original.ComponentFeatureCapabilitiesClient
type ComponentPurgeBody = original.ComponentPurgeBody
type ComponentPurgeBodyFilters = original.ComponentPurgeBodyFilters
type ComponentPurgeResponse = original.ComponentPurgeResponse
type ComponentPurgeStatusResponse = original.ComponentPurgeStatusResponse
type ComponentQuotaStatusClient = original.ComponentQuotaStatusClient
type ComponentsClient = original.ComponentsClient
type ComponentsResource = original.ComponentsResource
type ErrorFieldContract = original.ErrorFieldContract
type ErrorResponse = original.ErrorResponse
type ExportConfigurationsClient = original.ExportConfigurationsClient
type FavoritesClient = original.FavoritesClient
type InnerError = original.InnerError
type LinkProperties = original.LinkProperties
type ListAnnotation = original.ListAnnotation
type ListApplicationInsightsComponentAnalyticsItem = original.ListApplicationInsightsComponentAnalyticsItem
type ListApplicationInsightsComponentExportConfiguration = original.ListApplicationInsightsComponentExportConfiguration
type ListApplicationInsightsComponentFavorite = original.ListApplicationInsightsComponentFavorite
type ListApplicationInsightsComponentProactiveDetectionConfiguration = original.ListApplicationInsightsComponentProactiveDetectionConfiguration
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PrivateLinkScopedResource = original.PrivateLinkScopedResource
type ProactiveDetectionConfigurationsClient = original.ProactiveDetectionConfigurationsClient
type TagsResource = original.TagsResource
type WebTest = original.WebTest
type WebTestGeolocation = original.WebTestGeolocation
type WebTestListResult = original.WebTestListResult
type WebTestListResultIterator = original.WebTestListResultIterator
type WebTestListResultPage = original.WebTestListResultPage
type WebTestLocationsClient = original.WebTestLocationsClient
type WebTestProperties = original.WebTestProperties
type WebTestPropertiesConfiguration = original.WebTestPropertiesConfiguration
type WebTestsClient = original.WebTestsClient
type WebtestsResource = original.WebtestsResource
type WorkItemConfiguration = original.WorkItemConfiguration
type WorkItemConfigurationError = original.WorkItemConfigurationError
type WorkItemConfigurationsClient = original.WorkItemConfigurationsClient
type WorkItemConfigurationsListResult = original.WorkItemConfigurationsListResult
type WorkItemCreateConfiguration = original.WorkItemCreateConfiguration
type Workbook = original.Workbook
type WorkbookError = original.WorkbookError
type WorkbookProperties = original.WorkbookProperties
type WorkbookResource = original.WorkbookResource
type WorkbooksClient = original.WorkbooksClient
type WorkbooksListResult = original.WorkbooksListResult

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAPIKeysClient(subscriptionID string) APIKeysClient {
	return original.NewAPIKeysClient(subscriptionID)
}
func NewAPIKeysClientWithBaseURI(baseURI string, subscriptionID string) APIKeysClient {
	return original.NewAPIKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewAnalyticsItemsClient(subscriptionID string) AnalyticsItemsClient {
	return original.NewAnalyticsItemsClient(subscriptionID)
}
func NewAnalyticsItemsClientWithBaseURI(baseURI string, subscriptionID string) AnalyticsItemsClient {
	return original.NewAnalyticsItemsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAnnotationsClient(subscriptionID string) AnnotationsClient {
	return original.NewAnnotationsClient(subscriptionID)
}
func NewAnnotationsClientWithBaseURI(baseURI string, subscriptionID string) AnnotationsClient {
	return original.NewAnnotationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationInsightsComponentListResultIterator(page ApplicationInsightsComponentListResultPage) ApplicationInsightsComponentListResultIterator {
	return original.NewApplicationInsightsComponentListResultIterator(page)
}
func NewApplicationInsightsComponentListResultPage(getNextPage func(context.Context, ApplicationInsightsComponentListResult) (ApplicationInsightsComponentListResult, error)) ApplicationInsightsComponentListResultPage {
	return original.NewApplicationInsightsComponentListResultPage(getNextPage)
}
func NewComponentAvailableFeaturesClient(subscriptionID string) ComponentAvailableFeaturesClient {
	return original.NewComponentAvailableFeaturesClient(subscriptionID)
}
func NewComponentAvailableFeaturesClientWithBaseURI(baseURI string, subscriptionID string) ComponentAvailableFeaturesClient {
	return original.NewComponentAvailableFeaturesClientWithBaseURI(baseURI, subscriptionID)
}
func NewComponentCurrentBillingFeaturesClient(subscriptionID string) ComponentCurrentBillingFeaturesClient {
	return original.NewComponentCurrentBillingFeaturesClient(subscriptionID)
}
func NewComponentCurrentBillingFeaturesClientWithBaseURI(baseURI string, subscriptionID string) ComponentCurrentBillingFeaturesClient {
	return original.NewComponentCurrentBillingFeaturesClientWithBaseURI(baseURI, subscriptionID)
}
func NewComponentFeatureCapabilitiesClient(subscriptionID string) ComponentFeatureCapabilitiesClient {
	return original.NewComponentFeatureCapabilitiesClient(subscriptionID)
}
func NewComponentFeatureCapabilitiesClientWithBaseURI(baseURI string, subscriptionID string) ComponentFeatureCapabilitiesClient {
	return original.NewComponentFeatureCapabilitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewComponentQuotaStatusClient(subscriptionID string) ComponentQuotaStatusClient {
	return original.NewComponentQuotaStatusClient(subscriptionID)
}
func NewComponentQuotaStatusClientWithBaseURI(baseURI string, subscriptionID string) ComponentQuotaStatusClient {
	return original.NewComponentQuotaStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewComponentsClient(subscriptionID string) ComponentsClient {
	return original.NewComponentsClient(subscriptionID)
}
func NewComponentsClientWithBaseURI(baseURI string, subscriptionID string) ComponentsClient {
	return original.NewComponentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExportConfigurationsClient(subscriptionID string) ExportConfigurationsClient {
	return original.NewExportConfigurationsClient(subscriptionID)
}
func NewExportConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ExportConfigurationsClient {
	return original.NewExportConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFavoritesClient(subscriptionID string) FavoritesClient {
	return original.NewFavoritesClient(subscriptionID)
}
func NewFavoritesClientWithBaseURI(baseURI string, subscriptionID string) FavoritesClient {
	return original.NewFavoritesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProactiveDetectionConfigurationsClient(subscriptionID string) ProactiveDetectionConfigurationsClient {
	return original.NewProactiveDetectionConfigurationsClient(subscriptionID)
}
func NewProactiveDetectionConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ProactiveDetectionConfigurationsClient {
	return original.NewProactiveDetectionConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebTestListResultIterator(page WebTestListResultPage) WebTestListResultIterator {
	return original.NewWebTestListResultIterator(page)
}
func NewWebTestListResultPage(getNextPage func(context.Context, WebTestListResult) (WebTestListResult, error)) WebTestListResultPage {
	return original.NewWebTestListResultPage(getNextPage)
}
func NewWebTestLocationsClient(subscriptionID string) WebTestLocationsClient {
	return original.NewWebTestLocationsClient(subscriptionID)
}
func NewWebTestLocationsClientWithBaseURI(baseURI string, subscriptionID string) WebTestLocationsClient {
	return original.NewWebTestLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebTestsClient(subscriptionID string) WebTestsClient {
	return original.NewWebTestsClient(subscriptionID)
}
func NewWebTestsClientWithBaseURI(baseURI string, subscriptionID string) WebTestsClient {
	return original.NewWebTestsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkItemConfigurationsClient(subscriptionID string) WorkItemConfigurationsClient {
	return original.NewWorkItemConfigurationsClient(subscriptionID)
}
func NewWorkItemConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) WorkItemConfigurationsClient {
	return original.NewWorkItemConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkbooksClient(subscriptionID string) WorkbooksClient {
	return original.NewWorkbooksClient(subscriptionID)
}
func NewWorkbooksClientWithBaseURI(baseURI string, subscriptionID string) WorkbooksClient {
	return original.NewWorkbooksClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleApplicationTypeValues() []ApplicationType {
	return original.PossibleApplicationTypeValues()
}
func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}
func PossibleFavoriteSourceTypeValues() []FavoriteSourceType {
	return original.PossibleFavoriteSourceTypeValues()
}
func PossibleFavoriteTypeValues() []FavoriteType {
	return original.PossibleFavoriteTypeValues()
}
func PossibleFlowTypeValues() []FlowType {
	return original.PossibleFlowTypeValues()
}
func PossibleIngestionModeValues() []IngestionMode {
	return original.PossibleIngestionModeValues()
}
func PossibleItemScopePathValues() []ItemScopePath {
	return original.PossibleItemScopePathValues()
}
func PossibleItemScopeValues() []ItemScope {
	return original.PossibleItemScopeValues()
}
func PossibleItemTypeParameterValues() []ItemTypeParameter {
	return original.PossibleItemTypeParameterValues()
}
func PossibleItemTypeValues() []ItemType {
	return original.PossibleItemTypeValues()
}
func PossiblePurgeStateValues() []PurgeState {
	return original.PossiblePurgeStateValues()
}
func PossibleRequestSourceValues() []RequestSource {
	return original.PossibleRequestSourceValues()
}
func PossibleSharedTypeKindValues() []SharedTypeKind {
	return original.PossibleSharedTypeKindValues()
}
func PossibleWebTestKindValues() []WebTestKind {
	return original.PossibleWebTestKindValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
