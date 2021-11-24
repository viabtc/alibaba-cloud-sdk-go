package mse

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListClusterConnectionTypes invokes the mse.ListClusterConnectionTypes API synchronously
func (client *Client) ListClusterConnectionTypes(request *ListClusterConnectionTypesRequest) (response *ListClusterConnectionTypesResponse, err error) {
	response = CreateListClusterConnectionTypesResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterConnectionTypesWithChan invokes the mse.ListClusterConnectionTypes API asynchronously
func (client *Client) ListClusterConnectionTypesWithChan(request *ListClusterConnectionTypesRequest) (<-chan *ListClusterConnectionTypesResponse, <-chan error) {
	responseChan := make(chan *ListClusterConnectionTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterConnectionTypes(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListClusterConnectionTypesWithCallback invokes the mse.ListClusterConnectionTypes API asynchronously
func (client *Client) ListClusterConnectionTypesWithCallback(request *ListClusterConnectionTypesRequest, callback func(response *ListClusterConnectionTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterConnectionTypesResponse
		var err error
		defer close(result)
		response, err = client.ListClusterConnectionTypes(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListClusterConnectionTypesRequest is the request struct for api ListClusterConnectionTypes
type ListClusterConnectionTypesRequest struct {
	*requests.RpcRequest
}

// ListClusterConnectionTypesResponse is the response struct for api ListClusterConnectionTypes
type ListClusterConnectionTypesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	ErrorCode      string     `json:"ErrorCode" xml:"ErrorCode"`
	Code           int        `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	DynamicMessage string     `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           []DataItem `json:"Data" xml:"Data"`
}

// CreateListClusterConnectionTypesRequest creates a request to invoke ListClusterConnectionTypes API
func CreateListClusterConnectionTypesRequest() (request *ListClusterConnectionTypesRequest) {
	request = &ListClusterConnectionTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListClusterConnectionTypes", "", "")
	request.Method = requests.POST
	return
}

// CreateListClusterConnectionTypesResponse creates a response to parse from ListClusterConnectionTypes response
func CreateListClusterConnectionTypesResponse() (response *ListClusterConnectionTypesResponse) {
	response = &ListClusterConnectionTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
