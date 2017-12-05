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

func (client *Client) ListJobInstanceWorkers(request *ListJobInstanceWorkersRequest) (response *ListJobInstanceWorkersResponse, err error) {
	response = CreateListJobInstanceWorkersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListJobInstanceWorkersWithChan(request *ListJobInstanceWorkersRequest) (<-chan *ListJobInstanceWorkersResponse, <-chan error) {
	responseChan := make(chan *ListJobInstanceWorkersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobInstanceWorkers(request)
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

func (client *Client) ListJobInstanceWorkersWithCallback(request *ListJobInstanceWorkersRequest, callback func(response *ListJobInstanceWorkersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobInstanceWorkersResponse
		var err error
		defer close(result)
		response, err = client.ListJobInstanceWorkers(request)
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

type ListJobInstanceWorkersRequest struct {
	*requests.RpcRequest
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type ListJobInstanceWorkersResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	JobInstanceWorkers []struct {
		ApplicationId string `json:"ApplicationId" xml:"ApplicationId"`
		InstanceInfo  string `json:"InstanceInfo" xml:"InstanceInfo"`
		ContainerInfo string `json:"ContainerInfo" xml:"ContainerInfo"`
	} `json:"JobInstanceWorkers" xml:"JobInstanceWorkers"`
}

func CreateListJobInstanceWorkersRequest() (request *ListJobInstanceWorkersRequest) {
	request = &ListJobInstanceWorkersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListJobInstanceWorkers", "", "")
	return
}

func CreateListJobInstanceWorkersResponse() (response *ListJobInstanceWorkersResponse) {
	response = &ListJobInstanceWorkersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}