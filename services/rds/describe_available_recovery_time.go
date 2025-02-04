package rds

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

// DescribeAvailableRecoveryTime invokes the rds.DescribeAvailableRecoveryTime API synchronously
func (client *Client) DescribeAvailableRecoveryTime(request *DescribeAvailableRecoveryTimeRequest) (response *DescribeAvailableRecoveryTimeResponse, err error) {
	response = CreateDescribeAvailableRecoveryTimeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableRecoveryTimeWithChan invokes the rds.DescribeAvailableRecoveryTime API asynchronously
func (client *Client) DescribeAvailableRecoveryTimeWithChan(request *DescribeAvailableRecoveryTimeRequest) (<-chan *DescribeAvailableRecoveryTimeResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableRecoveryTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableRecoveryTime(request)
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

// DescribeAvailableRecoveryTimeWithCallback invokes the rds.DescribeAvailableRecoveryTime API asynchronously
func (client *Client) DescribeAvailableRecoveryTimeWithCallback(request *DescribeAvailableRecoveryTimeRequest, callback func(response *DescribeAvailableRecoveryTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableRecoveryTimeResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableRecoveryTime(request)
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

// DescribeAvailableRecoveryTimeRequest is the request struct for api DescribeAvailableRecoveryTime
type DescribeAvailableRecoveryTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	CrossBackupId        requests.Integer `position:"Query" name:"CrossBackupId"`
}

// DescribeAvailableRecoveryTimeResponse is the response struct for api DescribeAvailableRecoveryTime
type DescribeAvailableRecoveryTimeResponse struct {
	*responses.BaseResponse
	RecoveryEndTime   string `json:"RecoveryEndTime" xml:"RecoveryEndTime"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
	CrossBackupId     int    `json:"CrossBackupId" xml:"CrossBackupId"`
	DBInstanceId      string `json:"DBInstanceId" xml:"DBInstanceId"`
	RecoveryBeginTime string `json:"RecoveryBeginTime" xml:"RecoveryBeginTime"`
	RegionId          string `json:"RegionId" xml:"RegionId"`
}

// CreateDescribeAvailableRecoveryTimeRequest creates a request to invoke DescribeAvailableRecoveryTime API
func CreateDescribeAvailableRecoveryTimeRequest() (request *DescribeAvailableRecoveryTimeRequest) {
	request = &DescribeAvailableRecoveryTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAvailableRecoveryTime", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAvailableRecoveryTimeResponse creates a response to parse from DescribeAvailableRecoveryTime response
func CreateDescribeAvailableRecoveryTimeResponse() (response *DescribeAvailableRecoveryTimeResponse) {
	response = &DescribeAvailableRecoveryTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
