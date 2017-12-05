package push

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

func (client *Client) CheckDevices(request *CheckDevicesRequest) (response *CheckDevicesResponse, err error) {
	response = CreateCheckDevicesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CheckDevicesWithChan(request *CheckDevicesRequest) (<-chan *CheckDevicesResponse, <-chan error) {
	responseChan := make(chan *CheckDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDevices(request)
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

func (client *Client) CheckDevicesWithCallback(request *CheckDevicesRequest, callback func(response *CheckDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDevicesResponse
		var err error
		defer close(result)
		response, err = client.CheckDevices(request)
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

type CheckDevicesRequest struct {
	*requests.RpcRequest
	AppKey    string `position:"Query" name:"AppKey"`
	DeviceIds string `position:"Query" name:"DeviceIds"`
}

type CheckDevicesResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DeviceCheckInfos []struct {
		DeviceId  string `json:"DeviceId" xml:"DeviceId"`
		Available bool   `json:"Available" xml:"Available"`
	} `json:"DeviceCheckInfos" xml:"DeviceCheckInfos"`
}

func CreateCheckDevicesRequest() (request *CheckDevicesRequest) {
	request = &CheckDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "CheckDevices", "", "")
	return
}

func CreateCheckDevicesResponse() (response *CheckDevicesResponse) {
	response = &CheckDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}