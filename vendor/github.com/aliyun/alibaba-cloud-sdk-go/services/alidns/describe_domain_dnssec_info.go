package alidns

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

// DescribeDomainDnssecInfo invokes the alidns.DescribeDomainDnssecInfo API synchronously
// api document: https://help.aliyun.com/api/alidns/describedomaindnssecinfo.html
func (client *Client) DescribeDomainDnssecInfo(request *DescribeDomainDnssecInfoRequest) (response *DescribeDomainDnssecInfoResponse, err error) {
	response = CreateDescribeDomainDnssecInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainDnssecInfoWithChan invokes the alidns.DescribeDomainDnssecInfo API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedomaindnssecinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainDnssecInfoWithChan(request *DescribeDomainDnssecInfoRequest) (<-chan *DescribeDomainDnssecInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainDnssecInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainDnssecInfo(request)
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

// DescribeDomainDnssecInfoWithCallback invokes the alidns.DescribeDomainDnssecInfo API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedomaindnssecinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainDnssecInfoWithCallback(request *DescribeDomainDnssecInfoRequest, callback func(response *DescribeDomainDnssecInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainDnssecInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainDnssecInfo(request)
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

// DescribeDomainDnssecInfoRequest is the request struct for api DescribeDomainDnssecInfo
type DescribeDomainDnssecInfoRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeDomainDnssecInfoResponse is the response struct for api DescribeDomainDnssecInfo
type DescribeDomainDnssecInfoResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainName string `json:"DomainName" xml:"DomainName"`
	Status     string `json:"Status" xml:"Status"`
	DsRecord   string `json:"DsRecord" xml:"DsRecord"`
	Digest     string `json:"Digest" xml:"Digest"`
	DigestType string `json:"DigestType" xml:"DigestType"`
	Algorithm  string `json:"Algorithm" xml:"Algorithm"`
	PublicKey  string `json:"PublicKey" xml:"PublicKey"`
	KeyTag     string `json:"KeyTag" xml:"KeyTag"`
	Flags      string `json:"Flags" xml:"Flags"`
}

// CreateDescribeDomainDnssecInfoRequest creates a request to invoke DescribeDomainDnssecInfo API
func CreateDescribeDomainDnssecInfoRequest() (request *DescribeDomainDnssecInfoRequest) {
	request = &DescribeDomainDnssecInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainDnssecInfo", "alidns", "openAPI")
	return
}

// CreateDescribeDomainDnssecInfoResponse creates a response to parse from DescribeDomainDnssecInfo response
func CreateDescribeDomainDnssecInfoResponse() (response *DescribeDomainDnssecInfoResponse) {
	response = &DescribeDomainDnssecInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
