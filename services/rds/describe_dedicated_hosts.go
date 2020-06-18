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

// DescribeDedicatedHosts invokes the rds.DescribeDedicatedHosts API synchronously
// api document: https://help.aliyun.com/api/rds/describededicatedhosts.html
func (client *Client) DescribeDedicatedHosts(request *DescribeDedicatedHostsRequest) (response *DescribeDedicatedHostsResponse, err error) {
	response = CreateDescribeDedicatedHostsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDedicatedHostsWithChan invokes the rds.DescribeDedicatedHosts API asynchronously
// api document: https://help.aliyun.com/api/rds/describededicatedhosts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostsWithChan(request *DescribeDedicatedHostsRequest) (<-chan *DescribeDedicatedHostsResponse, <-chan error) {
	responseChan := make(chan *DescribeDedicatedHostsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDedicatedHosts(request)
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

// DescribeDedicatedHostsWithCallback invokes the rds.DescribeDedicatedHosts API asynchronously
// api document: https://help.aliyun.com/api/rds/describededicatedhosts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostsWithCallback(request *DescribeDedicatedHostsRequest, callback func(response *DescribeDedicatedHostsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDedicatedHostsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDedicatedHosts(request)
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

// DescribeDedicatedHostsRequest is the request struct for api DescribeDedicatedHosts
type DescribeDedicatedHostsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HostType             string           `position:"Query" name:"HostType"`
	HostStatus           string           `position:"Query" name:"HostStatus"`
	AllocationStatus     string           `position:"Query" name:"AllocationStatus"`
	DedicatedHostGroupId string           `position:"Query" name:"DedicatedHostGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OrderId              requests.Integer `position:"Query" name:"OrderId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeDedicatedHostsResponse is the response struct for api DescribeDedicatedHosts
type DescribeDedicatedHostsResponse struct {
	*responses.BaseResponse
	RequestId            string         `json:"RequestId" xml:"RequestId"`
	DedicatedHostGroupId string         `json:"DedicatedHostGroupId" xml:"DedicatedHostGroupId"`
	DedicatedHosts       DedicatedHosts `json:"DedicatedHosts" xml:"DedicatedHosts"`
}

// CreateDescribeDedicatedHostsRequest creates a request to invoke DescribeDedicatedHosts API
func CreateDescribeDedicatedHostsRequest() (request *DescribeDedicatedHostsRequest) {
	request = &DescribeDedicatedHostsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDedicatedHosts", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDedicatedHostsResponse creates a response to parse from DescribeDedicatedHosts response
func CreateDescribeDedicatedHostsResponse() (response *DescribeDedicatedHostsResponse) {
	response = &DescribeDedicatedHostsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
