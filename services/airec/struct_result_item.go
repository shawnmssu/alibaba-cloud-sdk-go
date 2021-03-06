package airec

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

// ResultItem is a nested struct in airec response
type ResultItem struct {
	Name               string                 `json:"Name" xml:"Name"`
	ErrorLevel         string                 `json:"ErrorLevel" xml:"ErrorLevel"`
	DefaultDisplay     bool                   `json:"DefaultDisplay" xml:"DefaultDisplay"`
	ItemType           string                 `json:"ItemType" xml:"ItemType"`
	SampleDisplay      bool                   `json:"SampleDisplay" xml:"SampleDisplay"`
	InstanceId         string                 `json:"InstanceId" xml:"InstanceId"`
	TotalProgress      int                    `json:"TotalProgress" xml:"TotalProgress"`
	SelectValue        string                 `json:"SelectValue" xml:"SelectValue"`
	SceneId            string                 `json:"SceneId" xml:"SceneId"`
	TraceId            string                 `json:"TraceId" xml:"TraceId"`
	SelectionOperation string                 `json:"SelectionOperation" xml:"SelectionOperation"`
	MetricData         map[string]interface{} `json:"MetricData" xml:"MetricData"`
	MatchInfo          string                 `json:"MatchInfo" xml:"MatchInfo"`
	Type               string                 `json:"Type" xml:"Type"`
	ErrorCount         int                    `json:"ErrorCount" xml:"ErrorCount"`
	MetricType         string                 `json:"MetricType" xml:"MetricType"`
	Message            string                 `json:"Message" xml:"Message"`
	Timestamp          string                 `json:"Timestamp" xml:"Timestamp"`
	ErrorType          string                 `json:"ErrorType" xml:"ErrorType"`
	ItemId             string                 `json:"ItemId" xml:"ItemId"`
	VersionId          string                 `json:"VersionId" xml:"VersionId"`
	GmtCreate          int64                  `json:"GmtCreate" xml:"GmtCreate"`
	SelectType         string                 `json:"SelectType" xml:"SelectType"`
	Platform           string                 `json:"Platform" xml:"Platform"`
	ErrorPercent       float64                `json:"ErrorPercent" xml:"ErrorPercent"`
	State              string                 `json:"State" xml:"State"`
	Appkey             string                 `json:"Appkey" xml:"Appkey"`
	GmtModified        int64                  `json:"GmtModified" xml:"GmtModified"`
	TraceInfo          string                 `json:"TraceInfo" xml:"TraceInfo"`
	Total              map[string]interface{} `json:"Total" xml:"Total"`
	Position           int                    `json:"Position" xml:"Position"`
	Weight             float64                `json:"Weight" xml:"Weight"`
	MetricRes          MetricRes              `json:"MetricRes" xml:"MetricRes"`
	Detail             []DetailItem           `json:"Detail" xml:"Detail"`
	DataPoints         []DataPointsItem       `json:"DataPoints" xml:"DataPoints"`
	SubProgressInfos   []SubProgressInfosItem `json:"SubProgressInfos" xml:"SubProgressInfos"`
	HistoryData        []HistoryDataItem      `json:"HistoryData" xml:"HistoryData"`
}
