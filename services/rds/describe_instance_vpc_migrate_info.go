package rds

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

// DescribeInstanceVpcMigrateInfo invokes the rds.DescribeInstanceVpcMigrateInfo API synchronously
// api document: https://help.aliyun.com/api/rds/describeinstancevpcmigrateinfo.html
func (client *Client) DescribeInstanceVpcMigrateInfo(request *DescribeInstanceVpcMigrateInfoRequest) (response *DescribeInstanceVpcMigrateInfoResponse, err error) {
	response = CreateDescribeInstanceVpcMigrateInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceVpcMigrateInfoWithChan invokes the rds.DescribeInstanceVpcMigrateInfo API asynchronously
// api document: https://help.aliyun.com/api/rds/describeinstancevpcmigrateinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceVpcMigrateInfoWithChan(request *DescribeInstanceVpcMigrateInfoRequest) (<-chan *DescribeInstanceVpcMigrateInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceVpcMigrateInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceVpcMigrateInfo(request)
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

// DescribeInstanceVpcMigrateInfoWithCallback invokes the rds.DescribeInstanceVpcMigrateInfo API asynchronously
// api document: https://help.aliyun.com/api/rds/describeinstancevpcmigrateinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceVpcMigrateInfoWithCallback(request *DescribeInstanceVpcMigrateInfoRequest, callback func(response *DescribeInstanceVpcMigrateInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceVpcMigrateInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceVpcMigrateInfo(request)
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

// DescribeInstanceVpcMigrateInfoRequest is the request struct for api DescribeInstanceVpcMigrateInfo
type DescribeInstanceVpcMigrateInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeInstanceVpcMigrateInfoResponse is the response struct for api DescribeInstanceVpcMigrateInfo
type DescribeInstanceVpcMigrateInfoResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	DBInstanceMigrateInfos DBInstanceMigrateInfos `json:"DBInstanceMigrateInfos" xml:"DBInstanceMigrateInfos"`
}

// CreateDescribeInstanceVpcMigrateInfoRequest creates a request to invoke DescribeInstanceVpcMigrateInfo API
func CreateDescribeInstanceVpcMigrateInfoRequest() (request *DescribeInstanceVpcMigrateInfoRequest) {
	request = &DescribeInstanceVpcMigrateInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceVpcMigrateInfo", "Rds", "openAPI")
	return
}

// CreateDescribeInstanceVpcMigrateInfoResponse creates a response to parse from DescribeInstanceVpcMigrateInfo response
func CreateDescribeInstanceVpcMigrateInfoResponse() (response *DescribeInstanceVpcMigrateInfoResponse) {
	response = &DescribeInstanceVpcMigrateInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}