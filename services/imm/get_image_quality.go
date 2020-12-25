package imm

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

// GetImageQuality invokes the imm.GetImageQuality API synchronously
func (client *Client) GetImageQuality(request *GetImageQualityRequest) (response *GetImageQualityResponse, err error) {
	response = CreateGetImageQualityResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageQualityWithChan invokes the imm.GetImageQuality API asynchronously
func (client *Client) GetImageQualityWithChan(request *GetImageQualityRequest) (<-chan *GetImageQualityResponse, <-chan error) {
	responseChan := make(chan *GetImageQualityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImageQuality(request)
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

// GetImageQualityWithCallback invokes the imm.GetImageQuality API asynchronously
func (client *Client) GetImageQualityWithCallback(request *GetImageQualityRequest, callback func(response *GetImageQualityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageQualityResponse
		var err error
		defer close(result)
		response, err = client.GetImageQuality(request)
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

// GetImageQualityRequest is the request struct for api GetImageQuality
type GetImageQualityRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	ImageUri string `position:"Query" name:"ImageUri"`
}

// GetImageQualityResponse is the response struct for api GetImageQuality
type GetImageQualityResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ImageUri     string       `json:"ImageUri" xml:"ImageUri"`
	ImageQuality ImageQuality `json:"ImageQuality" xml:"ImageQuality"`
}

// CreateGetImageQualityRequest creates a request to invoke GetImageQuality API
func CreateGetImageQualityRequest() (request *GetImageQualityRequest) {
	request = &GetImageQualityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetImageQuality", "", "")
	request.Method = requests.POST
	return
}

// CreateGetImageQualityResponse creates a response to parse from GetImageQuality response
func CreateGetImageQualityResponse() (response *GetImageQualityResponse) {
	response = &GetImageQualityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
