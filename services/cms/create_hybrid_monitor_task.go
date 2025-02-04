package cms

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

// CreateHybridMonitorTask invokes the cms.CreateHybridMonitorTask API synchronously
func (client *Client) CreateHybridMonitorTask(request *CreateHybridMonitorTaskRequest) (response *CreateHybridMonitorTaskResponse, err error) {
	response = CreateCreateHybridMonitorTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHybridMonitorTaskWithChan invokes the cms.CreateHybridMonitorTask API asynchronously
func (client *Client) CreateHybridMonitorTaskWithChan(request *CreateHybridMonitorTaskRequest) (<-chan *CreateHybridMonitorTaskResponse, <-chan error) {
	responseChan := make(chan *CreateHybridMonitorTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHybridMonitorTask(request)
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

// CreateHybridMonitorTaskWithCallback invokes the cms.CreateHybridMonitorTask API asynchronously
func (client *Client) CreateHybridMonitorTaskWithCallback(request *CreateHybridMonitorTaskRequest, callback func(response *CreateHybridMonitorTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHybridMonitorTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateHybridMonitorTask(request)
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

// CreateHybridMonitorTaskRequest is the request struct for api CreateHybridMonitorTask
type CreateHybridMonitorTaskRequest struct {
	*requests.RpcRequest
	LogSample             string                                  `position:"Query" name:"LogSample"`
	CollectTargetPath     string                                  `position:"Query" name:"CollectTargetPath"`
	Instances             *[]string                               `position:"Query" name:"Instances"  type:"Repeated"`
	LogSplit              string                                  `position:"Query" name:"LogSplit"`
	Description           string                                  `position:"Query" name:"Description"`
	TaskName              string                                  `position:"Query" name:"TaskName"`
	NetworkType           string                                  `position:"Query" name:"NetworkType"`
	ExtraInfo             string                                  `position:"Query" name:"ExtraInfo"`
	CollectInterval       string                                  `position:"Query" name:"CollectInterval"`
	TargetUserId          string                                  `position:"Query" name:"TargetUserId"`
	CollectTargetType     string                                  `position:"Query" name:"CollectTargetType"`
	AttachLabels          *[]CreateHybridMonitorTaskAttachLabels  `position:"Query" name:"AttachLabels"  type:"Repeated"`
	UploadRegion          string                                  `position:"Query" name:"UploadRegion"`
	CollectTimout         string                                  `position:"Query" name:"CollectTimout"`
	TaskType              string                                  `position:"Query" name:"TaskType"`
	MatchExpressRelation  string                                  `position:"Query" name:"MatchExpressRelation"`
	LogProcess            string                                  `position:"Query" name:"LogProcess"`
	SLSProcess            string                                  `position:"Query" name:"SLSProcess"`
	GroupId               string                                  `position:"Query" name:"GroupId"`
	TargetUserIdList      string                                  `position:"Query" name:"TargetUserIdList"`
	CollectTargetEndpoint string                                  `position:"Query" name:"CollectTargetEndpoint"`
	YARMConfig            string                                  `position:"Query" name:"YARMConfig"`
	LogFilePath           string                                  `position:"Query" name:"LogFilePath"`
	MatchExpress          *[]CreateHybridMonitorTaskMatchExpress  `position:"Query" name:"MatchExpress"  type:"Repeated"`
	Namespace             string                                  `position:"Query" name:"Namespace"`
	SLSProcessConfig      CreateHybridMonitorTaskSLSProcessConfig `position:"Query" name:"SLSProcessConfig"  type:"Struct"`
}

// CreateHybridMonitorTaskAttachLabels is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskAttachLabels struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

// CreateHybridMonitorTaskMatchExpress is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskMatchExpress struct {
	Function string `name:"Function"`
	Name     string `name:"Name"`
	Value    string `name:"Value"`
}

// CreateHybridMonitorTaskSLSProcessConfig is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskSLSProcessConfig struct {
	Filter     CreateHybridMonitorTaskSLSProcessConfigFilter            `name:"Filter" type:"Struct"`
	Express    *[]CreateHybridMonitorTaskSLSProcessConfigExpressItem    `name:"Express" type:"Repeated"`
	GroupBy    *[]CreateHybridMonitorTaskSLSProcessConfigGroupByItem    `name:"GroupBy" type:"Repeated"`
	Statistics *[]CreateHybridMonitorTaskSLSProcessConfigStatisticsItem `name:"Statistics" type:"Repeated"`
}

// CreateHybridMonitorTaskSLSProcessConfigFilter is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskSLSProcessConfigFilter struct {
	Filters  *[]CreateHybridMonitorTaskSLSProcessConfigFilterFiltersItem `name:"Filters" type:"Repeated"`
	Relation string                                                      `name:"Relation"`
}

// CreateHybridMonitorTaskSLSProcessConfigExpressItem is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskSLSProcessConfigExpressItem struct {
	Alias   string `name:"Alias"`
	Express string `name:"Express"`
}

// CreateHybridMonitorTaskSLSProcessConfigGroupByItem is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskSLSProcessConfigGroupByItem struct {
	SLSKeyName string `name:"SLSKeyName"`
	Alias      string `name:"Alias"`
}

// CreateHybridMonitorTaskSLSProcessConfigStatisticsItem is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskSLSProcessConfigStatisticsItem struct {
	SLSKeyName string `name:"SLSKeyName"`
	Function   string `name:"Function"`
	Alias      string `name:"Alias"`
	Parameter2 string `name:"Parameter2"`
	Parameter1 string `name:"Parameter1"`
}

// CreateHybridMonitorTaskSLSProcessConfigFilterFiltersItem is a repeated param struct in CreateHybridMonitorTaskRequest
type CreateHybridMonitorTaskSLSProcessConfigFilterFiltersItem struct {
	SLSKeyName string `name:"SLSKeyName"`
	Value      string `name:"Value"`
	Operator   string `name:"Operator"`
}

// CreateHybridMonitorTaskResponse is the response struct for api CreateHybridMonitorTask
type CreateHybridMonitorTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateCreateHybridMonitorTaskRequest creates a request to invoke CreateHybridMonitorTask API
func CreateCreateHybridMonitorTaskRequest() (request *CreateHybridMonitorTaskRequest) {
	request = &CreateHybridMonitorTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateHybridMonitorTask", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateHybridMonitorTaskResponse creates a response to parse from CreateHybridMonitorTask response
func CreateCreateHybridMonitorTaskResponse() (response *CreateHybridMonitorTaskResponse) {
	response = &CreateHybridMonitorTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
