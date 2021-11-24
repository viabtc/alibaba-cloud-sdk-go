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

// CreateNacosConfig invokes the mse.CreateNacosConfig API synchronously
func (client *Client) CreateNacosConfig(request *CreateNacosConfigRequest) (response *CreateNacosConfigResponse, err error) {
	response = CreateCreateNacosConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNacosConfigWithChan invokes the mse.CreateNacosConfig API asynchronously
func (client *Client) CreateNacosConfigWithChan(request *CreateNacosConfigRequest) (<-chan *CreateNacosConfigResponse, <-chan error) {
	responseChan := make(chan *CreateNacosConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNacosConfig(request)
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

// CreateNacosConfigWithCallback invokes the mse.CreateNacosConfig API asynchronously
func (client *Client) CreateNacosConfigWithCallback(request *CreateNacosConfigRequest, callback func(response *CreateNacosConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNacosConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateNacosConfig(request)
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

// CreateNacosConfigRequest is the request struct for api CreateNacosConfig
type CreateNacosConfigRequest struct {
	*requests.RpcRequest
	Type        string `position:"Query" name:"Type"`
	Content     string `position:"Query" name:"Content"`
	Tags        string `position:"Query" name:"Tags"`
	BetaIps     string `position:"Query" name:"BetaIps"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	DataId      string `position:"Query" name:"DataId"`
	AppName     string `position:"Query" name:"AppName"`
	NamespaceId string `position:"Query" name:"NamespaceId"`
	Group       string `position:"Query" name:"Group"`
	Desc        string `position:"Query" name:"Desc"`
}

// CreateNacosConfigResponse is the response struct for api CreateNacosConfig
type CreateNacosConfigResponse struct {
	*responses.BaseResponse
	HttpCode  string `json:"HttpCode" xml:"HttpCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateNacosConfigRequest creates a request to invoke CreateNacosConfig API
func CreateCreateNacosConfigRequest() (request *CreateNacosConfigRequest) {
	request = &CreateNacosConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "CreateNacosConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateNacosConfigResponse creates a response to parse from CreateNacosConfig response
func CreateCreateNacosConfigResponse() (response *CreateNacosConfigResponse) {
	response = &CreateNacosConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
