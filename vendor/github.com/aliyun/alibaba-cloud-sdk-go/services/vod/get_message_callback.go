package vod

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

// GetMessageCallback invokes the vod.GetMessageCallback API synchronously
// api document: https://help.aliyun.com/api/vod/getmessagecallback.html
func (client *Client) GetMessageCallback(request *GetMessageCallbackRequest) (response *GetMessageCallbackResponse, err error) {
	response = CreateGetMessageCallbackResponse()
	err = client.DoAction(request, response)
	return
}

// GetMessageCallbackWithChan invokes the vod.GetMessageCallback API asynchronously
// api document: https://help.aliyun.com/api/vod/getmessagecallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMessageCallbackWithChan(request *GetMessageCallbackRequest) (<-chan *GetMessageCallbackResponse, <-chan error) {
	responseChan := make(chan *GetMessageCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMessageCallback(request)
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

// GetMessageCallbackWithCallback invokes the vod.GetMessageCallback API asynchronously
// api document: https://help.aliyun.com/api/vod/getmessagecallback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMessageCallbackWithCallback(request *GetMessageCallbackRequest, callback func(response *GetMessageCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMessageCallbackResponse
		var err error
		defer close(result)
		response, err = client.GetMessageCallback(request)
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

// GetMessageCallbackRequest is the request struct for api GetMessageCallback
type GetMessageCallbackRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	ResourceRealOwnerId  requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	AppId                string           `position:"Query" name:"AppId"`
}

// GetMessageCallbackResponse is the response struct for api GetMessageCallback
type GetMessageCallbackResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	MessageCallback MessageCallback `json:"MessageCallback" xml:"MessageCallback"`
}

// CreateGetMessageCallbackRequest creates a request to invoke GetMessageCallback API
func CreateGetMessageCallbackRequest() (request *GetMessageCallbackRequest) {
	request = &GetMessageCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetMessageCallback", "vod", "openAPI")
	return
}

// CreateGetMessageCallbackResponse creates a response to parse from GetMessageCallback response
func CreateGetMessageCallbackResponse() (response *GetMessageCallbackResponse) {
	response = &GetMessageCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}