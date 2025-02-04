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

// Backup is a nested struct in rds response
type Backup struct {
	MetaStatus                string                 `json:"MetaStatus" xml:"MetaStatus"`
	CopyOnlyBackup            string                 `json:"CopyOnlyBackup" xml:"CopyOnlyBackup"`
	HostInstanceID            string                 `json:"HostInstanceID" xml:"HostInstanceID"`
	BackupMode                string                 `json:"BackupMode" xml:"BackupMode"`
	IsAvail                   int                    `json:"IsAvail" xml:"IsAvail"`
	BackupDBNames             string                 `json:"BackupDBNames" xml:"BackupDBNames"`
	StorageClass              string                 `json:"StorageClass" xml:"StorageClass"`
	Encryption                string                 `json:"Encryption" xml:"Encryption"`
	BackupSize                int64                  `json:"BackupSize" xml:"BackupSize"`
	BackupStartTime           string                 `json:"BackupStartTime" xml:"BackupStartTime"`
	Checksum                  string                 `json:"Checksum" xml:"Checksum"`
	ConsistentTime            int64                  `json:"ConsistentTime" xml:"ConsistentTime"`
	DBInstanceId              string                 `json:"DBInstanceId" xml:"DBInstanceId"`
	BackupScale               string                 `json:"BackupScale" xml:"BackupScale"`
	BackupDownloadURL         string                 `json:"BackupDownloadURL" xml:"BackupDownloadURL"`
	SlaveStatus               string                 `json:"SlaveStatus" xml:"SlaveStatus"`
	BackupExtractionStatus    string                 `json:"BackupExtractionStatus" xml:"BackupExtractionStatus"`
	DBInstanceComment         string                 `json:"DBInstanceComment" xml:"DBInstanceComment"`
	BackupIntranetDownloadURL string                 `json:"BackupIntranetDownloadURL" xml:"BackupIntranetDownloadURL"`
	BackupEndTime             string                 `json:"BackupEndTime" xml:"BackupEndTime"`
	BackupMethod              string                 `json:"BackupMethod" xml:"BackupMethod"`
	BackupType                string                 `json:"BackupType" xml:"BackupType"`
	TotalBackupSize           int64                  `json:"TotalBackupSize" xml:"TotalBackupSize"`
	BackupInitiator           string                 `json:"BackupInitiator" xml:"BackupInitiator"`
	BackupStatus              string                 `json:"BackupStatus" xml:"BackupStatus"`
	StoreStatus               string                 `json:"StoreStatus" xml:"StoreStatus"`
	BackupId                  string                 `json:"BackupId" xml:"BackupId"`
	BackupLocation            string                 `json:"BackupLocation" xml:"BackupLocation"`
	BackupDownloadLinkByDB    BackupDownloadLinkByDB `json:"BackupDownloadLinkByDB" xml:"BackupDownloadLinkByDB"`
}
