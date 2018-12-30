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

// GetRepoList invokes the cr.GetRepoList API synchronously
// api document: https://help.aliyun.com/api/cr/getrepolist.html
func (client *Client) GetRepoList(request *GetRepoListRequest) (response *GetRepoListResponse, err error) {
	response = CreateGetRepoListResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoListWithChan invokes the cr.GetRepoList API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoListWithChan(request *GetRepoListRequest) (<-chan *GetRepoListResponse, <-chan error) {
	responseChan := make(chan *GetRepoListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoList(request)
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

// GetRepoListWithCallback invokes the cr.GetRepoList API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoListWithCallback(request *GetRepoListRequest, callback func(response *GetRepoListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoListResponse
		var err error
		defer close(result)
		response, err = client.GetRepoList(request)
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

// GetRepoListRequest is the request struct for api GetRepoList
type GetRepoListRequest struct {
	*requests.RoaRequest
	Status         string           `position:"Query" name:"Status"`
	RepoNamePrefix string           `position:"Query" name:"RepoNamePrefix"`
	Page           requests.Integer `position:"Query" name:"Page"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// GetRepoListResponse is the response struct for api GetRepoList
type GetRepoListResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoListRequest creates a request to invoke GetRepoList API
func CreateGetRepoListRequest() (request *GetRepoListRequest) {
	request = &GetRepoListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "GetRepoList", "/repos", "", "")
	request.Method = requests.GET
	return
}

// CreateGetRepoListResponse creates a response to parse from GetRepoList response
func CreateGetRepoListResponse() (response *GetRepoListResponse) {
	response = &GetRepoListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
