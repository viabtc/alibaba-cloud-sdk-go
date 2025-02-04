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

// ListClusterDetailFeature invokes the mse.ListClusterDetailFeature API synchronously
func (client *Client) ListClusterDetailFeature(request *ListClusterDetailFeatureRequest) (response *ListClusterDetailFeatureResponse, err error) {
	response = CreateListClusterDetailFeatureResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterDetailFeatureWithChan invokes the mse.ListClusterDetailFeature API asynchronously
func (client *Client) ListClusterDetailFeatureWithChan(request *ListClusterDetailFeatureRequest) (<-chan *ListClusterDetailFeatureResponse, <-chan error) {
	responseChan := make(chan *ListClusterDetailFeatureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterDetailFeature(request)
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

// ListClusterDetailFeatureWithCallback invokes the mse.ListClusterDetailFeature API asynchronously
func (client *Client) ListClusterDetailFeatureWithCallback(request *ListClusterDetailFeatureRequest, callback func(response *ListClusterDetailFeatureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterDetailFeatureResponse
		var err error
		defer close(result)
		response, err = client.ListClusterDetailFeature(request)
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

// ListClusterDetailFeatureRequest is the request struct for api ListClusterDetailFeature
type ListClusterDetailFeatureRequest struct {
	*requests.RpcRequest
	InstanceId     string `position:"Query" name:"InstanceId"`
	RequestPars    string `position:"Query" name:"RequestPars"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// ListClusterDetailFeatureResponse is the response struct for api ListClusterDetailFeature
type ListClusterDetailFeatureResponse struct {
	*responses.BaseResponse
	Success   bool                             `json:"Success" xml:"Success"`
	Message   string                           `json:"Message" xml:"Message"`
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	ErrorCode string                           `json:"ErrorCode" xml:"ErrorCode"`
	HttpCode  string                           `json:"HttpCode" xml:"HttpCode"`
	Data      []DataInListClusterDetailFeature `json:"Data" xml:"Data"`
}

// CreateListClusterDetailFeatureRequest creates a request to invoke ListClusterDetailFeature API
func CreateListClusterDetailFeatureRequest() (request *ListClusterDetailFeatureRequest) {
	request = &ListClusterDetailFeatureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListClusterDetailFeature", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterDetailFeatureResponse creates a response to parse from ListClusterDetailFeature response
func CreateListClusterDetailFeatureResponse() (response *ListClusterDetailFeatureResponse) {
	response = &ListClusterDetailFeatureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
