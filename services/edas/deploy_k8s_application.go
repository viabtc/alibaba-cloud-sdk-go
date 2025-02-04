package edas

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

// DeployK8sApplication invokes the edas.DeployK8sApplication API synchronously
func (client *Client) DeployK8sApplication(request *DeployK8sApplicationRequest) (response *DeployK8sApplicationResponse, err error) {
	response = CreateDeployK8sApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// DeployK8sApplicationWithChan invokes the edas.DeployK8sApplication API asynchronously
func (client *Client) DeployK8sApplicationWithChan(request *DeployK8sApplicationRequest) (<-chan *DeployK8sApplicationResponse, <-chan error) {
	responseChan := make(chan *DeployK8sApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployK8sApplication(request)
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

// DeployK8sApplicationWithCallback invokes the edas.DeployK8sApplication API asynchronously
func (client *Client) DeployK8sApplicationWithCallback(request *DeployK8sApplicationRequest, callback func(response *DeployK8sApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployK8sApplicationResponse
		var err error
		defer close(result)
		response, err = client.DeployK8sApplication(request)
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

// DeployK8sApplicationRequest is the request struct for api DeployK8sApplication
type DeployK8sApplicationRequest struct {
	*requests.RoaRequest
	NasId                  string           `position:"Query" name:"NasId"`
	PackageVersionId       string           `position:"Query" name:"PackageVersionId"`
	BatchWaitTime          requests.Integer `position:"Query" name:"BatchWaitTime"`
	Envs                   string           `position:"Query" name:"Envs"`
	Annotations            string           `position:"Query" name:"Annotations"`
	CpuLimit               requests.Integer `position:"Query" name:"CpuLimit"`
	StorageType            string           `position:"Query" name:"StorageType"`
	ConfigMountDescs       string           `position:"Query" name:"ConfigMountDescs"`
	MemoryLimit            requests.Integer `position:"Query" name:"MemoryLimit"`
	ImageTag               string           `position:"Query" name:"ImageTag"`
	DeployAcrossZones      string           `position:"Query" name:"DeployAcrossZones"`
	DeployAcrossNodes      string           `position:"Query" name:"DeployAcrossNodes"`
	MemoryRequest          requests.Integer `position:"Query" name:"MemoryRequest"`
	Image                  string           `position:"Query" name:"Image"`
	PreStop                string           `position:"Query" name:"PreStop"`
	BuildPackId            string           `position:"Query" name:"BuildPackId"`
	EnableEmptyPushReject  requests.Boolean `position:"Query" name:"EnableEmptyPushReject"`
	LocalVolume            string           `position:"Query" name:"LocalVolume"`
	UpdateStrategy         string           `position:"Query" name:"UpdateStrategy"`
	Labels                 string           `position:"Query" name:"Labels"`
	UseBodyEncoding        requests.Boolean `position:"Query" name:"UseBodyEncoding"`
	ChangeOrderDesc        string           `position:"Query" name:"ChangeOrderDesc"`
	LosslessRuleFuncType   requests.Integer `position:"Query" name:"LosslessRuleFuncType"`
	EmptyDirs              string           `position:"Query" name:"EmptyDirs"`
	McpuLimit              requests.Integer `position:"Query" name:"McpuLimit"`
	LosslessRuleRelated    requests.Boolean `position:"Query" name:"LosslessRuleRelated"`
	RuntimeClassName       string           `position:"Query" name:"RuntimeClassName"`
	TrafficControlStrategy string           `position:"Query" name:"TrafficControlStrategy"`
	PostStart              string           `position:"Query" name:"PostStart"`
	CustomAffinity         string           `position:"Query" name:"CustomAffinity"`
	EnableLosslessRule     requests.Boolean `position:"Query" name:"EnableLosslessRule"`
	LosslessRuleWarmupTime requests.Integer `position:"Query" name:"LosslessRuleWarmupTime"`
	WebContainer           string           `position:"Query" name:"WebContainer"`
	EnableAhas             requests.Boolean `position:"Query" name:"EnableAhas"`
	SlsConfigs             string           `position:"Query" name:"SlsConfigs"`
	Readiness              string           `position:"Query" name:"Readiness"`
	Liveness               string           `position:"Query" name:"Liveness"`
	PackageVersion         string           `position:"Query" name:"PackageVersion"`
	EnvFroms               string           `position:"Query" name:"EnvFroms"`
	EdasContainerVersion   string           `position:"Query" name:"EdasContainerVersion"`
	PackageUrl             string           `position:"Query" name:"PackageUrl"`
	LosslessRuleDelayTime  requests.Integer `position:"Query" name:"LosslessRuleDelayTime"`
	MountDescs             string           `position:"Query" name:"MountDescs"`
	Replicas               requests.Integer `position:"Query" name:"Replicas"`
	CustomTolerations      string           `position:"Query" name:"CustomTolerations"`
	CpuRequest             requests.Integer `position:"Query" name:"CpuRequest"`
	WebContainerConfig     string           `position:"Query" name:"WebContainerConfig"`
	Command                string           `position:"Query" name:"Command"`
	Args                   string           `position:"Query" name:"Args"`
	JDK                    string           `position:"Query" name:"JDK"`
	UriEncoding            string           `position:"Query" name:"UriEncoding"`
	AppId                  string           `position:"Query" name:"AppId"`
	BatchTimeout           requests.Integer `position:"Query" name:"BatchTimeout"`
	PvcMountDescs          string           `position:"Query" name:"PvcMountDescs"`
	McpuRequest            requests.Integer `position:"Query" name:"McpuRequest"`
	VolumesStr             string           `position:"Query" name:"VolumesStr"`
	LosslessRuleAligned    requests.Boolean `position:"Query" name:"LosslessRuleAligned"`
	JavaStartUpConfig      string           `position:"Query" name:"JavaStartUpConfig"`
}

// DeployK8sApplicationResponse is the response struct for api DeployK8sApplication
type DeployK8sApplicationResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeployK8sApplicationRequest creates a request to invoke DeployK8sApplication API
func CreateDeployK8sApplicationRequest() (request *DeployK8sApplicationRequest) {
	request = &DeployK8sApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeployK8sApplication", "/pop/v5/k8s/acs/k8s_apps", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeployK8sApplicationResponse creates a response to parse from DeployK8sApplication response
func CreateDeployK8sApplicationResponse() (response *DeployK8sApplicationResponse) {
	response = &DeployK8sApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
