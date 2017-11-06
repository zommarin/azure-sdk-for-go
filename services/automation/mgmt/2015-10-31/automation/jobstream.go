package automation

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

// JobStreamClient is the automation Client
type JobStreamClient struct {
	ManagementClient
}

// NewJobStreamClient creates an instance of the JobStreamClient client.
func NewJobStreamClient(subscriptionID string) JobStreamClient {
	return NewJobStreamClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobStreamClientWithBaseURI creates an instance of the JobStreamClient client.
func NewJobStreamClientWithBaseURI(baseURI string, subscriptionID string) JobStreamClient {
	return JobStreamClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get retrieve the job stream identified by job stream id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// id. jobStreamID is the job stream id.
func (client JobStreamClient) Get(resourceGroupName string, automationAccountName string, jobID string, jobStreamID string) (result JobStream, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobStreamClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, automationAccountName, jobID, jobStreamID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobStreamClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobStreamClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobStreamClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobStreamClient) GetPreparer(resourceGroupName string, automationAccountName string, jobID string, jobStreamID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"jobStreamId":           autorest.Encode("path", jobStreamID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/streams/{jobStreamId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobStreamClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobStreamClient) GetResponder(resp *http.Response) (result JobStream, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByJob retrieve a list of jobs streams identified by job id.
//
// resourceGroupName is the resource group name. automationAccountName is the automation account name. jobID is the job
// Id. filter is the filter to apply on the operation.
func (client JobStreamClient) ListByJob(resourceGroupName string, automationAccountName string, jobID string, filter string) (result JobStreamListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.JobStreamClient", "ListByJob")
	}

	req, err := client.ListByJobPreparer(resourceGroupName, automationAccountName, jobID, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobStreamClient", "ListByJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.JobStreamClient", "ListByJob", resp, "Failure sending request")
		return
	}

	result, err = client.ListByJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobStreamClient", "ListByJob", resp, "Failure responding to request")
	}

	return
}

// ListByJobPreparer prepares the ListByJob request.
func (client JobStreamClient) ListByJobPreparer(resourceGroupName string, automationAccountName string, jobID string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"jobId":                 autorest.Encode("path", jobID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/streams", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByJobSender sends the ListByJob request. The method will close the
// http.Response Body if it receives an error.
func (client JobStreamClient) ListByJobSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByJobResponder handles the response to the ListByJob request. The method always
// closes the http.Response Body.
func (client JobStreamClient) ListByJobResponder(resp *http.Response) (result JobStreamListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByJobNextResults retrieves the next set of results, if any.
func (client JobStreamClient) ListByJobNextResults(lastResults JobStreamListResult) (result JobStreamListResult, err error) {
	req, err := lastResults.JobStreamListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.JobStreamClient", "ListByJob", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.JobStreamClient", "ListByJob", resp, "Failure sending next results request")
	}

	result, err = client.ListByJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.JobStreamClient", "ListByJob", resp, "Failure responding to next results request")
	}

	return
}

// ListByJobComplete gets all elements from the list without paging.
func (client JobStreamClient) ListByJobComplete(resourceGroupName string, automationAccountName string, jobID string, filter string, cancel <-chan struct{}) (<-chan JobStream, <-chan error) {
	resultChan := make(chan JobStream)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByJob(resourceGroupName, automationAccountName, jobID, filter)
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
			list, err = client.ListByJobNextResults(list)
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
