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

// AppendInstancesToPrometheusGlobalView invokes the arms.AppendInstancesToPrometheusGlobalView API synchronously
func (client *Client) AppendInstancesToPrometheusGlobalView(request *AppendInstancesToPrometheusGlobalViewRequest) (response *AppendInstancesToPrometheusGlobalViewResponse, err error) {
	response = CreateAppendInstancesToPrometheusGlobalViewResponse()
	err = client.DoAction(request, response)
	return
}

// AppendInstancesToPrometheusGlobalViewWithChan invokes the arms.AppendInstancesToPrometheusGlobalView API asynchronously
func (client *Client) AppendInstancesToPrometheusGlobalViewWithChan(request *AppendInstancesToPrometheusGlobalViewRequest) (<-chan *AppendInstancesToPrometheusGlobalViewResponse, <-chan error) {
	responseChan := make(chan *AppendInstancesToPrometheusGlobalViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AppendInstancesToPrometheusGlobalView(request)
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

// AppendInstancesToPrometheusGlobalViewWithCallback invokes the arms.AppendInstancesToPrometheusGlobalView API asynchronously
func (client *Client) AppendInstancesToPrometheusGlobalViewWithCallback(request *AppendInstancesToPrometheusGlobalViewRequest, callback func(response *AppendInstancesToPrometheusGlobalViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AppendInstancesToPrometheusGlobalViewResponse
		var err error
		defer close(result)
		response, err = client.AppendInstancesToPrometheusGlobalView(request)
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

// AppendInstancesToPrometheusGlobalViewRequest is the request struct for api AppendInstancesToPrometheusGlobalView
type AppendInstancesToPrometheusGlobalViewRequest struct {
	*requests.RpcRequest
	GlobalViewClusterId string `position:"Query" name:"GlobalViewClusterId"`
	GroupName           string `position:"Query" name:"GroupName"`
	Clusters            string `position:"Query" name:"Clusters"`
}

// AppendInstancesToPrometheusGlobalViewResponse is the response struct for api AppendInstancesToPrometheusGlobalView
type AppendInstancesToPrometheusGlobalViewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAppendInstancesToPrometheusGlobalViewRequest creates a request to invoke AppendInstancesToPrometheusGlobalView API
func CreateAppendInstancesToPrometheusGlobalViewRequest() (request *AppendInstancesToPrometheusGlobalViewRequest) {
	request = &AppendInstancesToPrometheusGlobalViewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "AppendInstancesToPrometheusGlobalView", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAppendInstancesToPrometheusGlobalViewResponse creates a response to parse from AppendInstancesToPrometheusGlobalView response
func CreateAppendInstancesToPrometheusGlobalViewResponse() (response *AppendInstancesToPrometheusGlobalViewResponse) {
	response = &AppendInstancesToPrometheusGlobalViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
