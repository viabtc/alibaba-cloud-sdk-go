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

// DescribeContactGroups invokes the arms.DescribeContactGroups API synchronously
func (client *Client) DescribeContactGroups(request *DescribeContactGroupsRequest) (response *DescribeContactGroupsResponse, err error) {
	response = CreateDescribeContactGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeContactGroupsWithChan invokes the arms.DescribeContactGroups API asynchronously
func (client *Client) DescribeContactGroupsWithChan(request *DescribeContactGroupsRequest) (<-chan *DescribeContactGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeContactGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeContactGroups(request)
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

// DescribeContactGroupsWithCallback invokes the arms.DescribeContactGroups API asynchronously
func (client *Client) DescribeContactGroupsWithCallback(request *DescribeContactGroupsRequest, callback func(response *DescribeContactGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeContactGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeContactGroups(request)
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

// DescribeContactGroupsRequest is the request struct for api DescribeContactGroups
type DescribeContactGroupsRequest struct {
	*requests.RpcRequest
	Size             requests.Integer `position:"Query" name:"Size"`
	IsDetail         requests.Boolean `position:"Query" name:"IsDetail"`
	Page             requests.Integer `position:"Query" name:"Page"`
	ContactGroupName string           `position:"Query" name:"ContactGroupName"`
}

// DescribeContactGroupsResponse is the response struct for api DescribeContactGroups
type DescribeContactGroupsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateDescribeContactGroupsRequest creates a request to invoke DescribeContactGroups API
func CreateDescribeContactGroupsRequest() (request *DescribeContactGroupsRequest) {
	request = &DescribeContactGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DescribeContactGroups", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeContactGroupsResponse creates a response to parse from DescribeContactGroups response
func CreateDescribeContactGroupsResponse() (response *DescribeContactGroupsResponse) {
	response = &DescribeContactGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
