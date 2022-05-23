package arms

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

// DeleteRetcodeApp invokes the arms.DeleteRetcodeApp API synchronously
func (client *Client) DeleteRetcodeApp(request *DeleteRetcodeAppRequest) (response *DeleteRetcodeAppResponse, err error) {
	response = CreateDeleteRetcodeAppResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRetcodeAppWithChan invokes the arms.DeleteRetcodeApp API asynchronously
func (client *Client) DeleteRetcodeAppWithChan(request *DeleteRetcodeAppRequest) (<-chan *DeleteRetcodeAppResponse, <-chan error) {
	responseChan := make(chan *DeleteRetcodeAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRetcodeApp(request)
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

// DeleteRetcodeAppWithCallback invokes the arms.DeleteRetcodeApp API asynchronously
func (client *Client) DeleteRetcodeAppWithCallback(request *DeleteRetcodeAppRequest, callback func(response *DeleteRetcodeAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRetcodeAppResponse
		var err error
		defer close(result)
		response, err = client.DeleteRetcodeApp(request)
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

// DeleteRetcodeAppRequest is the request struct for api DeleteRetcodeApp
type DeleteRetcodeAppRequest struct {
	*requests.RpcRequest
	AppId string `position:"Query" name:"AppId"`
}

// DeleteRetcodeAppResponse is the response struct for api DeleteRetcodeApp
type DeleteRetcodeAppResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRetcodeAppRequest creates a request to invoke DeleteRetcodeApp API
func CreateDeleteRetcodeAppRequest() (request *DeleteRetcodeAppRequest) {
	request = &DeleteRetcodeAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeleteRetcodeApp", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteRetcodeAppResponse creates a response to parse from DeleteRetcodeApp response
func CreateDeleteRetcodeAppResponse() (response *DeleteRetcodeAppResponse) {
	response = &DeleteRetcodeAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
