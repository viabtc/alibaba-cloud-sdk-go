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

// CreateVideoAnalyseTask invokes the imm.CreateVideoAnalyseTask API synchronously
func (client *Client) CreateVideoAnalyseTask(request *CreateVideoAnalyseTaskRequest) (response *CreateVideoAnalyseTaskResponse, err error) {
	response = CreateCreateVideoAnalyseTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVideoAnalyseTaskWithChan invokes the imm.CreateVideoAnalyseTask API asynchronously
func (client *Client) CreateVideoAnalyseTaskWithChan(request *CreateVideoAnalyseTaskRequest) (<-chan *CreateVideoAnalyseTaskResponse, <-chan error) {
	responseChan := make(chan *CreateVideoAnalyseTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVideoAnalyseTask(request)
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

// CreateVideoAnalyseTaskWithCallback invokes the imm.CreateVideoAnalyseTask API asynchronously
func (client *Client) CreateVideoAnalyseTaskWithCallback(request *CreateVideoAnalyseTaskRequest, callback func(response *CreateVideoAnalyseTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVideoAnalyseTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateVideoAnalyseTask(request)
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

// CreateVideoAnalyseTaskRequest is the request struct for api CreateVideoAnalyseTask
type CreateVideoAnalyseTaskRequest struct {
	*requests.RpcRequest
	GrabType        string           `position:"Query" name:"GrabType"`
	Project         string           `position:"Query" name:"Project"`
	StartTime       string           `position:"Query" name:"StartTime"`
	NotifyEndpoint  string           `position:"Query" name:"NotifyEndpoint"`
	NotifyTopicName string           `position:"Query" name:"NotifyTopicName"`
	EndTime         string           `position:"Query" name:"EndTime"`
	VideoUri        string           `position:"Query" name:"VideoUri"`
	SaveType        requests.Boolean `position:"Query" name:"SaveType"`
	Interval        string           `position:"Query" name:"Interval"`
	TgtUri          string           `position:"Query" name:"TgtUri"`
}

// CreateVideoAnalyseTaskResponse is the response struct for api CreateVideoAnalyseTask
type CreateVideoAnalyseTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	TaskType  string `json:"TaskType" xml:"TaskType"`
}

// CreateCreateVideoAnalyseTaskRequest creates a request to invoke CreateVideoAnalyseTask API
func CreateCreateVideoAnalyseTaskRequest() (request *CreateVideoAnalyseTaskRequest) {
	request = &CreateVideoAnalyseTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "CreateVideoAnalyseTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateVideoAnalyseTaskResponse creates a response to parse from CreateVideoAnalyseTask response
func CreateCreateVideoAnalyseTaskResponse() (response *CreateVideoAnalyseTaskResponse) {
	response = &CreateVideoAnalyseTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
