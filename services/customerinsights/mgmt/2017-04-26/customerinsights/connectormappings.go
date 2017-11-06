package customerinsights

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ConnectorMappingsClient is the the Azure Customer Insights management API provides a RESTful set of web services
// that interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type ConnectorMappingsClient struct {
	ManagementClient
}

// NewConnectorMappingsClient creates an instance of the ConnectorMappingsClient client.
func NewConnectorMappingsClient(subscriptionID string) ConnectorMappingsClient {
	return NewConnectorMappingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConnectorMappingsClientWithBaseURI creates an instance of the ConnectorMappingsClient client.
func NewConnectorMappingsClientWithBaseURI(baseURI string, subscriptionID string) ConnectorMappingsClient {
	return ConnectorMappingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a connector mapping or updates an existing connector mapping in the connector.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. connectorName is the name of
// the connector. mappingName is the name of the connector mapping. parameters is parameters supplied to the
// CreateOrUpdate Connector Mapping operation.
func (client ConnectorMappingsClient) CreateOrUpdate(resourceGroupName string, hubName string, connectorName string, mappingName string, parameters ConnectorMappingResourceFormat) (result ConnectorMappingResourceFormat, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: mappingName,
			Constraints: []validation.Constraint{{Target: "mappingName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "mappingName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "mappingName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ConnectorMapping", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ConnectorMapping.EntityTypeName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ConnectorMapping.MappingProperties", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ConnectorMapping.MappingProperties.ErrorManagement", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ConnectorMapping.MappingProperties.Format", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.ConnectorMapping.MappingProperties.Format.FormatType", Name: validation.Null, Rule: true, Chain: nil}}},
							{Target: "parameters.ConnectorMapping.MappingProperties.Availability", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.ConnectorMapping.MappingProperties.Availability.Interval", Name: validation.Null, Rule: true, Chain: nil}}},
							{Target: "parameters.ConnectorMapping.MappingProperties.Structure", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ConnectorMapping.MappingProperties.CompleteOperation", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "customerinsights.ConnectorMappingsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, hubName, connectorName, mappingName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConnectorMappingsClient) CreateOrUpdatePreparer(resourceGroupName string, hubName string, connectorName string, mappingName string, parameters ConnectorMappingResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectorName":     autorest.Encode("path", connectorName),
		"hubName":           autorest.Encode("path", hubName),
		"mappingName":       autorest.Encode("path", mappingName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorMappingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConnectorMappingsClient) CreateOrUpdateResponder(resp *http.Response) (result ConnectorMappingResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a connector mapping in the connector.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. connectorName is the name of
// the connector. mappingName is the name of the connector mapping.
func (client ConnectorMappingsClient) Delete(resourceGroupName string, hubName string, connectorName string, mappingName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, hubName, connectorName, mappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConnectorMappingsClient) DeletePreparer(resourceGroupName string, hubName string, connectorName string, mappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectorName":     autorest.Encode("path", connectorName),
		"hubName":           autorest.Encode("path", hubName),
		"mappingName":       autorest.Encode("path", mappingName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorMappingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConnectorMappingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a connector mapping in the connector.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. connectorName is the name of
// the connector. mappingName is the name of the connector mapping.
func (client ConnectorMappingsClient) Get(resourceGroupName string, hubName string, connectorName string, mappingName string) (result ConnectorMappingResourceFormat, err error) {
	req, err := client.GetPreparer(resourceGroupName, hubName, connectorName, mappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConnectorMappingsClient) GetPreparer(resourceGroupName string, hubName string, connectorName string, mappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectorName":     autorest.Encode("path", connectorName),
		"hubName":           autorest.Encode("path", hubName),
		"mappingName":       autorest.Encode("path", mappingName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorMappingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConnectorMappingsClient) GetResponder(resp *http.Response) (result ConnectorMappingResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByConnector gets all the connector mappings in the specified connector.
//
// resourceGroupName is the name of the resource group. hubName is the name of the hub. connectorName is the name of
// the connector.
func (client ConnectorMappingsClient) ListByConnector(resourceGroupName string, hubName string, connectorName string) (result ConnectorMappingListResult, err error) {
	req, err := client.ListByConnectorPreparer(resourceGroupName, hubName, connectorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "ListByConnector", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByConnectorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "ListByConnector", resp, "Failure sending request")
		return
	}

	result, err = client.ListByConnectorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "ListByConnector", resp, "Failure responding to request")
	}

	return
}

// ListByConnectorPreparer prepares the ListByConnector request.
func (client ConnectorMappingsClient) ListByConnectorPreparer(resourceGroupName string, hubName string, connectorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectorName":     autorest.Encode("path", connectorName),
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-26"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByConnectorSender sends the ListByConnector request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorMappingsClient) ListByConnectorSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByConnectorResponder handles the response to the ListByConnector request. The method always
// closes the http.Response Body.
func (client ConnectorMappingsClient) ListByConnectorResponder(resp *http.Response) (result ConnectorMappingListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByConnectorNextResults retrieves the next set of results, if any.
func (client ConnectorMappingsClient) ListByConnectorNextResults(lastResults ConnectorMappingListResult) (result ConnectorMappingListResult, err error) {
	req, err := lastResults.ConnectorMappingListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "ListByConnector", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByConnectorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "ListByConnector", resp, "Failure sending next results request")
	}

	result, err = client.ListByConnectorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorMappingsClient", "ListByConnector", resp, "Failure responding to next results request")
	}

	return
}

// ListByConnectorComplete gets all elements from the list without paging.
func (client ConnectorMappingsClient) ListByConnectorComplete(resourceGroupName string, hubName string, connectorName string, cancel <-chan struct{}) (<-chan ConnectorMappingResourceFormat, <-chan error) {
	resultChan := make(chan ConnectorMappingResourceFormat)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByConnector(resourceGroupName, hubName, connectorName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByConnectorNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
