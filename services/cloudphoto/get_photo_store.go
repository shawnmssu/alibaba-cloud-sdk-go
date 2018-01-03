package cloudphoto

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

func (client *Client) GetPhotoStore(request *GetPhotoStoreRequest) (response *GetPhotoStoreResponse, err error) {
	response = CreateGetPhotoStoreResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetPhotoStoreWithChan(request *GetPhotoStoreRequest) (<-chan *GetPhotoStoreResponse, <-chan error) {
	responseChan := make(chan *GetPhotoStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPhotoStore(request)
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

func (client *Client) GetPhotoStoreWithCallback(request *GetPhotoStoreRequest, callback func(response *GetPhotoStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPhotoStoreResponse
		var err error
		defer close(result)
		response, err = client.GetPhotoStore(request)
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

type GetPhotoStoreRequest struct {
	*requests.RpcRequest
	StoreName string `position:"Query" name:"StoreName"`
}

type GetPhotoStoreResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Action     string `json:"Action" xml:"Action"`
	PhotoStore struct {
		Id               int    `json:"Id" xml:"Id"`
		Name             string `json:"Name" xml:"Name"`
		Remark           string `json:"Remark" xml:"Remark"`
		AutoCleanEnabled bool   `json:"AutoCleanEnabled" xml:"AutoCleanEnabled"`
		AutoCleanDays    int    `json:"AutoCleanDays" xml:"AutoCleanDays"`
		DefaultQuota     int    `json:"DefaultQuota" xml:"DefaultQuota"`
		Ctime            int    `json:"Ctime" xml:"Ctime"`
		Mtime            int    `json:"Mtime" xml:"Mtime"`
		Buckets          []struct {
			Name   string `json:"Name" xml:"Name"`
			Region string `json:"Region" xml:"Region"`
			State  string `json:"State" xml:"State"`
			Acl    string `json:"Acl" xml:"Acl"`
		} `json:"Buckets" xml:"Buckets"`
	} `json:"PhotoStore" xml:"PhotoStore"`
}

func CreateGetPhotoStoreRequest() (request *GetPhotoStoreRequest) {
	request = &GetPhotoStoreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotoStore", "", "")
	return
}

func CreateGetPhotoStoreResponse() (response *GetPhotoStoreResponse) {
	response = &GetPhotoStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}