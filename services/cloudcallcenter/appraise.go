package cloudcallcenter

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

// Appraise invokes the cloudcallcenter.Appraise API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/appraise.html
func (client *Client) Appraise(request *AppraiseRequest) (response *AppraiseResponse, err error) {
	response = CreateAppraiseResponse()
	err = client.DoAction(request, response)
	return
}

// AppraiseWithChan invokes the cloudcallcenter.Appraise API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/appraise.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AppraiseWithChan(request *AppraiseRequest) (<-chan *AppraiseResponse, <-chan error) {
	responseChan := make(chan *AppraiseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Appraise(request)
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

// AppraiseWithCallback invokes the cloudcallcenter.Appraise API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/appraise.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AppraiseWithCallback(request *AppraiseRequest, callback func(response *AppraiseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AppraiseResponse
		var err error
		defer close(result)
		response, err = client.Appraise(request)
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

// AppraiseRequest is the request struct for api Appraise
type AppraiseRequest struct {
	*requests.RpcRequest
	Source          string `position:"Query" name:"Source"`
	Type            string `position:"Query" name:"Type"`
	RamId           string `position:"Query" name:"RamId"`
	Acid            string `position:"Query" name:"Acid"`
	PressKey        string `position:"Query" name:"PressKey"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SkillGroupId    string `position:"Query" name:"SkillGroupId"`
	PressKeyMapping string `position:"Query" name:"PressKeyMapping"`
}

// AppraiseResponse is the response struct for api Appraise
type AppraiseResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateAppraiseRequest creates a request to invoke Appraise API
func CreateAppraiseRequest() (request *AppraiseRequest) {
	request = &AppraiseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "Appraise", "", "")
	request.Method = requests.POST
	return
}

// CreateAppraiseResponse creates a response to parse from Appraise response
func CreateAppraiseResponse() (response *AppraiseResponse) {
	response = &AppraiseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}