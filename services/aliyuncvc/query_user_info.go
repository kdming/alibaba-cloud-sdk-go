package aliyuncvc

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

// QueryUserInfo invokes the aliyuncvc.QueryUserInfo API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/queryuserinfo.html
func (client *Client) QueryUserInfo(request *QueryUserInfoRequest) (response *QueryUserInfoResponse, err error) {
	response = CreateQueryUserInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryUserInfoWithChan invokes the aliyuncvc.QueryUserInfo API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/queryuserinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUserInfoWithChan(request *QueryUserInfoRequest) (<-chan *QueryUserInfoResponse, <-chan error) {
	responseChan := make(chan *QueryUserInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUserInfo(request)
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

// QueryUserInfoWithCallback invokes the aliyuncvc.QueryUserInfo API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/queryuserinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUserInfoWithCallback(request *QueryUserInfoRequest, callback func(response *QueryUserInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUserInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryUserInfo(request)
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

// QueryUserInfoRequest is the request struct for api QueryUserInfo
type QueryUserInfoRequest struct {
	*requests.RpcRequest
}

// QueryUserInfoResponse is the response struct for api QueryUserInfo
type QueryUserInfoResponse struct {
	*responses.BaseResponse
	ErrorCode int      `json:"ErrorCode" xml:"ErrorCode"`
	Message   string   `json:"Message" xml:"Message"`
	Success   bool     `json:"Success" xml:"Success"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	UserInfo  UserInfo `json:"UserInfo" xml:"UserInfo"`
}

// CreateQueryUserInfoRequest creates a request to invoke QueryUserInfo API
func CreateQueryUserInfoRequest() (request *QueryUserInfoRequest) {
	request = &QueryUserInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-09-19", "QueryUserInfo", "aliyuncvc", "openAPI")
	return
}

// CreateQueryUserInfoResponse creates a response to parse from QueryUserInfo response
func CreateQueryUserInfoResponse() (response *QueryUserInfoResponse) {
	response = &QueryUserInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
