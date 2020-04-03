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

// CreateAlertContactGroup invokes the arms.CreateAlertContactGroup API synchronously
// api document: https://help.aliyun.com/api/arms/createalertcontactgroup.html
func (client *Client) CreateAlertContactGroup(request *CreateAlertContactGroupRequest) (response *CreateAlertContactGroupResponse, err error) {
	response = CreateCreateAlertContactGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAlertContactGroupWithChan invokes the arms.CreateAlertContactGroup API asynchronously
// api document: https://help.aliyun.com/api/arms/createalertcontactgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlertContactGroupWithChan(request *CreateAlertContactGroupRequest) (<-chan *CreateAlertContactGroupResponse, <-chan error) {
	responseChan := make(chan *CreateAlertContactGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlertContactGroup(request)
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

// CreateAlertContactGroupWithCallback invokes the arms.CreateAlertContactGroup API asynchronously
// api document: https://help.aliyun.com/api/arms/createalertcontactgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAlertContactGroupWithCallback(request *CreateAlertContactGroupRequest, callback func(response *CreateAlertContactGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlertContactGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateAlertContactGroup(request)
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

// CreateAlertContactGroupRequest is the request struct for api CreateAlertContactGroup
type CreateAlertContactGroupRequest struct {
	*requests.RpcRequest
	ContactGroupName string `position:"Query" name:"ContactGroupName"`
	ProxyUserId      string `position:"Query" name:"ProxyUserId"`
	ContactIds       string `position:"Query" name:"ContactIds"`
}

// CreateAlertContactGroupResponse is the response struct for api CreateAlertContactGroup
type CreateAlertContactGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ContactGroupId string `json:"ContactGroupId" xml:"ContactGroupId"`
}

// CreateCreateAlertContactGroupRequest creates a request to invoke CreateAlertContactGroup API
func CreateCreateAlertContactGroupRequest() (request *CreateAlertContactGroupRequest) {
	request = &CreateAlertContactGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateAlertContactGroup", "arms", "openAPI")
	return
}

// CreateCreateAlertContactGroupResponse creates a response to parse from CreateAlertContactGroup response
func CreateCreateAlertContactGroupResponse() (response *CreateAlertContactGroupResponse) {
	response = &CreateAlertContactGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
