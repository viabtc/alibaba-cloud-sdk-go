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

// DeleteNacosConfigs invokes the mse.DeleteNacosConfigs API synchronously
func (client *Client) DeleteNacosConfigs(request *DeleteNacosConfigsRequest) (response *DeleteNacosConfigsResponse, err error) {
	response = CreateDeleteNacosConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNacosConfigsWithChan invokes the mse.DeleteNacosConfigs API asynchronously
func (client *Client) DeleteNacosConfigsWithChan(request *DeleteNacosConfigsRequest) (<-chan *DeleteNacosConfigsResponse, <-chan error) {
	responseChan := make(chan *DeleteNacosConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNacosConfigs(request)
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

// DeleteNacosConfigsWithCallback invokes the mse.DeleteNacosConfigs API asynchronously
func (client *Client) DeleteNacosConfigsWithCallback(request *DeleteNacosConfigsRequest, callback func(response *DeleteNacosConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNacosConfigsResponse
		var err error
		defer close(result)
		response, err = client.DeleteNacosConfigs(request)
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

// DeleteNacosConfigsRequest is the request struct for api DeleteNacosConfigs
type DeleteNacosConfigsRequest struct {
	*requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	NamespaceId string `position:"Query" name:"NamespaceId"`
	Ids         string `position:"Query" name:"Ids"`
}

// DeleteNacosConfigsResponse is the response struct for api DeleteNacosConfigs
type DeleteNacosConfigsResponse struct {
	*responses.BaseResponse
	HttpCode  string `json:"HttpCode" xml:"HttpCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteNacosConfigsRequest creates a request to invoke DeleteNacosConfigs API
func CreateDeleteNacosConfigsRequest() (request *DeleteNacosConfigsRequest) {
	request = &DeleteNacosConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteNacosConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteNacosConfigsResponse creates a response to parse from DeleteNacosConfigs response
func CreateDeleteNacosConfigsResponse() (response *DeleteNacosConfigsResponse) {
	response = &DeleteNacosConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
