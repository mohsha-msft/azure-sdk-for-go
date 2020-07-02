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

package digitaltwins

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/digitaltwins/2020-05-31/digitaltwins"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type Client = original.Client
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type EventRoute = original.EventRoute
type EventRouteCollection = original.EventRouteCollection
type EventRouteCollectionIterator = original.EventRouteCollectionIterator
type EventRouteCollectionPage = original.EventRouteCollectionPage
type EventRoutesClient = original.EventRoutesClient
type IncomingRelationship = original.IncomingRelationship
type IncomingRelationshipCollection = original.IncomingRelationshipCollection
type IncomingRelationshipCollectionIterator = original.IncomingRelationshipCollectionIterator
type IncomingRelationshipCollectionPage = original.IncomingRelationshipCollectionPage
type InnerError = original.InnerError
type ListModelData = original.ListModelData
type ModelData = original.ModelData
type ModelsClient = original.ModelsClient
type PagedModelDataCollection = original.PagedModelDataCollection
type PagedModelDataCollectionIterator = original.PagedModelDataCollectionIterator
type PagedModelDataCollectionPage = original.PagedModelDataCollectionPage
type QueryClient = original.QueryClient
type QueryResult = original.QueryResult
type QuerySpecification = original.QuerySpecification
type RelationshipCollection = original.RelationshipCollection
type RelationshipCollectionIterator = original.RelationshipCollectionIterator
type RelationshipCollectionPage = original.RelationshipCollectionPage
type SetObject = original.SetObject

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewEventRouteCollectionIterator(page EventRouteCollectionPage) EventRouteCollectionIterator {
	return original.NewEventRouteCollectionIterator(page)
}
func NewEventRouteCollectionPage(getNextPage func(context.Context, EventRouteCollection) (EventRouteCollection, error)) EventRouteCollectionPage {
	return original.NewEventRouteCollectionPage(getNextPage)
}
func NewEventRoutesClient() EventRoutesClient {
	return original.NewEventRoutesClient()
}
func NewEventRoutesClientWithBaseURI(baseURI string) EventRoutesClient {
	return original.NewEventRoutesClientWithBaseURI(baseURI)
}
func NewIncomingRelationshipCollectionIterator(page IncomingRelationshipCollectionPage) IncomingRelationshipCollectionIterator {
	return original.NewIncomingRelationshipCollectionIterator(page)
}
func NewIncomingRelationshipCollectionPage(getNextPage func(context.Context, IncomingRelationshipCollection) (IncomingRelationshipCollection, error)) IncomingRelationshipCollectionPage {
	return original.NewIncomingRelationshipCollectionPage(getNextPage)
}
func NewModelsClient() ModelsClient {
	return original.NewModelsClient()
}
func NewModelsClientWithBaseURI(baseURI string) ModelsClient {
	return original.NewModelsClientWithBaseURI(baseURI)
}
func NewPagedModelDataCollectionIterator(page PagedModelDataCollectionPage) PagedModelDataCollectionIterator {
	return original.NewPagedModelDataCollectionIterator(page)
}
func NewPagedModelDataCollectionPage(getNextPage func(context.Context, PagedModelDataCollection) (PagedModelDataCollection, error)) PagedModelDataCollectionPage {
	return original.NewPagedModelDataCollectionPage(getNextPage)
}
func NewQueryClient() QueryClient {
	return original.NewQueryClient()
}
func NewQueryClientWithBaseURI(baseURI string) QueryClient {
	return original.NewQueryClientWithBaseURI(baseURI)
}
func NewRelationshipCollectionIterator(page RelationshipCollectionPage) RelationshipCollectionIterator {
	return original.NewRelationshipCollectionIterator(page)
}
func NewRelationshipCollectionPage(getNextPage func(context.Context, RelationshipCollection) (RelationshipCollection, error)) RelationshipCollectionPage {
	return original.NewRelationshipCollectionPage(getNextPage)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
