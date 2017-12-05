package emr

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

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
	response = CreateDescribeClusterResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeClusterWithChan(request *DescribeClusterRequest) (<-chan *DescribeClusterResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCluster(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeClusterWithCallback(request *DescribeClusterRequest, callback func(response *DescribeClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterResponse
		var err error
		defer close(result)
		response, err = client.DescribeCluster(request)
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

type DescribeClusterRequest struct {
	*requests.RpcRequest
	Id              string `position:"Query" name:"Id"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type DescribeClusterResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	ClusterInfo struct {
		Id                     string `json:"Id" xml:"Id"`
		RegionId               string `json:"RegionId" xml:"RegionId"`
		ZoneId                 string `json:"ZoneId" xml:"ZoneId"`
		Name                   string `json:"Name" xml:"Name"`
		CreateType             string `json:"CreateType" xml:"CreateType"`
		StartTime              int64  `json:"StartTime" xml:"StartTime"`
		StopTime               int64  `json:"StopTime" xml:"StopTime"`
		LogEnable              bool   `json:"LogEnable" xml:"LogEnable"`
		LogPath                string `json:"LogPath" xml:"LogPath"`
		UserId                 string `json:"UserId" xml:"UserId"`
		Status                 string `json:"Status" xml:"Status"`
		HighAvailabilityEnable bool   `json:"HighAvailabilityEnable" xml:"HighAvailabilityEnable"`
		ChargeType             string `json:"ChargeType" xml:"ChargeType"`
		ExpiredTime            int64  `json:"ExpiredTime" xml:"ExpiredTime"`
		Period                 int    `json:"Period" xml:"Period"`
		RunningTime            int    `json:"RunningTime" xml:"RunningTime"`
		MasterNodeTotal        int    `json:"MasterNodeTotal" xml:"MasterNodeTotal"`
		MasterNodeInService    int    `json:"MasterNodeInService" xml:"MasterNodeInService"`
		CoreNodeTotal          int    `json:"CoreNodeTotal" xml:"CoreNodeTotal"`
		CoreNodeInService      int    `json:"CoreNodeInService" xml:"CoreNodeInService"`
		TaskNodeTotal          int    `json:"TaskNodeTotal" xml:"TaskNodeTotal"`
		TaskNodeInService      int    `json:"TaskNodeInService" xml:"TaskNodeInService"`
		ShowSoftwareInterface  bool   `json:"ShowSoftwareInterface" xml:"ShowSoftwareInterface"`
		CreateResource         string `json:"CreateResource" xml:"CreateResource"`
		VpcId                  string `json:"VpcId" xml:"VpcId"`
		VSwitchId              string `json:"VSwitchId" xml:"VSwitchId"`
		NetType                string `json:"NetType" xml:"NetType"`
		UserDefinedEmrEcsRole  string `json:"UserDefinedEmrEcsRole" xml:"UserDefinedEmrEcsRole"`
		IoOptimized            bool   `json:"IoOptimized" xml:"IoOptimized"`
		InstanceGeneration     string `json:"InstanceGeneration" xml:"InstanceGeneration"`
		ImageId                string `json:"ImageId" xml:"ImageId"`
		SecurityGroupId        string `json:"SecurityGroupId" xml:"SecurityGroupId"`
		SecurityGroupName      string `json:"SecurityGroupName" xml:"SecurityGroupName"`
		BootstrapFailed        bool   `json:"BootstrapFailed" xml:"BootstrapFailed"`
		Configurations         string `json:"Configurations" xml:"Configurations"`
		EasEnable              bool   `json:"EasEnable" xml:"EasEnable"`
		FailReason             struct {
			ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
			ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
			RequestId string `json:"RequestId" xml:"RequestId"`
		} `json:"FailReason" xml:"FailReason"`
		SoftwareInfo struct {
			EmrVer      string `json:"EmrVer" xml:"EmrVer"`
			ClusterType string `json:"ClusterType" xml:"ClusterType"`
			Softwares   []struct {
				DisplayName string `json:"DisplayName" xml:"DisplayName"`
				Name        string `json:"Name" xml:"Name"`
				OnlyDisplay bool   `json:"OnlyDisplay" xml:"OnlyDisplay"`
				StartTpe    int    `json:"StartTpe" xml:"StartTpe"`
				Version     string `json:"Version" xml:"Version"`
			} `json:"Softwares" xml:"Softwares"`
		} `json:"SoftwareInfo" xml:"SoftwareInfo"`
		EcsOrderInfoList []struct {
			NodeType       string `json:"NodeType" xml:"NodeType"`
			InstanceType   string `json:"InstanceType" xml:"InstanceType"`
			CpuCore        int    `json:"CpuCore" xml:"CpuCore"`
			MemoryCapacity int    `json:"MemoryCapacity" xml:"MemoryCapacity"`
			DiskType       string `json:"DiskType" xml:"DiskType"`
			DiskCapacity   int    `json:"DiskCapacity" xml:"DiskCapacity"`
			DiskCount      int    `json:"DiskCount" xml:"DiskCount"`
			BandWidth      string `json:"BandWidth" xml:"BandWidth"`
			Nodes          []struct {
				ZoneId         string `json:"ZoneId" xml:"ZoneId"`
				InstanceId     string `json:"InstanceId" xml:"InstanceId"`
				Status         string `json:"Status" xml:"Status"`
				PubIp          string `json:"PubIp" xml:"PubIp"`
				InnerIp        string `json:"InnerIp" xml:"InnerIp"`
				ExpiredTime    string `json:"ExpiredTime" xml:"ExpiredTime"`
				EmrExpiredTime string `json:"EmrExpiredTime" xml:"EmrExpiredTime"`
				DaemonInfos    []struct {
					Name string `json:"Name" xml:"Name"`
				} `json:"DaemonInfos" xml:"DaemonInfos"`
				DiskInfos []struct {
					Device   string `json:"Device" xml:"Device"`
					DiskName string `json:"DiskName" xml:"DiskName"`
					DiskId   string `json:"DiskId" xml:"DiskId"`
					Type     string `json:"Type" xml:"Type"`
					Size     int    `json:"Size" xml:"Size"`
				} `json:"DiskInfos" xml:"DiskInfos"`
			} `json:"Nodes" xml:"Nodes"`
		} `json:"EcsOrderInfoList" xml:"EcsOrderInfoList"`
		BootstrapActionList []struct {
			Name string `json:"Name" xml:"Name"`
			Path string `json:"Path" xml:"Path"`
			Arg  string `json:"Arg" xml:"Arg"`
		} `json:"BootstrapActionList" xml:"BootstrapActionList"`
	} `json:"ClusterInfo" xml:"ClusterInfo"`
}

func CreateDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeCluster", "", "")
	return
}

func CreateDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}