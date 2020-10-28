package sgw

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

// FileShare is a nested struct in sgw response
type FileShare struct {
	Name                   string `json:"Name" xml:"Name"`
	DiskId                 string `json:"DiskId" xml:"DiskId"`
	DiskType               string `json:"DiskType" xml:"DiskType"`
	Protocol               string `json:"Protocol" xml:"Protocol"`
	Size                   int64  `json:"Size" xml:"Size"`
	Enabled                bool   `json:"Enabled" xml:"Enabled"`
	State                  string `json:"State" xml:"State"`
	TotalUpload            int64  `json:"TotalUpload" xml:"TotalUpload"`
	TotalDownload          int64  `json:"TotalDownload" xml:"TotalDownload"`
	OssBucketName          string `json:"OssBucketName" xml:"OssBucketName"`
	OssEndpoint            string `json:"OssEndpoint" xml:"OssEndpoint"`
	OssBucketSsl           bool   `json:"OssBucketSsl" xml:"OssBucketSsl"`
	LocalPath              string `json:"LocalPath" xml:"LocalPath"`
	CacheMode              string `json:"CacheMode" xml:"CacheMode"`
	Address                string `json:"Address" xml:"Address"`
	SerialNumber           string `json:"SerialNumber" xml:"SerialNumber"`
	IndexId                string `json:"IndexId" xml:"IndexId"`
	RemoteSync             bool   `json:"RemoteSync" xml:"RemoteSync"`
	PollingInterval        int    `json:"PollingInterval" xml:"PollingInterval"`
	IgnoreDelete           bool   `json:"IgnoreDelete" xml:"IgnoreDelete"`
	FeLimit                int    `json:"FeLimit" xml:"FeLimit"`
	BeLimit                int    `json:"BeLimit" xml:"BeLimit"`
	InPlace                bool   `json:"InPlace" xml:"InPlace"`
	Browsable              bool   `json:"Browsable" xml:"Browsable"`
	Squash                 string `json:"Squash" xml:"Squash"`
	RwUserList             string `json:"RwUserList" xml:"RwUserList"`
	RoUserList             string `json:"RoUserList" xml:"RoUserList"`
	RwClientList           string `json:"RwClientList" xml:"RwClientList"`
	RoClientList           string `json:"RoClientList" xml:"RoClientList"`
	OssUsed                int64  `json:"OssUsed" xml:"OssUsed"`
	Used                   int64  `json:"Used" xml:"Used"`
	InRate                 int64  `json:"InRate" xml:"InRate"`
	OutRate                int64  `json:"OutRate" xml:"OutRate"`
	LagPeriod              int64  `json:"LagPeriod" xml:"LagPeriod"`
	DirectIO               bool   `json:"DirectIO" xml:"DirectIO"`
	NfsFullPath            string `json:"NfsFullPath" xml:"NfsFullPath"`
	FileNumLimit           int64  `json:"FileNumLimit" xml:"FileNumLimit"`
	FsSizeLimit            int64  `json:"FsSizeLimit" xml:"FsSizeLimit"`
	ServerSideEncryption   bool   `json:"ServerSideEncryption" xml:"ServerSideEncryption"`
	ServerSideCmk          string `json:"ServerSideCmk" xml:"ServerSideCmk"`
	ClientSideEncryption   bool   `json:"ClientSideEncryption" xml:"ClientSideEncryption"`
	ClientSideCmk          string `json:"ClientSideCmk" xml:"ClientSideCmk"`
	KmsRotatePeriod        string `json:"KmsRotatePeriod" xml:"KmsRotatePeriod"`
	OssHealth              string `json:"OssHealth" xml:"OssHealth"`
	PathPrefix             string `json:"PathPrefix" xml:"PathPrefix"`
	FastReclaim            bool   `json:"FastReclaim" xml:"FastReclaim"`
	SupportArchive         bool   `json:"SupportArchive" xml:"SupportArchive"`
	RemainingMetaSpace     int64  `json:"RemainingMetaSpace" xml:"RemainingMetaSpace"`
	MnsHealth              string `json:"MnsHealth" xml:"MnsHealth"`
	ExpressSyncId          string `json:"ExpressSyncId" xml:"ExpressSyncId"`
	WindowsAcl             bool   `json:"WindowsAcl" xml:"WindowsAcl"`
	AccessBasedEnumeration bool   `json:"AccessBasedEnumeration" xml:"AccessBasedEnumeration"`
	NfsV4Optimization      bool   `json:"NfsV4Optimization" xml:"NfsV4Optimization"`
	BucketsStub            bool   `json:"BucketsStub" xml:"BucketsStub"`
	BucketInfos            string `json:"BucketInfos" xml:"BucketInfos"`
	ObsoleteBuckets        string `json:"ObsoleteBuckets" xml:"ObsoleteBuckets"`
	TransferAcceleration   bool   `json:"TransferAcceleration" xml:"TransferAcceleration"`
	DownloadLimit          int    `json:"DownloadLimit" xml:"DownloadLimit"`
	RemoteSyncDownload     bool   `json:"RemoteSyncDownload" xml:"RemoteSyncDownload"`
	PartialSyncPaths       string `json:"PartialSyncPaths" xml:"PartialSyncPaths"`
	SyncProgress           int    `json:"SyncProgress" xml:"SyncProgress"`
	UploadQueue            int64  `json:"UploadQueue" xml:"UploadQueue"`
	DownloadQueue          int64  `json:"DownloadQueue" xml:"DownloadQueue"`
	DownloadRate           int64  `json:"DownloadRate" xml:"DownloadRate"`
	ActiveMessages         int64  `json:"ActiveMessages" xml:"ActiveMessages"`
}