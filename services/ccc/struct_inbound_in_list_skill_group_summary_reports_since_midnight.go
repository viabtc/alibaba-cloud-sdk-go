package ccc

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

// InboundInListSkillGroupSummaryReportsSinceMidnight is a nested struct in ccc response
type InboundInListSkillGroupSummaryReportsSinceMidnight struct {
	CallsOffered                              int64   `json:"CallsOffered" xml:"CallsOffered"`
	CallsHandled                              int64   `json:"CallsHandled" xml:"CallsHandled"`
	HandleRate                                float64 `json:"HandleRate" xml:"HandleRate"`
	TotalRingTime                             int64   `json:"TotalRingTime" xml:"TotalRingTime"`
	MaxRingTime                               int64   `json:"MaxRingTime" xml:"MaxRingTime"`
	AverageRingTime                           int64   `json:"AverageRingTime" xml:"AverageRingTime"`
	ServiceLevel20                            float64 `json:"ServiceLevel20" xml:"ServiceLevel20"`
	TotalTalkTime                             int64   `json:"TotalTalkTime" xml:"TotalTalkTime"`
	MaxTalkTime                               string  `json:"MaxTalkTime" xml:"MaxTalkTime"`
	AverageTalkTime                           int64   `json:"AverageTalkTime" xml:"AverageTalkTime"`
	TotalWorkTime                             int64   `json:"TotalWorkTime" xml:"TotalWorkTime"`
	MaxWorkTime                               int64   `json:"MaxWorkTime" xml:"MaxWorkTime"`
	AverageWorkTime                           int64   `json:"AverageWorkTime" xml:"AverageWorkTime"`
	SatisfactionIndex                         float64 `json:"SatisfactionIndex" xml:"SatisfactionIndex"`
	SatisfactionSurveysOffered                int64   `json:"SatisfactionSurveysOffered" xml:"SatisfactionSurveysOffered"`
	SatisfactionSurveysResponded              int64   `json:"SatisfactionSurveysResponded" xml:"SatisfactionSurveysResponded"`
	InComingQueueOfQueueCount                 int64   `json:"InComingQueueOfQueueCount" xml:"InComingQueueOfQueueCount"`
	AnsweredByAgentOfQueueCount               int64   `json:"AnsweredByAgentOfQueueCount" xml:"AnsweredByAgentOfQueueCount"`
	GiveUpByAgentOfQueueCount                 int64   `json:"GiveUpByAgentOfQueueCount" xml:"GiveUpByAgentOfQueueCount"`
	AbandonedInQueueOfQueueCount              int64   `json:"AbandonedInQueueOfQueueCount" xml:"AbandonedInQueueOfQueueCount"`
	OverFlowInQueueOfQueueCount               int64   `json:"OverFlowInQueueOfQueueCount" xml:"OverFlowInQueueOfQueueCount"`
	QueueWaitTimeDuration                     int64   `json:"QueueWaitTimeDuration" xml:"QueueWaitTimeDuration"`
	AnsweredByAgentOfQueueWaitTimeDuration    int64   `json:"AnsweredByAgentOfQueueWaitTimeDuration" xml:"AnsweredByAgentOfQueueWaitTimeDuration"`
	QueueMaxWaitTimeDuration                  int64   `json:"QueueMaxWaitTimeDuration" xml:"QueueMaxWaitTimeDuration"`
	AnsweredByAgentOfQueueMaxWaitTimeDuration int64   `json:"AnsweredByAgentOfQueueMaxWaitTimeDuration" xml:"AnsweredByAgentOfQueueMaxWaitTimeDuration"`
}
