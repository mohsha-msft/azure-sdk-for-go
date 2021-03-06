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

package templatespecs

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2019-06-01-preview/templatespecs"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type Kind = original.Kind

const (
	KindTemplate             Kind = original.KindTemplate
	KindTemplateSpecArtifact Kind = original.KindTemplateSpecArtifact
)

type TemplateSpecExpandKind = original.TemplateSpecExpandKind

const (
	Versions TemplateSpecExpandKind = original.Versions
)

type Artifact = original.Artifact
type AzureResourceBase = original.AzureResourceBase
type BaseClient = original.BaseClient
type BasicArtifact = original.BasicArtifact
type Client = original.Client
type Error = original.Error
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Properties = original.Properties
type SystemData = original.SystemData
type TemplateArtifact = original.TemplateArtifact
type TemplateSpec = original.TemplateSpec
type UpdateModel = original.UpdateModel
type VersionInfo = original.VersionInfo
type VersionProperties = original.VersionProperties
type VersionTemplatespecs = original.VersionTemplatespecs
type VersionUpdateModel = original.VersionUpdateModel
type VersionsClient = original.VersionsClient
type VersionsListResult = original.VersionsListResult
type VersionsListResultIterator = original.VersionsListResultIterator
type VersionsListResultPage = original.VersionsListResultPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewVersionsClient(subscriptionID string) VersionsClient {
	return original.NewVersionsClient(subscriptionID)
}
func NewVersionsClientWithBaseURI(baseURI string, subscriptionID string) VersionsClient {
	return original.NewVersionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVersionsListResultIterator(page VersionsListResultPage) VersionsListResultIterator {
	return original.NewVersionsListResultIterator(page)
}
func NewVersionsListResultPage(cur VersionsListResult, getNextPage func(context.Context, VersionsListResult) (VersionsListResult, error)) VersionsListResultPage {
	return original.NewVersionsListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleTemplateSpecExpandKindValues() []TemplateSpecExpandKind {
	return original.PossibleTemplateSpecExpandKindValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
