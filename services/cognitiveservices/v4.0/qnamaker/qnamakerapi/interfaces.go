package qnamakerapi

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

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v4.0/qnamaker"
	"github.com/Azure/go-autorest/autorest"
)

// EndpointSettingsClientAPI contains the set of methods on the EndpointSettingsClient type.
type EndpointSettingsClientAPI interface {
	GetSettings(ctx context.Context) (result qnamaker.EndpointSettingsDTO, err error)
	UpdateSettings(ctx context.Context, endpointSettingsPayload qnamaker.EndpointSettingsDTO) (result autorest.Response, err error)
}

var _ EndpointSettingsClientAPI = (*qnamaker.EndpointSettingsClient)(nil)

// EndpointKeysClientAPI contains the set of methods on the EndpointKeysClient type.
type EndpointKeysClientAPI interface {
	GetKeys(ctx context.Context) (result qnamaker.EndpointKeysDTO, err error)
	RefreshKeys(ctx context.Context, keyType string) (result qnamaker.EndpointKeysDTO, err error)
}

var _ EndpointKeysClientAPI = (*qnamaker.EndpointKeysClient)(nil)

// AlterationsClientAPI contains the set of methods on the AlterationsClient type.
type AlterationsClientAPI interface {
	Get(ctx context.Context) (result qnamaker.WordAlterationsDTO, err error)
	Replace(ctx context.Context, wordAlterations qnamaker.WordAlterationsDTO) (result autorest.Response, err error)
}

var _ AlterationsClientAPI = (*qnamaker.AlterationsClient)(nil)

// KnowledgebaseClientAPI contains the set of methods on the KnowledgebaseClient type.
type KnowledgebaseClientAPI interface {
	Create(ctx context.Context, createKbPayload qnamaker.CreateKbDTO) (result qnamaker.Operation, err error)
	Delete(ctx context.Context, kbID string) (result autorest.Response, err error)
	Download(ctx context.Context, kbID string, environment qnamaker.EnvironmentType, source string, changedSince string) (result qnamaker.QnADocumentsDTO, err error)
	GetDetails(ctx context.Context, kbID string) (result qnamaker.KnowledgebaseDTO, err error)
	ListAll(ctx context.Context) (result qnamaker.KnowledgebasesDTO, err error)
	Publish(ctx context.Context, kbID string) (result autorest.Response, err error)
	Replace(ctx context.Context, kbID string, replaceKb qnamaker.ReplaceKbDTO) (result autorest.Response, err error)
	Update(ctx context.Context, kbID string, updateKb qnamaker.UpdateKbOperationDTO) (result qnamaker.Operation, err error)
}

var _ KnowledgebaseClientAPI = (*qnamaker.KnowledgebaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	GetDetails(ctx context.Context, operationID string) (result qnamaker.Operation, err error)
}

var _ OperationsClientAPI = (*qnamaker.OperationsClient)(nil)
