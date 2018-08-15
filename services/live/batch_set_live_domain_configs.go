package live

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

// BatchSetLiveDomainConfigs invokes the live.BatchSetLiveDomainConfigs API synchronously
// api document: https://help.aliyun.com/api/live/batchsetlivedomainconfigs.html
func (client *Client) BatchSetLiveDomainConfigs(request *BatchSetLiveDomainConfigsRequest) (response *BatchSetLiveDomainConfigsResponse, err error) {
	response = CreateBatchSetLiveDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchSetLiveDomainConfigsWithChan invokes the live.BatchSetLiveDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/live/batchsetlivedomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchSetLiveDomainConfigsWithChan(request *BatchSetLiveDomainConfigsRequest) (<-chan *BatchSetLiveDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *BatchSetLiveDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchSetLiveDomainConfigs(request)
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

// BatchSetLiveDomainConfigsWithCallback invokes the live.BatchSetLiveDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/live/batchsetlivedomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchSetLiveDomainConfigsWithCallback(request *BatchSetLiveDomainConfigsRequest, callback func(response *BatchSetLiveDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchSetLiveDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.BatchSetLiveDomainConfigs(request)
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

// BatchSetLiveDomainConfigsRequest is the request struct for api BatchSetLiveDomainConfigs
type BatchSetLiveDomainConfigsRequest struct {
	*requests.RpcRequest
	Functions     string           `position:"Query" name:"Functions"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchSetLiveDomainConfigsResponse is the response struct for api BatchSetLiveDomainConfigs
type BatchSetLiveDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchSetLiveDomainConfigsRequest creates a request to invoke BatchSetLiveDomainConfigs API
func CreateBatchSetLiveDomainConfigsRequest() (request *BatchSetLiveDomainConfigsRequest) {
	request = &BatchSetLiveDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "BatchSetLiveDomainConfigs", "live", "openAPI")
	return
}

// CreateBatchSetLiveDomainConfigsResponse creates a response to parse from BatchSetLiveDomainConfigs response
func CreateBatchSetLiveDomainConfigsResponse() (response *BatchSetLiveDomainConfigsResponse) {
	response = &BatchSetLiveDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}