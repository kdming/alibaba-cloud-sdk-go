package arms

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

// StopAlert invokes the arms.StopAlert API synchronously
// api document: https://help.aliyun.com/api/arms/stopalert.html
func (client *Client) StopAlert(request *StopAlertRequest) (response *StopAlertResponse, err error) {
	response = CreateStopAlertResponse()
	err = client.DoAction(request, response)
	return
}

// StopAlertWithChan invokes the arms.StopAlert API asynchronously
// api document: https://help.aliyun.com/api/arms/stopalert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopAlertWithChan(request *StopAlertRequest) (<-chan *StopAlertResponse, <-chan error) {
	responseChan := make(chan *StopAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopAlert(request)
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

// StopAlertWithCallback invokes the arms.StopAlert API asynchronously
// api document: https://help.aliyun.com/api/arms/stopalert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopAlertWithCallback(request *StopAlertRequest, callback func(response *StopAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopAlertResponse
		var err error
		defer close(result)
		response, err = client.StopAlert(request)
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

// StopAlertRequest is the request struct for api StopAlert
type StopAlertRequest struct {
	*requests.RpcRequest
	AlertId     string `position:"Query" name:"AlertId"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
}

// StopAlertResponse is the response struct for api StopAlert
type StopAlertResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	StopAlertIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateStopAlertRequest creates a request to invoke StopAlert API
func CreateStopAlertRequest() (request *StopAlertRequest) {
	request = &StopAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "StopAlert", "arms", "openAPI")
	return
}

// CreateStopAlertResponse creates a response to parse from StopAlert response
func CreateStopAlertResponse() (response *StopAlertResponse) {
	response = &StopAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
