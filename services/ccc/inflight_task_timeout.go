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

// InflightTaskTimeout invokes the ccc.InflightTaskTimeout API synchronously
func (client *Client) InflightTaskTimeout(request *InflightTaskTimeoutRequest) (response *InflightTaskTimeoutResponse, err error) {
	response = CreateInflightTaskTimeoutResponse()
	err = client.DoAction(request, response)
	return
}

// InflightTaskTimeoutWithChan invokes the ccc.InflightTaskTimeout API asynchronously
func (client *Client) InflightTaskTimeoutWithChan(request *InflightTaskTimeoutRequest) (<-chan *InflightTaskTimeoutResponse, <-chan error) {
	responseChan := make(chan *InflightTaskTimeoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InflightTaskTimeout(request)
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

// InflightTaskTimeoutWithCallback invokes the ccc.InflightTaskTimeout API asynchronously
func (client *Client) InflightTaskTimeoutWithCallback(request *InflightTaskTimeoutRequest, callback func(response *InflightTaskTimeoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InflightTaskTimeoutResponse
		var err error
		defer close(result)
		response, err = client.InflightTaskTimeout(request)
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

// InflightTaskTimeoutRequest is the request struct for api InflightTaskTimeout
type InflightTaskTimeoutRequest struct {
	*requests.RpcRequest
	InstanceId      string           `position:"Query" name:"InstanceId"`
	InstanceOwnerId requests.Integer `position:"Query" name:"InstanceOwnerId"`
	TaskId          string           `position:"Query" name:"TaskId"`
}

// InflightTaskTimeoutResponse is the response struct for api InflightTaskTimeout
type InflightTaskTimeoutResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateInflightTaskTimeoutRequest creates a request to invoke InflightTaskTimeout API
func CreateInflightTaskTimeoutRequest() (request *InflightTaskTimeoutRequest) {
	request = &InflightTaskTimeoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "InflightTaskTimeout", "", "")
	request.Method = requests.POST
	return
}

// CreateInflightTaskTimeoutResponse creates a response to parse from InflightTaskTimeout response
func CreateInflightTaskTimeoutResponse() (response *InflightTaskTimeoutResponse) {
	response = &InflightTaskTimeoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
