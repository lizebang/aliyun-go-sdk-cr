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

// GetRepoTags invokes the cr.GetRepoTags API synchronously
// api document: https://help.aliyun.com/api/cr/getrepotags.html
func (client *Client) GetRepoTags(request *GetRepoTagsRequest) (response *GetRepoTagsResponse, err error) {
	response = CreateGetRepoTagsResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoTagsWithChan invokes the cr.GetRepoTags API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagsWithChan(request *GetRepoTagsRequest) (<-chan *GetRepoTagsResponse, <-chan error) {
	responseChan := make(chan *GetRepoTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoTags(request)
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

// GetRepoTagsWithCallback invokes the cr.GetRepoTags API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepotags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoTagsWithCallback(request *GetRepoTagsRequest, callback func(response *GetRepoTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoTagsResponse
		var err error
		defer close(result)
		response, err = client.GetRepoTags(request)
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

// GetRepoTagsRequest is the request struct for api GetRepoTags
type GetRepoTagsRequest struct {
	*requests.RoaRequest
	RepoNamespace string           `position:"Path" name:"RepoNamespace"`
	RepoName      string           `position:"Path" name:"RepoName"`
	Page          requests.Integer `position:"Query" name:"Page"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// GetRepoTagsResponse is the response struct for api GetRepoTags
type GetRepoTagsResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoTagsRequest creates a request to invoke GetRepoTags API
func CreateGetRepoTagsRequest() (request *GetRepoTagsRequest) {
	request = &GetRepoTagsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "GetRepoTags", "/repos/[RepoNamespace]/[RepoName]/tags", "", "")
	request.Method = requests.GET
	return
}

// CreateGetRepoTagsResponse creates a response to parse from GetRepoTags response
func CreateGetRepoTagsResponse() (response *GetRepoTagsResponse) {
	response = &GetRepoTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
