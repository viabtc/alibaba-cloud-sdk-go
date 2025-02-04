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

// CreateVGroup invokes the mse.CreateVGroup API synchronously
func (client *Client) CreateVGroup(request *CreateVGroupRequest) (response *CreateVGroupResponse, err error) {
	response = CreateCreateVGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVGroupWithChan invokes the mse.CreateVGroup API asynchronously
func (client *Client) CreateVGroupWithChan(request *CreateVGroupRequest) (<-chan *CreateVGroupResponse, <-chan error) {
	responseChan := make(chan *CreateVGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVGroup(request)
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

// CreateVGroupWithCallback invokes the mse.CreateVGroup API asynchronously
func (client *Client) CreateVGroupWithCallback(request *CreateVGroupRequest, callback func(response *CreateVGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateVGroup(request)
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

// CreateVGroupRequest is the request struct for api CreateVGroup
type CreateVGroupRequest struct {
	*requests.RpcRequest
	PrimaryUser         string `position:"Query" name:"PrimaryUser"`
	Name                string `position:"Query" name:"Name"`
	AcceptLanguage      string `position:"Query" name:"AcceptLanguage"`
	SeataServerUniqueId string `position:"Query" name:"SeataServerUniqueId"`
	Region              string `position:"Query" name:"Region"`
}

// CreateVGroupResponse is the response struct for api CreateVGroup
type CreateVGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateCreateVGroupRequest creates a request to invoke CreateVGroup API
func CreateCreateVGroupRequest() (request *CreateVGroupRequest) {
	request = &CreateVGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "CreateVGroup", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVGroupResponse creates a response to parse from CreateVGroup response
func CreateCreateVGroupResponse() (response *CreateVGroupResponse) {
	response = &CreateVGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
