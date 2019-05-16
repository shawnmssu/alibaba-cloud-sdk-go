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

package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateProductTopic invokes the iot.UpdateProductTopic API synchronously
// api document: https://help.aliyun.com/api/iot/updateproducttopic.html
func (client *Client) UpdateProductTopic(request *UpdateProductTopicRequest) (response *UpdateProductTopicResponse, err error) {
	response = CreateUpdateProductTopicResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateProductTopicWithChan invokes the iot.UpdateProductTopic API asynchronously
// api document: https://help.aliyun.com/api/iot/updateproducttopic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProductTopicWithChan(request *UpdateProductTopicRequest) (<-chan *UpdateProductTopicResponse, <-chan error) {
	responseChan := make(chan *UpdateProductTopicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateProductTopic(request)
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

// UpdateProductTopicWithCallback invokes the iot.UpdateProductTopic API asynchronously
// api document: https://help.aliyun.com/api/iot/updateproducttopic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProductTopicWithCallback(request *UpdateProductTopicRequest, callback func(response *UpdateProductTopicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateProductTopicResponse
		var err error
		defer close(result)
		response, err = client.UpdateProductTopic(request)
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

// UpdateProductTopicRequest is the request struct for api UpdateProductTopic
type UpdateProductTopicRequest struct {
	*requests.RpcRequest
	AccessKeyId    string `position:"Query" name:"AccessKeyId"`
	IotInstanceId  string `position:"Query" name:"IotInstanceId"`
	Desc           string `position:"Query" name:"Desc"`
	Operation      string `position:"Query" name:"Operation"`
	TopicShortName string `position:"Query" name:"TopicShortName"`
	TopicId        string `position:"Query" name:"TopicId"`
}

// UpdateProductTopicResponse is the response struct for api UpdateProductTopic
type UpdateProductTopicResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateProductTopicRequest creates a request to invoke UpdateProductTopic API
func CreateUpdateProductTopicRequest() (request *UpdateProductTopicRequest) {
	request = &UpdateProductTopicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateProductTopic", "iot", "openAPI")
	return
}

// CreateUpdateProductTopicResponse creates a response to parse from UpdateProductTopic response
func CreateUpdateProductTopicResponse() (response *UpdateProductTopicResponse) {
	response = &UpdateProductTopicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
