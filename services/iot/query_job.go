package iot

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

// QueryJob invokes the iot.QueryJob API synchronously
func (client *Client) QueryJob(request *QueryJobRequest) (response *QueryJobResponse, err error) {
	response = CreateQueryJobResponse()
	err = client.DoAction(request, response)
	return
}

// QueryJobWithChan invokes the iot.QueryJob API asynchronously
func (client *Client) QueryJobWithChan(request *QueryJobRequest) (<-chan *QueryJobResponse, <-chan error) {
	responseChan := make(chan *QueryJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryJob(request)
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

// QueryJobWithCallback invokes the iot.QueryJob API asynchronously
func (client *Client) QueryJobWithCallback(request *QueryJobRequest, callback func(response *QueryJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryJobResponse
		var err error
		defer close(result)
		response, err = client.QueryJob(request)
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

// QueryJobRequest is the request struct for api QueryJob
type QueryJobRequest struct {
	*requests.RpcRequest
	JobId         string `position:"Query" name:"JobId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QueryJobResponse is the response struct for api QueryJob
type QueryJobResponse struct {
	*responses.BaseResponse
	RequestId    string         `json:"RequestId" xml:"RequestId"`
	Success      bool           `json:"Success" xml:"Success"`
	Code         string         `json:"Code" xml:"Code"`
	ErrorMessage string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryJob `json:"Data" xml:"Data"`
}

// CreateQueryJobRequest creates a request to invoke QueryJob API
func CreateQueryJobRequest() (request *QueryJobRequest) {
	request = &QueryJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryJob", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryJobResponse creates a response to parse from QueryJob response
func CreateQueryJobResponse() (response *QueryJobResponse) {
	response = &QueryJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}