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

// DeleteOfficeConversionTask invokes the imm.DeleteOfficeConversionTask API synchronously
func (client *Client) DeleteOfficeConversionTask(request *DeleteOfficeConversionTaskRequest) (response *DeleteOfficeConversionTaskResponse, err error) {
	response = CreateDeleteOfficeConversionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteOfficeConversionTaskWithChan invokes the imm.DeleteOfficeConversionTask API asynchronously
func (client *Client) DeleteOfficeConversionTaskWithChan(request *DeleteOfficeConversionTaskRequest) (<-chan *DeleteOfficeConversionTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteOfficeConversionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteOfficeConversionTask(request)
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

// DeleteOfficeConversionTaskWithCallback invokes the imm.DeleteOfficeConversionTask API asynchronously
func (client *Client) DeleteOfficeConversionTaskWithCallback(request *DeleteOfficeConversionTaskRequest, callback func(response *DeleteOfficeConversionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteOfficeConversionTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteOfficeConversionTask(request)
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

// DeleteOfficeConversionTaskRequest is the request struct for api DeleteOfficeConversionTask
type DeleteOfficeConversionTaskRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

// DeleteOfficeConversionTaskResponse is the response struct for api DeleteOfficeConversionTask
type DeleteOfficeConversionTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteOfficeConversionTaskRequest creates a request to invoke DeleteOfficeConversionTask API
func CreateDeleteOfficeConversionTaskRequest() (request *DeleteOfficeConversionTaskRequest) {
	request = &DeleteOfficeConversionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DeleteOfficeConversionTask", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteOfficeConversionTaskResponse creates a response to parse from DeleteOfficeConversionTask response
func CreateDeleteOfficeConversionTaskResponse() (response *DeleteOfficeConversionTaskResponse) {
	response = &DeleteOfficeConversionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
