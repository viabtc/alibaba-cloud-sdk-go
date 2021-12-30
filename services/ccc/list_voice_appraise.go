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

// ListVoiceAppraise invokes the ccc.ListVoiceAppraise API synchronously
func (client *Client) ListVoiceAppraise(request *ListVoiceAppraiseRequest) (response *ListVoiceAppraiseResponse, err error) {
	response = CreateListVoiceAppraiseResponse()
	err = client.DoAction(request, response)
	return
}

// ListVoiceAppraiseWithChan invokes the ccc.ListVoiceAppraise API asynchronously
func (client *Client) ListVoiceAppraiseWithChan(request *ListVoiceAppraiseRequest) (<-chan *ListVoiceAppraiseResponse, <-chan error) {
	responseChan := make(chan *ListVoiceAppraiseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVoiceAppraise(request)
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

// ListVoiceAppraiseWithCallback invokes the ccc.ListVoiceAppraise API asynchronously
func (client *Client) ListVoiceAppraiseWithCallback(request *ListVoiceAppraiseRequest, callback func(response *ListVoiceAppraiseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVoiceAppraiseResponse
		var err error
		defer close(result)
		response, err = client.ListVoiceAppraise(request)
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

// ListVoiceAppraiseRequest is the request struct for api ListVoiceAppraise
type ListVoiceAppraiseRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListVoiceAppraiseResponse is the response struct for api ListVoiceAppraise
type ListVoiceAppraiseResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Notice         string      `json:"Notice" xml:"Notice"`
	ContactFlow    ContactFlow `json:"ContactFlow" xml:"ContactFlow"`
}

// CreateListVoiceAppraiseRequest creates a request to invoke ListVoiceAppraise API
func CreateListVoiceAppraiseRequest() (request *ListVoiceAppraiseRequest) {
	request = &ListVoiceAppraiseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListVoiceAppraise", "", "")
	request.Method = requests.POST
	return
}

// CreateListVoiceAppraiseResponse creates a response to parse from ListVoiceAppraise response
func CreateListVoiceAppraiseResponse() (response *ListVoiceAppraiseResponse) {
	response = &ListVoiceAppraiseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
