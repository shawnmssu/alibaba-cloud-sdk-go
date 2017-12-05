package itaas

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

func (client *Client) FetchCommand(request *FetchCommandRequest) (response *FetchCommandResponse, err error) {
	response = CreateFetchCommandResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) FetchCommandWithChan(request *FetchCommandRequest) (<-chan *FetchCommandResponse, <-chan error) {
	responseChan := make(chan *FetchCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FetchCommand(request)
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

func (client *Client) FetchCommandWithCallback(request *FetchCommandRequest, callback func(response *FetchCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FetchCommandResponse
		var err error
		defer close(result)
		response, err = client.FetchCommand(request)
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

type FetchCommandRequest struct {
	*requests.RpcRequest
	Uid           string `position:"Query" name:"Uid"`
	AuthType      string `position:"Query" name:"AuthType"`
	Sid           string `position:"Query" name:"Sid"`
	Language      string `position:"Query" name:"Language"`
	PushType      string `position:"Query" name:"PushType"`
	PushId        string `position:"Query" name:"PushId"`
	Cid           string `position:"Query" name:"Cid"`
	Umid          string `position:"Query" name:"Umid"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type FetchCommandResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	PushId       int64  `json:"PushId" xml:"PushId"`
	PushType     int    `json:"PushType" xml:"PushType"`
	CommandType  string `json:"CommandType" xml:"CommandType"`
	Content      string `json:"Content" xml:"Content"`
	FeedbackUrl  string `json:"FeedbackUrl" xml:"FeedbackUrl"`
}

func CreateFetchCommandRequest() (request *FetchCommandRequest) {
	request = &FetchCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "FetchCommand", "", "")
	return
}

func CreateFetchCommandResponse() (response *FetchCommandResponse) {
	response = &FetchCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}