package ccc

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

// SaveWebRTCStats invokes the ccc.SaveWebRTCStats API synchronously
func (client *Client) SaveWebRTCStats(request *SaveWebRTCStatsRequest) (response *SaveWebRTCStatsResponse, err error) {
	response = CreateSaveWebRTCStatsResponse()
	err = client.DoAction(request, response)
	return
}

// SaveWebRTCStatsWithChan invokes the ccc.SaveWebRTCStats API asynchronously
func (client *Client) SaveWebRTCStatsWithChan(request *SaveWebRTCStatsRequest) (<-chan *SaveWebRTCStatsResponse, <-chan error) {
	responseChan := make(chan *SaveWebRTCStatsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveWebRTCStats(request)
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

// SaveWebRTCStatsWithCallback invokes the ccc.SaveWebRTCStats API asynchronously
func (client *Client) SaveWebRTCStatsWithCallback(request *SaveWebRTCStatsRequest, callback func(response *SaveWebRTCStatsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveWebRTCStatsResponse
		var err error
		defer close(result)
		response, err = client.SaveWebRTCStats(request)
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

// SaveWebRTCStatsRequest is the request struct for api SaveWebRTCStats
type SaveWebRTCStatsRequest struct {
	*requests.RpcRequest
	CallId        string           `position:"Query" name:"CallId"`
	RecordTime    requests.Integer `position:"Query" name:"RecordTime"`
	CallStartTime requests.Integer `position:"Query" name:"CallStartTime"`
	Uid           string           `position:"Query" name:"Uid"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Stats         string           `position:"Query" name:"Stats"`
	TenantId      string           `position:"Query" name:"TenantId"`
	CalleeNumber  string           `position:"Query" name:"CalleeNumber"`
	CallerNumber  string           `position:"Query" name:"CallerNumber"`
}

// SaveWebRTCStatsResponse is the response struct for api SaveWebRTCStats
type SaveWebRTCStatsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RowCount       int64  `json:"RowCount" xml:"RowCount"`
}

// CreateSaveWebRTCStatsRequest creates a request to invoke SaveWebRTCStats API
func CreateSaveWebRTCStatsRequest() (request *SaveWebRTCStatsRequest) {
	request = &SaveWebRTCStatsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "SaveWebRTCStats", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveWebRTCStatsResponse creates a response to parse from SaveWebRTCStats response
func CreateSaveWebRTCStatsResponse() (response *SaveWebRTCStatsResponse) {
	response = &SaveWebRTCStatsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
