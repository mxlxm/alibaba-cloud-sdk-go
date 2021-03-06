package imm

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

// GetTagSet invokes the imm.GetTagSet API synchronously
// api document: https://help.aliyun.com/api/imm/gettagset.html
func (client *Client) GetTagSet(request *GetTagSetRequest) (response *GetTagSetResponse, err error) {
	response = CreateGetTagSetResponse()
	err = client.DoAction(request, response)
	return
}

// GetTagSetWithChan invokes the imm.GetTagSet API asynchronously
// api document: https://help.aliyun.com/api/imm/gettagset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTagSetWithChan(request *GetTagSetRequest) (<-chan *GetTagSetResponse, <-chan error) {
	responseChan := make(chan *GetTagSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTagSet(request)
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

// GetTagSetWithCallback invokes the imm.GetTagSet API asynchronously
// api document: https://help.aliyun.com/api/imm/gettagset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTagSetWithCallback(request *GetTagSetRequest, callback func(response *GetTagSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTagSetResponse
		var err error
		defer close(result)
		response, err = client.GetTagSet(request)
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

// GetTagSetRequest is the request struct for api GetTagSet
type GetTagSetRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
}

// GetTagSetResponse is the response struct for api GetTagSet
type GetTagSetResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SetId      string `json:"SetId" xml:"SetId"`
	Status     string `json:"Status" xml:"Status"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	ModifyTime string `json:"ModifyTime" xml:"ModifyTime"`
	Photos     int    `json:"Photos" xml:"Photos"`
}

// CreateGetTagSetRequest creates a request to invoke GetTagSet API
func CreateGetTagSetRequest() (request *GetTagSetRequest) {
	request = &GetTagSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetTagSet", "imm", "openAPI")
	return
}

// CreateGetTagSetResponse creates a response to parse from GetTagSet response
func CreateGetTagSetResponse() (response *GetTagSetResponse) {
	response = &GetTagSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
