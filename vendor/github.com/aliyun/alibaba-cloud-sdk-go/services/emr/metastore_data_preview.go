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

// MetastoreDataPreview invokes the emr.MetastoreDataPreview API synchronously
// api document: https://help.aliyun.com/api/emr/metastoredatapreview.html
func (client *Client) MetastoreDataPreview(request *MetastoreDataPreviewRequest) (response *MetastoreDataPreviewResponse, err error) {
	response = CreateMetastoreDataPreviewResponse()
	err = client.DoAction(request, response)
	return
}

// MetastoreDataPreviewWithChan invokes the emr.MetastoreDataPreview API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredatapreview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDataPreviewWithChan(request *MetastoreDataPreviewRequest) (<-chan *MetastoreDataPreviewResponse, <-chan error) {
	responseChan := make(chan *MetastoreDataPreviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MetastoreDataPreview(request)
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

// MetastoreDataPreviewWithCallback invokes the emr.MetastoreDataPreview API asynchronously
// api document: https://help.aliyun.com/api/emr/metastoredatapreview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MetastoreDataPreviewWithCallback(request *MetastoreDataPreviewRequest, callback func(response *MetastoreDataPreviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MetastoreDataPreviewResponse
		var err error
		defer close(result)
		response, err = client.MetastoreDataPreview(request)
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

// MetastoreDataPreviewRequest is the request struct for api MetastoreDataPreview
type MetastoreDataPreviewRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DbName          string           `position:"Query" name:"DbName"`
	TableName       string           `position:"Query" name:"TableName"`
}

// MetastoreDataPreviewResponse is the response struct for api MetastoreDataPreview
type MetastoreDataPreviewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Samples   string `json:"Samples" xml:"Samples"`
}

// CreateMetastoreDataPreviewRequest creates a request to invoke MetastoreDataPreview API
func CreateMetastoreDataPreviewRequest() (request *MetastoreDataPreviewRequest) {
	request = &MetastoreDataPreviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "MetastoreDataPreview", "emr", "openAPI")
	return
}

// CreateMetastoreDataPreviewResponse creates a response to parse from MetastoreDataPreview response
func CreateMetastoreDataPreviewResponse() (response *MetastoreDataPreviewResponse) {
	response = &MetastoreDataPreviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}