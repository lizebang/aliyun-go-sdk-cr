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

// GetRepoListByNamespace invokes the cr.GetRepoListByNamespace API synchronously
// api document: https://help.aliyun.com/api/cr/getrepolistbynamespace.html
func (client *Client) GetRepoListByNamespace(request *GetRepoListByNamespaceRequest) (response *GetRepoListByNamespaceResponse, err error) {
	response = CreateGetRepoListByNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoListByNamespaceWithChan invokes the cr.GetRepoListByNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepolistbynamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoListByNamespaceWithChan(request *GetRepoListByNamespaceRequest) (<-chan *GetRepoListByNamespaceResponse, <-chan error) {
	responseChan := make(chan *GetRepoListByNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepoListByNamespace(request)
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

// GetRepoListByNamespaceWithCallback invokes the cr.GetRepoListByNamespace API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepolistbynamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoListByNamespaceWithCallback(request *GetRepoListByNamespaceRequest, callback func(response *GetRepoListByNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoListByNamespaceResponse
		var err error
		defer close(result)
		response, err = client.GetRepoListByNamespace(request)
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

// GetRepoListByNamespaceRequest is the request struct for api GetRepoListByNamespace
type GetRepoListByNamespaceRequest struct {
	*requests.RoaRequest
	Status         string           `position:"Query" name:"Status"`
	RepoNamespace  string           `position:"Path" name:"RepoNamespace"`
	RepoNamePrefix string           `position:"Query" name:"RepoNamePrefix"`
	Page           requests.Integer `position:"Query" name:"Page"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// GetRepoListByNamespaceResponse is the response struct for api GetRepoListByNamespace
type GetRepoListByNamespaceResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoListByNamespaceRequest creates a request to invoke GetRepoListByNamespace API
func CreateGetRepoListByNamespaceRequest() (request *GetRepoListByNamespaceRequest) {
	request = &GetRepoListByNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "GetRepoListByNamespace", "/repos/[RepoNamespace]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetRepoListByNamespaceResponse creates a response to parse from GetRepoListByNamespace response
func CreateGetRepoListByNamespaceResponse() (response *GetRepoListByNamespaceResponse) {
	response = &GetRepoListByNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
