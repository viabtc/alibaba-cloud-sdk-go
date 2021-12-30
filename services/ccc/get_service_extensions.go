package ccc

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

// GetServiceExtensions invokes the ccc.GetServiceExtensions API synchronously
func (client *Client) GetServiceExtensions(request *GetServiceExtensionsRequest) (response *GetServiceExtensionsResponse, err error) {
	response = CreateGetServiceExtensionsResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceExtensionsWithChan invokes the ccc.GetServiceExtensions API asynchronously
func (client *Client) GetServiceExtensionsWithChan(request *GetServiceExtensionsRequest) (<-chan *GetServiceExtensionsResponse, <-chan error) {
	responseChan := make(chan *GetServiceExtensionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceExtensions(request)
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

// GetServiceExtensionsWithCallback invokes the ccc.GetServiceExtensions API asynchronously
func (client *Client) GetServiceExtensionsWithCallback(request *GetServiceExtensionsRequest, callback func(response *GetServiceExtensionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceExtensionsResponse
		var err error
		defer close(result)
		response, err = client.GetServiceExtensions(request)
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

// GetServiceExtensionsRequest is the request struct for api GetServiceExtensions
type GetServiceExtensionsRequest struct {
	*requests.RpcRequest
	ServiceType string `position:"Query" name:"ServiceType"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

// GetServiceExtensionsResponse is the response struct for api GetServiceExtensions
type GetServiceExtensionsResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	Code              string            `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	HttpStatusCode    int               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ServiceExtensions ServiceExtensions `json:"ServiceExtensions" xml:"ServiceExtensions"`
}

// CreateGetServiceExtensionsRequest creates a request to invoke GetServiceExtensions API
func CreateGetServiceExtensionsRequest() (request *GetServiceExtensionsRequest) {
	request = &GetServiceExtensionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetServiceExtensions", "", "")
	request.Method = requests.POST
	return
}

// CreateGetServiceExtensionsResponse creates a response to parse from GetServiceExtensions response
func CreateGetServiceExtensionsResponse() (response *GetServiceExtensionsResponse) {
	response = &GetServiceExtensionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
