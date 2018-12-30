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

// DeleteRepo invokes the cr.DeleteRepo API synchronously
// api document: https://help.aliyun.com/api/cr/deleterepo.html
func (client *Client) DeleteRepo(request *DeleteRepoRequest) (response *DeleteRepoResponse, err error) {
	response = CreateDeleteRepoResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepoWithChan invokes the cr.DeleteRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepoWithChan(request *DeleteRepoRequest) (<-chan *DeleteRepoResponse, <-chan error) {
	responseChan := make(chan *DeleteRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepo(request)
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

// DeleteRepoWithCallback invokes the cr.DeleteRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/deleterepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteRepoWithCallback(request *DeleteRepoRequest, callback func(response *DeleteRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepoResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepo(request)
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

// DeleteRepoRequest is the request struct for api DeleteRepo
type DeleteRepoRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

// DeleteRepoResponse is the response struct for api DeleteRepo
type DeleteRepoResponse struct {
	*responses.BaseResponse
}

// CreateDeleteRepoRequest creates a request to invoke DeleteRepo API
func CreateDeleteRepoRequest() (request *DeleteRepoRequest) {
	request = &DeleteRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Cr", "2016-06-07", "DeleteRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteRepoResponse creates a response to parse from DeleteRepo response
func CreateDeleteRepoResponse() (response *DeleteRepoResponse) {
	response = &DeleteRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
