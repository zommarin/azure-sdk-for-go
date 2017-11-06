// +build go1.9

// Copyright 2017 Microsoft Corporation
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
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package documentdb

import original "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"

type OperationsClient = original.OperationsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type DatabaseAccountsClient = original.DatabaseAccountsClient
type DatabaseAccountKind = original.DatabaseAccountKind

const (
	GlobalDocumentDB	DatabaseAccountKind	= original.GlobalDocumentDB
	MongoDB			DatabaseAccountKind	= original.MongoDB
	Parse			DatabaseAccountKind	= original.Parse
)

type DatabaseAccountOfferType = original.DatabaseAccountOfferType

const (
	Standard DatabaseAccountOfferType = original.Standard
)

type DefaultConsistencyLevel = original.DefaultConsistencyLevel

const (
	BoundedStaleness	DefaultConsistencyLevel	= original.BoundedStaleness
	ConsistentPrefix	DefaultConsistencyLevel	= original.ConsistentPrefix
	Eventual		DefaultConsistencyLevel	= original.Eventual
	Session			DefaultConsistencyLevel	= original.Session
	Strong			DefaultConsistencyLevel	= original.Strong
)

type KeyKind = original.KeyKind

const (
	Primary			KeyKind	= original.Primary
	PrimaryReadonly		KeyKind	= original.PrimaryReadonly
	Secondary		KeyKind	= original.Secondary
	SecondaryReadonly	KeyKind	= original.SecondaryReadonly
)

type ConsistencyPolicy = original.ConsistencyPolicy
type DatabaseAccount = original.DatabaseAccount
type DatabaseAccountConnectionString = original.DatabaseAccountConnectionString
type DatabaseAccountCreateUpdateParameters = original.DatabaseAccountCreateUpdateParameters
type DatabaseAccountCreateUpdateProperties = original.DatabaseAccountCreateUpdateProperties
type DatabaseAccountListConnectionStringsResult = original.DatabaseAccountListConnectionStringsResult
type DatabaseAccountListKeysResult = original.DatabaseAccountListKeysResult
type DatabaseAccountListReadOnlyKeysResult = original.DatabaseAccountListReadOnlyKeysResult
type DatabaseAccountPatchParameters = original.DatabaseAccountPatchParameters
type DatabaseAccountProperties = original.DatabaseAccountProperties
type DatabaseAccountRegenerateKeyParameters = original.DatabaseAccountRegenerateKeyParameters
type DatabaseAccountsListResult = original.DatabaseAccountsListResult
type FailoverPolicies = original.FailoverPolicies
type FailoverPolicy = original.FailoverPolicy
type Location = original.Location
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Resource = original.Resource

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseAccountsClient(subscriptionID string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClient(subscriptionID)
}
func NewDatabaseAccountsClientWithBaseURI(baseURI string, subscriptionID string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClientWithBaseURI(baseURI, subscriptionID)
}
