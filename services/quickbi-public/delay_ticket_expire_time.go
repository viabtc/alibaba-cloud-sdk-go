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

// DelayTicketExpireTime invokes the quickbi_public.DelayTicketExpireTime API synchronously
func (client *Client) DelayTicketExpireTime(request *DelayTicketExpireTimeRequest) (response *DelayTicketExpireTimeResponse, err error) {
	response = CreateDelayTicketExpireTimeResponse()
	err = client.DoAction(request, response)
	return
}

// DelayTicketExpireTimeWithChan invokes the quickbi_public.DelayTicketExpireTime API asynchronously
func (client *Client) DelayTicketExpireTimeWithChan(request *DelayTicketExpireTimeRequest) (<-chan *DelayTicketExpireTimeResponse, <-chan error) {
	responseChan := make(chan *DelayTicketExpireTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DelayTicketExpireTime(request)
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

// DelayTicketExpireTimeWithCallback invokes the quickbi_public.DelayTicketExpireTime API asynchronously
func (client *Client) DelayTicketExpireTimeWithCallback(request *DelayTicketExpireTimeRequest, callback func(response *DelayTicketExpireTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DelayTicketExpireTimeResponse
		var err error
		defer close(result)
		response, err = client.DelayTicketExpireTime(request)
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

// DelayTicketExpireTimeRequest is the request struct for api DelayTicketExpireTime
type DelayTicketExpireTimeRequest struct {
	*requests.RpcRequest
	ExpireTime  requests.Integer `position:"Query" name:"ExpireTime"`
	Ticket      string           `position:"Query" name:"Ticket"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	SignType    string           `position:"Query" name:"SignType"`
}

// DelayTicketExpireTimeResponse is the response struct for api DelayTicketExpireTime
type DelayTicketExpireTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDelayTicketExpireTimeRequest creates a request to invoke DelayTicketExpireTime API
func CreateDelayTicketExpireTimeRequest() (request *DelayTicketExpireTimeRequest) {
	request = &DelayTicketExpireTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2021-03-25", "DelayTicketExpireTime", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDelayTicketExpireTimeResponse creates a response to parse from DelayTicketExpireTime response
func CreateDelayTicketExpireTimeResponse() (response *DelayTicketExpireTimeResponse) {
	response = &DelayTicketExpireTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
