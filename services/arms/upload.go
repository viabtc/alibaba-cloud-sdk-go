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

// Upload invokes the arms.Upload API synchronously
func (client *Client) Upload(request *UploadRequest) (response *UploadResponse, err error) {
	response = CreateUploadResponse()
	err = client.DoAction(request, response)
	return
}

// UploadWithChan invokes the arms.Upload API asynchronously
func (client *Client) UploadWithChan(request *UploadRequest) (<-chan *UploadResponse, <-chan error) {
	responseChan := make(chan *UploadResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Upload(request)
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

// UploadWithCallback invokes the arms.Upload API asynchronously
func (client *Client) UploadWithCallback(request *UploadRequest, callback func(response *UploadResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadResponse
		var err error
		defer close(result)
		response, err = client.Upload(request)
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

// UploadRequest is the request struct for api Upload
type UploadRequest struct {
	*requests.RpcRequest
	FileName string `position:"Query" name:"FileName"`
	File     string `position:"Query" name:"File"`
	Pid      string `position:"Query" name:"Pid"`
	Version  string `position:"Query" name:"Version"`
}

// UploadResponse is the response struct for api Upload
type UploadResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Fid        string `json:"Fid" xml:"Fid"`
	FileName   string `json:"FileName" xml:"FileName"`
	Version    string `json:"Version" xml:"Version"`
	UploadTime string `json:"UploadTime" xml:"UploadTime"`
}

// CreateUploadRequest creates a request to invoke Upload API
func CreateUploadRequest() (request *UploadRequest) {
	request = &UploadRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "Upload", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUploadResponse creates a response to parse from Upload response
func CreateUploadResponse() (response *UploadResponse) {
	response = &UploadResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
