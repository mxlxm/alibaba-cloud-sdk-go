package ecs

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

func (client *Client) AddIpRange(request *AddIpRangeRequest) (response *AddIpRangeResponse, err error) {
	response = CreateAddIpRangeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddIpRangeWithChan(request *AddIpRangeRequest) (<-chan *AddIpRangeResponse, <-chan error) {
	responseChan := make(chan *AddIpRangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddIpRange(request)
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

func (client *Client) AddIpRangeWithCallback(request *AddIpRangeRequest, callback func(response *AddIpRangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddIpRangeResponse
		var err error
		defer close(result)
		response, err = client.AddIpRange(request)
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

type AddIpRangeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpAddress            string `position:"Query" name:"IpAddress"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type AddIpRangeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAddIpRangeRequest() (request *AddIpRangeRequest) {
	request = &AddIpRangeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AddIpRange", "", "")
	return
}

func CreateAddIpRangeResponse() (response *AddIpRangeResponse) {
	response = &AddIpRangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}