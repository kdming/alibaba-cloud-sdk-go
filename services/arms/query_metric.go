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

// QueryMetric invokes the arms.QueryMetric API synchronously
// api document: https://help.aliyun.com/api/arms/querymetric.html
func (client *Client) QueryMetric(request *QueryMetricRequest) (response *QueryMetricResponse, err error) {
	response = CreateQueryMetricResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMetricWithChan invokes the arms.QueryMetric API asynchronously
// api document: https://help.aliyun.com/api/arms/querymetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMetricWithChan(request *QueryMetricRequest) (<-chan *QueryMetricResponse, <-chan error) {
	responseChan := make(chan *QueryMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMetric(request)
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

// QueryMetricWithCallback invokes the arms.QueryMetric API asynchronously
// api document: https://help.aliyun.com/api/arms/querymetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMetricWithCallback(request *QueryMetricRequest, callback func(response *QueryMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMetricResponse
		var err error
		defer close(result)
		response, err = client.QueryMetric(request)
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

// QueryMetricRequest is the request struct for api QueryMetric
type QueryMetricRequest struct {
	*requests.RpcRequest
	EndTime       requests.Integer      `position:"Query" name:"EndTime"`
	OrderBy       string                `position:"Query" name:"OrderBy"`
	StartTime     requests.Integer      `position:"Query" name:"StartTime"`
	Filters       *[]QueryMetricFilters `position:"Query" name:"Filters"  type:"Repeated"`
	ProxyUserId   string                `position:"Query" name:"ProxyUserId"`
	Measures      *[]string             `position:"Query" name:"Measures"  type:"Repeated"`
	IntervalInSec requests.Integer      `position:"Query" name:"IntervalInSec"`
	Metric        string                `position:"Query" name:"Metric"`
	Limit         requests.Integer      `position:"Query" name:"Limit"`
	Dimensions    *[]string             `position:"Query" name:"Dimensions"  type:"Repeated"`
	Order         string                `position:"Query" name:"Order"`
}

// QueryMetricFilters is a repeated param struct in QueryMetricRequest
type QueryMetricFilters struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// QueryMetricResponse is the response struct for api QueryMetric
type QueryMetricResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateQueryMetricRequest creates a request to invoke QueryMetric API
func CreateQueryMetricRequest() (request *QueryMetricRequest) {
	request = &QueryMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "QueryMetric", "arms", "openAPI")
	return
}

// CreateQueryMetricResponse creates a response to parse from QueryMetric response
func CreateQueryMetricResponse() (response *QueryMetricResponse) {
	response = &QueryMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
