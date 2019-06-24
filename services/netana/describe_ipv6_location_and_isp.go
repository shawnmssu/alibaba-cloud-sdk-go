package netana

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

// DescribeIpv6LocationAndIsp invokes the netana.DescribeIpv6LocationAndIsp API synchronously
// api document: https://help.aliyun.com/api/netana/describeipv6locationandisp.html
func (client *Client) DescribeIpv6LocationAndIsp(request *DescribeIpv6LocationAndIspRequest) (response *DescribeIpv6LocationAndIspResponse, err error) {
	response = CreateDescribeIpv6LocationAndIspResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpv6LocationAndIspWithChan invokes the netana.DescribeIpv6LocationAndIsp API asynchronously
// api document: https://help.aliyun.com/api/netana/describeipv6locationandisp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIpv6LocationAndIspWithChan(request *DescribeIpv6LocationAndIspRequest) (<-chan *DescribeIpv6LocationAndIspResponse, <-chan error) {
	responseChan := make(chan *DescribeIpv6LocationAndIspResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpv6LocationAndIsp(request)
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

// DescribeIpv6LocationAndIspWithCallback invokes the netana.DescribeIpv6LocationAndIsp API asynchronously
// api document: https://help.aliyun.com/api/netana/describeipv6locationandisp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeIpv6LocationAndIspWithCallback(request *DescribeIpv6LocationAndIspRequest, callback func(response *DescribeIpv6LocationAndIspResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpv6LocationAndIspResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpv6LocationAndIsp(request)
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

// DescribeIpv6LocationAndIspRequest is the request struct for api DescribeIpv6LocationAndIsp
type DescribeIpv6LocationAndIspRequest struct {
	*requests.RpcRequest
	IpAddress            string           `position:"Query" name:"IpAddress"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
}

// DescribeIpv6LocationAndIspResponse is the response struct for api DescribeIpv6LocationAndIsp
type DescribeIpv6LocationAndIspResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	IpAddress   string `json:"IpAddress" xml:"IpAddress"`
	Country     string `json:"Country" xml:"Country"`
	City        string `json:"City" xml:"City"`
	ISP         string `json:"ISP" xml:"ISP"`
	CountryCode string `json:"CountryCode" xml:"CountryCode"`
	CityCode    string `json:"CityCode" xml:"CityCode"`
	ISPCode     string `json:"ISPCode" xml:"ISPCode"`
}

// CreateDescribeIpv6LocationAndIspRequest creates a request to invoke DescribeIpv6LocationAndIsp API
func CreateDescribeIpv6LocationAndIspRequest() (request *DescribeIpv6LocationAndIspRequest) {
	request = &DescribeIpv6LocationAndIspRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Netana", "2018-10-18", "DescribeIpv6LocationAndIsp", "Netana", "openAPI")
	return
}

// CreateDescribeIpv6LocationAndIspResponse creates a response to parse from DescribeIpv6LocationAndIsp response
func CreateDescribeIpv6LocationAndIspResponse() (response *DescribeIpv6LocationAndIspResponse) {
	response = &DescribeIpv6LocationAndIspResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}