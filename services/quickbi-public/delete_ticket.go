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

// DeleteTicket invokes the quickbi_public.DeleteTicket API synchronously
func (client *Client) DeleteTicket(request *DeleteTicketRequest) (response *DeleteTicketResponse, err error) {
	response = CreateDeleteTicketResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTicketWithChan invokes the quickbi_public.DeleteTicket API asynchronously
func (client *Client) DeleteTicketWithChan(request *DeleteTicketRequest) (<-chan *DeleteTicketResponse, <-chan error) {
	responseChan := make(chan *DeleteTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTicket(request)
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

// DeleteTicketWithCallback invokes the quickbi_public.DeleteTicket API asynchronously
func (client *Client) DeleteTicketWithCallback(request *DeleteTicketRequest, callback func(response *DeleteTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTicketResponse
		var err error
		defer close(result)
		response, err = client.DeleteTicket(request)
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

// DeleteTicketRequest is the request struct for api DeleteTicket
type DeleteTicketRequest struct {
	*requests.RpcRequest
	Ticket      string `position:"Query" name:"Ticket"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
}

// DeleteTicketResponse is the response struct for api DeleteTicket
type DeleteTicketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteTicketRequest creates a request to invoke DeleteTicket API
func CreateDeleteTicketRequest() (request *DeleteTicketRequest) {
	request = &DeleteTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2021-03-25", "DeleteTicket", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteTicketResponse creates a response to parse from DeleteTicket response
func CreateDeleteTicketResponse() (response *DeleteTicketResponse) {
	response = &DeleteTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
