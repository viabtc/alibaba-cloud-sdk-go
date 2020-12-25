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

// GetOfficeConversionTask invokes the imm.GetOfficeConversionTask API synchronously
func (client *Client) GetOfficeConversionTask(request *GetOfficeConversionTaskRequest) (response *GetOfficeConversionTaskResponse, err error) {
	response = CreateGetOfficeConversionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetOfficeConversionTaskWithChan invokes the imm.GetOfficeConversionTask API asynchronously
func (client *Client) GetOfficeConversionTaskWithChan(request *GetOfficeConversionTaskRequest) (<-chan *GetOfficeConversionTaskResponse, <-chan error) {
	responseChan := make(chan *GetOfficeConversionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOfficeConversionTask(request)
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

// GetOfficeConversionTaskWithCallback invokes the imm.GetOfficeConversionTask API asynchronously
func (client *Client) GetOfficeConversionTaskWithCallback(request *GetOfficeConversionTaskRequest, callback func(response *GetOfficeConversionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOfficeConversionTaskResponse
		var err error
		defer close(result)
		response, err = client.GetOfficeConversionTask(request)
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

// GetOfficeConversionTaskRequest is the request struct for api GetOfficeConversionTask
type GetOfficeConversionTaskRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

// GetOfficeConversionTaskResponse is the response struct for api GetOfficeConversionTask
type GetOfficeConversionTaskResponse struct {
	*responses.BaseResponse
	RequestId       string     `json:"RequestId" xml:"RequestId"`
	TgtType         string     `json:"TgtType" xml:"TgtType"`
	Status          string     `json:"Status" xml:"Status"`
	Percent         int        `json:"Percent" xml:"Percent"`
	PageCount       int        `json:"PageCount" xml:"PageCount"`
	TaskId          string     `json:"TaskId" xml:"TaskId"`
	TgtUri          string     `json:"TgtUri" xml:"TgtUri"`
	ImageSpec       string     `json:"ImageSpec" xml:"ImageSpec"`
	NotifyTopicName string     `json:"NotifyTopicName" xml:"NotifyTopicName"`
	NotifyEndpoint  string     `json:"NotifyEndpoint" xml:"NotifyEndpoint"`
	ExternalID      string     `json:"ExternalID" xml:"ExternalID"`
	CreateTime      string     `json:"CreateTime" xml:"CreateTime"`
	FinishTime      string     `json:"FinishTime" xml:"FinishTime"`
	SrcUri          string     `json:"SrcUri" xml:"SrcUri"`
	FailDetail      FailDetail `json:"FailDetail" xml:"FailDetail"`
}

// CreateGetOfficeConversionTaskRequest creates a request to invoke GetOfficeConversionTask API
func CreateGetOfficeConversionTaskRequest() (request *GetOfficeConversionTaskRequest) {
	request = &GetOfficeConversionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetOfficeConversionTask", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOfficeConversionTaskResponse creates a response to parse from GetOfficeConversionTask response
func CreateGetOfficeConversionTaskResponse() (response *GetOfficeConversionTaskResponse) {
	response = &GetOfficeConversionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
