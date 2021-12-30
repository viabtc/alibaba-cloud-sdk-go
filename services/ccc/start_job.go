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

// StartJob invokes the ccc.StartJob API synchronously
func (client *Client) StartJob(request *StartJobRequest) (response *StartJobResponse, err error) {
	response = CreateStartJobResponse()
	err = client.DoAction(request, response)
	return
}

// StartJobWithChan invokes the ccc.StartJob API asynchronously
func (client *Client) StartJobWithChan(request *StartJobRequest) (<-chan *StartJobResponse, <-chan error) {
	responseChan := make(chan *StartJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartJob(request)
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

// StartJobWithCallback invokes the ccc.StartJob API asynchronously
func (client *Client) StartJobWithCallback(request *StartJobRequest, callback func(response *StartJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartJobResponse
		var err error
		defer close(result)
		response, err = client.StartJob(request)
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

// StartJobRequest is the request struct for api StartJob
type StartJobRequest struct {
	*requests.RpcRequest
	GroupId              string           `position:"Query" name:"GroupId"`
	JobJson              string           `position:"Query" name:"JobJson"`
	CallingNumber        *[]string        `position:"Query" name:"CallingNumber"  type:"Repeated"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	SelfHostedCallCenter requests.Boolean `position:"Query" name:"SelfHostedCallCenter"`
	ScenarioId           string           `position:"Query" name:"ScenarioId"`
}

// StartJobResponse is the response struct for api StartJob
type StartJobResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	TaskIds        []KeyValuePair `json:"TaskIds" xml:"TaskIds"`
}

// CreateStartJobRequest creates a request to invoke StartJob API
func CreateStartJobRequest() (request *StartJobRequest) {
	request = &StartJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "StartJob", "", "")
	request.Method = requests.POST
	return
}

// CreateStartJobResponse creates a response to parse from StartJob response
func CreateStartJobResponse() (response *StartJobResponse) {
	response = &StartJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
