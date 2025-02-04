package mse

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

// Gateways is a nested struct in mse response
type Gateways struct {
	Name            string        `json:"Name" xml:"Name"`
	Region          string        `json:"Region" xml:"Region"`
	Vswitch2        string        `json:"Vswitch2" xml:"Vswitch2"`
	ChargeType      string        `json:"ChargeType" xml:"ChargeType"`
	GmtCreate       string        `json:"GmtCreate" xml:"GmtCreate"`
	Tag             string        `json:"Tag" xml:"Tag"`
	InstanceId      string        `json:"InstanceId" xml:"InstanceId"`
	SupportWasm     bool          `json:"SupportWasm" xml:"SupportWasm"`
	Replica         int           `json:"Replica" xml:"Replica"`
	ArmsOn          bool          `json:"ArmsOn" xml:"ArmsOn"`
	Status          int           `json:"Status" xml:"Status"`
	EndDate         string        `json:"EndDate" xml:"EndDate"`
	GatewayType     string        `json:"GatewayType" xml:"GatewayType"`
	GmtModified     string        `json:"GmtModified" xml:"GmtModified"`
	Id              int64         `json:"Id" xml:"Id"`
	CurrentVersion  string        `json:"CurrentVersion" xml:"CurrentVersion"`
	Spec            string        `json:"Spec" xml:"Spec"`
	Upgrade         bool          `json:"Upgrade" xml:"Upgrade"`
	GatewayUniqueId string        `json:"GatewayUniqueId" xml:"GatewayUniqueId"`
	PrimaryUser     string        `json:"PrimaryUser" xml:"PrimaryUser"`
	AhasOn          bool          `json:"AhasOn" xml:"AhasOn"`
	MustUpgrade     bool          `json:"MustUpgrade" xml:"MustUpgrade"`
	AppVersion      string        `json:"AppVersion" xml:"AppVersion"`
	StatusDesc      string        `json:"StatusDesc" xml:"StatusDesc"`
	LatestVersion   string        `json:"LatestVersion" xml:"LatestVersion"`
	InitConfig      InitConfig    `json:"InitConfig" xml:"InitConfig"`
	Slb             []Slb         `json:"Slb" xml:"Slb"`
	InternetSlb     []InternetSlb `json:"InternetSlb" xml:"InternetSlb"`
}
