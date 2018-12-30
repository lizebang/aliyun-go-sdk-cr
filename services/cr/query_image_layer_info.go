package cr

// MIT License

// Copyright (c) 2018 Li Zebang

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryImageLayerInfo invokes the cr.QueryImageLayerInfo API synchronously
// api document: https://help.aliyun.com/api/cr/query-image-layer-info.html
func (client *Client) QueryImageLayerInfo(request *QueryImageLayerInfoRequest) (response *QueryImageLayerInfoResponse, err error) {
	response = CreateQueryImageLayerInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryImageLayerInfoWithChan invokes the cr.QueryImageLayerInfo API asynchronously
// api document: https://help.aliyun.com/api/cr/query-image-layer-info.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryImageLayerInfoWithChan(request *QueryImageLayerInfoRequest) (<-chan *QueryImageLayerInfoResponse, <-chan error) {
	responseChan := make(chan *QueryImageLayerInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryImageLayerInfo(request)
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

// QueryImageLayerInfoWithCallback invokes the cr.QueryImageLayerInfo API asynchronously
// api document: https://help.aliyun.com/api/cr/query-image-layer-info.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryImageLayerInfoWithCallback(request *QueryImageLayerInfoRequest, callback func(response *QueryImageLayerInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryImageLayerInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryImageLayerInfo(request)
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

// QueryImageLayerInfoRequest is the request struct for api QueryImageLayerInfo
type QueryImageLayerInfoRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Tag           string `position:"Path" name:"Tag"`
}

// QueryImageLayerInfoResponse is the response struct for api QueryImageLayerInfo
type QueryImageLayerInfoResponse struct {
	*responses.BaseResponse
}

// CreateQueryImageLayerInfoRequest creates a request to invoke QueryImageLayerInfo API
func CreateQueryImageLayerInfoRequest() (request *QueryImageLayerInfoRequest) {
	request = &QueryImageLayerInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "QueryImageLayerInfo", "/repos/[RepoNamespace]/[RepoName]/tags/[Tag]/layers", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryImageLayerInfoResponse creates a response to parse from QueryImageLayerInfo response
func CreateQueryImageLayerInfoResponse() (response *QueryImageLayerInfoResponse) {
	response = &QueryImageLayerInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
