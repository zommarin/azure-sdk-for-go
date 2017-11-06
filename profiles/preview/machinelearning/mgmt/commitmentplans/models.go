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

package commitmentplans

import original "github.com/Azure/azure-sdk-for-go/services/machinelearning/mgmt/2016-05-01-preview/commitmentplans"

type UsageHistoryClient = original.UsageHistoryClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type CommitmentAssociationsClient = original.CommitmentAssociationsClient
type GroupClient = original.GroupClient
type CommitmentAssociation = original.CommitmentAssociation
type CommitmentAssociationListResult = original.CommitmentAssociationListResult
type CommitmentAssociationProperties = original.CommitmentAssociationProperties
type CommitmentPlan = original.CommitmentPlan
type ListResult = original.ListResult
type MoveCommitmentAssociationRequest = original.MoveCommitmentAssociationRequest
type PatchPayload = original.PatchPayload
type PlanQuantity = original.PlanQuantity
type PlanUsageHistory = original.PlanUsageHistory
type PlanUsageHistoryListResult = original.PlanUsageHistoryListResult
type Properties = original.Properties
type Resource = original.Resource
type ResourceSku = original.ResourceSku

func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageHistoryClient(subscriptionID string) UsageHistoryClient {
	return original.NewUsageHistoryClient(subscriptionID)
}
func NewUsageHistoryClientWithBaseURI(baseURI string, subscriptionID string) UsageHistoryClient {
	return original.NewUsageHistoryClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
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
func NewCommitmentAssociationsClient(subscriptionID string) CommitmentAssociationsClient {
	return original.NewCommitmentAssociationsClient(subscriptionID)
}
func NewCommitmentAssociationsClientWithBaseURI(baseURI string, subscriptionID string) CommitmentAssociationsClient {
	return original.NewCommitmentAssociationsClientWithBaseURI(baseURI, subscriptionID)
}
