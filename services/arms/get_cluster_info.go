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

// GetClusterInfo invokes the arms.GetClusterInfo API synchronously
func (client *Client) GetClusterInfo(request *GetClusterInfoRequest) (response *GetClusterInfoResponse, err error) {
	response = CreateGetClusterInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetClusterInfoWithChan invokes the arms.GetClusterInfo API asynchronously
func (client *Client) GetClusterInfoWithChan(request *GetClusterInfoRequest) (<-chan *GetClusterInfoResponse, <-chan error) {
	responseChan := make(chan *GetClusterInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetClusterInfo(request)
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

// GetClusterInfoWithCallback invokes the arms.GetClusterInfo API asynchronously
func (client *Client) GetClusterInfoWithCallback(request *GetClusterInfoRequest, callback func(response *GetClusterInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetClusterInfoResponse
		var err error
		defer close(result)
		response, err = client.GetClusterInfo(request)
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

// GetClusterInfoRequest is the request struct for api GetClusterInfo
type GetClusterInfoRequest struct {
	*requests.RpcRequest
	QueryUserId string `position:"Query" name:"QueryUserId"`
	ClusterId   string `position:"Query" name:"ClusterId"`
}

// GetClusterInfoResponse is the response struct for api GetClusterInfo
type GetClusterInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetClusterInfoRequest creates a request to invoke GetClusterInfo API
func CreateGetClusterInfoRequest() (request *GetClusterInfoRequest) {
	request = &GetClusterInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetClusterInfo", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetClusterInfoResponse creates a response to parse from GetClusterInfo response
func CreateGetClusterInfoResponse() (response *GetClusterInfoResponse) {
	response = &GetClusterInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
