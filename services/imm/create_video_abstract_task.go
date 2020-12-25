package imm

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

// CreateVideoAbstractTask invokes the imm.CreateVideoAbstractTask API synchronously
func (client *Client) CreateVideoAbstractTask(request *CreateVideoAbstractTaskRequest) (response *CreateVideoAbstractTaskResponse, err error) {
	response = CreateCreateVideoAbstractTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVideoAbstractTaskWithChan invokes the imm.CreateVideoAbstractTask API asynchronously
func (client *Client) CreateVideoAbstractTaskWithChan(request *CreateVideoAbstractTaskRequest) (<-chan *CreateVideoAbstractTaskResponse, <-chan error) {
	responseChan := make(chan *CreateVideoAbstractTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVideoAbstractTask(request)
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

// CreateVideoAbstractTaskWithCallback invokes the imm.CreateVideoAbstractTask API asynchronously
func (client *Client) CreateVideoAbstractTaskWithCallback(request *CreateVideoAbstractTaskRequest, callback func(response *CreateVideoAbstractTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVideoAbstractTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateVideoAbstractTask(request)
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

// CreateVideoAbstractTaskRequest is the request struct for api CreateVideoAbstractTask
type CreateVideoAbstractTaskRequest struct {
	*requests.RpcRequest
	TargetVideoUri  string           `position:"Query" name:"TargetVideoUri"`
	Project         string           `position:"Query" name:"Project"`
	NotifyEndpoint  string           `position:"Query" name:"NotifyEndpoint"`
	NotifyTopicName string           `position:"Query" name:"NotifyTopicName"`
	VideoUri        string           `position:"Query" name:"VideoUri"`
	AbstractLength  requests.Integer `position:"Query" name:"AbstractLength"`
	TargetClipsUri  string           `position:"Query" name:"TargetClipsUri"`
}

// CreateVideoAbstractTaskResponse is the response struct for api CreateVideoAbstractTask
type CreateVideoAbstractTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateVideoAbstractTaskRequest creates a request to invoke CreateVideoAbstractTask API
func CreateCreateVideoAbstractTaskRequest() (request *CreateVideoAbstractTaskRequest) {
	request = &CreateVideoAbstractTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateVideoAbstractTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateVideoAbstractTaskResponse creates a response to parse from CreateVideoAbstractTask response
func CreateCreateVideoAbstractTaskResponse() (response *CreateVideoAbstractTaskResponse) {
	response = &CreateVideoAbstractTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
