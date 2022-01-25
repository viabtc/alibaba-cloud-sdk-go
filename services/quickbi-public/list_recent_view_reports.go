package quickbi_public

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

// ListRecentViewReports invokes the quickbi_public.ListRecentViewReports API synchronously
func (client *Client) ListRecentViewReports(request *ListRecentViewReportsRequest) (response *ListRecentViewReportsResponse, err error) {
	response = CreateListRecentViewReportsResponse()
	err = client.DoAction(request, response)
	return
}

// ListRecentViewReportsWithChan invokes the quickbi_public.ListRecentViewReports API asynchronously
func (client *Client) ListRecentViewReportsWithChan(request *ListRecentViewReportsRequest) (<-chan *ListRecentViewReportsResponse, <-chan error) {
	responseChan := make(chan *ListRecentViewReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRecentViewReports(request)
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

// ListRecentViewReportsWithCallback invokes the quickbi_public.ListRecentViewReports API asynchronously
func (client *Client) ListRecentViewReportsWithCallback(request *ListRecentViewReportsRequest, callback func(response *ListRecentViewReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRecentViewReportsResponse
		var err error
		defer close(result)
		response, err = client.ListRecentViewReports(request)
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

// ListRecentViewReportsRequest is the request struct for api ListRecentViewReports
type ListRecentViewReportsRequest struct {
	*requests.RpcRequest
	QueryMode   string           `position:"Query" name:"QueryMode"`
	TreeType    string           `position:"Query" name:"TreeType"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	SignType    string           `position:"Query" name:"SignType"`
	OffsetDay   requests.Integer `position:"Query" name:"OffsetDay"`
	Keyword     string           `position:"Query" name:"Keyword"`
	UserId      string           `position:"Query" name:"UserId"`
}

// ListRecentViewReportsResponse is the response struct for api ListRecentViewReports
type ListRecentViewReportsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateListRecentViewReportsRequest creates a request to invoke ListRecentViewReports API
func CreateListRecentViewReportsRequest() (request *ListRecentViewReportsRequest) {
	request = &ListRecentViewReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2020-08-09", "ListRecentViewReports", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRecentViewReportsResponse creates a response to parse from ListRecentViewReports response
func CreateListRecentViewReportsResponse() (response *ListRecentViewReportsResponse) {
	response = &ListRecentViewReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
